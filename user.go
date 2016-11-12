package system

import (
	"os/user"
	"regexp"
)

func CurrentUserName() (string, error) {
	return userCurrent()
}

var reg = regexp.MustCompile("\\d+")

func GetUser(name string) (*user.User, error) {

	if name == "" {
		return user.Current()
	}

	if reg.Match([]byte(name)) {
		return user.LookupId(name)
	}

	return user.Lookup(name)

}
