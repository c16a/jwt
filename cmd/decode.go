package cmd

import (
	"github.com/c16a/jwt/lib"
	"github.com/spf13/cobra"
)

var tokenToBeDecoded string
var hmacSecret string

func init() {
	decodeCommand.Flags().StringVarP(&tokenToBeDecoded, "token", "t", "", "token to be decoded")
	decodeCommand.Flags().StringVarP(&hmacSecret, "secret", "s", "", "optional hmac secret")
	rootCmd.AddCommand(decodeCommand)
}

var decodeCommand = &cobra.Command{
	Use: "decode",
	Short: "Decodes a JWT token",
	Run: func(cmd *cobra.Command, args []string) {
		lib.ParseToken(tokenToBeDecoded, hmacSecret)
	},
}