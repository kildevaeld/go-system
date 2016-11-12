// +build darwin freebsd netbsd openbsd windows !linux

package system

func uname() (Utsname, error) {
	return Utsname{}, ErrCouldNotDetectUname
}
