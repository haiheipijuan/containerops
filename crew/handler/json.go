/*
Copyright 2014 Huawei Technologies Co., Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
}

func JSON(code int, msg interface{}) (int, []byte) {
	var result []byte
	switch code {
	case http.StatusOK:
		result, _ = json.Marshal(msg)
	case http.StatusCreated, http.StatusNoContent, http.StatusAccepted:
		// POST/PUT/DELETE success, return nothing
		result, _ = json.Marshal(map[string]interface{}{"result": "success"})
	default:
		result, _ = json.Marshal(map[string]interface{}{"error": msg})
	}
	return code, result
}
