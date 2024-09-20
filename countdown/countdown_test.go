package countdown

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	sleepInterval := 3 * time.Second
	startInt := 3
	t.Run("testing output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		/*
			init a buffer, so that the prints from Countdown can be
			written to it
		*/
		sleeper := emptySleeper{}
		Countdown(buffer, sleeper, 3)
		got := buffer.String()
		want := "3...\n2...\n1...\nGO!!!"
		if got != want {
			generateErrorMessage(t, got, want)
		}
	})
	t.Run("testing sleep accumulation", func(t *testing.T) {

		sleepAccumulator := sleepTimeAccumulation{
			totalSleepTime: 0 * time.Second,
		}
		sleeper := &ConfigurableSleeper{
			Duration:  sleepInterval,
			SleepFunc: sleepAccumulator.addSleep,
		}
		buffer := &bytes.Buffer{}
		Countdown(buffer, sleeper, startInt)
		got := int(sleepAccumulator.totalSleepTime)
		want := startInt * int(sleepInterval)
		if got != want {
			t.Errorf("got: %s\nwant: %s",
				convertIntToString(got),
				convertIntToString(want))
		}
	})

	t.Run("testing order of operations", func(t *testing.T) {
		writeSpy := writerSpy{}
		Countdown(&writeSpy, &writeSpy, startInt)
		got := writeSpy.ops
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}

type writerSpy struct {
	ops []string
}

func (w *writerSpy) Write(p []byte) (n int, err error) {
	buffer := bytes.Buffer{}
	n, err = buffer.Write(p)
	w.ops = append(w.ops, "write")
	return n, err
}

func (w *writerSpy) Sleep() {
	w.ops = append(w.ops, "sleep")
}

type sleepTimeAccumulation struct {
	totalSleepTime time.Duration
}

func (s *sleepTimeAccumulation) addSleep(duration time.Duration) {
	s.totalSleepTime += duration
}

func generateErrorMessage(t testing.TB, got, want string) {
	t.Helper()
	t.Errorf("got: %s\nwant:%s", got, want)
}

func convertIntToString(i int) (s string) {
	return fmt.Sprintf("%d", i)
}
