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
	CertificateService certificateServiceInterface = &certificateService{}
)

type certificateService struct {
}

type certificateServiceInterface interface {
	GetAllCertificates(skip, take int, filter model.CertificateFilter) ([]model.Certificate, error)
	GetCertificateById(id int) (*model.Certificate, error)
	CreateCertificate(cert model.Certificate, userId int) (*model.Certificate, error)
	UpdateCertificate(model.Certificate) (*model.Certificate, error)
	DeleteCertificate(int) error
	GetCertificatesCount(filter model.CertificateFilter) (int, error)
	GetCertificatePdfReportByIdAndWin(id int, win string) ([]byte, error)
}

func (s *certificateService) GetCertificateById(id int) (*model.Certificate, error) {
	certDb := []model.CertificateDb{}
	err := database.Client.Select(&certDb, `SELECT * FROM certificates WHERE id = $1`, id)
	if err != nil || len(certDb) == 0 {
		loger.ErrorLog.Println("Error getting certificate by id: ", err)
		return nil, err
	}

	cert, err := getJsonCertificate(certDb[0])
	if err != nil {
		loger.ErrorLog.Println("Error getting certificate, error getting json from db: ", err)
		return nil, err
	}

	return cert, nil
}

func (s *certificateService) GetAllCertificates(skip, take int, filter model.CertificateFilter) ([]model.Certificate, error) {
	query := queryBuilderForCertificates(skip, take, filter, false)
	certsDb := []model.CertificateDb{}
	err := database.Client.Select(&certsDb, query)
	if err != nil {
		loger.ErrorLog.Println("Error getting all certificates: ", err)
		return nil, err
	}

	result := []model.Certificate{}
	for _, c := range certsDb {
		cert, err := getJsonCertificate(c)
		if err != nil {
			loger.ErrorLog.Println("Error getting all certificates, error getting json from db: ", err)
			return nil, err
		}

		result = append(result, *cert)
	}

	return result, nil
}

func queryBuilderForCertificates(skip, take int, filter model.CertificateFilter, isCount bool) string {
	var query string
	if isCount {
		query = `SELECT count(*) FROM certificates WHERE `
	} else {
		query = `SELECT * FROM certificates WHERE `
	}

	if filter.TypeVehicle != "" {
		query = query + ` type_vehicle LIKE ` + "'%" + filter.TypeVehicle + "%'" + ` AND `
	}
	if filter.Variant != "" {
		query = query + ` variant LIKE ` + "'%" + filter.Variant + "'%" + ` AND `
	}
	if filter.VersionVehicle != "" {
		query = query + ` version_vehicle LIKE ` + "'%" + filter.VersionVehicle + "'%" + ` AND `
	}
	if filter.EstimatedProductionYear != 0 {
		query = query + ` CAST(estimated_production_year AS TEXT) LIKE ` + "'" + strconv.Itoa(filter.EstimatedProductionYear) + "'%" + ` AND `
	}
	if filter.EngineCode != "" {
		query = query + ` CAST(engine_code AS TEXT) LIKE ` + "'%" + filter.EngineCode + "'%" + ` AND `
	}
	if filter.EngineCapacity != 0 {
		query = query + ` CAST(engine_capacity AS TEXT) LIKE ` + "'" + strconv.Itoa(filter.EngineCapacity) + "'%" + ` AND `
	}
	if filter.EnginePower != 0 {
		query = query + ` CAST(engine_power AS TEXT) LIKE ` + "'" + strconv.Itoa(filter.EnginePower) + "'%" + ` AND `
	}
	if filter.Fuel != "" {
		query = query + ` fuel LIKE ` + "'%" + filter.Fuel + "'%" + ` AND `
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

func getJsonCertificate(certDb model.CertificateDb) (*model.Certificate, error) {
	cert := model.Certificate{}

	cert.Id = certDb.Id
	cert.Brand = certDb.Brand.String
	cert.TypeVehicle = certDb.TypeVehicle.String
	cert.Variant = certDb.Variant.String
	cert.VersionVehicle = certDb.VersionVehicle.String
	cert.CommercialName = certDb.CommercialName.String
	cert.EstimatedProductionYear = certDb.EstimatedProductionYear.String
	cert.MaxMass = certDb.MaxMass.String
	cert.RunningMass = certDb.RunningMass.String
	cert.Category = certDb.Category.String
	cert.BodyworkCode = certDb.BodyworkCode.String
	cert.AxlesTyresNum = certDb.AxlesTyresNum.String
	cert.Length = certDb.Length.String
	cert.Width = certDb.Width.String
	cert.Height = certDb.Height.String
	cert.TyreWheel = certDb.TyreWheel.String
	cert.EngineCode = certDb.EngineCode.String
	cert.EngineCapacity = certDb.EngineCapacity.String
	cert.EnginePower = certDb.EnginePower.String
	cert.Fuel = certDb.Fuel.String
	cert.PowerWeightRatio = certDb.PowerWeightRatio.String
	cert.SeatNumber = certDb.SeatNumber.String
	cert.StandingNumber = certDb.StandingNumber.String
	cert.MaxSpeed = certDb.MaxSpeed.String
	cert.GasLevel = certDb.GasLevel.String
	cert.MaxLadenMassAxios = certDb.MaxLadenMassAxios.String
	cert.NumberWvta = certDb.NumberWvta.String
	cert.PollutionCert = certDb.PollutionCert.String
	cert.NoiseCert = certDb.NoiseCert.String
	cert.CouplingDeviceApproval = certDb.CouplingDeviceApproval.String
	cert.CreatedAt = certDb.CreatedAt
	cert.UpdatedAt = certDb.UpdatedAt

	createdBy, err := UsersService.GetUserByID(int(certDb.CreatedBy.Int64))
	if err != nil {
		loger.ErrorLog.Println("Error getting FuelConsumption json from FuelConsumptionDb, error getting user by id: ", err)
		return nil, err
	}
	cert.CreatedBy = *createdBy

	return &cert, nil
}

func (s *certificateService) CreateCertificate(cert model.Certificate, userId int) (*model.Certificate, error) {
	_, err := database.Client.Exec(`INSERT INTO certificates 
    (brand, type_vehicle, variant, version_vehicle, commercial_name, estimated_production_year, max_mass, running_mass, category, bodywork_code, axles_tyres_num, length, width, height, tyre_wheel, engine_code, engine_capacity, engine_power, fuel, power_weight_ratio, seat_number, standing_number, max_speed, gas_level, max_laden_mass_axios, number_wvta, pollution_cert, noise_cert, coupling_device_approval, created_by, created_at, updated_at) 
    	VALUES 
   ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32)`,
		cert.Brand, cert.TypeVehicle, cert.Variant, cert.VersionVehicle, cert.CommercialName, cert.EstimatedProductionYear, cert.MaxMass, cert.RunningMass, cert.Category, cert.BodyworkCode, cert.AxlesTyresNum, cert.Length, cert.Width, cert.Height, cert.TyreWheel, cert.EngineCode, cert.EngineCapacity, cert.EnginePower, cert.Fuel, cert.PowerWeightRatio, cert.SeatNumber, cert.StandingNumber, cert.MaxSpeed, cert.GasLevel, cert.MaxLadenMassAxios, cert.NumberWvta, cert.PollutionCert, cert.NoiseCert, cert.CouplingDeviceApproval, userId, time.Now(), time.Now())
	if err != nil {
		loger.ErrorLog.Println("Error creating certificate: ", err)
		return nil, err
	}

	//todo return that FuelConsumption if there is need fot that
	return nil, err
}

func (s *certificateService) DeleteCertificate(id int) error {
	_, err := database.Client.Exec(`DELETE from certificates WHERE id = $1`, id)
	if err != nil {
		loger.ErrorLog.Println("Error deleting certificate: ", err)
		return err
	}

	return nil
}

func (s *certificateService) GetCertificatesCount(filter model.CertificateFilter) (int, error) {
	query := queryBuilderForCertificates(0, 0, filter, true)
	count := []int{}
	err := database.Client.Select(&count, query)
	if err != nil || len(count) == 0 {
		loger.ErrorLog.Println("Error getting count of certificates: ", err)
		return 0, err
	}

	return count[0], nil
}

func (s *certificateService) UpdateCertificate(cert model.Certificate) (*model.Certificate, error) {
	_, err := database.Client.Exec(`UPDATE certificates SET brand = $1, type_vehicle = $2, variant = $3, version_vehicle = $4, commercial_name = $5, estimated_production_year = $6, max_mass = $7, running_mass = $8, category = $9, bodywork_code = $10, axles_tyres_num = $11, length = $12, width = $13, height = $14, tyre_wheel = $15, engine_code = $16, engine_capacity = $17, engine_power = $18, fuel = $19, power_weight_ratio = $20, seat_number = $21, standing_number = $22, max_speed = $23, gas_level = $24, max_laden_mass_axios = $25, number_wvta = $26, pollution_cert = $27, noise_cert = $28,  coupling_device_approval = $29, updated_at = $30 WHERE id = $31`,
		cert.Brand, cert.TypeVehicle, cert.Variant, cert.VersionVehicle, cert.CommercialName, cert.EstimatedProductionYear, cert.MaxMass, cert.RunningMass, cert.Category, cert.BodyworkCode, cert.AxlesTyresNum, cert.Length, cert.Width, cert.Height, cert.TyreWheel, cert.EngineCode, cert.EngineCapacity, cert.EnginePower, cert.Fuel, cert.PowerWeightRatio, cert.SeatNumber, cert.StandingNumber, cert.MaxSpeed, cert.GasLevel, cert.MaxLadenMassAxios, cert.NumberWvta, cert.PollutionCert, cert.NoiseCert, cert.CouplingDeviceApproval, time.Now(), cert.Id)
	if err != nil {
		loger.ErrorLog.Println("Error updating FuelConsumption: ", err)
		return nil, err
	}

	//todo return that FuelConsumption if there is need fot that
	return nil, err
}

func (s *certificateService) GetCertificatePdfReportByIdAndWin(id int, win string) ([]byte, error) {
	cert, err := s.GetCertificateById(id)
	if err != nil {
		loger.ErrorLog.Println("Error getting certificate by id: ", err)
		return nil, err
	}

	return PdfService.CreateCertificate(cert, win)
}
