package models

type ErrorMultiple struct {
	SoftErrors    []string `json:"soft_errors" 2:"error 1, error 2"` // Не критичные ошибки
	CriticalError []string `json:"critical_error" 2:"error 1"`       // Критичные ошибки
}

func (e *ErrorMultiple) AddSoft(err error) {
	e.SoftErrors = append(e.SoftErrors, err.Error())
}

func (e *ErrorMultiple) AddCritical(err error) {
	e.CriticalError = append(e.CriticalError, err.Error())
}
