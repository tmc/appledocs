// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintOperation] class.
var PrintOperationClass objc.Class

func init() {
	PrintOperationClass = objc.GetClass("NSPrintOperation")
}

type PrintOperation struct {
	objc.ID
}

func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{
		ID: objc.ID(ptr),
	}
}




