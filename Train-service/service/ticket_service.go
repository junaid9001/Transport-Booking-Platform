package service

import (
	"fmt"

	domainerrors "github.com/nabeel-mp/tripneo/train-service/domain_errors"
	"github.com/nabeel-mp/tripneo/train-service/repository"
	"github.com/nabeel-mp/tripneo/train-service/utils"
)

func GetTicket(bookingID, userID string) (interface{}, error) {
	// 1. Verify booking ownership and preload the schedule and train
	booking, err := repository.GetBookingByID(bookingID)
	if err != nil {
		return nil, err
	}
	if booking.UserID != userID {
		return nil, domainerrors.ErrUnauthorized
	}
	if booking.Status != "CONFIRMED" {
		return nil, domainerrors.ErrBookingNotConfirmed
	}

	// 2. Fetch ticket record
	ticket, err := repository.GetTicketByBookingID(bookingID)
	if err != nil {
		return nil, err
	}

	// 3. Fetch passengers
	passengers, err := repository.GetPassengersByBookingID(bookingID)
	if err != nil {
		return nil, err
	}

	// NEW LOGIC: Use the booking's specific stations
	// In the segment-based model, a booking belongs to a specific pair of stations
	// which may be different from the train's start and end points.
	fromStation := booking.FromStation.Name // Assuming these are preloaded/available in your model
	toStation := booking.ToStation.Name

	type TicketResponse struct {
		TicketNumber string      `json:"ticket_number"`
		PNR          string      `json:"pnr"`
		QRCodeURL    string      `json:"qr_code_url"`
		TrainName    string      `json:"train_name"`
		TrainNumber  string      `json:"train_number"`
		From         string      `json:"from"`
		To           string      `json:"to"`
		DepartureAt  interface{} `json:"departure_at"`
		ArrivalAt    interface{} `json:"arrival_at"`
		Class        string      `json:"class"`
		Passengers   interface{} `json:"passengers"`
		Status       string      `json:"status"`
	}

	return TicketResponse{
		TicketNumber: ticket.TicketNumber,
		PNR:          booking.PNR,
		QRCodeURL:    ticket.QRCodeURL,
		TrainName:    booking.TrainSchedule.Train.TrainName,
		TrainNumber:  booking.TrainSchedule.Train.TrainNumber,
		From:         fromStation,
		To:           toStation,
		DepartureAt:  booking.DepartureTime, // Use the specific station departure time
		ArrivalAt:    booking.ArrivalTime,   // Use the specific station arrival time
		Class:        booking.SeatClass,
		Passengers:   passengers,
		Status:       booking.Status,
	}, nil
}

// VerifyTicket validates a QR ticket's HMAC token.
func VerifyTicket(bookingID, token string) (interface{}, error) {
	booking, err := repository.GetBookingByID(bookingID)
	if err != nil {
		return nil, err
	}

	// Validate HMAC token to ensure the QR wasn't faked
	valid := utils.VerifyQRToken(bookingID, token)
	if !valid {
		return nil, fmt.Errorf("invalid or tampered QR token")
	}

	passengers, err := repository.GetPassengersByBookingID(bookingID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"valid":      true,
		"booking_id": bookingID,
		"pnr":        booking.PNR,
		"train":      booking.TrainSchedule.Train.TrainNumber,
		"route":      fmt.Sprintf("%s -> %s", booking.FromStation.Code, booking.ToStation.Code),
		"status":     booking.Status,
		"passengers": passengers,
	}, nil
}
