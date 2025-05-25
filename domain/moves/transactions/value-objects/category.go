package transactionVals

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCategory = errors.New("invalid category, please try again")
)

var categoryMap = map[string]bool{
	"Alimentación":  true,
	"Transporte":    true,
	"Vivienda":      true,
	"Servicios":     true,
	"Salud":         true,
	"Educación":     true,
	"Ocio":          true,
	"Viajes":        true,
	"Ropa":          true,
	"Tecnología":    true,
	"Ahorro":        true,
	"Inversiones":   true,
	"Deudas":        true,
	"Impuestos":     true,
	"Mascotas":      true,
	"Donaciones":    true,
	"Regalos":       true,
	"Suscripciones": true,
	"Hogar":         true,
	"Otros":         true,
}

type Category struct {
	Val string
}

func NewCategory(category string) (Category, error) {
	if _, exists := categoryMap[category]; !exists {
		fmt.Println("category passed: ", category)
		return Category{}, ErrInvalidCategory
	}

	return Category{Val: category}, nil
}
