package local_test

import (
	"testing"

	"github.com/serge64/invite/internal/repository/local"
)

func TestChatRepository_Add(t *testing.T) {
	r := local.NewChatRepository()

	testcases := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name: "valid",
			id:   "key",
		},
		{
			name:     "no valid",
			id:       "key",
			expected: local.ErrKeyNotUnique,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Add(tc.id)
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func TestChatRepository_Values(t *testing.T) {
	r := local.NewChatRepository()
	_ = r.Add("id")
	expected := "id"
	values := r.Values()
	if values[0] != expected {
		t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", expected, values[0])
	}
}

func TestChatRepository_ValuesEmpty(t *testing.T) {
	r := local.NewChatRepository()
	values := r.Values()
	if len(values) != 0 {
		t.Errorf("Values not equals:\n- expected: 0\n- actual: %d", len(values))
	}
}

func TestChatRepository_Exists(t *testing.T) {
	r := local.NewChatRepository()
	_ = r.Add("key")

	testcases := []struct {
		name     string
		id       string
		expected bool
	}{
		{
			name:     "valid",
			id:       "key",
			expected: true,
		},
		{
			name: "no valid",
			id:   "no valid key",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ok := r.Exists(tc.id)
			if ok != tc.expected {
				t.Errorf("Values not equals:\n- expected: %t\n- actual: %t", tc.expected, ok)
			}
		})
	}
}

func TestChatRepository_Delete(t *testing.T) {
	r := local.NewChatRepository()
	_ = r.Add("key")

	testcases := []struct {
		name     string
		id       string
		expected error
	}{
		{
			name: "valid",
			id:   "key",
		},
		{
			name:     "no valid",
			id:       "key",
			expected: local.ErrKeyNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := r.Delete(tc.id)
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func BenchmarkChatRepository_Add(b *testing.B) {
	r := local.NewChatRepository()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Add("key")
		}
	})
}

func BenchmarkChatRepository_Values(b *testing.B) {
	r := local.NewChatRepository()
	_ = r.Add("key")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Values()
		}
	})
}

func BenchmarkChatRepository_Exists(b *testing.B) {
	r := local.NewChatRepository()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Exists("key")
		}
	})
}

func BenchmarkChatRepository_Delete(b *testing.B) {
	r := local.NewChatRepository()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = r.Delete("key")
		}
	})
}
