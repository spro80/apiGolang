package main

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application/orderUseCase"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
	"github.com/spro80/apiGolang/app/interfaces/web"
)

//"github.com/spro80/apiGolang/app/infraestructure/mongo/mongo_connection"

func main() {
	fmt.Println("Init in main")

	// Database connection
	//connectionMongo := mongo_connection.GetConnection()
	//fmt.Println(connectionMongo)
	//Controller

	orderUseCase := orderUseCase.NewOrderUseCase()
	orderController := order.NewOrderController(orderUseCase)
	//WebServer
	var port string = "8080"
	web.NewWebServer()
	web.InitRoutes(orderController)
	web.Start(port)
	//web.Start(config.GetString("web.port"))
}
