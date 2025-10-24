// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayerMediaSelectionCriteria

// ExampleNewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics demonstrates how to create a PlayerMediaSelectionCriteria instance using NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics.
// Creates media selection criteria with the preferred languages and media characteristics.
func ExampleNewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics() {
	_ = avfoundation.NewPlayerMediaSelectionCriteriaWithPreferredLanguagesPreferredMediaCharacteristics(
		[]avfoundation.string{}, // preferredLanguages []string
		[]avfoundation.string{}, // preferredMediaCharacteristics []string
	)
	// Output:
}
// ExampleNewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics demonstrates how to create a PlayerMediaSelectionCriteria instance using NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics.
// Creates media selection criteria with the principal media characteristics, and preferred languages and media characteristics.
func ExampleNewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics() {
	_ = avfoundation.NewPlayerMediaSelectionCriteriaWithPrincipalMediaCharacteristicsPreferredLanguagesPreferredMediaCharacteristics(
		[]avfoundation.string{}, // principalMediaCharacteristics []string
		[]avfoundation.string{}, // preferredLanguages []string
		[]avfoundation.string{}, // preferredMediaCharacteristics []string
	)
	// Output:
}
