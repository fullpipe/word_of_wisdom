package client

import (
	"net"

	"github.com/fullpipe/word_of_wisdom/server"
	"github.com/urfave/cli/v2"
)

func NewClient() *cli.Command {
	return &cli.Command{
		Name:   "client",
		Action: clientAction,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost:8080",
			},
		},
	}
}

func clientAction(cCtx *cli.Context) error {
	conn, err := net.Dial("tcp", cCtx.String("host"))
	if err != nil {
		return err
	}
	defer conn.Close()

	return server.StartClient(cCtx.Context, conn)
}
