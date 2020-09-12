package color

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestColor_HSLA(t *testing.T) {
	t.Run("red", func(t *testing.T) {
		color, _ := ParseString("#ff0000")
		h, s, l, a := color.HSLA()
		assert.Equal(t, 0.0, h)
		assert.Equal(t, 1.0, s)
		assert.Equal(t, 0.5, l)
		assert.Equal(t, 1.0, a)
	})
	t.Run("blue", func(t *testing.T) {
		color, _ := ParseString("#0000ff")
		h, s, l, a := color.HSLA()
		assert.Equal(t, 240.0, h)
		assert.Equal(t, 1.0, s)
		assert.Equal(t, 0.5, l)
		assert.Equal(t, 1.0, a)
	})
	t.Run("light_blue", func(t *testing.T) {
		color, _ := ParseString("#b7b7d5")
		h, s, l, a := color.HSLA()
		assert.Equal(t, 240.0, h)
		assert.Equal(t, 0.26, math.Round(s*100)/100.0)
		assert.Equal(t, 0.78, math.Round(l*100)/100.0)
		assert.Equal(t, 1.0, math.Round(a*100)/100)
	})
}

func TestParseHSLA(t *testing.T) {
	tests := []string{"#ff0000", "#b7b7d5", "#c4c4c4"}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			color, _ := ParseString(test)
			h, s, l, a := color.HSLA()
			assert.True(t, ParseHSLA(h, s, l, a).Equal(color))
		})
	}
}
