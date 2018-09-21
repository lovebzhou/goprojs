package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/lovebzhou/goprojs/all/b2s"
	"github.com/lovebzhou/goprojs/all/c2s"

	_ "net/http/pprof"

	"github.com/lovebzhou/goprojs/all/api"
)

func init() {
	log.Println("init in main")
}

func main() {
	// load configuration
	configFile := flag.String("config", "./all.yml", "Configuration file path.")

	var cfg Config
	if err := cfg.FromFile(*configFile); err != nil {
		fmt.Fprintf(os.Stderr, "jackal: %v\n", err)
		return
	}

	if cfg.Debug.Port > 0 {
		go initDebugServer(cfg.Debug.Port)
	}

	log.Printf("%v", cfg)

	go c2s.StartService(cfg.C2S)
	go api.StartService(cfg.API)
	b2s.StartService(cfg.B2S)
}

func initDebugServer(port int) {
	debugSrv := &http.Server{}
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("%v", err)
	}
	debugSrv.Serve(ln)
}
