package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	// Rutas
	routes.SetupRoutes(r)

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
