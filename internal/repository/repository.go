package repository

import (
	"context"

	"github.com/serge64/invite/internal/entity"
)

type Repository struct {
	Chat  ChatRepository
	Guest GuestRepository
}

type ChatRepository interface {
	Add(context.Context, string) error
	Chats(context.Context) []string
	Exists(context.Context, string) bool
	Delete(context.Context, string) error
}

type GuestRepository interface {
	Create(context.Context, entity.Guest) error
	Find(context.Context, string) (entity.Guest, bool)
	Guests(context.Context) []entity.Guest
	Delete(context.Context, string) error
}
