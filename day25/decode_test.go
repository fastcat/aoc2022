package day25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decode(t *testing.T) {
	for _, tt := range examples {
		t.Run(tt.snafu, func(t *testing.T) {
			assert.Equal(t, tt.value, decode(tt.snafu))
		})
	}
}
