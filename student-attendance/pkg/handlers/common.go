package handlers

import (
	"net/http"
	"student-attendance/pkg/services"
)

type CommonHandler struct {
}

func (comHand CommonHandler) GetApiCallCount(w http.ResponseWriter, r *http.Request) []byte {
	return services.AllApiCounts()
}
