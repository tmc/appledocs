// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintOperation] class.
var printOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}

type _PrintOperationClass struct {
	class objc.Class
}

// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation

type PrintOperation struct {
	objectivec.Object
}

// PrintOperationFrom constructs a [PrintOperation] from an unsafe.Pointer.
//
// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{objectivec.Object{objc.ID(ptr)}}
}



