/*
Copyright © 2023 João Freire <jvwfreire@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	utils_aave "main/utils/aave"
	utils_compound "main/utils/compound"
)

// borrowerListCmd represents the borrowerList command
var borrowerListCmd = &cobra.Command{
	Use:   "borrowerList",
	Short: "Generates a list of borrowers",
	Long: `Generates a list of borrowers from the AAVE-v2 lending pool. 
Outputs a .csv file with the following format: "txHash, address, blockNumber".
This file is used to run commands such as the "queryHealth".
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "aave" {
			fmt.Println("Generating AAVE-V2 borrower list...")
			utils_aave.GenerateBorrowerList()
		} else if args[0] == "compound" {
			fmt.Println("Generating Compound borrower list...")
			utils_compound.GenerateBorrowerList()
		} else {
			fmt.Println("Invalid argument. Please use 'aave' as the argument.")
		}
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
