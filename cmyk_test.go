package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColor_CMYKAWhole(t *testing.T) {
	t.Run("red", func(t *testing.T) {
		color, _ := ParseString("#ff0000")
		c, m, y, k, a := color.CMYKAWhole()
		fmt.Println(c, m, y, k)
		assert.Equal(t, uint8(0), c)
		assert.Equal(t, uint8(100), m)
		assert.Equal(t, uint8(100), y)
		assert.Equal(t, uint8(0), k)
		assert.Equal(t, uint8(100), a)
	})
	t.Run("black", func(t *testing.T) {
		color, _ := ParseString("#000000")
		c, m, y, k, a := color.CMYKAWhole()
		assert.Equal(t, uint8(0), c)
		assert.Equal(t, uint8(0), m)
		assert.Equal(t, uint8(0), y)
		assert.Equal(t, uint8(100), k)
		assert.Equal(t, uint8(100), a)
	})
}

func TestParseCMYKAWhole(t *testing.T) {
	t.Run("red", func(t *testing.T) {
		color := ParseCMYKAWhole(0, 100, 100, 0, 100)
		assert.Equal(t, uint8(255), color.Red)
		assert.Equal(t, uint8(0), color.Green)
		assert.Equal(t, uint8(0), color.Blue)
		assert.Equal(t, uint8(255), color.Alpha)
	})
	t.Run("black", func(t *testing.T) {
		color := ParseCMYKAWhole(0, 0, 0, 100, 100)
		assert.Equal(t, uint8(0), color.Red)
		assert.Equal(t, uint8(0), color.Green)
		assert.Equal(t, uint8(0), color.Blue)
		assert.Equal(t, uint8(255), color.Alpha)
	})
}
