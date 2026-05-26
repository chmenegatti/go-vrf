package utilities

import (
	"fmt"
	"regexp"
)

// validName: ASCII letters, digits, underscore, hyphen, dot.
// Must not contain "..", may not be empty, max 100 chars.
var validName = regexp.MustCompile(`^[A-Za-z0-9_.-]{1,100}$`)

// ValidateFileName ensures the caller-supplied name is a safe basename
// — never a path. It is the chokepoint for any value flowing into
// os.Open / os.Create.
func ValidateFileName(name string) error {
	if name == "" {
		return fmt.Errorf("name is empty")
	}
	if !validName.MatchString(name) {
		return fmt.Errorf("name %q contains invalid characters (allowed: A-Z a-z 0-9 _ - .)", name)
	}
	for i := 0; i < len(name)-1; i++ {
		if name[i] == '.' && name[i+1] == '.' {
			return fmt.Errorf("name %q contains %q sequence", name, "..")
		}
	}
	return nil
}
