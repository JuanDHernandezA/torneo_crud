package routes

import (
	"github.com/gorilla/mux"
	"github.com/JuanDHernandezA/torneo_crud/controllers"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/usuarios", controllers.ObtenerAdmins).Methods("GET")
	r.HandleFunc("/usuarios", controllers.CrearAdministrador).Methods("POST")
}
