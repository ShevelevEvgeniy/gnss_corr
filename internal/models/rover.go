package models

import "github.com/google/uuid"

type Rover struct {
	RoverID uuid.UUID `db:"id"`
	Name    string    `db:"name"`
}
