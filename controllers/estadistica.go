package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerEstadisticas(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM estadistica")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var estadisticas []models.Estadistica
	for rows.Next() {
		var u models.Estadistica
		err := rows.Scan(&u.Id_estadistica, &u.Cantidad, &u.Id_partido, &u.Id_jugador, &u.Id_tipo_estadistica)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		estadisticas = append(estadisticas, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estadisticas)
}