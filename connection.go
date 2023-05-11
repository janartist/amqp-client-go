package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"net"
	"time"
)

// NewConnection 全新tcp连接
func NewConnection(uri string) (*Connection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return &Connection{conn: conn}, err
}

type Connection struct {
	Uri  string
	conn *amqp.Connection
}

func (c *Connection) Read(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) Write(b []byte) (n int, err error) {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) Close() error {
	return c.GetConnection().Close()
}

func (c *Connection) LocalAddr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) RemoteAddr() net.Addr {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) SetDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) SetReadDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) SetWriteDeadline(t time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (c *Connection) GetConnection() *amqp.Connection {
	return c.conn
}

func (c *Connection) NewChannel() (*channel, error) {
	//new open channel
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, err
	}
	return &channel{c, ch}, nil
}
