//go:build linux

package godaemon

import (
	"os/exec"
	"syscall"
)

func startDaemonProcess(name string, args []string, attr *ProcAttr) error {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid:   true,
		Pgid:      0,
		Pdeathsig: syscall.Signal(0),
	}

	if attr != nil {
		cmd.Env = attr.Env
		cmd.Dir = attr.Dir
	}

	return cmd.Start()
}
