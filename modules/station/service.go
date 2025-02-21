package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/HaikalRFadhilahh/api-mrt-go-bwa/common/client"
)

// Interface For Service Station
type Service interface {
	GetAllStation() (response []StationResponse, err error)
	CheckSheduleByStation(id string) (response []ScheduleResponse, err error)
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

func (s *service) CheckSheduleByStation(id string) (response []ScheduleResponse, err error) {
	// URL Endpoint Config
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	// Hit URL
	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return nil, err
	}

	// Create Data For Station
	var schedule []Schedule
	if err := json.Unmarshal(byteResponse, &schedule); err != nil {
		return nil, err
	}

	// Create To Take Data With Same ID Request
	var validStation Schedule

	// Selected Data With ID Station
	for _, i := range schedule {
		if i.StationId == id {
			validStation = i
		}
	}

	// Check Valid Station Exist Data
	if validStation.StationId == "" {
		return nil, errors.New("data station not found")
	}

	// Convert Data Struct to Array Response
	// Mapping Schedule Bundaran HI
	bundaranHITime, err := ConvertTimeStringToArrayTime(validStation.ScheduleBundaranHI)
	if err != nil {
		return nil, err
	}

	// Mapping Schedule Lebak Bulus
	lebakBulusTime, err := ConvertTimeStringToArrayTime(validStation.ScheduleLebakBulus)
	if err != nil {
		return nil, err
	}

	// Mapping Data To Response Schedule Handler DTO
	return MappingStationTimeToResponseSchedule(ScheduleStationTime{StationName: "Stasiun Bundaran HI", Time: bundaranHITime}, ScheduleStationTime{StationName: "Stasiun Lebak Bulus Grab", Time: lebakBulusTime})
}

// Function To Mapping Response Schedule
func MappingStationTimeToResponseSchedule(data ...ScheduleStationTime) (response []ScheduleResponse, err error) {
	// Mapping Variadic Data Paramater
	for _, d := range data {
		for _, i := range d.Time {
			if i.Format("15:04") > time.Now().Format("15:04") {
				response = append(response, ScheduleResponse{
					StationName: d.StationName,
					Time:        i.Format("15:04"),
				})
			}
		}
	}

	return response, nil
}

// Function Helper To Mapping Data Time and trim split String
func ConvertTimeStringToArrayTime(waktu string) (response []time.Time, err error) {
	if waktu == "" {
		return
	}
	for _, i := range strings.Split(strings.ReplaceAll(waktu, " ", ""), ",") {
		parsedTime, err := time.Parse("15:04", i)
		if err != nil {
			return nil, err
		}
		response = append(response, parsedTime)
	}
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
