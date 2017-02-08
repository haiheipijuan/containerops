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

func PostTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var team models.Team
	err = json.Unmarshal(reqBody, &team)
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	if team.Name == "Owner" {
		return JSON(http.StatusBadRequest, "Team name can not be Owner!")
	}

	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] save team error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func DeleteTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	teamID := ctx.Params(":team")
	err := models.GetDB().Where("id = ?", teamID).Delete(models.Team{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteTeamV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var team models.Team
	err = json.Unmarshal(reqBody, &team)
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	team.ID, _ = strconv.ParseInt(ctx.Params(":team"), 10, 64)

	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] save team error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	teamID := ctx.Params(":team")

	var team models.Team
	err := models.GetTeam().Where("id = ?", teamID).First(&team).Error
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, team)
}

func GetTeamListV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")

	var teams []models.Team
	err := models.GetTeam().Where("organization = ?", orgID).Find(&teams).Error
	if err != nil {
		log.Errorf("[handler.GetTeamListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, teams)
}

func PostTeamRoleAssignV1handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostTeamRoleAssignV1handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var role models.Role
	err = json.Unmarshal(reqBody, &role)
	if err != nil {
		log.Errorf("[handler.PostTeamRoleAssignV1handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	teamID := ctx.Params(":team")

	var team models.Team
	err = models.GetTeam().Where("id = ?", teamID).First(&team).Error
	if err != nil {
		log.Errorf("[handler.PostTeamRoleAssignV1handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	team.Role = role
	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PostTeamRoleAssignV1handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}
