package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	customLogger "github.com/innvtseeds/wdic-server/library/logger"
)

var myLogger = customLogger.NewLogger()

/*
Steps are as follows
1. Establish the secret key from the ENV Variable
2. Start preparing the token payload. I have kept the creation generic to allow for any type of token creation
3. Generate the a new token using jwt.NewWithClaims
4. Conver that token to string
5. Return the string if no error, else return the error
*/
func GenerateToken(payload map[string]interface{}) (*string, error) {

	SECRET_KEY := os.Getenv("JWT_KEY")
	if SECRET_KEY == "" {
		myLogger.Error("LIBRARY :: JWT :: SECRET KEY MISSING")
		return nil, errors.New("JWT SECRET KEY MISSING")
	}

	var hmacSampleSecret = []byte(SECRET_KEY)

	tokenPayload := jwt.MapClaims{
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour).Unix(),
	}
	// Iterate through variadic payload and add to tokenPayload
	// Add key-value pairs from the payload to tokenPayload
	for key, value := range payload {
		tokenPayload[key] = value
	}

	myLogger.Info("LIBRARY :: JWT TOKEN GENERATION :: PAYLOAD ::  ", tokenPayload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenPayload)

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		myLogger.Error("LIBRARY :: JWT :: FAILED TO CONVERT JWT TO STRING", err)
		return nil, err
	}

	return &tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	SECRET_KEY := os.Getenv("JWT_KEY")
	if SECRET_KEY == "" {
		myLogger.Error("LIBRARY :: JWT :: SECRET KEY MISSING")
		return nil, errors.New("JWT SECRET KEY MISSING")
	}

	hmacSampleSecret := []byte(SECRET_KEY)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	switch {
	case token.Valid:
		return token, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		myLogger.Error("LIBRARY :: JWT PARSING :: That's not even a token")
		return nil, jwt.ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		myLogger.Error("LIBRARY :: JWT PARSING :: Invalid Signature")
		return nil, jwt.ErrTokenSignatureInvalid
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		myLogger.Error("LIBRARY :: JWT PARSING :: Timing is everything, token expired")
		return nil, jwt.ErrTokenExpired
	default:
		myLogger.Error("LIBRARY :: JWT PARSING :: Couldn't handle this token:", err)
		return nil, err
	}
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	SECRET_KEY := os.Getenv("JWT_KEY")
	if SECRET_KEY == "" {
		myLogger.Error("LIBRARY :: JWT :: SECRET KEY MISSING")
		return nil, errors.New("JWT SECRET KEY MISSING")
	}

	hmacSampleSecret := []byte(SECRET_KEY)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil {
		myLogger.Error("LIBRARY :: JWT DECODING :: Failed to decode token:", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		myLogger.Error("LIBRARY :: JWT DECODING :: Unable to extract claims from token")
		return nil, errors.New("Unable to extract claims from token")
	}

	return claims, nil
}
