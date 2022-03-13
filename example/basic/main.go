package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"runtime/trace"

	"github.com/rusq/tracer"
)

var tracefile = flag.String("t", "trace.out", "trace `filename`")

func main() {
	flag.Parse()

	if err := run(*tracefile); err != nil {
		log.Fatal(err)
	}
}

func run(tracefile string) error {
	// initialise the Tracer
	t := tracer.New(tracefile)
	if err := t.Start(); err != nil {
		return err
	}
	// Close, closes the trace file, if not called, trace
	// file most likely will be corrupt.  You can also use
	// t.End()
	defer t.Close()

	// do something
	fmt.Println(add(context.Background(), 2, 2))

	return nil
}

// add is the sample function that shows trace tasks, logging, and regions.
func add(ctx context.Context, xs ...int) int {
	ctx, task := trace.NewTask(ctx, "add")
	defer task.End()

	var result int
	trace.WithRegion(ctx, "quickmath", func() {
		for _, x := range xs {
			result += x
		}
	})
	trace.Logf(ctx, "info", "len(xs)=%d, res=%d", len(xs), result)

	return result
}
