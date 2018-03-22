// +build !windows,!plan9,!solaris

package terminal

import (
	"os"

	"golang.org/x/sys/unix"
)

func getWinsize() (int, int, error) {

	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, os.NewSyscallError("GetWinsize", err)
	}

	return int(ws.Col), int(ws.Row), nil
}
