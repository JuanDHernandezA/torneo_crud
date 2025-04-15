package models

type TrRegistrarEquipo struct {
	Equipo 		Equipo		`json:"equipo"`
	Jugadores 	[]Jugador	`json:"jugadores"`
}