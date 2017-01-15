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

// Application Table define.
type Application struct {
	ID        int64      `json:"id" gorm:"primary_key"` // Application ID
	Name      string     `json:"name"`                  // Application name
	Private   bool       `json:"private"`               // Application is public or private
	Project   int64      `json:"project_id"`            // Which project the application belongs to
	Teams     []Team     `json:"teams"`                 // Different teams have differt permissions
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of Application in database.
func (r *Application) TableName() string {
	return "application"
}

// Return *gorm.DB with table Application
func GetApplication() *gorm.DB {
	return db.Model(&Application{})
}
