package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"srbolabApp/model"
	"srbolabApp/service"
	"strconv"
)

func CreateIrregularity(w http.ResponseWriter, r *http.Request) {
	var irregularity model.Irregularity
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&irregularity)
	if err != nil {
		log.Println("unable to retrieve the parsed code")
		SetErrorResponse(w, err)
		return
	}

	token, err := GetTokenFromRequest(r)
	if err != nil {
		log.Println("unable to retrieve token from request")
		SetErrorResponse(w, err)
		return
	}

	userId, err := service.UsersService.GetUserIDByToken(token)
	if err != nil {
		log.Println("error getting user from token")
		SetErrorResponse(w, err)
		return
	}

	createdIrregularity, err := service.IrregularityService.CreateIrregularity(irregularity, userId)
	if err != nil {
		log.Println("error creating irregularity")
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
		SetErrorResponse(w, err)
		return
	}

	takeParam := queryParams["take"][0]
	take, err := strconv.Atoi(takeParam)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}
	irregularities, err := service.IrregularityService.GetAllIrregularities(skip, take)
	if err != nil {
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
