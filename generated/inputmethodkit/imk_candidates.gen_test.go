// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit_test

import (
	"github.com/tmc/appledocs/generated/inputmethodkit"
)

// Suppress unused import errors
var _ = inputmethodkit.NewIMKCandidates

// ExampleNewIMKCandidatesWithServerPanelType demonstrates how to create a IMKCandidates instance using NewIMKCandidatesWithServerPanelType.
// Returns the initialized   object.
func ExampleNewIMKCandidatesWithServerPanelType() {
	_ = inputmethodkit.NewIMKCandidatesWithServerPanelType(
		inputmethodkit.IMKServer{}, // server IMKServer
		inputmethodkit.IMKCandidatePanelType{}, // panelType IMKCandidatePanelType
	)
	// Output:
}
// ExampleNewIMKCandidatesWithServerPanelTypeStyleType demonstrates how to create a IMKCandidates instance using NewIMKCandidatesWithServerPanelTypeStyleType.
func ExampleNewIMKCandidatesWithServerPanelTypeStyleType() {
	_ = inputmethodkit.NewIMKCandidatesWithServerPanelTypeStyleType(
		inputmethodkit.IMKServer{}, // server IMKServer
		inputmethodkit.IMKCandidatePanelType{}, // panelType IMKCandidatePanelType
		inputmethodkit.IMKStyleType{}, // style IMKStyleType
	)
	// Output:
}
