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

func PostModuleV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostModuleV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var module models.Module
	err = json.Unmarshal(reqBody, &module)
	if err != nil {
		log.Errorf("[handler.PostModuleV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetModule().Save(&module).Error
	if err != nil {
		log.Errorf("[handler.PostModuleV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func DeleteModuleV1Handler(ctx *macaron.Context) (int, []byte) {
	modueID := ctx.Params(":module")
	err := models.GetDB().Where("id = ?", modueID).Delete(models.Module{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteModuleV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutModuleV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutModuleV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var module models.Module
	err = json.Unmarshal(reqBody, &module)
	if err != nil {
		log.Errorf("[handler.PutModuleV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	module.ID, _ = strconv.ParseInt(ctx.Params(":module"), 10, 64)

	err = models.GetModule().Save(&module).Error
	if err != nil {
		log.Errorf("[handler.PutModuleV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetModuleV1Handler(ctx *macaron.Context) (int, []byte) {
	modueID := ctx.Params(":module")

	var module models.Module
	err := models.GetModule().Where("id = ?", modueID).First(&module).Error
	if err != nil {
		log.Errorf("[handler.GetModuleV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, module)
}

func GetModuleListV1Handler(ctx *macaron.Context) (int, []byte) {
	appID := ctx.Params(":application")

	var modules []models.Module
	err := models.GetModule().Where("application = ?", appID).Find(&modules).Error
	if err != nil {
		log.Errorf("[handler.GetApplicationListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, modules)
}
