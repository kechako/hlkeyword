package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

var colorTable = map[string]color.Attribute{
	"black":   color.FgBlack,
	"red":     color.FgRed,
	"green":   color.FgGreen,
	"yellow":  color.FgYellow,
	"blue":    color.FgBlue,
	"magenta": color.FgMagenta,
	"cyan":    color.FgCyan,
	"white":   color.FgWhite,
}
var bgColorTable = map[string]color.Attribute{
	"black":   color.BgBlack,
	"red":     color.BgRed,
	"green":   color.BgGreen,
	"yellow":  color.BgYellow,
	"blue":    color.BgBlue,
	"magenta": color.BgMagenta,
	"cyan":    color.BgCyan,
	"white":   color.BgWhite,
}

// main
func _main() (int, error) {
	var err error

	colorName := flag.String("color", "red", "Highlight color. (black, red, green, yellow, blue, magenta, cyan, white)")
	bgColorName := flag.String("bgcolor", "", "Highlight background color. (black, red, green, yellow, blue, magenta, cyan, white)")
	flag.Parse()

	colorAttr, ok := colorTable[*colorName]
	if !ok {
		return 1, fmt.Errorf("Invalid color is specified.")
	}
	c := color.New(colorAttr)

	args := flag.Args()
	if len(args) < 1 {
		return 1, fmt.Errorf("Keyword is not specified.")
	}
	keyword := args[0]

	var fp *os.File
	if len(args) > 1 {
		fp, err = os.Open(args[1])
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
			c.Print(keyword)
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
