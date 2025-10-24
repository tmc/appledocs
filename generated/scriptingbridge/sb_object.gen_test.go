// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge_test

import (
	"github.com/tmc/appledocs/generated/scriptingbridge"
)

// Suppress unused import errors
var _ = scriptingbridge.NewSBObject

// ExampleNewSBObject demonstrates how to create a SBObject instance.
// Initializes and returns an instance of an   subclass.
func ExampleNewSBObject() {
	_ = scriptingbridge.NewSBObject()
	// Output:
}
// ExampleSBObject_Get demonstrates using Get on a SBObject instance.
// Forces evaluation of the receiver, causing the real object to be returned   immediately.
func ExampleSBObject_Get() {
	obj := scriptingbridge.NewSBObject()
	_ = obj.Get()
	// Output:
	}

// ExampleSBObject_LastError demonstrates using LastError on a SBObject instance.
// The error from the last event this object sent, or nil if it succeeded.
func ExampleSBObject_LastError() {
	obj := scriptingbridge.NewSBObject()
	_ = obj.LastError()
	// Output:
	}


