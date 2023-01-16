package services

import (
	"compression/src/models"
	"compression/src/repository"
)

type CompressionService struct {
	CompressionRepository *repository.CompressionRepository
}

func (s *CompressionService) GetAll() ([]models.Compression, error) {

	compressionsDb, err := s.CompressionRepository.GetAll()

	if err != nil {

		return nil, err
	}

	var compressions []models.Compression

	for _, c := range compressionsDb {

		compressions = append(compressions, c.ToView())
	}

	return compressions, nil
}

func (s *CompressionService) GetById(id string) (models.Compression, error) {

	CompressionDb, err := s.CompressionRepository.GetById(id)

	if err != nil {

		return models.Compression{}, err
	}

	compression := CompressionDb.ToView()

	return compression, nil
}

func (s *CompressionService) Create(compression models.Compression) (string, error) {

	compressionDb := compression.ToDB()

	id, err := s.CompressionRepository.Create(compressionDb)

	if err != nil {

		return "", err
	}

	return id, nil
}
