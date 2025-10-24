// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface_test

import (
	"github.com/tmc/appledocs/generated/securityinterface"
)

// Suppress unused import errors
var _ = securityinterface.NewSFAuthorizationPluginView

// ExampleSFAuthorizationPluginView_Callbacks demonstrates using Callbacks on a SFAuthorizationPluginView instance.
// Returns the authorization callbacks structure with which this instance was initialized.
func ExampleSFAuthorizationPluginView_Callbacks() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.Callbacks()
	// Output:
	}

// ExampleSFAuthorizationPluginView_DidActivate demonstrates using DidActivate on a SFAuthorizationPluginView instance.
// Tells the authorization plug-in when its user interface has become active.
func ExampleSFAuthorizationPluginView_DidActivate() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	obj.DidActivate()
	// Output:
	}

// ExampleSFAuthorizationPluginView_DidDeactivate demonstrates using DidDeactivate on a SFAuthorizationPluginView instance.
// Tells the authorization plug-in that its user interface has been deactivated.
func ExampleSFAuthorizationPluginView_DidDeactivate() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	obj.DidDeactivate()
	// Output:
	}

// ExampleSFAuthorizationPluginView_DisplayView demonstrates using DisplayView on a SFAuthorizationPluginView instance.
// Displays the user interface provided by the authorization plug-in view subclass.
func ExampleSFAuthorizationPluginView_DisplayView() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	obj.DisplayView()
	// Output:
	}

// ExampleSFAuthorizationPluginView_EngineRef demonstrates using EngineRef on a SFAuthorizationPluginView instance.
// Returns the authorization engine handle with which this instance was initialized.
func ExampleSFAuthorizationPluginView_EngineRef() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.EngineRef()
	// Output:
	}

// ExampleSFAuthorizationPluginView_FirstKeyView demonstrates using FirstKeyView on a SFAuthorizationPluginView instance.
// Returns the first view in the keyboard loop of the view.
func ExampleSFAuthorizationPluginView_FirstKeyView() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.FirstKeyView()
	// Output:
	}

// ExampleSFAuthorizationPluginView_FirstResponder demonstrates using FirstResponder on a SFAuthorizationPluginView instance.
// Returns the view that should get focus for keyboard events.
func ExampleSFAuthorizationPluginView_FirstResponder() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.FirstResponder()
	// Output:
	}

// ExampleSFAuthorizationPluginView_LastError demonstrates using LastError on a SFAuthorizationPluginView instance.
// Returns the last error that occurred during evaluation.
func ExampleSFAuthorizationPluginView_LastError() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.LastError()
	// Output:
	}

// ExampleSFAuthorizationPluginView_LastKeyView demonstrates using LastKeyView on a SFAuthorizationPluginView instance.
// Returns the last view in the keyboard loop of the view.
func ExampleSFAuthorizationPluginView_LastKeyView() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	_ = obj.LastKeyView()
	// Output:
	}

// ExampleSFAuthorizationPluginView_UpdateView demonstrates using UpdateView on a SFAuthorizationPluginView instance.
// Tells the authorization plug-in to get and display the appropriate view in the authorization plug-in’s user interface.
func ExampleSFAuthorizationPluginView_UpdateView() {
	obj := securityinterface.NewSFAuthorizationPluginView()
	obj.UpdateView()
	// Output:
	}

