package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JuanDHernandezA/torneo_crud/config"
	"github.com/JuanDHernandezA/torneo_crud/models"
)

func ObtenerAdmins(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT * FROM administrador")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []models.Administrador
	for rows.Next() {
		var u models.Administrador
		err := rows.Scan(&u.Id_admin, &u.Nombre_admin, &u.Contraseña)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		admins = append(admins, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(admins)
}

func CrearAdministrador(w http.ResponseWriter, r *http.Request) {
	var u models.Administrador
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO administrador (nombre_admin, contraseña) VALUES ($1, $2) RETURNING id_admin"
	err = config.DB.QueryRow(query, u.Nombre_admin, u.Contraseña).Scan(&u.Id_admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}
