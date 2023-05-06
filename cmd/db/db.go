/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/cheekspot/cmd"
	logging "github.com/taylormonacelli/cheekspot/cmd/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"

	// "gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
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
		run()
	},
}

var Logger *zap.Logger

func init() {
	cmd.RootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type User struct {
	gorm.Model
	Name          string
	BirthDate     time.Time
	Address       string
	SummerAddress string
}

func run() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// logger.Logger.Info("Hello, world!")

	// Auto Migrate
	db.AutoMigrate(&User{}, &User{})

	// Migrate specific table
	db.Table("users").AutoMigrate(&User{})

	// Create a new user
	user := User{Name: "John Doe", BirthDate: time.Now().AddDate(0, -1, 0), Address: "123 Main St"}
	db.Create(&user)

	// Query for users born within the last month
	var users []User
	db.Where("birth_date > ?", time.Now().AddDate(0, -2, 0)).Find(&users)

	// Print the results
	for _, u := range users {
		fmt.Println(u.Name, u.BirthDate, u.Address)
	}
}
