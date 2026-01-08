package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/atotto/clipboard"
)

var CLI struct {
	From  string `arg:"" help:"Startime (HHMM)"`
	To    string `arg:"" help:"Endtime (HHMM)"`
	Pause int    `arg:"" help:"Pause in minutes" default:"0"`
}

func main() {
	ctx := kong.Parse(&CLI,
		kong.Name("Time diff tool"),
		kong.Description("A simple tool to helpt get the amount of hous from a start too and end time with an added pause"),
		kong.UsageOnError(),
	)

	// layout acorging to time package constants https://pkg.go.dev/time#pkg-constants
	layout := "1504"

	fmt.Printf("from is: %s\n", CLI.From)

	fmt.Printf("to is: %s\n", CLI.To)
	fromTime, err := time.Parse(layout, CLI.From)
	if err != nil {
		ctx.FatalIfErrorf((fmt.Errorf("input for 'from' time wrong: use (HHMM)")))
	}

	toTime, err := time.Parse(layout, CLI.To)
	if err != nil {
		ctx.FatalIfErrorf((fmt.Errorf("input for 'to' time wrong: use (HHMM)")))
	}

	diff := toTime.Sub(fromTime)
	diff = diff - (time.Duration(CLI.Pause) * time.Minute)

	decimalHours := diff.Hours()

	output := fmt.Sprintf("%.2f", decimalHours)
	outputDanish := strings.Replace(output, ".", ",", 1)

	err = clipboard.WriteAll(outputDanish)
	if err != nil {
		fmt.Printf("Couldn't copy to clipboard: %v\n", err)
	}

	fmt.Println("The time differense is")
	fmt.Printf("%s\n", outputDanish)
}
