package restfulAPI

import (
	"log"
	"net/http"
	"strconv"
	"utils"
)

const (
	Version                  = "1.0"
	DefaultListenPort uint16 = 80
)

type APIServer struct {
	Version string
	Port    uint16
}

func NewApiServer(cfg *utils.Config) *APIServer {
	s := &APIServer{
		Version: Version,
		Port:    DefaultListenPort,
	}

	if cfg != nil {
		port := cfg.GetInt("api-port")
		if port > 0 {
			s.Port = uint16(port)
		}
	}

	return s
}

func (s *APIServer) Start() error {
	log.Println("Start monitor warn webhook server ...")
	router := NewRouter()
	addr := ":" + strconv.Itoa(int(s.Port))
	log.Printf("Webhook Server listen on : %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalln("Start webhook server failed by :", err.Error())
	}

	return err
}

func (s *APIServer) Shutdown() {
	log.Println("Shutdown webhook server ...\n")
}
