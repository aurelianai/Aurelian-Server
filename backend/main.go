package main

import (
	"AELS/config"
	"AELS/persistence"
	"AELS/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	for _, v := range []string{"JWT_SECRET", "USER_CREATE_ACCESS_KEY"} {
		if os.Getenv(v) == "" {
			fmt.Printf("Required environment variable %s not defined!\n", v)
			os.Exit(1)
		}
	}

	if len(os.Args) != 2 {
		fmt.Printf("Pass the path of the config file as a cmd line arg!")
	}

	if err := config.Config.InitAndValidate(os.Args[1]); err != nil {
		fmt.Printf("An error occurred parsing the config file '%s'", os.Args[1])
	} else {
		fmt.Printf("Successfully Parsed Config File '%s'", os.Args[1])
	}

	persistence.InitAndMigrate()

	r := mux.NewRouter()
	routes.SetupRoutes(r)
	fmt.Println("Starting Aurelian Entperise Language Server on 0.0.0.0:2140")
	log.Fatal(http.ListenAndServe(":2140",
		handlers.LoggingHandler(os.Stdout, r),
	))
}
