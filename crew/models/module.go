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

// Module Table define.
type Module struct {
	ID               int64      `json:"id" gorm:"primary_key"` // Module ID
	Name             string     `json:"name"`                  // Module name
	Private          bool       `json:"private"`               // Module is public or private
	Application      int64      `json:"application_id"`        // Which application the module belongs to
	Teams            []Team     `json:"teams"`                 // Different teams have differt permissions
	GitResponsityUrl string     `json:"git_responsity_url"`    // One Module correspond one git responsity url
	CreatedAt        time.Time  `json:"-"`
	UpdatedAt        time.Time  `json:"-"`
	DeletedAt        *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of Module in database.
func (r *Module) TableName() string {
	return "module"
}

// Return *gorm.DB with table Module
func GetModule() *gorm.DB {
	return db.Model(&Module{})
}