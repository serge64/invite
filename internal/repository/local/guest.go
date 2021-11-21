package local

import (
	"context"
	"sync"

	"github.com/serge64/invite/internal/entity"
)

type GuestRepository struct {
	db  LocalStorage
	mu  *sync.Mutex
	buf []entity.Guest
}

func NewGuestRepository() GuestRepository {
	return GuestRepository{
		db:  NewLocalStorage(),
		mu:  &sync.Mutex{},
		buf: make([]entity.Guest, 0, 100),
	}
}

func (r GuestRepository) Create(_ context.Context, g entity.Guest) error {
	return r.db.Set(string(g.Token), g)
}

func (r GuestRepository) Find(_ context.Context, token string) (entity.Guest, bool) {
	var guest entity.Guest
	g, ok := r.db.Get(token)
	if g != nil {
		guest = g.(entity.Guest)
	}
	return guest, ok
}

func (r GuestRepository) Guests(_ context.Context) []entity.Guest {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buf = r.buf[:0]
	for _, v := range r.db.Values() {
		r.buf = append(r.buf, v.(entity.Guest))
	}
	return r.buf
}

func (r GuestRepository) Delete(_ context.Context, token string) error {
	return r.db.Delete(token)
}
