/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"soundLabCli/utils"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create a todo",
	Long:  `This command will create todo`,
	Run: func(cmd *cobra.Command, args []string) {

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		credentials := utils.Credentials{
			Username: username,
			Password: password,
		}
		fmt.Printf("Login in %+v\n", credentials)

		resp := utils.Login(credentials)
		fmt.Printf("Token created: %+v\n", resp.Token)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "provide username")
	loginCmd.Flags().StringP("password", "p", "", "provide password")
}
