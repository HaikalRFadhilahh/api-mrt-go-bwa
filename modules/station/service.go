package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HaikalRFadhilahh/api-mrt-go-bwa/common/client"
)

// Interface For Service Station
type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

// Struct Impl Interface
type service struct {
	client *http.Client
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	// URL Endpoint Config
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	// Create Data For Station
	var stations []Station
	if err := json.Unmarshal(byteResponse, &stations); err != nil {
		return nil, err
	}

	// Loop Byte Mashal Json Data
	for _, i := range stations {
		response = append(response, StationResponse{
			Id:   i.Id,
			Name: i.Name,
		})
	}

	// Return
	return 
}

// Factory Service
func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
