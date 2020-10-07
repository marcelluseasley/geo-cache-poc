package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// LogEntry ...
type LogEntry struct {
	event time.Time
	latitude  string
	longitude string
}

func main() {

	//le := LogEntry{}

	regex := *regexp.MustCompile(`([0-2][0-9]|(3)[0-1])/(\w{3})/(\d{4}):(\d{2}):(\d{2}):(\d{2}).+geoLongitude=(-?\d+(\.\d+)?)&geoLatitude=(-?\d+(\.\d+)?)`)

	file, err := os.Open("log.txt")
	if err != nil {
		log.Fatal("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

/*
1 - Day
3 - Month
4 - Year
5 - Hour
6 - Minute
7 - Second
8 - Longitude
10 - Latitude
*/

	for _, eachLine := range text {
		res := regex.FindAllStringSubmatch(eachLine, -1)
		fmt.Println(len(res[0]))
		for i, j := range res[0] {
			fmt.Println(i, j)
		}
		break
	}

}

// CoordPrecisionOfX returns a string representation of a geographical coordinate to a specific number of decimal places (truncation, no rounding)
func CoordPrecisionOfX(f float64, x int) string {
	s64 := strconv.FormatFloat(f, 'f', -1, 64)

	decIndex := strings.Index(s64, ".")

	if decIndex == -1 {
		return s64
	}
	if decIndex+x+1 >= len(s64) {
		return s64[0:]
	}
	return s64[0 : decIndex+x+1]
}
