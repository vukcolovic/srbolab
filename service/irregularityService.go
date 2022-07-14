package service

import (
	"srbolabApp/database"
	"srbolabApp/loger"
	"srbolabApp/model"
	"strconv"
	"strings"
	"time"
)

var (
	IrregularityService irregularityServiceInterface = &irregularityService{}
)

type irregularityService struct {
}

type irregularityServiceInterface interface {
	GetIrregularityByID(id int) (*model.Irregularity, error)
	GetAllIrregularities(skip, take int, filter model.IrregularityFilter) ([]model.Irregularity, error)
	CreateIrregularity(model.Irregularity, int) (*model.Irregularity, error)
	DeleteIrregularity(int) error
	GetIrregularitiesCount(model.IrregularityFilter) (int, error)
	UpdateIrregularity(model.Irregularity) (*model.Irregularity, error)
	ChangeCorrected(model.Irregularity, bool, int) error
}

func (s *irregularityService) GetIrregularityByID(id int) (*model.Irregularity, error) {
	irregularitiesDb := []model.IrregularityDb{}
	err := database.Client.Select(&irregularitiesDb, `SELECT * FROM irregularities WHERE id = $1`, id)
	if err != nil || len(irregularitiesDb) == 0 {
		loger.ErrorLog.Println("Error getting irregularity by id: ", err)
		return nil, err
	}

	irrJson, err := getJsonIrregularity(irregularitiesDb[0])
	if err != nil {
		loger.ErrorLog.Println("Error getting Irregularity, error getting json from db: ", err)
		return nil, err
	}

	return irrJson, nil
}

func (s *irregularityService) GetAllIrregularities(skip, take int, filter model.IrregularityFilter) ([]model.Irregularity, error) {
	query := queryBuilderForIrregularities(skip, take, filter, false)
	irregularities := []model.IrregularityDb{}
	err := database.Client.Select(&irregularities, query)
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

func (s *irregularityService) GetIrregularitiesCount(filter model.IrregularityFilter) (int, error) {
	query := queryBuilderForIrregularities(0, 0, filter, true)
	count := []int{}
	err := database.Client.Select(&count, query)
	if err != nil || len(count) == 0 {
		loger.ErrorLog.Println("Error getting count of irregularities: ", err)
		return 0, err
	}

	return count[0], nil
}

func queryBuilderForIrregularities(skip, take int, filter model.IrregularityFilter, isCount bool) string {
	var query string
	if isCount {
		query = `SELECT count(*) FROM irregularities WHERE `
	} else {
		query = `SELECT * FROM irregularities WHERE `
	}

	if filter.Subject != "" {
		query = query + ` subject = ` + "'" + filter.Subject + "'" + ` AND `
	}
	if filter.Checked != "" {
		query = query + ` corrected = ` + filter.Checked + ` AND `
	}
	if filter.Level != nil {
		query = query + ` level_id = ` + strconv.Itoa(filter.Level.Id) + ` AND `
	}
	if filter.Controller != nil {
		query = query + ` controller_id = ` + strconv.Itoa(filter.Controller.Id) + ` AND `
	}
	if !filter.DateFrom.IsZero() {
		query = query + ` created_at >= ` + "'" + filter.DateFrom.Format("2006-01-02") + "'" + "::date" + ` AND `
	}
	if !filter.DateTo.IsZero() {
		query = query + ` created_at < ` + "'" + filter.DateTo.Format("2006-01-02") + "'" + "::date + " + "'1 day'::interval" + ` AND `
	}

	if strings.HasSuffix(query, " WHERE ") {
		query = strings.TrimRight(query, " WHERE ")
	}
	query = strings.TrimRight(query, "AND ")

	if !isCount {
		query = query + ` ORDER BY id desc OFFSET ` + strconv.Itoa(skip) + ` LIMIT ` + strconv.Itoa(take)
	}

	return query
}

func (s *irregularityService) UpdateIrregularity(irregularity model.Irregularity) (*model.Irregularity, error) {
	_, err := database.Client.Exec(`UPDATE irregularities SET subject = $1, level_id = $2, controller_id = $3, description = $4, notice = $5, updated_at = $6 WHERE id = $7`,
		irregularity.Subject, irregularity.Level.Id, irregularity.Controller.Id, irregularity.Description, irregularity.Notice, time.Now(), irregularity.Id)
	if err != nil {
		loger.ErrorLog.Println("Error updating Irregularity: ", err)
		return nil, err
	}

	//todo return that irregularity if there is need fot that
	return nil, err
}

func (s *irregularityService) ChangeCorrected(irregularity model.Irregularity, corrected bool, userId int) error {
	_, err := database.Client.Exec(`UPDATE irregularities SET corrected = $1, corrected_by = $2, updated_at = $3 WHERE id = $4`,
		corrected, userId, time.Now(), irregularity.Id)
	if err != nil {
		loger.ErrorLog.Println("Error changing corrected Irregularity: ", err)
		return err
	}

	return err
}
