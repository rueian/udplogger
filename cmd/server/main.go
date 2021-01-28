package main

import (
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"syscall"
	"udplogger/pkg/msg"
	"udplogger/pkg/util"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	ln, err := net.ListenPacket("udp", ":7777")
	if err != nil {
		logger.Fatal("listen failed", zap.Error(err))
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdown
		ln.Close()
	}()

	buf := make([]byte, 1472)
	m := msg.Message{}
	err = util.RateLogging(logger, func() error {
		n, _, err := ln.ReadFrom(buf)
		if err != nil {
			return err
		}
		m.FromDatagram(buf[:n])
		return nil
	})
	logger.Info("closing", zap.Error(err))
}
