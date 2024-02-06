package test_test

import (
	"errors"
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/internal/repo/user/mock"
	"git.legchelife.ru/root/template/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParamGetList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	paramRepo := user_mock.NewMockIParam(ctl)
	mocResp := []user.Param{
		{
			ID:     1,
			UserID: 1,
			Name:   "test",
		},
		{
			ID:     2,
			UserID: 1,
			Name:   "test2",
		},
	}
	paramRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(mocResp, nil).Times(1)

	fakeRepo := user.NewFake(nil, paramRepo, nil)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	in := 1

	resp, err := useCase.UserRepo.Param.GetList(nil, in)
	//проверка на ошибку
	if err != nil {
		require.NoError(t, err)
	}
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}

func TestParamGetListError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	paramRepo := user_mock.NewMockIParam(ctl)
	var mocResp []user.Param
	mocErr := errors.New("test error")
	paramRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	fakeRepo := user.NewFake(nil, paramRepo, nil)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	in := 1

	resp, err := useCase.UserRepo.Param.GetList(nil, in)
	//проверка на ошибку
	if err != nil {
		require.Error(t, err)
	}
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}
