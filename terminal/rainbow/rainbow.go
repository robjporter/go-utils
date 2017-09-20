package rainbow

import (
	"encoding/hex"
	"fmt"
	"strings"
)

var (
	before   = []byte("\033[")
	after    = []byte("m")
	reset    = []byte("\033[0;00m")
	fgcolors = fgTermRGB[16:232]
	bgcolors = bgTermRGB[16:232]
)

func New() {

}

func ForegroundGradient(in string, fr, fg, fb, ir, ig, ib uint) string {
	tmp := ""
	red := fr
	green := fg
	blue := fb
	for i := 0; i < len(in); i++ {
		letter := in[i : i+1]
		tmp += colourise(letter, red, green, blue, 0, 0, 0)
		red += ir
		green += ig
		blue += ib
	}
	return tmp
}

func BackgroundGradient(in string, fr, fg, fb, br, bg, bb, ir, ig, ib uint) string {
	tmp := ""
	red := br
	green := bg
	blue := bb
	for i := 0; i < len(in); i++ {
		letter := in[i : i+1]
		tmp += colourise(letter, fr, fg, fb, red, green, blue)
		red += ir
		green += ig
		blue += ib
	}
	return tmp
}

func Foreground(in string, fr, fg, fb uint) string {
	return colourise(in, fr, fg, fb, 0, 0, 0)
}
func Background(in string, fr, fg, fb, br, bg, bb uint) string {
	return colourise(in, fr, fg, fb, br, bg, bb)
}

func colourise(in string, fr, fg, fb, br, bg, bb uint) string {
	colour := rgb(fr, fg, fb, br, bg, bb)
	return string(append(append(append(append(before, colour...), after...), in...), reset...))
}

func color(r, g, b uint, foreground bool) []byte {
	approxR := ((uint16(r) * 5) / 255)
	approxG := ((uint16(g) * 5) / 255)
	approxB := ((uint16(b) * 5) / 255)
	i := 36*approxR + 6*approxG + approxB
	if foreground {
		return fgcolors[i]
	}
	return bgcolors[i]
}

func rgb(fr, fg, fb, br, bg, bb uint) []byte {
	fore := append(color(fr, fg, fb, true), byte(';'))
	back := color(br, bg, bb, false)
	return append(fore, back...)
}

func Convert(hexer string) (uint, uint, uint) {
	if strings.HasPrefix(hexer, "#") {
		hexer = hexer[1:len(hexer)]
	}
	if len(hexer) == 3 {
		hexer = fmt.Sprintf("%c%c%c%c%c%c", hexer[0], hexer[0], hexer[1], hexer[1], hexer[2], hexer[2])
	}
	d, _ := hex.DecodeString(hexer)
	return uint(d[0]), uint(d[1]), uint(d[2])
}
