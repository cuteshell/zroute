package transport

// TransType transport type
type TransType int32

// transport type
const (
	HTTP TransType = iota + 1
	Serial
	TCPClient
	TCPServer
	UDPClient
	UDPServer
)

func (t TransType) String() string {
	return [...]string{
		"Http", "Serial", "TCPClient",
		"TCPServer", "UDPClient", "UDPServer"}[t-1]
}
