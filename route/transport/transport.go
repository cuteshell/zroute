package transport

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type Transport interface {
	Reader
	Writer
	Open() error
	Close() error
}
