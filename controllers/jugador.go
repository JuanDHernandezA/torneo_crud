package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerJugadores(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM jugador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var jugadores []models.Jugador
	for rows.Next() {
		var u models.Jugador
		err := rows.Scan(&u.Id_jugador, &u.Nombre_jugador, &u.Numero_jugador, &u.Id_posicion, &u.Es_capitan, &u.Id_equipo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jugadores = append(jugadores, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jugadores)
}