package health

import (
	"api/openapi"
	"encoding/json"
	"net/http"
)

type HealthCheckServer struct{}

func (s *HealthCheckServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(openapi.Message{
		Message: "OK",
	})
}
