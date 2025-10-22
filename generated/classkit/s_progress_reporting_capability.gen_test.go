// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit_test

import (
	"github.com/tmc/appledocs/generated/classkit"
)

// Suppress unused import errors
var _ = classkit.NewSProgressReportingCapability

// ExampleNewSProgressReportingCapabilityWithKindDetails demonstrates how to create a SProgressReportingCapability instance using NewSProgressReportingCapabilityWithKindDetails.
// Creates a new progress reporting capability of the given type with a descriptive string.
func ExampleNewSProgressReportingCapabilityWithKindDetails() {
	_ = classkit.NewSProgressReportingCapabilityWithKindDetails(
		classkit.SProgressReportingCapabilityKind{}, // kind SProgressReportingCapabilityKind
		"details", // details string
	)
	// Output:
}
