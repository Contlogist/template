package rechan

import "git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/models"

type Multiple struct {
	Data   interface{}
	Errors models.ErrorMultiple
}

func (*Multiple) SendError(reChan chan Multiple, errorMultiple models.ErrorMultiple) {
	if reChan == nil {
		return
	}
	select {
	case _, ok := <-reChan:
		if !ok {
			return
		}
	default:
		if len(errorMultiple.CriticalError) > 0 {
			reChan <- Multiple{Errors: errorMultiple}
			return
		} else {
			return
		}
	}
}

func (*Multiple) SendData(reChan chan Multiple, data interface{}) {
	if reChan == nil {
		return
	}
	select {
	case _, ok := <-reChan:
		if !ok {
			return
		}
	default:
		reChan <- Multiple{Data: data}
	}
}
