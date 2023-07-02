package natsrouter

import (
	"encoding/json"
	"errors"
	uc "git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/models/rechan"
	"github.com/nats-io/nats.go"
)

type RecMessage struct {
	From    string
	To      string
	Payload []byte
}

type DecodePayload struct {
	Action string      //serviceName.objectName
	Method string      //POST
	Body   interface{} //json
}

func NewRouter(reChan chan rechan.Nats, msg *nats.Msg, usecase uc.Repo) {
	//получаем из сообщения данные в виде json и преобразуем в структуру
	message := RecMessage{
		From:    msg.Subject,
		To:      msg.Reply,
		Payload: msg.Data,
	}
	//payload в структуру
	decodePayload := DecodePayload{}
	err := json.Unmarshal(message.Payload, &decodePayload)
	if err != nil {
		reChan <- rechan.Nats{
			Error: err,
		}
	}

	switch decodePayload.Action {
	case "Template.subject":
		newExampleRoutes(reChan, decodePayload, usecase)
	default:
		err = errors.New("action not found")
		reChan <- rechan.Nats{
			Error: err,
		}
	}
}
