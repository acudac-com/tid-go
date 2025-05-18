package tid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Uses the seconds of the current time to apply a year delta to the returned
// time. This helps with handling bursts of id generations without any
// conflicts.
func timeWithYearJumps() time.Time {
	t := time.Now()
	dur := time.Duration((-30+t.Second())*1e9) * 3600 * 24 * 365
	return t.Add(dur)
}

// Returns a random unix time based id. Includes jitter to prevent db hotspots
// and handle bursts of new ids being generated.
func Unix() string {
	return rev(base36(jitter(timeWithYearJumps().Unix())))
}

// Returns a random unix millisecond time based id. Includes jitter to prevent
// db hotspots and handle bursts of new ids being generated.
func Milli() string {
	return rev(base36(jitter(timeWithYearJumps().UnixMilli())))
}

// Returns a random unix microsecond time based id. Includes jitter to prevent
// db hotspots and handle bursts of new ids being generated.
func Micro() string {
	return rev(base36(jitter(timeWithYearJumps().UnixMicro())))
}

// Returns a random unix microsecond time based id. Includes jitter to prevent
// db hotspots and handle bursts of new ids being generated.
func Nano() string {
	return rev(base36(jitter(timeWithYearJumps().UnixNano())))
}

func jitter(value int64) int64 {
	dif := -1e6 + rand.Intn(2e6)
	value = value + int64(dif)
	return value
}

func base36(value int64) string {
	return strconv.FormatInt(value, 36)
}

func rev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

const (
	Y5138Unix  int64 = 99999999999       // 19xtf1tr = 8 chars
	Y5138Milli int64 = 99999999999000    // zg3d62qe0 = 9 chars
	Y5138Micro int64 = 99999999999000000 // rcn1hsrx0w0 = 11 chars
)

// Returns a unix time based id that ensures later values are lexacographically
// smaller to appear first when listed from a database.
func UnixLatestFirst() string {
	dif := Y5138Unix - time.Now().Unix()
	return fmt.Sprintf("%08s", base36(dif))
}

// Returns a unix millisecond time based id that ensures later values are
// lexacographically smaller to appear first when listed from a database.
func MilliLatestFirst() string {
	dif := Y5138Milli - time.Now().UnixMilli()
	return fmt.Sprintf("%09s", base36(dif))
}

// Returns a unix microsecond time based id that ensures later values are
// lexacographically smaller to appear first when listed from a database.
func MicroLatestFirst() string {
	dif := Y5138Micro - time.Now().UnixMicro()
	return fmt.Sprintf("%011s", base36(dif))
}
