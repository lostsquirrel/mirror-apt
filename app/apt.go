package app

import (
	"context"
	"log"

	"mirror-apt/tools"
	"net/http"
	"time"
)

type Instance struct {
	httpServer *http.Server
	aptHandler APTHandler
	addr       string
}

func NewInstance() *Instance {
	config := tools.LoadEnvConfig()
	path := tools.GetEnv("TARGET_PATH", "/data/apt-mirror/mirror/mirrors.ustc.edu.cn/")
	aptHandler := APTHandler{
		sourceBase: config,
		targetBase: path,
	}
	s := &Instance{
		// just in case you need some setup here
		aptHandler: aptHandler,
		addr:       tools.GetEnv("ADDR", ":8888"),
	}

	return s
}

func (s *Instance) Start() { // Startup all dependencies

	s.httpServer = &http.Server{Addr: s.addr, Handler: s.aptHandler}
	err := s.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Println("Http Server stopped unexpected")
		s.Shutdown()
	} else {
		log.Fatal("Http Server stopped")
	}
}

func (s *Instance) Shutdown() {
	if s.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			log.Fatal("Failed to shutdown http server gracefully")
		} else {
			s.httpServer = nil
		}
	}
}
