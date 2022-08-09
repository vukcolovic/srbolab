package model

import (
	"database/sql"
	"time"
)

type Certificate struct {
	Id                      int       `json:"id"`
	Brand                   string    `json:"brand"`
	TypeVehicle             string    `json:"type_vehicle"`
	Variant                 string    `json:"variant"`
	VersionVehicle          string    `json:"version_vehicle"`
	CommercialName          string    `json:"commercial_name"`
	EstimatedProductionYear int       `json:"estimated_production_year"`
	MaxMass                 int       `json:"max_mass"`
	RunningMass             int       `json:"running_mass"`
	Category                string    `json:"category"`
	BodyworkCode            string    `json:"bodywork_code"`
	AxlesTyresNum           string    `json:"axles_tyres_num"`
	Length                  int       `json:"length"`
	Width                   int       `json:"width"`
	Height                  int       `json:"height"`
	TyreWheel               string    `json:"tyre_wheel"`
	EngineCode              string    `json:"engine_code"`
	EngineCapacity          int       `json:"engine_capacity"`
	EnginePower             int       `json:"engine_power"`
	Fuel                    string    `json:"fuel"`
	PowerWeightRatio        string    `json:"power_weight_ratio"`
	SeatNumber              int       `json:"seat_number"`
	StandingNumber          int       `json:"standing_number"`
	MaxSpeed                int       `json:"max_speed"`
	GasLevel                string    `json:"gas_level"`
	MaxLadenMassAxios       string    `json:"max_laden_mass_axios"`
	NumberWvta              string    `json:"number_wvta"`
	PollutionCert           string    `json:"pollution_cert"`
	NoiseCert               string    `json:"noise_cert"`
	CouplingDeviceApproval  string    `json:"coupling_device_approval"`
	CreatedBy               User      `json:"created_by"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type CertificateDb struct {
	Id                      int            `db:"id"`
	Brand                   sql.NullString `db:"brand"`
	TypeVehicle             sql.NullString `db:"type_vehicle"`
	Variant                 sql.NullString `db:"variant"`
	VersionVehicle          sql.NullString `db:"version_vehicle"`
	CommercialName          sql.NullString `db:"commercial_name"`
	EstimatedProductionYear sql.NullInt64  `db:"estimated_production_year"`
	MaxMass                 sql.NullInt64  `db:"max_mass"`
	RunningMass             sql.NullInt64  `db:"running_mass"`
	Category                sql.NullString `db:"category"`
	BodyworkCode            sql.NullString `db:"bodywork_code"`
	AxlesTyresNum           sql.NullString `db:"axles_tyres_num"`
	Length                  sql.NullInt64  `db:"length"`
	Width                   sql.NullInt64  `db:"width"`
	Height                  sql.NullInt64  `db:"height"`
	TyreWheel               sql.NullString `db:"tyre_wheel"`
	EngineCode              sql.NullString `db:"engine_code"`
	EngineCapacity          sql.NullInt64  `db:"engine_capacity"`
	EnginePower             sql.NullInt64  `db:"engine_power"`
	Fuel                    sql.NullString `db:"fuel"`
	PowerWeightRatio        sql.NullString `db:"power_weight_ratio"`
	SeatNumber              sql.NullInt64  `db:"seat_number"`
	StandingNumber          sql.NullInt64  `db:"standing_number"`
	MaxSpeed                sql.NullInt64  `db:"max_speed"`
	GasLevel                sql.NullString `db:"gas_level"`
	MaxLadenMassAxios       sql.NullString `db:"max_laden_mass_axios"`
	NumberWvta              sql.NullString `db:"number_wvta"`
	PollutionCert           sql.NullString `db:"pollution_cert"`
	NoiseCert               sql.NullString `db:"noise_cert"`
	CouplingDeviceApproval  sql.NullString `db:"coupling_device_approval"`
	CreatedBy               sql.NullInt64  `db:"created_by"`
	CreatedAt               time.Time      `db:"created_at"`
	UpdatedAt               time.Time      `db:"updated_at"`
}

//filter Certificate
type CertificateFilter struct {
	TypeVehicle             string `json:"type_vehicle"`
	Variant                 string `json:"variant"`
	VersionVehicle          string `json:"version_vehicle"`
	EstimatedProductionYear int    `json:"estimated_production_year"`
	EngineCode              string `json:"engine_code"`
	EngineCapacity          int    `json:"engine_capacity"`
	EnginePower             int    `json:"engine_power"`
	Fuel                    string `json:"fuel"`
}