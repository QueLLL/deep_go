package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
const (
	maskByte0 = 0xff << (8 * iota)
	maskByte1
	maskByte2
	maskByte3
)

func ToLittleEndian(number uint32) uint32 {
	return ((number >> 24) & maskByte0) |
		((number >> 8) & maskByte1) |
		((number << 8) & maskByte2) |
		((number << 24) & maskByte3)
}

func TestConversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
