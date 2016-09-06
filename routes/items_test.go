package routes

import (
	"testing"
)

func TestAllItems(t *testing.T) {
	resetItems()
	if len(allItems) != 20 {
		t.Error("20 items not inserted")
	}

	if len(mapItems) != 20 {
		t.Error("20 items not inserted in map")
	}
}
