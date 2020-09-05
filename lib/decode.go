package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"os"
	"time"
)

//ParseToken parses a JWT with support for HMAC and RSA
//hmacSecret - optional param if token is signed with HMAC
//publicKeyFile - mandatory param if token is signed with RSA
func ParseToken(tokenToBeDecoded, hmacSecret, publicKeyFile string, w io.Writer) error {
	tokenString := tokenToBeDecoded

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			printAlgorithm(token, w)
			printTokenDetails(token, w)

			hmacSampleSecret := []byte(hmacSecret)
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		}

		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			printAlgorithm(token, w)
			printTokenDetails(token, w)

			if len(publicKeyFile) <= 0 {
				return nil, errors.New("public key is mandatory for RSA decoding")
			}
			_, err := os.Stat(publicKeyFile)
			if os.IsNotExist(err) {
				return nil, errors.New("could not find public key file")
			}
			publicKeyBytes, err := ioutil.ReadFile(publicKeyFile)
			if err != nil {
				return nil, errors.New("could not read public key file")
			}
			return jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
		}

		return nil, errors.New("unknown signing algorithm used")
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jsonString, err := prettyJSON(claims)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, jsonString)
	} else {
		fmt.Fprintln(w, err)
	}

	return nil
}

func printAlgorithm(token *jwt.Token, w io.Writer) {
	c := color.New(color.FgGreen).Add(color.Bold)
	fmt.Fprintf(w, "\nSigned with: %s\n", c.Sprintf(token.Method.Alg()))
}

func printTokenDetails(token *jwt.Token, w io.Writer) {

	var issuedAt int64
	var expiresAt int64

	c := color.New(color.FgGreen).Add(color.Bold)
	if claims, ok := token.Claims.(jwt.StandardClaims); ok {
		issuedAt = claims.IssuedAt
		expiresAt = claims.ExpiresAt
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			i, isIntVal := claims["iat"].(float64)
			if isIntVal {
				issuedAt = int64(i)
			}

			e, isIntVal := claims["exp"].(float64)
			if isIntVal {
				expiresAt = int64(e)
			}
		}
	}
	fmt.Fprintf(w, "Issued at: %s\n", c.Sprintf(convertTimestampToLocalString(issuedAt)))
	fmt.Fprintf(w, "Expires at: %s\n", c.Sprintf(convertTimestampToLocalString(expiresAt)))
}

func convertTimestampToLocalString(t int64) string {
	utc := time.Unix(t, 0)
	return utc.Format(time.RFC1123)
}

const (
	empty  = ""
	indent = "    "
)

func prettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, indent)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}
