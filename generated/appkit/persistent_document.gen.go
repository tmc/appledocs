
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PersistentDocument] class.
var PersistentDocumentClass _PersistentDocumentClass

func init() {
	PersistentDocumentClass = _PersistentDocumentClass{objc.GetClass("NSPersistentDocument")}
}

type _PersistentDocumentClass struct {
	objc.Class
}

// An interface definition for the [PersistentDocument] class.
type IPersistentDocument interface {
	ID() objc.ID
}

type PersistentDocument struct {
	id objc.ID
}

func PersistentDocumentFrom(ptr unsafe.Pointer) PersistentDocument {
	return PersistentDocument{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PersistentDocument) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentDocumentClass) Alloc() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PersistentDocumentClass) New() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPersistentDocument creates and returns a new initialized instance.
func NewPersistentDocument() PersistentDocument {
	return PersistentDocumentClass.New()
}

// Init initializes the instance.
func (p_ PersistentDocument) Init() PersistentDocument {
	rv := objc.Send[PersistentDocument](p_.ID(), selInit)
	return rv
}
