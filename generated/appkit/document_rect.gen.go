
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [documentRect] class.
var documentRectClass _documentRectClass

func init() {
	documentRectClass = _documentRectClass{objc.GetClass("documentRect")}
}

type _documentRectClass struct {
	objc.Class
}

// An interface definition for the [documentRect] class.
type IdocumentRect interface {
	ID() objc.ID
}

type documentRect struct {
	id objc.ID
}

func documentRectFrom(ptr unsafe.Pointer) documentRect {
	return documentRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ documentRect) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _documentRectClass) Alloc() documentRect {
	rv := objc.Send[documentRect](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _documentRectClass) New() documentRect {
	rv := objc.Send[documentRect](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdocumentRect creates and returns a new initialized instance.
func NewdocumentRect() documentRect {
	return documentRectClass.New()
}

// Init initializes the instance.
func (d_ documentRect) Init() documentRect {
	rv := objc.Send[documentRect](d_.ID(), selInit)
	return rv
}
