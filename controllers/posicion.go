package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerPosiciones(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM posicion")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posiciones []models.Posicion
	for rows.Next() {
		var u models.Posicion
		err := rows.Scan(&u.Id_posicion, &u.Nombre_posicion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posiciones = append(posiciones, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posiciones)
}