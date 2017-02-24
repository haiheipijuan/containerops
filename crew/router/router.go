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
			m.Post("/", handler.PostUserRegisterV1Handler)
			m.Get("/", handler.GetUserListV1Handler)

			m.Put("/:user", handler.PutUserResetV1Handler)
			m.Get("/:username", handler.GetUserExistV1Handler)
			m.Post("/login", handler.PostUserLoginV1Handler)
		})

		m.Group("/organization", func() {
			m.Post("/", handler.PostOrganizationV1Handler)
			m.Get("/", handler.GetOrganizationListV1Handler)

			m.Group("/:organization", func() {
				m.Delete("/", handler.DeleteOrganizationV1Handler)
				m.Put("/", handler.PutOrganizationV1Handler)
				m.Get("/", handler.GetOrganizationV1Handler)
			})
		})

		m.Group("/team", func() {
			m.Post("/", handler.PostTeamV1Handler)

			m.Group("/:team", func() {
				m.Delete("/", handler.DeleteTeamV1Handler)
				m.Put("/", handler.PutTeamV1Handler)
				m.Get("/", handler.GetTeamV1Handler)
			})

			m.Get("/:organization", handler.GetTeamListV1Handler)

			m.Put("/roleassign", handler.PutTeamRoleAssignV1handler)
		})

		m.Group("/project", func() {
			m.Post("/", handler.PostProjectV1Handler)

			m.Group("/:project", func() {
				m.Delete("/", handler.DeleteProjectV1Handler)
				m.Put("/", handler.PutProjectV1Handler)
				m.Get("/", handler.GetProjectV1Handler)

				m.Put("/teamassign/:team", handler.PutProjectTeamAssignV1handler)
			})

			m.Get("/:organization", handler.GetProjectListV1Handler)
		})

		m.Group("/application", func() {
			m.Post("/", handler.PostApplicationV1Handler)

			m.Group("/:application", func() {
				m.Delete("/", handler.DeleteApplicationV1Handler)
				m.Put("/", handler.PutApplicationV1Handler)
				m.Get("/", handler.GetApplicationV1Handler)

				m.Put("/teamassign/:team", handler.PutApplicationTeamAssignV1handler)
			})

			m.Get("/:project", handler.GetApplicationListV1Handler)
		})

		m.Group("/module", func() {
			m.Post("/", handler.PostModuleV1Handler)

			m.Group("/:module", func() {
				m.Delete("/", handler.DeleteModuleV1Handler)
				m.Put("/", handler.PutModuleV1Handler)
				m.Get("/", handler.GetModuleV1Handler)

				m.Put("/teamassign/:team", handler.PutModuleTeamAssignV1handler)
			})

			m.Get("/:application", handler.GetModuleListV1Handler)
		})

		m.Group("/workflow", func() {
			m.Post("/", handler.PostWrokflowV1Handler)

			m.Group("/:workflow", func() {
				m.Delete("/", handler.DeleteWrokflowV1Handler)
				m.Put("/", handler.PutWrokflowV1Handler)
				m.Get("/", handler.GetWrokflowV1Handler)
			})

			m.Get("/:organization", handler.GetWrokflowListV1Handler)
		})

		m.Group("/component", func() {
			m.Post("/", handler.PostComponentV1Handler)

			m.Group("/:component", func() {
				m.Delete("/", handler.DeleteComponentV1Handler)
				m.Put("/", handler.PutComponentV1Handler)
				m.Get("/", handler.GetComponentV1Handler)
			})

			m.Get("/:organization", handler.GetComponentListV1Handler)
		})

		// m.Group("/permission", func() {
		// 	m.Get("/list/:team", handler.GetTeamPermissionListV1Handler)
		// })

		m.Get("/role", handler.GetRoleListV1Handler)
	})
}
