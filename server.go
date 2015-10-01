package zeroless

type Server struct {
	sock
	port int
}

func NewServer(port int) *Server {
	server := Server{sock{}, port}
	server.setInit(&server)
	return &server
}

func (this Server) init(s socket) error {
	return bindZmqSock(s, this.port)
}
