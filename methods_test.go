package appledocs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseExternalID(t *testing.T) {
	tests := []struct {
		name         string
		externalID   string
		wantSymType  string
		wantClass    string
		wantMemType  string
		wantMemName  string
		wantOk       bool
	}{
		{
			name:         "class symbol",
			externalID:   "c:objc(cs)SKPaymentQueue",
			wantSymType:  "cs",
			wantClass:    "SKPaymentQueue",
			wantMemType:  "",
			wantMemName:  "",
			wantOk:       true,
		},
		{
			name:         "instance method",
			externalID:   "c:objc(cs)SKPaymentQueue(im)restoreCompletedTransactions",
			wantSymType:  "cs",
			wantClass:    "SKPaymentQueue",
			wantMemType:  "im",
			wantMemName:  "restoreCompletedTransactions",
			wantOk:       true,
		},
		{
			name:         "class method",
			externalID:   "c:objc(cs)SKPaymentQueue(cm)canMakePayments",
			wantSymType:  "cs",
			wantClass:    "SKPaymentQueue",
			wantMemType:  "cm",
			wantMemName:  "canMakePayments",
			wantOk:       true,
		},
		{
			name:         "property",
			externalID:   "c:objc(cs)SKPaymentQueue(py)delegate",
			wantSymType:  "cs",
			wantClass:    "SKPaymentQueue",
			wantMemType:  "py",
			wantMemName:  "delegate",
			wantOk:       true,
		},
		{
			name:         "protocol",
			externalID:   "c:objc(pl)NSObject",
			wantSymType:  "pl",
			wantClass:    "NSObject",
			wantMemType:  "",
			wantMemName:  "",
			wantOk:       true,
		},
		{
			name:         "invalid format",
			externalID:   "not-a-valid-id",
			wantOk:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSymType, gotClass, gotMemType, gotMemName, gotOk := parseExternalID(tt.externalID)

			if gotOk != tt.wantOk {
				t.Errorf("parseExternalID() ok = %v, want %v", gotOk, tt.wantOk)
				return
			}

			if !tt.wantOk {
				return
			}

			if gotSymType != tt.wantSymType {
				t.Errorf("parseExternalID() symbolType = %v, want %v", gotSymType, tt.wantSymType)
			}
			if gotClass != tt.wantClass {
				t.Errorf("parseExternalID() className = %v, want %v", gotClass, tt.wantClass)
			}
			if gotMemType != tt.wantMemType {
				t.Errorf("parseExternalID() memberType = %v, want %v", gotMemType, tt.wantMemType)
			}
			if gotMemName != tt.wantMemName {
				t.Errorf("parseExternalID() memberName = %v, want %v", gotMemName, tt.wantMemName)
			}
		})
	}
}

func TestBuildSignature(t *testing.T) {
	tests := []struct {
		name      string
		fragments []Fragment
		want      string
	}{
		{
			name: "simple method",
			fragments: []Fragment{
				{Kind: "keyword", Text: "func"},
				{Kind: "text", Text: " "},
				{Kind: "identifier", Text: "myMethod"},
				{Kind: "text", Text: "()"},
			},
			want: "func myMethod()",
		},
		{
			name: "method with parameters",
			fragments: []Fragment{
				{Kind: "text", Text: "- ("},
				{Kind: "keyword", Text: "void"},
				{Kind: "text", Text: ")"},
				{Kind: "identifier", Text: "setDelegate"},
				{Kind: "text", Text: ":("},
				{Kind: "typeIdentifier", Text: "id"},
				{Kind: "text", Text: ")delegate"},
			},
			want: "- (void)setDelegate:(id)delegate",
		},
		{
			name:      "empty fragments",
			fragments: []Fragment{},
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildSignature(tt.fragments)
			if got != tt.want {
				t.Errorf("buildSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscoverMethods(t *testing.T) {
	// Check if cache directory exists
	cacheDir := filepath.Join(os.Getenv("HOME"), ".appledocs", "cache", "developer.apple.com", "tutorials", "data", "documentation")
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found, skipping integration test")
	}

	fsys, err := Open(cacheDir)
	if err != nil {
		t.Fatalf("Failed to open cache directory: %v", err)
	}

	classes, err := DiscoverMethods(fsys)
	if err != nil {
		t.Fatalf("DiscoverMethods() error = %v", err)
	}

	if len(classes) == 0 {
		t.Error("DiscoverMethods() returned no classes")
	}

	// Check if we found any known classes (if they exist in cache)
	t.Logf("Discovered %d classes", len(classes))

	// Look for a specific class if it exists
	if classMethods, ok := classes["SKPaymentQueue"]; ok {
		t.Logf("Found SKPaymentQueue with:")
		t.Logf("  - %d instance methods", len(classMethods.InstanceMethods))
		t.Logf("  - %d class methods", len(classMethods.ClassMethods))
		t.Logf("  - %d properties", len(classMethods.Properties))
		t.Logf("  - %d initializers", len(classMethods.Initializers))

		// Verify we have some methods
		if len(classMethods.InstanceMethods) == 0 && len(classMethods.ClassMethods) == 0 {
			t.Error("SKPaymentQueue has no methods")
		}

		// Check that method info is populated
		if len(classMethods.InstanceMethods) > 0 {
			method := classMethods.InstanceMethods[0]
			if method.ClassName != "SKPaymentQueue" {
				t.Errorf("Method className = %v, want SKPaymentQueue", method.ClassName)
			}
			if method.Name == "" {
				t.Error("Method name is empty")
			}
			if method.ExternalID == "" {
				t.Error("Method externalID is empty")
			}
		}
	}

	// Sample a few classes to verify structure
	count := 0
	for className, classMethods := range classes {
		if count >= 3 {
			break
		}
		t.Logf("Class: %s (externalID: %s)", className, classMethods.ClassExternalID)
		t.Logf("  Instance methods: %d", len(classMethods.InstanceMethods))
		t.Logf("  Class methods: %d", len(classMethods.ClassMethods))
		t.Logf("  Properties: %d", len(classMethods.Properties))
		t.Logf("  Initializers: %d", len(classMethods.Initializers))
		count++
	}
}

func TestDiscoverMethodsForFramework(t *testing.T) {
	// Check if cache directory exists
	cacheDir := filepath.Join(os.Getenv("HOME"), ".appledocs", "cache", "developer.apple.com", "tutorials", "data", "documentation")
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		t.Skip("Cache directory not found, skipping integration test")
	}

	// Test with a specific framework subdirectory
	frameworkDir := filepath.Join(cacheDir, "StoreKit")
	if _, err := os.Stat(frameworkDir); os.IsNotExist(err) {
		t.Skip("StoreKit framework not found in cache, skipping test")
	}

	fsys, err := Open(frameworkDir)
	if err != nil {
		t.Fatalf("Failed to open framework directory: %v", err)
	}

	classes, err := DiscoverMethods(fsys)
	if err != nil {
		t.Fatalf("DiscoverMethods() error = %v", err)
	}

	if len(classes) == 0 {
		t.Error("DiscoverMethods() returned no classes for StoreKit")
	}

	t.Logf("Discovered %d classes in StoreKit framework", len(classes))

	// Check for known StoreKit classes
	knownClasses := []string{"SKPaymentQueue", "SKPayment", "SKPaymentTransaction"}
	for _, expectedClass := range knownClasses {
		if classMethods, ok := classes[expectedClass]; ok {
			t.Logf("Found %s:", expectedClass)
			t.Logf("  Methods: %d instance, %d class",
				len(classMethods.InstanceMethods), len(classMethods.ClassMethods))
			t.Logf("  Properties: %d", len(classMethods.Properties))
		}
	}
}
