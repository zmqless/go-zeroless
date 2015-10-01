package zeroless

type Address struct {
	Ip   string
	Port int
}

type Client struct {
	sock
	addresses []Address
}

func NewClient() *Client {
	client := Client{sock{}, make([]Address, 0)}
	client.setInit(&client)
	return &client
}

func (this Client) init(s socket) error {
	var err error

	for _, address := range this.Addresses() {
		err = connectZmqSock(s, address.Ip, address.Port)

		if err != nil {
			return err
		}
	}

	return nil
}

func (this Client) Addresses() []Address {
	return this.addresses
}

func (this *Client) Connect(ip string, port int) *Client {
	if this.Ready() {
		connectZmqSock(this.getZmqSock(), ip, port)
	}

	this.addresses = append(this.addresses, Address{ip, port})

	return this
}

func (this *Client) ConnectLocal(port int) *Client {
	return this.Connect("127.0.0.1", port)
}

func (this *Client) Disconnect(ip string, port int) *Client {
	index := -1

	for i, address := range this.addresses {
		if address.Ip == ip && address.Port == port {
			index = i
			disconnectZmqSock(this.getZmqSock(), address.Ip, address.Port)
			break
		}
	}

	if index != -1 {
		begin := this.addresses[:index]
		end := this.addresses[(index + 1):]
		this.addresses = append(begin, end...)
	}

	return this
}

func (this *Client) DisconnectLocal(port int) *Client {
	return this.Disconnect("127.0.0.1", port)
}
