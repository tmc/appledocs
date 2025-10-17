// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentDocument] class.
var persistentDocumentClass = _PersistentDocumentClass{objc.GetClass("NSPersistentDocument")}

type _PersistentDocumentClass struct {
	class objc.Class
}

// A document object that can integrate with Core Data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPersistentDocument

type PersistentDocument struct {
	Document
}

// PersistentDocumentFrom constructs a [PersistentDocument] from an unsafe.Pointer.
//
// A document object that can integrate with Core Data.
func PersistentDocumentFrom(ptr unsafe.Pointer) PersistentDocument {
	return PersistentDocument{
		Document: DocumentFrom(ptr),
	}
}



