package driver

// Driver protocol to communicate
type Driver interface {
	OnCreate() error
	OnScan() error
	Command() error
	SendFrame() error
	RecvFrame() error
	WaitResponse() error
	OnDestroy() error
}
