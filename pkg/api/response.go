package api

type PaginationMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type Paginated struct {
	Items interface{}    `json:"items"`
	Meta  PaginationMeta `json:"meta"`
}

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"error,omitempty"`
}

func Success(data any, message string) *APIResponse {
	return &APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func PaginatedSuccess(items any, limit, offset, total int, message string) *APIResponse {
	return &APIResponse{
		Status:  "success",
		Message: message,
		Data: Paginated{
			Items: items,
			Meta: PaginationMeta{
				Limit:  limit,
				Offset: offset,
				Total:  total,
			},
		},
	}
}

func Error(message string, errors error) *APIResponse {
	return &APIResponse{
		Status:  "error",
		Message: message,
		Errors:  errors.Error(),
	}
}
