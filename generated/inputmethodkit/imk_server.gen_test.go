// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit_test

import (
	"github.com/tmc/appledocs/generated/inputmethodkit"
)

// Suppress unused import errors
var _ = inputmethodkit.NewIMKServer

// ExampleIMKServer_Bundle demonstrates using Bundle on a IMKServer instance.
// Returns an   object for the input method.
func ExampleIMKServer_Bundle() {
	obj := inputmethodkit.NewIMKServer()
	_ = obj.Bundle()
	// Output:
	}

// ExampleIMKServer_LastKeyEventWasDeadKey demonstrates using LastKeyEventWasDeadKey on a IMKServer instance.
func ExampleIMKServer_LastKeyEventWasDeadKey() {
	obj := inputmethodkit.NewIMKServer()
	_ = obj.LastKeyEventWasDeadKey()
	// Output:
	}

// ExampleIMKServer_PaletteWillTerminate demonstrates using PaletteWillTerminate on a IMKServer instance.
func ExampleIMKServer_PaletteWillTerminate() {
	obj := inputmethodkit.NewIMKServer()
	_ = obj.PaletteWillTerminate()
	// Output:
	}


