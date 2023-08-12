package main

import (
	"AELS/persistence"
	"AELS/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var REQUIRED_ENV_VARS = []string{"INFERENCE_BACKEND", "JWT_SECRET", "USER_CREATE_ACCESS_KEY"}

func main() {
	for _, v := range REQUIRED_ENV_VARS {
		if os.Getenv(v) == "" {
			fmt.Printf("Required environment variable %s not defined!\n", v)
			os.Exit(1)
		}
	}

	persistence.InitAndMigrate()

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	fmt.Println("Starting Aurelian Entperise Language Server on 0.0.0.0:2140")
	log.Fatal(http.ListenAndServe(":2140",
		handlers.LoggingHandler(os.Stdout, r),
	))
}
