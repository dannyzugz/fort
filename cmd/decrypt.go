/*
Copyright © 2024 Daniel Aguilar danny.godev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/DanyZugz/Encryptdata/internal/controllers"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt called")
		pss := controllers.ValidatePass()
		controllers.Decrypt(args[0], pss)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

}
