package main

import (
	"fmt"

	"github.com/luizhirasawa/Go-api.git/config"
	"github.com/luizhirasawa/Go-api.git/router"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
