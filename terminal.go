package st

import "fmt"

// RGB colors {R, G, B}
var (
	Red    = []int{210, 82, 82}
	Yellow = []int{255, 198, 109}
	Green  = []int{97, 150, 71}
	Blue   = []int{104, 151, 187}
	White  = []int{255, 255, 255}
	Black  = []int{0, 0, 0}
)

// Colorful Println. Works only with unix terminal, but not Windows CLI
//	Example:
//	 	ColorPrint("Hello, world!", true, false, Red...)
func ColorPrint(val interface{}, foreground bool, background bool, color ...int) {
	// No color
	if !foreground && !background {
		fmt.Print(val)
		return
	}

	// Text and background color
	if foreground && background {
		if len(color) < 6 {
			fmt.Print(val)
			return
		}
		fmt.Print(fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], color[3], color[4], color[5], val))
		return
	}

	// Only text color
	if foreground {
		if len(color) < 3 {
			fmt.Print(val)
			return
		}
		fmt.Print(fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}

	// Only background color
	if background {
		if len(color) < 3 {
			fmt.Print(val)
			return
		}
		fmt.Print(fmt.Sprintf("\033[48;2;%d;%d;%dm%v\033[0m", color[0], color[1], color[2], val))
		return
	}
}

// Colorful Println. Works only with unix terminal, but not Windows CLI
func ColorPrintln(val interface{}, foreground bool, background bool, color ...int) {
	ColorPrint(val, foreground, background, color...)
	fmt.Println()
}