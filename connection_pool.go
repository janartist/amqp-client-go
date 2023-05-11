package rabbitmq

import (
	"log"
	"net"

	"rabbitmq-client/pool"
)

type Rabbitmq struct {
	Uri        string `json:"uri" yaml:"uri"`
	InitialCap int    `json:"initial_cap" yaml:"initial_cap"`
	MaxCap     int    `json:"max_cap" yaml:"max_cap"`
	Pool       pool.Pool
}

func connectionPoolFactory(uri string) pool.Factory {
	return func() (net.Conn, error) {
		return NewConnection(uri)
	}
}

// NewConnection 全新Rabbitmq实例
func NewRabbitmq(r *Rabbitmq) (*Rabbitmq, error) {
	p, err := pool.NewChannelPool(r.InitialCap, r.MaxCap, connectionPoolFactory(r.Uri))
	if err != nil {
		log.Fatal(err)
	}
	r.Pool = p
	return r, nil
}

func (r *Rabbitmq) GetConnection() (*Connection, error) {
	// 使用连接池
	conn, connection, err := r.getActiveConnection()
	if err != nil {
		return nil, err
	}
	if connection.GetConnection().IsClosed() {
		if pc, ok := conn.(*pool.PoolConn); ok {
			pc.MarkUnusable()
		}
		_, connection, err = r.getActiveConnection()
		if err != nil {
			return nil, err
		}
	}
	return connection, err
}

func (r *Rabbitmq) getActiveConnection() (net.Conn, *Connection, error) {
	conn, err := r.Pool.Get()
	defer r.Pool.Close()
	if err != nil {
		log.Fatal(err)
		return nil, nil, err
	}
	return conn, conn.(*Connection), nil
}
