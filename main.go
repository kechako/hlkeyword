package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// main
func _main() (int, error) {
	var err error

	if len(os.Args) < 2 {
		return 1, fmt.Errorf("Keyword is not specified.")
	}
	keyword := os.Args[1]

	var fp *os.File
	if len(os.Args) > 2 {
		fp, err = os.Open(os.Args[2])
		if err != nil {
			return 2, errors.Wrap(err, "Can not open file.")
		}
		defer fp.Close()
	} else {
		fp = os.Stdin
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := scanner.Text()
		for index := strings.Index(text, keyword); index >= 0; index = strings.Index(text, keyword) {
			fmt.Print(text[0:index])
			color.Red(keyword)
			text = text[index+len(keyword) : len(text)]
		}
		if len(text) > 0 {
			fmt.Print(text)
		}
		fmt.Println("")
	}
	err = scanner.Err()
	if err != nil {
		return 3, errors.Wrap(err, "Can not read input.")
	}

	return 0, nil
}

// Entry point
func main() {
	status, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Error] %v\n", err)
		os.Exit(status)
	}
}
