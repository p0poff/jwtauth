package model

import (
	"repo"
)

var userRepo repo.UserManual

type User struct {
	Login string
	Username string
	pass string
}

func getUserRepo(u repo.IUser, login, pass string) (map[string]string, error) {
	return u.GetUser(login, pass)
}

func GetUser(login, pass string) (User, error) {
	res, err := getUserRepo(userRepo, login, pass)
	if err != nil{
		return User{}, err
	}
	return User{res["login"], res["username"], res["pass"]}, nil
}