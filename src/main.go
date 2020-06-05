package main

import (
	"fmt"
	"params"
	"server"
)

func main() {
	fmt.Println("jwt go...")
	param := params.InitParams()
	err := server.GetServer(param)
	if err != nil {
        fmt.Println(err)
    }
}