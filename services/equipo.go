package services

import (
	"database/sql"

	"github.com/JuanDHernandezA/torneo_crud/models"
)

func CrearEquipo(tx *sql.Tx, equipo models.Equipo) (models.Equipo, error) {

	err := tx.QueryRow("INSERT INTO equipo (nombre_equipo) VALUES ($1) RETURNING id_equipo",equipo.Nombre_equipo).Scan(&equipo.Id_equipo)
	if err != nil {
		return equipo,err
	}

	return equipo, nil
}