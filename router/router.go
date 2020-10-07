// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"actionflow/config"
	"actionflow/controller"
)

const (
	timeout = 5 * time.Second
)

type Router struct {
	config config.Config
	engine *gin.Engine
}

func Run(addr string, cfg *config.Config) error {
	r := Router{}

	if err := r.initRouter(cfg); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	if err := r.setRoute(); err != nil {
		return errors.Wrap(err, "failed to route")
	}

	return r.runRouter(addr)
}

func (r *Router) initRouter(cfg *config.Config) error {
	gin.SetMode(gin.ReleaseMode)

	r.config = *cfg

	r.engine = gin.New()
	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())

	return nil
}

func (r *Router) setRoute() error {
	ctrl := controller.NewController()

	accounts := r.engine.Group("/accounts")
	accounts.GET(":id", ctrl.GetAccount)

	cfg := r.engine.Group("/config")
	cfg.GET("server/version", ctrl.GetServerVersion)

	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return nil
}

func (r Router) runRouter(addr string) error {
	srv := &http.Server{
		Addr:           addr,
		Handler:        r.engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to listen and serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "failed to shutdown")
	}

	<-ctx.Done()

	return nil
}
