package main

import (
	"github.com/peywong-galaxy/s-gw/httpmod"
	"github.com/peywong-galaxy/s-gw/log"
	"go.uber.org/zap"
	"sync"
)

func init() {
	log.Logger.Info("logger structure assigned successfully", zap.String("func","NewLogger"))

}

func main() {
	log.Logger.Info("Main func start")

	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		defer wg.Done()
		httpmod.LaunchHttp()
	}()

	wg.Wait()
}
