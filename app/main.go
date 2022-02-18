package main

import (
	"fmt"

	"github.com/spro80/apiGolang/app/interfaces/web"
)

func main() {
	fmt.Println("Init in main")

	var port string = "9090"
	//WebServer
	//web.NewWebServer()
	//web.InitRoutes(orderUseCase)
	//web.Start(config.GetString("web.port"))
	web.Start(port)
}
