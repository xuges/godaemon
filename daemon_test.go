package godaemon

import (
	"runtime"
	"testing"
)

func Test_StartDaemonProcess(t *testing.T) {
	switch runtime.GOOS {
	case "windows":
		err := StartDaemonProcess("notepad.exe", nil, nil)
		if err != nil {
			t.Fatal(err)
		}

	case "linux":
		err := StartDaemonProcess("gedit", nil, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}
