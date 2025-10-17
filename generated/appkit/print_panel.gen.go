// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintPanel] class.
var printPanelClass = _PrintPanelClass{objc.GetClass("NSPrintPanel")}

type _PrintPanelClass struct {
	class objc.Class
}

// An interface definition for the [PrintPanel] class.
type IPrintPanel interface {
	objectivec.IObject
}

// The Print panel that queries the user for information about a print job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel

type PrintPanel struct {
	objectivec.Object
}

// PrintPanelFrom constructs a [PrintPanel] from an unsafe.Pointer.
//
// The Print panel that queries the user for information about a print job.
func PrintPanelFrom(ptr unsafe.Pointer) PrintPanel {
	return PrintPanel{objectivec.Object{objc.ID(ptr)}}
}



