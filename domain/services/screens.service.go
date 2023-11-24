package services

import (
	"dollar-ussd/domain/models/screens"
	"fmt"
	"log"
)

type Screen struct {
	Detail   string
	Type     string
	Position string
}

func (scr Screen) CreateScreenDetail() ([]Screen, error) {

	s := screens.Screen{}

	scrns, err := s.GetAllScreens()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	screens2 := []Screen{}
	for _, screen := range scrns {
		det := ""
		for _, option := range screen.Options {
			det = fmt.Sprintf("%s \n\n %s", det, option.OptionDetail)
		}

		scr := Screen{
			Detail:   det,
			Type:     string(screen.Type),
			Position: screen.NextPosition,
		}

		screens2 = append(screens2, scr)

	}

	return screens2, nil
}
