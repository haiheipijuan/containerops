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

func PostOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostOrganizationV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var org models.Organization
	err = json.Unmarshal(reqBody, &org)
	if err != nil {
		log.Errorf("[handler.PostOrganizationV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetOrganization().Save(&org).Error
	if err != nil {
		log.Errorf("[handler.PostOrganizationV1Handler] save orgnazation error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	// create default owner team
	var users []models.User
	users = append(users, models.GetUserByName(org.Owner))

	var team models.Team
	team.Name = "Owner"
	team.Organization = org.ID
	team.Role = models.GetRoleByName("Owner")
	team.Users = users

	err = models.GetTeam().Save(&team).Error
	if err != nil {
		log.Errorf("[handler.PostOrganizationV1Handler] save team error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, "success")
}

func DeleteOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")
	err := models.GetDB().Where("id = ?", orgID).Delete(models.Organization{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteOrganizationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")
	orgName := ctx.Params(":name")
	err := models.GetOrganization().Where("id = ?", orgID).Update("name", orgName).Error
	if err != nil {
		log.Errorf("[handler.PutOrganizationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetOrganizationV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")

	var org models.Organization
	err := models.GetOrganization().Where("id = ?", orgID).First(&org).Error
	if err != nil {
		log.Errorf("[handler.GetOrganizationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, org)
}

func GetOrganizationListV1Handler(ctx *macaron.Context) (int, []byte) {
	var orgs []models.Organization
	err := models.GetOrganization().Find(&orgs).Error
	if err != nil {
		log.Errorf("[handler.GetOrganizationV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, orgs)
}
