package colours

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BOLD           = "\033[1m"
	ITALIC         = "\033[3m"
	UNDERLINE      = "\033[4m"
	STRIKETHROUGH  = "\033[9m"
	REVERSED       = "\u001b[7m"
	RESET          = "\033[0m"
	BLINK          = "\x1b[5m"
	TICK           = "✔"
	CROSS          = "✖"
	COPYRIGHT      = "©"
	REGISTREDTM    = "®"
	TRADEMARK      = "™"
	BULLET         = "•"
	ARROWLEFT      = "←"
	ARROWRIGHT     = "→"
	ARROWUP        = "↑"
	ARROWDOWN      = "↓"
	ARROWLEFTRIGHT = "↔"
	INFINITY       = "∞"
	CELSIUS        = "℃"
	FAHRENHEIT     = "℉"
	SUNSHINE       = "☀"
	CLOUDY         = "☁"
	RAIN           = "☂"
	SNOW           = "☃"
	STARBLACK      = "★"
	STARWHITE      = "☆"
	PHONEBLACK     = "☎"
	PHONEWHITE     = "☏"
	POINTLEFT      = "☚"
	POINTRIGHT     = "☛"
	POINTUP        = "☝"
	POINTDOWN      = "☟"
	DEATH          = "☠"
	SMILEY         = "☺"
	HEART          = "♡"
	DIAMOND        = "♢"
	SPADE          = "♤"
	CLUB           = "♧"
	BLOCK          = "█"
	BLOCK2         = " "
)

const (
	BLACK         = "0"
	BRIGHTBLACK   = "0;1"
	RED           = "1"
	BRIGHTRED     = "1;1"
	GREEN         = "2"
	BRIGHTGREEN   = "2;1"
	YELLOW        = "3"
	BRIGHTYELLOW  = "3;1"
	BLUE          = "4"
	BRIGHTBLUE    = "4;1"
	MAGENTA       = "5"
	BRIGHTMAGENTA = "5;1"
	CYAN          = "6"
	BRIGHTCYAN    = "6;1"
	WHITE         = "7"
	BRIGHTWHITE   = "7;1"
)

var (
	Output *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func getColor(code string) string {
	//return fmt.Sprintf("\033[3%sm", code)
	return fmt.Sprintf("\u001b[3%sm", code)
}

func getBgColor(code string) string {
	return fmt.Sprintf("\u001b[4%sm", code)
}

func Bold(str string) string {
	return fmt.Sprintf("%s%s%s", BOLD, str, RESET)
}

func Underline(str string) string {
	return fmt.Sprintf("%s%s%s", UNDERLINE, str, RESET)
}

func Italic(str string) string {
	return fmt.Sprintf("%s%s%s", ITALIC, str, RESET)
}

func Background(str string, color string) string {
	return fmt.Sprintf("%s%s%s", getBgColor(color), str, RESET)
}

func Block(count int, color string) string {
	str := strings.Repeat(BLOCK2, count)
	return fmt.Sprintf("%s%s%s", getBgColor(color), str, RESET)
}

func Color(str string, color string) string {
	return fmt.Sprintf("%s%s%s", getColor(color), str, RESET)
}

func Black(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BLACK), str, RESET)
}

func BrightBlack(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTBLACK), str, RESET)
}

func Red(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(RED), str, RESET)
}

func BrightRed(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTRED), str, RESET)
}

func Green(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(GREEN), str, RESET)
}

func BrightGreen(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTGREEN), str, RESET)
}

func Yellow(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(YELLOW), str, RESET)
}

func BrightYellow(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTYELLOW), str, RESET)
}

func Blue(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BLUE), str, RESET)
}

func BrightBlue(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTBLUE), str, RESET)
}

func Magenta(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(MAGENTA), str, RESET)
}

func BrightMagenta(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTMAGENTA), str, RESET)
}

func Cyan(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(CYAN), str, RESET)
}

func BrightCyan(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTCYAN), str, RESET)
}

func White(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(WHITE), str, RESET)
}

func BrightWhite(str string) string {
	return fmt.Sprintf("%s%s%s", getColor(BRIGHTWHITE), str, RESET)
}

func printBox(backgroundCode string, fontCode string, str string) string {
	return fmt.Sprintf("%s%s\n\n %s \n%s", backgroundCode, fontCode, str, RESET)
}

func printSmallBox(backgroundCode string, fontCode string, str string) string {
	return fmt.Sprintf("%s%s\n%s%s", backgroundCode, fontCode, str, RESET)
}

func BlackPanel(str string) string {
	return printBox(getBgColor(BRIGHTBLACK), getColor(BRIGHTWHITE), str)
}

func RedPanel(str string) string {
	return printBox(getBgColor(RED), getColor(BRIGHTWHITE), str)
}

func GreenPanel(str string) string {
	return printBox(getBgColor(GREEN), getColor(BRIGHTWHITE), str)
}

func YellowPanel(str string) string {
	return printBox(getBgColor(YELLOW), getColor(BRIGHTWHITE), str)
}

func BluePanel(str string) string {
	return printBox(getBgColor(BRIGHTBLUE), getColor(BRIGHTWHITE), str)
}

func MagentaPanel(str string) string {
	return printBox(getBgColor(MAGENTA), getColor(BRIGHTWHITE), str)
}

func CyanPanel(str string) string {
	return printBox(getBgColor(CYAN), getColor(BRIGHTWHITE), str)
}

func WhitePanel(str string) string {
	return printBox(getBgColor(WHITE), getColor(BRIGHTWHITE), str)
}

func BlackSmallPanel(str string) string {
	return printSmallBox(getBgColor(BRIGHTBLACK), getColor(BRIGHTWHITE), str)
}

func RedSmallPanel(str string) string {
	return printSmallBox(getBgColor(RED), getColor(BRIGHTWHITE), str)
}

func GreenSmallPanel(str string) string {
	return printSmallBox(getBgColor(GREEN), getColor(BRIGHTWHITE), str)
}

func YellowSmallPanel(str string) string {
	return printSmallBox(getBgColor(YELLOW), getColor(BRIGHTWHITE), str)
}

func BlueSmallPanel(str string) string {
	return printSmallBox(getBgColor(BRIGHTBLUE), getColor(BRIGHTWHITE), str)
}

func MagentaSmallPanel(str string) string {
	return printSmallBox(getBgColor(MAGENTA), getColor(BRIGHTWHITE), str)
}

func CyanSmallPanel(str string) string {
	return printSmallBox(getBgColor(CYAN), getColor(BRIGHTWHITE), str)
}

func WhiteSmallPanel(str string) string {
	return printSmallBox(getBgColor(WHITE), getColor(BRIGHTWHITE), str)
}

func Title(str string) string {
	str = strings.TrimSpace(str)
	strr := getColor(BRIGHTWHITE) + str + RESET + "\n"
	strr += getColor(WHITE) + strings.Repeat("=", len(str)) + RESET
	return strr
}

func CustomTitle(str string, titleColor string, underlineColor string) string {
	str = strings.TrimSpace(str)
	strr := getColor(titleColor) + str + RESET + "\n"
	strr += getColor(underlineColor) + strings.Repeat("=", len(str)) + RESET
	return strr
}

func Info(str string) string {
	return BrightBlue(str)
}

func Success(str string) string {
	return Green(str)
}

func Warning(str string) string {
	return BrightRed(str)
}

func Error(str string) string {
	return Red(str)
}

func PrintQColor(foreground int, background int, str string) {

}

func Highlight(str, substr string, color string) string {
	hiSubstr := Color(substr, color)
	return strings.Replace(str, substr, hiSubstr, -1)
}

func MoveTo(str string, x int, y int) (out string) {
	//x, y = GetXY(x, y)

	return fmt.Sprintf("\033[%d;%dH%s", y, x, str)
}

func Reversed(str string) string {
	return fmt.Sprintf("%s%s%s", REVERSED, str, RESET)
}

// TO LOOK AT

func Blink(str string) string {
	return fmt.Sprintf("%s%s%s", BLINK, str, RESET)
}

func StrikeThrough(str string) string {
	return fmt.Sprintf("%s%s%s", STRIKETHROUGH, str, RESET)
}

func BannerPrintLineS(s string, number int) string {
	str := ""
	for i := 0; i < number; i++ {
		str += s
	}
	return str
}

func BannerPrintLineCommentS(s string, comment string, number int) string {
	str := strings.ToUpper(comment)
	for i := 0; i < number-len(comment); i++ {
		str += s
	}
	return str
}
