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

func ObtenerEquipos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM equipo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var equipos []models.Equipo
	for rows.Next() {
		var equipo models.Equipo
		err := rows.Scan(&equipo.Id_equipo, &equipo.Nombre_equipo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		equipos = append(equipos, equipo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipos)
}

func CrearEquipo(w http.ResponseWriter, r *http.Request) {
	var equipo models.Equipo

	dec := json.NewDecoder(io.NopCloser(io.Reader(r.Body)))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&equipo); err == nil {
		fmt.Println(equipo)

		if tx, errorTrDB := config.DB.Begin(); errorTrDB == nil { //inicia la transacción
			if response, err := services.CrearEquipo(tx, equipo); err == nil {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
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

	} else {
		http.Error(w, "Error: formato incorrecto", http.StatusBadRequest)
		return
	}
}
