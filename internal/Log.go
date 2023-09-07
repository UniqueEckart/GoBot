package internal

import "fmt"

func Log(log string, loglevel int) {
	if loglevel == 0 {
		fmt.Printf("[INFO] %v", log)
	}
	if loglevel == 1 {
		fmt.Printf("[ERROR] ")
	}
}

func LogFormat(log string, loglevel int, format string) {
	if loglevel == 0 && format != "" {
		fmt.Printf(log, format)
	}
}
