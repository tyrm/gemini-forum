package gemini

import (
	"fmt"
	"testing"
)

func TestRegexValidUsername(t *testing.T) {
	tables := []struct {
		x string
		n bool
	}{
		{"user1234", true},
		{"12asdf33", true},
		{"xXxALB1C0RExXx", true},
		{"john.smith", true},
		{"john-smith", true},
		{"john_smith", true},
		{";DROP table \"users\";", false},
		{"jordy#1", false},
	}

	for i, table := range tables {
		i := i
		table := table
		name := fmt.Sprintf("[%d] Testing %s", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			match := reValidUsername.MatchString(table.x)
			if match != table.n {
				t.Errorf("[%d] regex match on %s failed, got: %v, want: %v,", i, table.x, match, table.n)
			}
		})
	}
}
