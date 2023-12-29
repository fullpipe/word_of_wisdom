package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/fullpipe/word_of_wisdom/server"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewServer() *cli.Command {
	return &cli.Command{
		Name:   "server",
		Action: serverAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost:8080",
			},
			&cli.IntFlag{
				Name:  "complexity",
				Value: 2,
				Action: func(ctx *cli.Context, v int) error {
					if v < 1 || v > 10 {
						return fmt.Errorf("complexity value %v out of range[1-10]", v)
					}

					return nil
				},
			},
		},
	}
}

func serverAction(cCtx *cli.Context) error {
	listener, err := net.Listen("tcp", cCtx.String("host"))
	if err != nil {
		return err
	}
	defer listener.Close()

	logrus.Infof("Server is listening %s", cCtx.String("host"))

	go func() {
		<-cCtx.Context.Done()
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if errors.Is(err, net.ErrClosed) {
			return nil
		}

		if err != nil {
			logrus.Error(err)
			continue
		}

		go func() {
			c := server.NewConnection(conn, cCtx.Int("complexity"))
			err := c.Run(cCtx.Context)
			if err != nil {
				logrus.Error(err)
			}
		}()
	}
}
