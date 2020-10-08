# MPX Caching Test

This small program simulates MPX caching.

It accepts a filename, optional TTL, and optional precision of decimal places as flags and lists the number of "new" vs. "duplicate" hits to MPX.

## Default values
- filename: `log.txt`
- TTL: `3`
- precision: `2`

### Example
```
$ go run main.go -h
Usage of main:
    -f string
        name of log file (default "log.txt")
    -p int
        precision of latitude/longitude (default 2)
    -ttl int
        cache TTL (default 3)
```

```
$ go run main.go -f=log.txt -p 3 -ttl=3
Cache Hits / New MPX requests (lat/long precision of 3) and TTL: 3 minutes
New: 200        Duplicate: 85 