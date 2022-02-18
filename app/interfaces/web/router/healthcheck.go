package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
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

func InitRoutes() {
	routes.NewHealthHandler(echoServer)
}

func Start(port string) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  3 * time.Minute,
		WriteTimeout: 3 * time.Minute,
	}
	log.Info("App listen in port %s", port)
	log.Info(server)
	//echoServer.Logger.Fatal(echoServer.StartServer(server))
}
