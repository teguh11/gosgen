/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("db called")
		driver, _ := cmd.Flags().GetString("driver")
		// host, _ := cmd.Flags().GetString("host")
		// port, _ := cmd.Flags().GetString("port")
		// username, _ := cmd.Flags().GetString("username")
		// password, _ := cmd.Flags().GetString("password")
		// dbname, _ := cmd.Flags().GetString("dbname")

		if driver == "mysql" {
			fmt.Println("mysql")
		} else if driver == "pgsql" {
			fmt.Println("pgsql")
		} else {
			fmt.Println("Driver not support")
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.Flags().String("driver", "", "driver database to generate (mysql, postgres)")
	dbCmd.Flags().String("host", "localhost", "host for connection (localhost)")
	dbCmd.Flags().String("port", "3306", "port for connection (3306)")
	dbCmd.Flags().String("username", "", "username for connection (empty)")
	dbCmd.Flags().String("password", "", "password for connection (empty)")
	dbCmd.Flags().String("dbname", "", "database for connection (empty)")
	dbCmd.MarkFlagRequired("driver")
	// dbCmd.MarkFlagRequired("host")
	// dbCmd.MarkFlagRequired("port")
	// dbCmd.MarkFlagRequired("username")
	// dbCmd.MarkFlagRequired("password")
	// dbCmd.MarkFlagRequired("dbname")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
