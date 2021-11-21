package entity

type Status uint8

const (
	StatusNoResponded Status = iota
	StatusPositive
	StatusNegative
)

func ValidateStatus(value string) bool {
	return value == "true" || value == "false"
}

func ConvertToStatus(value string) Status {
	if value == "true" {
		return StatusPositive
	}
	return StatusNegative
}
