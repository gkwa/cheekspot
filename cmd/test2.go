/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/k0kubun/pp"

	"github.com/spf13/cobra"
)

// test2Cmd represents the test2 command
var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test2 called")
		test2()
	},
}

func init() {
	RootCmd.AddCommand(test2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func test2() {
	jsonStr := `{
        
    }`
	msg, err := jsonStrToSNSMessage(jsonStr)
	if err != nil {
		panic(err)
	}
	pp.Printf("my message %s", msg)
}

func jsonStrToSNSMessage(jsonStr string) (types.Message, error) {
	var msg types.Message
	err := json.Unmarshal([]byte(jsonStr), &msg)
	if err != nil {
		return types.Message{}, err
	}
	return msg, nil
}
