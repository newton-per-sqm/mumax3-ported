package engine

import (
	"fmt"
	"github.com/mumax/3/util"
	"time"
)

func init() {
	DeclFunc("Run", Run, "Run the simulation for a time in seconds")
	DeclFunc("Steps", Steps, "Run the simulation for a number of time steps")
	DeclFunc("RunWhile", RunWhile, "Run while condition function is true")
	DeclFunc("SetSolver", SetSolver, "Set solver type. 1:Euler, 2:Heun")
	DeclVar("t", &Time, "Total simulated time (s)")
	DeclVar("step", &NSteps, "Total number of time steps taken")
	DeclROnly("dt", &Dt_si, "Last solver time step (s)")
	DeclVar("MinDt", &MinDt, "Minimum time step the solver can take (s)")
	DeclVar("MaxDt", &MaxDt, "Maximum time step the solver can take (s)")
	DeclVar("MaxErr", &MaxErr, "Maximum error per step the solver can tolerate")
	DeclVar("Headroom", &Headroom, "Solver headroom")
	DeclVar("FixDt", &FixDt, "Set a fixed time step, 0 disables fixed step")
	SetSolver(HEUN)
}

var (
	//Solver     = NewSolver(Torque.Set, M.normalize, 1e-15, &GammaLL, HeunStep)
	Time       float64             // time in seconds
	pause      = true              // set pause at any time to stop running after the current step
	postStep   []func()            // called on after every time step
	Inject     = make(chan func()) // injects code in between time steps. Used by web interface.
	solvertype int
)

func SetSolver(typ int) {
	switch typ {
	default:
		util.Fatalf("SetSolver: unknown solver type: %v", typ)
	case 1:
		stepper = EulerStep
	case 2:
		stepper = HeunStep
	}
	solvertype = typ
}

const (
	EULER = 1
	HEUN  = 2
)

// Run the simulation for a number of seconds.
func Run(seconds float64) {
	//GUI.Set("runtime", seconds)
	stop := Time + seconds
	RunWhile(func() bool { return Time < stop })
}

// Run the simulation for a number of steps.
func Steps(n int) {
	//GUI.Set("runsteps", n)
	stop := NSteps + n
	RunWhile(func() bool { return NSteps < stop })
}

// Runs as long as condition returns true.
func RunWhile(condition func() bool) {
	//GUI.Set("maxtorque", "--") // only updated on page refresh, avoid showing out-of-date value
	pause = false
	for condition() && !pause {
		select {
		default:
			step()
		// accept tasks form Inject channel
		case f := <-Inject:
			f()
		}
	}
	pause = true
}

func break_() {
	pause = true
}

// exit finished simulation this long after browser was closed
var Timeout = 3 * time.Second

func RunInteractive() {
	gui_.RunInteractive()
}

func nop() {}

// Enter interactive mode. Simulation is now exclusively controlled by web GUI
func (g *guistate) RunInteractive() {

	// periodically wake up Run so it may exit on timeout
	go func() {
		for {
			Inject <- nop
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("entering interactive mode")
	g.UpdateKeepAlive()
	for time.Since(g.KeepAlive()) < Timeout {
		f := <-Inject
		f()
	}
	fmt.Println("browser disconnected, exiting")
}

func step() {
	Step(M.Buffer())
	for _, f := range postStep {
		f()
	}
	DoOutput()
}

// Register function f to be called after every time step.
// Typically used, e.g., to manipulate the magnetization.
func PostStep(f func()) {
	postStep = append(postStep, f)
}

// inject code into engine and wait for it to complete.
func InjectAndWait(task func()) {
	ready := make(chan int)
	Inject <- func() { task(); ready <- 1 }
	<-ready
}
