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

// User Table define.
type User struct {
	ID           int64      `json:"id" gorm:"primary_key"` // User ID
	Name         string     `json:"name"`                  // User name
	Nickname     string     `json:"nickname"`              // User nickname
	Password     string     `json:"password"`              // User password for login
	HeadImageUrl string     `json:"head_image_url"`        // User headimageurl
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
}

// TableName is return the table name of User in database.
func (r *User) TableName() string {
	return "user"
}

// Return *gorm.DB with table User
func GetUser() *gorm.DB {
	return db.Model(&User{})
}

func getUserByName(name string) User {
	var user User
	if err := GetUser().Where("name = ?", name).First(&user).Error; err != nil {
		log.Errorf("getUserByName error: %v\n", err)
	}
	return user
}
