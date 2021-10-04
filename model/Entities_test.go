package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parse
		Success bool   // parser resul
		Type    string // the input typ
		Value   string // the input valu
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX05ABC8E", false, "TX", "ABC8E", 5},
		{"NN04111A", false, "NN", "111A", 4},
		{"TX06ELKQJ", false, "TX", "ELKQJ", 5},
		{"TX0PELKQJ", false, "TX", "ELKQJ", 5},
		{"ZA03ELK", false, "ZA", "ELQ", 3},
	}

	for _, c := range cases {
		_, err := Parser(c.Input)
		assert.Equal(t, err == nil, c.Success)
	}
}
