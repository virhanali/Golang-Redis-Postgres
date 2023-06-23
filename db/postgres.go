package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var postgresDB *sqlx.DB
var redisClient *redis.Client

func InitPostgres() {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalln("Failed to connect to PostgreSQL:", err)
	}

	postgresDB = db
}

func GetPostgresDB() *sqlx.DB {
	return postgresDB
}

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("Failed to connect to redis:", err)
	}
}

func GetRedisDB() *redis.Client {
	return redisClient
}
