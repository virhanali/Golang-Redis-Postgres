package response

type BasePayload struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

type MetaTpl struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalData int `json:"total_data"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalData   int `json:"total_data"`
}
