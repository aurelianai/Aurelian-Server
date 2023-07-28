package main

import (
	"AELS/persistence"
	"AELS/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	persistence.InitAndMigrate()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	log.Fatal(http.ListenAndServe(":2140",
		handlers.LoggingHandler(os.Stdout, r),
	))
}
