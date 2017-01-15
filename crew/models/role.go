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

// Role Table define.
type Role struct {
	ID          int64        `json:"id" gorm:"primary_key"` // Role ID
	Name        string       `json:"name"`                  // Role Name
	Permissions []Permission `json:"permissions"`           // Permissions the role has
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
	DeletedAt   *time.Time   `json:"-" sql:"index"`
}

// TableName is return the table name of Role in database.
func (r *Role) TableName() string {
	return "role"
}

// Return *gorm.DB with table Role
func GetRole() *gorm.DB {
	return db.Model(&Role{})
}

// Generate default roles, and insert them to db
func defaultRoles() {
	var count int
	if err := GetRole().Count(&count).Error; err != nil {
		log.Errorf("defaultPermissions error: %v\n", err)
		return
	}
	log.Debugf("defaultRoles count=%v\n", count)

	if count <= 0 {
		GetRole().Save(&Role{Name: "ReadRole", Permissions: []Permission{getPermissionByName("Read")}})
		GetRole().Save(&Role{Name: "RWRole", Permissions: []Permission{getPermissionByName("Read"), getPermissionByName("Write")}})
		GetRole().Save(&Role{Name: "Admin", Permissions: []Permission{getPermissionByName("Read"), getPermissionByName("Write"), getPermissionByName("Delete")}})
		GetRole().Save(&Role{Name: "Owner", Permissions: []Permission{getPermissionByName("Read"), getPermissionByName("Write"), getPermissionByName("Delete")}})
	}
}

func getRoleByName(name string) Role {
	var role Role
	if err := GetRole().Where("name = ?", name).First(&role).Error; err != nil {
		log.Errorf("getRoleByName error: %v\n", err)
	}
	return role
}
