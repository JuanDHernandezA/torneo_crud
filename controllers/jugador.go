package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
	"github.com/JuanDHernandezA/torneo_crud/services"
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
		var jugador models.Jugador
		err := rows.Scan(&jugador.Id_jugador, &jugador.Nombre_jugador, &jugador.Numero_jugador, &jugador.Id_posicion, &jugador.Es_capitan, &jugador.Id_equipo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jugadores = append(jugadores, jugador)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jugadores)
}

func CrearJugador(w http.ResponseWriter, r *http.Request) {
	var jugador models.Jugador

	dec := json.NewDecoder(io.NopCloser(io.Reader(r.Body)))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&jugador); err == nil {
		fmt.Println(jugador)

		if tx, errorTrDB := config.DB.Begin(); errorTrDB == nil { //inicia la transacción
			if response, err := services.CrearJugador(tx, jugador); err == nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			} else {
				tx.Rollback()
				http.Error(w, "Error creando jugador: "+err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tx.Commit(); err != nil {
				http.Error(w, "Error finalizando transacción: "+err.Error(), http.StatusInternalServerError)
				return
			}
			
		} else {
			http.Error(w, errorTrDB.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "Error: formato incorrecto", http.StatusBadRequest)
		return
	}
}
