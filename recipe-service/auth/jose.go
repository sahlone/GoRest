package auth

import (
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/model"
	. "gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

var issuer = "urn:dummy:recipe:idp"
var subject = "userId"
var audience = "dummy-audience"

var claims = jwt.Claims{
	Subject:  subject,
	Issuer:   issuer,
	Audience: jwt.Audience{audience},
}

// THis should not be here . It should come through some key management system like AWS KMS
var signingKey = "ZaSzQZIZVHaHe7qy2t0QZDvM4YnnD57dAu27xR8b1fHpTAgmxaNL54Qt8894WPx1"

// This should come from database and not be static
var username = "admin"
var password = "password"

func Authorize(user, pass string) (string, error) {

	if user == username && pass == password {
		var symKey = []byte(signingKey)

		signer, err := NewSigner(SigningKey{Algorithm: HS256, Key: symKey},
			(&SignerOptions{}).WithType("JWT"))
		if err != nil {
			logger.Error("failed to create signer", err)
			return "", err
		}
		return jwt.Signed(signer).Claims(claims).CompactSerialize()
	} else {
		return "", model.ErrCredsInvalid
	}

}

func ValidateToken(token string) error {

	jsonWebSignature, err := ParseSigned(token)
	if err != nil {
		logger.Info("invalid token recieved . Token :%v, Err:%v", token, err)
		return model.ErrTokenInvalid
	}

	_, err = jsonWebSignature.Verify([]byte(signingKey))

	if err != nil {
		logger.Info("token signature failed . Token :%v, Err:%v", token, err)
		return model.ErrTokenInvalid
	}

	return nil
}
