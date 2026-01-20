package utils

import "math"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type Pagination struct {
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func SuccessWithPagination(
	data interface{},
	page, limit int,
	total int64,
	message string,
) Response {

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	meta := Pagination{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: totalPage,
	}

	return Response{
		Success: true,
		Data:    data,
		Meta:    meta,
		Message: message,
	}

}

func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}
