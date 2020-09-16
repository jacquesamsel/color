package color

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrorInvalidColor = errors.New("invalid color")

	hexBase = 16
)

// Color represents a color with 8 bits allocated to each of red, green, blue, alpha
type Color struct {
	Red, Green, Blue, Alpha uint8
}

// ParseString parses a hex color.
// Supports the formats  #000000, #00000000 (with alpha)
func ParseString(s string) (Color, error) {
	s = strings.ToLower(s[1:]) // remove the #, convert to lower case
	if !(len(s) == 6 || len(s) == 8) {
		return Color{}, ErrorInvalidColor
	}
	var red, green, blue, alpha int64
	alpha = 255 // opaque by default
	var err error
	if len(s) == 6 { // #000000
		red, err = strconv.ParseInt(s[0:2], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
		green, err = strconv.ParseInt(s[2:4], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
		blue, err = strconv.ParseInt(s[4:6], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
	}
	if len(s) == 8 { // #00000000
		red, err = strconv.ParseInt(s[0:2], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
		green, err = strconv.ParseInt(s[2:4], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
		blue, err = strconv.ParseInt(s[4:6], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
		alpha, err = strconv.ParseInt(s[6:8], hexBase, 9)
		if err != nil {
			return Color{}, ErrorInvalidColor
		}
	}
	return Color{
		Red:   uint8(red),
		Green: uint8(green),
		Blue:  uint8(blue),
		Alpha: uint8(alpha),
	}, nil
}

// ParseInt64 converts an int64 into a color. For example, 0xffffff
func ParseInt64(i int64) (Color, error) {
	return ParseString(strconv.FormatInt(i, hexBase))
}

func (color Color) String() string {
	build := "#"
	build += padString(strconv.FormatInt(int64(color.Red), 16), '0', 2)
	build += padString(strconv.FormatInt(int64(color.Green), 16), '0', 2)
	build += padString(strconv.FormatInt(int64(color.Blue), 16), '0', 2)
	if color.Alpha != 255 {
		build += padString(strconv.FormatInt(int64(color.Alpha), 16), '0', 2)
	}
	return build
}

func padString(s string, c rune, length int) string {
	for len(s) < length {
		s = string(c) + s
	}
	return s
}

// Equal checks the equality of the red, green, blue and alpha channels.
func (color Color) Equal(c Color) bool {
	return color.Red == c.Red && color.Green == c.Green && color.Blue == c.Blue && color.Alpha == c.Alpha
}

func (color *Color) MarshalJSON() ([]byte, error) {
	return []byte("\"" + color.String() + "\""), nil
}

func (color *Color) UnmarshalJSON(b []byte) (err error) {
	*color, err = ParseString(string(b[1 : len(b)-1]))
	return err
}

func (color *Color) MarshalText() ([]byte, error) {
	return []byte(color.String()), nil
}

func (color *Color) UnmarshalText(d []byte) (err error) {
	*color, err = ParseString(string(d))
	return
}
