package kllid

import (
	"testing"
)

func TestNew(t *testing.T) {
	id := New(10)
	if len(id) != 10 {
		t.Errorf("id incorrect length")
	}
}
