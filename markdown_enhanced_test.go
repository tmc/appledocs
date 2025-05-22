package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestComplexMarkdownContent tests markdown generation with complex content structures
func TestComplexMarkdownContent(t *testing.T) {
	tests := []struct {
		name           string
		doc            *DocJSONData
		expectedInText []string
		notInText      []string
	}{
		{
			name: "full documentation with all sections",
			doc: &DocJSONData{
				Metadata: Metadata{
					Title:       "CompleteFramework",
					Role:        "framework",
					RoleHeading: "Framework",
					Modules: []Module{
						{Name: "CompleteFramework"},
					},
					Platforms: []Platform{
						{Name: "iOS", IntroducedAt: "13.0", Beta: false},
						{Name: "macOS", IntroducedAt: "10.15", Beta: true},
						{Name: "watchOS", IntroducedAt: "6.0", Deprecated: true},
					},
				},
				Abstract: []TextContent{
					{Type: "text", Text: "A comprehensive framework for testing markdown generation."},
				},
				PrimaryContentSections: []ContentSection{
					{
						Kind: "declarations",
						Declarations: []Declaration{
							{
								Languages: []string{"swift"},
								Tokens: []Fragment{
									{Text: "import "},
									{Text: "CompleteFramework"},
								},
							},
							{
								Languages: []string{"objective-c"},
								Tokens: []Fragment{
									{Text: "@import "},
									{Text: "CompleteFramework"},
									{Text: ";"},
								},
							},
						},
					},
					{
						Kind: "content",
						Content: []ContentBlock{
							{
								Type:  "heading",
								Level: 2,
								Text:  "Overview",
							},
							{
								Type: "paragraph",
								InlineContent: []InlineContent{
									{Type: "text", Text: "This framework provides "},
									{Type: "emphasis", InlineContent: []InlineContent{
										{Type: "text", Text: "powerful"},
									}},
									{Type: "text", Text: " capabilities for "},
									{Type: "strong", InlineContent: []InlineContent{
										{Type: "text", Text: "advanced development"},
									}},
									{Type: "text", Text: "."},
								},
							},
							{
								Type: "unorderedList",
								Items: []Item{
									{
										Content: []ContentBlock{
											{
												Type: "paragraph",
												InlineContent: []InlineContent{
													{Type: "text", Text: "Feature one"},
												},
											},
										},
									},
									{
										Content: []ContentBlock{
											{
												Type: "paragraph",
												InlineContent: []InlineContent{
													{Type: "text", Text: "Feature two with "},
													{Type: "codeVoice", Code: "inline code"},
												},
											},
										},
									},
								},
							},
							{
								Type: "orderedList",
								Items: []Item{
									{
										Content: []ContentBlock{
											{
												Type: "paragraph",
												InlineContent: []InlineContent{
													{Type: "text", Text: "First step"},
												},
											},
										},
									},
									{
										Content: []ContentBlock{
											{
												Type: "paragraph",
												InlineContent: []InlineContent{
													{Type: "text", Text: "Second step"},
												},
											},
										},
									},
								},
							},
							{
								Type: "aside",
								Name: "Important",
								Content: []ContentBlock{
									{
										Type: "paragraph",
										InlineContent: []InlineContent{
											{Type: "text", Text: "This is an important note about the framework."},
										},
									},
								},
							},
							{
								Type:  "codeListing",
								Style: "swift",
								Content: []ContentBlock{
									{
										Type: "paragraph",
										InlineContent: []InlineContent{
											{Type: "text", Text: "let framework = CompleteFramework()"},
										},
									},
									{
										Type: "paragraph",
										InlineContent: []InlineContent{
											{Type: "text", Text: "framework.initialize()"},
										},
									},
								},
							},
						},
					},
				},
				TopicSections: []TopicSection{
					{
						Title: "Core Classes",
						Identifiers: []string{
							"doc://com.apple.documentation/documentation/CompleteFramework/MainClass",
							"doc://com.apple.documentation/documentation/CompleteFramework/HelperClass",
						},
					},
					{
						Title: "Protocols",
						Identifiers: []string{
							"doc://com.apple.documentation/documentation/CompleteFramework/MainProtocol",
						},
					},
				},
				RelationshipsSections: []RelationshipSection{
					{
						Title: "Inherits From",
						Type:  "inheritsFrom",
						Identifiers: []string{
							"doc://com.apple.documentation/documentation/Foundation/NSObject",
						},
					},
				},
				SeeAlsoSections: []SeeAlsoSection{
					{
						Title: "Related Documentation",
						Identifiers: []string{
							"doc://com.apple.documentation/documentation/RelatedFramework",
						},
					},
				},
				References: map[string]Reference{
					"doc://com.apple.documentation/documentation/CompleteFramework/MainClass": {
						Title: "MainClass",
						URL:   "doc://com.apple.documentation/documentation/CompleteFramework/MainClass",
						Role:  "class",
						Abstract: []TextContent{
							{Type: "text", Text: "The main class of the framework."},
						},
					},
					"doc://com.apple.documentation/documentation/CompleteFramework/HelperClass": {
						Title: "HelperClass",
						URL:   "doc://com.apple.documentation/documentation/CompleteFramework/HelperClass",
						Role:  "class",
						Abstract: []TextContent{
							{Type: "text", Text: "A helper class that provides utility functions."},
						},
					},
					"doc://com.apple.documentation/documentation/CompleteFramework/MainProtocol": {
						Title: "MainProtocol",
						URL:   "doc://com.apple.documentation/documentation/CompleteFramework/MainProtocol",
						Role:  "protocol",
						Abstract: []TextContent{
							{Type: "text", Text: "The main protocol definition."},
						},
					},
					"doc://com.apple.documentation/documentation/Foundation/NSObject": {
						Title: "NSObject",
						URL:   "doc://com.apple.documentation/documentation/Foundation/NSObject",
						Role:  "class",
					},
					"doc://com.apple.documentation/documentation/RelatedFramework": {
						Title: "RelatedFramework",
						URL:   "doc://com.apple.documentation/documentation/RelatedFramework",
						Role:  "framework",
						Abstract: []TextContent{
							{Type: "text", Text: "A related framework for extended functionality."},
						},
					},
				},
			},
			expectedInText: []string{
				"# CompleteFramework",
				"---",
				"**Type:** `Framework`",
				"**Framework:** `CompleteFramework`",
				"**Platform Availability:**",
				"✅iOS **13.0**",
				"🧪macOS **10.15** (Beta)",
				"⚠️watchOS **6.0** (Deprecated)",
				"## Overview",
				"A comprehensive framework for testing markdown generation.",
				"## Declaration",
				"### Swift",
				"```swift",
				"import CompleteFramework",
				"```",
				"### Objective-C",
				"```objectivec",
				"@import CompleteFramework;",
				"```",
				"*powerful*",
				"**advanced development**",
				"- Feature one",
				"- Feature two with `inline code`",
				"1. First step",
				"2. Second step",
				"> **Important**",
				"> This is an important note",
				"```swift",
				"let framework = CompleteFramework()",
				"framework.initialize()",
				"## Topics",
				"### Core Classes",
				"**[MainClass](/documentation/documentation/CompleteFramework/MainClass.md)** 🏛️",
				"The main class of the framework.",
				"**[HelperClass](/documentation/documentation/CompleteFramework/HelperClass.md)** 🏛️",
				"A helper class that provides utility functions.",
				"### Protocols",
				"**[MainProtocol](/documentation/documentation/CompleteFramework/MainProtocol.md)** 📋",
				"The main protocol definition.",
				"## Relationships",
				"### Inherits From",
				"- [NSObject](/documentation/documentation/Foundation/NSObject.md)",
				"## See Also",
				"### Related Documentation",
				"- [RelatedFramework](/documentation/documentation/RelatedFramework.md) - A related framework for extended functionality.",
			},
			notInText: []string{
				"null",
				"undefined",
				"<script>",
				"<html>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			err := writeMarkdownContent(&buf, tt.doc)
			if err != nil {
				t.Errorf("writeMarkdownContent() error = %v", err)
			}

			result := buf.String()

			// Check for expected content
			for _, expected := range tt.expectedInText {
				if !strings.Contains(result, expected) {
					t.Errorf("Expected text '%s' not found in output", expected)
				}
			}

			// Check for content that should not be present
			for _, notExpected := range tt.notInText {
				if strings.Contains(result, notExpected) {
					t.Errorf("Unexpected text '%s' found in output", notExpected)
				}
			}

			// Verify proper markdown structure
			lines := strings.Split(result, "\n")
			hasTitle := false
			for _, line := range lines {
				if strings.HasPrefix(line, "# ") {
					hasTitle = true
					break
				}
			}
			if !hasTitle {
				t.Error("Markdown should have a title (line starting with '# ')")
			}
		})
	}
}

// TestMarkdownTables tests table generation in markdown
func TestMarkdownTables(t *testing.T) {
	doc := &DocJSONData{
		Metadata: Metadata{Title: "TableTest"},
		PrimaryContentSections: []ContentSection{
			{
				Kind: "content",
				Content: []ContentBlock{
					{
						Type: "table",
						Content: []ContentBlock{
							{
								Type: "tableRow",
								Content: []ContentBlock{
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "text", Text: "Property"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "text", Text: "Type"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "text", Text: "Description"},
										},
									},
								},
							},
							{
								Type: "tableRow",
								Content: []ContentBlock{
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "codeVoice", Code: "name"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "codeVoice", Code: "String"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "text", Text: "The name of the object"},
										},
									},
								},
							},
							{
								Type: "tableRow",
								Content: []ContentBlock{
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "codeVoice", Code: "isEnabled"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "codeVoice", Code: "Bool"},
										},
									},
									{
										Type: "tableCell",
										InlineContent: []InlineContent{
											{Type: "text", Text: "Whether the object is enabled"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	var buf strings.Builder
	err := writeMarkdownContent(&buf, doc)
	if err != nil {
		t.Errorf("writeMarkdownContent() error = %v", err)
	}

	result := buf.String()

	expectedTableElements := []string{
		"|  Property | Type | Description |",
		"| --- | --- | --- |",
		"|  `name` | `String` | The name of the object |",
		"|  `isEnabled` | `Bool` | Whether the object is enabled |",
	}

	for _, expected := range expectedTableElements {
		if !strings.Contains(result, expected) {
			t.Errorf("Expected table element '%s' not found in output", expected)
		}
	}
}

// TestMarkdownCodeBlocks tests code block generation with different languages
func TestMarkdownCodeBlocks(t *testing.T) {
	tests := []struct {
		name         string
		style        string
		content      string
		expectedLang string
	}{
		{
			name:         "Swift code",
			style:        "swift",
			content:      "let value = 42",
			expectedLang: "swift",
		},
		{
			name:         "Objective-C code",
			style:        "objective-c",
			content:      "NSString *value = @\"test\";",
			expectedLang: "objectivec",
		},
		{
			name:         "C++ code",
			style:        "c++",
			content:      "std::string value = \"test\";",
			expectedLang: "cpp",
		},
		{
			name:         "Shell command",
			style:        "bash",
			content:      "echo \"Hello World\"",
			expectedLang: "bash",
		},
		{
			name:         "No style specified",
			style:        "",
			content:      "generic code",
			expectedLang: "swift", // default
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := &DocJSONData{
				Metadata: Metadata{Title: "CodeTest"},
				PrimaryContentSections: []ContentSection{
					{
						Kind: "content",
						Content: []ContentBlock{
							{
								Type:  "codeListing",
								Style: tt.style,
								Content: []ContentBlock{
									{
										Type: "paragraph",
										InlineContent: []InlineContent{
											{Type: "text", Text: tt.content},
										},
									},
								},
							},
						},
					},
				},
			}

			var buf strings.Builder
			err := writeMarkdownContent(&buf, doc)
			if err != nil {
				t.Errorf("writeMarkdownContent() error = %v", err)
			}

			result := buf.String()

			expectedStart := "```" + tt.expectedLang
			expectedEnd := "```"

			if !strings.Contains(result, expectedStart) {
				t.Errorf("Expected code block start '%s' not found", expectedStart)
			}
			if !strings.Contains(result, tt.content) {
				t.Errorf("Expected code content '%s' not found", tt.content)
			}
			if !strings.Contains(result, expectedEnd) {
				t.Errorf("Expected code block end '%s' not found", expectedEnd)
			}
		})
	}
}

// TestMarkdownLinks tests link generation and formatting
func TestMarkdownLinks(t *testing.T) {
	tests := []struct {
		name         string
		inline       InlineContent
		refs         map[string]Reference
		expectedLink string
	}{
		{
			name: "reference link with abstract",
			inline: InlineContent{
				Type:       "reference",
				Identifier: "doc://com.apple.documentation/documentation/SwiftUI/View",
			},
			refs: map[string]Reference{
				"doc://com.apple.documentation/documentation/SwiftUI/View": {
					Title: "View",
					URL:   "doc://com.apple.documentation/documentation/SwiftUI/View",
					Abstract: []TextContent{
						{Type: "text", Text: "A type that represents part of your app's user interface."},
					},
				},
			},
			expectedLink: "[View](/documentation/documentation/SwiftUI/View.md)",
		},
		{
			name: "reference link with beta indicator",
			inline: InlineContent{
				Type:       "reference",
				Identifier: "doc://com.apple.documentation/documentation/SwiftUI/BetaView",
			},
			refs: map[string]Reference{
				"doc://com.apple.documentation/documentation/SwiftUI/BetaView": {
					Title: "BetaView",
					URL:   "doc://com.apple.documentation/documentation/SwiftUI/BetaView",
					Beta:  true,
				},
			},
			expectedLink: "[BetaView (Beta)](/documentation/documentation/SwiftUI/BetaView.md)",
		},
		{
			name: "missing reference fallback",
			inline: InlineContent{
				Type:       "reference",
				Identifier: "doc://com.apple.documentation/documentation/Missing/Class",
			},
			refs:         map[string]Reference{},
			expectedLink: "[Class](/documentation/documentation/Missing/Class.md)",
		},
		{
			name: "external link",
			inline: InlineContent{
				Type:       "link",
				Identifier: "https://developer.apple.com",
				InlineContent: []InlineContent{
					{Type: "text", Text: "Apple Developer"},
				},
			},
			refs:         map[string]Reference{},
			expectedLink: "[Apple Developer](https://developer.apple.com)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			writeInlineContent(&buf, tt.inline, tt.refs)
			result := buf.String()

			if result != tt.expectedLink {
				t.Errorf("Expected link '%s', got '%s'", tt.expectedLink, result)
			}
		})
	}
}

// TestMarkdownImageHandling tests image embedding in markdown
func TestMarkdownImageHandling(t *testing.T) {
	inline := InlineContent{
		Type:       "image",
		Identifier: "test-image",
	}

	refs := map[string]Reference{
		"test-image": {
			Title: "Test Image",
			URL:   "/documentation/images/test.png",
			Abstract: []TextContent{
				{Type: "text", Text: "A test image for documentation"},
			},
		},
	}

	var buf strings.Builder
	writeInlineContent(&buf, inline, refs)
	result := buf.String()

	expectedImage := "![Test Image](https://developer.apple.com/documentation/images/test.png)"
	if result != expectedImage {
		t.Errorf("Expected image '%s', got '%s'", expectedImage, result)
	}
}

// TestMarkdownSpecialCharacterEscaping tests proper escaping of markdown special characters
func TestMarkdownSpecialCharacterEscaping(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "angle brackets",
			input:    "<ViewController>",
			expected: "&lt;ViewController&gt;",
		},
		{
			name:     "normal text",
			input:    "This is normal text",
			expected: "This is normal text",
		},
		{
			name:     "code voice with backticks",
			input:    "code with ` backtick",
			expected: "`` code with ` backtick ``",
		},
		{
			name:     "code voice without backticks",
			input:    "simple code",
			expected: "`simple code`",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf strings.Builder
			
			if tt.name == "code voice with backticks" || tt.name == "code voice without backticks" {
				inline := InlineContent{
					Type: "codeVoice",
					Code: tt.input,
				}
				writeInlineContent(&buf, inline, nil)
			} else {
				inline := InlineContent{
					Type: "text",
					Text: tt.input,
				}
				writeInlineContent(&buf, inline, nil)
			}

			result := buf.String()
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

// TestEnhancedFrameworkIndex tests the improved framework index generation
func TestEnhancedFrameworkIndex(t *testing.T) {
	tempDir := t.TempDir()

	// Create test JSON files representing various frameworks
	testFrameworks := []struct {
		name    string
		classes []string
	}{
		{
			name:    "SwiftUI",
			classes: []string{"View", "App", "Scene", "WindowGroup", "NavigationView"},
		},
		{
			name:    "UIKit",
			classes: []string{"UIView", "UIViewController", "UIButton", "UILabel"},
		},
		{
			name:    "Foundation",
			classes: []string{"NSObject", "NSString", "NSArray", "NSDictionary"},
		},
		{
			name:    "EndpointSecurity",
			classes: []string{"es_message_t", "es_event_t", "es_client_t"},
		},
	}

	jsonFiles := []string{}

	// Create framework files
	for _, framework := range testFrameworks {
		// Create main framework file
		frameworkDir := filepath.Join(tempDir, "tutorials", "data", "documentation")
		err := os.MkdirAll(frameworkDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create framework directory: %v", err)
		}

		frameworkFile := filepath.Join(frameworkDir, framework.name+".json")
		frameworkDoc := DocJSONData{
			Metadata: Metadata{Title: framework.name, Role: "framework"},
			Abstract: []TextContent{{Type: "text", Text: fmt.Sprintf("The %s framework", framework.name)}},
		}
		frameworkData, _ := json.MarshalIndent(frameworkDoc, "", "  ")
		err = os.WriteFile(frameworkFile, frameworkData, 0644)
		if err != nil {
			t.Fatalf("Failed to create framework file: %v", err)
		}
		jsonFiles = append(jsonFiles, frameworkFile)

		// Create class files
		classDir := filepath.Join(frameworkDir, framework.name)
		err = os.MkdirAll(classDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create class directory: %v", err)
		}

		for _, className := range framework.classes {
			classFile := filepath.Join(classDir, className+".json")
			classDoc := DocJSONData{
				Metadata: Metadata{Title: className, Role: "symbol"},
				Abstract: []TextContent{{Type: "text", Text: fmt.Sprintf("The %s class", className)}},
			}
			classData, _ := json.MarshalIndent(classDoc, "", "  ")
			err = os.WriteFile(classFile, classData, 0644)
			if err != nil {
				t.Fatalf("Failed to create class file: %v", err)
			}
			jsonFiles = append(jsonFiles, classFile)
		}
	}

	// Test framework index creation
	err := createFrameworkIndex(tempDir, jsonFiles)
	if err != nil {
		t.Errorf("createFrameworkIndex() error = %v", err)
	}

	// Read and verify the index
	indexPath := filepath.Join(tempDir, "index.md")
	indexContent, err := os.ReadFile(indexPath)
	if err != nil {
		t.Errorf("Failed to read index file: %v", err)
	}

	indexStr := string(indexContent)

	// Verify main structure
	expectedElements := []string{
		"# Apple Documentation",
		"## Contents",
		"## App Frameworks",
		"### [SwiftUI](tutorials/data/documentation/SwiftUI.md)",
		"<details>",
		"<summary>Major Classes</summary>",
		"- [View](tutorials/data/documentation/SwiftUI/View.md)",
		"- [App](tutorials/data/documentation/SwiftUI/App.md)",
		"</details>",
		"### [UIKit](tutorials/data/documentation/UIKit.md)",
		"- [UIView](tutorials/data/documentation/UIKit/UIView.md)",
		"## System",
		"### [EndpointSecurity](tutorials/data/documentation/EndpointSecurity.md)",
		"## EndpointSecurity Framework Detail",
		"**[EndpointSecurity Documentation](tutorials/data/documentation/EndpointSecurity.md)**",
		"<details>",
		"<summary>EndpointSecurity Classes and Types</summary>",
		"- [es_message_t](tutorials/data/documentation/EndpointSecurity/es_message_t.md)",
		"*Generated on",
	}

	for _, expected := range expectedElements {
		if !strings.Contains(indexStr, expected) {
			t.Errorf("Expected element '%s' not found in index", expected)
		}
	}

	// Verify proper categorization
	if !strings.Contains(indexStr, "App Frameworks") {
		t.Error("Expected 'App Frameworks' category")
	}
	if !strings.Contains(indexStr, "System") {
		t.Error("Expected 'System' category")
	}

	// Verify collapsible sections are properly formatted
	detailsCount := strings.Count(indexStr, "<details>")
	summaryCount := strings.Count(indexStr, "<summary>")
	if detailsCount != summaryCount {
		t.Errorf("Mismatch between <details> (%d) and <summary> (%d) tags", detailsCount, summaryCount)
	}

	// Verify that classes are limited appropriately (should show "... and X more classes" if >15)
	if strings.Contains(indexStr, "... and") && !strings.Contains(indexStr, "more classes") {
		t.Error("Class limit indicator is malformed")
	}
}

// BenchmarkComplexMarkdownGeneration benchmarks markdown generation with complex documents
func BenchmarkComplexMarkdownGeneration(b *testing.B) {
	// Create a complex document similar to real Apple docs
	doc := &DocJSONData{
		Metadata: Metadata{
			Title:       "ComplexFramework",
			Role:        "framework",
			RoleHeading: "Framework",
			Modules:     []Module{{Name: "ComplexFramework"}},
			Platforms: []Platform{
				{Name: "iOS", IntroducedAt: "13.0"},
				{Name: "macOS", IntroducedAt: "10.15"},
				{Name: "watchOS", IntroducedAt: "6.0"},
				{Name: "tvOS", IntroducedAt: "13.0"},
			},
		},
		Abstract: []TextContent{
			{Type: "text", Text: "A complex framework for benchmarking markdown generation performance."},
		},
		PrimaryContentSections: make([]ContentSection, 10), // 10 sections
		TopicSections:          make([]TopicSection, 5),     // 5 topic sections
		References:             make(map[string]Reference),
	}

	// Populate with realistic content
	for i := 0; i < 10; i++ {
		doc.PrimaryContentSections[i] = ContentSection{
			Kind: "content",
			Content: []ContentBlock{
				{
					Type:  "heading",
					Level: 2,
					Text:  fmt.Sprintf("Section %d", i+1),
				},
				{
					Type: "paragraph",
					InlineContent: []InlineContent{
						{Type: "text", Text: fmt.Sprintf("This is the content for section %d with various inline elements like ", i+1)},
						{Type: "codeVoice", Code: "codeExample"},
						{Type: "text", Text: " and "},
						{Type: "emphasis", InlineContent: []InlineContent{
							{Type: "text", Text: "emphasized text"},
						}},
						{Type: "text", Text: "."},
					},
				},
				{
					Type: "unorderedList",
					Items: []Item{
						{
							Content: []ContentBlock{
								{
									Type: "paragraph",
									InlineContent: []InlineContent{
										{Type: "text", Text: fmt.Sprintf("List item %d.1", i+1)},
									},
								},
							},
						},
						{
							Content: []ContentBlock{
								{
									Type: "paragraph",
									InlineContent: []InlineContent{
										{Type: "text", Text: fmt.Sprintf("List item %d.2", i+1)},
									},
								},
							},
						},
					},
				},
			},
		}
	}

	// Add topic sections with references
	for i := 0; i < 5; i++ {
		identifiers := make([]string, 10) // 10 items per section
		for j := 0; j < 10; j++ {
			identifier := fmt.Sprintf("doc://com.apple.documentation/documentation/ComplexFramework/Class%d_%d", i, j)
			identifiers[j] = identifier
			
			// Add reference
			doc.References[identifier] = Reference{
				Title: fmt.Sprintf("Class%d_%d", i, j),
				URL:   identifier,
				Abstract: []TextContent{
					{Type: "text", Text: fmt.Sprintf("Description for class %d_%d", i, j)},
				},
			}
		}
		
		doc.TopicSections[i] = TopicSection{
			Title:       fmt.Sprintf("Topic Group %d", i+1),
			Identifiers: identifiers,
		}
	}

	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		_ = writeMarkdownContent(&buf, doc)
	}
}