package _interface

type Screen struct {
	ScreenID string `json:"id"`
	Details  string `json:"details"`
	//Options  Option `json:"options"`
	InputType string `json:"input_type"`
	Position  string `json:"position"`
}

type Option struct {
	OptionNumber int    `json:"option_number"`
	OptionDetail string `json:"option_detail"`
}

type InputDetail struct {
	Selection string `json:"selection"`
	Screen    string `json:"screen"`
	Type      string `json:"type"`
}
