package usecase

import (
	"context"
	"log"
	"strconv"

	"p4/internal/repository"
	"p4/pkg/modules"

	"github.com/redis/go-redis/v9"
)

type UserUseCase struct {
	repo  repository.UserRepository
	redis *redis.Client
}

func NewUserUseCase(r *repository.Repositories, redis *redis.Client) *UserUseCase {
	return &UserUseCase{
		repo:  r,
		redis: redis,
	}
}

func (u *UserUseCase) GetUsers() ([]modules.User, error) {
	users, err := u.repo.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserUseCase) NewUser(newUser modules.User) (int, error) {
	id, err := u.repo.NewUser(newUser)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (u *UserUseCase) DeleteUser(id int) (*modules.User, error) {
	deletedUser, err := u.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	return deletedUser, nil
}

func (u *UserUseCase) GetUserByid(id int) (*modules.User, error) {
	r := u.redis
	key := "user" + strconv.Itoa(id)
	ctx := context.Background()

	var cacheUser modules.User
	if err := r.HGetAll(ctx, key).Scan(&cacheUser); err != nil {
		log.Println(err)
	}

	if cacheUser.Id == id {
		log.Println("from redis cache")
		return &cacheUser, nil
	}

	userwithID, err := u.repo.GetUserByid(id)
	if err != nil {
		return nil, err
	}
	if _, err := r.Pipelined(ctx, func(p redis.Pipeliner) error {
		r.HSet(ctx, key, "id", userwithID.Id)
		r.HSet(ctx, key, "name", userwithID.Name)
		r.HSet(ctx, key, "age", userwithID.Age)
		r.HSet(ctx, key, "hobby", userwithID.Hobby)
		r.HSet(ctx, key, "profession", userwithID.Profession)
		return nil
	}); err != nil {
		log.Println(err)
	}
	log.Println("set to redis cache")

	return userwithID, nil
}

func (u *UserUseCase) UpdateUser(id int, userToUpdate modules.User) (*modules.User, error) {
	updatedUser, err := u.repo.UpdateUser(id, userToUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
