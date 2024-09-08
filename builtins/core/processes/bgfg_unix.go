//go:build !windows && !plan9 && !js
// +build !windows,!plan9,!js

package processes

import (
	"errors"
	"os"
	"syscall"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/state"
	"github.com/lmorg/murex/lang/types"
	"golang.org/x/sys/unix"
)

func mkbg(p *lang.Process) error {
	fid, err := p.Parameters.Uint32(0)
	if err != nil {
		return errors.New("invalid parameters. Expecting either a code block or FID of a stopped process")
	}

	f, err := lang.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	if f.State.Get() != state.Stopped {
		return errors.New("FID is not a stopped process. Run `jobs` or `fid-list` to see a list of stopped processes")
	}

	if f.SystemProcess.External() {
		err = f.SystemProcess.Signal(syscall.SIGCONT)
		if err != nil {
			return err
		}
	}
	go unstop(f)

	updateTree(f, true)

	f.State.Set(state.Executing)

	lang.ShowPrompt <- true
	return nil
}

func cmdForeground(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	fid, err := p.Parameters.Uint32(0)
	if err != nil {
		return err
	}

	f, err := lang.GlobalFIDs.Proc(fid)
	if err != nil {
		return err
	}

	lang.HidePrompt <- true
	go unstop(f)
	updateTree(f, false)

	lang.ForegroundProc.Set(f)
	f.State.Set(state.Executing)

	if f.SystemProcess.External() {
		//if p.SystemProcess.ForcedTTY() {
		lang.UnixPidToFg(f)
		//}
		err = f.SystemProcess.Signal(syscall.SIGCONT)
		if err != nil {
			return err
		}
		unix.IoctlSetPointerInt(int(os.Stdin.Fd()), unix.TIOCSPGRP, f.SystemProcess.Pid())
	}

	<-f.Context.Done()
	return nil
}

func unstop(p *lang.Process) {
	p.WaitForStopped <- true
}
