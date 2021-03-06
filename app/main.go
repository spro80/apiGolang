package main

import (
	"fmt"

	"github.com/spro80/apiGolang/app/application/orderUseCase"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGet"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGetAll"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiLogin"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/userControllerGetAll"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/userControllerGetById"
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

	//useCase
	orderUseCase := orderUseCase.NewOrderUseCase()
	orderController := order.NewOrderController(orderUseCase)

	//controller
	templateApiController := templateApiGet.NewTemplateApiController()
	templateApiGetAllController := templateApiGetAll.NewTemplateApiGetAllController()
	templateApiLoginController := templateApiLogin.NewTemplateApiLoginController()

	apiUserGetAllController := userControllerGetAll.NewUserControllerGetAll(service)
	userControllerGetById := userControllerGetById.NewUserControllerGetBydId(serviceGetUsersById)

	//WebServer
	var port string = "8080"
	web.NewWebServer()
	web.InitRoutes(orderController, templateApiController, templateApiGetAllController, templateApiLoginController, apiUserGetAllController, userControllerGetById)
	web.Start(port)
	//web.Start(config.GetString("web.port"))
	//ordenes, eventos
}
