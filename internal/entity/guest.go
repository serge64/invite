package entity

type Guest struct {
	Token      Token  `json:"token,omitempty"`
	SingleMode bool   `json:"single_mode"`
	Name1      string `json:"name1"`
	Name2      string `json:"name2"`
	Status     Status `json:"status"`
	Choice1    string `json:"choice1"`
	Choice2    string `json:"choice2"`
}
