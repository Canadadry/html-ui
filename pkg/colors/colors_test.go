package colors

import (
	"testing"
)

func TestFromStringSuccess(t *testing.T) {
	tests := []struct {
		in   string
		expR uint8
		expG uint8
		expB uint8
	}{
		{"black", 0, 0, 0},
		{"darkblue", 0, 0, 139},
		{"#ff0080", 255, 0, 128},
		{"#123", 17, 34, 51},
		{"rgb(17,34,51)", 17, 34, 51},
	}
	for _, tt := range tests {
		result, err := FromString(tt.in)
		if err != nil {
			t.Fatalf("[%s] conversion failed %v", tt.in, err)
		}
		if result.R != tt.expR {
			t.Fatalf("[%s] R failed got %d exp %d", tt.in, result.R, tt.expR)
		}
		if result.G != tt.expG {
			t.Fatalf("[%s] G failed got %d exp %d", tt.in, result.G, tt.expG)
		}
		if result.B != tt.expB {
			t.Fatalf("[%s] B failed got %d exp %d", tt.in, result.B, tt.expB)
		}
	}
}

func TestFromStringErrro(t *testing.T) {
	tests := []struct {
		in string
	}{
		{""},
		{"#12"},
		{"#1234"},
		{"rgb("},
		{"rgb(17,34,51#"},
		{"rgb(17,34#51)"},
		{"rgb(17,34,#51)"},
		{"rgb(17,34,-51)"},
	}
	for _, tt := range tests {
		_, err := FromString(tt.in)
		if err == nil {
			t.Fatalf("[%s] conversion should have failed", tt.in)
		}
	}
}
