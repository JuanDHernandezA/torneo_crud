package services

import (
	"database/sql"

	"github.com/JuanDHernandezA/torneo_crud/models"
)

func CrearJugador(tx *sql.Tx, jugador models.Jugador) (models.Jugador, error) {

	err := tx.QueryRow("INSERT INTO jugador (nombre_jugador, numero_jugador, id_posicion, es_capitan, id_equipo) VALUES ($1, $2, $3, $4, $5) RETURNING id_jugador",jugador.Nombre_jugador, jugador.Numero_jugador, jugador.Id_posicion, jugador.Es_capitan, jugador.Id_equipo).Scan(&jugador.Id_jugador)
	if err != nil {
		return jugador,err
	}

	return jugador, nil
}