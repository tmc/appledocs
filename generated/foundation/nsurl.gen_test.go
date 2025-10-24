// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURL

// ExampleURL_FileReferenceURL demonstrates using FileReferenceURL on a URL instance.
// Returns a new file reference URL that points to the same resource as the receiver.
//
// Note: This example is not executed because FileReferenceURL crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleURL_FileReferenceURL() {
	obj := foundation.NewURL()
	_ = obj.FileReferenceURL()
	}

// ExampleURL_IsFileReferenceURL demonstrates using IsFileReferenceURL on a URL instance.
// Returns whether the URL is a file reference URL.
func ExampleURL_IsFileReferenceURL() {
	obj := foundation.NewURL()
	_ = obj.IsFileReferenceURL()
	// Output:
	}

// ExampleURL_RemoveAllCachedResourceValues demonstrates using RemoveAllCachedResourceValues on a URL instance.
// Removes all cached resource values and temporary resource values from the URL object.
func ExampleURL_RemoveAllCachedResourceValues() {
	obj := foundation.NewURL()
	obj.RemoveAllCachedResourceValues()
	// Output:
	}

// ExampleURL_StartAccessingSecurityScopedResource demonstrates using StartAccessingSecurityScopedResource on a URL instance.
// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
func ExampleURL_StartAccessingSecurityScopedResource() {
	obj := foundation.NewURL()
	_ = obj.StartAccessingSecurityScopedResource()
	// Output:
	}

// ExampleURL_StopAccessingSecurityScopedResource demonstrates using StopAccessingSecurityScopedResource on a URL instance.
// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
func ExampleURL_StopAccessingSecurityScopedResource() {
	obj := foundation.NewURL()
	obj.StopAccessingSecurityScopedResource()
	// Output:
	}

