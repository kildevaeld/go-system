package system 

import (
	"strings"
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