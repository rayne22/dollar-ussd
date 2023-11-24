package services

import (
	"dollar-ussd/domain/models/screens"
	"fmt"
	"log"
)

type Screen struct {
	Detail          string
	Type            string
	Position        string
	CurrentPosition string
}

func (scr Screen) CreateScreenDetail() ([]Screen, error) {

	s := screens.Screen{}

	scrns, err := s.GetAllScreens()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	finalScrn := []Screen{}

	for _, screen := range scrns {
		screens2 := []Screen{}
		det := ""
		for _, option := range screen.Options {

			if option.OptionNumber != nil {
				num := *option.OptionNumber
				det = fmt.Sprintf("%s \n\n %d) %s", det, num, option.OptionDetail)
			} else {
				det = fmt.Sprintf("%s \n\n %s", det, option.OptionDetail)
			}

		}
		scr2 := Screen{
			Detail:          det,
			Type:            string(screen.Type),
			Position:        screen.NextPosition,
			CurrentPosition: screen.Position,
		}

		screens2 = append(screens2, scr2)

		finalScrn = append(finalScrn, screens2...)

	}

	return finalScrn, nil
}
