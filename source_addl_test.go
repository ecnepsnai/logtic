package logtic

import "testing"

func TestEscapeCharacters(t *testing.T) {
	check := func(in, expect string) {
		result := escapeCharacters(in)
		if result != expect {
			t.Errorf("Incorrect result for escaped string. Got '%s' expected '%s'", result, expect)
		}
	}

	check("Hello\nWorld!", "Hello\\nWorld!")
	check("Hello\\nWorld!", "Hello\\nWorld!")
}
