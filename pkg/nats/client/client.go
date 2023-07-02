package natsclient

//func GetMessage(conf natsconnect.Config, router func(msg *nats.Msg)) {
//	conn, err := natsconnect.Connect(conf)
//	if err != nil {
//		logrus.Error(err)
//	}
//	defer conn.Close()
//
//	sub, err := conn.Subscribe(conf.Theme, router)
//	if err != nil {
//		logrus.Error(err)
//	}
//	err = sub.AutoUnsubscribe(1)
//	if err != nil {
//		logrus.Error(err)
//	}
//}
