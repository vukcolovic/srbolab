package service

import (
	"srbolabApp/database"
	"srbolabApp/loger"
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
	DeleteIrregularity(int) error
}

func (s *irregularityService) GetAllIrregularities(skip, take int) ([]model.Irregularity, error) {
	irregularities := []model.IrregularityDb{}
	err := database.Client.Select(&irregularities, `SELECT * FROM irregularities ORDER BY id desc OFFSET $1 LIMIT $2`, skip, take)
	if err != nil {
		loger.ErrorLog.Println("Error getting all Irregularities: ", err)
		return nil, err
	}

	result := []model.Irregularity{}
	for _, irr := range irregularities {
		irrJson, err := getJsonIrregularity(irr)
		if err != nil {
			loger.ErrorLog.Println("Error getting all Irregularities, error getting json from db: ", err)
			return nil, err
		}

		result = append(result, *irrJson)
	}

	return result, nil
}

func (s *irregularityService) CreateIrregularity(irregularity model.Irregularity, userId int) (*model.Irregularity, error) {
	_, err := database.Client.Exec(`INSERT INTO irregularities (subject, level_id, controller_id, created_by, description, notice, corrected, corrected_date, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		irregularity.Subject, irregularity.Level.Id, irregularity.Controller.Id, userId, irregularity.Description, irregularity.Notice, irregularity.Corrected, irregularity.CorrectedDate, time.Now(), time.Now())
	if err != nil {
		loger.ErrorLog.Println("Error creating Irregularity: ", err)
		return nil, err
	}

	//todo return that irregularity if there is need fot that
	return nil, err
}

func (s *irregularityService) DeleteIrregularity(id int) error {
	_, err := database.Client.Exec(`DELETE FROM irregularities WHERE id = $1`, id)
	if err != nil {
		loger.ErrorLog.Println("Error deleting Irregularity: ", err)
		return err
	}

	return nil
}

func getJsonIrregularity(i model.IrregularityDb) (*model.Irregularity, error) {
	jsonIrr := model.Irregularity{}

	jsonIrr.Id = i.Id
	jsonIrr.Subject = i.Subject
	level, err := EnumerationService.GetAllIrregularityLevelById(i.Level)
	if err != nil {
		loger.ErrorLog.Println("Error getting IrregularityJson from IrregularityDb, error getting IrregularityLevel by id: ", err)
		return nil, err
	}
	jsonIrr.Level = *level

	usersMap := make(map[int64]*model.User)

	controllor, err := UsersService.GetUserByID(int(i.Controller.Int64))
	if err != nil {
		loger.ErrorLog.Println("Error getting IrregularityJson from IrregularityDb, error getting user by id: ", err)
		return nil, err
	}
	usersMap[int64(controllor.Id)] = controllor
	jsonIrr.Controller = *controllor
	if i.CreatedBy.Valid {
		createdBy, ok := usersMap[i.CreatedBy.Int64]
		if !ok {
			createdBy, err = UsersService.GetUserByID(int(i.CreatedBy.Int64))
			if err != nil {
				loger.ErrorLog.Println("Error getting IrregularityJson from IrregularityDb, error getting user by id: ", err)
				return nil, err
			}
			usersMap[int64(createdBy.Id)] = createdBy
		}
		jsonIrr.CreatedBy = *createdBy
	}

	jsonIrr.Description = i.Description
	jsonIrr.Corrected = i.Corrected
	jsonIrr.Notice = i.Notice

	if i.CorrectedBy.Valid {
		correctedBy, ok := usersMap[i.CorrectedBy.Int64]
		if !ok {
			correctedBy, err = UsersService.GetUserByID(int(i.CorrectedBy.Int64))
			if err != nil {
				loger.ErrorLog.Println("Error getting IrregularityJson from IrregularityDb, error getting user by id: ", err)
				return nil, err
			}
			usersMap[int64(correctedBy.Id)] = correctedBy
		}
		jsonIrr.CorrectedBy = *correctedBy
	}

	jsonIrr.CorrectedDate = i.CorrectedDate
	jsonIrr.CreatedAt = i.CreatedAt
	jsonIrr.UpdatedAt = i.UpdatedAt

	return &jsonIrr, nil
}
