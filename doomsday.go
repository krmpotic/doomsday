package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var yearMin, yearMax int
var helpFlag = flag.Bool("d", false, "print doomsdays")

func init() {
	flag.IntVar(&yearMin, "s", 1900, "min year")
	flag.IntVar(&yearMax, "e", 2100, "max year")

	rand.Seed(time.Now().Unix())
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

func doomsday(year int) time.Weekday {
	return date(year, time.March, 14).Weekday()
}

func help() {
	for i := yearMin - yearMin%100; i <= yearMax; i += 100 {
		fmt.Printf("%d - %s\n", i, doomsday(i))
	}

	prevMonth := time.December
	for t := date(1999, 1, 1); t.Before(date(2000, 1, 1)); t = t.AddDate(0, 0, 1) {
		if t.Weekday() == doomsday(1999) {
			if prevMonth != t.Month() {
				fmt.Println()
				fmt.Printf("%10s", t.Month())
				prevMonth = t.Month()
			}
			fmt.Printf("%3d", t.Day())
		}
	}
	fmt.Println()
}

func main() {
	flag.Parse()
	if *helpFlag {
		help()
		return
	}

	t := date(yearMin+rand.Intn(yearMax-yearMin),
		time.January,
		1+rand.Intn(365))

	fmt.Printf("%d\n", t.Year())
	guessWeekday(doomsday(t.Year()))
	fmt.Printf("%d.%d.%d\n", t.Day(), t.Month(), t.Year())
	guessWeekday(t.Weekday())
}

func guessWeekday(goal time.Weekday) {
	for {
		var guess int
		fmt.Printf("Sun-Sat [0-6]? ")
		fmt.Scanf("%d", &guess)
		if guess == int(goal) {
			fmt.Printf("Correct, it was a %s\n", goal)
			break
		}
	}
}
