/*
Copyright © 2024 Daniel Aguilar danny.godev@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/dannyzugz/fort/internal/controllers"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "reverses encryption to make data readable and accessible again",
	Long:  ` `,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt called")
		pss := controllers.ValidatePass()
		controllers.Decrypt(args[0], pss)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

}
