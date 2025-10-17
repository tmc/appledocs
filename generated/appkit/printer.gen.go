// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Printer] class.
var printerClass = _PrinterClass{objc.GetClass("NSPrinter")}

type _PrinterClass struct {
	class objc.Class
}

// An interface definition for the [Printer] class.
type IPrinter interface {
	objectivec.IObject
}

// An object that describes a printer’s capabilities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter

type Printer struct {
	objectivec.Object
}

// PrinterFrom constructs a [Printer] from an unsafe.Pointer.
//
// An object that describes a printer’s capabilities.
func PrinterFrom(ptr unsafe.Pointer) Printer {
	return Printer{objectivec.Object{objc.ID(ptr)}}
}



