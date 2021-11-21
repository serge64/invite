package usecase

import (
	"strings"
	"sync"

	"github.com/serge64/invite/internal/entity"
)

var (
	once sync.Once
	buf  strings.Builder
)

func makeBuffer() {
	buf = strings.Builder{}
}

func GuestToMessage(g entity.Guest, url string) string {
	once.Do(makeBuffer)
	token := string(g.Token)

	buf.Reset()
	buf.Grow(250)

	writeToken(token)
	writeLink(url, token)
	writeGuestName(g)
	writeStatus(g.Status)

	if g.SingleMode {
		writeAdditionPerson(g.Name2)
	}

	writeChoice1(g)

	if g.Name2 != "" {
		writeChoice2(g)
	}

	return buf.String()
}

func writeToken(token string) {
	_, _ = buf.WriteString("токен: ")
	_, _ = buf.WriteString(token)
	_, _ = buf.WriteString("\n")
}

func writeLink(url, token string) {
	_, _ = buf.WriteString("ссылка: ")
	_, _ = buf.WriteString(url)
	_, _ = buf.WriteString(token)
	_, _ = buf.WriteString("\n")
}

func writeGuestName(g entity.Guest) {
	_, _ = buf.WriteString("гость: ")
	_, _ = buf.WriteString(g.Name1)
	if !g.SingleMode {
		_, _ = buf.WriteString(" и ")
		_, _ = buf.WriteString(g.Name2)
	}
	_, _ = buf.WriteString("\n")
}

func writeStatus(status entity.Status) {
	_, _ = buf.WriteString("статус: ")
	switch status {
	case entity.StatusPositive:
		_, _ = buf.WriteString("приду")
	case entity.StatusNegative:
		_, _ = buf.WriteString("не приду")
	default:
		_, _ = buf.WriteString("не определен")
	}
	_, _ = buf.WriteString("\n")
}

func writeAdditionPerson(name string) {
	_, _ = buf.WriteString("плюс один: ")
	_, _ = buf.WriteString(validateField(name))
	_, _ = buf.WriteString("\n")
}

func writeChoice1(g entity.Guest) {
	_, _ = buf.WriteString("выбор ")
	_, _ = buf.WriteString(g.Name1)
	_, _ = buf.WriteString(": ")
	_, _ = buf.WriteString(validateField(g.Choice1))
}

func writeChoice2(g entity.Guest) {
	_, _ = buf.WriteString("\nвыбор ")
	_, _ = buf.WriteString(g.Name2)
	_, _ = buf.WriteString(": ")
	_, _ = buf.WriteString(validateField(g.Choice2))
}

func validateField(value string) string {
	if value == "" {
		return "-"
	}
	return value
}
