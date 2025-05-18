package tid_test

import (
	"testing"
	"time"

	"github.com/acudac-com/tid-go"
)

func TestUnix(t *testing.T) {
	set := map[string]struct{}{}
	i := 0
	clashFor := 0
	clashForMax := 0
	for {
		i++
		gen := tid.Unix()
		if i%10000 == 0 {
			println(i, gen)
		}
		if _, ok := set[gen]; ok {
			clashFor++
			if clashFor > clashForMax {
				clashForMax = clashFor
				println(i, "new max clashes", clashForMax)
			}
			if clashForMax > 3 {
				break
			}
		} else {
			clashFor = 0
		}
		set[gen] = struct{}{}
		time.Sleep(10 * time.Microsecond)
	}
}

func TestUnixLatestFirst(t *testing.T) {
	for {
		gen := tid.MicroLatestFirst()
		println(gen)
		time.Sleep(1 * time.Microsecond)
	}
}

func TestE(t *testing.T) {
	println(1e5 == 100000)
}
