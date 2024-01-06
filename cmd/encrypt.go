/*
Copyright © 2024 Daniel Aguilar danny.godev@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/DanyZugz/Encryptdata/internal/controllers"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypt called")
		dir, _ := os.Getwd()
		fmt.Println(dir)

		pss := controllers.GetPassword()

		controllers.Encrypt(args[0], pss)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

}
