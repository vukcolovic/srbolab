package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status            Status
	Data              string
	ErrorMessage      string
	MessageParameters []string
}

type Status string

const (
	ResponseSuccess = "success" //All went well, and (usually) some data was returned. (Required fields: status, data)
	ResponseError   = "error"   //An error occurred in processing the request, i.e. an exception was thrown. (Required: status, message	- Optionals: code, data)
)

func SetErrorResponse(w http.ResponseWriter, error error) {
	setResponse(w, Response{Status: ResponseError, ErrorMessage: error.Error()})
}

func SetSuccessResponse(w http.ResponseWriter, body interface{}) {
	data, err := json.Marshal(body)
	if err != nil {
		//loger.Instance().Error("ResponseError creating Response", loger.AdditionalFields{"ResponseError": err, "Data": body, "Status": ResponseSuccess})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SetSuccessResponseWithoutParsingBody(w, string(data))
}

func SetSuccessResponseWithoutParsingBody(w http.ResponseWriter, body string) {
	response := Response{Status: ResponseSuccess}
	response.Data = string(body)
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	setResponse(w, response)
}

func setResponse(w http.ResponseWriter, response Response) {
	enc := json.NewEncoder(w)
	err := enc.Encode(response)
	if err != nil {
		//loger.Instance().Error("ResponseError creating JSON Response", loger.AdditionalFields{"ResponseError": err, "Response": response, "Status": ResponseSuccess})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
