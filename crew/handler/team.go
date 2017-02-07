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
	"time"

	"github.com/Huawei/containerops/models"
	"gopkg.in/macaron.v1"
)

type Team struct {
	ID           int64      `json:"id" gorm:"primary_key"` // Team ID
	Name         string     `json:"name"`                  // Team name
	Organization int64      `json:"organization_id"`       // Which organization the team belongs to
	Users        []User     `json:"users"`                 // Users the team has
	Role         Role       `json:"role"`                  // Role the team is
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
}

func PostTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] parse request body error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	var team models.Team
	err = json.Unmarshal(reqBody, &team)
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] json unmarshal error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PostTeamV1Handler] save team error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	return Result(http.StatusOK, "success")
}

func DeleteTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	teamID := ctx.Params(":team")
	err := models.GetDB().Where("id = ?", teamID).Delete(models.Team{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteTeamV1Handler] error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}
	return Result(http.StatusOK, "success")
}

func PutTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] parse request body error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	var team models.Team
	err = json.Unmarshal(reqBody, &team)
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] json unmarshal error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	team.ID = ctx.Params(":team")

	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] save team error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}
	return Result(http.StatusOK, "success")
}

func GetTeamV1Handler(ctx *macaron.Context) (int, []byte) {
	teamID := ctx.Params(":team")

	var team modesl.Team
	err := modesl.GetTeam().Where("id = ?", teamID).First(&team).Error
	if err != nil {
		log.Errorf("[handler.PutTeamV1Handler] save team error:%v\n", err)
		return Result(http.StatusBadRequest, err)
	}

	return Result(http.StatusOK, "")
}

func GetTeamListV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func GetTeamUserListV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func PostTeamRoleAssignV1handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func GetTeamRoleV1handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}

func GetTeamPermissionListV1Handler(ctx *macaron.Context) (int, []byte) {
	return Result(http.StatusOK, "")
}
