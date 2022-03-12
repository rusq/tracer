# tracer
Trace wrapper package - wraps the boilerplace code for the standard
`runtime/trace`.

## Getting started

Minimalistic example:
```go
import (
	"log"

	"github.com/rusq/tracer"
)

func main() {
	flag.Parse()

	if err := run("trace.out"); err != nil {
		log.Fatal(err)
	}
}

func run(tracefile string) error {
	t := tracer.New(tracefile)
	if err := t.Start(); err != nil {
		return err
	}
	defer t.Close()

    // do something

	return nil
}
```

For complete example see [example/basic](./example/basic).

## But why all this?
Adding tracing to your program is a great opportunity to get the insights on the
performance and locks.

Initialising tracing is usually done by adding the similar code to the beginning
of your `main` function:
```go
	tf, err = os.Create(tracefile)
	if err != nil {
		log.Fatalf("failed to create trace output file: %s", err)
	}
	if err := trace.Start(tf); err != nil {
		log.Fatalf("failed to start trace: %s", err)
	}
```

Then, adding similar code to the end of the main:
```go
	if trace.IsEnabled() {
		trace.Stop()
	}
	if err := tf.Close(); err != nil {
		log.Fatalf("trace file failed to close: %s", err)
	}
```

But what happens if the programme crashes and burns?  Thanks for asking!  If the
program crashes and burns, i.e. by calling the `log.Fatal()` before the trace
file is properly closed, it will be corrupt, and `go tool trace` will not be
able to parse it.

The easiest solution is to extract the logic into a run function:

```go
func main() {
    if err:=run(); err!=nil{
        log.Fatal(err)
    }
}

func run() error {
	tf, err = os.Create(tracefile)
	if err != nil {
		log.Fatalf("failed to create trace output file: %s", err)
	}
    defer tf.Close()
	if err := trace.Start(tf); err != nil {
		log.Fatalf("failed to start trace: %s", err)
	}
    defer tf.Stop()

    // call your functions
    // do your magic
    // stop all wars
    // make this world a happier place

    return nil
}
```
With this approach, unless the programme panics somewhere, `Stop` and `Close`
would be inevitably called (in this order).

This package simply wraps the boilerplate code of opening/closing the trace file
and starting/stopping the tracer.

I just got tired of adding the same code over and over and over again, that is why.
