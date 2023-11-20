package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	for i := 1; i <= 8; i++ {
		var maxVal uint64 = 1 << (i * 7)
		ns := time.Duration(maxVal)
		//us := time.Duration(safeMultiply(int64(ns), 1e3))
		ms := time.Duration(safeMultiply(int64(ns), 1e6))
		fmt.Printf("%d,%s,%s,%s\n", i, formatNS(ns), formatUS(maxVal), formatNS(ms))
	}
	return nil
}

func formatNS(t time.Duration) string {
	switch {
	case t == -1:
		return "         "
	case t < time.Microsecond:
		return fmt.Sprintf("% 7.1fns", float64(t))
	case t < time.Millisecond:
		return fmt.Sprintf("% 7.1fus", float64(t)/1e3)
	case t < time.Second:
		return fmt.Sprintf("% 7.1fms", float64(t)/1e6)
	case t < time.Minute:
		return fmt.Sprintf("% 7.1f s", float64(t)/1e9)
	case t < time.Hour:
		return fmt.Sprintf("% 7.1f m", float64(t)/1e9/60)
	case t < 24*time.Hour:
		return fmt.Sprintf("% 7.1f h", float64(t)/1e9/60/60)
	case t < 24*time.Hour*365:
		return fmt.Sprintf("% 7.1f d", float64(t)/1e9/60/60/24)
	default:
		return fmt.Sprintf("% 7.1f y", float64(t)/1e9/60/60/24/365)
	}
}

func formatTime(val uint64, unit string) string {
	var scale int
	switch unit {
	case "ns":
		scale = 1
	case "us":
		scale = 1e3
	case "ms":
		scale = 1e6
	}
	var (
		usec = uint64(1e3 / scale)
		msec = uint64(1e6 / scale)
		sec  = uint64(1e9 / scale)
		min  = 60 * sec
		hour = 60 * min
		day  = 24 * hour
		year = 365 * hour
	)
	switch {
	case val < usec:
		return fmt.Sprintf("% 7.1fns", float64(val))
	case val < msec:
		return fmt.Sprintf("% 7.1fus", float64(val)/float64(usec))
	case val < sec:
		return fmt.Sprintf("% 7.1fms", float64(val)/float64(msec))
	case val < min:
		return fmt.Sprintf("% 7.1f s", float64(val)/float64(sec))
	case val < hour:
		return fmt.Sprintf("% 7.1f m", float64(val)/float64(min))
	case val < day:
		return fmt.Sprintf("% 7.1f h", float64(val)/float64(hour))
	case val < year:
		return fmt.Sprintf("% 7.1f d", float64(val)/float64(day))
	default:
		return fmt.Sprintf("% 7.1f y", float64(val)/float64(year))
	}
}

func formatUS(val uint64) string {
	const (
		msec = 1e3
		sec  = 1e6
		min  = 60 * sec
		hour = 60 * min
		day  = 24 * hour
		year = 365 * hour
	)
	switch {
	case val < msec:
		return fmt.Sprintf("% 7.1fus", float64(val))
	case val < sec:
		return fmt.Sprintf("% 7.1fms", float64(val)/msec)
	case val < min:
		return fmt.Sprintf("% 7.1f s", float64(val)/sec)
	case val < hour:
		return fmt.Sprintf("% 7.1f m", float64(val)/min)
	case val < day:
		return fmt.Sprintf("% 7.1f h", float64(val)/hour)
	case val < 365*24*60*60*1e6:
		return fmt.Sprintf("% 7.1f d", float64(val)/1e6/60/60/24)
	default:
		return fmt.Sprintf("% 7.1f y", float64(val)/1e6/60/60/24/365)
	}
}

// multiplies a*b and returns the result unless it overflows in which case -1
// is returned.
// see https://stackoverflow.com/a/50744801/62383
func safeMultiply(a, b int64) int64 {
	const mostPositive = 1<<63 - 1
	const mostNegative = -(mostPositive + 1)

	result := a * b
	if a == 0 || b == 0 || a == 1 || b == 1 {
		return result
	}
	if a == mostNegative || b == mostNegative {
		return -1
	}
	if result/b != a {
		return -1
	}
	return result
}
