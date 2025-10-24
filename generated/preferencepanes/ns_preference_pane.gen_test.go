// Code generated from Apple documentation for PreferencePanes. DO NOT EDIT.

package preferencepanes_test

import (
	"github.com/tmc/appledocs/generated/preferencepanes"
)

// Suppress unused import errors
var _ = preferencepanes.NewPreferencePane

// ExamplePreferencePane_AssignMainView demonstrates using AssignMainView on a PreferencePane instance.
// Locates and assigns the preference pane’s main view from the nib file loaded by  .
func ExamplePreferencePane_AssignMainView() {
	obj := preferencepanes.NewPreferencePane()
	obj.AssignMainView()
	// Output:
	}

// ExamplePreferencePane_DidSelect demonstrates using DidSelect on a PreferencePane instance.
// Notifies the preference pane that the main app has just displayed the preference pane’s main view.
func ExamplePreferencePane_DidSelect() {
	obj := preferencepanes.NewPreferencePane()
	obj.DidSelect()
	// Output:
	}

// ExamplePreferencePane_DidUnselect demonstrates using DidUnselect on a PreferencePane instance.
// Notifies the preference pane that the main app has just stopped displaying the preference pane’s main view.
func ExamplePreferencePane_DidUnselect() {
	obj := preferencepanes.NewPreferencePane()
	obj.DidUnselect()
	// Output:
	}

// ExamplePreferencePane_LoadMainView demonstrates using LoadMainView on a PreferencePane instance.
// Loads the preference pane’s user interface into its main view.
func ExamplePreferencePane_LoadMainView() {
	obj := preferencepanes.NewPreferencePane()
	_ = obj.LoadMainView()
	// Output:
	}

// ExamplePreferencePane_MainViewDidLoad demonstrates using MainViewDidLoad on a PreferencePane instance.
// Notifies the preference pane that the main view is set up and prepared to be displayed.
func ExamplePreferencePane_MainViewDidLoad() {
	obj := preferencepanes.NewPreferencePane()
	obj.MainViewDidLoad()
	// Output:
	}

// ExamplePreferencePane_WillSelect demonstrates using WillSelect on a PreferencePane instance.
// Notifies the preference pane that the main app is about to display the preference pane’s main view.
func ExamplePreferencePane_WillSelect() {
	obj := preferencepanes.NewPreferencePane()
	obj.WillSelect()
	// Output:
	}

// ExamplePreferencePane_WillUnselect demonstrates using WillUnselect on a PreferencePane instance.
// Notifies the preference pane that the main app is about to stop displaying the preference pane’s main view.
func ExamplePreferencePane_WillUnselect() {
	obj := preferencepanes.NewPreferencePane()
	obj.WillUnselect()
	// Output:
	}




