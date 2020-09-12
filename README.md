# color
Color formats implemented in golang (GO)

[![GoDoc](https://godoc.org/github.com/ocuppi/color?status.svg)](https://godoc.org/github.com/ocuppi/color) 
[![Go report](http://goreportcard.com/badge/ocuppi/color)](http://goreportcard.com/report/ocuppi/color)
[![Coverage](http://gocover.io/_badge/github.com/ocuppi/color)](https://gocover.io/github.com/ocuppi/color)
## Supported formats
### RGBA
Red, green, blue, alpha
```go
// Get RGBA values
r, g, b, a := color.RGBA()
// Parse RGBA
color2 := color.Color{255, 255, 255, 255}
```
### HSLA
Hue, saturation, lightness, alpha
> :warning: the colors are internally represented as RGBA, so there may be a slight loss in precision if they are converted.
```go
// Get HSLA values
h, s, l, a := color.HSLA()
// Parse HSLA float values (0.0-1.0)
color2 := color.ParseHSLA(1.0, 1.0, 1.0, 1.0)
// Parse HSLA whole numbers values (0-360), (0-100) * 3
color3 := color.ParseHSLAWhole(360, 100, 100, 100)
```
### CMYKA
Cyan, magenta, yellow, black, alpha. Alpha is not often used in situations involving CMYK, but it is available
> :warning: the colors are internally represented as RGBA, so there may be a slight loss in precision if they are converted.
```go
// Get CMYKA values
c, m, y, k, a := color.CMYKA()
// Parse CMYKA float values (0-1.0)
color2 := color.ParseCMYKA(1.0, 1.0, 1.0, 0.0)
// Parse CMYKA whole values (0-100)
color3 := color.ParseCMYKAWhole(100, 100, 100, 0)
```
