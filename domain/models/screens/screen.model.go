package screens

import (
	"dollar-ussd/domain/enums"
	"gorm.io/gorm"
)

type Screen struct {
	gorm.Model
	Options      []Option        `json:"options"`
	Type         enums.InputType `json:"type"`
	Position     string          `json:"position"`
	NextPosition string          `json:"next_position"`
}

type Option struct {
	gorm.Model
	ScreenID     uint   `json:"screen_id"`
	OptionNumber *int   `json:"option_number"`
	OptionDetail string `json:"option_detail"`
}
