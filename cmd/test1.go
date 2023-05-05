/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"

	"gorm.io/gorm"

	// "gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

	mysns "github.com/taylormonacelli/cheekspot/cmd/aws/sns"
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
	RootCmd.AddCommand(test1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type MultiString []string

func (s *MultiString) Scan(src interface{}) error {
	str, ok := src.(string)
	if !ok {
		return errors.New("failed to scan multistring field - source is not a string")
	}
	*s = strings.Split(str, ",")
	return nil
}

func (s MultiString) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}
	return strings.Join(s, ","), nil
}

func parseInstanceInfo(jsonStr string) (mysns.ExtendedInstanceIdentityDocument, error) {
	var info mysns.ExtendedInstanceIdentityDocument
	err := json.Unmarshal([]byte(jsonStr), &info)
	if err != nil {
		return mysns.ExtendedInstanceIdentityDocument{}, err
	}
	return info, nil
}

func createStructFromStr() (mysns.ExtendedInstanceIdentityDocument, error) {
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
	return meta, err
}

func test1() error {
	doc, err := createStructFromStr()
	if err != nil {
		fmt.Printf("fail at %s", err)
		return err
	}

	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&mysns.ExtendedInstanceIdentityDocument{})
	db.Create(&doc)

	var instances []mysns.ExtendedInstanceIdentityDocument
	db.Where("Epochtime = ?", doc.Epochtime).Find(&instances)
	pp.Println(instances)

	var doc1 mysns.ExtendedInstanceIdentityDocument
	db.First(&doc1, 1)

	pp.Println(doc1)

	jsBytes, _ := json.MarshalIndent(doc1, "", "    ")
	jsonStr := string(jsBytes)
	fmt.Println(jsonStr)

	return nil
}
