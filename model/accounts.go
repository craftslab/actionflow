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

package model

import (
	"errors"
)

type Account struct {
	DisplayName string `json:"displayname"`
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	UserName    string `json:"username"`
}

var accounts = []Account{
	{
		DisplayName: "Administrator",
		Email:       "",
		Id:          1,
		Name:        "Administrator",
		UserName:    "admin",
	},
	{
		DisplayName: "Super John",
		Email:       "john.doe@example.com",
		Id:          2,
		Name:        "John Doe",
		UserName:    "john",
	},
}

func GetAccount(id int) (Account, error) {
	for _, v := range accounts {
		if id == v.Id {
			return v, nil
		}
	}

	return Account{}, errors.New("invalid id")
}

func QueryAccount(q string) ([]Account, error) {
	if q == "" {
		return []Account{}, errors.New("invalid query")
	}

	var buf []Account

	for k, v := range accounts {
		if q == v.UserName {
			buf = append(buf, accounts[k])
		}
	}

	return buf, nil
}
