package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spro80/apiGolang/app/infraestructure/controllers/order"
	"github.com/spro80/apiGolang/app/interfaces/web/routes"
)

// "github.com/spro80/apiGolang/app/infraestructure/controllers/useCaseOne"

var echoServer *echo.Echo

func NewWebServer() {
	echoServer = echo.New()
	echoServer.HideBanner = true
	//echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORS())
	echoServer.Use(middleware.RequestID())
	//echoServer.Validator = json_validator.NewJsonValidator()
}

func InitRoutes(saveOrderUseCase order.OrderControllerInterface) {
	routes.NewHealthHandler(echoServer)
	routes.NewSaveOrderHandler(echoServer, saveOrderUseCase)
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
