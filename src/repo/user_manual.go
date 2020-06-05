package repo

import (
	"errors"
)

type UserManual map[string]map[string]string

func (d UserManual) GetUser(login, pass string) (map[string]string, error) {
	data := make(map[string]map[string]string)
	data["pop"] = map[string]string{"login": "pop", "pass": "pass", "username": "popoff"}
	data["foo"] = map[string]string{"login": "foo", "pass": "pass", "username": "bar"}

	res, ok := data[login]
	if ok {
		if res["pass"] == pass {
			return res, nil
		}
	}

	return map[string]string{}, errors.New("User not found")
}