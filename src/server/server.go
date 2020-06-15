package server

import (
	"fmt"
	"net/http"
	"encoding/json"
	"params"
	"errors"
	"model"
	"my_jwt"
)

type auth struct {
	Login string
	Pass string
}

type response struct {
	Code int `json:"code"`
	Result map[string]string `json:"result"`
}

func (resp response) get(code int, res map[string]string) (string, error) {
	resp.Code = code
	resp.Result = res
	slcB, err := json.Marshal(&resp)
	return string(slcB), err
}

func (resp response) getError(error string) string {
	_r, _ := resp.get(400, map[string]string{"error":error})
	return _r
}

var jwtClaims my_jwt.Claims

func getClaims(r *http.Request) (auth, error) {
	var a auth
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
        return a, err
    }
    if a.Login == "" || a.Pass == "" {
    	return a, errors.New("Empty login or pass")
    }
    return a, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request, param params.Init) {
	var resp response

	a, err := getClaims(r)
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	user, err := model.GetUser(a.Login, a.Pass)
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	token, err := jwtClaims.GenToken(user.Username, param)
 	if err != nil{
		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}
 	
 	
 	strR, err := resp.get(200, map[string]string{"token": token})
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, strR)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "check")
}

func GetServer(param params.Init) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "error")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		loginHandler(w, r, param)
	})
	http.HandleFunc("/check", checkHandler)
	err := http.ListenAndServe(":9000", nil)
	return err
}