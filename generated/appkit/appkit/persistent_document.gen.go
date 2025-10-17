// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PersistentDocument] class.
var PersistentDocumentClass objc.Class

func init() {
	PersistentDocumentClass = objc.GetClass("NSPersistentDocument")
}

type PersistentDocument struct {
	objc.ID
}

func PersistentDocumentFrom(ptr unsafe.Pointer) PersistentDocument {
	return PersistentDocument{
		ID: objc.ID(ptr),
	}
}




