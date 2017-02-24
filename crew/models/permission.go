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

// import (
// 	"time"

// 	"github.com/jinzhu/gorm"
// )

// // Permission Table define.
// type Permission struct {
// 	ID        int64      `json:"id" gorm:"primary_key"` // Permission ID
// 	Name      string     `json:"name"`                  // Permission name
// 	CreatedAt time.Time  `json:"-"`
// 	UpdatedAt time.Time  `json:"-"`
// 	DeletedAt *time.Time `json:"-" sql:"index"`
// }

// // TableName is return the table name of Permission in database.
// func (r *Permission) TableName() string {
// 	return "permission"
// }

// // Return *gorm.DB with table Permission
// func GetPermission() *gorm.DB {
// 	return db.Model(&Permission{})
// }

// // Generate default peermissions, and insert them to db
// func defaultPermissions() {
// 	var count int
// 	if err := GetPermission().Count(&count).Error; err != nil {
// 		log.Errorf("defaultPermissions error: %v\n", err)
// 		return
// 	}
// 	log.Debugf("defaultPermissions count=%v\n", count)

// 	if count <= 0 {
// 		GetPermission().Save(&Permission{Name: "Read"})
// 		GetPermission().Save(&Permission{Name: "Write"})
// 		GetPermission().Save(&Permission{Name: "Delete"})
// 	}
// }

// func getPermissionByName(name string) Permission {
// 	var permission Permission
// 	if err := GetPermission().Where("name = ?", name).First(&permission).Error; err != nil {
// 		log.Errorf("getPermissionByName error: %v\n", err)
// 	}
// 	return permission
// }
