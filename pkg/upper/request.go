package upper

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/rechan"
)

// DoRequest - метод обработки запросов.
// Принимает дженерик типа ответа, контекст и функцию запроса которую запускает в горутине.
// В основном используется для запросов к БД.
func DoRequest[T any](ctx *context.Base, requestFunction func() (T, error)) (*T, error) {
	reChan := make(chan rechan.Base)
	defer close(reChan)
	go handleRequest(reChan, requestFunction)
	return handleResponse[T](ctx, reChan)
}

// handleRequest функция обработки запроса, принимает канал для ответа и функцию запроса.
func handleRequest[T any](reChan chan rechan.Base, requestFunction func() (T, error)) {
	c := rechan.Base{}
	data, err := requestFunction()
	if err != nil {
		c.SendError(reChan, "Error: ", err)
	} else {
		c.SendData(reChan, data)
	}
}

// handleResponse функция обработки ответа, принимает контекст и канал для ответа.
func handleResponse[T any](ctx *context.Base, reChan chan rechan.Base) (*T, error) {
	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			r := re.Data.(T)
			return &r, nil
		}
	}
}
