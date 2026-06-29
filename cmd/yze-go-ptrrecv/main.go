// Command yze-go-ptrrecv runs the ptrrecv analyzer as a standalone go/analysis
// checker (text, -json, and -fix output, and as a `go vet -vettool`).
package main

import (
	ptrrecv "github.com/gomatic/yze-go-ptrrecv"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// run is the analysis entry point, indirected so the binary's wiring is testable
// without invoking the real driver (which loads packages and exits the process).
var run = singlechecker.Main

func main() { run(ptrrecv.Analyzer) }
