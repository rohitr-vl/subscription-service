package services

import (
	"encoding/json"
	"student-attendance/pkg/api_middleware"
)

func AllApiCounts() []byte {
	jsonVal, err := json.Marshal(api_middleware.Counter)
	if err != nil {
		return []byte(`{"error":"no counts found"}`)
	} else {
		return jsonVal
	}
}
