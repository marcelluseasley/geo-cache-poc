package main

import (
	//"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/icza/backscanner"
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

// LogEntry ...
type LogEntry struct {
	event     time.Time
	latitude  string
	longitude string
}

func (le *LogEntry) coordString() string {
	return fmt.Sprintf("%s_%s", le.latitude, le.longitude)
}

func main() {

	filePtr := flag.String("f", "log.txt", "name of log file")
	precPtr := flag.Int("p", 2, "precision of latitude/longitude")
	cachePtr := flag.Int("ttl", 3, "cache TTL")
	flag.Parse()
	precision := *precPtr
	cacheTTL := *cachePtr
	var newCount int
	var dupCount int
	cache := make(map[string]time.Time)

	file, err := os.Open(*filePtr)
	if err != nil {
		log.Fatalf("failed to open: %s", *filePtr)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatalf("cannot get file info: %v", err)
	}
	defer file.Close()

	// use backscanner package since logs are stored in reverse chronological order in log file.
	// backscanner reads lines from end of file to start of file, thus placing the lines in chronological order
	scanner := backscanner.New(file, int(fi.Size()))
	var text []string

	for {
		line, _, err := scanner.Line()
		if err != nil {
			break
		}
		text = append(text, line)
	}

	/*
	   regexp match groups
	   1 - Day
	   3 - Month
	   4 - Year
	   5 - Hour
	   6 - Minute
	   7 - Second
	   8 - Longitude
	   10 - Latitude
	*/
	regex := *regexp.MustCompile(`([0-2][0-9]|(3)[0-1])/(\w{3})/(\d{4}):(\d{2}):(\d{2}):(\d{2}).+geoLongitude=(-?\d+(\.\d+)?)&geoLatitude=(-?\d+(\.\d+)?)`)
	for _, eachLine := range text[1:] {

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
		le := LogEntry{event: t, latitude: coordPrecisionOfX(latitude, precision), longitude: coordPrecisionOfX(longitude, precision)}

		// caching logic
		key := le.coordString()
		//fmt.Println(key, le.event.Local())
		if tStamp, ok := cache[key]; !ok {
			cache[key] = le.event
			newCount++
		} else {
			if cacheTTL == -1 {
				dupCount++
					cache[key] = le.event
			} else {
				if le.event.Before(tStamp.Add(time.Minute * time.Duration(cacheTTL))) {
					dupCount++
					cache[key] = le.event // update cache with new timestamp (restart TTL)?
				} else {
					cache[key] = le.event
					newCount++
				}
			}

		}

	}

	fmt.Printf("Cache Hits / New MPX requests (lat/long precision of %d) and TTL: %d minutes\n", precision, cacheTTL)
	fmt.Printf("New: %d\tDuplicate: %d\n", newCount, dupCount)

}

// coordPrecisionOfX returns a string representation of a geographical coordinate to a specific number of decimal places (truncation, no rounding)
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
