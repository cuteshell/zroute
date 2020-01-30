package common

type Closable interface {
	Close() error
}

type Runnable interface {
	Start() error

	Closable
}
