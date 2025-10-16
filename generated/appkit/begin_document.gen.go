
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [beginDocument] class.
var beginDocumentClass _beginDocumentClass

func init() {
	beginDocumentClass = _beginDocumentClass{objc.GetClass("beginDocument")}
}

type _beginDocumentClass struct {
	objc.Class
}

// An interface definition for the [beginDocument] class.
type IbeginDocument interface {
	ID() objc.ID
}

type beginDocument struct {
	id objc.ID
}

func beginDocumentFrom(ptr unsafe.Pointer) beginDocument {
	return beginDocument{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ beginDocument) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _beginDocumentClass) Alloc() beginDocument {
	rv := objc.Send[beginDocument](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _beginDocumentClass) New() beginDocument {
	rv := objc.Send[beginDocument](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewbeginDocument creates and returns a new initialized instance.
func NewbeginDocument() beginDocument {
	return beginDocumentClass.New()
}

// Init initializes the instance.
func (b_ beginDocument) Init() beginDocument {
	rv := objc.Send[beginDocument](b_.ID(), selInit)
	return rv
}
