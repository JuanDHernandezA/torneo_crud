package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerEquiposTorneos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM equipo_torneo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipos_torneos []models.Equipo_torneo
	for rows.Next() {
		var u models.Equipo_torneo
		err := rows.Scan(&u.Id_equipo_toneo, &u.Id_equipo, &u.Id_torneo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		equipos_torneos = append(equipos_torneos, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipos_torneos)
}