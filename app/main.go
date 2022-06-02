package main

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application/orderUseCase"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/apiGetUserById"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGet"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGetAll"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGetAllServices"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiLogin"
	"github.com/spro80/apiGolang/app/infraestructure/services"
	"github.com/spro80/apiGolang/app/interfaces/web"
)

//"github.com/spro80/apiGolang/app/infraestructure/mongo/mongo_connection"

func main() {
	fmt.Println("Init in main")

	// Database connection
	//connectionMongo := mongo_connection.GetConnection()
	//fmt.Println(connectionMongo)
	//Controller

	//services
	service := services.NewGetUsers()
	serviceGetUsersById := services.NewGetUsersById()

	//cosmos
	orderUseCase := orderUseCase.NewOrderUseCase()
	orderController := order.NewOrderController(orderUseCase)

	templateApiController := templateApiGet.NewTemplateApiController()
	templateApiGetAllController := templateApiGetAll.NewTemplateApiGetAllController()
	templateApiLoginController := templateApiLogin.NewTemplateApiLoginController()
	templateApiGetAllServicesController := templateApiGetAllServices.NewTemplateApiGetAllServicesController(service)

	getUserByIdController := apiGetUserById.NewControllerApiGetUserBydId(serviceGetUsersById)

	//WebServer
	var port string = "8080"
	web.NewWebServer()
	web.InitRoutes(orderController, templateApiController, templateApiGetAllController, templateApiLoginController, templateApiGetAllServicesController, getUserByIdController)
	web.Start(port)
	//web.Start(config.GetString("web.port"))
	//ordenes, eventos
}
