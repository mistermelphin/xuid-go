package xuid

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

// ToUUIDString represents XUID as UUID string. It's a short version of id.ToUUID().String(). Panics in case the XUID is invalid
func (id XUID) ToUUIDString() string {
	return id.ToUUID().String()
}

// ToUUID converts XUID as UUID. Panics in case the XUID is invalid
func (id XUID) ToUUID() uuid.UUID {
	u, err := ToUUID(id)
	if err != nil {
		panic(err)
	}
	return u
}

// ToUUID converts XUID as UUID
func ToUUID(xuid XUID) (uuid.UUID, error) {
	uuidBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(string(xuid))
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to decode: %w", err)
	}

	u, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to read: %w", err)
	}

	return u, nil
}

func parseUUID(s string) (XUID, error) {
	uid, err := uuid.Parse(s)
	if err != nil {
		return "", fmt.Errorf("parse UUID: %w", err)
	}
	return FromUUID(uid)
}

// FromUUID converts UUID to XUID
func FromUUID(id uuid.UUID) (XUID, error) {
	uuidBytes, err := id.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("fail to encode: %w", err)
	}

	x := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(uuidBytes)
	return XUID(x), nil
}

// MustFromUUID converts UUID to XUID. Panics in case of wrong format
func MustFromUUID(id uuid.UUID) XUID {
	x, err := FromUUID(id)
	if err != nil {
		panic(err)
	}
	return x
}
