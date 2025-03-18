package models

type Administrador struct {
	Id_admin     	int    `json:"id"`
	Nombre_admin 	string `json:"nombre"`
	Contraseña  	string `json:"email"`
}
