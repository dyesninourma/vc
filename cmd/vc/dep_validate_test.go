package main

import "testing"

func TestValidateDependency(t *testing.T) {
	cases := []struct {
		name      string
		issue     string
		dependsOn string
		wantErr   bool
	}{
		{"normal", "vc-1", "vc-2", false},
		{"missing left", "", "vc-2", true},
		{"missing right", "vc-1", "", true},
		{"self dependency", "vc-1", "vc-1", true},
	}
	for _, tc := range cases {
		err := validateDependency(tc.issue, tc.dependsOn)
		if (err != nil) != tc.wantErr {
			t.Fatalf("%s: err = %v, wantErr %v", tc.name, err, tc.wantErr)
		}
	}
}
