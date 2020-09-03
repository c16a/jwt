package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/color"
	"time"
)

func ParseToken(tokenToBeDecoded string, hmacSecret string) {
	tokenString := tokenToBeDecoded

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		PrintAlgorithm(token)
		PrintTokenDetails(token)

		hmacSampleSecret := []byte(hmacSecret)
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jsonString, err := PrettyJson(claims)
		if err != nil {
			panic(err)
		}
		fmt.Println(jsonString)
	} else {
		fmt.Println(err)
	}
}

func PrintAlgorithm(token *jwt.Token) {
	c := color.New(color.FgGreen).Add(color.Bold)
	fmt.Printf("\nSigned with: %s\n", c.Sprintf(token.Method.Alg()))
}

func PrintTokenDetails(token *jwt.Token) {

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
	fmt.Printf("Issued at: %s\n", c.Sprintf(convertTimestampToLocalString(issuedAt)))
	fmt.Printf("Expires at: %s\n", c.Sprintf(convertTimestampToLocalString(expiresAt)))
}

func convertTimestampToLocalString(t int64) string {
	utc := time.Unix(t, 0)
	return utc.Format(time.RFC1123)
}

const (
	empty  = ""
	indent = "    "
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, indent)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}