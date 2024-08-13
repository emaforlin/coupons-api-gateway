package config

import (
	"fmt"
	"os"
)

// var (
// 	minTokenTTL     = 10 * time.Minute
// 	maxTokenTTL     = 3 * time.Hour
// 	defaultTokenTTL = 15 * time.Minute
// )

func MustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

// func NormalizeTime(ttlTime time.Duration) time.Duration {
// 	var ttl = ttlTime

// 	if ttl <= 0 {
// 		return defaultTokenTTL
// 	}

// 	if ttl < minTokenTTL {
// 		ttl = minTokenTTL
// 	} else if ttl > maxTokenTTL {
// 		ttl = maxTokenTTL
// 	}
// 	return ttl
// }
