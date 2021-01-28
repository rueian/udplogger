package main

import (
	"go.uber.org/zap"
	"net"
	"udplogger/pkg/msg"
	"udplogger/pkg/util"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn, err := net.Dial("udp", "127.0.0.1:7777")
	if err != nil {
		logger.Fatal("dial failed", zap.Error(err))
	}

	b := make([]byte, 1472)
	err = util.RateLogging(logger, func() error {
		m := msg.Message{}
		m.ToDatagram(b)
		conn.Write(b)
		return nil
	})
	logger.Info("closing", zap.Error(err))
}
