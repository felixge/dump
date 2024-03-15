// samples-to-call-counts is a tool to explore some haphazard claims I made
// about the nyquist theorem and the number of samples required to accurately
// estimate the number of times a function is called.

package main

import (
	"cmp"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"slices"
	"strings"
	"time"
)

func main() {
	cpuF := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	if *cpuF != "" {
		f, err := os.Create(*cpuF)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	start := time.Now()
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	simDuration := 10 * time.Minute
	samplePeriod := 10 * time.Millisecond
	steps := time.Duration(50)

	cw := csv.NewWriter(os.Stdout)
	cw.Write([]string{"call duration (ms)", "error A (%)"})

	for callDur := samplePeriod / steps; callDur <= samplePeriod*2; callDur += samplePeriod / steps {
		callDist := NormalDistribution(r, callDur, callDur/10)
		a := &Application{loop: []*Call{
			{Function: "A", Duration: callDist},
			{Function: "B", Duration: callDist},
		}}
		p := &Profiler{app: a, period: NormalDistribution(r, 10*time.Millisecond, 1*time.Microsecond)}
		q := &Queue{}
		q.Push(a.Run(0)...)
		q.Push(p.Sample(0)...)
		simulate(q, simDuration)
		errorA := float64(a.CallStats.calls["A"]-p.CallStats.calls["A"]) / float64(a.CallStats.calls["A"])
		cw.Write([]string{
			fmt.Sprintf("%f", callDur.Seconds()*1000),
			fmt.Sprintf("%v", errorA*100),
		})

		// fmt.Printf("seed: %d\n", seed)
		// fmt.Printf("app: %s\n", a.CallStats.String())
		// fmt.Printf("profiler: %s\n", p.CallStats.String())
		// fmt.Printf("error A: %.1f%%\n", errorA*100)
		// fmt.Printf("sim duration: %s\n", simDuration)
		// fmt.Printf("real duration: %s\n", time.Since(start))
		cw.Flush()
	}
	dt := time.Since(start)
	fmt.Printf("dt: %v\n", dt)
}

// simulate runs the simulation for the given queue of events, stopping after
// maxT time has passed.
func simulate(q *Queue, maxT time.Duration) {
	for q.Len() > 0 && q.NextTime() < maxT {
		e := q.Pop()
		q.Push(e.Run(e.Time)...)
	}
}

type Application struct {
	loop      []*Call
	idx       int
	CallStats CallStats
}

func (a *Application) Run(t time.Duration) []Event {
	call := a.loop[a.idx%len(a.loop)]
	a.CallStats.Inc(call.Function)
	a.idx++
	return []Event{{Time: t + call.Duration(), Run: a.Run}}
}

// Call returns the currently executing call.
func (a *Application) Call() *Call {
	return a.loop[(a.idx-1)%len(a.loop)]
}

type Profiler struct {
	app       *Application
	period    func() time.Duration
	prevCall  string
	CallStats CallStats
}

func (p *Profiler) Sample(t time.Duration) []Event {
	call := p.app.Call()
	if call.Function != p.prevCall {
		p.CallStats.Inc(call.Function)
		p.prevCall = call.Function
	}
	return []Event{{Time: t + p.period(), Run: p.Sample}}
}

type CallStats struct {
	calls map[string]int
}

func (c *CallStats) Inc(fn string) {
	if c.calls == nil {
		c.calls = map[string]int{}
	}
	c.calls[fn]++
}

func (c CallStats) String() string {
	var keys []string
	for k := range c.calls {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	var calls []string
	for _, k := range keys {
		count := c.calls[k]
		calls = append(calls, fmt.Sprintf("%s: %d", k, count))
	}
	return strings.Join(calls, " ")
}

type Queue struct {
	events []Event
	dead   []bool
}

// NextTime returns the time of the next event in the queue.
func (q *Queue) NextTime() time.Duration {
	return q.events[0].Time
}

// Push adds the given events to the queue.
func (q *Queue) Push(events ...Event) {
	q.events = append(q.events, events...)
	// TODO: use a heap data structure
	slices.SortFunc(q.events, func(a, b Event) int {
		return cmp.Compare(a.Time, b.Time)
	})
}

// Pop removes and returns the next event from the queue.
func (q *Queue) Pop() Event {
	e := q.events[0]
	q.events = q.events[1:]
	return e
}

// Len returns the number of events in the queue.
func (q *Queue) Len() int {
	return len(q.events)
}

// Event represents a simulated event.
type Event struct {
	// Time is the offset of the event from the start of the simulation.
	Time time.Duration
	// Run executes the event and returns the next set of events that should
	// occur in response.
	Run func(t time.Duration) []Event
}

// Loop simulates a series of function calls
type Loop struct {
	// Calls is the list of function calls that occurred during the loop.
	Calls []Call
}

// Call represents a simulated function call.
type Call struct {
	// Function is the name of the function being called.
	Function string
	// Duration is the amount of time the function took to execute.
	Duration func() time.Duration
}

// FixedDuration returns a function that always returns the duration d.
func FixedDuration(d time.Duration) func() time.Duration {
	return func() time.Duration { return d }
}

// NormalDistribution returns a function that returns a duration from a normal
// distribution with the given mean and standard deviation.
func NormalDistribution(r *rand.Rand, mean, stdev time.Duration) func() time.Duration {
	return func() time.Duration {
		return time.Duration(rand.NormFloat64()*float64(stdev) + float64(mean))
	}
}
