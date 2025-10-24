// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit_test

import (
	"github.com/tmc/appledocs/generated/inputmethodkit"
)

// Suppress unused import errors
var _ = inputmethodkit.NewIMKCandidates

// ExampleIMKCandidates_Attributes demonstrates using Attributes on a IMKCandidates instance.
// Returns a dictionary of the style attributes used for the candidates window..
func ExampleIMKCandidates_Attributes() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.Attributes()
	// Output:
	}

// ExampleIMKCandidates_CandidateFrame demonstrates using CandidateFrame on a IMKCandidates instance.
func ExampleIMKCandidates_CandidateFrame() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.CandidateFrame()
	// Output:
	}

// ExampleIMKCandidates_ClearSelection demonstrates using ClearSelection on a IMKCandidates instance.
func ExampleIMKCandidates_ClearSelection() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.ClearSelection()
	// Output:
	}

// ExampleIMKCandidates_DismissesAutomatically demonstrates using DismissesAutomatically on a IMKCandidates instance.
// Returns the state of the flag that determines whether the candidates window dismisses automatically.
func ExampleIMKCandidates_DismissesAutomatically() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.DismissesAutomatically()
	// Output:
	}

// ExampleIMKCandidates_Hide demonstrates using Hide on a IMKCandidates instance.
// Hides a candidates window, if it is visible.
func ExampleIMKCandidates_Hide() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.Hide()
	// Output:
	}

// ExampleIMKCandidates_HideChild demonstrates using HideChild on a IMKCandidates instance.
func ExampleIMKCandidates_HideChild() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.HideChild()
	// Output:
	}

// ExampleIMKCandidates_IsVisible demonstrates using IsVisible on a IMKCandidates instance.
// Returns whether or not the candidates window is visible.
func ExampleIMKCandidates_IsVisible() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.IsVisible()
	// Output:
	}

// ExampleIMKCandidates_PanelType demonstrates using PanelType on a IMKCandidates instance.
// Returns the style of the candidates window.
func ExampleIMKCandidates_PanelType() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.PanelType()
	// Output:
	}

// ExampleIMKCandidates_SelectedCandidate demonstrates using SelectedCandidate on a IMKCandidates instance.
func ExampleIMKCandidates_SelectedCandidate() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.SelectedCandidate()
	// Output:
	}

// ExampleIMKCandidates_SelectedCandidateString demonstrates using SelectedCandidateString on a IMKCandidates instance.
func ExampleIMKCandidates_SelectedCandidateString() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.SelectedCandidateString()
	// Output:
	}

// ExampleIMKCandidates_SelectionKeys demonstrates using SelectionKeys on a IMKCandidates instance.
// Returns an array of   objects where each   object represents a virtual key code.
func ExampleIMKCandidates_SelectionKeys() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.SelectionKeys()
	// Output:
	}

// ExampleIMKCandidates_SelectionKeysKeylayout demonstrates using SelectionKeysKeylayout on a IMKCandidates instance.
// Returns the key layout that maps virtual key codes to selection keys.
func ExampleIMKCandidates_SelectionKeysKeylayout() {
	obj := inputmethodkit.NewIMKCandidates()
	_ = obj.SelectionKeysKeylayout()
	// Output:
	}

// ExampleIMKCandidates_ShowCandidates demonstrates using ShowCandidates on a IMKCandidates instance.
func ExampleIMKCandidates_ShowCandidates() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.ShowCandidates()
	// Output:
	}

// ExampleIMKCandidates_ShowChild demonstrates using ShowChild on a IMKCandidates instance.
func ExampleIMKCandidates_ShowChild() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.ShowChild()
	// Output:
	}

// ExampleIMKCandidates_UpdateCandidates demonstrates using UpdateCandidates on a IMKCandidates instance.
// Updates the candidates that are displayed in the candidates window.
func ExampleIMKCandidates_UpdateCandidates() {
	obj := inputmethodkit.NewIMKCandidates()
	obj.UpdateCandidates()
	// Output:
	}

