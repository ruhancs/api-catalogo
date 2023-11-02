package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Category struct {
	Name        string    `json:"name" valid:"required,stringlength(4|25)"`
	Description string    `json:"description" valid:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewCategory(name, description string, createdAt time.Time) (*Category,error) {
	category := &Category{
		Name: name,
		Description: description,
		CreatedAt: createdAt,
	}

	err := category.Validate()
	if err != nil{
		return nil,err
	}

	return category,nil
}

func (c *Category) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}