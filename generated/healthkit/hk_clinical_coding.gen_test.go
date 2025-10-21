// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKClinicalCoding


// ExampleNewHKClinicalCodingWithSystemVersionCode demonstrates how to create a HKClinicalCoding instance using NewHKClinicalCodingWithSystemVersionCode.
// Creates a clinical coding with the specified system, version, and code.
func ExampleNewHKClinicalCodingWithSystemVersionCode() {
	_ = healthkit.NewHKClinicalCodingWithSystemVersionCode(
		"system", // system string
		"version", // version string
		"code", // code string
	)
	// Output:
}


