package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/axbarsan/doggo/evaluator"
	"github.com/axbarsan/doggo/lexer"
	"github.com/axbarsan/doggo/object"
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
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, fmt.Sprintf("There are %d errors in your code.\n", len(errors)))
	io.WriteString(out, " parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
