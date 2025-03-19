package models

type Torneo struct {
	Id_torneo int    `json:"id_torneo"`
	Nombre_torneo string `json:"nombre_torneo"`
	Fecha_inicio string `json:"fecha_inicio"`
	Fecha_fin string `json:"fecha_fin"`
	Id_tipo_torneo int `json:"id_tipo_torneo"`
	Id_admin int `json:"id_admin"`
	Id_estado_torneo int `json:"id_estado_torneo"`
}