package database

import "testing"

func Test_getDsn(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDsn(); got != tt.want {
				t.Errorf("getDsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
