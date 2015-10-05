package LifeCycle

type Sleep interface {
	Sleep() string
}

type WakeUp interface {
	WakeUp() string
}

type LifeCycle interface {
	Sleep
	WakeUp
}

func GenericSleep(s Sleep) string {
	return s.Sleep()
}

func GenericWakeUp(w WakeUp) string {
	return w.WakeUp()
}
