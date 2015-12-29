package model

import (
	"time"
)

type Identifiable struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Base struct{}
