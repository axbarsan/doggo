package runner

import (
	"fmt"
	"io"
	"strings"

	"github.com/axbarsan/doggo/internal/evaluator"
	"github.com/axbarsan/doggo/internal/lexer"
	"github.com/axbarsan/doggo/internal/object"
	"github.com/axbarsan/doggo/internal/parser"
)

type Runner struct {
	env *object.Environment
}

func New() *Runner {
	env := object.NewEnvironment()

	r := &Runner{
		env: env,
	}

	return r
}

func (r *Runner) Run(code string) string {
	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return getParserErrors(p.Errors())
	}

	evaluated := evaluator.Eval(program, r.env)
	if evaluated != nil {
		return evaluated.Inspect()
	}

	return ""
}

func getParserErrors(errors []string) string {
	buf := new(strings.Builder)

	io.WriteString(buf, fmt.Sprintf("There are %d errors in your code.\n", len(errors)))
	io.WriteString(buf, " parser errors: \n")
	for _, msg := range errors {
		io.WriteString(buf, fmt.Sprintf("\t%s\n", msg))
	}

	return buf.String()
}
