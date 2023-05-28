package server

import (
	"auth/server/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	users []*models.User
}

func NewUserRepository() *UserRepository {

	p1, _ := bcrypt.GenerateFromPassword([]byte("111111111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("123123123"), bcrypt.DefaultCost)

	users := []*models.User{
		{

			ID:       1,
			Name:     "user1",
			Email:    "first@gmail.com",
			Password: string(p1),
		},
		{
			ID:       2,
			Name:     "user2",
			Email:    "second@email.com",
			Password: string(p2),
		},
	}

	return &UserRepository{users: users}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserById(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
