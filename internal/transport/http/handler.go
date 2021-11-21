package httphandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serge64/invite/internal/usecase"
)

type Handler struct {
	router *mux.Router
	uc     usecase.UseCase
}

func NewHandler(uc usecase.UseCase) Handler {
	h := Handler{
		router: mux.NewRouter(),
		uc:     uc,
	}
	h.configureRouting()
	return h
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h Handler) configureRouting() {
	// TODO: added middleware and endpoints
}
