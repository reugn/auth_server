package auth

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v4"
	"github.com/reugn/auth-server/repository"
)

// TokenType represents a token type.
type TokenType int

const (
	// BearerToken is an opaque string, not intended to have any meaning to clients using it.
	// Some servers will issue tokens that are a short `string` of hexadecimal characters,
	// while others may use structured tokens such as JSON Web Tokens.
	BearerToken TokenType = iota

	// BasicToken is a string where credentials is the base64 encoding of id and
	// password joined by a single colon :
	BasicToken
)

// ToString converts the TokenType to a string.
func (t TokenType) ToString() string {
	return [...]string{"Bearer", "Basic"}[t]
}

// Claims is the custom JWT claims container.
type Claims struct {
	jwt.StandardClaims
	Username string              `json:"user"`
	Role     repository.UserRole `json:"role"`
}

// AccessToken represents an access token.
type AccessToken struct {
	Token   string `json:"access_token"`
	Type    string `json:"token_type"`
	Expires int64  `json:"expires_in"`
}

// Marshal marshals the AccessToken to a JSON string.
func (t *AccessToken) Marshal() string {
	jsonByteArray, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(jsonByteArray)
}
