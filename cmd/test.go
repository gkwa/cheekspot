/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"

	mysns "github.com/taylormonacelli/cheekspot/cmd/aws/sns"
	"github.com/taylormonacelli/cheekspot/cmd/logging"

	"gorm.io/gorm"

	// "gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

// testCmd represents the test1 command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		test()
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test() error {
	jsonStr := `{
        "accountId": "123456789012",
        "architecture": "arm64",
        "availabilityZone": "us-east-1a",
        "billingProducts": [
            "product1",
            "product2"
        ],
        "devpayProductCodes": [
            "devpayCode1"
        ],
        "imageId": "ami-0123456789abcdef",
        "instanceId": "i-0123456789abcdef0",
        "instanceType": "m5.large",
        "kernelId": "aki-0123456789abcdef",
        "marketplaceProductCodes": [
            "marketplaceCode1",
            "marketplaceCode2"
        ],
        "pendingTime": "2023-05-03T09:00:00Z",
        "privateIp": "10.0.0.10",
        "ramdiskId": "ari-0123456789abcdef",
        "region": "us-east-1",
        "version": "2022-04-01"
    }`
	logging.Logger.Debug("Debug message from subcommand")

	var doc mysns.ExtendedInstanceIdentityDocument

	err := json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&mysns.ExtendedInstanceIdentityDocument{})
	db.Create(&doc)

	var docs []mysns.ExtendedInstanceIdentityDocument
	db.Find(&docs)
	pp.Println(docs)

	jsonBytes, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonBytes))

	logging.Logger.Info("Info message from subcommand")
	logging.Logger.Debug("Debug message from subcommand")

	return nil
}
