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
	"net/http"

	"gopkg.in/macaron.v1"
)

func PostApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func DeleteApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func PutApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func GetApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func GetApplicationListV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}
