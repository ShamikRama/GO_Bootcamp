package db

import "Data_Api/models"

type Store interface {
	GetPlaces(limit, offset int) ([]models.Place, int, error)
	GetNearestPlaces(lat, lon float64, limit int) ([]models.Place, error)
}
