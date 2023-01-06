package day25

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encode(t *testing.T) {
	for _, tt := range examples {
		t.Run(strconv.Itoa(tt.value), func(t *testing.T) {
			assert.Equal(t, tt.snafu, encode(tt.value))
		})
	}
}
