package utils

import (
	"fmt"
	"os"
)

/*
	Private functions
 */

func check_no_message(e error, allowPanic bool) {
	if e != nil {
		if allowPanic {
			panic(e)
		}
	}
}

func check_print_messages(e error, allowPanic bool, messages ...interface{}) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "", messages)
		if allowPanic {
			panic(e)
		}
	}
}

/*
	Public function
 */

func ContainsString(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func Check(e error) {
	check_no_message(e, false)
}

func CheckAndShow(e error, messages ...interface{}) {
	check_print_messages(e, true, messages)
}
