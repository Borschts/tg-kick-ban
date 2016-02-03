package backend

type Backend struct {
	conn int
}

func NewBackend(conn_str string) (Backend, error) {
	conn := 0
	return Backend{
		conn: conn,
	}, nil
}

func (b *Backend) Close() error {
	return nil
}
