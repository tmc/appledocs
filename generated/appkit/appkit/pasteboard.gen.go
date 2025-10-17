// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Pasteboard] class.
var PasteboardClass objc.Class

func init() {
	PasteboardClass = objc.GetClass("NSPasteboard")
}

type Pasteboard struct {
	objc.ID
}

func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{
		ID: objc.ID(ptr),
	}
}


// Sets the given string as the representation for the specified type for the first item on the receiver. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string string, dataType unsafe.Pointer) bool {
	sel := objc.RegisterName("setString:forType:")
	ret := p_.ID.Send(sel, string, dataType)
	return ret != 0
}

