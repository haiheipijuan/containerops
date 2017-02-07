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

// Organization Table define.
type Organization struct {
	ID    int64  `json:"id" gorm:"primary_key"` // Organization ID
	Name  string `json:"name"`                  // Organization name
	Owner string `json:"owner"`                 // Who create the organization
	// OwnerTeam Team       `json:"owner_team"`            // User who in the owner team can operate the organization
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of Organization in database.
func (r *Organization) TableName() string {
	return "organization"
}

// Return *gorm.DB with table Organization
func GetOrganization() *gorm.DB {
	return db.Model(&Organization{})
}
