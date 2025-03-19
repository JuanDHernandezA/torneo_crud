package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerTiposEstadistica(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM tipo_estadistica")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipos_estadistica []models.Tipo_estadistica
	for rows.Next() {
		var u models.Tipo_estadistica
		err := rows.Scan(&u.Id_tipo_estadistica, &u.Nombre_tipo_estadistica)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tipos_estadistica = append(tipos_estadistica, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos_estadistica)
}