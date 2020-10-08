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
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"actionflow/config"
)

func TestRunRouter(t *testing.T) {
	auth := func(user, pass string) string {
		base := user + ":" + pass
		return "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
	}

	c := config.Config{}
	r := Router{}

	err := r.initRouter(&c)
	assert.Equal(t, nil, err)

	err = r.setRoute()
	assert.Equal(t, nil, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/auth/login", nil)
	req.Header.Set("Authorization", auth("admin", "admin"))
	r.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "admin", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/config/server/version", nil)
	req.Header.Set("Authorization", auth("admin", "admin"))
	r.engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "\""+config.Version+"-build-"+config.Build+"\"", w.Body.String())
}
