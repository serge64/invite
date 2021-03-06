package usecase_test

import (
	"log"
	"strings"
	"testing"

	"github.com/serge64/invite/internal/entity"
	"github.com/serge64/invite/internal/usecase"
)

func TestGuest_GuestToMessage(t *testing.T) {
	token := entity.GenerateToken()
	url := "http://localhost/invite/"
	link := url + string(token)

	testcases := []struct {
		name     string
		guest    entity.Guest
		expected string
	}{
		{
			name: "single mode off - positivie",
			guest: entity.Guest{
				Token:   token,
				Name1:   "Гость1",
				Name2:   "Гость2",
				Status:  entity.StatusPositive,
				Choice1: "водка",
				Choice2: "водка",
			},
			expected: ExpectedString(
				"токен", string(token),
				"ссылка", link,
				"гость", "Гость1 и Гость2",
				"статус", "приду",
				"выбор Гость1", "водка",
				"выбор Гость2", "водка",
			),
		},
		{
			name: "single mode on - negative",
			guest: entity.Guest{
				Token:      token,
				SingleMode: true,
				Name1:      "Гость1",
				Name2:      "Гость2",
				Status:     entity.StatusNegative,
				Choice1:    "водка",
			},
			expected: ExpectedString(
				"токен", string(token),
				"ссылка", link,
				"гость", "Гость1",
				"статус", "не приду",
				"плюс один", "Гость2",
				"выбор Гость1", "водка",
				"выбор Гость2", "-",
			),
		},
		{
			name: "single mode on - no responded",
			guest: entity.Guest{
				Token:      token,
				SingleMode: true,
				Name1:      "Гость1",
			},
			expected: ExpectedString(
				"токен", string(token),
				"ссылка", link,
				"гость", "Гость1",
				"статус", "не определен",
				"плюс один", "-",
				"выбор Гость1", "-",
			),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			text := usecase.GuestToMessage(tc.guest, url)
			if tc.expected != text {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, text)
			}
		})
	}
}

func ExpectedString(opts ...string) string {
	switch {
	case len(opts) == 0:
		return ""
	case len(opts)%2 != 0:
		log.Fatal("the number of transmitted values must be even")
	}

	buf := strings.Builder{}
	sep1 := ": "
	sep2 := "\n"

	n := (len(opts)-1)*(len(sep1)+len(sep2)) + len(sep1)
	for i := 0; i < len(opts); i++ {
		n += len(opts[i])
	}

	buf.Grow(n)

	for i := 0; i < len(opts)-1; i = i + 2 {
		_, _ = buf.WriteString(opts[i])
		_, _ = buf.WriteString(sep1)
		_, _ = buf.WriteString(opts[i+1])
		if i < len(opts)-2 {
			_, _ = buf.WriteString(sep2)
		}
	}

	return buf.String()
}

func BenchmarkGuest_GuestToString(b *testing.B) {
	g := entity.Guest{
		Token:      entity.GenerateToken(),
		SingleMode: true,
		Name1:      "Гость1",
		Name2:      "Гость2",
		Status:     entity.StatusNegative,
		Choice1:    "водка",
		Choice2:    "водка",
	}

	for i := 0; i < b.N; i++ {
		_ = usecase.GuestToMessage(g, "localhost")
	}
}
