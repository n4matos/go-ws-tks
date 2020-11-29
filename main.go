package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nathanzeras/go-ws-tks/app/handlers"
	"github.com/nathanzeras/go-ws-tks/config"
)

func main() {

	r := mux.NewRouter()

	config.LoadEnv()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	r.HandleFunc("/api/v1/recebimento", config.Use(handlers.RecebimentoGeral, config.BasicAuth)).Methods("POST")
	//r.HandleFunc("/api/v1/search", config.Use(handlers.Consulta)).Methods("GET")
	r.HandleFunc("/api/autorizacoes/{cpf:[0-9]+}/{procedimento:[0-9]+}", config.Use(handlers.Search, config.BasicAuth)).Methods("GET")
	//r.HandleFunc("/api/v1/recebimentox", config.Use(handlers.Recebimento, config.BasicAuth)).Methods("POST")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, r) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
