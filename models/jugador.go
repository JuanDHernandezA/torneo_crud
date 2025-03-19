package models

type Jugador struct {
	Id_jugador int    `json:"id_jugador"`
	Nombre_jugador string `json:"nombre_jugador"`
	Numero_jugador int `json:"numero_jugador"`
	Id_posicion int `json:"id_posicion"`
	Es_capitan bool `json:"es_capitan"`
	Id_equipo int `json:"id_equipo"`
}