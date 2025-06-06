package godaemon

type ProcAttr struct {
	Dir string
	Env []string
}

func StartDaemonProcess(name string, args []string, attr *ProcAttr) error {
	return startDaemonProcess(name, args, attr)
}
