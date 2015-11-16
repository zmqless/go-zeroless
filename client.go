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

func (this *Client) Connect(ip string, port int) error {
	if !this.Ready() {
		this.addresses = append(this.addresses, Address{ip, port})
		return nil
	}

	return connectZmqSock(this.getZmqSock(), ip, port)
}

func (this *Client) ConnectLocal(port int) error {
	return this.Connect("127.0.0.1", port)
}

func (this *Client) Disconnect(ip string, port int) error {
	index := -1
	var err error

	for i, address := range this.addresses {
		if address.Ip == ip && address.Port == port {
			index = i
			err = disconnectZmqSock(this.getZmqSock(), address.Ip, address.Port)

			if err != nil {
				return err
			}

			break
		}
	}

	if index != -1 {
		begin := this.addresses[:index]
		end := this.addresses[(index + 1):]
		this.addresses = append(begin, end...)
	}

	return nil
}

func (this *Client) DisconnectLocal(port int) error {
	return this.Disconnect("127.0.0.1", port)
}
