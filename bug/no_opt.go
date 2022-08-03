package bug

type NoOpt struct{}

func NewNoOpt() *NoOpt {
	return &NoOpt{}
}

func (n *NoOpt) Ensure(error) {}
