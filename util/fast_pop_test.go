package util

import (
	"testing"
)

func TestFastPopString(t *testing.T) {
	slice := []string{"foo", "bar", "baz", "qux"}

	tables := []struct {
		x    string
		n    map[string]bool
		nLen int
	}{
		{"quz", map[string]bool{"foo": true, "bar": true, "baz": true, "qux": true}, 4},
		{"bar", map[string]bool{"foo": true, "bar": false, "baz": true, "qux": true}, 3},
		{"beep", map[string]bool{"foo": true, "bar": false, "baz": true, "qux": true}, 3},
		{"foo", map[string]bool{"foo": false, "bar": false, "baz": true, "qux": true}, 2},
		{"qux", map[string]bool{"foo": false, "bar": false, "baz": true, "qux": false}, 1},
		{"baz", map[string]bool{"foo": false, "bar": false, "baz": false, "qux": false}, 0},
		{"thud", map[string]bool{"foo": false, "bar": false, "baz": false, "qux": false}, 0},
	}

	for _, table := range tables {
		slice = FastPopString(slice, table.x)

		// Check length is correct
		if len(slice) != table.nLen {
			t.Errorf("slice length was incorrect, got: %v, want: %v.", len(slice), table.nLen)
		}

		for key, shouldExist := range table.n {
			index := -1

			for i, v := range slice {
				if v == key {
					index = i
				}
			}

			if shouldExist && index == -1 {
				t.Errorf("element %s missing from slice. should be present.", key)
			}

			if !shouldExist && index != -1 {
				t.Errorf("element %s exists in slice. should not be present.", key)
			}
		}
	}
}
