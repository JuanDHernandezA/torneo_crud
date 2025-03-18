package models

type Administrador struct {
	Id_admin     int    `json:"id_admin"`
	Nombre_admin string `json:"nombre_admin"`
	Contraseña   string `json:"contraseña"`
}
