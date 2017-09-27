package io

import (
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["echo"] = cmdOut
	proc.GoFunctions["out"] = cmdOut
	proc.GoFunctions["err"] = cmdErr
	proc.GoFunctions["print"] = cmdPrint
}

func cmdOut(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	if f, _ := p.Parameters.String(0); f == "-n" {
		_, err = p.Stdout.Write(p.Parameters.ByteAllRange(0, -1))
		return
	}
	_, err = p.Stdout.Writeln(p.Parameters.ByteAll())
	return
}

func cmdErr(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)
	p.ExitNum = 1
	if f, _ := p.Parameters.String(0); f == "-n" {
		_, err = p.Stdout.Write(p.Parameters.ByteAllRange(0, -1))
		return
	}
	_, err = p.Stderr.Writeln(p.Parameters.ByteAll())
	return
}

func cmdPrint(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)
	_, err = fmt.Println(p.Parameters.ByteAll())
	return
}

/*
func cmdTimeStamp(pid string) (err error) {
	//out.StdOut =
	//return
}
*/