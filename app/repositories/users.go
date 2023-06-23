package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"ginredis/app/entities"
	"ginredis/response"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type users struct {
	postgres *sqlx.DB
	redis    *redis.Client
}

func NewListRepository(db *sqlx.DB, redis *redis.Client) IUser {
	return &users{
		postgres: db,
		redis:    redis,
	}
}

type UserFilter struct {
	Page  int
	Limit int

	Name string
}

func (r *users) ListUsers(ctx context.Context, filter UserFilter) ([]entities.Users, *response.MetaTpl, error) {
	var (
		conditions []string
		values     []interface{}
	)

	query := `SELECT COUNT(u.id) OVER() as total, 
			u.id, u.name, u.password, u.email	
			FROM users u`

	if len(filter.Name) > 0 {
		conditions = append(conditions, "u.name ILIKE ?")
		values = append(values, "%"+filter.Name+"%")
	}

	if len(conditions) > 0 && len(values) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(conditions, " AND "))
	}

	query = fmt.Sprintf("%s ORDER BY u.id ASC", query)

	limit := 10
	if filter.Limit > 0 {
		limit = filter.Limit
	}

	page := 1
	if filter.Page > 0 {
		page = filter.Page
	}

	offset := limit * (page - 1)
	query = fmt.Sprintf("%s LIMIT ? OFFSET ? ", query)
	values = append(values, limit)
	values = append(values, offset)

	query = r.postgres.Rebind(query)

	var result []struct {
		Total int `db:"total"`
		entities.Users
	}

	cacheKey := fmt.Sprintf("data users:%d:%d", limit, offset)
	cachedData, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		err = json.Unmarshal([]byte(cachedData), &result)
		if err != nil {
			return nil, nil, err
		}
	} else {
		err := r.postgres.SelectContext(ctx, &result, query, values...)
		if err != nil {
			return nil, nil, err
		}

		// Save the data to Redis cache
		cachedData, err := json.Marshal(result)
		if err == nil {
			err = r.redis.Set(ctx, cacheKey, cachedData, 30*time.Second).Err()
			if err != nil {
				log.Println("Failed to save data to Redis cache:", err)
			}
		}
	}

	users := make([]entities.Users, len(result))
	pagination := new(response.MetaTpl)

	if len(result) == 0 {
		return users, pagination, nil
	}

	for i, v := range result {
		users[i] = v.Users
	}

	pagination.Limit = limit
	pagination.Page = page
	pagination.TotalData = result[0].Total

	return users, pagination, nil
}
