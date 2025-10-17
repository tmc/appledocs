// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pasteboard] class.
var pasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}

type _PasteboardClass struct {
	class objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	objectivec.IObject
	SetStringForType(string string, dataType unsafe.Pointer) bool
}

// An object that transfers data to and from the pasteboard server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard

type Pasteboard struct {
	objectivec.Object
}

// PasteboardFrom constructs a [Pasteboard] from an unsafe.Pointer.
//
// An object that transfers data to and from the pasteboard server.
func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{objectivec.Object{objc.ID(ptr)}}
}

// Sets the given string as the representation for the specified type for the first item on the receiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string string, dataType unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), string, dataType)
	return rv
}


