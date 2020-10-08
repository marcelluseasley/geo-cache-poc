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

// Month constants

var months = map[string]time.Month{
	"Jan": 1,
	"Feb": 2, 
	"Mar": 3,
	"Apr": 4,
	"May": 5,
	"Jun": 6,
	"Jul": 7,
	"Aug": 8,
	"Sep": 9,
	"Oct": 10,
	"Nov": 11,
	"Dec": 12,
}


const Precision = 2

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
		
		// create time object
		year, err := strconv.Atoi(res[0][4])
		if err != nil {
			log.Fatal("invalid year")
			continue
		}

		month := months[res[0][3]]

		day, err := strconv.Atoi(res[0][1])
		if err != nil {
			log.Fatal("invalid day")
			continue
		}

		hours, err := strconv.Atoi(res[0][5])
		if err != nil {
			log.Fatal("invalid hour")
			continue
		}

		minutes, err := strconv.Atoi(res[0][6])
		if err != nil {
			log.Fatal("invalid minute")
			continue
		}

		seconds, err := strconv.Atoi(res[0][7])
		if err != nil {
			log.Fatal("invalid seconds")
			continue
		}
		latitude := res[0][10]
		longitude := res[0][8]
		t := time.Date(year, month, day, hours, minutes, seconds, 0, time.UTC)
		fmt.Printf("%s lat=%s long=%s\n", t.Local(), coordPrecisionOfX(latitude, Precision), coordPrecisionOfX(longitude, Precision))
		
	}

}

// CoordPrecisionOfX returns a string representation of a geographical coordinate to a specific number of decimal places (truncation, no rounding)
func coordPrecisionOfX(coord string, x int) string {
	//s64 := strconv.FormatFloat(f, 'f', -1, 64)

	decIndex := strings.Index(coord, ".")

	if decIndex == -1 {
		return coord
	}
	if decIndex+x+1 >= len(coord) {
		return coord[0:]
	}
	return coord[0 : decIndex+x+1]
}
