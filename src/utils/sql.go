package utilities

import (
	"fmt"
	"strings"
)

// EscapeLiteral returns a safely-quoted SQL string literal for s.
//
// Single quotes are doubled (ANSI SQL). NUL bytes are rejected because
// many SQL parsers and drivers treat them as string terminators, which
// can be exploited to truncate the literal.
//
// The result includes the surrounding single quotes, e.g. EscapeLiteral("a'b")
// returns "'a”b'". Callers should NOT add their own quotes.
func EscapeLiteral(s string) (string, error) {
	if strings.ContainsRune(s, 0) {
		return "", fmt.Errorf("string contains NUL byte")
	}
	return "'" + strings.ReplaceAll(s, "'", "''") + "'", nil
}

// MustEscapeLiteral is EscapeLiteral but panics on error. Use only when
// inputs are known not to contain NUL bytes (e.g. UUIDs, fixed enums).
func MustEscapeLiteral(s string) string {
	out, err := EscapeLiteral(s)
	if err != nil {
		panic(err)
	}
	return out
}
