package handlers

import (
	"net/http"
	"srbolabApp/service"
)

func ListIrregularityLevels(w http.ResponseWriter, r *http.Request) {
	levels, err := service.EnumerationService.GetAllIrregularityLevels()
	if err != nil {
		SetErrorResponse(w, err)
		return
	}

	SetSuccessResponse(w, levels)
}
