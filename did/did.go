package did

import (
	"errors"
	"strings"
)

// DID represents a Decentralized Identifier
type DID struct {
	Method string
	ID     string
}

// Parse parses a DID string and returns a DID struct
func Parse(did string) (*DID, error) {
	if !strings.HasPrefix(did, "did:") {
		return nil, errors.New("invalid DID: must start with 'did:'")
	}

	parts := strings.SplitN(did[4:], ":", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid DID: must contain method and ID")
	}

	return &DID{
		Method: parts[0],
		ID:     parts[1],
	}, nil
}

// NewDID creates a new DID instance
func NewDID(method, id string) *DID {
	return &DID{
		Method: method,
		ID:     id,
	}
}
