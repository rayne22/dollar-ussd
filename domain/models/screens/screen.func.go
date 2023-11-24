package screens

import (
	"dollar-ussd/configs"
	"log"
)

func (s Screen) CreateScreen() (Screen, error) {
	// posts to DB
	err := configs.GetDB().Model(&Screen{}).Create(&s).Error
	if err != nil {
		log.Println(err)
		return Screen{}, err
	}
	return s, nil
}

func (s Screen) GetAllScreens() ([]Screen, error) {
	screens := []Screen{}
	// posts to DB
	err := configs.GetDB().Model(&Screen{}).Preload("Options").Find(&screens).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return screens, nil
}
