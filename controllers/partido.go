package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerPartidos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM partido")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var partidos []models.Partido
	for rows.Next() {
		var u models.Partido
		err := rows.Scan(&u.Id_partido, &u.Id_local, &u.Id_visitante, &u.Fecha)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		partidos = append(partidos, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partidos)
}