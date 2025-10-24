package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// FrameworkInfo holds metadata about a single framework.
type FrameworkInfo struct {
	// Name is the framework name (e.g., "AppKit", "Foundation")
	Name string

	// PackageName is the lowercase package name (e.g., "appkit", "foundation")
	PackageName string

	// Prefix is the common type prefix for C-style frameworks (e.g., "CG" for CoreGraphics, "CF" for CoreFoundation)
	// Empty for Objective-C frameworks
	Prefix string

	// ImportPath is the full import path for this framework
	// e.g., "github.com/tmc/appledocs/generated/appkit"
	ImportPath string
}

// FrameworkRegistry manages metadata about all known frameworks.
type FrameworkRegistry struct {
	// frameworks maps framework names to their metadata
	frameworks map[string]*FrameworkInfo

	// packageToFramework maps package names to framework names
	packageToFramework map[string]string

	// baseModule is the base import path for all generated frameworks
	// e.g., "github.com/tmc/appledocs/generated"
	baseModule string
}

// NewFrameworkRegistry creates a new framework registry.
func NewFrameworkRegistry(baseModule string) *FrameworkRegistry {
	return &FrameworkRegistry{
		frameworks:         make(map[string]*FrameworkInfo),
		packageToFramework: make(map[string]string),
		baseModule:         baseModule,
	}
}

// Register adds a framework to the registry.
func (r *FrameworkRegistry) Register(name, prefix string) *FrameworkInfo {
	packageName := strings.ToLower(name)
	info := &FrameworkInfo{
		Name:        name,
		PackageName: packageName,
		Prefix:      prefix,
		ImportPath:  r.baseModule + "/" + packageName,
	}
	r.frameworks[name] = info
	r.packageToFramework[packageName] = name
	return info
}

// Get retrieves framework info by name.
func (r *FrameworkRegistry) Get(name string) (*FrameworkInfo, bool) {
	info, ok := r.frameworks[name]
	return info, ok
}

// GetByPackage retrieves framework info by package name.
func (r *FrameworkRegistry) GetByPackage(pkgName string) (*FrameworkInfo, bool) {
	name, ok := r.packageToFramework[pkgName]
	if !ok {
		return nil, false
	}
	return r.Get(name)
}

// GetImportPath returns the import path for a framework name.
// Returns empty string if framework not found.
// UNUSED: Commented out as unreachable code
/*
func (r *FrameworkRegistry) GetImportPath(name string) string {
	if info, ok := r.Get(name); ok {
		return info.ImportPath
	}
	return ""
}
*/

// GetImportPathByPackage returns the import path for a package name.
// Returns empty string if package not found.
func (r *FrameworkRegistry) GetImportPathByPackage(pkgName string) string {
	if info, ok := r.GetByPackage(pkgName); ok {
		return info.ImportPath
	}
	return ""
}

// GetPrefix returns the framework prefix for a framework name.
// Returns empty string if framework not found or has no prefix.
func (r *FrameworkRegistry) GetPrefix(name string) string {
	if info, ok := r.Get(name); ok {
		return info.Prefix
	}
	return ""
}

// All returns all registered frameworks sorted by name.
func (r *FrameworkRegistry) All() []*FrameworkInfo {
	infos := make([]*FrameworkInfo, 0, len(r.frameworks))
	for _, info := range r.frameworks {
		infos = append(infos, info)
	}
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].Name < infos[j].Name
	})
	return infos
}

// DiscoverFromDirectory discovers frameworks by scanning a generated directory.
// It looks for subdirectories that contain Go packages and registers them.
func (r *FrameworkRegistry) DiscoverFromDirectory(generatedDir string) error {
	// Read directory entries
	entries, err := os.ReadDir(generatedDir)
	if err != nil {
		// If directory doesn't exist, that's ok - we'll use defaults
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read generated directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		packageName := entry.Name()

		// Skip special directories
		if packageName == "objc" || packageName == "objectivec" {
			continue
		}

		// Check if this looks like a framework package (has .go files)
		pkgPath := filepath.Join(generatedDir, packageName)
		pkgEntries, err := os.ReadDir(pkgPath)
		if err != nil {
			continue
		}

		hasGoFiles := false
		for _, pe := range pkgEntries {
			if !pe.IsDir() && strings.HasSuffix(pe.Name(), ".go") {
				hasGoFiles = true
				break
			}
		}

		if !hasGoFiles {
			continue
		}

		// Convert package name to framework name (e.g., "appkit" -> "AppKit")
		frameworkName := packageNameToFrameworkName(packageName)

		// Determine prefix by checking known patterns
		prefix := determineFrameworkPrefix(frameworkName)

		r.Register(frameworkName, prefix)
	}

	return nil
}

// packageNameToFrameworkName converts a lowercase package name to a framework name.
// Examples: "appkit" -> "AppKit", "coregraphics" -> "CoreGraphics"
func packageNameToFrameworkName(pkgName string) string {
	// Special cases for known frameworks
	knownFrameworks := map[string]string{
		"appkit":                 "AppKit",
		"foundation":             "Foundation",
		"coregraphics":           "CoreGraphics",
		"corefoundation":         "CoreFoundation",
		"coreaudio":              "CoreAudio",
		"coredata":               "CoreData",
		"coreimage":              "CoreImage",
		"corevideo":              "CoreVideo",
		"coretext":               "CoreText",
		"coremedia":              "CoreMedia",
		"coremidi":               "CoreMIDI",
		"corebluetooth":          "CoreBluetooth",
		"corelocation":           "CoreLocation",
		"coremotion":             "CoreMotion",
		"coreanimation":          "CoreAnimation",
		"quartzcore":             "QuartzCore",
		"quartz":                 "Quartz",
		"iokit":                  "IOKit",
		"metalkit":               "MetalKit",
		"metal":                  "Metal",
		"scenekit":               "SceneKit",
		"spritekit":              "SpriteKit",
		"gamekit":                "GameKit",
		"mapkit":                 "MapKit",
		"cloudkit":               "CloudKit",
		"avfoundation":           "AVFoundation",
		"avkit":                  "AVKit",
		"webkit":                 "WebKit",
		"usernotifications":      "UserNotifications",
		"uniformtypeidentifiers": "UniformTypeIdentifiers",
		"naturallanguage":        "NaturalLanguage",
		"vision":                 "Vision",
		"coreml":                 "CoreML",
		"createml":               "CreateML",
		"accessibility":          "Accessibility",
		"accounts":               "Accounts",
		"addressbook":            "AddressBook",
		"adservices":             "AdServices",
		"adsupport":              "AdSupport",
		"appintents":             "AppIntents",
		"screencapturekit":       "ScreenCaptureKit",
		"objectivec":             "ObjectiveC",
	}

	if framework, ok := knownFrameworks[pkgName]; ok {
		return framework
	}

	// Default: capitalize first letter
	if len(pkgName) == 0 {
		return pkgName
	}
	return strings.ToUpper(pkgName[:1]) + pkgName[1:]
}

// determineFrameworkPrefix determines the C-style prefix for a framework.
// Returns empty string for Objective-C frameworks.
func determineFrameworkPrefix(frameworkName string) string {
	// Known C-style framework prefixes
	prefixes := map[string]string{
		"CoreGraphics":   "CG",
		"CoreFoundation": "CF",
		"CoreAudio":      "CA",
		"CoreImage":      "CI",
		"CoreVideo":      "CV",
		"CoreText":       "CT",
		"CoreMedia":      "CM",
		"CoreMIDI":       "MIDI",
		"IOKit":          "IO",
		"Quartz":         "Q",
		"Metal":          "MTL",
	}

	if prefix, ok := prefixes[frameworkName]; ok {
		return prefix
	}

	// No prefix for Objective-C frameworks
	return ""
}

// globalRegistry is the global framework registry instance.
var globalRegistry *FrameworkRegistry

// initializeFrameworkRegistry initializes the global framework registry.
// It first tries to discover frameworks from the generated directory,
// then falls back to registering commonly used frameworks.
func initializeFrameworkRegistry(baseModule, generatedDir string) error {
	globalRegistry = NewFrameworkRegistry(baseModule)

	// Try to discover from generated directory first
	if generatedDir != "" {
		if err := globalRegistry.DiscoverFromDirectory(generatedDir); err != nil {
			return err
		}
	}

	// Register commonly used frameworks that may not exist yet
	// This ensures the registry works even before frameworks are generated
	commonFrameworks := []struct {
		name   string
		prefix string
	}{
		{"ObjectiveC", ""},
		{"Foundation", ""},
		{"AppKit", ""},
		{"CoreGraphics", "CG"},
		{"CoreFoundation", "CF"},
		{"CoreAudio", "CA"},
		{"QuartzCore", ""},
		{"CloudKit", ""},
		{"UserNotifications", ""},
		{"UniformTypeIdentifiers", ""},
	}

	for _, fw := range commonFrameworks {
		// Only register if not already discovered
		if _, exists := globalRegistry.Get(fw.name); !exists {
			globalRegistry.Register(fw.name, fw.prefix)
		}
	}

	return nil
}

// GetFrameworkInfo returns framework info from the global registry.
// UNUSED: Commented out as unreachable code
/*
func GetFrameworkInfo(name string) (*FrameworkInfo, bool) {
	if globalRegistry == nil {
		return nil, false
	}
	return globalRegistry.Get(name)
}
*/

// GetFrameworkImportPath returns the import path for a framework.
// UNUSED: Commented out as unreachable code
/*
func GetFrameworkImportPath(name string) string {
	if globalRegistry == nil {
		return ""
	}
	return globalRegistry.GetImportPath(name)
}
*/

// GetFrameworkPrefix returns the prefix for a framework.
func GetFrameworkPrefix(name string) string {
	if globalRegistry == nil {
		return ""
	}
	return globalRegistry.GetPrefix(name)
}

// ExtractFrameworkFromType extracts the framework package name from a Go type string.
// Examples:
//
//	"coregraphics.CGRect" -> "coregraphics"
//	"foundation.String" -> "foundation"
//	"int" -> ""
//	"[]appkit.Window" -> "appkit"
func ExtractFrameworkFromType(goType string) string {
	// Strip array/slice/pointer prefixes
	goType = strings.TrimLeft(goType, "*[]")

	// Check for package-qualified type
	if idx := strings.Index(goType, "."); idx != -1 {
		return goType[:idx]
	}

	return ""
}

// GetImportPathFromType extracts the import path from a Go type string.
// Returns empty string if the type doesn't require an import.
// Examples:
//
//	"coregraphics.CGRect" -> "github.com/tmc/appledocs/generated/coregraphics"
//	"foundation.String" -> "github.com/tmc/appledocs/generated/foundation"
//	"int" -> ""
func GetImportPathFromType(goType string) string {
	pkgName := ExtractFrameworkFromType(goType)
	if pkgName == "" {
		return ""
	}

	// Special cases
	if pkgName == "unsafe" {
		return "" // unsafe is built-in, no import needed
	}

	if pkgName == "objc" {
		// objc.* types come from our wrapper package (generated/objc)
		// This is the cached selector wrapper, not a framework
		if globalRegistry != nil {
			return globalRegistry.baseModule + "/objc"
		}
		return ""
	}

	if globalRegistry == nil {
		return ""
	}

	return globalRegistry.GetImportPathByPackage(pkgName)
}
