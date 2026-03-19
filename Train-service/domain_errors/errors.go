package domainerrors

import "errors"

var (
	ErrTrainNotFound       = errors.New("train not found")
	ErrNoSeatsAvailable    = errors.New("no confirmed seats available")
	ErrSeatAlreadyLocked   = errors.New("seat is temporarily held by another user")
	ErrSeatAlreadyBooked   = errors.New("seat is already booked")
	ErrBookingNotFound     = errors.New("booking not found")
	ErrBookingNotConfirmed = errors.New("booking is not yet confirmed")
	ErrUnauthorized        = errors.New("you do not own this booking")
	ErrCannotCancel        = errors.New("only pending or confirmed bookings can be cancelled")
)
