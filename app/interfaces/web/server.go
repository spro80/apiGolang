package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGet"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGetAll"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiGetAllServices"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/templateApiLogin"
	"github.com/spro80/apiGolang/app/interfaces/web/routes"
)

var echoServer *echo.Echo

func NewWebServer() {
	echoServer = echo.New()
	echoServer.HideBanner = true
	//echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.RequestID())
	//echoServer.Validator = json_validator.NewJsonValidator()
}

func InitRoutes(saveOrderUseCase order.OrderControllerInterface, templateApiGetController templateApiGet.TemplateApiGetControllerInterface, templateApiGetAllController templateApiGetAll.TemplateApiGetAllControllerInterface, templateApiLoginController templateApiLogin.TemplateApiLoginControllerInterface, templateApiGetAllServicesController templateApiGetAllServices.TemplateApiGetAllServicesControllerInterface) {
	routes.NewHealthHandler(echoServer)
	routes.NewPingHandler(echoServer)
	routes.NewProcessOrderHandler(echoServer, saveOrderUseCase)
	routes.NewTemplateApiHandler(echoServer, templateApiGetController)
	routes.NewTemplateApiGetAllHandler(echoServer, templateApiGetAllController)
	routes.NewTemplateApiLoginHandler(echoServer, templateApiLoginController)
	routes.NewTemplateApiGetAllServicesHandler(echoServer, templateApiGetAllServicesController)
}

func Start(port string) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	log.Info("App listen in port %s", port)
	echoServer.Logger.Fatal(echoServer.StartServer(server))
}

/*
func HealtCheck() {
	fmt.Println("Prueba Healtcheck")
}
func Start(port string) {
	fmt.Println("Init in Start")
	router := mux.NewRouter()
	//router.HandleFunc("/health", HealtCheck()).Methods("GET")
	PORT := port //os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
*/
