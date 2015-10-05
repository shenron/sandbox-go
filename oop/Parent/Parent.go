package Parent

type Parent struct {
	firstname, name string
}

func New(firstname, name string) *Parent {
	return &Parent{firstname: firstname, name: name}
}

func (p Parent) GetName() string {
	return p.name
}

func (p Parent) GetFirstName() string {
	return p.firstname
}

func (p Parent) Sleep() string {
	return "late"
}

func (p Parent) WakeUp() string {
	return "early"
}
