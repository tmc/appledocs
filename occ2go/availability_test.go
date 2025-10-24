package occ2go

import "testing"

func TestAvailability_IOSOnly(t *testing.T) {
	tests := []struct {
		name      string
		platforms map[string]string
		want      bool
	}{
		{
			name: "iOS only (not on macOS)",
			platforms: map[string]string{
				"iOS":     "3.0",
				"iPadOS":  "3.0",
				"tvOS":    "9.0",
				"visionOS": "1.0",
			},
			want: true,
		},
		{
			name: "macOS and iOS (universal)",
			platforms: map[string]string{
				"macOS": "10.14",
				"iOS":   "12.0",
			},
			want: false,
		},
		{
			name:      "no platforms (assume universal)",
			platforms: map[string]string{},
			want:      false,
		},
		{
			name: "macOS only",
			platforms: map[string]string{
				"macOS": "10.14",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Availability{
				IntroducedAt: tt.platforms,
			}
			if got := a.IOSOnly(); got != tt.want {
				t.Errorf("IOSOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}
