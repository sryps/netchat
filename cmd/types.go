package cmd

var port int
var addr string

type ProtocolType int

const (
	Unknown ProtocolType = iota
	GRPC
	HTTP
	TCP
)

func (p ProtocolType) String() string {
	switch p {
	case GRPC:
		return "gRPC"
	case HTTP:
		return "HTTP"
	case TCP:
		return "TCP"
	default:
		return "Unknown"
	}
}
