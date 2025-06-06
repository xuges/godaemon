package godaemon

import (
	"os/exec"
	"syscall"
)

func startDaemonProcess(name string, args []string, attr *ProcAttr) error {
	const DETACHED_PROCESS = 0x00000008

	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags:    DETACHED_PROCESS,
		NoInheritHandles: true,
	}

	if attr != nil {
		cmd.Env = attr.Env
		cmd.Dir = attr.Dir
	}

	return cmd.Start()
}
