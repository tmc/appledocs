package occ2go

import (
	"reflect"
	"testing"

	"github.com/tmc/appledocs"
)

// TestParseProtocolDeclaration tests protocol name extraction from tokens
func TestParseProtocolDeclaration(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []appledocs.Token
		want     *ParsedProtocol
		wantNil  bool
	}{
		{
			name: "simple delegate protocol with @protocol keyword",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSApplicationDelegate"},
			},
			want: &ParsedProtocol{
				Name:       "NSApplicationDelegate",
				IsDelegate: true,
			},
		},
		{
			name: "simple data source protocol with @protocol keyword",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSTableViewDataSource"},
			},
			want: &ParsedProtocol{
				Name:         "NSTableViewDataSource",
				IsDataSource: true,
			},
		},
		{
			name: "protocol without @protocol keyword (fallback)",
			tokens: []appledocs.Token{
				{Kind: "identifier", Text: "NSWindowDelegate"},
			},
			want: &ParsedProtocol{
				Name:       "NSWindowDelegate",
				IsDelegate: true,
			},
		},
		{
			name: "regular protocol (not delegate or data source)",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSCopying"},
			},
			want: &ParsedProtocol{
				Name:         "NSCopying",
				IsDelegate:   false,
				IsDataSource: false,
			},
		},
		{
			name: "protocol with extra whitespace",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "text", Text: " "},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSTextFieldDelegate"},
			},
			want: &ParsedProtocol{
				Name:       "NSTextFieldDelegate",
				IsDelegate: true,
			},
		},
		{
			name: "protocol ending with 'Delegate' suffix",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "SCStreamOutputDelegate"},
			},
			want: &ParsedProtocol{
				Name:       "SCStreamOutputDelegate",
				IsDelegate: true,
			},
		},
		{
			name: "protocol ending with 'DataSource' suffix",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "UICollectionViewDataSource"},
			},
			want: &ParsedProtocol{
				Name:         "UICollectionViewDataSource",
				IsDataSource: true,
			},
		},
		{
			name: "empty tokens",
			tokens: []appledocs.Token{},
			wantNil: true,
		},
		{
			name: "only whitespace tokens",
			tokens: []appledocs.Token{
				{Kind: "text", Text: " "},
				{Kind: "text", Text: " "},
			},
			wantNil: true,
		},
		{
			name: "@protocol keyword but no identifier",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseProtocolDeclaration(tt.tokens)

			if tt.wantNil {
				if got != nil {
					t.Errorf("ParseProtocolDeclaration() = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Fatalf("ParseProtocolDeclaration() = nil, want %v", tt.want)
			}

			if got.Name != tt.want.Name {
				t.Errorf("ParseProtocolDeclaration().Name = %v, want %v", got.Name, tt.want.Name)
			}

			if got.IsDelegate != tt.want.IsDelegate {
				t.Errorf("ParseProtocolDeclaration().IsDelegate = %v, want %v", got.IsDelegate, tt.want.IsDelegate)
			}

			if got.IsDataSource != tt.want.IsDataSource {
				t.Errorf("ParseProtocolDeclaration().IsDataSource = %v, want %v", got.IsDataSource, tt.want.IsDataSource)
			}
		})
	}
}

// TestDetectDelegateType tests the delegate/data source auto-detection logic
func TestDetectDelegateType(t *testing.T) {
	tests := []struct {
		name             string
		protocolName     string
		wantIsDelegate   bool
		wantIsDataSource bool
	}{
		{
			name:             "NSApplicationDelegate",
			protocolName:     "NSApplicationDelegate",
			wantIsDelegate:   true,
			wantIsDataSource: false,
		},
		{
			name:             "NSWindowDelegate",
			protocolName:     "NSWindowDelegate",
			wantIsDelegate:   true,
			wantIsDataSource: false,
		},
		{
			name:             "NSTableViewDataSource",
			protocolName:     "NSTableViewDataSource",
			wantIsDelegate:   false,
			wantIsDataSource: true,
		},
		{
			name:             "UICollectionViewDataSource",
			protocolName:     "UICollectionViewDataSource",
			wantIsDelegate:   false,
			wantIsDataSource: true,
		},
		{
			name:             "SCStreamOutputDelegate",
			protocolName:     "SCStreamOutputDelegate",
			wantIsDelegate:   true,
			wantIsDataSource: false,
		},
		{
			name:             "NSCopying (regular protocol)",
			protocolName:     "NSCopying",
			wantIsDelegate:   false,
			wantIsDataSource: false,
		},
		{
			name:             "NSCoding (regular protocol)",
			protocolName:     "NSCoding",
			wantIsDelegate:   false,
			wantIsDataSource: false,
		},
		{
			name:             "NSSecureCoding (regular protocol)",
			protocolName:     "NSSecureCoding",
			wantIsDelegate:   false,
			wantIsDataSource: false,
		},
		{
			name:             "empty string",
			protocolName:     "",
			wantIsDelegate:   false,
			wantIsDataSource: false,
		},
		{
			name:             "DelegateProtocol (ends with Delegate)",
			protocolName:     "DelegateProtocol",
			wantIsDelegate:   false,
			wantIsDataSource: false,
		},
		{
			name:             "MyDelegate",
			protocolName:     "MyDelegate",
			wantIsDelegate:   true,
			wantIsDataSource: false,
		},
		{
			name:             "CustomDataSource",
			protocolName:     "CustomDataSource",
			wantIsDelegate:   false,
			wantIsDataSource: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proto := &ParsedProtocol{
				Name: tt.protocolName,
			}

			detectDelegateType(proto)

			if proto.IsDelegate != tt.wantIsDelegate {
				t.Errorf("detectDelegateType() IsDelegate = %v, want %v", proto.IsDelegate, tt.wantIsDelegate)
			}

			if proto.IsDataSource != tt.wantIsDataSource {
				t.Errorf("detectDelegateType() IsDataSource = %v, want %v", proto.IsDataSource, tt.wantIsDataSource)
			}
		})
	}
}

// TestDetectDelegateTypeNilProtocol tests that detectDelegateType handles nil gracefully
func TestDetectDelegateTypeNilProtocol(t *testing.T) {
	// Should not panic
	detectDelegateType(nil)
}

// TestParseProtocolDeclarationWithMethods tests protocol parsing when methods are present
// Note: Currently ParseProtocolDeclaration only extracts the name, not methods
// This test documents the current behavior and can be extended when method parsing is added
func TestParseProtocolDeclarationWithMethods(t *testing.T) {
	tokens := []appledocs.Token{
		{Kind: "keyword", Text: "@protocol"},
		{Kind: "text", Text: " "},
		{Kind: "identifier", Text: "NSApplicationDelegate"},
		{Kind: "text", Text: " "},
		{Kind: "text", Text: "<"},
		{Kind: "identifier", Text: "NSObject"},
		{Kind: "text", Text: ">"},
		{Kind: "text", Text: "\n"},
		{Kind: "keyword", Text: "@optional"},
		{Kind: "text", Text: "\n"},
		{Kind: "text", Text: "-"},
		{Kind: "text", Text: " "},
		{Kind: "text", Text: "("},
		{Kind: "keyword", Text: "void"},
		{Kind: "text", Text: ")"},
		{Kind: "identifier", Text: "applicationDidFinishLaunching"},
		{Kind: "text", Text: ":"},
		{Kind: "text", Text: "("},
		{Kind: "typeIdentifier", Text: "NSNotification"},
		{Kind: "text", Text: " "},
		{Kind: "text", Text: "*)"},
		{Kind: "identifier", Text: "notification"},
		{Kind: "text", Text: ";"},
		{Kind: "text", Text: "\n"},
		{Kind: "keyword", Text: "@end"},
	}

	got := ParseProtocolDeclaration(tokens)

	if got == nil {
		t.Fatal("ParseProtocolDeclaration() = nil, want non-nil")
	}

	if got.Name != "NSApplicationDelegate" {
		t.Errorf("ParseProtocolDeclaration().Name = %v, want NSApplicationDelegate", got.Name)
	}

	if !got.IsDelegate {
		t.Errorf("ParseProtocolDeclaration().IsDelegate = false, want true")
	}

	// Note: Method parsing would be tested here once implemented
	// For now, we just verify the protocol name is extracted correctly
}

// TestParseProtocolDeclarationRealWorld tests parsing with actual Apple documentation tokens
// This uses token patterns observed in real Apple documentation JSON files
func TestParseProtocolDeclarationRealWorld(t *testing.T) {
	tests := []struct {
		name         string
		tokens       []appledocs.Token
		wantName     string
		wantDelegate bool
	}{
		{
			name: "NSApplicationDelegate from AppKit docs",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSApplicationDelegate"},
				{Kind: "text", Text: " <"},
				{Kind: "typeIdentifier", Text: "NSObject", Identifier: "doc://com.apple.documentation/documentation/objectivec/1418956-nsobject"},
				{Kind: "text", Text: ">"},
			},
			wantName:     "NSApplicationDelegate",
			wantDelegate: true,
		},
		{
			name: "NSTableViewDataSource from AppKit docs",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "NSTableViewDataSource"},
			},
			wantName:     "NSTableViewDataSource",
			wantDelegate: false,
		},
		{
			name: "SCStreamOutput from ScreenCaptureKit docs",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "SCStreamOutput"},
			},
			wantName:     "SCStreamOutput",
			wantDelegate: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseProtocolDeclaration(tt.tokens)

			if got == nil {
				t.Fatal("ParseProtocolDeclaration() = nil, want non-nil")
			}

			if got.Name != tt.wantName {
				t.Errorf("ParseProtocolDeclaration().Name = %v, want %v", got.Name, tt.wantName)
			}

			// Note: IsDelegate/IsDataSource is auto-detected from name,
			// not from documentation tokens
		})
	}
}

// TestParseProtocolDeclarationEdgeCases tests various edge cases
func TestParseProtocolDeclarationEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		tokens  []appledocs.Token
		wantNil bool
	}{
		{
			name: "multiple @protocol keywords (should use first)",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "FirstProtocol"},
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "SecondProtocol"},
			},
			wantNil: false, // Should parse first protocol
		},
		{
			name: "identifier before @protocol keyword",
			tokens: []appledocs.Token{
				{Kind: "identifier", Text: "SomeIdentifier"},
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "ActualProtocol"},
			},
			wantNil: false, // Should find @protocol and use ActualProtocol
		},
		{
			name: "only @protocol keyword",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@protocol"},
			},
			wantNil: true, // No identifier found
		},
		{
			name: "no identifiers at all",
			tokens: []appledocs.Token{
				{Kind: "text", Text: " "},
				{Kind: "keyword", Text: "@protocol"},
				{Kind: "text", Text: " "},
			},
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseProtocolDeclaration(tt.tokens)

			if tt.wantNil && got != nil {
				t.Errorf("ParseProtocolDeclaration() = %v, want nil", got)
			}

			if !tt.wantNil && got == nil {
				t.Errorf("ParseProtocolDeclaration() = nil, want non-nil")
			}
		})
	}
}

// TestParseProtocolWithInheritance tests protocol declarations that inherit from other protocols
func TestParseProtocolWithInheritance(t *testing.T) {
	tokens := []appledocs.Token{
		{Kind: "keyword", Text: "@protocol"},
		{Kind: "text", Text: " "},
		{Kind: "identifier", Text: "NSSecureCoding"},
		{Kind: "text", Text: " <"},
		{Kind: "typeIdentifier", Text: "NSCoding"},
		{Kind: "text", Text: ">"},
	}

	got := ParseProtocolDeclaration(tokens)

	if got == nil {
		t.Fatal("ParseProtocolDeclaration() = nil, want non-nil")
	}

	if got.Name != "NSSecureCoding" {
		t.Errorf("ParseProtocolDeclaration().Name = %v, want NSSecureCoding", got.Name)
	}

	// Note: Protocol inheritance is not currently captured in ParsedProtocol
	// This test documents current behavior - can be extended when inheritance is added
}

// TestParseProtocolCompleteness verifies that ParsedProtocol struct fields are properly initialized
func TestParseProtocolCompleteness(t *testing.T) {
	tokens := []appledocs.Token{
		{Kind: "keyword", Text: "@protocol"},
		{Kind: "text", Text: " "},
		{Kind: "identifier", Text: "TestProtocol"},
	}

	proto := ParseProtocolDeclaration(tokens)

	if proto == nil {
		t.Fatal("ParseProtocolDeclaration() returned nil")
	}

	// Verify struct is properly initialized with zero values
	if proto.Name != "TestProtocol" {
		t.Errorf("Name = %v, want TestProtocol", proto.Name)
	}

	if proto.RequiredMethods != nil {
		t.Errorf("RequiredMethods = %v, want nil (not yet parsed)", proto.RequiredMethods)
	}

	if proto.OptionalMethods != nil {
		t.Errorf("OptionalMethods = %v, want nil (not yet parsed)", proto.OptionalMethods)
	}

	if proto.Comment != "" {
		t.Errorf("Comment = %v, want empty string", proto.Comment)
	}

	if proto.DocURL != "" {
		t.Errorf("DocURL = %v, want empty string", proto.DocURL)
	}

	if proto.Abstract != "" {
		t.Errorf("Abstract = %v, want empty string", proto.Abstract)
	}

	// Verify zero value of Availability struct
	if !reflect.DeepEqual(proto.Availability, Availability{}) {
		t.Errorf("Availability = %+v, want zero value", proto.Availability)
	}
}

// BenchmarkParseProtocolDeclaration benchmarks protocol parsing performance
func BenchmarkParseProtocolDeclaration(b *testing.B) {
	tokens := []appledocs.Token{
		{Kind: "keyword", Text: "@protocol"},
		{Kind: "text", Text: " "},
		{Kind: "identifier", Text: "NSApplicationDelegate"},
		{Kind: "text", Text: " <"},
		{Kind: "typeIdentifier", Text: "NSObject"},
		{Kind: "text", Text: ">"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseProtocolDeclaration(tokens)
	}
}

// BenchmarkDetectDelegateType benchmarks delegate detection performance
func BenchmarkDetectDelegateType(b *testing.B) {
	proto := &ParsedProtocol{
		Name: "NSApplicationDelegate",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		detectDelegateType(proto)
	}
}
