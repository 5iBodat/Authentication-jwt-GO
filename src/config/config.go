package config

import "time"

type JWTconfig struct {
	SecretKey     string
	TokenLifespan time.Duration
}

var JwtConfig = JWTconfig{
	SecretKey:     "jwt-secret-key",
	TokenLifespan: time.Hour * 24,
}
