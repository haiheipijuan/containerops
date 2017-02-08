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

	"github.com/Huawei/containerops/crew/models"
	"gopkg.in/macaron.v1"
)

func PostUserRegisterV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostUserRegisterV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var user models.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Errorf("[handler.PostUserRegisterV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetUser().Save(&user).Error
	if err != nil {
		log.Errorf("[handler.PostUserRegisterV1Handler] save user error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func PostUserLoginV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostUserLoginV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var user models.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Errorf("[handler.PostUserLoginV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var userLogin models.User
	err = models.GetUser().Where("name = ? AND password = ?", user.Name, user.Password).Find(&userLogin).Error
	if err != nil {
		log.Errorf("[handler.PostUserLoginV1Handler] get user error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if userLogin.ID != 0 {
		return JSON(http.StatusOK, "success")
	}

	return JSON(http.StatusOK, "failed")
}

func PutUserResetV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutUserResetV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var user models.User
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Errorf("[handler.PutUserResetV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetUser().Where("name = ?", user.Name).Update("password", user.Password).Error
	if err != nil {
		log.Errorf("[handler.PutUserResetV1Handler] update user password error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func GetUserExistV1Handler(ctx *macaron.Context) (int, []byte) {
	username := ctx.Params(":username")

	var user models.User
	err := models.GetUser().Where("name = ?", username).First(&user).Error
	if err != nil {
		log.Errorf("[handler.GetUserExistV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if user.ID != 0 {
		return JSON(http.StatusOK, true)
	}

	return JSON(http.StatusOK, false)
}
