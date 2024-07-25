package xuid

import (
	"fmt"
	"github.com/google/uuid"
)

// XUID is a short URL-friendly compressed UUID.
//
// More information about XUID: https://xuid.org
type XUID string

// String represents XUID as string
func (id XUID) String() string {
	return string(id)
}

// Parse string to XUID. It supports a 22-len XUID format, or 32,36,38,45 UUID format of string
func Parse(s string) (XUID, error) {
	l := len(s)
	switch l {
	case 22:
		x := XUID(s)
		if _, err := ToUUID(x); err != nil {
			return "", fmt.Errorf("invalid xuid: %w", err)
		}
		return x, nil
	case 32, 36, 38, 45:
		return parseUUID(s)
	default:
		return "", fmt.Errorf("invalid length (%d) of id", l)
	}
}

// MustParse string to XUID. It supports a 22-len XUID format, or 32,36,38,45 UUID format of string. Panics in case
// the provided id has wrong format
func MustParse(s string) XUID {
	x, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return x
}

// New generates a new random XUID
func New() XUID {
	return MustFromUUID(uuid.New())
}
