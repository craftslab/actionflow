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
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"actionflow/config"
)

type Router struct {
	config config.Config
	router *gin.Engine
}

func Run(addr string, cfg *config.Config) error {
	r := Router{}

	if err := r.init(cfg); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	if err := r.route(); err != nil {
		return errors.Wrap(err, "failed to route")
	}

	return r.run(addr)
}

func (r *Router) init(cfg *config.Config) error {
	gin.SetMode(gin.ReleaseMode)

	r.config = *cfg
	r.router = gin.Default()

	return nil
}

func (r *Router) route() error {
	r.router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return nil
}

func (r Router) run(addr string) error {
	s := &http.Server{
		Addr:           addr,
		Handler:        r.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to listen and serve")
	}

	return nil
}
