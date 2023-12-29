package server

import (
	"context"
	"encoding/gob"
	"net"

	"github.com/fullpipe/word_of_wisdom/puzzle"
	"github.com/sirupsen/logrus"
)

func StartClient(ctx context.Context, conn net.Conn) error {
	go func() {
		<-ctx.Done()
		logrus.Info("shutting down connection...")
		conn.Close()
	}()

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	var p puzzle.Puzzle
	err := dec.Decode(&p)
	if err != nil {
		logrus.Fatal("unable to read puzzle:", err)
	}

	logrus.Infof("solving puzzle with Complexity: %d", p.Complexity)

	solver := puzzle.NewSolver()
	solution := solver.Solve(p)

	logrus.Infof("solved with %d attempts", solution.Attempts)

	err = enc.Encode(solution)
	if err != nil {
		logrus.Fatal("unable to send solution:", err)
	}

	var m Message
	err = dec.Decode(&m)
	if err != nil {
		logrus.Fatal("decode error:", err)
	}

	logrus.Info("wisdom: ", m.Message)

	return nil
}
