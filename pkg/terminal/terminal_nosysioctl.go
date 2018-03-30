// +build windows plan9 solaris

package terminal

func getWinsize() (int, int, error) {
	return 79, 24, nil
}
