package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"srbolabApp/loger"
	"srbolabApp/model"
	"srbolabApp/service"
	"strconv"
)

func ListFuelConsumptions(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	skipParam := queryParams["skip"][0]
	skip, err := strconv.Atoi(skipParam)
	if err != nil {
		loger.ErrorLog.Println("Unable to retrieve skip param: ", err)
		SetErrorResponse(w, err)
		return
	}
	takeParam := queryParams["take"][0]
	take, err := strconv.Atoi(takeParam)
	if err != nil {
		loger.ErrorLog.Println("Unable to retrieve take param: ", err)
		SetErrorResponse(w, err)
		return
	}

	//var filter model.IrregularityFilter
	//decoder := json.NewDecoder(r.Body)
	//err = decoder.Decode(&filter)
	//if err != nil {
	//	loger.ErrorLog.Println("Unable to decode filter object: ", err)
	//	SetErrorResponse(w, err)
	//	return
	//}

	fuelCons, err := service.FuelConsumptionService.GetAllFuelConsumptions(skip, take)
	if err != nil {
		loger.ErrorLog.Println("Error getting fuel consumtpions: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, fuelCons)
}

func DeleteFuelConsumption(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		//loger.Error("error retrieving examination request param vin")
		SetErrorResponse(w, errors.New("error delete fuel consumtpion"))
		return
	}

	idInt, err := strconv.Atoi(id)

	err = service.FuelConsumptionService.DeleteFuelConsumption(idInt)
	if err != nil {
		//loger.Error("error deleting examination request")
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, nil)
}

func CreateFuelConsumption(w http.ResponseWriter, r *http.Request) {
	var fuelConsumption model.FuelConsumption
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&fuelConsumption)
	if err != nil {
		loger.ErrorLog.Println("Error decoding FuelConsumption: ", err)
		SetErrorResponse(w, err)
		return
	}

	token, err := GetTokenFromRequest(r)
	if err != nil {
		loger.ErrorLog.Println("Unable to retrieve token from requeste: ", err)
		SetErrorResponse(w, err)
		return
	}

	userId, err := service.UsersService.GetUserIDByToken(token)
	if err != nil {
		loger.ErrorLog.Println("Error getting user from token: ", err)
		SetErrorResponse(w, err)
		return
	}

	createdFuelConsumption, err := service.FuelConsumptionService.CreateFuelConsumption(fuelConsumption, userId)
	if err != nil {
		loger.ErrorLog.Println("Error creating FuelConsumption: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, createdFuelConsumption)
}
