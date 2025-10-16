
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [endDocument] class.
var endDocumentClass _endDocumentClass

func init() {
	endDocumentClass = _endDocumentClass{objc.GetClass("endDocument")}
}

type _endDocumentClass struct {
	objc.Class
}

// An interface definition for the [endDocument] class.
type IendDocument interface {
	ID() objc.ID
}

type endDocument struct {
	id objc.ID
}

func endDocumentFrom(ptr unsafe.Pointer) endDocument {
	return endDocument{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ endDocument) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _endDocumentClass) Alloc() endDocument {
	rv := objc.Send[endDocument](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _endDocumentClass) New() endDocument {
	rv := objc.Send[endDocument](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewendDocument creates and returns a new initialized instance.
func NewendDocument() endDocument {
	return endDocumentClass.New()
}

// Init initializes the instance.
func (e_ endDocument) Init() endDocument {
	rv := objc.Send[endDocument](e_.ID(), selInit)
	return rv
}
