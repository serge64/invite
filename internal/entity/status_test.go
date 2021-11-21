package entity_test

import (
	"testing"

	"github.com/serge64/invite/internal/entity"
)

func TestStatus_ValidateStatus(t *testing.T) {
	testcases := []struct {
		name     string
		status   string
		expected bool
	}{
		{
			name:     "valid true",
			status:   "true",
			expected: true,
		},
		{
			name:     "valid false",
			status:   "false",
			expected: true,
		},
		{
			name:   "empty status",
			status: "",
		},
		{
			name:   "no valid status",
			status: "novalid",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			b := entity.ValidateStatus(tc.status)
			if tc.expected != b {
				t.Errorf("Expected value to be '%t' but got '%t'", tc.expected, b)
			}
		})
	}
}

func TestStatus_ConvertToStatus(t *testing.T) {
	testcases := []struct {
		name     string
		status   string
		expected entity.Status
	}{
		{
			name:     "true",
			status:   "true",
			expected: entity.StatusPositive,
		},
		{
			name:     "false",
			status:   "false",
			expected: entity.StatusNegative,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			s := entity.ConvertToStatus(tc.status)
			if tc.expected != s {
				t.Errorf("Expected value to be '%d' but got '%d'", tc.expected, s)
			}
		})
	}
}
