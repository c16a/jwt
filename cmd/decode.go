package cmd

import (
	"github.com/c16a/jwt/lib"
	"github.com/spf13/cobra"
)

var tokenToBeDecoded string
var hmacSecret string
var publicKeyFile string

func init() {

	// HMAC support
	decodeCommand.Flags().StringVarP(&tokenToBeDecoded, "token", "t", "", "token to be decoded")
	decodeCommand.Flags().StringVarP(&hmacSecret, "secret", "s", "", "optional hmac secret")

	// RSA support
	decodeCommand.Flags().StringVarP(&publicKeyFile, "publicKeyFile", "", "", "public key file path to decode RSA tokens")

	rootCmd.AddCommand(decodeCommand)
}

var decodeCommand = &cobra.Command{
	Use: "decode",
	Short: "Decodes a JWT token",
	RunE: func(cmd *cobra.Command, args []string) error {
		return lib.ParseToken(tokenToBeDecoded, hmacSecret, publicKeyFile)
	},
}