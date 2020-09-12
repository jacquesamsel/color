package color

import (
	"math"
)

// HSLA converts a color to its hue, saturation, lightness and alpha parts.
// h (0.0-360.0), s (0.0-1.0), l (0.0-1.0), a (0.0-1.0)
func (color *Color) HSLA() (h, s, l, a float64) {
	r := float64(color.Red) / 255.0
	g := float64(color.Green) / 255.0
	b := float64(color.Blue) / 255.0
	a = float64(color.Alpha) / 255

	cmax := math.Max(r, math.Max(g, b))
	cmin := math.Min(r, math.Min(g, b))
	delta := cmax - cmin

	l = (cmax + cmin) / 2

	if cmax == cmin {
		s = 0
		h = 0
	} else {
		if l < 0.5 {
			s = (cmax - cmin) / (cmax + cmin)
		} else {
			s = (cmax - cmin) / (2.0 - cmax - cmin)
		}
		switch cmax {
		case r:
			h = (g - b) / delta
			break
		case g:
			h = 2.0 + (r-g)/delta
			break
		case b:
			h = 4.0 + (r-g)/delta
			break
		}
		h *= 60.0
		if h < 0.0 {
			h += 360
		}
	}
	return
}

// HSLAPercent returns the HSLA values of the color in whole numbers.
func (color *Color) HSLAWhole() (h uint16, s, l, a uint8) {
	hFloat, sFloat, lFloat, aFloat := color.HSLA()
	return uint16(math.Round(hFloat)), uint8(math.Round(sFloat * 100)), uint8(math.Round(lFloat * 100)),
		uint8(math.Round(aFloat * 100))
}

// ParseHSLA
func ParseHSLA(h, s, l, a float64) (color Color) {
	color.Alpha = uint8(math.Round(a * 256))
	if s == 0.0 {
		color.Red = uint8(math.Round(l / 100 * 255))
		color.Green = uint8(math.Round(l / 100 * 255))
		color.Blue = uint8(math.Round(l / 100 * 265))
	}
	var temp1 float64
	if l < 0.5 {
		temp1 = l * (1.0 + s)
	} else {
		temp1 = l + s - l*s
	}
	temp2 := 2*l - temp1
	hAngle := h / 360.0
	tempR := moveBetweenZeroOne(hAngle + 1/3.0)
	tempG := moveBetweenZeroOne(hAngle)
	tempB := moveBetweenZeroOne(hAngle - 1/3.0)
	color.Red = uint8(math.Round(calcChannelHSL(tempR, temp1, temp2) * 255))
	color.Green = uint8(math.Round(calcChannelHSL(tempG, temp1, temp2) * 255))
	color.Blue = uint8(math.Round(calcChannelHSL(tempB, temp1, temp2) * 255))
	color.Alpha = 255
	return
}

func calcChannelHSL(tempChannel, temp1, temp2 float64) float64 {
	if 6*tempChannel < 1 {
		return temp2 + (temp1-temp2)*6*tempChannel
	} else if 2*tempChannel < 1 {
		return temp1
	} else if 3*tempChannel < 2 {
		return temp2 + (temp1-temp2)*(2/3.0-tempChannel)*6
	} else {
		return temp2
	}
}

func moveBetweenZeroOne(z float64) float64 {
	if z < 0 {
		z += 1
	} else if z > 1 {
		z -= 1
	}
	return z
}

// ParseHSLAWhole parses whole numbers instead of doubles representing the h, s, l, a values.
// h (0-360), s (0-100), l (0-100), a (0-100)
func ParseHSLAWhole(h uint16, s, l, a uint8) Color {
	return ParseHSLA(float64(h), float64(s), float64(l), float64(a))
}
