package system

type Utsname struct {
	Sysname    string
	Nodename   string
	Release    string
	Version    string
	Machine    string
	Domainname string
}

func Uname() (Utsname, error) {
	return uname()
}
