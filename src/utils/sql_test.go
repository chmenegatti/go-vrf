package utilities

import "testing"

func TestEscapeLiteral(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"plain", "hello", "'hello'", false},
		{"empty", "", "''", false},
		{"single quote", "O'Brien", "'O''Brien'", false},
		{"injection attempt", "'; DROP TABLE users;--", "'''; DROP TABLE users;--'", false},
		{"many quotes", "a'b'c", "'a''b''c'", false},
		{"backslash kept literal", `a\b`, `'a\b'`, false},
		{"unicode", "café", "'café'", false},
		{"nul byte rejected", "a\x00b", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := EscapeLiteral(tc.in)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got %q", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
