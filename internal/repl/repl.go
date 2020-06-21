package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/axbarsan/doggo/internal/runner"
)

// REPL stands for 'Read Eval Print Loop'.
// This should read the input, send it to
// the interpreter for evaluation, and
// print the result/output of the interpreter.

const PROMPT = ">> "

// Start parses each line of the file and returns
// the result to the output stream.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	r := runner.New()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		output := r.Run(line)

		io.WriteString(out, output)
		if output != "" {
			io.WriteString(out, "\n")
		}
	}
}
