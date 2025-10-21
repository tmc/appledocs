// Code generated from Apple documentation for AutomaticAssessmentConfiguration. DO NOT EDIT.

package automaticassessmentconfiguration_test

import (
	"github.com/tmc/appledocs/generated/automaticassessmentconfiguration"
)

// Suppress unused import errors
var _ = automaticassessmentconfiguration.NewAEAssessmentSession

// ExampleNewAEAssessmentSessionWithConfiguration demonstrates how to create a AEAssessmentSession instance using NewAEAssessmentSessionWithConfiguration.
// Creates a new assessment session.
func ExampleNewAEAssessmentSessionWithConfiguration() {
	_ = automaticassessmentconfiguration.NewAEAssessmentSessionWithConfiguration(
		automaticassessmentconfiguration.AEAssessmentConfiguration{}, // configuration AEAssessmentConfiguration
	)
	// Output:
}

