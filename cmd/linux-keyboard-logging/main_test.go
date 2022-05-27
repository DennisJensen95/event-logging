package main

import (
	"testing"
)

// test function
func TestKeyMap(t *testing.T) {
	key_map := getKeyMap()
	// t.Log("Map: ", key_map)
	if key_map != nil && len(key_map) > 0 {
		actualString := key_map["0x0041"]
		expectedString := "A"
		if actualString != expectedString {
			t.Errorf("Expected String(%s) is not same as"+
				" actual string (%s)", expectedString, actualString)
		}
	}
}
