package cmd

import (
	"fmt"
	"gonote/config"
	"gonote/router"
)

func Start() {
	config.InitConfig()
	router.InitRouter()
}

func Clear() {
	fmt.Println("==================Clear===================")
}
