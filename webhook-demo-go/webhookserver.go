package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"restfulAPI"
	"runtime"
	"syscall"
	"utils"
)

var (
	ErrorBadConfig = errors.New("BadConfigFile")
	config         = flag.String("c", "etc/conf.ini", "Config file")
)

type IServer interface {
	Start() error
	Shutdown()
}

func stopSignal(s IServer) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh,
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL)

	go func() {
		<-signalCh
		s.Shutdown()
		os.Exit(0)
	}()
}

func main() {
	log.Println("Welcome to ucloud monitor webhook demo ...")

	flag.Parse()
	cfg, err := utils.NewConfig(*config)
	if err != nil {
		log.Fatal(ErrorBadConfig.Error())
	}

	utils.SetGlobalConf(cfg)

	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	log.Printf("Monitor webhook demo use %d process cores\n", cpuNum)

	s := restfulAPI.NewApiServer(cfg)
	stopSignal(s)

	s.Start()
}
