package main

import (
	"time"
	"github.com/breathman/graylog-example/common"
	"log"
)

func main() {
	logger, err := common.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for range time.Tick(5 * time.Second) {
		ctxLog := logger.NewPrefix("ticker")
		i++
		if i % 2 == 0 {
			ctxLog.Error("planned error")
			continue
		}
		ctxLog.Infof("message number %d", i)
	}
}