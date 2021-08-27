package httpx

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	engine *gin.Engine
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}

func (s *Server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}

func (s *Server) GraceRun(addr ...string) {
	if len(addr) == 0 {
		addr = append(addr, ":8080")
	}
	servers := make([]*http.Server, 0)
	for _, port := range addr {
		if port == "" {
			continue
		}
		v := &http.Server{
			Addr:    port,
			Handler: s.engine,
		}
		go func(v *http.Server, port string) {
			log.Printf("Listening and serving HTTP on %s\n", port)
			if err := v.ListenAndServe(); err != nil {
				log.Fatalf("Listening run fatal: %v\n", err.Error())
			}
		}(v, port)
		servers = append(servers, v)
	}

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, v := range servers {
		if err := v.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
	}
}

func NewServer(mode string) *Server {
	engine := gin.New()

	gin.SetMode(mode)
	gin.DisableBindValidation()

	return &Server{
		engine: engine,
	}
}
