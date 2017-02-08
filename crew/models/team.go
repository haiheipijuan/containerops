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
	ID           int64      `json:"id" gorm:"primary_key"` // Team ID
	Name         string     `json:"name"`                  // Team name
	Organization int64      `json:"org_id"`                // Which organization the team belongs to
	Users        []User     `json:"users"`                 // Users the team has
	Role         Role       `json:"role"`                  // Role the team is
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of Team in database.
func (r *Team) TableName() string {
	return "team"
}

// Return *gorm.DB with table Team
func GetTeam() *gorm.DB {
	return db.Model(&Team{})
}
