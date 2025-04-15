package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
	"github.com/JuanDHernandezA/torneo_crud/services"
)

func TrRegistrarEquipo(w http.ResponseWriter, r *http.Request) {
	var TrRegistrarEquipo models.TrRegistrarEquipo
	dec := json.NewDecoder(io.NopCloser(io.Reader(r.Body)))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&TrRegistrarEquipo); err == nil {

		if tx, errorTrDB := config.DB.Begin(); errorTrDB == nil { //inicia la transacción
			if response, err := services.CrearEquipo(tx, TrRegistrarEquipo.Equipo); err == nil {

				TrRegistrarEquipo.Equipo.Id_equipo = response.Id_equipo

				for i, jugador := range TrRegistrarEquipo.Jugadores {

					jugador.Id_equipo = response.Id_equipo

					if response, err := services.CrearJugador(tx, jugador); err == nil {

						TrRegistrarEquipo.Jugadores[i].Id_jugador = response.Id_jugador

					} else {
						tx.Rollback()
						http.Error(w, "Error creando jugador: "+err.Error(), http.StatusInternalServerError)
						return
					}
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(TrRegistrarEquipo)

			} else {
				tx.Rollback()
				http.Error(w, "Error creando equipo: "+err.Error(), http.StatusInternalServerError)
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
	}
}
