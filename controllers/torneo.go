package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerTorneos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM torneo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var torneos []models.Torneo
	for rows.Next() {
		var u models.Torneo
		err := rows.Scan(&u.Id_torneo, &u.Nombre_torneo, &u.Fecha_inicio, &u.Fecha_fin, &u.Id_tipo_torneo, &u.Id_admin, &u.Id_estado_torneo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		torneos = append(torneos, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(torneos)
}