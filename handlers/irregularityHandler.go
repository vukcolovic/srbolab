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

func CreateIrregularity(w http.ResponseWriter, r *http.Request) {
	var irregularity model.Irregularity
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&irregularity)
	if err != nil {
		loger.ErrorLog.Println("Error decoding Irregularity: ", err)
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

	createdIrregularity, err := service.IrregularityService.CreateIrregularity(irregularity, userId)
	if err != nil {
		loger.ErrorLog.Println("Error creating irregularity: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, createdIrregularity)
}

func ListIrregularities(w http.ResponseWriter, r *http.Request) {
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

	var filter model.IrregularityFilter
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&filter)
	if err != nil {
		loger.ErrorLog.Println("Unable to decode filter object: ", err)
		SetErrorResponse(w, err)
		return
	}

	irregularities, err := service.IrregularityService.GetAllIrregularities(skip, take, filter)
	if err != nil {
		loger.ErrorLog.Println("Error getting irregularities: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, irregularities)
}

func DeleteIrregularity(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		//loger.Error("error retrieving examination request param vin")
		SetErrorResponse(w, errors.New("error delete irregularity"))
		return
	}

	idInt, err := strconv.Atoi(id)

	err = service.IrregularityService.DeleteIrregularity(idInt)
	if err != nil {
		//loger.Error("error deleting examination request")
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, nil)
}

func CountIrregularities(w http.ResponseWriter, r *http.Request) {
	var filter model.IrregularityFilter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&filter)
	if err != nil {
		loger.ErrorLog.Println("Unable to decode filter object: ", err)
		SetErrorResponse(w, err)
		return
	}

	count, err := service.IrregularityService.GetIrregularitiesCount(filter)
	if err != nil {
		loger.ErrorLog.Println("Error getting irregularitie count: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, count)
}
