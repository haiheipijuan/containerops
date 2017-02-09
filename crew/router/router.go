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

package router

import (
	"gopkg.in/macaron.v1"

	"github.com/Huawei/containerops/crew/handler"
)

// SetRouters is crew router's definition fucntion.
func SetRouters(m *macaron.Macaron) {
	m.Group("/v1", func() {
		m.Group("/user", func() {
			m.Post("/register", handler.PostUserRegisterV1Handler)
			m.Post("/login", handler.PostUserLoginV1Handler)
			m.Put("/reset", handler.PutUserResetV1Handler)
			m.Get("/exist/:username", handler.GetUserExistV1Handler)
		})

		m.Group("/organization", func() {
			m.Post("/", handler.PostOrganizationV1Handler)
			m.Delete("/:organization", handler.DeleteOrganizationV1Handler)
			m.Put("/:organization/:name", handler.PutOrganizationV1Handler)
			m.Get("/:organization", handler.GetOrganizationV1Handler)

			m.Get("/list", handler.GetOrganizationListV1Handler)
		})

		m.Group("/team", func() {
			m.Post("/", handler.PostTeamV1Handler)
			m.Delete("/:team", handler.DeleteTeamV1Handler)
			m.Put("/:team", handler.PutTeamV1Handler)
			m.Get("/:team", handler.GetTeamV1Handler)

			m.Get("/list/:organization", handler.GetTeamListV1Handler)

			m.Group("/role", func() {
				m.Post("/assign/:team", handler.PostTeamRoleAssignV1handler)
			})
		})

		m.Group("/project", func() {
			m.Post("/", handler.PostProjectV1Handler)
			m.Delete("/:project", handler.DeleteProjectV1Handler)
			m.Put("/:project", handler.PutProjectV1Handler)
			m.Get("/:project", handler.GetProjectV1Handler)

			m.Get("/list/:/organization", handler.GetProjectListV1Handler)
		})

		m.Group("/application", func() {
			m.Post("/", handler.PostApplicationV1Handler)
			m.Delete("/:application", handler.DeleteApplicationV1Handler)
			m.Put("/:application", handler.PutApplicationV1Handler)
			m.Get("/:application", handler.GetApplicationV1Handler)

			m.Get("/list/:project", handler.GetApplicationListV1Handler)
		})

		m.Group("/module", func() {
			m.Post("/", handler.PostModuleV1Handler)
			m.Delete("/:module", handler.DeleteModuleV1Handler)
			m.Put("/:module", handler.PutModuleV1Handler)
			m.Get("/:module", handler.GetModuleV1Handler)

			m.Get("/list/:application", handler.GetModuleListV1Handler)
		})

		m.Group("/workflow", func() {
			m.Post("/", handler.PostWrokflowV1Handler)
			m.Delete("/:workflow", handler.DeleteWrokflowV1Handler)
			m.Put("/:workflow", handler.PutWrokflowV1Handler)
			m.Get("/:workflow", handler.GetWrokflowV1Handler)

			m.Get("/list/:organization", handler.GetWrokflowListV1Handler)
		})

		m.Group("/component", func() {
			m.Post("/", handler.PostComponentV1Handler)
			m.Delete("/:component", handler.DeleteComponentV1Handler)
			m.Put("/:component", handler.PutComponentV1Handler)
			m.Get("/:component", handler.GetComponentV1Handler)

			m.Get("/list/:organization", handler.GetComponentListV1Handler)
		})

		// m.Group("/permission", func() {
		// 	m.Get("/list/:team", handler.GetTeamPermissionListV1Handler)
		// })

		m.Group("/role", func() {
			m.Get("/list", handler.GetRoleListV1Handler)
		})
	})
}
