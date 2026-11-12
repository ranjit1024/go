package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func line_filter() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr)
	}
}
