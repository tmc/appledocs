
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [documentVisibleRect] class.
var documentVisibleRectClass _documentVisibleRectClass

func init() {
	documentVisibleRectClass = _documentVisibleRectClass{objc.GetClass("documentVisibleRect")}
}

type _documentVisibleRectClass struct {
	objc.Class
}

// An interface definition for the [documentVisibleRect] class.
type IdocumentVisibleRect interface {
	ID() objc.ID
}

type documentVisibleRect struct {
	id objc.ID
}

func documentVisibleRectFrom(ptr unsafe.Pointer) documentVisibleRect {
	return documentVisibleRect{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ documentVisibleRect) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _documentVisibleRectClass) Alloc() documentVisibleRect {
	rv := objc.Send[documentVisibleRect](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _documentVisibleRectClass) New() documentVisibleRect {
	rv := objc.Send[documentVisibleRect](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdocumentVisibleRect creates and returns a new initialized instance.
func NewdocumentVisibleRect() documentVisibleRect {
	return documentVisibleRectClass.New()
}

// Init initializes the instance.
func (d_ documentVisibleRect) Init() documentVisibleRect {
	rv := objc.Send[documentVisibleRect](d_.ID(), selInit)
	return rv
}
