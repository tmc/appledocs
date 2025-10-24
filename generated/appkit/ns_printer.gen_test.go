// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewPrinter

// ExampleNewPrinterWithType demonstrates how to create a Printer instance using NewPrinterWithType.
// Creates and returns a printer object initialized to the first available printer with the specified make and model information.
func ExampleNewPrinterWithType() {
	_ = appkit.NewPrinterWithType(
		appkit.PrinterTypeName /* typedef */{}, // type PrinterTypeName /* typedef */
	)
	// Output:
}
