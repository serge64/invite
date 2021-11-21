package entity

import (
	"fmt"
	"strings"
)

type Guest struct {
	Token      Token  `json:"token,omitempty"`
	SingleMode bool   `json:"single_mode"`
	Name1      string `json:"name1"`
	Name2      string `json:"name2"`
	Status     Status `json:"status"`
	Choice1    string `json:"choice1"`
	Choice2    string `json:"choice2"`
}

func (g Guest) ToString(url string) string {
	text := make([]string, 0, 7)
	builder := strings.Builder{}
	token := string(g.Token)

	builder.Grow(len(url) + CodeSize)

	_, _ = builder.WriteString(url)
	_, _ = builder.WriteString(token)

	text = append(text, fmt.Sprintf("токен: %s", token))
	text = append(text, fmt.Sprintf("ссылка: [%s](%s)", builder.String(), builder.String()))

	builder.Reset()
	builder.Grow(30)

	_, _ = builder.WriteString(g.Name1)
	if !g.SingleMode {
		_, _ = builder.WriteString(" и ")
		_, _ = builder.WriteString(g.Name2)
	}

	text = append(text, fmt.Sprintf("гость: %s", builder.String()))
	text = append(text, fmt.Sprintf("статус: %s", g.Status))

	if g.SingleMode {
		text = append(text, fmt.Sprintf("плюс один: %s", validateField(g.Name2)))
	}

	text = append(text, fmt.Sprintf("выбор %s: %s", g.Name1, validateField(g.Choice1)))

	if g.Name2 != "" {
		text = append(text, fmt.Sprintf("выбор %s: %s", g.Name2, validateField(g.Choice2)))
	}

	return strings.Join(text, "\n")
}

func validateField(value string) string {
	if value == "" {
		return "-"
	}
	return value
}
