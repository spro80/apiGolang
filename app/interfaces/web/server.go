package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

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
