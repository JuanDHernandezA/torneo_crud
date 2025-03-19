package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerTiposTorneo(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM tipo_torneo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipos_torneo []models.Tipo_torneo
	for rows.Next() {
		var u models.Tipo_torneo
		err := rows.Scan(&u.Id_tipo_torneo, &u.Nombre_tipo_torneo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tipos_torneo = append(tipos_torneo, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos_torneo)
}