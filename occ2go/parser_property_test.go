package occ2go

import (
	"strings"
	"testing"

	"github.com/tmc/appledocs"
)

// TestParsePropertyDeclaration tests property parsing with various type patterns
func TestParsePropertyDeclaration(t *testing.T) {
	tests := []struct {
		name       string
		tokens     []appledocs.Token
		wantType   string
		wantName   string
		wantErr    bool
	}{
		{
			name: "simple property",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@property"},
				{Kind: "text", Text: " ("},
				{Kind: "keyword", Text: "readonly"},
				{Kind: "text", Text: ") "},
				{Kind: "typeIdentifier", Text: "NSString"},
				{Kind: "text", Text: " * "},
				{Kind: "identifier", Text: "title"},
				{Kind: "text", Text: ";"},
			},
			wantType: "NSString *",
			wantName: "title",
		},
		{
			name: "generic array property",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@property"},
				{Kind: "text", Text: " ("},
				{Kind: "keyword", Text: "readonly"},
				{Kind: "text", Text: ", "},
				{Kind: "keyword", Text: "copy"},
				{Kind: "text", Text: ") "},
				{Kind: "typeIdentifier", Text: "NSArray"},
				{Kind: "text", Text: "<NSButton *> * "},
				{Kind: "identifier", Text: "buttons"},
				{Kind: "text", Text: ";"},
			},
			wantType: "NSArray <NSButton *> *",
			wantName: "buttons",
		},
		{
			name: "generic array with __kindof",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@property"},
				{Kind: "text", Text: " ("},
				{Kind: "keyword", Text: "readonly"},
				{Kind: "text", Text: ") "},
				{Kind: "typeIdentifier", Text: "NSArray"},
				{Kind: "text", Text: "<__kindof NSView *> * "},
				{Kind: "identifier", Text: "subviews"},
				{Kind: "text", Text: ";"},
			},
			wantType: "NSArray <__kindof NSView *> *",
			wantName: "subviews",
		},
		{
			name: "generic dictionary property",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@property"},
				{Kind: "text", Text: " ("},
				{Kind: "keyword", Text: "copy"},
				{Kind: "text", Text: ") "},
				{Kind: "typeIdentifier", Text: "NSDictionary"},
				{Kind: "text", Text: "<NSString *, id> * "},
				{Kind: "identifier", Text: "userInfo"},
				{Kind: "text", Text: ";"},
			},
			wantType: "NSDictionary <NSString *, id> *",
			wantName: "userInfo",
		},
		{
			name: "non-pointer property",
			tokens: []appledocs.Token{
				{Kind: "keyword", Text: "@property"},
				{Kind: "text", Text: " "},
				{Kind: "keyword", Text: "NSInteger"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "count"},
				{Kind: "text", Text: ";"},
			},
			wantType: "NSInteger",
			wantName: "count",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prop, err := ParsePropertyDeclaration(tt.tokens)
			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			if prop == nil {
				t.Fatal("Expected property, got nil")
			}

			// Normalize whitespace for comparison
			gotType := strings.Join(strings.Fields(prop.Type), " ")
			wantType := strings.Join(strings.Fields(tt.wantType), " ")

			if gotType != wantType {
				t.Errorf("Type mismatch:\n  got:  %q\n  want: %q", gotType, wantType)
			}
			if prop.Name != tt.wantName {
				t.Errorf("Name mismatch: got %q, want %q", prop.Name, tt.wantName)
			}
			// ObjCType should match Type
			gotObjCType := strings.Join(strings.Fields(prop.ObjCType), " ")
			if gotObjCType != gotType {
				t.Errorf("ObjCType should match Type:\n  ObjCType: %q\n  Type:     %q", gotObjCType, gotType)
			}
		})
	}
}
