package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	flagMoney       float64
	flagStartHour   int
	flagJobDuration time.Duration
)

func main() {
	flag.Float64Var(&flagMoney, "money", 130.00, "your earnings per day")
	flag.IntVar(&flagStartHour, "start-hour", 8, "your earnings per day")
	flag.DurationVar(&flagJobDuration, "duration", time.Hour*8, "hours per day you work")
	flag.Parse()

	now := time.Now()

	// TODO: Customize starting hour. Example: -start 8:30
	start := time.Date(now.Year(), now.Month(), now.Day(), flagStartHour, 0, 0, 0, time.UTC)

	// TODO: Use math/big instead of float64s
	totalSeconds := flagJobDuration.Seconds()
	seconds := now.Sub(start).Seconds()
	earnings := flagMoney * clamp01(seconds/totalSeconds)

	// TODO: Customize currency
	// TODO: Customize currency fmt ($x, 10€, ...)
	fmt.Printf("You are %.2f€ richer by now...\n", earnings)
}

func clamp01(value float64) float64 {
	if value < 0.0 {
		return 0.0
	}
	if value > 1.0 {
		return 1.0
	}

	return value
}
