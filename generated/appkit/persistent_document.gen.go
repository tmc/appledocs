// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentDocument] class.
var (
	persistentDocumentClass     _PersistentDocumentClass
	persistentDocumentClassOnce sync.Once
)

func getPersistentDocumentClass() _PersistentDocumentClass {
	persistentDocumentClassOnce.Do(func() {
		persistentDocumentClass = _PersistentDocumentClass{objc.GetClass("NSPersistentDocument")}
	})
	return persistentDocumentClass
}

type _PersistentDocumentClass struct {
	class objc.Class
}

// An interface definition for the [PersistentDocument] class.
type IPersistentDocument interface {
	IDocument
}

// A document object that can integrate with Core Data.
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

// Alloc allocates a new instance without initialization.
func (pc _PersistentDocumentClass) Alloc() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentDocumentClass) New() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentDocument) Init() PersistentDocument {
	rv := objc.Send[PersistentDocument](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentDocument) Autorelease() PersistentDocument {
	rv := objc.Send[PersistentDocument](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentDocument creates a new PersistentDocument instance.
func NewPersistentDocument() PersistentDocument {
	return getPersistentDocumentClass().New()
}




