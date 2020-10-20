package repository

import (
	"fmt"
	"gopkg.in/reform.v1"
	"vue_back/blog/model"
)

type UserRepository struct {
	reform *reform.DB
}

func NewUserRepository(reform *reform.DB) *UserRepository {
	return &UserRepository{reform: reform}
}

func (u *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := model.User{}
	err := u.reform.FindOneTo(&user, "email", email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}
