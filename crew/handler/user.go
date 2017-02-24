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
	"strconv"

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

	return JSON(http.StatusCreated, "success")
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
	err = models.GetUser().Where("name = ?", user.Name).Find(&userLogin).Error
	if err != nil {
		if err.Error() == "record not found" {
			return JSON(http.StatusBadRequest, "user not exist")
		}
		log.Errorf("[handler.PostUserLoginV1Handler] get user error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if userLogin.Password != user.Password {
		return JSON(http.StatusBadRequest, "passowrd error")
	}

	return JSON(http.StatusOK, "success")
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

	userID := ctx.Params(":user")

	var existUser models.User
	err = models.GetUser().Where("id = ?", userID).First(&existUser).Error
	if err != nil {
		log.Errorf("[handler.PutUserResetV1Handler] get exist user error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if user.Name != "" && user.Name != existUser.Name {
		log.Errorf("[handler.PutUserResetV1Handler] username can not be update.\n")
		return JSON(http.StatusBadRequest, "username can not be update.")
	}

	err = models.GetUser().Where("id = ?", userID).Updates(user).Error
	if err != nil {
		log.Errorf("[handler.PutUserResetV1Handler] get exist user error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusCreated, "success")
}

func GetUserExistV1Handler(ctx *macaron.Context) (int, []byte) {
	username := ctx.Params(":username")

	var user models.User
	err := models.GetUser().Where("name = ?", username).First(&user).Error
	if err != nil {
		if err.Error() == "record not found" {
			return JSON(http.StatusOK, false)
		}
		log.Errorf("[handler.GetUserExistV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if user.ID != 0 {
		return JSON(http.StatusOK, true)
	}

	return JSON(http.StatusOK, false)
}

func GetUserListV1Handler(ctx *macaron.Context) (int, []byte) {
	pageIndex, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))

	log.Debugf("pageIndex=%v,pageSize=%v\n", pageIndex, pageSize)

	var users []models.User
	var err error
	if pageIndex > 0 {
		err = models.GetUser().Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&users).Error
	} else {
		err = models.GetUser().Limit(pageSize).Find(&users).Error
	}
	if err != nil {
		log.Errorf("[handler.GetUserListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, users)
}
