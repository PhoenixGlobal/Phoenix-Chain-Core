package common

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

// Report gives off a warning requesting the user to submit an issue to the github tracker.
func Report(extra ...interface{}) {
	fmt.Fprintln(os.Stderr, "You've encountered a sought after, hard to reproduce bug. Please report this to the developers <3 https://github.com/ethereum/go-ethereum/issues")
	fmt.Fprintln(os.Stderr, extra...)

	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "%v:%v\n", file, line)

	debug.PrintStack()

	fmt.Fprintln(os.Stderr, "#### BUG! PLEASE REPORT ####")
}

// PrintDepricationWarning prinst the given string in a box using fmt.Println.
func PrintDepricationWarning(str string) {
	line := strings.Repeat("#", len(str)+4)
	emptyLine := strings.Repeat(" ", len(str))
	fmt.Printf(`
%s
# %s #
# %s #
# %s #
%s

`, line, emptyLine, str, emptyLine, line)
}

/*func CurrentGoRoutineID() string {
	bytes := debug.Stack()
	for i, ch := range bytes {
		if ch == '\n' || ch == '\r' {
			bytes = bytes[0:i]
			break
		}
	}
	line := string(bytes)
	var valid = regexp.MustCompile(`goroutine\s(\d+)\s+\[`)

	if params := valid.FindAllStringSubmatch(line, -1); params != nil {
		return params[0][1]
	} else {
		return ""
	}
}*/
