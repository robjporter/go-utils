package display

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"../../go/as"
)

var (
	Font3 = map[rune][][]bool{
		'0': {{true, true, true}, {true, false, true}, {true, false, true}, {true, false, true}, {true, true, true}},
		'1': {{false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
		'2': {{true, true, true}, {false, false, true}, {true, true, true}, {true, false, false}, {true, true, true}},
		'3': {{true, true, true}, {false, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
		'4': {{true, false, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
		'5': {{true, true, true}, {true, false, false}, {true, true, true}, {false, false, true}, {true, true, true}},
		'6': {{true, true, true}, {true, false, false}, {true, true, true}, {true, false, true}, {true, true, true}},
		'7': {{true, true, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
		'8': {{true, true, true}, {true, false, true}, {true, true, true}, {true, false, true}, {true, true, true}},
		'9': {{true, true, true}, {true, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
		':': {{false}, {true}, {false}, {true}, {false}}}
	Font5 = map[rune][][]bool{
		'0': {{true, true, true, true, true}, {true, true, false, true, true}, {true, true, false, true, true}, {true, true, false, true, true}, {true, true, true, true, true}},
		'1': {{false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
		'2': {{true, true, true}, {false, false, true}, {true, true, true}, {true, false, false}, {true, true, true}},
		'3': {{true, true, true}, {false, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
		'4': {{true, false, true}, {true, false, true}, {true, true, true}, {false, false, true}, {false, false, true}},
		'5': {{true, true, true}, {true, false, false}, {true, true, true}, {false, false, true}, {true, true, true}},
		'6': {{true, true, true}, {true, false, false}, {true, true, true}, {true, false, true}, {true, true, true}},
		'7': {{true, true, true}, {false, false, true}, {false, false, true}, {false, false, true}, {false, false, true}},
		'8': {{true, true, true}, {true, false, true}, {true, true, true}, {true, false, true}, {true, true, true}},
		'9': {{true, true, true}, {true, false, true}, {true, true, true}, {false, false, true}, {true, true, true}},
		':': {{false}, {true}, {false}, {true}, {false}}}
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetTerminalSize() (int, int, error) {
	height := 0
	width := 0
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err == nil {
		splits := strings.Split(string(out), " ")
		if len(splits) == 2 {
			height = as.ToInt(splits[0])
			width = as.ToInt(strings.TrimRight(splits[1], "\n"))
		}
	}

	return height, width, err
}

// StrThickLine returns a thick line (using '=')
func ThickLine(n int) (l string) {
	l = strings.Repeat("=", n)
	return l + "\n"
}

// StrThinLine returns a thin line (using '-')
func ThinLine(n int) (l string) {
	l = strings.Repeat("-", n)
	return l + "\n"
}

// StrThinLine returns a thin line (using '-')
func SpecialLine(n int) (l string) {
	l = strings.Repeat("*", n)
	return l + "\n"
}

// StrSpaces returns a line with spaces
func SpaceLine(n int) (l string) {
	l = strings.Repeat(" ", n)
	return
}
