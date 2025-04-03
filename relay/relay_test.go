package relay

import (
	"testing"
)

func TestGlobalIDEncoding(t *testing.T) {
	id := "123"
	typeName := "User"

	globalID := ToGlobalID(typeName, id)
	decodedType, decodedID, err := FromGlobalID(globalID)

	if err != nil {
		t.Fatalf("error decoding global ID: %v", err)
	}
	if decodedType != typeName || decodedID != id {
		t.Fatalf("expected %s:%s, got %s:%s", typeName, id, decodedType, decodedID)
	}
}
