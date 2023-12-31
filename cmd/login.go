/*
Copyright © 2023 Ricardo Santis
*/
package cmd

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
	"soundLabCli/utils"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long:  `This command will login`,
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
		Db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte("login"), []byte(resp.Token))
			return err
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "provide username")
	loginCmd.Flags().StringP("password", "p", "", "provide password")
}
