// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ItemProviderClass _ItemProviderClass

func init() {
	ItemProviderClass = _ItemProviderClass{objc.GetClass("NSItemProvider")}
}

type _ItemProviderClass struct {
	class objc.Class
}

type ItemProvider struct {
	objc.ID
}

func ItemProviderFrom(ptr unsafe.Pointer) ItemProvider {
	return ItemProvider{
		ID: objc.ID(ptr),
	}
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


