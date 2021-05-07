package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"mirror-apt/tools"
	"net/http"
	"time"
)

type Instance struct {
	httpServer *http.Server
}

func NewInstance() *Instance {
	s := &Instance{
		// just in case you need some setup here
	}

	return s
}

func (s *Instance) Start() { // Startup all dependencies
	aptHandler := APTHandler{
		sourceBase: map[string]string{
			"ubuntu": "https://mirrors.aliyun.com",
			"debian": "https://mirrors.ustc.edu.cn",
			"debian-security": "https://mirrors.ustc.edu.cn",
			"pve": "http://download.proxmox.com/debian",
			"corosync-3": "http://download.proxmox.com/debian",
			"ceph-nautilus": "http://download.proxmox.wiki/debian",
			"docker-ce": "https://mirrors.ustc.edu.cn",
		},
		targetBase: "/data/apt-mirror/mirror/mirrors.ustc.edu.cn/",
	}
	addr := tools.GetEnv("ADDR", ":8888")
	// we can gracefully shut it down again
	s.httpServer = &http.Server{Addr: addr, Handler: aptHandler}
	err := s.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.WithError(err).Error("Http Server stopped unexpected")
		s.Shutdown()
	} else {
		logrus.WithError(err).Info("Http Server stopped")
	}
}

func (s *Instance) Shutdown() {
	if s.httpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := s.httpServer.Shutdown(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to shutdown http server gracefully")
		} else {
			s.httpServer = nil
		}
	}
}
