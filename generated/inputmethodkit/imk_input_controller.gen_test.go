// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit_test

import (
	"github.com/tmc/appledocs/generated/inputmethodkit"
)

// Suppress unused import errors
var _ = inputmethodkit.NewIMKInputController

// ExampleIMKInputController_CancelComposition demonstrates using CancelComposition on a IMKInputController instance.
// Stops the current composition and replaces marked text with the original text.
func ExampleIMKInputController_CancelComposition() {
	obj := inputmethodkit.NewIMKInputController()
	obj.CancelComposition()
	// Output:
	}

// ExampleIMKInputController_Client demonstrates using Client on a IMKInputController instance.
// Returns the client object associated with the input controller.
func ExampleIMKInputController_Client() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.Client()
	// Output:
	}

// ExampleIMKInputController_Delegate demonstrates using Delegate on a IMKInputController instance.
// Returns the delegate for input controller  object.
func ExampleIMKInputController_Delegate() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.Delegate()
	// Output:
	}

// ExampleIMKInputController_HidePalettes demonstrates using HidePalettes on a IMKInputController instance.
// Informs an input method that it should  close any visible user interface.
func ExampleIMKInputController_HidePalettes() {
	obj := inputmethodkit.NewIMKInputController()
	obj.HidePalettes()
	// Output:
	}

// ExampleIMKInputController_InputControllerWillClose demonstrates using InputControllerWillClose on a IMKInputController instance.
func ExampleIMKInputController_InputControllerWillClose() {
	obj := inputmethodkit.NewIMKInputController()
	obj.InputControllerWillClose()
	// Output:
	}

// ExampleIMKInputController_Menu demonstrates using Menu on a IMKInputController instance.
// Returns a menu of commands that are specific to an input method.
func ExampleIMKInputController_Menu() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.Menu()
	// Output:
	}

// ExampleIMKInputController_ReplacementRange demonstrates using ReplacementRange on a IMKInputController instance.
// Returns the range in the client document that the text should replace.
func ExampleIMKInputController_ReplacementRange() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.ReplacementRange()
	// Output:
	}

// ExampleIMKInputController_SelectionRange demonstrates using SelectionRange on a IMKInputController instance.
// Returns where the range of the selection that should be placed inside marked text.
func ExampleIMKInputController_SelectionRange() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.SelectionRange()
	// Output:
	}

// ExampleIMKInputController_Server demonstrates using Server on a IMKInputController instance.
// Returns the server object that manages the input controller.
func ExampleIMKInputController_Server() {
	obj := inputmethodkit.NewIMKInputController()
	_ = obj.Server()
	// Output:
	}

// ExampleIMKInputController_UpdateComposition demonstrates using UpdateComposition on a IMKInputController instance.
// Informs the input controller that the composition has changed.
func ExampleIMKInputController_UpdateComposition() {
	obj := inputmethodkit.NewIMKInputController()
	obj.UpdateComposition()
	// Output:
	}

