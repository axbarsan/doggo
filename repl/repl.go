package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/axbarsan/doggo/lexer"
	"github.com/axbarsan/doggo/parser"
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

	for {
		_, _ = fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())

			continue
		}

		_, _ = io.WriteString(out, program.String())
		_, _ = io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, _ = io.WriteString(out, "\t"+msg+"\n")
	}
}
