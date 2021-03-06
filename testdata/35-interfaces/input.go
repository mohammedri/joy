package main

var _ Component = (*component)(nil)

// Component interface
type Component interface {
	String() string
}

type component struct {
	component string
}

func (c component) String() string {
	return c.component
}

// New component
func New(name string) Component {
	return &component{name}
}

func main() {
	c := New("awesomeness")
	println(c.String())
}
