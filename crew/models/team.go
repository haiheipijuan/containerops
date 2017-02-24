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

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Team Table define.
type Team struct {
	ID             int64      `json:"id" gorm:"primary_key"` // Team ID
	Name           string     `json:"name"`                  // Team name
	Users          string     `json:"users"`                 // Users the team has
	OrganizationID int64      `json:"org_id"`                // Which organization the team belongs to
	RoleID         int64      `json:"role_id"`               // Role the team is
	PorjectID      int64      `json:"project_id"`            // Which project the team be assigned to
	ApplicationID  int64      `json:"application_id"`        // Which application the team be assigned to
	ModuleID       int64      `json:"module_id"`             // Which module the team be assigned to
	WorkflowID     int64      `json:"workflow_id"`           // Which workflow the team be assigned to
	ComponentID    int64      `json:"component_id"`          // Which component the team be assigned to
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of Team in database.
func (r *Team) TableName() string {
	return "team"
}

// Return *gorm.DB with table Team
func GetTeam() *gorm.DB {
	return db.Model(&Team{})
}
