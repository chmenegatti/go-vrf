package utilities

import "testing"

func TestValidateFileName(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		wantErr bool
	}{
		{"plain", "DB-Shared_10", false},
		{"with dot", "T0-Cluster_4.json", false},
		{"empty", "", true},
		{"slash", "etc/passwd", true},
		{"backslash", `etc\passwd`, true},
		{"dotdot", "../etc/passwd", true},
		{"dotdot only", "..", true},
		{"leading dot ok", ".hidden", false},
		{"null byte", "a\x00b", true},
		{"space", "name with space", true},
		{"unicode", "café", true},
		{"too long", string(make([]byte, 101)), true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateFileName(tc.in)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ValidateFileName(%q) err=%v, wantErr=%v", tc.in, err, tc.wantErr)
			}
		})
	}
}
