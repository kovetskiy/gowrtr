package generator

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/moznion/gowrtr/internal/errmsg"
)

// Root is a code generator for the entry point.
type Root struct {
	Statements     []Statement
	Gofmt          bool
	GofmtOptions   []string
	Goimports      bool
	SyntaxChecking bool
}

// NewRoot generates a new `Root`.
func NewRoot(statements ...Statement) *Root {
	return &Root{
		Statements: statements,
	}
}

// AddStatements adds statements to Root.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) AddStatements(statements ...Statement) *Root {
	return &Root{
		Statements:     append(g.Statements, statements...),
		Gofmt:          g.Gofmt,
		GofmtOptions:   g.GofmtOptions,
		Goimports:      g.Goimports,
		SyntaxChecking: g.SyntaxChecking,
	}
}

// EnableGofmt enables `gofmt`. If `gofmt` is enabled, it applies `gofmt` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableGofmt(gofmtOptions ...string) *Root {
	return &Root{
		Statements:     g.Statements,
		Gofmt:          true,
		GofmtOptions:   gofmtOptions,
		Goimports:      g.Goimports,
		SyntaxChecking: g.SyntaxChecking,
	}
}

// EnableGoimports enables `goimports`. If `goimports` is enabled, it applies `goimports` on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableGoimports() *Root {
	return &Root{
		Statements:     g.Statements,
		Gofmt:          g.Gofmt,
		GofmtOptions:   g.GofmtOptions,
		Goimports:      true,
		SyntaxChecking: g.SyntaxChecking,
	}
}

// EnableSyntaxChecking enables syntax checking. If this option is enabled, it checks the syntax of the code on code generation phase.
// This method returns a *new* `Root`; it means this method acts as immutable.
func (g *Root) EnableSyntaxChecking() *Root {
	return &Root{
		Statements:     g.Statements,
		Gofmt:          g.Gofmt,
		GofmtOptions:   g.GofmtOptions,
		Goimports:      g.Goimports,
		SyntaxChecking: true,
	}
}

// Generate generates golang code according to registered statements.
func (g *Root) Generate(indentLevel int) (string, error) {
	generatedCode := ""

	for _, statement := range g.Statements {
		gen, err := statement.Generate(indentLevel)
		if err != nil {
			return "", err
		}
		generatedCode += gen
	}

	if g.SyntaxChecking {
		_, err := g.applyGofmt(generatedCode, "-e")
		if err != nil {
			return "", err
		}
	}

	if g.Gofmt {
		var err error
		generatedCode, err = g.applyGofmt(generatedCode, g.GofmtOptions...)
		if err != nil {
			return "", err
		}
	}

	if g.Goimports {
		var err error
		generatedCode, err = g.applyGoimports(generatedCode)
		if err != nil {
			return "", err
		}
	}

	return generatedCode, nil
}

func (g *Root) applyGofmt(generatedCode string, gofmtOptions ...string) (string, error) {
	return applyCodeFormatter(generatedCode, "gofmt", gofmtOptions...)
}

func (g *Root) applyGoimports(generatedCode string) (string, error) {
	return applyCodeFormatter(generatedCode, "goimports")
}

func applyCodeFormatter(generatedCode string, formatterCmdName string, formatterOpts ...string) (string, error) {
	echoCmd := exec.Command("echo", generatedCode)
	formatterCmd := exec.Command(formatterCmdName, formatterOpts...)

	r, w := io.Pipe()
	echoCmd.Stdout = w
	formatterCmd.Stdin = r

	var out, errout bytes.Buffer
	formatterCmd.Stdout = &out
	formatterCmd.Stderr = &errout

	echoCmd.Start()
	if err := formatterCmd.Start(); err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}
	echoCmd.Wait()
	w.Close()
	err := formatterCmd.Wait()

	if err != nil {
		cmds := []string{formatterCmdName}
		return "", errmsg.CodeFormatterError(strings.Join(append(cmds, formatterOpts...), " "), errout.String(), err)
	}

	return out.String(), err
}