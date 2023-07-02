package rechan

import (
	"github.com/sirupsen/logrus"
)

type Base struct {
	Data  interface{}
	Error error
}

func (*Base) SendError(reChan chan Base, title string, err error) {
	if reChan == nil {
		return
	}
	logrus.Error("SendError - ", title, ": ", err.Error())
	select {
	case _, ok := <-reChan:
		if !ok {
			return
		}
	default:
		reChan <- Base{Error: err}
	}
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
