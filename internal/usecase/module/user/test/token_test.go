package test_test

import (
	"fmt"
	"git.legchelife.ru/root/template/internal/repo"
	"git.legchelife.ru/root/template/internal/repo/user"
	"git.legchelife.ru/root/template/internal/repo/user/mock"
	"git.legchelife.ru/root/template/internal/usecase"
	"git.legchelife.ru/root/template/pkg/models/context"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTokenGet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{
		{
			ID:           1,
			Name:         "Roman",
			Email:        "r.abramov@contlogist.ru",
			CompanyID:    1,
			Password:     "123456",
			RefreshToken: nil,
		},
	}
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, nil).Times(1)
	userRepo.EXPECT().Put(gomock.Any(), gomock.Any()).Return(true, nil).Times(1)

	tokenMockResp := &user.Token{
		Access: user.AccessToken{
			Token: "test",
			Hours: 12,
		},
		Refresh: user.RefreshToken{
			Token: "test",
			Hours: 12,
		},
	}
	tokenRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(tokenMockResp, nil).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.ru"
	inPasword := "123456"

	ginCtx := &gin.Context{}
	ctx := context.Base{}.Create(ginCtx)

	resp, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPasword)
	//проверка на ошибку
	if err != nil {
		require.NoError(t, err)
	}
	fmt.Println(resp)
	//проверка на совпадение
	require.Equal(t, tokenMockResp, resp)
}

func TestTokenGetErrorTokenGet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{
		{
			ID:           1,
			Name:         "Roman",
			Email:        "r.abramov@contlogist.ru",
			CompanyID:    1,
			Password:     "123456",
			RefreshToken: nil,
		},
	}
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, nil).Times(1)

	putErr := fmt.Errorf("ошибка генерации токена")
	tokenRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, putErr).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.tu"
	inPass := "123456"

	ginCtx := &gin.Context{}
	ctx := context.Base{}.Create(ginCtx)

	_, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPass)
	require.EqualError(t, err, "ошибка генерации токена")
}

func TestTokenGetUserError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{}
	returnErr := fmt.Errorf("test error")
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, returnErr).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.ru"
	inPass := "123456"

	ginCtx := &gin.Context{}
	ctx := context.Base{}.Create(ginCtx)

	_, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPass)
	require.Error(t, err)
}

func TestTokenGetErrorUserNotFound(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{}
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, nil).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.ru"
	inPass := "123456"

	ginCtx := &gin.Context{}
	ctx := context.Base{}.Create(ginCtx)

	_, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPass)
	require.EqualError(t, err, "пользователь не найден или неверный пароль")
}

func TestTokenGetErrorManyUsers(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{
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
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, nil).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.ru"
	inPass := "123456"

	ginCtx := &gin.Context{}

	ctx := context.Base{}.Create(ginCtx)

	_, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPass)
	require.EqualError(t, err, "найдено более одного пользователя")
}

func TestTokenGetErrorUserPut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userRepo := user_mock.NewMockIUser(ctl)
	tokenRepo := user_mock.NewMockIToken(ctl)

	userMockResp := []user.User{
		{
			ID:           1,
			Name:         "Roman",
			Email:        "r.abramov@contlogist.ru",
			CompanyID:    1,
			Password:     "123456",
			RefreshToken: nil,
		},
	}
	userRepo.EXPECT().GetList(gomock.Any(), gomock.Any()).Return(userMockResp, nil).Times(1)

	tokenMockResp := &user.Token{
		Access: user.AccessToken{
			Token: "test",
			Hours: 12,
		},
		Refresh: user.RefreshToken{
			Token: "test",
			Hours: 12,
		},
	}
	tokenRepo.EXPECT().Get(gomock.Any(), gomock.Any(), gomock.Any()).Return(tokenMockResp, nil).Times(1)

	putErr := fmt.Errorf("ошибка обновления пользователя")
	userRepo.EXPECT().Put(gomock.Any(), gomock.Any()).Return(false, putErr).Times(1)

	fakeRepo := user.NewFake(userRepo, nil, tokenRepo)
	rp := repo.NewFake(fakeRepo, nil, nil)
	useCase := uc.New(rp)

	inEmail := "r.abramov@contlogist.ru"
	inPass := "123456"

	ginCtx := &gin.Context{}
	ctx := context.Base{}.Create(ginCtx)

	_, err := useCase.UserRepo.Token.Get(&ctx, inEmail, inPass)
	require.EqualError(t, err, "ошибка обновления пользователя")
}
