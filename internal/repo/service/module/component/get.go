package service_module_component

import (
	"git.legchelife.ru/root/template/ent"
	"git.legchelife.ru/root/template/pkg/models/context"
	"git.legchelife.ru/root/template/pkg/models/rechan"
)

func (s *Component) GetTest(ctx context.Base) (*ent.User, error) {
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

		c.SendData(reChan, res)
	}()

	select {
	case <-ctx.Context.Done():
		return nil, ctx.Context.Err()
	case re := <-reChan:
		if re.Error != nil {
			return nil, re.Error
		} else {
			return re.Data.(*ent.User), nil
		}
	}
}
