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

var createMarketplaceCmd = &cobra.Command{
	Use:   "createMarketplace",
	Short: "Creates a marketplace",
	Long:  `This command will create a new marketplace`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		marketName := utils.MarketName{
			Name: name,
		}
		err := Db.View(func(txn *badger.Txn) error {
			item, _ := txn.Get([]byte("login"))
			item.Value(func(val []byte) error {
				fmt.Printf("Creating market with name %+v\n, token: %+v\n", marketName, string(val))

				if string(val) == "" {
					fmt.Println("Error: Please login first")
					return nil
				}

				resp, error := utils.CreateMarketplace(marketName, string(val))
				if error == nil {
					fmt.Printf("Market created: %+v\n", resp)
				}
				return nil
			})
			return nil
		})
		if err != nil {
			println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createMarketplaceCmd)
	createMarketplaceCmd.Flags().StringP("name", "n", "", "provide name")
}
