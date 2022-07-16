package service

import (
	"srbolabApp/database"
	"srbolabApp/loger"
	"srbolabApp/model"
	"srbolabApp/util"
	"time"
)

var (
	FuelConsumptionService fuelConsumptionServiceInterface = &fuelConsumptionService{}
)

type fuelConsumptionService struct {
}

type fuelConsumptionServiceInterface interface {
	GetAllFuelConsumptions(skip, take int) ([]model.FuelConsumption, error)
	//GetFuelConsumptionByID(id int) (*model.FuelConsumption, error)
	CreateFuelConsumption(fs model.FuelConsumption, userId int) (*model.FuelConsumption, error)
	//UpdateFuelConsumption(bool, model.FuelConsumption) (*model.FuelConsumption, error)
	DeleteFuelConsumption(int) error
	//GetFuelConsumptionCount() (int, error)
}

//func (s *userService) GetUserIDByToken(token string) (int, error) {
//	claims := &jwt.StandardClaims{}
//
//	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
//		return []byte("secret"), nil
//	})
//	if err != nil {
//		loger.ErrorLog.Println("Error getting user by token, error parse claims: ", err)
//		return 0, err
//	}
//
//	id, err := strconv.Atoi(claims.Id)
//	if err != nil {
//		loger.ErrorLog.Println("Error getting user by token: ", err)
//		return 0, err
//	}
//
//	return id, nil
//}

func (s *fuelConsumptionService) GetAllFuelConsumptions(skip, take int) ([]model.FuelConsumption, error) {
	fuelConsDb := []model.FuelConsumptionDb{}
	err := database.Client.Select(&fuelConsDb, `SELECT * FROM fuel_consumption ORDER BY id desc OFFSET $1 LIMIT $2`, skip, take)
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

func getJsonFuelConsumption(fcDb model.FuelConsumptionDb) (*model.FuelConsumption, error) {
	jsonFc := model.FuelConsumption{}

	jsonFc.Id = fcDb.Id
	jsonFc.FuelType = fcDb.FuelType
	jsonFc.DateConsumption = util.Date{fcDb.DateConsumption}
	jsonFc.Price = fcDb.Price
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
	_, err := database.Client.Exec(`INSERT INTO fuel_consumption (date_consumption, fuel_type, liter, price, car_registration, poured_by, created_by, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		fuelCons.DateConsumption, fuelCons.FuelType, fuelCons.Liter, fuelCons.Price, fuelCons.CarRegistration, fuelCons.PouredBy.Id, userId, time.Now(), time.Now())
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

//func (s *userService) GetUsersCount() (int, error) {
//	count := []int{}
//	err := database.Client.Select(&count, `SELECT COUNT(id) FROM users WHERE deleted = false`)
//	if err != nil || len(count) == 0 {
//		loger.ErrorLog.Println("Error getting count of users: ", err)
//		return 0, err
//	}
//
//	return count[0], nil
//}
