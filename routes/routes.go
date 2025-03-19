package routes

import (
	"github.com/JuanDHernandezA/torneo_crud/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {

	//administrador
	r.HandleFunc("/administrador", controllers.ObtenerAdmins).Methods("GET")
	r.HandleFunc("/administrador", controllers.CrearAdministrador).Methods("POST")

	//equipo
	r.HandleFunc("/equipo", controllers.ObtenerEquipos).Methods("GET")

	//equipo_torneo
	r.HandleFunc("/equipo_torneo", controllers.ObtenerEquiposTorneos).Methods("GET")

	//estadistica
	r.HandleFunc("/estadistica", controllers.ObtenerEstadisticas).Methods("GET")

	//estado_torneo
	r.HandleFunc("/estado_torneo", controllers.ObtenerEstadosTorneo).Methods("GET")

	//jugador
	r.HandleFunc("/jugador", controllers.ObtenerJugadores).Methods("GET")

	//partido
	r.HandleFunc("/partido", controllers.ObtenerPartidos).Methods("GET")

	//posicion
	r.HandleFunc("/posicion", controllers.ObtenerPosiciones).Methods("GET")
	
	//tipo_estadistica
	r.HandleFunc("/tipo_estadistica", controllers.ObtenerTiposEstadistica).Methods("GET")

	//tipo_torneo
	r.HandleFunc("/tipo_torneo", controllers.ObtenerTiposTorneo).Methods("GET")

	//torneo
	r.HandleFunc("/torneo", controllers.ObtenerTorneos).Methods("GET")
}
