package routes

import (
	"github.com/JuanDHernandezA/torneo_crud/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/administrador", controllers.ObtenerAdmins).Methods("GET")
	r.HandleFunc("/administrador", controllers.CrearAdministrador).Methods("POST")
}
