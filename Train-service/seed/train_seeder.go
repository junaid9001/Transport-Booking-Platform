package seed

import (
	"encoding/json"
	"os"

	"github.com/lib/pq"
	"github.com/nabeel-mp/tripneo/train-service/models"
	"gorm.io/gorm"
)

func SeedStations(tx *gorm.DB) error {
	bytes, err := os.ReadFile("data/stations.json")
	if err != nil {
		return err
	}

	var stations []models.Station
	if err := json.Unmarshal(bytes, &stations); err != nil {
		return err
	}

	for _, s := range stations {
		if err := tx.Where("code = ?", s.Code).FirstOrCreate(&s).Error; err != nil {
			return err
		}
	}
	return nil
}

func SeedTrains(tx *gorm.DB) error {
	bytes, err := os.ReadFile("data/trains.json")
	if err != nil {
		return err
	}

	var rawTrains []struct {
		TrainNumber string  `json:"train_number"`
		TrainName   string  `json:"train_name"`
		DaysOfWeek  []int32 `json:"days_of_week"`
		IsActive    bool    `json:"is_active"`
		Stops       []struct {
			StationCode string `json:"station_code"`
			Sequence    int    `json:"sequence"`
			Arrival     string `json:"arrival"`
			Departure   string `json:"departure"`
			DayOffset   int    `json:"day_offset"`
			Distance    int    `json:"distance"`
		} `json:"stops"`
	}

	if err := json.Unmarshal(bytes, &rawTrains); err != nil {
		return err
	}

	for _, r := range rawTrains {
		train := models.Train{
			TrainNumber: r.TrainNumber,
			TrainName:   r.TrainName,
			DaysOfWeek:  pq.Int32Array(r.DaysOfWeek),
			IsActive:    r.IsActive,
		}
		if err := tx.Where("train_number = ?", train.TrainNumber).FirstOrCreate(&train).Error; err != nil {
			return err
		}

		for _, stop := range r.Stops {
			var station models.Station
			tx.Where("code = ?", stop.StationCode).First(&station)

			trainStop := models.TrainStop{
				TrainID:       train.ID,
				StationID:     station.ID,
				StopSequence:  stop.Sequence,
				ArrivalTime:   stop.Arrival,
				DepartureTime: stop.Departure,
				DayOffset:     stop.DayOffset,
				DistanceKm:    stop.Distance,
			}
			tx.Where("train_id = ? AND stop_sequence = ?", train.ID, stop.Sequence).FirstOrCreate(&trainStop)
		}
	}
	return nil
}
