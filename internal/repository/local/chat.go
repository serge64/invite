package local

import (
	"context"
	"sync"
)

type ChatRepository struct {
	db  LocalStorage
	mu  *sync.Mutex
	buf []string
}

func NewChatRepository() ChatRepository {
	return ChatRepository{
		db:  NewLocalStorage(),
		mu:  &sync.Mutex{},
		buf: make([]string, 0, 100),
	}
}

func (r ChatRepository) Add(_ context.Context, id string) error {
	return r.db.Set(id, id)
}

func (r ChatRepository) Chats(_ context.Context) []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buf = r.buf[:0]
	for _, v := range r.db.Values() {
		r.buf = append(r.buf, v.(string))
	}
	return r.buf
}

func (r ChatRepository) Exists(_ context.Context, id string) bool {
	_, ok := r.db.Get(id)
	return ok
}

func (r ChatRepository) Delete(_ context.Context, id string) error {
	return r.db.Delete(id)
}
