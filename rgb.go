package color

func (color *Color) RGBA() (r, g, b, a uint8) {
	return color.Red, color.Green, color.Blue, color.Alpha
}
