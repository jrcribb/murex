package typemgmt

import (
	"errors"
	"os"
	"regexp"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["set"] = cmdSet
	proc.GoFunctions["!set"] = cmdUnset
	proc.GoFunctions["global"] = cmdGlobal
	proc.GoFunctions["!global"] = cmdUnglobal
	proc.GoFunctions["export"] = cmdExport
	proc.GoFunctions["!export"] = cmdUnexport
	proc.GoFunctions["unset"] = cmdUnexport
}

var (
	rxSet     = regexp.MustCompile(`(?sm)^([_a-zA-Z0-9]+)\s*=(.*$)`)
	rxVarName = regexp.MustCompile(`^([_a-zA-Z0-9]+)$`)
)

func cmdSet(p *proc.Process) error      { return set(p, p) }
func cmdUnset(p *proc.Process) error    { return unset(p, p) }
func cmdGlobal(p *proc.Process) error   { return set(p, proc.ShellProcess) }
func cmdUnglobal(p *proc.Process) error { return unset(p, proc.ShellProcess) }

func set(p *proc.Process, scope *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	//var overrideDataType bool
	dataType, _ := p.Parameters.String(0)
	var params string

	switch p.Parameters.Len() {
	case 0:
		return errors.New("Missing variable name")
	case 1:
		if p.IsMethod {
			dataType = p.Stdin.GetDataType()
		} else {
			dataType = types.String
		}
		params, _ = p.Parameters.String(0)
	case 2:
		params, _ = p.Parameters.String(1)
	case 3:
		return errors.New("Too many parameters. Have you quoted the variable data using either single quotes, double quotes, or parentheses?")
	}

	// Set variable as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `set` / `global` as method")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		b = utils.CrLfTrim(b)
		return scope.Parent.Variables.Set(params, string(b), dataType)
	}

	// Set variable as parameters:
	if rxVarName.MatchString(params) {
		return scope.Parent.Variables.Set(params, "", dataType)
	}

	// Define an empty variable
	match := rxSet.FindAllStringSubmatch(params, -1)
	return scope.Parent.Variables.Set(match[0][1], match[0][2], dataType)
}

func unset(p *proc.Process, scope *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = scope.Variables.Unset(varName)
	return err
}

func cmdExport(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
	}

	params := p.Parameters.StringAll()

	// Set env as method:
	if p.IsMethod {
		if !rxVarName.MatchString(params) {
			return errors.New("Invalid variable name; unexpected parameters for calling `export` as method")
		}
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		return os.Setenv(params, string(b))
	}

	// Set env as parameters:
	if rxVarName.MatchString(params) {
		return os.Setenv(params, "")
	}

	match := rxSet.FindAllStringSubmatch(params, -1)
	return os.Setenv(match[0][1], match[0][2])
}

func cmdUnexport(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing variable name")
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = os.Unsetenv(varName)
	return err
}
