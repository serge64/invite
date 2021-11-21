package local_test

import (
	"context"
	"testing"

	"github.com/serge64/invite/internal/entity"
	"github.com/serge64/invite/internal/repository/local"
)

var (
	guest = entity.Guest{}
)

func init() {
	guest.Token = entity.GenerateToken()
	guest.Name1 = "Гость1"
}

func TestGuestRepository_Create(t *testing.T) {
	r := local.NewGuestRepository()

	testcases := []struct {
		name     string
		guest    entity.Guest
		expected error
	}{
		{
			name:  "valid",
			guest: guest,
		},
		{
			name:     "no valid",
			guest:    guest,
			expected: local.ErrKeyNotUnique,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Create(context.TODO(), tc.guest)
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func TestGuestRepository_Find(t *testing.T) {
	r := local.NewGuestRepository()
	_ = r.Create(context.TODO(), guest)

	testcases := []struct {
		name     string
		token    entity.Token
		expected entity.Guest
	}{
		{
			name:     "valid",
			token:    guest.Token,
			expected: guest,
		},
		{
			name:     "no valid",
			token:    entity.GenerateToken(),
			expected: entity.Guest{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			g, _ := r.Find(context.TODO(), string(tc.token))
			if g.Token != tc.expected.Token {
				t.Errorf("Values not equals:\n- expected: %#v\n- actual: %#v", tc.expected, g)
			}
		})
	}
}

func TestGuestRepository_Guests(t *testing.T) {
	r := local.NewGuestRepository()
	_ = r.Create(context.TODO(), guest)
	values := r.Guests(context.TODO())
	if values[0].Token != guest.Token {
		t.Errorf("Values not equals:\n- expected: %#v\n- actual: %#v", guest, values[0])
	}
}

func TestGuestRepository_GuestsEmpty(t *testing.T) {
	r := local.NewGuestRepository()
	values := r.Guests(context.TODO())
	if len(values) != 0 {
		t.Errorf("Values not equals:\n- expected: 0\n- actual: %d", len(values))
	}
}

func TestGuestRepository_Delete(t *testing.T) {
	r := local.NewGuestRepository()
	_ = r.Create(context.TODO(), guest)

	testcases := []struct {
		name     string
		token    entity.Token
		expected error
	}{
		{
			name:  "valid",
			token: guest.Token,
		},
		{
			name:     "no valid",
			token:    entity.GenerateToken(),
			expected: local.ErrKeyNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Delete(context.TODO(), string(tc.token))
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func BenchmarkGuestRepository_Create(b *testing.B) {
	r := local.NewGuestRepository()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Create(context.TODO(), guest)
		}
	})
}

func BenchmarkGuestRepository_Guests(b *testing.B) {
	r := local.NewGuestRepository()
	_ = r.Create(context.TODO(), guest)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Guests(context.TODO())
		}
	})
}

func BenchmarkGuestRepository_Find(b *testing.B) {
	r := local.NewGuestRepository()
	t := string(guest.Token)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = r.Find(context.TODO(), t)
		}
	})
}

func BenchmarkGuestRepository_Delete(b *testing.B) {
	r := local.NewGuestRepository()
	t := string(guest.Token)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Delete(context.TODO(), t)
		}
	})
}
