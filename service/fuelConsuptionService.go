package service

import (
	"srbolabApp/database"
	"srbolabApp/loger"
	"srbolabApp/model"
	"srbolabApp/util"
	"strconv"
	"strings"
	"time"
)

var (
	FuelConsumptionService fuelConsumptionServiceInterface = &fuelConsumptionService{}
)

type fuelConsumptionService struct {
}

type fuelConsumptionServiceInterface interface {
	GetAllFuelConsumptions(skip, take int, filter model.FuelConsumptionFilter) ([]model.FuelConsumption, error)
	GetFuelConsumptionByID(id int) (*model.FuelConsumption, error)
	CreateFuelConsumption(fs model.FuelConsumption, userId int) (*model.FuelConsumption, error)
	UpdateFuelConsumption(model.FuelConsumption) (*model.FuelConsumption, error)
	DeleteFuelConsumption(int) error
	GetFuelConsumptionCount(filter model.FuelConsumptionFilter) (int, error)
	CountSumPrice(filter model.FuelConsumptionFilter) (float64, error)
}

func (s *fuelConsumptionService) GetFuelConsumptionByID(id int) (*model.FuelConsumption, error) {
	fuelConsumptionsDb := []model.FuelConsumptionDb{}
	err := database.Client.Select(&fuelConsumptionsDb, `SELECT * FROM fuel_consumption WHERE id = $1`, id)
	if err != nil || len(fuelConsumptionsDb) == 0 {
		loger.ErrorLog.Println("Error getting fuel consumption by id: ", err)
		return nil, err
	}

	fsJson, err := getJsonFuelConsumption(fuelConsumptionsDb[0])
	if err != nil {
		loger.ErrorLog.Println("Error getting fuel consumption, error getting json from db: ", err)
		return nil, err
	}

	return fsJson, nil
}

func (s *fuelConsumptionService) GetAllFuelConsumptions(skip, take int, filter model.FuelConsumptionFilter) ([]model.FuelConsumption, error) {
	query := queryBuilderForFuelConsumptions(skip, take, filter, false, false)
	fuelConsDb := []model.FuelConsumptionDb{}
	err := database.Client.Select(&fuelConsDb, query)
	if err != nil {
		loger.ErrorLog.Println("Error getting all fuel consumptions: ", err)
		return nil, err
	}

	result := []model.FuelConsumption{}
	for _, fc := range fuelConsDb {
		fcJson, err := getJsonFuelConsumption(fc)
		if err != nil {
			loger.ErrorLog.Println("Error getting all FuelConsumtpio, error getting json from db: ", err)
			return nil, err
		}

		result = append(result, *fcJson)
	}

	return result, nil
}

func queryBuilderForFuelConsumptions(skip, take int, filter model.FuelConsumptionFilter, isCount bool, isSum bool) string {
	var query string
	if isCount {
		query = `SELECT count(*) FROM fuel_consumption WHERE `
	} else if isSum {
		query = `SELECT sum(price) FROM fuel_consumption WHERE `
	} else {
		query = `SELECT * FROM fuel_consumption WHERE `
	}

	if filter.CarRegistration != "" {
		query = query + ` car_registration = ` + "'" + filter.CarRegistration + "'" + ` AND `
	}
	if filter.PouredBy != nil {
		query = query + ` poured_by = ` + strconv.Itoa(filter.PouredBy.Id) + ` AND `
	}
	if !filter.DateFrom.IsZero() {
		query = query + ` date_consumption >= ` + "'" + filter.DateFrom.Format("2006-01-02") + "'" + "::date" + ` AND `
	}
	if !filter.DateTo.IsZero() {
		query = query + ` date_consumption < ` + "'" + filter.DateTo.Format("2006-01-02") + "'" + "::date + " + "'1 day'::interval" + ` AND `
	}

	if strings.HasSuffix(query, " WHERE ") {
		query = strings.TrimRight(query, " WHERE ")
	}
	query = strings.TrimRight(query, "AND ")

	if !isCount && !isSum {
		query = query + ` ORDER BY id desc OFFSET ` + strconv.Itoa(skip) + ` LIMIT ` + strconv.Itoa(take)
	}

	return query
}

func getJsonFuelConsumption(fcDb model.FuelConsumptionDb) (*model.FuelConsumption, error) {
	jsonFc := model.FuelConsumption{}

	jsonFc.Id = fcDb.Id
	jsonFc.FuelType = fcDb.FuelType
	jsonFc.DateConsumption = util.Date{fcDb.DateConsumption}
	jsonFc.Price = fcDb.Price
	jsonFc.Liter = fcDb.Liter
	jsonFc.CarRegistration = fcDb.CarRegistration
	jsonFc.CreatedAt = fcDb.CreatedAt
	jsonFc.UpdatedAt = fcDb.UpdatedAt

	usersMap := make(map[int64]*model.User)

	pouredBy, err := UsersService.GetUserByID(int(fcDb.PouredBy.Int64))
	if err != nil {
		loger.ErrorLog.Println("Error getting FuelConsumption json from FuelConsumptionDb, error getting user by id: ", err)
		return nil, err
	}
	usersMap[int64(pouredBy.Id)] = pouredBy
	jsonFc.PouredBy = *pouredBy
	if fcDb.CreatedBy.Valid {
		createdBy, ok := usersMap[fcDb.CreatedBy.Int64]
		if !ok {
			createdBy, err = UsersService.GetUserByID(int(fcDb.CreatedBy.Int64))
			if err != nil {
				loger.ErrorLog.Println("Error getting FuelConsumption json from FuelConsumptionDb, error getting user by id: ", err)
				return nil, err
			}
			usersMap[int64(createdBy.Id)] = createdBy
		}
		jsonFc.CreatedBy = *createdBy
	}

	return &jsonFc, nil
}

func (s *fuelConsumptionService) CreateFuelConsumption(fuelCons model.FuelConsumption, userId int) (*model.FuelConsumption, error) {
	_, err := database.Client.Exec(`INSERT INTO fuel_consumption (date_consumption, fuel_type, liter, price, car_registration, poured_by, bill_file, created_by, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		fuelCons.DateConsumption.Time, fuelCons.FuelType, fuelCons.Liter, fuelCons.Price, fuelCons.CarRegistration, fuelCons.PouredBy.Id, fuelCons.BillFile, userId, time.Now(), time.Now())
	if err != nil {
		loger.ErrorLog.Println("Error creating FuelConsumption: ", err)
		return nil, err
	}

	//todo return that FuelConsumption if there is need fot that
	return nil, err
}

func (s *fuelConsumptionService) DeleteFuelConsumption(fsId int) error {
	_, err := database.Client.Exec(`DELETE from fuel_consumption WHERE id = $1`, fsId)
	if err != nil {
		loger.ErrorLog.Println("Error deleting fuel consumption: ", err)
		return err
	}

	return nil
}

func (s *fuelConsumptionService) GetFuelConsumptionCount(filter model.FuelConsumptionFilter) (int, error) {
	query := queryBuilderForFuelConsumptions(0, 0, filter, true, false)
	count := []int{}
	err := database.Client.Select(&count, query)
	if err != nil || len(count) == 0 {
		loger.ErrorLog.Println("Error getting count of fuel consumptions: ", err)
		return 0, err
	}

	return count[0], nil
}

func (s *fuelConsumptionService) UpdateFuelConsumption(fuelConsumption model.FuelConsumption) (*model.FuelConsumption, error) {
	_, err := database.Client.Exec(`UPDATE fuel_consumption SET car_registration = $1, liter = $2, price = $3, date_consumption = $4, poured_by = $5, fuel_type = $6, updated_at = $7 WHERE id = $8`,
		fuelConsumption.CarRegistration, fuelConsumption.Liter, fuelConsumption.Price, fuelConsumption.DateConsumption.Time, fuelConsumption.PouredBy.Id, fuelConsumption.FuelType, time.Now(), fuelConsumption.Id)
	if err != nil {
		loger.ErrorLog.Println("Error updating FuelConsumption: ", err)
		return nil, err
	}

	//todo return that FuelConsumption if there is need fot that
	return nil, err
}

func (s *fuelConsumptionService) CountSumPrice(filter model.FuelConsumptionFilter) (float64, error) {
	query := queryBuilderForFuelConsumptions(0, 0, filter, false, true)
	sum := []float64{}
	err := database.Client.Select(&sum, query)
	if err != nil || len(sum) == 0 {
		loger.ErrorLog.Println("Error getting sum prices for fuel consumptions: ", err)
		return 0, err
	}

	return sum[0], nil
}
