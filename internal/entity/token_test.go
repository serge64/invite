package entity_test

import (
	"testing"

	"github.com/serge64/invite/internal/entity"
)

func TestToken_GenerateToken(t *testing.T) {
	token := entity.GenerateToken()
	if len(token) != entity.CodeSize {
		t.Errorf("Expected value to be %d but got %d", entity.CodeSize, len(token))
	}
}

func TestToken_ValidateToken(t *testing.T) {
	testcases := []struct {
		name  string
		token entity.Token
		valid bool
	}{
		{
			name:  "valid",
			token: entity.GenerateToken(),
			valid: true,
		},
		{
			name:  "no valid: no equals by len",
			token: entity.Token(""),
		},
		{
			name:  "no valid: no match #1",
			token: entity.Token("!@#$%!@#$%!@#$%1"),
		},
		{
			name:  "no valid: no match #2",
			token: entity.Token("0123456789ABCDE@"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ok := entity.ValidateToken(tc.token)
			if tc.valid {
				if !ok {
					t.Errorf("Expected value to be 'true' but got 'false'")
				}
			} else {
				if ok {
					t.Errorf("Expected value to be 'false' but got 'true'")
				}
			}
		})
	}
}

func BenchmarkToken_GenerateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = entity.GenerateToken()
	}
}

func BenchmarkToken_ValidateToken(b *testing.B) {
	t := entity.GenerateToken()
	for i := 0; i < b.N; i++ {
		_ = entity.ValidateToken(t)
	}
}
