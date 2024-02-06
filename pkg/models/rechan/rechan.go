package rechan

import (
	"errors"
	"github.com/upper/db/v4"
)

type Base struct {
	Data  interface{}
	Error error
}

func (*Base) SendData(reChan chan Base, data interface{}) {
	if reChan == nil {
		return
	}
	select {
	case _, ok := <-reChan:
		if !ok {
			return
		}
	default:
		reChan <- Base{Data: data}
	}
}

func (*Base) SendError(reChan chan Base, title string, err error) {
	if reChan == nil {
		return
	}
	select {
	case _, ok := <-reChan:
		if !ok {
			return
		}
	default:
		reChan <- Base{Error: errorTranslate(err)}
	}
}

func errorTranslate(err error) error {
	switch {
	case errors.Is(err, db.ErrNoMoreRows):
		err = errors.New("пользователь не найден")
	}

	return err
}
