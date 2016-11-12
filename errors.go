package system

import "errors"

var (
	ErrCouldNotDetectUsername = errors.New("Could not detect default username")
	ErrCouldNotDetectUname    = errors.New("Could get uname from current platform")
)
