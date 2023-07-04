/*
Copyright © 2023 Ricardo Santis
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "soundLabCli",
	Short: "soundLabCli is an admin cli tool for soundLabConnect project",
	Long:  "soundLabCli is an admin cli tool for soundLabConnect project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		cmd.Parent().PersistentFlags().Set("token", token)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("token", "", "Bearer token")
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(createMarketplaceCmd)
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read config file:", err)
	}
}
