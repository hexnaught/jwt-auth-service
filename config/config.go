package config

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// App represents the app config that may need to be accessed frquently,
// for example. caching config is much faster than constantly grabbing vars.
type App struct {
	BCryptCostFactor int    // The cost factor we use in bcrypt method calls
	Debug            bool   // Are we running in debug mode - Might not use this, might use compile directives
	JWTSecret        string // Secret used for generating, signing and validating JWT tokens
	JWTTTL           int    // Time to live for JWT tokens (non-refresh)
}

// LoadAppConfig parses/loads app config from env vars
func LoadAppConfig() *App {

	// Attempt to load the bcrypt cost we want to use from env vars
	// Set the cost we want to use depending on if it is too high, too low or not set
	bcryptCostStr := os.Getenv("BCRYPT_COST_FACTOR")
	bcryptCost, err := strconv.Atoi(bcryptCostStr)
	if err != nil {
		bcryptCost = bcrypt.DefaultCost
	} else if bcryptCost < bcrypt.DefaultCost {
		bcryptCost = bcrypt.DefaultCost
	} else if bcryptCost > bcrypt.MaxCost {
		bcryptCost = bcrypt.MaxCost
	}

	// Default debug mode to false, can be overwritten to try by setting the env var to anything
	// other than "false", "no", "0", "" (Blank/Empty String)
	debugMode := false
	debugModeStr, isSet := os.LookupEnv("DEBUG_MODE")
	if isSet && debugModeStr != "false" && debugModeStr != "no" && debugModeStr != "0" {
		debugMode = true
	}

	jwtTTLStr := os.Getenv("JWT_TTL")
	jwtTTLMinutes, err := strconv.Atoi(jwtTTLStr)
	if err != nil {
		jwtTTLMinutes = 15
	}

	return &App{
		BCryptCostFactor: bcryptCost,
		Debug:            debugMode,
		JWTSecret:        os.Getenv("JWT_SECRET"),
		JWTTTL:           jwtTTLMinutes,
	}
}
