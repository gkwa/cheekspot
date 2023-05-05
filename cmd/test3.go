/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

// test3Cmd represents the test3 command
var test3Cmd = &cobra.Command{
	Use:   "test3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test3 called")
		test3()
	},
}

func init() {
	rootCmd.AddCommand(test3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test3() {
	jsonStr := `{
        "accountId": "348759328109",
        "architecture": "arm64",
        "availabilityZone": "us-east-1c",
        "billingProducts": [
            "bp-8f5a09f1"
        ],
        "devpayProductCodes": null,
        "epochtime": 1682801174,
        "imageId": "ami-0f4836e0909f7315f",
        "instanceId": "i-0388847dffe58da42",
        "instanceType": "m5a.4xlarge",
        "kernelId": null,
        "marketplaceProductCodes": null,
        "pendingTime": "2023-04-29T15:45:23Z",
        "privateIp": "10.1.2.15",
        "ramdiskId": null,
        "region": "us-east-1",
        "version": "2022-11-07"
    }`
	doc, err := genIdentiyDocFromJsonStr(jsonStr)
	if err != nil {
		panic(err)
	}

	// Print the instance identity document
	fmt.Println("InstanceIdentityDocument:")
	fmt.Println("  AccountID:", aws.ToString(&doc.AccountID))
	fmt.Println("  Architecture:", aws.ToString(&doc.Architecture))
	fmt.Println("  AvailabilityZone:", aws.ToString(&doc.AvailabilityZone))
	fmt.Println("  ImageID:", aws.ToString(&doc.ImageID))
	fmt.Println("  InstanceID:", aws.ToString(&doc.InstanceID))
	fmt.Println("  InstanceType:", aws.ToString(&doc.InstanceType))
	fmt.Println("  PrivateIP:", aws.ToString(&doc.PrivateIP))
	fmt.Println("  Region:", aws.ToString(&doc.Region))
	fmt.Println("  Version:", aws.ToString(&doc.Version))
}

func genIdentiyDocFromJsonStr(jsonStr string) (imds.InstanceIdentityDocument, error) {
	// create a map to hold the decoded JSON data
	var doc imds.InstanceIdentityDocument

	// unmarshal the JSON data into the map
	err := json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		panic(err)
	}
	return doc, err
}
