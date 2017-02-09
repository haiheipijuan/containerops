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

func PostComponentV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostComponentV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var component models.Component
	err = json.Unmarshal(reqBody, &component)
	if err != nil {
		log.Errorf("[handler.PostComponentV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetComponent().Save(&component).Error
	if err != nil {
		log.Errorf("[handler.PostComponentV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func DeleteComponentV1Handler(ctx *macaron.Context) (int, []byte) {
	componentID := ctx.Params(":component")
	err := models.GetDB().Where("id = ?", componentID).Delete(models.Component{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteComponentV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutComponentV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutComponentV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var component models.Component
	err = json.Unmarshal(reqBody, &component)
	if err != nil {
		log.Errorf("[handler.PutComponentV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	component.ID, _ = strconv.ParseInt(ctx.Params(":component"), 10, 64)

	err = models.GetComponent().Save(&component).Error
	if err != nil {
		log.Errorf("[handler.PutComponentV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetComponentV1Handler(ctx *macaron.Context) (int, []byte) {
	componentID := ctx.Params(":component")

	var component models.Component
	err := models.GetComponent().Where("id = ?", componentID).First(&component).Error
	if err != nil {
		log.Errorf("[handler.GetComponentV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, component)
}

func GetComponentListV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":org_id")

	var components []models.Component
	err := models.GetComponent().Where("organization = ?", orgID).Find(&components).Error
	if err != nil {
		log.Errorf("[handler.GetComponentListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, components)
}
