// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Document] class.
var DocumentClass objc.Class

func init() {
	DocumentClass = objc.GetClass("NSDocument")
}

type Document struct {
	objc.ID
}

func DocumentFrom(ptr unsafe.Pointer) Document {
	return Document{
		ID: objc.ID(ptr),
	}
}


// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc Document) AllowedClassesForRestorableStateKeyPath(keyPath string) unsafe.Pointer {
	sel := objc.RegisterName("allowedClassesForRestorableStateKeyPath:")
	ret := objc.ID(DocumentClass).Send(sel, keyPath)
	return unsafe.Pointer(ret)
}
// Saves the interface-related state of the document. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	sel := objc.RegisterName("encodeRestorableStateWithCoder:")
	d_.ID.Send(sel, coder)
}
// Validates the specified user interface item that the receiver manages. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item unsafe.Pointer) bool {
	sel := objc.RegisterName("validateUserInterfaceItem:")
	ret := d_.ID.Send(sel, item)
	return ret != 0
}


