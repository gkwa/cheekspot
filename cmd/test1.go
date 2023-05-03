/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"gorm.io/gorm"

	// "gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

// test1Cmd represents the test1 command
var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test1 called")
		test1()
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type InstanceInfo struct {
	AccountId        string `json:"accountId"`
	Architecture     string `json:"architecture"`
	AvailabilityZone string `json:"availabilityZone"`
	// BillingProducts         []string `json:"billingProducts"`
	// DevpayProductCodes      []string `json:"devpayProductCodes"`
	Epochtime    int64  `json:"epochtime"`
	ImageId      string `json:"imageId"`
	InstanceId   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
	KernelId     string `json:"kernelId"`
	// MarketplaceProductCodes []string `json:"marketplaceProductCodes"`
	PendingTime string `json:"pendingTime"`
	PrivateIp   string `json:"privateIp"`
	RamdiskId   string `json:"ramdiskId"`
	Region      string `json:"region"`
	Version     string `json:"version"`
}

func parseInstanceInfo(jsonStr string) (InstanceInfo, error) {
	var info InstanceInfo
	err := json.Unmarshal([]byte(jsonStr), &info)
	if err != nil {
		return InstanceInfo{}, err
	}
	return info, nil
}

func createStructFromStr() (InstanceInfo, error) {
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
	meta, err := parseInstanceInfo(jsonStr)
	if err != nil {
		panic(err)
	}
	meta.Epochtime = 1654321987
	pp.Println(meta)
	return meta, err
}

func test1() error {
	meta, err := createStructFromStr()
	if err != nil {
		fmt.Printf("fail at %s", err)
		return err
	}

	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// logger.Logger.Info("Hello, world!")
	db.AutoMigrate(&InstanceInfo{})

	// Migrate specific table
	db.Table("instances").AutoMigrate(&InstanceInfo{})

	// Create a new meta record
	db.Create(&meta)

	// Query for users born within the last month
	var instances []InstanceInfo
	db.Where("Epochtime >= ?", meta.Epochtime).Find(&instances)

	// Print the results
	for _, instance := range instances {
		pp.Println(instance)
	}
	return nil
}
