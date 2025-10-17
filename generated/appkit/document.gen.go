// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Document] class.
var documentClass = _DocumentClass{objc.GetClass("NSDocument")}

type _DocumentClass struct {
	class objc.Class
}

// An interface definition for the [Document] class.
type IDocument interface {
	objectivec.IObject
	EncodeRestorableStateWithCoder(coder unsafe.Pointer)
	ValidateUserInterfaceItem(item unsafe.Pointer) bool
}

// An abstract class that defines the interface for macOS documents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument

type Document struct {
	objectivec.Object
}

// DocumentFrom constructs a [Document] from an unsafe.Pointer.
//
// An abstract class that defines the interface for macOS documents.
func DocumentFrom(ptr unsafe.Pointer) Document {
	return Document{objectivec.Object{objc.ID(ptr)}}
}

// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}
// Saves the interface-related state of the document. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}
// Validates the specified user interface item that the receiver manages. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}


