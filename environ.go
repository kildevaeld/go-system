package system

import (
	"errors"
	"regexp"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
)

type Environ []string

func (self *Environ) Add(env ...string) {
	*self = append(*self, env...)
}

func (self Environ) ToMap() map[string]string {
	env := make(map[string]string)
	for _, e := range self {
		a := strings.SplitN(e, "=", 2)
		env[a[0]] = a[1]
	}
	return env
}

func (self Environ) Get(key string) string {
	for _, e := range self {
		i := strings.Index(e, "=")
		if key == e[:i] {
			return e[:i]
		}
	}
	return ""
}

func (self Environ) Has(key string) bool {
	return self.Get(key) != ""
}

var regu = regexp.MustCompile("\\$([a-zA-Z_]+)")

func (self Environ) Expand(str string) string {
	return regu.ReplaceAllStringFunc(str, func(str string) string {
		e := self.Get(str[1:])
		if e == "" {
			return str
		}
		return e
	})
}

func (self Environ) TryExpand(str string) (string, error) {
	var err error
	str = reg.ReplaceAllStringFunc(str, func(str string) string {
		e := self.Get(str[1:])
		if e == "" {
			err = multierror.Append(err, errors.New("could not expand "+str))
		}
		return e
	})

	return str, err
}

func MapToEnviron(m map[string]string) Environ {
	var out Environ
	for k, v := range m {
		out = append(out, k+"="+v)
	}
	return out
}

func MergeEnviron(env ...Environ) Environ {
	out := make(map[string]string)
	for _, e := range env {
		m := e.ToMap()
		for k, v := range m {
			out[k] = v
		}
	}
	return MapToEnviron(out)
}
