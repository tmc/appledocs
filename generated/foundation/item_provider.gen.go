// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ItemProvider] class.
var itemProviderClass = _ItemProviderClass{objc.GetClass("NSItemProvider")}

type _ItemProviderClass struct {
	class objc.Class
}

// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider

type ItemProvider struct {
	objectivec.Object
}

// ItemProviderFrom constructs a [ItemProvider] from an unsafe.Pointer.
//
// An item provider for conveying data or a file between processes during drag-and-drop or copy-and-paste activities, or from a host app to an app extension.
func ItemProviderFrom(ptr unsafe.Pointer) ItemProvider {
	return ItemProvider{objectivec.Object{objc.ID(ptr)}}
}

// Returns a Boolean value indicating whether an item provider can load objects of a specified class. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/canLoadObject(ofClass:)-3eig9
func (i_ ItemProvider) CanLoadObjectOfClass(aClass unsafe.Pointer) bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("canLoadObjectOfClass:"), aClass)
	return rv
}
// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSItemProvider/registerFileRepresentation(forTypeIdentifier:fileOptions:visibility:loadHandler:)
func (i_ ItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions unsafe.Pointer, visibility unsafe.Pointer, loadHandler unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:"), typeIdentifier, fileOptions, visibility, loadHandler)
}


