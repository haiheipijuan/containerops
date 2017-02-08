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
	"net/http"

	"github.com/Huawei/containerops/crew/models"
	"gopkg.in/macaron.v1"
)

func GetRoleListV1Handler(ctx *macaron.Context) (int, []byte) {
	var roles []models.Role
	err := models.GetRole().Find(&roles).Error
	if err != nil {
		log.Errorf("[role.GetRoleListV1Handler] error : %v\n", err)
		return JSON(http.StatusBadRequest, err)
	}
	return JSON(http.StatusOK, roles)
}
