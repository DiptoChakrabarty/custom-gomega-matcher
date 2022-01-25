package human

type Person struct {
	Age       int
	FirstName string
	LastName  string
}

func (p *Person) IsChild() bool {
	return p.Age < 10
}
