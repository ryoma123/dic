package dic

import (
	"fmt"
	"os"
)

const (
	exitErr      = 1
	errPrefix    = "Error"
	noticePrefix = "Notice"
)

func setError(msg error) {
	fmt.Printf("%s: %s\n", errPrefix, msg)
	os.Exit(exitErr)
}

func setNotice(msg string) {
	fmt.Printf("%s: %s\n", noticePrefix, msg)
}
