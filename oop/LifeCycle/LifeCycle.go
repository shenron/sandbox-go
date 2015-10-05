package LifeCycle

type Sleep interface {
	Sleep() string
}

type StandUp interface {
	StandUp() string
}

type LifeCycle interface {
	Sleep
	StandUp
}

func GenericSleep(s Sleep) string {
	return s.Sleep()
}
