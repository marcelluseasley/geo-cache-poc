# MPX Caching Test

This small program simulates MPX caching.

It accepts a filename, optional TTL, and optional precision of decimal places as flags and lists the number of "new" vs. "duplicate" hits to MPX. If TTL is set to `-1`, then it isn't taken into account and ANY duplicate entry will be counted as such, regardless of timestamp.

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
```
### Note, by manually changing TTL from minutes(`time.Minute`) to seconds(`time.Second`) (for testing purposes), the number of new entries really increases as precision increases and TTL decreases.