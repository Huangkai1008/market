package constants

import "time"

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

const FPath = "conf/conf.toml"

const (
	JwtExpireDuration = 6 * time.Hour
	JwtIssuer         = "market"
)
