package entity

type Status uint8

const (
	StatusNoResponded Status = iota
	StatusPositive
	StatusNegative

	statusNoRespondedStr string = "не определен"
	statusPositiveStr    string = "приду"
	statusNegativeStr    string = "не приду"
)

func StatusValid(value string) bool {
	return value == "true" || value == "false"
}

func ConvertToStatus(value string) Status {
	if value == "true" {
		return StatusPositive
	}
	return StatusNegative
}

func (s Status) String() string {
	switch s {
	case StatusPositive:
		return statusPositiveStr
	case StatusNegative:
		return statusNegativeStr
	default:
		return statusNoRespondedStr
	}
}
