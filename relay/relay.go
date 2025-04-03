package relay

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// ToGlobalID encodes a type name and ID into a global ID string.
func ToGlobalID(typeName, id string) string {
	raw := fmt.Sprintf("%s:%s", typeName, id)
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

// FromGlobalID decodes a global ID string into its type name and ID.
func FromGlobalID(globalID string) (typeName, id string, err error) {
	decoded, err := base64.StdEncoding.DecodeString(globalID)
	if err != nil {
		return "", "", fmt.Errorf("failed to decode global ID: %w", err)
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid global ID format")
	}

	return parts[0], parts[1], nil
}
