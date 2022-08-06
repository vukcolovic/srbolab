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

func GetCertificateByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	certIdParam, ok := vars["id"]
	if !ok {
		loger.ErrorLog.Println("missing parameter id")
		SetErrorResponse(w, errors.New("certificate not found"))
		return
	}

	certId, err := strconv.Atoi(certIdParam)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	user, err := service.CertificateService.GetCertificateById(certId)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, user)
}

func ListCertificates(w http.ResponseWriter, r *http.Request) {
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

	var filter model.CertificateFilter
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&filter)
	if err != nil {
		loger.ErrorLog.Println("Unable to decode filter object: ", err)
		SetErrorResponse(w, err)
		return
	}

	fuelCons, err := service.CertificateService.GetAllCertificates(skip, take, filter)
	if err != nil {
		loger.ErrorLog.Println("Error getting certificates: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, fuelCons)
}

func DeleteCertificate(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		loger.ErrorLog.Println("missing parameter id")
		SetErrorResponse(w, errors.New("error delete certificate"))
		return
	}

	idInt, err := strconv.Atoi(id)

	err = service.CertificateService.DeleteCertificate(idInt)
	if err != nil {
		loger.ErrorLog.Println("error delete certificate")
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, nil)
}

func CreateCertificate(w http.ResponseWriter, r *http.Request) {
	var certificate model.Certificate
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&certificate)
	if err != nil {
		loger.ErrorLog.Println("Error decoding certificate: ", err)
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

	createdCertificate, err := service.CertificateService.CreateCertificate(certificate, userId)
	if err != nil {
		loger.ErrorLog.Println("Error creating certificate: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, createdCertificate)
}

func CountCertificates(w http.ResponseWriter, r *http.Request) {
	var filter model.CertificateFilter
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&filter)
	if err != nil {
		loger.ErrorLog.Println("Unable to decode certificate filter object: ", err)
		SetErrorResponse(w, err)
		return
	}

	count, err := service.CertificateService.GetCertificatesCount(filter)
	if err != nil {
		loger.ErrorLog.Println("Error getting certificates count: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, count)
}

func UpdateCertificate(w http.ResponseWriter, r *http.Request) {
	var cert model.Certificate
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&cert)
	if err != nil {
		loger.ErrorLog.Println("Error decoding certificate: ", err)
		SetErrorResponse(w, err)
		return
	}

	updatedCertificate, err := service.CertificateService.UpdateCertificate(cert)
	if err != nil {
		loger.ErrorLog.Println("Error updating certificate: ", err)
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, updatedCertificate)
}
