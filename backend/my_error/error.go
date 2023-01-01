package myError

import "errors"

var (
	PointerIsNil         = errors.New("Pointer is nil")
	ClaimsTransFailed    = errors.New("Claims trans failed")
	TokenClaimsIsInvalid = errors.New("TokenClaims is invalid")
	JwtExpireTime        = errors.New("Jwt expire time")
)
