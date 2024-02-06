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

func TestUserGet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := 1
	mocResp := &user.User{
		ID:           1,
		Name:         "Roman",
		Email:        "r.abramov@contlogist.ru",
		CompanyID:    1,
		Password:     "123456",
		RefreshToken: nil,
	}
	userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mocResp, nil).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Get(nil, in)
	//проверка на ошибку
	if err != nil {
		require.NoError(t, err)
	}
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}

func TestUserGetError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := 1
	mocResp := &user.User{}
	mocErr := errors.New("test error")
	userRepo.EXPECT().Get(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Get(nil, in)
	//проверка на ошибку
	if err != nil {
		require.Error(t, err)
	}
	//проверка на несовпадение
	require.NotEqual(t, mocResp, resp)
}

func TestUserGetList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := user.UserFilter{}
	mocResp := []user.User{
		{
			ID:           1,
			Name:         "Roman",
			Email:        "r.abramov@contlogist.ru",
			CompanyID:    1,
			Password:     "123456",
			RefreshToken: nil,
		},
		{
			ID:           2,
			Name:         "Ekaterina",
			Email:        "e.abramova@contlogist.ru",
			CompanyID:    1,
			Password:     "123456",
			RefreshToken: nil,
		},
	}
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(mocResp, nil).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.GetList(nil, in)
	//проверка на ошибку
	if err != nil {
		require.NoError(t, err)
	}
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}

func TestUserGetListError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := user.UserFilter{}
	var mocResp []user.User
	mocErr := errors.New("test error")
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.GetList(nil, in)
	//проверка на ошибку
	if err != nil {
		require.Error(t, err)
	}
	//проверка на несовпадение
	require.Equal(t, mocResp, resp)
}

func TestUserPost(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := user.User{
		ID:           1,
		Name:         "Roman",
		Email:        "r.abramov@contlogist.ru",
		CompanyID:    1,
		Password:     "123456",
		RefreshToken: nil,
	}
	mocResp := 1
	userRepo.EXPECT().Post(gomock.Any(), gomock.Any()).Return(&mocResp, nil).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Post(nil, in)
	//проверка на ошибку
	if err != nil {
		require.NoError(t, err)
	}
	//проверка на совпадение
	require.Equal(t, &mocResp, resp)
}

func TestUserPostError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := user.User{}
	var mocResp *int = nil
	mocErr := errors.New("test error")
	userRepo.EXPECT().Post(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Post(nil, in)
	//проверка на ошибку
	require.Error(t, err)
	//проверка на ожидание
	require.EqualValues(t, mocResp, resp)
}

func TestUserPut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := &user.User{
		ID:           1,
		Name:         "Roman",
		Email:        "r.abramov@contlogist.ru",
		CompanyID:    1,
		Password:     "123456",
		RefreshToken: nil,
	}

	mocResp := true
	userRepo.EXPECT().Put(gomock.Any(), gomock.Any()).Return(mocResp, nil).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Put(nil, in)
	//проверка на ошибку
	require.NoError(t, err)
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}

func TestUserPutError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := &user.User{}
	mocResp := false
	mocErr := errors.New("test error")
	userRepo.EXPECT().Put(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Put(nil, in)
	//проверка на ошибку
	require.Error(t, err)
	//проверка на ожидание
	require.Equal(t, mocResp, resp)
}

func TestUserDelete(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := 1
	mocResp := true
	userRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(mocResp, nil).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Delete(nil, in)
	//проверка на ошибку
	require.NoError(t, err)
	//проверка на совпадение
	require.Equal(t, mocResp, resp)
}

func TestUserDeleteError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	in := 1
	mocResp := false
	mocErr := errors.New("test error")
	userRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(mocResp, mocErr).Times(1)

	userFakeRepo := user.NewFake(userRepo, nil, nil)
	rp := repo.NewFake(userFakeRepo, nil, nil)
	useCase := uc.New(rp)

	resp, err := useCase.UserRepo.User.Delete(nil, in)
	//проверка на ошибку
	require.Error(t, err)
	//проверка на ожидание
	require.Equal(t, mocResp, resp)
}
