package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `example:"Aluno Teste" binding:"required" validate:"nonzero"`
	CPF        string `example:"01234567890"  binding:"required"  validate:"len=11,regexp=^[0-9]*$"`
	RG         string `example:"123456789"   binding:"required"   validate:"len=9,regexp=^[0-9]*$"`
}

func (c *Student) ValidateData() error {
	if err := validator.Validate(c); err != nil {
		return err
	}
	return nil
}

func (c *Student) ToStudentView() StudentView {
	return StudentView{
		Id:        int(c.ID),
		Name:      c.Name,
		CPF:       c.CPF,
		RG:        c.RG,
		CreatedAt: c.CreatedAt,
	}
}
