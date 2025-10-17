// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ItemProvider] class.
var ItemProviderClass objc.Class

func init() {
	ItemProviderClass = objc.GetClass("NSItemProvider")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSItemProvider/canLoadObject(ofClass:)-3eig9
func (i_ ItemProvider) CanLoadObjectOfClass(aClass unsafe.Pointer) bool {
	sel := objc.RegisterName("canLoadObjectOfClass:")
	ret := i_.ID.Send(sel, aClass)
	return ret != 0
}
// Registers a file-backed representation for an item, specifying file options, item visibility, and a load handler. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/NSItemProvider/registerFileRepresentation(forTypeIdentifier:fileOptions:visibility:loadHandler:)
func (i_ ItemProvider) RegisterFileRepresentationForTypeIdentifierFileOptionsVisibilityLoadHandler(typeIdentifier string, fileOptions unsafe.Pointer, visibility unsafe.Pointer, loadHandler unsafe.Pointer) {
	sel := objc.RegisterName("registerFileRepresentationForTypeIdentifier:fileOptions:visibility:loadHandler:")
	i_.ID.Send(sel, typeIdentifier, fileOptions, visibility, loadHandler)
}

