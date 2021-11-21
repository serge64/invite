package local

import "sync"

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

func (r ChatRepository) Add(id string) error {
	return r.db.Set(id, id)
}

func (r ChatRepository) Values() []string {
	list := r.db.Values()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.buf = r.buf[:0]
	for _, v := range list {
		r.buf = append(r.buf, v.(string))
	}
	return r.buf
}

func (r ChatRepository) Exists(id string) bool {
	_, ok := r.db.Get(id)
	return ok
}

func (r ChatRepository) Delete(id string) error {
	return r.db.Delete(id)
}
