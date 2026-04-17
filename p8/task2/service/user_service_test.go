package service

import (
	"fmt"
	"p8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Batman"}
	mockRepo.EXPECT().GetUserByID(1).Return(user, nil)

	result, err := userService.GetUserByID(1)
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 1, Name: "Superman"}
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.CreateUser(user)
	assert.NoError(t, err)
}

func TestRegisterUserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 2, Name: "Ironman"}
	mockRepo.EXPECT().GetByEmai("ironman@email.com").Return(user, nil)

	err := userService.RegisterUser(user, "ironman@email.com")
	assert.Error(t, err)
}

func TestRegisterUserNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 7, Name: "Spiderman"}
	mockRepo.EXPECT().GetByEmai("spiderman@email.com").Return(nil, nil)
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.RegisterUser(user, "spiderman@email.com")
	assert.NoError(t, err)
}

func TestRegisterUserRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 11, Name: "Venom"}
	mockRepo.EXPECT().GetByEmai("venom@email.com").Return(nil, assert.AnError)

	err := userService.RegisterUser(user, "venom@email.com")
	assert.Error(t, err)
}

func TestUpdateUsernameEmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	err := userService.UpdateUserName(5, "")
	assert.Error(t, err)
}

func TestUpdateUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 20, Name: "Name"}
	mockRepo.EXPECT().GetUserByID(20).Return(user, nil)
	fmt.Println("\nUser before update:", user)

	mockRepo.EXPECT().UpdateUser(user).Return(nil)

	err := userService.UpdateUserName(20, "NewName")
	fmt.Print("User after update: ", user, "\n\n")
	assert.NoError(t, err)
}

func TestDeleteAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	err := userService.DeleteUser(1)
	assert.Error(t, err)
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().DeleteUser(3).Return(nil)

	err := userService.DeleteUser(3)
	assert.NoError(t, err)
}

func TestDeleteUserRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().DeleteUser(9).Return(assert.AnError)

	err := userService.DeleteUser(9)
	assert.Error(t, err)
}
