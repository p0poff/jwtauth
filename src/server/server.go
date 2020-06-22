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

func getToken(r *http.Request) (string, error) {
	return "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6InBvcG9mZiIsImV4cCI6MTU5MjgyODM0NCwiaWF0IjoxNTkyODI3MzQ0LCJpc3MiOiJqd3QifQ.zmI78fbsoFWr6TBY9PAnyRBNQ2HHIEXfKwk2IPAr3Is", nil
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

 	fmt.Fprintf(w, strR)
}

func checkHandler(w http.ResponseWriter, r *http.Request, param params.Init) {
	var resp response

	token, err := getToken(r)
	if err != nil {
		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
	}

	isValid, err := jwtClaims.IsValid(token, param)
	if err != nil {
		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
	}

	res := map[string]string{"res": ""}

	if isValid {
		res["res"] = "valid"
	} else {
		res["res"] = "no valid"
	}
	
	strR, err := resp.get(200, res)
 	if err != nil {
 		http.Error(w, resp.getError(err.Error()), http.StatusBadRequest)
 		return
 	}

 	fmt.Fprintf(w, strR)
}

func handlerWrapper(param params.Init, f func(http.ResponseWriter, *http.Request, params.Init)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r, param)
	}
}

func GetServer(param params.Init) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "error")
	})
	
	http.HandleFunc("/login", handlerWrapper(param, loginHandler))

	http.HandleFunc("/check", handlerWrapper(param, checkHandler))
	err := http.ListenAndServe(":9000", nil)
	return err
}