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

	user := &repository.User{ID: 1, Name: "Test 1"}
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

	user := &repository.User{ID: 2, Name: "Test 2"}
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.CreateUser(user)
	assert.NoError(t, err)
}

func TestRegisterUserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 3, Name: "Test 3"}
	mockRepo.EXPECT().GetByEmail("3@test.com").Return(user, nil)

	err := userService.RegisterUser(user, "3@test.com")
	assert.Error(t, err)
}

func TestRegisterUserNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 4, Name: "Test 4"}
	mockRepo.EXPECT().GetByEmail("4@test.com").Return(nil, nil)
	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := userService.RegisterUser(user, "4@test.com")
	assert.NoError(t, err)
}

func TestRegisterUserRepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 5, Name: "Test 5"}
	mockRepo.EXPECT().GetByEmail("5@test.com").Return(nil, assert.AnError)

	err := userService.RegisterUser(user, "5@test.com")
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

func TestUpdateUsernameNotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	mockRepo.EXPECT().GetUserByID(6).Return(nil, assert.AnError)

	err := userService.UpdateUserName(6, "Test 6 new")
	assert.Error(t, err)
}

func TestUpdateUsernameFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 7, Name: "Test 7"}
	mockRepo.EXPECT().GetUserByID(7).Return(user, nil)
	mockRepo.EXPECT().UpdateUser(user).Return(assert.AnError)

	err := userService.UpdateUserName(7, "Test 7 new")
	assert.Error(t, err)
}

func TestUpdateUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	user := &repository.User{ID: 6, Name: "Test 6"}
	mockRepo.EXPECT().GetUserByID(6).Return(user, nil)
	fmt.Println("\nUser before update:", user)

	mockRepo.EXPECT().UpdateUser(user).Return(nil)

	err := userService.UpdateUserName(6, "Test 6 new")
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

	user := &repository.User{ID: 19, Name: "Test 19"}
	mockRepo.EXPECT().GetUserByID(19).Return(user, nil)

	mockRepo.EXPECT().DeleteUser(19).Return(nil)
	mockRepo.EXPECT().GetUserByID(19).Return(nil, nil)

	beforeDelete, err := userService.GetUserByID(19)
	fmt.Println("\nUser before delete:", beforeDelete)

	err = userService.DeleteUser(19)
	result, err := userService.GetUserByID(19)
	fmt.Print("User after delete:", result, "\n\n")

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
