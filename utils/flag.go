package utils

import (
	"bufio"
	"strings"
)

func GetFlagValue(flag string, input string) string {
	flagWithSeparator := flag + " "
	index := strings.Index(input, flagWithSeparator)
	if index == -1 {
		return ""
	}

	value := input[index+len(flagWithSeparator):]
	scanner := bufio.NewScanner(strings.NewReader(value))
	scanner.Split(bufio.ScanWords)
	if scanner.Scan() {
		value = scanner.Text()
	}

	return value
}
