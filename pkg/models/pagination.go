package models

type PaginationRequest struct {
	Page  uint `json:"page" 2:"1"`   // Номер страницы
	Count uint `json:"count" 2:"10"` // Количество элементов на странице
}

type PaginationResponse struct {
	TotalCount uint        `json:"total_count" 2:"1"`         // Общее количество элементов
	TotalPages uint        `json:"total_pages" 2:"1"`         // Общее количество страниц
	Page       uint        `json:"page" 2:"1"`                // Номер страницы
	Count      uint        `json:"count" 2:"10"`              // Количество элементов на странице
	Data       interface{} `json:"data" swaggertype:"object"` // Данные
}
