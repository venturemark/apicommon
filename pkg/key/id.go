package key

type ID struct {
	f float64
	s string
}

func (i *ID) F() float64 {
	return i.f
}

func (i *ID) S() string {
	return i.s
}
