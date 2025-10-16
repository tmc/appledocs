
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isDocumentEdited] class.
var isDocumentEditedClass _isDocumentEditedClass

func init() {
	isDocumentEditedClass = _isDocumentEditedClass{objc.GetClass("isDocumentEdited")}
}

type _isDocumentEditedClass struct {
	objc.Class
}

// An interface definition for the [isDocumentEdited] class.
type IisDocumentEdited interface {
	ID() objc.ID
}

type isDocumentEdited struct {
	id objc.ID
}

func isDocumentEditedFrom(ptr unsafe.Pointer) isDocumentEdited {
	return isDocumentEdited{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isDocumentEdited) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isDocumentEditedClass) Alloc() isDocumentEdited {
	rv := objc.Send[isDocumentEdited](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isDocumentEditedClass) New() isDocumentEdited {
	rv := objc.Send[isDocumentEdited](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisDocumentEdited creates and returns a new initialized instance.
func NewisDocumentEdited() isDocumentEdited {
	return isDocumentEditedClass.New()
}

// Init initializes the instance.
func (i_ isDocumentEdited) Init() isDocumentEdited {
	rv := objc.Send[isDocumentEdited](i_.ID(), selInit)
	return rv
}
