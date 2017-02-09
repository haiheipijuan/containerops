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

func PostWrokflowV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostWrokflowV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var workflow models.Workflow
	err = json.Unmarshal(reqBody, &workflow)
	if err != nil {
		log.Errorf("[handler.PostWrokflowV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetWorkflow().Save(&workflow).Error
	if err != nil {
		log.Errorf("[handler.PostWrokflowV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func DeleteWrokflowV1Handler(ctx *macaron.Context) (int, []byte) {
	workflowID := ctx.Params(":workflow")
	err := models.GetDB().Where("id = ?", workflowID).Delete(models.Workflow{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteWrokflowV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutWrokflowV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutWrokflowV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var workflow models.Workflow
	err = json.Unmarshal(reqBody, &workflow)
	if err != nil {
		log.Errorf("[handler.PutWrokflowV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	workflow.ID, _ = strconv.ParseInt(ctx.Params(":workflow"), 10, 64)

	err = models.GetComponent().Save(&workflow).Error
	if err != nil {
		log.Errorf("[handler.PutWrokflowV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetWrokflowV1Handler(ctx *macaron.Context) (int, []byte) {
	workflowID := ctx.Params(":workflow")

	var workflow models.Workflow
	err := models.GetWorkflow().Where("id = ?", workflowID).First(&workflow).Error
	if err != nil {
		log.Errorf("[handler.GetWrokflowV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, workflow)
}

func GetWrokflowListV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")

	var workflows []models.Workflow
	err := models.GetWorkflow().Where("organization = ?", orgID).Find(&workflows).Error
	if err != nil {
		log.Errorf("[handler.GetWrokflowListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, workflows)
}
