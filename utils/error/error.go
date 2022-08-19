package error

type newErr struct {
	Err string
}

func (n newErr) Error() string {
	return n.Err
}

func NotFound(err string) error {
	return newErr{Err: err}
}
