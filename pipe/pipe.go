package pipe

type Pipe[T any] struct {
	err error
	t   *T
}

func New[T any](t *T) *Pipe[T] {
	return &Pipe[T]{t: t}
}

func (p *Pipe[T]) SetError(err error) *T {
	p.err = err
	return p.t
}

func (p *Pipe[T]) Piping(fn func() error) *T {
	if p.err != nil {
		return p.t
	}

	if err := fn(); err != nil {
		p.err = err
	}

	return p.t
}

func (p *Pipe[T]) Error() error {
	return p.err
}
