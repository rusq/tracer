# Basic tracer usage example

This example shows the basic initialisation of the tracer and some `runtime/trace`
features.

## Running

Start it the usual way:
   
    go run .

## Examine the trace file

If executed with default parameters, the `trace.out` file will have been
generated.  Open it with:

    go tool trace trace.out

The browser will open, navigate to **User-defined tasks** and **User-defined
regions** to see the trace of the execution.
