package entity

import (
	"crypto/rand"
	"sort"
	"sync"
)

type Token string

const (
	CodeSize   int    = 16
	dictionary string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	amount     byte   = byte(len(dictionary))
)

var (
	bytes []byte
	once  sync.Once
)

func makeBytes() {
	bytes = make([]byte, CodeSize)
}

func GenerateToken() Token {
	once.Do(makeBytes)
	_, _ = rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%amount]
	}
	return Token(bytes)
}

func ValidateToken(t Token) bool {
	if len(t) != CodeSize {
		return false
	}
	return match(t)
}

func match(value Token) bool {
	for _, v := range value {
		if !bsearch(int(v)) {
			return false
		}
	}
	return true
}

func bsearch(a int) bool {
	i := sort.Search(
		len(dictionary),
		func(i int) bool {
			return a <= int(dictionary[i])
		},
	)
	return i <= len(dictionary) && a == int(dictionary[i])
}
