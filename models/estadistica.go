package models

type Estadistica struct {
	Id_estadistica 		int	`json:"id_estadistica"`
	Cantidad			int	`json:"cantidad"`
	Id_partido 			int	`json:"id_partido"`
	Id_jugador 			int	`json:"id_jugador"`
	Id_tipo_estadistica int	`json:"id_tipo_estadistica"`
}