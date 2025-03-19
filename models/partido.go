package models

type Partido struct {
	Id_partido    int    `json:"id_partido"`
	Id_local	  int    `json:"id_local"`
	Id_visitante  int    `json:"id_visitante"`
	Fecha         string `json:"fecha"`
}