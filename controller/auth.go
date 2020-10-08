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

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Accounts = gin.Accounts{
		"admin": "admin",
	}
)

// GetLogin godoc
// @Summary Login authorization
// @Description Login authorization
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Basic authorization"
// @Success 200 {string} username
// @Failure 400 {object} util.HTTPError
// @Failure 404 {object} util.HTTPError
// @Failure 500 {object} util.HTTPError
// @Router /auth/login [get]
func (c *Controller) GetLogin(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.MustGet(gin.AuthUserKey).(string))
}
