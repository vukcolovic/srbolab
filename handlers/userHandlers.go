package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"srbolabApp/errorUtils"
	"srbolabApp/loger"
	"srbolabApp/model"
	"srbolabApp/service"
	"strconv"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("unable to retrieve the just parsed code")
		SetErrorResponse(w, err)
		return
	}

	createdUser, err := service.UsersService.CreateUser(user)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, createdUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("unable to retrieve the just parsed code")
		SetErrorResponse(w, err)
		return
	}

	token, err := service.UsersService.Login(user)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	cookie := &http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
		Path:    "/",
	}

	http.SetCookie(w, cookie)
	SetSuccessResponse(w, token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now().Add(-time.Minute * 1),
		Path:    "/",
	}

	http.SetCookie(w, cookie)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.UsersService.GetAllUsers()
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, users)
}

func GetUserByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userIdParam, ok := vars["id"]
	if !ok {
		loger.ErrorLog.Println(errorUtils.ERR_MISSING_REQ_PARAM)
		SetErrorResponse(w, errors.New("user not found"))
		return
	}

	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	user, err := service.UsersService.GetUserByID(userId)
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, user)
}

//func (repo *DBRepo) CreateTemplate(w http.ResponseWriter, req *http.Request) {
//	payloadTemplate := internal.ExaminationRequest{}
//	decoder := json.NewDecoder(req.Body)
//	err := decoder.Decode(&payloadTemplate)
//	if err != nil {
//		loger.Error("unable to retrieve the just parsed code")
//		SetErrorResponse(w, errorUtils.New("unable to retrieve the just parsed code"))
//		return
//	}
//
//	template, err := db.Create(&payloadTemplate)
//	if err != nil {
//		loger.Error("error creating template")
//		SetErrorResponse(w, errorUtils.New("error creating template"))
//		return
//	}
//
//	SetSuccessResponse(w, template)
//}
//
//func (repo *DBRepo) UpdateTemplate(w http.ResponseWriter, req *http.Request) {
//	payload := internal.ExaminationRequest{}
//	decoder := json.NewDecoder(req.Body)
//	err := decoder.Decode(&payload)
//	if err != nil {
//		loger.Error("unable to retrieve the just parsed code")
//		SetErrorResponse(w, errorUtils.New("unable to retrieve the just parsed code"))
//		return
//	}
//
//	template, err := db.Update(&payload)
//	if err != nil {
//		loger.Error("error updating examination request")
//		SetErrorResponse(w, errorUtils.New("error updating examination request"))
//		return
//	}
//
//	SetSuccessResponse(w, template)
//}
//
//func (repo *DBRepo) DeleteTemplate(w http.ResponseWriter, req *http.Request) {
//	vars := mux.Vars(req)
//	templateID, ok := vars["vin"]
//	if !ok {
//		loger.Error("error retrieving examination request param vin")
//		SetErrorResponse(w, errorUtils.New("examination request not found"))
//		return
//	}
//
//	err := db.DeleteTemplate(templateID)
//	if err != nil {
//		loger.Error("error deleting examination request")
//		SetErrorResponse(w, errorUtils.New("erro deleting examination request"))
//		return
//	}
//
//	SetSuccessResponse(w, nil)
//}
