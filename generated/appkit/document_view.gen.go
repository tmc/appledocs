
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [documentView] class.
var documentViewClass _documentViewClass

func init() {
	documentViewClass = _documentViewClass{objc.GetClass("documentView")}
}

type _documentViewClass struct {
	objc.Class
}

// An interface definition for the [documentView] class.
type IdocumentView interface {
	ID() objc.ID
}

type documentView struct {
	id objc.ID
}

func documentViewFrom(ptr unsafe.Pointer) documentView {
	return documentView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ documentView) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _documentViewClass) Alloc() documentView {
	rv := objc.Send[documentView](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _documentViewClass) New() documentView {
	rv := objc.Send[documentView](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdocumentView creates and returns a new initialized instance.
func NewdocumentView() documentView {
	return documentViewClass.New()
}

// Init initializes the instance.
func (d_ documentView) Init() documentView {
	rv := objc.Send[documentView](d_.ID(), selInit)
	return rv
}
