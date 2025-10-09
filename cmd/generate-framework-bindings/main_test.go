package main

import (
	"testing"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		wantMajor int
		wantMinor int
		wantOK    bool
	}{
		{"simple major", "10", 10, 0, true},
		{"major.minor", "10.14", 10, 14, true},
		{"single digit minor", "10.9", 10, 9, true},
		{"large versions", "15.4", 15, 4, true},
		{"zero version", "0.0", 0, 0, true},
		{"invalid format", "invalid", 0, 0, false},
		{"empty string", "", 0, 0, false},
		{"extra dots cause failure", "10.14.5", 0, 0, false}, // Patch versions not supported
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMajor, gotMinor, gotOK := parseVersion(tt.version)
			if gotMajor != tt.wantMajor || gotMinor != tt.wantMinor || gotOK != tt.wantOK {
				t.Errorf("parseVersion(%q) = (%d, %d, %v), want (%d, %d, %v)",
					tt.version, gotMajor, gotMinor, gotOK, tt.wantMajor, tt.wantMinor, tt.wantOK)
			}
		})
	}
}

func TestCompareVersionStrings(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want int
	}{
		// Equal
		{"equal same", "10.14", "10.14", 0},
		{"equal zero", "0.0", "0.0", 0},

		// Less than
		{"major less", "10.14", "11.0", -1},
		{"minor less", "10.9", "10.14", -1},
		{"single vs major.minor", "10", "10.1", -1},

		// Greater than
		{"major greater", "11.0", "10.14", 1},
		{"minor greater", "10.14", "10.9", 1},
		{"major.minor vs single", "10.1", "10", 1},

		// Edge cases
		{"10.10 vs 10.9", "10.10", "10.9", 1}, // Semantic comparison, not lexicographic
		{"15.4 vs 15.0", "15.4", "15.0", 1},
		{"10.0 vs 10", "10.0", "10", 0}, // Both are 10.0

		// Invalid versions fallback to lexicographic
		{"invalid both", "abc", "def", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareVersionStrings(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("compareVersionStrings(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExtractAvailability(t *testing.T) {
	tests := []struct {
		name   string
		input  []Platform
		want Availability
	}{
		{
			name: "single platform",
			input: []Platform{
				{Name: "macOS", IntroducedAt: "10.14", Beta: false},
			},
			want: Availability{
				IntroducedAt: map[string]string{"macOS": "10.14"},
				DeprecatedAt: map[string]string{},
				Beta:         false,
			},
		},
		{
			name: "multiple platforms",
			input: []Platform{
				{Name: "macOS", IntroducedAt: "10.14"},
				{Name: "iOS", IntroducedAt: "12.0"},
			},
			want: Availability{
				IntroducedAt: map[string]string{
					"macOS": "10.14",
					"iOS":   "12.0",
				},
				DeprecatedAt: map[string]string{},
			},
		},
		{
			name: "deprecated platform",
			input: []Platform{
				{Name: "macOS", IntroducedAt: "10.0", DeprecatedAt: "10.14"},
			},
			want: Availability{
				IntroducedAt: map[string]string{"macOS": "10.0"},
				DeprecatedAt: map[string]string{"macOS": "10.14"},
			},
		},
		{
			name: "beta API",
			input: []Platform{
				{Name: "macOS", IntroducedAt: "15.4", Beta: true},
			},
			want: Availability{
				IntroducedAt: map[string]string{"macOS": "15.4"},
				DeprecatedAt: map[string]string{},
				Beta:         true,
			},
		},
		{
			name: "unavailable platform skipped",
			input: []Platform{
				{Name: "macOS", IntroducedAt: "10.14"},
				{Name: "iOS", Unavailable: true},
			},
			want: Availability{
				IntroducedAt: map[string]string{"macOS": "10.14"},
				DeprecatedAt: map[string]string{},
			},
		},
		{
			name:  "empty platforms",
			input: []Platform{},
			want: Availability{
				IntroducedAt: map[string]string{},
				DeprecatedAt: map[string]string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractAvailability(tt.input)

			// Compare IntroducedAt maps
			if len(got.IntroducedAt) != len(tt.want.IntroducedAt) {
				t.Errorf("IntroducedAt length mismatch: got %d, want %d",
					len(got.IntroducedAt), len(tt.want.IntroducedAt))
			}
			for k, v := range tt.want.IntroducedAt {
				if got.IntroducedAt[k] != v {
					t.Errorf("IntroducedAt[%q] = %q, want %q", k, got.IntroducedAt[k], v)
				}
			}

			// Compare DeprecatedAt maps
			if len(got.DeprecatedAt) != len(tt.want.DeprecatedAt) {
				t.Errorf("DeprecatedAt length mismatch: got %d, want %d",
					len(got.DeprecatedAt), len(tt.want.DeprecatedAt))
			}
			for k, v := range tt.want.DeprecatedAt {
				if got.DeprecatedAt[k] != v {
					t.Errorf("DeprecatedAt[%q] = %q, want %q", k, got.DeprecatedAt[k], v)
				}
			}

			// Compare Beta flag
			if got.Beta != tt.want.Beta {
				t.Errorf("Beta = %v, want %v", got.Beta, tt.want.Beta)
			}
		})
	}
}

func TestAvailabilityIsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		avail Availability
		want  bool
	}{
		{
			name:  "empty availability",
			avail: Availability{IntroducedAt: map[string]string{}, DeprecatedAt: map[string]string{}},
			want:  true,
		},
		{
			name:  "has introduced version",
			avail: Availability{IntroducedAt: map[string]string{"macOS": "10.14"}, DeprecatedAt: map[string]string{}},
			want:  false,
		},
		{
			name:  "has deprecated version",
			avail: Availability{IntroducedAt: map[string]string{}, DeprecatedAt: map[string]string{"macOS": "10.14"}},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.avail.IsEmpty()
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAvailabilityPlatforms(t *testing.T) {
	avail := Availability{
		IntroducedAt: map[string]string{
			"macOS":   "10.14",
			"iOS":     "12.0",
			"watchOS": "5.0",
		},
		DeprecatedAt: map[string]string{},
	}

	platforms := avail.Platforms()

	// Should be sorted
	want := []string{"iOS", "macOS", "watchOS"}
	if len(platforms) != len(want) {
		t.Fatalf("Platforms() length = %d, want %d", len(platforms), len(want))
	}

	for i, p := range platforms {
		if p != want[i] {
			t.Errorf("Platforms()[%d] = %q, want %q", i, p, want[i])
		}
	}
}

func TestFindMinimumMacOSVersion(t *testing.T) {
	tests := []struct {
		name  string
		funcs []*ParsedFunction
		want  string
	}{
		{
			name:  "no functions",
			funcs: []*ParsedFunction{},
			want:  "",
		},
		{
			name: "single function",
			funcs: []*ParsedFunction{
				{
					Name: "TestFunc",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "10.14"},
						DeprecatedAt: map[string]string{},
					},
				},
			},
			want: "10.14",
		},
		{
			name: "multiple functions - find minimum",
			funcs: []*ParsedFunction{
				{
					Name: "NewFunc",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "15.4"},
						DeprecatedAt: map[string]string{},
					},
				},
				{
					Name: "OldFunc",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "10.0"},
						DeprecatedAt: map[string]string{},
					},
				},
				{
					Name: "MidFunc",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "10.14"},
						DeprecatedAt: map[string]string{},
					},
				},
			},
			want: "10.0",
		},
		{
			name: "no macOS versions",
			funcs: []*ParsedFunction{
				{
					Name: "IOSOnlyFunc",
					Availability: Availability{
						IntroducedAt: map[string]string{"iOS": "12.0"},
						DeprecatedAt: map[string]string{},
					},
				},
			},
			want: "",
		},
		{
			name: "version 10.10 vs 10.9",
			funcs: []*ParsedFunction{
				{
					Name: "Func1",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "10.10"},
						DeprecatedAt: map[string]string{},
					},
				},
				{
					Name: "Func2",
					Availability: Availability{
						IntroducedAt: map[string]string{"macOS": "10.9"},
						DeprecatedAt: map[string]string{},
					},
				},
			},
			want: "10.9", // Should use semantic comparison
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinimumMacOSVersion(tt.funcs)
			if got != tt.want {
				t.Errorf("findMinimumMacOSVersion() = %q, want %q", got, tt.want)
			}
		})
	}
}
