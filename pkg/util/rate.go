package util

import (
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

func RateLogging(logger *zap.Logger, fn func() error) (err error) {
	c1 := int64(0)
	go func() {
		c2 := int64(0)
		for {
			time.Sleep(time.Second)
			c3 := atomic.LoadInt64(&c1)
			logger.Info("checkpoint", zap.Int64("rate", c3-c2))
			c2 = c3
		}
	}()
	for {
		atomic.AddInt64(&c1, 1)
		if err = fn(); err != nil {
			return err
		}
	}
}
