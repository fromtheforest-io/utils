package uuid

import "testing"

func TestUuid(t *testing.T) {
	uuid1 := Uuid()
	uuid2 := Uuid()

	if uuid1 == "" {
		t.Errorf("Expected a non-empty UUID, got an empty string")
	}

	if uuid2 == "" {
		t.Errorf("Expected a non-empty UUID, got an empty string")
	}

	if uuid1 == uuid2 {
		t.Errorf("Expected unique UUIDs, but got the same: %s", uuid1)
	}
}
