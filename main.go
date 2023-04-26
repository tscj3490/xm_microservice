package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tscj3490/xm_microservice/handlers"
)

func main() {
	// Create new Router
	router := mux.NewRouter()
	// route properly to respective handlers
	router.Handle("/companies", handlers.CreateCompanyHandler()).Methods("POST")
	router.Handle("/companies/{id}", handlers.GetCompanyHandler()).Methods("GET")
	router.Handle("/companies/{id}", handlers.DeleteCompanyHandler()).Methods("DELETE")
	router.Handle("/companies/{id}", handlers.UpdateCompanyHandler()).Methods("PUT")

	// Create new server and assign the router
	server := http.Server{
		Addr:    ":9090",
		Handler: handlers.AuthHandler(router),
	}
	fmt.Println("Staring Company Microservice on Port 9090")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to start HTTP Server")
	}

	// Start Server on defined port/host.
	// isAzureDeployment := os.Getenv("AZURE_DEPLOYMENT")
	// if isAzureDeployment == "TRUE" {
	// 	fmt.Println("Running on Azure")
	// 	err := server.ListenAndServe()
	// 	if err != nil {
	// 		fmt.Println("Failed to start HTTP Server")
	// 	}
	// } else {
	// 	err := server.ListenAndServeTLS("server.crt", "server.key")
	// 	if err != nil {
	// 		fmt.Printf("Failed to start HTTPS server: %s", err.Error())
	// 	}
	// }
}
