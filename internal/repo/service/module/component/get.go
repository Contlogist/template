package service_module_component

import (
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/rechan"
	"github.com/sirupsen/logrus"
)

func (s *Component) GetTest(ctx context.Base) (bool, error) {
	ctx.SetTimeout(2)
	reChan := make(chan rechan.Base, 0)
	defer close(reChan)

	go func() {
		c := rechan.Base{}

		res, err := s.db.User.Create().
			SetAddress("test").
			SetName("test").
			SetAge(1).
			Save(ctx.Context)

		if err != nil {
			c.SendError(reChan, "AutoUpdateAuthToken", err)
		}

		logrus.Info(res.ID)

		c.SendData(reChan, true)
	}()

	select {
	case <-ctx.Context.Done():
		return false, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return false, re.Error
		} else {
			return re.Data.(bool), nil
		}
	}
}
