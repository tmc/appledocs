package main

import "github.com/tmc/appledocs/occ2go"

// PropertyOverride represents a manually-defined property for classes where
// Apple's documentation is incomplete or missing properties that exist in the
// actual Objective-C runtime.
type PropertyOverride struct {
	Name         string
	Type         string
	Attributes   []string // e.g., "readonly", "readwrite", "nonatomic"
	Comment      string
	Availability occ2go.Availability
}

// propertyOverrides maps framework name -> class name -> property name -> override
// Use this for properties that exist in Objective-C but are missing from Apple's
// JSON documentation exports.
var propertyOverrides = map[string]map[string][]PropertyOverride{
	"Virtualization": {
		"VZVirtualMachineConfiguration": {
			{
				Name: "cpuCount",
				Type: "NSUInteger", // Maps to uint
				Attributes: []string{
					"readwrite",
					"nonatomic",
				},
				Comment: "The number of CPUs for the virtual machine. Must be between minimumAllowedCPUCount and maximumAllowedCPUCount.",
				Availability: occ2go.Availability{
					IntroducedAt: map[string]string{
						"macOS": "11.0",
					},
				},
			},
			{
				Name: "memorySize",
				Type: "unsigned long long", // Maps to uint64
				Attributes: []string{
					"readwrite",
					"nonatomic",
				},
				Comment: "The memory size in bytes for the virtual machine. Must be a multiple of 1MB and between minimumAllowedMemorySize and maximumAllowedMemorySize.",
				Availability: occ2go.Availability{
					IntroducedAt: map[string]string{
						"macOS": "11.0",
					},
				},
			},
		},
	},
}

// MergePropertyOverrides adds manual property overrides to a parsed class.
// This is used during generation to supplement Apple's incomplete documentation.
func MergePropertyOverrides(framework, className string, class *occ2go.ParsedClass) {
	frameworkOverrides, ok := propertyOverrides[framework]
	if !ok {
		return
	}

	classOverrides, ok := frameworkOverrides[className]
	if !ok {
		return
	}

	// Check if properties already exist (don't override documented properties)
	existingProps := make(map[string]bool)
	for _, prop := range class.Properties {
		existingProps[prop.Name] = true
	}

	// Add override properties that don't exist
	for _, override := range classOverrides {
		if !existingProps[override.Name] {
			class.Properties = append(class.Properties, &occ2go.ParsedProperty{
				Name:         override.Name,
				Type:         override.Type,
				Attributes:   override.Attributes,
				Comment:      override.Comment,
				Availability: override.Availability,
				DocURL:       "", // No doc URL for manual overrides
				Abstract:     override.Comment,
			})
		}
	}
}
