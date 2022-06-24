// Package tracer wraps the boilerplace code for runtime/trace.
package tracer

import (
	"fmt"
	"os"
	"runtime/trace"
)

type Info struct {
	Filename string

	tf *os.File
}

// New creates a new tracer instance.  Tracer should be started with
// Start.
func New(tracefile string) *Info {
	return &Info{Filename: tracefile}
}

// Start starts tracing.  If the Tracer was initialised with an empty filename,
// Start it will not start the trace and will return nil.
func (t *Info) Start() error {
	if t.Filename == "" {
		return nil
	}

	var err error
	t.tf, err = os.Create(t.Filename)
	if err != nil {
		return fmt.Errorf("failed to create trace output file: %w", err)
	}
	if err := trace.Start(t.tf); err != nil {
		return fmt.Errorf("failed to start trace: %w", err)
	}
	return nil
}

// Close ends the trace session (it's a wrapper around End).
func (t *Info) Close() error {
	return t.End()
}

// End ends the trace session.  If the tracing has not been initialised, it does
// nothing and returns nil.
func (t *Info) End() (err error) {
	defer func() {
		// sometimes tracer panics, we don't want to crash the caller.
		if r := recover(); r != nil {
			err = fmt.Errorf("trace panic recovered: %v", r)
		}
	}()
	if t.tf == nil {
		return
	}
	if trace.IsEnabled() {
		trace.Stop()
	}
	if err := t.tf.Close(); err != nil {
		return err
	}
	t.tf = nil
	return
}
