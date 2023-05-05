/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/spf13/cobra"
)

// test5Cmd represents the test5 command
var test5Cmd = &cobra.Command{
	Use:   "test5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test5 called")
		test5()
	},
}

func init() {
	rootCmd.AddCommand(test5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type MyInstanceIdentityDocument struct {
	imds.InstanceIdentityDocument

	// Change these fields from []string to string
	DevpayProductCodes      string `json:"devpayProductCodes,omitempty"`
	MarketplaceProductCodes string `json:"marketplaceProductCodes,omitempty"`
	BillingProducts         string `json:"billingProducts,omitempty"`
}

func (i *MyInstanceIdentityDocument) MarshalJSON() ([]byte, error) {
	// Convert the fields that need to be changed to a comma-delimited string
	modifiedFields := []string{i.DevpayProductCodes, i.MarketplaceProductCodes, i.BillingProducts}
	for index, field := range modifiedFields {
		modifiedFields[index] = strings.Join(strings.Fields(field), ",")
	}

	// Call MarshalJSON on the embedded type
	jsonBytes, err := json.Marshal(&i.InstanceIdentityDocument)
	if err != nil {
		return nil, err
	}

	// Replace the original fields with the modified fields in the JSON bytes
	modifiedBytes := []byte(strings.ReplaceAll(string(jsonBytes), fmt.Sprintf(`"%v"`, modifiedFields[0]), fmt.Sprintf(`"%v"`, strings.Join(strings.Fields(modifiedFields[0]), ","))))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), fmt.Sprintf(`"%v"`, modifiedFields[1]), fmt.Sprintf(`"%v"`, strings.Join(strings.Fields(modifiedFields[1]), ","))))
	modifiedBytes = []byte(strings.ReplaceAll(string(modifiedBytes), fmt.Sprintf(`"%v"`, modifiedFields[2]), fmt.Sprintf(`"%v"`, strings.Join(strings.Fields(modifiedFields[2]), ","))))

	return modifiedBytes, nil
}

func test5() {
	// Parse the time string into a time.Time value
	pendingTime, err := time.Parse(time.RFC3339, "2022-05-04T22:33:57Z")
	if err != nil {
		// Handle the error
	}
	// Create an instance of the modified struct
	myDoc := &MyInstanceIdentityDocument{InstanceIdentityDocument: imds.InstanceIdentityDocument{
		AccountID:               "123456789012",
		Architecture:            "x86_64",
		AvailabilityZone:        "us-west-2a",
		ImageID:                 "ami-0c55b159cbfafe1f0",
		InstanceID:              "i-0e76fa926c788c919",
		InstanceType:            "t2.micro",
		KernelID:                "",
		PendingTime:             pendingTime,
		PrivateIP:               "172.31.16.132",
		RamdiskID:               "",
		Region:                  "us-west-2",
		Version:                 "2022-05-04",
		DevpayProductCodes:      []string{"code1", "code2"},
		MarketplaceProductCodes: []string{"code3", "code4"},
		BillingProducts:         []string{"product1", "product2"},
	}}

	// Marshal the modified struct to JSON

	jsonBytes, err := myDoc.MarshalJSON()
	if err != nil {
		// handle the error here
	}
	fmt.Println(string(jsonBytes))
}
