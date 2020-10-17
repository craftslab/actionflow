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
	"errors"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"actionflow/config"
	"actionflow/docs"
	"actionflow/router"
)

var (
	app  = kingpin.New("actionflow", "Action Flow").Version(config.Version + "-build-" + config.Build)
	addr = app.Flag("addr", "Server listen address").Default(":9090").String()
)

func Run() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	if err := initDoc(); err != nil {
		log.Fatalf("failed to init doc: %v", err)
	}

	if err := runServer(cfg); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

	log.Println("server exiting")
}

func initConfig() (*config.Config, error) {
	c := config.New()
	if c == nil {
		return nil, errors.New("faile to new config")
	}

	return c, nil
}

func initDoc() error {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = *addr
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	return nil
}

func runServer(cfg *config.Config) error {
	r := router.New()
	if r == nil {
		return errors.New("faile to new router")
	}

	return r.Run(*addr, cfg)
}
