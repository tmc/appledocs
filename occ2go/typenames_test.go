package occ2go

import "testing"

func TestStripTypeQualifiers(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"__kindof NSView *", "NSView *"},
		{"const char *", "char *"},
		{"NSString *", "NSString *"},
		{"[]__kindof AVCaptureControl", "[]AVCaptureControl"},
		{"__kindof NSView", "NSView"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := StripTypeQualifiers(tt.input)
			if got != tt.want {
				t.Errorf("StripTypeQualifiers(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPointerType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"NSString *", true},
		{"NSView *", true},
		{"int", false},
		{"id", false},
		{"BOOL", false},
		{"char *", true},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsPointerType(tt.input)
			if got != tt.want {
				t.Errorf("IsPointerType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestStripPointer(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSString *", "NSString"},
		{"NSView *", "NSView"},
		{"int", "int"},
		{"char *", "char"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := StripPointer(tt.input)
			if got != tt.want {
				t.Errorf("StripPointer(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsBlockType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"void (^)(void)", true},
		{"void (^)(NSInteger)", true},
		{"NSString *", false},
		{"int", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsBlockType(tt.input)
			if got != tt.want {
				t.Errorf("IsBlockType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsProtocolType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"id<NSCopying>", true},
		{"id<NSFetchRequestResult>", true},
		{"id", false},
		{"NSString", false},
		{"NSView<Protocol>", false}, // Not id<Protocol> pattern
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsProtocolType(tt.input)
			if got != tt.want {
				t.Errorf("IsProtocolType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsArrayType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"[]int", true},
		{"[]NSString", true},
		{"[]unsafe.Pointer", true},
		{"NSArray<NSString *>", false},
		{"int", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsArrayType(tt.input)
			if got != tt.want {
				t.Errorf("IsArrayType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetArrayElementType(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"[]int", "int"},
		{"[]NSString", "NSString"},
		{"[]unsafe.Pointer", "unsafe.Pointer"},
		{"int", ""},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetArrayElementType(tt.input)
			if got != tt.want {
				t.Errorf("GetArrayElementType(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsGenericType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"NSArray<NSString *>", true},
		{"NSDictionary<NSString *, id>", true},
		{"NSSet<NSNumber *>", true},
		{"NSString", false},
		{"int", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsGenericType(tt.input)
			if got != tt.want {
				t.Errorf("IsGenericType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestExtractGenericElementType(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSArray<NSString *>", "NSString"},
		{"NSArray<NSString *> *", "NSString"},
		{"NSArray<NSButton *>", "NSButton"},
		{"NSArray<__kindof NSView *>", "NSView"},
		{"NSArray<NSView<NSCollectionViewElement> *>", "NSView"},
		{"NSDictionary<K, V>", ""}, // Not supported
		{"NSString", ""},           // Not an NSArray
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ExtractGenericElementType(tt.input)
			if got != tt.want {
				t.Errorf("ExtractGenericElementType(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestStripProtocolConformance(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSView<NSCollectionViewElement>", "NSView"},
		{"NSString<NSCopying>", "NSString"},
		{"id<NSCopying>", "id<NSCopying>"}, // Special case
		{"NSString", "NSString"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := StripProtocolConformance(tt.input)
			if got != tt.want {
				t.Errorf("StripProtocolConformance(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestStripPackageQualification(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"foundation.NSString", "NSString"},
		{"uniformtypeidentifiers.UTType", "UTType"},
		{"appkit.NSView", "NSView"},
		{"NSString", "NSString"},
		{"NSView.Property", "NSView.Property"}, // Not a package (Property is capitalized)
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := StripPackageQualification(tt.input)
			if got != tt.want {
				t.Errorf("StripPackageQualification(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNormalizeTypeName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"__kindof NSView *", "NSView *"},
		{"foundation.NSString *", "NSString *"},
		{"const char *", "char *"},
		{"NSView<Protocol> *", "NSView *"}, // Strips protocol but keeps pointer
		{"__kindof foundation.NSButton *", "NSButton *"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := NormalizeTypeName(tt.input)
			if got != tt.want {
				t.Errorf("NormalizeTypeName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPrimitiveType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"int", true},
		{"BOOL", true},
		{"bool", true},
		{"NSInteger", true},
		{"NSUInteger", true},
		{"CGFloat", true},
		{"double", true},
		{"float", true},
		{"uint64_t", true},
		{"NSString", false},
		{"id", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsPrimitiveType(tt.input)
			if got != tt.want {
				t.Errorf("IsPrimitiveType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsSpecialType(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"id", true},
		{"Class", true},
		{"SEL", true},
		{"NSString", false},
		{"int", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsSpecialType(tt.input)
			if got != tt.want {
				t.Errorf("IsSpecialType(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
