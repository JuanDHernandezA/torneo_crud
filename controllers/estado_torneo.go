package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerEstadosTorneo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM estado_torneo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var estados_torneo []models.Estado_torneo
	for rows.Next() {
		var u models.Estado_torneo
		err := rows.Scan(&u.Id_estado_torneo, &u.Nombre_estado)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		estados_torneo = append(estados_torneo, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estados_torneo)
}