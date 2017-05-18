package system 

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