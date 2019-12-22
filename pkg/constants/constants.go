package constants

import "time"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const FPath = "configs/conf.toml"

const (
	JwtExpireDuration = 6 * time.Hour
	JwtIssuer         = "market"
)
