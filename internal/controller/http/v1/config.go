package v1

import "encoding/json"

//описываем модель ответа

type Request[T any] struct {
	Data T `json:"data"`
}

func (r *Request[T]) FromJSON(data []byte) error {
	return json.Unmarshal(data, r)
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"error"`
}
