package occ2go

import (
	"testing"
	"github.com/tmc/appledocs"
)

func TestParseStructFieldWithVoidPtr(t *testing.T) {
	doc := &appledocs.Document{
		Metadata: appledocs.Metadata{
			ExternalID: "c:@SA@TestStruct@FI@info",
		},
		PrimaryContentSections: []appledocs.ContentSection{
			{
				Declarations: []appledocs.Declaration{
					{
						Tokens: []appledocs.Token{
							{Kind: "keyword", Text: "void"},
							{Kind: "text", Text: "*"},
							{Kind: "identifier", Text: "info"},
							{Kind: "text", Text: ";"},
						},
					},
				},
			},
		},
	}

	field, err := ParseStructField(doc)
	if err != nil {
		t.Fatalf("ParseStructField failed: %v", err)
	}

	if field.Name != "Info" {
		t.Errorf("Expected field name 'Info', got '%s'", field.Name)
	}

	if field.Type != "void *" {
		t.Errorf("Expected field type 'void *', got '%s'", field.Type)
	}
}
