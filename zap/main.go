package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()


	url := "http://marmotedu.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

}
