
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [documentCursor] class.
var documentCursorClass _documentCursorClass

func init() {
	documentCursorClass = _documentCursorClass{objc.GetClass("documentCursor")}
}

type _documentCursorClass struct {
	objc.Class
}

// An interface definition for the [documentCursor] class.
type IdocumentCursor interface {
	ID() objc.ID
}

type documentCursor struct {
	id objc.ID
}

func documentCursorFrom(ptr unsafe.Pointer) documentCursor {
	return documentCursor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ documentCursor) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _documentCursorClass) Alloc() documentCursor {
	rv := objc.Send[documentCursor](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _documentCursorClass) New() documentCursor {
	rv := objc.Send[documentCursor](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdocumentCursor creates and returns a new initialized instance.
func NewdocumentCursor() documentCursor {
	return documentCursorClass.New()
}

// Init initializes the instance.
func (d_ documentCursor) Init() documentCursor {
	rv := objc.Send[documentCursor](d_.ID(), selInit)
	return rv
}
