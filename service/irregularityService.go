package service

import (
	"srbolabApp/database"
	"srbolabApp/model"
	"time"
)

var (
	IrregularityService irregularityServiceInterface = &irregularityService{}
)

type irregularityService struct {
}

type irregularityServiceInterface interface {
	GetAllIrregularities(skip, take int) ([]model.Irregularity, error)
	CreateIrregularity(model.Irregularity, int) (*model.Irregularity, error)
}

func (s *irregularityService) GetAllIrregularities(skip, take int) ([]model.Irregularity, error) {
	return []model.Irregularity{}, nil
}

func (s *irregularityService) CreateIrregularity(irregularity model.Irregularity, userId int) (*model.Irregularity, error) {
	_, err := database.Client.Exec(`INSERT INTO public.irregularities (subject, level_id, controller_id, created_by, description, notice, corrected, corrected_by, corrected_date, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		irregularity.Subject, irregularity.Level.Id, irregularity.Controller.Id, userId, irregularity.Description, irregularity.Notice, irregularity.Corrected, irregularity.CorrectedBy.Id, irregularity.CorrectedDate, time.Now(), time.Now())
	if err != nil {
		//loger.Instance().Error("error inserting format", loger.AdditionalFields{"Error": err, "DbKey": formatToAdd.DbKey})
		return nil, err
	}

	//todo return that irregularity if there is need fot that
	return nil, err
}
