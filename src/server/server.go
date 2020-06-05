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
	
	a, err := getClaims(r)
 	if err != nil {
 		http.Error(w, err.Error(), http.StatusBadRequest)
 		return
 	}

 	user, err := model.GetUser(a.Login, a.Pass)
 	if err != nil {
 		http.Error(w, err.Error(), http.StatusBadRequest)
 		return
 	}

 	token, err := jwtClaims.GenToken(user.Username, param)
 	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
 		return
 	}
    fmt.Fprintf(w, "Person: %+v\n", user)

    fmt.Fprintf(w, token)
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