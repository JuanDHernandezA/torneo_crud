package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	// Rutas
	routes.SetupRoutes(r)

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"}, // Permitir Angular
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(r)))
}
