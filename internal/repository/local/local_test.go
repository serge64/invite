package local_test

import (
	"testing"

	"github.com/serge64/invite/internal/repository/local"
)

func TestLocalStorage_Set(t *testing.T) {
	db := local.NewLocalStorage()

	testcases := []struct {
		name     string
		key      string
		value    interface{}
		expected error
	}{
		{
			name:  "valid",
			key:   "key",
			value: "value",
		},
		{
			name:     "no valid",
			key:      "key",
			value:    "value",
			expected: local.ErrKeyNotUnique,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := db.Set(tc.key, tc.value)
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func TestLocalStorage_Get(t *testing.T) {
	db := local.NewLocalStorage()
	_ = db.Set("key", "value")

	testcases := []struct {
		name     string
		key      string
		expected string
	}{
		{
			name:     "valid",
			key:      "key",
			expected: "value",
		},
		{
			name:     "no valid",
			key:      "invalidkey",
			expected: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			value, _ := db.Get(tc.key)
			if value == nil {
				value = ""
			} else {
				value = value.(string)
			}
			if value != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, value)
			}
		})
	}
}

func TestLocalStorage_Delete(t *testing.T) {
	db := local.NewLocalStorage()
	_ = db.Set("key", "value")

	testcases := []struct {
		name     string
		key      string
		expected error
	}{
		{
			name: "valid",
			key:  "key",
		},
		{
			name:     "no valid",
			key:      "invalidkey",
			expected: local.ErrKeyNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := db.Delete(tc.key)
			if err != tc.expected {
				t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", tc.expected, err)
			}
		})
	}
}

func TestLocalStorage_Keys(t *testing.T) {
	db := local.NewLocalStorage()
	_ = db.Set("key", "value")
	expected := "key"
	keys := db.Keys()
	if expected != keys[0] {
		t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", expected, keys[0])
	}
}

func TestLocalStorage_KeysEmpty(t *testing.T) {
	db := local.NewLocalStorage()
	keys := db.Keys()
	if len(keys) != 0 {
		t.Errorf("Values not equals:\n- expected: 0\n- actual: %d", len(keys))
	}
}

func TestLocalStorage_Values(t *testing.T) {
	db := local.NewLocalStorage()
	_ = db.Set("key", "value")
	expected := "value"
	values := db.Values()
	if expected != values[0].(string) {
		t.Errorf("Values not equals:\n- expected: %s\n- actual: %s", expected, values[0].(string))
	}
}

func TestLocalStorage_ValuesEmpty(t *testing.T) {
	db := local.NewLocalStorage()
	values := db.Values()
	if len(values) != 0 {
		t.Errorf("Values not equals:\n- expected: 0\n- actual: %d", len(values))
	}
}

func BenchmarkLocalStorage_SetGet(b *testing.B) {
	db := local.NewLocalStorage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = db.Set("key", struct{}{})
		}
	})
}

func BenchmarkLocalStorage_Gett(b *testing.B) {
	db := local.NewLocalStorage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = db.Get("key")
		}
	})
}

func BenchmarkLocalStorage_Delete(b *testing.B) {
	db := local.NewLocalStorage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = db.Set("key", struct{}{})
			_ = db.Delete("key")
		}
	})
}

func BenchmarkLocalStorage_Keys(b *testing.B) {
	db := local.NewLocalStorage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = db.Set("key", struct{}{})
			_ = db.Set("key2", struct{}{})
			_ = db.Keys()
		}
	})
}

func BenchmarkLocalStorage_Values(b *testing.B) {
	db := local.NewLocalStorage()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = db.Set("key", struct{}{})
			_ = db.Set("key2", struct{}{})
			_ = db.Values()
		}
	})
}
