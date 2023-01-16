package services

import (
	"crud/src/database"
	"crud/src/models"
)

func GetAll() ([]models.Product, error) {

	products, err := database.GetAll()

	if err != nil || products == nil {

		return []models.Product{}, err
	}

	return products, nil
}

func Get(id string) (models.Product, error) {

	product, err := database.GetById(id)

	if err != nil {

		return models.Product{}, nil
	}

	return product, nil
}

func Create(product models.Product) (string, error) {

	id, err := database.Create(product)

	if err != nil {

		return "", err
	}

	return id, nil
}
