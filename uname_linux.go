// +build linux

package system

import "syscall"

func charsToString(ca [65]int8) string {
	s := make([]byte, len(ca))
	var lens int
	for ; lens < len(ca); lens++ {
		if ca[lens] == 0 {
			break
		}
		s[lens] = uint8(ca[lens])
	}
	return string(s[0:lens])
}

func uname() (Utsname, error) {
	out := Utsname{}
	n := syscall.Utsname{}
	if err := syscall.Uname(&n); err != nil {
		return out, err
	}

	out.Sysname = charsToString(n.Sysname)
	out.Nodename = charsToString(n.Nodename)
	out.Release = charsToString(n.Release)
	out.Version = charsToString(n.Version)
	out.Machine = charsToString(n.Machine)
	out.Domainname = charsToString(n.Domainname)

	return out, nil
}
