// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPrintInfo

// ExampleNewPrintInfo demonstrates how to create a PrintInfo instance.
// Creates a printing information object.
func ExampleNewPrintInfo() {
	_ = appkit.NewPrintInfo()
	// Output:
}
// ExamplePrintInfo_Dictionary demonstrates using Dictionary on a PrintInfo instance.
// Returns the print info’s dictionary that contains the printing attributes.
func ExamplePrintInfo_Dictionary() {
	obj := appkit.NewPrintInfo()
	_ = obj.Dictionary()
	// Output:
	}

// ExamplePrintInfo_PMPageFormat demonstrates using PMPageFormat on a PrintInfo instance.
// Returns a Core Printing object configured with the print info’s page format information.
func ExamplePrintInfo_PMPageFormat() {
	obj := appkit.NewPrintInfo()
	obj.PMPageFormat()
	// Output:
	}

// ExamplePrintInfo_PMPrintSession demonstrates using PMPrintSession on a PrintInfo instance.
// Returns a Core Printing object configured with the print info’s session information.
func ExamplePrintInfo_PMPrintSession() {
	obj := appkit.NewPrintInfo()
	obj.PMPrintSession()
	// Output:
	}

// ExamplePrintInfo_PMPrintSettings demonstrates using PMPrintSettings on a PrintInfo instance.
// Returns a Core Printing object configured with the print info’s print settings information
func ExamplePrintInfo_PMPrintSettings() {
	obj := appkit.NewPrintInfo()
	obj.PMPrintSettings()
	// Output:
	}

// ExamplePrintInfo_SetUpPrintOperationDefaultValues demonstrates using SetUpPrintOperationDefaultValues on a PrintInfo instance.
// Validates the attributes encapsulated by the print info.
func ExamplePrintInfo_SetUpPrintOperationDefaultValues() {
	obj := appkit.NewPrintInfo()
	obj.SetUpPrintOperationDefaultValues()
	// Output:
	}

// ExamplePrintInfo_UpdateFromPMPageFormat demonstrates using UpdateFromPMPageFormat on a PrintInfo instance.
// Synchronizes the print info’s page format information with information from its associated page format object.
func ExamplePrintInfo_UpdateFromPMPageFormat() {
	obj := appkit.NewPrintInfo()
	obj.UpdateFromPMPageFormat()
	// Output:
	}

// ExamplePrintInfo_UpdateFromPMPrintSettings demonstrates using UpdateFromPMPrintSettings on a PrintInfo instance.
// Synchronizes the print info’s print settings information with information from its associated print settings object.
func ExamplePrintInfo_UpdateFromPMPrintSettings() {
	obj := appkit.NewPrintInfo()
	obj.UpdateFromPMPrintSettings()
	// Output:
	}

