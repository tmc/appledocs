// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Printer] class.
var PrinterClass objc.Class

func init() {
	PrinterClass = objc.GetClass("NSPrinter")
}

type Printer struct {
	objc.ID
}

func PrinterFrom(ptr unsafe.Pointer) Printer {
	return Printer{
		ID: objc.ID(ptr),
	}
}




