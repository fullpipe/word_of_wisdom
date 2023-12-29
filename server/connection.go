package server

import (
	"context"
	"encoding/gob"
	"net"

	"github.com/fullpipe/word_of_wisdom/puzzle"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ConnectionState int

const (
	Initial ConnectionState = 1 << iota
	WaitingSolution
	Connected
)

type Connection struct {
	conn       net.Conn
	state      ConnectionState
	complexity int
}

func NewConnection(conn net.Conn, complexity int) *Connection {
	return &Connection{
		conn:       conn,
		state:      Initial,
		complexity: complexity,
	}
}

func (c *Connection) Run(serverCtx context.Context) error {
	defer c.conn.Close()
	_, cancel := context.WithCancel(serverCtx)
	go func() {
		<-serverCtx.Done()
		logrus.Info("shutting down client connection...")
		c.conn.Close()
	}()

	enc := gob.NewEncoder(c.conn)
	dec := gob.NewDecoder(c.conn)
	p := puzzle.NewPuzzle(c.complexity)

	for {
		switch c.state {
		case Initial:
			logrus.Info("sending puzzle to client")
			err := enc.Encode(p)
			if err != nil {
				cancel()
				return errors.Wrap(err, "unable to start puzzle solving")
			}

			c.state = WaitingSolution
		case WaitingSolution:
			logrus.Info("validating puzzle solution")
			var solution puzzle.Solution
			err := dec.Decode(&solution)
			if err != nil {
				cancel()
				return errors.Wrap(err, "unable to decode solution")
			}

			validator := puzzle.NewValidator()
			if !validator.Validate(p, solution) {
				cancel()
				return errors.New("invalid solution")
			}

			c.state = Connected
		case Connected:
			logrus.Info("client connected...sending wisdom")
			err := enc.Encode(&Message{Message: GetRandomWisdom()})
			if err != nil {
				cancel()
				return errors.Wrap(err, "unable to encode wisdom")
			}

			cancel()

			return nil
		}
	}
}

func HandleConnection(ctx context.Context, conn net.Conn) error {
	connection := &Connection{
		conn:  conn,
		state: Initial,
	}

	return connection.Run(ctx)
}
