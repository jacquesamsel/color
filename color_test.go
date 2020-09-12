package color

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseString(t *testing.T) {
	t.Run("valid_rgb", func(t *testing.T) {
		color, err := ParseString("#ff0000")
		assert.Nil(t, err)
		assert.Equal(t, uint8(255), color.Red)
		assert.Equal(t, uint8(0), color.Green)
		assert.Equal(t, uint8(0), color.Blue)
		assert.Equal(t, uint8(255), color.Alpha)
	})
	t.Run("valid_rgba", func(t *testing.T) {
		color, err := ParseString("#3bde6291")
		assert.Nil(t, err)
		assert.Equal(t, uint8(59), color.Red)
		assert.Equal(t, uint8(222), color.Green)
		assert.Equal(t, uint8(98), color.Blue)
		assert.Equal(t, uint8(145), color.Alpha)
	})
	t.Run("invalid_rgb_chars", func(t *testing.T) {
		_, err := ParseString("#00z000")
		assert.True(t, errors.Is(err, ErrorInvalidColor))
	})
	t.Run("invalid_rgb_length", func(t *testing.T) {
		_, err := ParseString("#3bde629")
		assert.True(t, errors.Is(err, ErrorInvalidColor))
	})
}

func TestColor_String(t *testing.T) {
	t.Run("rgb", func(t *testing.T) {
		c := &Color{
			Red:   255,
			Green: 0,
			Blue:  0,
			Alpha: 255,
		}
		assert.Equal(t, "#ff0000", c.String())
	})
	t.Run("rgba", func(t *testing.T) {
		c := &Color{
			Red:   255,
			Green: 0,
			Blue:  0,
			Alpha: 100,
		}
		assert.Equal(t, "#ff000064", c.String())
	})
}
