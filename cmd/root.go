/*
Copyright © 2023 Ricardo Santis
*/
package cmd

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/spf13/cobra"
	"os"
)

var Db, err = badger.Open(badger.DefaultOptions("/tmp/badger"))

var rootCmd = &cobra.Command{
	Use:   "soundLabCli",
	Short: "soundLabCli is an admin cli tool for soundLabConnect project",
	Long:  "soundLabCli is an admin cli tool for soundLabConnect project",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createMarketplaceCmd)
}
