package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerEquipos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM equipo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipos []models.Equipo
	for rows.Next() {
		var u models.Equipo
		err := rows.Scan(&u.Id_equipo, &u.Nombre_equipo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		equipos = append(equipos, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipos)
}