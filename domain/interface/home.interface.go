package _interface

import (
	"dollar-ussd/domain/services"
	"dollar-ussd/utils"
	"fmt"
	"log"
	"strings"
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
			Position:  screen.CurrentPosition,
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

		getPosition := strings.Split(selection.Screen, ",")

		for _, v := range getPosition {
			finalPosition := strings.Split(v, "-")

			if len(finalPosition) > 1 {
				foundSelection := utils.ValueIsEqual(finalPosition, selection.Selection)
				if foundSelection == selection.Selection {
					if scrn.Position == finalPosition[1] {

						nxtScreen = scrn

						fmt.Println("WHAT IS JO>>>>>>>>>>>>>>?", scrn.Position, finalPosition[1])
						break
					}
				}
			} else {
				if scrn.Position == finalPosition[0] {

					nxtScreen = scrn
					//fmt.Println("WHAT IS JO>>>>>>>>>>>>>>?", scrn.Position, finalPosition[0])
					break
				}
			}

		}

	}

	return nxtScreen, nil
}
