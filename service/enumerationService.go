package service

import (
	"log"
	"srbolabApp/database"
	"srbolabApp/model"
)

var (
	EnumerationService enumerationServiceInterface = &enumerationService{}
)

type enumerationService struct {
}

type enumerationServiceInterface interface {
	GetAllIrregularityLevels() ([]model.IrregularityLevel, error)
}

func (e *enumerationService) GetAllIrregularityLevels() ([]model.IrregularityLevel, error) {
	levels := []model.IrregularityLevel{}
	err := database.Client.Select(&levels, `SELECT * FROM public.irregularity_levels`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return levels, nil
}
