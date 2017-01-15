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

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/Huawei/containerops/crew/models"
)

// databasecmd is subcommand which migrate/backup/restore crew's database.
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database subcommand migrate/backup/restore Crew's database.",
	Long:  ``,
}

// migrateDatabaseCmd is subcommand migrate crew's database.
var migrateDatabaseCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate subcommand migrate crew's database.",
	Long:  ``,
	Run:   migrateDatabase,
}

// backupDatabaseCmd is subcommand backup crew's database.
var backupDatabaseCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup subcommand backup crew's database.",
	Long:  ``,
	Run:   backupDatabase,
}

// restoreDatabaseCmd is subcommand restore crew's database.
var restoreDatabaseCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore subcommand restore crew's database.",
	Long:  ``,
	Run:   restoreDatabase,
}

// init()
func init() {
	RootCmd.AddCommand(databaseCmd)

	databaseCmd.AddCommand(migrateDatabaseCmd)
	databaseCmd.AddCommand(backupDatabaseCmd)
	databaseCmd.AddCommand(restoreDatabaseCmd)
}

// migrateDatabase is auto-migrate database of Crew.
func migrateDatabase(cmd *cobra.Command, args []string) {
	models.Migrate()
}

// backupDatabase is
func backupDatabase(cmd *cobra.Command, args []string) {

}

// restoreDatabase is
func restoreDatabase(cmd *cobra.Command, args []string) {

}
