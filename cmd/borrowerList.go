/*
Copyright © 2023 João Freire <jvwfreire@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	utils_aave "main/utils/aave"
)

// borrowerListCmd represents the borrowerList command
var borrowerListCmd = &cobra.Command{
	Use:   "borrowerList",
	Short: "Generates a list of borrowers",
	Long: `Generates a list of borrowers from the AAVE-v2 lending pool. 
Outputs a .csv file with the following format: "txHash, address, blockNumber".
This file is used to run commands such as the "queryHealth".
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating AAVE-V2 borrower list...")
		utils_aave.GenerateBorrowerList()
	},
}

func init() {
	rootCmd.AddCommand(borrowerListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// borrowerListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// borrowerListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
