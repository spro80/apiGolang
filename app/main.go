package main

import (
	"fmt"

	"github.com/spro80/apiGolang/app/interfaces/web"
)

func main() {
	fmt.Println("Init in main")

	//WebServer
	var port string = "8080"
	web.NewWebServer()
	web.InitRoutes()
	web.Start(port)
	//web.Start(config.GetString("web.port"))
}
