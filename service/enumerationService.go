package service

import (
	"errors"
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
	GetAllIrregularityLevelById(id int) (*model.IrregularityLevel, error)
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

func (e *enumerationService) GetAllIrregularityLevelById(id int) (*model.IrregularityLevel, error) {
	levels := []model.IrregularityLevel{}
	err := database.Client.Select(&levels, `SELECT * FROM public.irregularity_levels WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(levels) < 1 {
		log.Println("No level by id ", id)
		return nil, errors.New("no level by id")
	}

	return &levels[0], nil
}
