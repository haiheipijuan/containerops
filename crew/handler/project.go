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

func PostProjectV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PostProjectV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var project models.Project
	err = json.Unmarshal(reqBody, &project)
	if err != nil {
		log.Errorf("[handler.PostProjectV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	err = models.GetProject().Save(&project).Error
	if err != nil {
		log.Errorf("[handler.PostProjectV1Handler] save project error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func DeleteProjectV1Handler(ctx *macaron.Context) (int, []byte) {
	proID := ctx.Params(":project")
	err := models.GetDB().Where("id = ?", proID).Delete(models.Project{}).Error
	if err != nil {
		log.Errorf("[handler.DeleteProjectV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func PutProjectV1Handler(ctx *macaron.Context) (int, []byte) {
	reqBody, err := ctx.Req.Body().Bytes()
	if err != nil {
		log.Errorf("[handler.PutProjectV1Handler] parse request body error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	var project models.Project
	err = json.Unmarshal(reqBody, &project)
	if err != nil {
		log.Errorf("[handler.PutProjectV1Handler] json unmarshal error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	project.ID, _ = strconv.ParseInt(ctx.Params(":project"), 10, 64)

	err = models.GetProject().Save(&project).Error
	if err != nil {
		log.Errorf("[handler.PutProjectV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, "success")
}

func GetProjectV1Handler(ctx *macaron.Context) (int, []byte) {
	proID := ctx.Params(":project")

	var project models.Project
	err := models.GetProject().Where("id = ?", proID).First(&project).Error
	if err != nil {
		log.Errorf("[handler.GetProjectV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	return JSON(http.StatusOK, project)
}

func GetProjectListV1Handler(ctx *macaron.Context) (int, []byte) {
	orgID := ctx.Params(":organization")

	var projects []models.Project
	err := models.GetProject().Where("organization = ?", orgID).Find(&projects).Error
	if err != nil {
		log.Errorf("[handler.GetProjectListV1Handler] error:%v\n", err)
		return JSON(http.StatusBadRequest, err)
	}

	// var team models.Team
	// err = models.GetTeam().Where("id = ?", teamID).First(&team).Error
	// if err != nil {
	// 	log.Errorf("[handler.GetProjectListV1Handler] error:%v\n", err)
	// 	return JSON(http.StatusBadRequest, err)
	// }

	// var resultPorjects []models.Project
	// if team.Name != "Owner" {
	// 	// normal team only can get public project
	// 	for _, v := range projects {
	// 		if !v.Private {
	// 			resultPorjects = append(resultPorjects, v)
	// 		}
	// 	}
	// } else {
	// 	// owner team can get all project
	// 	resultPorjects = append(resultPorjects, projects...)
	// }

	return JSON(http.StatusOK, projects)
}
