package models

import "time"

type StudentView struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	RG        string    `json:"rg"`
	CreatedAt time.Time `json:"created_at"`
}
