// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration_test

import (
	"github.com/tmc/appledocs/generated/automaticassessmentconfiguration"
)

// Suppress unused import errors
var _ = automaticassessmentconfiguration.NewAEAssessmentApplication

// ExampleNewAEAssessmentApplicationWithBundleIdentifier demonstrates how to create a AEAssessmentApplication instance using NewAEAssessmentApplicationWithBundleIdentifier.
// Creates a representation of an app using its bundle identifier.
func ExampleNewAEAssessmentApplicationWithBundleIdentifier() {
	_ = automaticassessmentconfiguration.NewAEAssessmentApplicationWithBundleIdentifier(
		"bundleIdentifier", // bundleIdentifier string
	)
	// Output:
}
// ExampleNewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier demonstrates how to create a AEAssessmentApplication instance using NewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier.
// Creates a representation of an app using its bundle and team identifiers.
func ExampleNewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier() {
	_ = automaticassessmentconfiguration.NewAEAssessmentApplicationWithBundleIdentifierTeamIdentifier(
		"bundleIdentifier", // bundleIdentifier string
		"teamIdentifier", // teamIdentifier string
	)
	// Output:
}
