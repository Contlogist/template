package natsconnect

import (
	"context"
	uc "git.legchelife.ru/gitlab-instance-7d441567/catalog_m/internal/usecase"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

// Server -.
type Server struct {
	Js       nats.JetStreamContext
	Messages chan *nats.Msg
	Subj     string
}

func Connect(url string) (nats.JetStreamContext, error) {
	nc, _ := nats.Connect(url)
	js, _ := nc.JetStream()
	i, err := js.AddStream(&nats.StreamConfig{
		Name:     "Template",
		Subjects: []string{"Template.*"},
		MaxBytes: 1024,
	})
	if err != nil {
		logrus.Error("AddStream: ", err)
	} else {
		logrus.Info("AddStream: ", i)
	}

	return js, nil
}

func New(js nats.JetStreamContext, usecase uc.Repo) {
	s := &Server{
		Js:   js,
		Subj: "Template",
	}
	s.start(usecase)
}

func (s *Server) start(usecase uc.Repo) {
	// создаем подписку на сообщения для адресата userServer

	go func() {
		sub, err := s.Js.SubscribeSync("Template")
		if err != nil {
			logrus.Error("SubscribeSync: ", err)
		}
		ctx := context.Background()
		for {
			select {
			case <-ctx.Done():
				// контекст отменен, выходим из цикла
				return
			default:
				// получение нового сообщения
				msg, err := sub.NextMsgWithContext(ctx)
				if err != nil {
					logrus.Error("NextMsgWithContext: ", err)
				}

				// обработка полученного сообщения
				logrus.Info("Получили сообщение: ", string(msg.Data))

				// отправка подтверждения обработки сообщения
				msg.Ack()
				count, _ := s.Js.StreamInfo("Template")
				logrus.Info("StreamInfo: ", count.State.Msgs)
			}
		}
	}()
}
