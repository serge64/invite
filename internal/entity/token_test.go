package entity_test

import (
	"testing"

	"github.com/serge64/invite/internal/entity"
)

func TestToken_IsValid(t *testing.T) {
	testcases := []struct {
		name  string
		token entity.Token
		valid bool
	}{
		{
			name:  "valid",
			token: entity.NewToken(),
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
			ok := tc.token.IsValid()
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

func TestToken_String(t *testing.T) {
	token := entity.NewToken()
	if token.String() == "" {
		t.Error("Expected value to be no empty but got empty")
	}
}

func BenchmarkToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = entity.NewToken().IsValid()
	}
}
