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

func PostApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostApplicationV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var app models.Application
	err = json.Unmarshal(reqBody, &app)
	if err != nil {
		log.Errorf("[handler.PostApplicationV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetApplication().Save(&app).Error
	if err != nil {
		log.Errorf("[handler.PostApplicationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func DeleteApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	appID := ctx.Params(":application")
	err := models.GetDB().Where("id = ?", appID).Delete(models.Application{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteApplicationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutApplicationV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var app models.Application
	err = json.Unmarshal(reqBody, &app)
	if err != nil {
		log.Errorf("[handler.PutApplicationV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	app.ID, _ = strconv.ParseInt(ctx.Params(":application"), 10, 64)

	err = models.GetProject().Save(&app).Error
	if err != nil {
		log.Errorf("[handler.PutApplicationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetApplicationV1Handler(ctx *macaron.Context) (int, []byte) {
	appID := ctx.Params(":application")

	var app models.Application
	err := models.GetApplication().Where("id = ?", appID).First(&app).Error
	if err != nil {
		log.Errorf("[handler.GetApplicationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, app)
}

func GetApplicationListV1Handler(ctx *macaron.Context) (int, []byte) {
	proID := ctx.Params(":project")

	var applications []models.Application
	err := models.GetApplication().Where("project = ?", proID).Find(&applications).Error
	if err != nil {
		log.Errorf("[handler.GetApplicationListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, applications)
}
