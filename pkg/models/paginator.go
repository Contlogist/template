package models

type Paginator struct {
	Limit      int         `json:"limit"`      // Количество элементов на странице
	TotalCount int         `json:"totalCount"` // Общее количество элементов
	Page       int         `json:"page"`       // Номер текущей страницы
	TotalPages int         `json:"totalPages"` // Общее количество страниц
	Items      interface{} `json:"items"`      // Элементы
}
