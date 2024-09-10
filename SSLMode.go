package postgres

import (
	"strings"
)

// PGSSLMODE
const (
	SSLModeDisable SSLMode = iota
	SSLModeRequire
	SSLModeVerifyCA
	SSLModeVerifyFull
)

// Convert string to SSLMode
//
// empty: disable
//
// default: require
func SSLModeFromString(mode string) SSLMode {
	switch v := strings.ToLower(mode); v {
	case "verify-ca":
		return SSLModeVerifyCA
	case "verify-full":
		return SSLModeVerifyFull
	case "disable", "":
		return SSLModeDisable
	default:
		return SSLModeRequire
	}
}

type SSLMode uint8

// Return SSL mode string
//
// default: disable
func (m SSLMode) String() string {
	switch m {
	case SSLModeRequire:
		return "require"
	case SSLModeVerifyCA:
		return "verify-ca"
	case SSLModeVerifyFull:
		return "verify-full"
	default:
		return "disable"
	}
}
