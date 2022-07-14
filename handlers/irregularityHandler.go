package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"srbolabApp/errorUtils"
	"srbolabApp/loger"
	"srbolabApp/model"
	"srbolabApp/service"
	"strconv"
)

func GetIrregularityByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	irrgularityIdParam, ok := vars["id"]
	if !ok {
		loger.ErrorLog.Println(errorUtils.ERR_MISSING_REQ_PARAM)
		SetErrorResponse(w, errors.New("irregularity not found"))
		return
	}

	irregularityId, err := strconv.Atoi(irrgularityIdParam)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	user, err := service.IrregularityService.GetIrregularityByID(irregularityId)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, user)
}

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

func UpdateIrregularities(w http.ResponseWriter, r *http.Request) {
	var irregularity model.Irregularity
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&irregularity)
	if err != nil {
		loger.ErrorLog.Println("Error decoding Irregularity: ", err)
		SetErrorResponse(w, err)
		return
	}

	updatedIrregularity, err := service.IrregularityService.UpdateIrregularity(irregularity)
	if err != nil {
		loger.ErrorLog.Println("Error updating irregularity: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, updatedIrregularity)
}

func ChangeCorrected(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	if len(queryParams["corrected"]) == 0 {
		//loger.Error("error retrieving examination request param vin")
		SetErrorResponse(w, errors.New("error changing corrected"))
		return
	}
	corrected := queryParams["corrected"][0]

	correctedBool, _ := strconv.ParseBool(corrected)

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

	err = service.IrregularityService.ChangeCorrected(irregularity, correctedBool, userId)
	if err != nil {
		loger.ErrorLog.Println("Error changing corrected irregularity: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, nil)
}
