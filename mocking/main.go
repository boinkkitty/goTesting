package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Interface for dependency injection
type Sleeper interface {
	Sleep()
}

// SpySleeper tracks calls
type SpySleeper struct {
	Calls int
}

// Implement sleeper method
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Default Sleeper with 1 unit of time
type DefaultSleeper struct{}

// Sleep for 1 second
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Spy that tracks operations for countdown
type SpyCountdownOperations struct {
	Calls []string
}

// Append sleep when sleeping
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Append write when writing
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Mock
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration) // spytime's sleep
}

type SpyTime struct {
	durationSlept time.Duration
}

// Spy takes in a duration
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Also a sleeper! it sleeps for its own duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}
