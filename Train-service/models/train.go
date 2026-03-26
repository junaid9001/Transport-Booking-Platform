package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Station struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"size:100;not null"`
	Code      string    `gorm:"size:10;not null;unique"` // e.g. ERS
	City      string    `gorm:"size:50;not null"`
	CreatedAt time.Time
}

type Train struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TrainNumber string        `gorm:"size:10;uniqueIndex;not null"`
	TrainName   string        `gorm:"size:200;not null"`
	DaysOfWeek  pq.Int32Array `gorm:"type:integer[];not null"`
	IsActive    bool          `gorm:"default:true"`
	Stops       []TrainStop   `gorm:"foreignKey:TrainID"` // Has Many relation
	CreatedAt   time.Time
}

type TrainStop struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	TrainID       uuid.UUID `gorm:"type:uuid;not null;index"`
	StationID     uuid.UUID `gorm:"type:uuid;not null"`
	Station       Station   `gorm:"foreignKey:StationID"`
	StopSequence  int       `gorm:"not null"`        // 1, 2, 3...
	ArrivalTime   string    `gorm:"size:5;not null"` // HH:MM
	DepartureTime string    `gorm:"size:5;not null"` // HH:MM
	DayOffset     int       `gorm:"default:0"`       // 0=Same day, 1=Next day
	DistanceKm    int       `gorm:"not null;default:0"`
}
