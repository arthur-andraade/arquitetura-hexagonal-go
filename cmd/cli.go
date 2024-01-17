package cmd

import (
	"arquitetura-hexagonal/adapters/cli"
	"fmt"

	"github.com/spf13/cobra"
)

// Variables
var action string
var productId string
var productName string
var productPrice float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI to do actions in products",
	Long: `This CLI allow to do this actions in product's table: create, enable, disable and get

		create - It will create the new product
		enable - It will change status of product to ENABLED
		disable - It will change status of product to DISABLED
		get - It will return the product saved (default action)
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "get", "Action in product's table")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "Product ID")
	cliCmd.Flags().StringVarP(&productName, "product", "n", "", "Product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "Product price")

}
