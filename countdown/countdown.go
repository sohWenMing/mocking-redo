package countdown

import (
	"fmt"
	"io"
	"time"
)

func Countdown(w io.Writer, s Sleeper, startInt int) {
	currentInt := startInt
	for {
		if currentInt > 0 {
			w.Write([]byte(generateCountString(currentInt)))
			s.Sleep()
			currentInt--
		} else {
			w.Write([]byte("GO!!!"))
			break
		}
	}
}

func generateCountString(currentInt int) (countString string) {
	return fmt.Sprintf("%d...\n", currentInt)
}

type Sleeper interface {
	Sleep()
}

type emptySleeper struct{}

func (e emptySleeper) Sleep() {
	return
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

var MainSleeper = ConfigurableSleeper{
	Duration:  1 * time.Second,
	SleepFunc: time.Sleep,
}
