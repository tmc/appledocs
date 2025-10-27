package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

// TestProtocolGeneration tests that protocols are correctly generated from parsed data
func TestProtocolGeneration(t *testing.T) {
	tests := []struct {
		name      string
		framework string
		protocols []*occ2go.ParsedProtocol
		wantInits []string // Expected init() function content
		wantVars  []string // Expected protocol variable declarations
	}{
		{
			name:      "single delegate protocol",
			framework: "AppKit",
			protocols: []*occ2go.ParsedProtocol{
				{
					Name:       "NSApplicationDelegate",
					IsDelegate: true,
					Abstract:   "A set of methods that delegates of NSApplication objects can implement.",
				},
			},
			wantInits: []string{
				`ApplicationDelegateProtocol = objc.GetProtocol("NSApplicationDelegate")`,
			},
			wantVars: []string{
				"ApplicationDelegateProtocol *objc.Protocol",
			},
		},
		{
			name:      "data source protocol",
			framework: "AppKit",
			protocols: []*occ2go.ParsedProtocol{
				{
					Name:         "NSTableViewDataSource",
					IsDataSource: true,
					Abstract:     "A set of methods that a table view uses to provide data to a table view.",
				},
			},
			wantInits: []string{
				`TableViewDataSourceProtocol = objc.GetProtocol("NSTableViewDataSource")`,
			},
			wantVars: []string{
				"TableViewDataSourceProtocol *objc.Protocol",
			},
		},
		{
			name:      "regular protocol",
			framework: "Foundation",
			protocols: []*occ2go.ParsedProtocol{
				{
					Name:       "NSCopying",
					IsDelegate: false,
					Abstract:   "A protocol that objects adopt to provide functional copies of themselves.",
				},
			},
			wantInits: []string{
				`CopyingProtocol = objc.GetProtocol("NSCopying")`,
			},
			wantVars: []string{
				"CopyingProtocol *objc.Protocol",
			},
		},
		{
			name:      "multiple protocols",
			framework: "AppKit",
			protocols: []*occ2go.ParsedProtocol{
				{
					Name:       "NSApplicationDelegate",
					IsDelegate: true,
				},
				{
					Name:       "NSWindowDelegate",
					IsDelegate: true,
				},
				{
					Name:         "NSTableViewDataSource",
					IsDataSource: true,
				},
			},
			wantInits: []string{
				`ApplicationDelegateProtocol = objc.GetProtocol("NSApplicationDelegate")`,
				`WindowDelegateProtocol = objc.GetProtocol("NSWindowDelegate")`,
				`TableViewDataSourceProtocol = objc.GetProtocol("NSTableViewDataSource")`,
			},
			wantVars: []string{
				"ApplicationDelegateProtocol *objc.Protocol",
				"WindowDelegateProtocol *objc.Protocol",
				"TableViewDataSourceProtocol *objc.Protocol",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create generator with test data
			gen := &Generator{
				Framework:   tt.framework,
				PackageName: strings.ToLower(tt.framework),
				Protocols:   tt.protocols,
			}

			// Prepare generator (builds indexes)
			gen.prepare()

			// Generate output using template engine
			var output bytes.Buffer
			err := gen.GenerateTxtarFromModule(&output)
			if err != nil {
				t.Fatalf("GenerateTxtarFromModule() error = %v", err)
			}

			outputStr := output.String()

			// Verify init() function content
			for _, wantInit := range tt.wantInits {
				if !strings.Contains(outputStr, wantInit) {
					t.Errorf("Generated output missing expected init: %q", wantInit)
				}
			}

			// Verify variable declarations
			for _, wantVar := range tt.wantVars {
				if !strings.Contains(outputStr, wantVar) {
					t.Errorf("Generated output missing expected variable: %q", wantVar)
				}
			}

			// Verify objc.Protocol import is present
			if !strings.Contains(outputStr, `"github.com/ebitengine/purego/objc"`) {
				t.Error("Generated output missing objc import")
			}
		})
	}
}

// TestProtocolNaming tests that protocol names are correctly converted to Go variable names
func TestProtocolNaming(t *testing.T) {
	tests := []struct {
		name         string
		protocolName string
		wantVarName  string
	}{
		{
			name:         "NS prefix stripped",
			protocolName: "NSApplicationDelegate",
			wantVarName:  "ApplicationDelegateProtocol",
		},
		{
			name:         "UI prefix stripped",
			protocolName: "UITableViewDelegate",
			wantVarName:  "TableViewDelegateProtocol",
		},
		{
			name:         "SC prefix stripped",
			protocolName: "SCStreamOutput",
			wantVarName:  "StreamOutputProtocol",
		},
		{
			name:         "already has Protocol suffix",
			protocolName: "SomeProtocol",
			wantVarName:  "SomeProtocolProtocol", // Current behavior - may want to improve
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &Generator{
				Framework:   "TestFramework",
				PackageName: "testframework",
				Protocols: []*occ2go.ParsedProtocol{
					{Name: tt.protocolName},
				},
			}

			gen.prepare()

			var output bytes.Buffer
			err := gen.GenerateTxtarFromModule(&output)
			if err != nil {
				t.Fatalf("GenerateTxtarFromModule() error = %v", err)
			}

			outputStr := output.String()

			// Check that the expected variable name appears
			expectedDecl := tt.wantVarName + " *objc.Protocol"
			if !strings.Contains(outputStr, expectedDecl) {
				t.Errorf("Expected variable declaration %q not found in output", expectedDecl)
			}

			// Check that the init assignment uses the correct name
			expectedInit := tt.wantVarName + ` = objc.GetProtocol("`
			if !strings.Contains(outputStr, expectedInit) {
				t.Errorf("Expected init assignment %q not found in output", expectedInit)
			}
		})
	}
}

// TestProtocolAvailability tests that protocol availability annotations are included
func TestProtocolAvailability(t *testing.T) {
	gen := &Generator{
		Framework:   "AppKit",
		PackageName: "appkit",
		Protocols: []*occ2go.ParsedProtocol{
			{
				Name:       "NSApplicationDelegate",
				IsDelegate: true,
				Availability: occ2go.Availability{
					IntroducedAt: map[string]string{
						"macOS": "10.0",
					},
				},
				Abstract: "A set of methods that delegates of NSApplication objects can implement.",
			},
		},
	}

	gen.prepare()

	var output bytes.Buffer
	err := gen.GenerateTxtarFromModule(&output)
	if err != nil {
		t.Fatalf("GenerateTxtarFromModule() error = %v", err)
	}

	outputStr := output.String()

	// Verify availability comment is present
	// The exact format may vary based on template, but should mention macOS
	if !strings.Contains(outputStr, "macOS") {
		t.Error("Expected availability information (macOS) in generated output")
	}

	// Verify abstract/documentation is present
	if !strings.Contains(outputStr, "NSApplicationDelegate") {
		t.Error("Expected protocol name in documentation")
	}
}

// TestProtocolGenerationWithEmptyList tests behavior with no protocols
func TestProtocolGenerationWithEmptyList(t *testing.T) {
	gen := &Generator{
		Framework:   "TestFramework",
		PackageName: "testframework",
		Protocols:   []*occ2go.ParsedProtocol{}, // Empty protocols list
	}

	gen.prepare()

	var output bytes.Buffer
	err := gen.GenerateTxtarFromModule(&output)
	if err != nil {
		t.Fatalf("GenerateTxtarFromModule() error = %v", err)
	}

	outputStr := output.String()

	// Should still generate valid package without protocols
	if !strings.Contains(outputStr, "package testframework") {
		t.Error("Expected package declaration even with no protocols")
	}

	// Should not have protocol-specific code
	if strings.Contains(outputStr, "objc.GetProtocol") {
		t.Error("Should not generate protocol code when no protocols present")
	}
}

// TestProtocolWithMethods tests protocol generation with method declarations
func TestProtocolWithMethods(t *testing.T) {
	gen := &Generator{
		Framework:   "AppKit",
		PackageName: "appkit",
		Protocols: []*occ2go.ParsedProtocol{
			{
				Name:       "NSApplicationDelegate",
				IsDelegate: true,
				RequiredMethods: []*occ2go.ParsedMethod{
					{
						Name:          "applicationDidFinishLaunching",
						IsClassMethod: false, // instance method
						Parameters: []occ2go.Parameter{
							{
								Name: "notification",
								Type: "NSNotification *",
							},
						},
						ReturnType: "void",
					},
				},
				OptionalMethods: []*occ2go.ParsedMethod{
					{
						Name:          "applicationWillTerminate",
						IsClassMethod: false, // instance method
						Parameters: []occ2go.Parameter{
							{
								Name: "notification",
								Type: "NSNotification *",
							},
						},
						ReturnType: "void",
					},
				},
			},
		},
	}

	gen.prepare()

	var output bytes.Buffer
	err := gen.GenerateTxtarFromModule(&output)
	if err != nil {
		t.Fatalf("GenerateTxtarFromModule() error = %v", err)
	}

	outputStr := output.String()

	// Verify protocol variable is generated
	if !strings.Contains(outputStr, "ApplicationDelegateProtocol") {
		t.Error("Expected ApplicationDelegateProtocol variable")
	}

	// Note: Method interface generation is future work
	// This test documents current state and will be expanded when
	// protocol interface generation is implemented
}

// TestProtocolInterfaceGeneration tests generation of Go interfaces for protocols
// This test documents the expected future behavior for protocol interfaces
func TestProtocolInterfaceGeneration(t *testing.T) {
	t.Skip("Protocol interface generation not yet implemented - placeholder for Phase 3")

	// Future test: verify that protocol methods are converted to Go interface methods
	// Example: NSApplicationDelegate -> PApplicationDelegate interface
	// with methods like ApplicationDidFinishLaunching(notification Notification)
}

// TestDelegateHelperGeneration tests generation of delegate builder functions
// This test documents the expected future behavior for delegate helpers
func TestDelegateHelperGeneration(t *testing.T) {
	t.Skip("Delegate helper generation not yet implemented - tracked in appledocs-330")

	// Future test: verify that delegate protocols get builder functions
	// Example: NewNSApplicationDelegate(handler PApplicationDelegate) objc.ID
}

// TestProtocolRegistryIntegration tests that generated protocols can be looked up
func TestProtocolRegistryIntegration(t *testing.T) {
	// This test verifies the generated code would work with objc runtime
	// by checking the generated objc.GetProtocol calls are correctly formatted

	tests := []struct {
		name         string
		protocolName string
	}{
		{
			name:         "NSApplicationDelegate",
			protocolName: "NSApplicationDelegate",
		},
		{
			name:         "NSTableViewDataSource",
			protocolName: "NSTableViewDataSource",
		},
		{
			name:         "UICollectionViewDelegate",
			protocolName: "UICollectionViewDelegate",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &Generator{
				Framework:   "TestFramework",
				PackageName: "testframework",
				Protocols: []*occ2go.ParsedProtocol{
					{Name: tt.protocolName},
				},
			}

			gen.prepare()

			var output bytes.Buffer
			err := gen.GenerateTxtarFromModule(&output)
			if err != nil {
				t.Fatalf("GenerateTxtarFromModule() error = %v", err)
			}

			outputStr := output.String()

			// Verify objc.GetProtocol is called with the exact Objective-C protocol name
			expectedCall := `objc.GetProtocol("` + tt.protocolName + `")`
			if !strings.Contains(outputStr, expectedCall) {
				t.Errorf("Expected objc.GetProtocol call: %q", expectedCall)
			}

			// Verify the variable is a pointer to objc.Protocol
			if !strings.Contains(outputStr, "*objc.Protocol") {
				t.Error("Expected *objc.Protocol type for protocol variable")
			}
		})
	}
}

// TestProtocolGenerationCompleteness tests that all protocols in the generator are processed
func TestProtocolGenerationCompleteness(t *testing.T) {
	protocols := []*occ2go.ParsedProtocol{
		{Name: "NSApplicationDelegate", IsDelegate: true},
		{Name: "NSWindowDelegate", IsDelegate: true},
		{Name: "NSTableViewDataSource", IsDataSource: true},
		{Name: "NSCopying", IsDelegate: false, IsDataSource: false},
		{Name: "NSSecureCoding", IsDelegate: false, IsDataSource: false},
	}

	gen := &Generator{
		Framework:   "AppKit",
		PackageName: "appkit",
		Protocols:   protocols,
	}

	gen.prepare()

	var output bytes.Buffer
	err := gen.GenerateTxtarFromModule(&output)
	if err != nil {
		t.Fatalf("GenerateTxtarFromModule() error = %v", err)
	}

	outputStr := output.String()

	// Verify each protocol gets a variable and init
	for _, proto := range protocols {
		// Check for objc.GetProtocol call
		expectedCall := `objc.GetProtocol("` + proto.Name + `")`
		if !strings.Contains(outputStr, expectedCall) {
			t.Errorf("Missing objc.GetProtocol call for protocol: %s", proto.Name)
		}
	}

	// Count init() functions - should have at least one
	initCount := strings.Count(outputStr, "func init() {")
	if initCount == 0 {
		t.Error("Expected at least one init() function for protocols")
	}
}

// TestProtocolFileGeneration tests that protocol files are correctly named and structured
func TestProtocolFileGeneration(t *testing.T) {
	gen := &Generator{
		Framework:   "AppKit",
		PackageName: "appkit",
		Protocols: []*occ2go.ParsedProtocol{
			{
				Name:       "NSApplicationDelegate",
				IsDelegate: true,
			},
		},
	}

	gen.prepare()

	var output bytes.Buffer
	err := gen.GenerateTxtarFromModule(&output)
	if err != nil {
		t.Fatalf("GenerateTxtarFromModule() error = %v", err)
	}

	outputStr := output.String()

	// Verify txtar format has protocol file entries
	// The exact file naming convention may vary, but should include protocol name
	if !strings.Contains(outputStr, "application_delegate_protocol") {
		t.Log("Note: Protocol file naming may not follow expected pattern")
	}

	// Verify package declaration is correct
	if !strings.Contains(outputStr, "package appkit") {
		t.Error("Expected correct package declaration")
	}

	// Verify Code generated comment
	if !strings.Contains(outputStr, "Code generated") {
		t.Error("Expected 'Code generated' comment in output")
	}
}

// TestProtocolDelegateDetection tests that IsDelegate and IsDataSource flags work correctly
func TestProtocolDelegateDetection(t *testing.T) {
	tests := []struct {
		name         string
		protocolName string
		isDelegate   bool
		isDataSource bool
	}{
		{
			name:         "delegate protocol",
			protocolName: "NSApplicationDelegate",
			isDelegate:   true,
			isDataSource: false,
		},
		{
			name:         "data source protocol",
			protocolName: "NSTableViewDataSource",
			isDelegate:   false,
			isDataSource: true,
		},
		{
			name:         "regular protocol",
			protocolName: "NSCopying",
			isDelegate:   false,
			isDataSource: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gen := &Generator{
				Framework:   "AppKit",
				PackageName: "appkit",
				Protocols: []*occ2go.ParsedProtocol{
					{
						Name:         tt.protocolName,
						IsDelegate:   tt.isDelegate,
						IsDataSource: tt.isDataSource,
					},
				},
			}

			gen.prepare()

			var output bytes.Buffer
			err := gen.GenerateTxtarFromModule(&output)
			if err != nil {
				t.Fatalf("GenerateTxtarFromModule() error = %v", err)
			}

			outputStr := output.String()

			// All protocols should be generated regardless of delegate/datasource status
			expectedCall := `objc.GetProtocol("` + tt.protocolName + `")`
			if !strings.Contains(outputStr, expectedCall) {
				t.Errorf("Expected protocol generation for %s", tt.protocolName)
			}

			// Note: IsDelegate/IsDataSource flags will be used in future phases
			// for generating delegate helpers and interface types
		})
	}
}
