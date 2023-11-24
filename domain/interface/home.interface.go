package _interface

import (
	"dollar-ussd/domain/services"
	"log"
)

func (s *Screen) FetchScreen() (Screen, error) {

	return Screen{}, nil
}

func (s *Screen) FetchScreens() ([]Screen, error) {
	scr := services.Screen{}

	screens := []Screen{}

	scrns, err := scr.CreateScreenDetail()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, screen := range scrns {
		scrn := Screen{
			ScreenID:  screen.Position,
			InputType: screen.Type,
			Details:   screen.Detail,
		}

		screens = append(screens, scrn)
	}
	return screens, nil
}

func (s *Screen) SelectionProcessing(selection InputDetail) (Screen, error) {

	resp, err := s.FetchScreens()
	if err != nil {
		return Screen{}, err
	}

	nxtScreen := Screen{}

	for _, scrn := range resp {

		if scrn.ScreenID == selection.Selection {
			nxtScreen = scrn
			break
		}
	}

	return nxtScreen, nil
}
