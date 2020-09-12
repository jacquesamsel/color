package color

import (
	"math"
)

// CMYKA returns the cyan, magenta, yellow, black, and alpha values of the color
func (color *Color) CMYKA() (c, m, y, k, a float64) {
	rFloat := float64(color.Red) / 255
	gFloat := float64(color.Green) / 255
	bFloat := float64(color.Blue) / 255

	k = 1 - math.Max(rFloat, math.Max(gFloat, bFloat))
	c = (1 - rFloat - k) / (1 - k)
	m = (1 - gFloat - k) / (1 - k)
	y = (1 - bFloat - k) / (1 - k)
	a = float64(color.Alpha) / 255
	return
}

func (color *Color) CMYKAWhole() (c, m, y, k, a uint8) {
	cFloat, mFloat, yFloat, kFloat, aFloat := color.CMYKA()
	return toPercent(cFloat), toPercent(mFloat), toPercent(yFloat), toPercent(kFloat), toPercent(aFloat)
}

func toPercent(f float64) uint8 {
	return uint8(math.Round(f * 100))
}

func toFloat(percent uint8) float64 {
	return float64(percent) / 100
}

func ParseCMYKA(c, m, y, k, a float64) (color Color) {
	color.Red = uint8(math.Round(255 * (1 - c) * (1 - k)))
	color.Green = uint8(math.Round(255 * (1 - m) * (1 - k)))
	color.Blue = uint8(math.Round(255 * (1 - y) * (1 - k)))
	color.Alpha = uint8(math.Round(255 * a))
	return
}

func ParseCMYKAWhole(c, m, y, k, a uint8) Color {
	return ParseCMYKA(toFloat(c), toFloat(m), toFloat(y), toFloat(k), toFloat(a))
}
