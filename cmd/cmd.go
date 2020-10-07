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

package cmd

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/alecthomas/kingpin.v2"

	"actionflow/config"
)

var (
	app  = kingpin.New("actionflow", "Action Flow").Author(Author).Version(Version)
	addr = app.Flag("addr", "Server listen address").Default(":9090").String()
)

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	runServer(&cfg)

	log.Println("server finished.")
}

func initConfig() (config.Config, error) {
	cfg := config.NewConfig()

	return cfg, nil
}

func runServer(cfg *config.Config) {
	s := &http.Server{
		Addr:           *addr,
		Handler:        setupRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
