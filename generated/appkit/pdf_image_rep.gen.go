
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFImageRep] class.
var PDFImageRepClass _PDFImageRepClass

func init() {
	PDFImageRepClass = _PDFImageRepClass{objc.GetClass("NSPDFImageRep")}
}

type _PDFImageRepClass struct {
	objc.Class
}

// An interface definition for the [PDFImageRep] class.
type IPDFImageRep interface {
	ID() objc.ID
}

type PDFImageRep struct {
	id objc.ID
}

func PDFImageRepFrom(ptr unsafe.Pointer) PDFImageRep {
	return PDFImageRep{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PDFImageRep) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PDFImageRepClass) Alloc() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PDFImageRepClass) New() PDFImageRep {
	rv := objc.Send[PDFImageRep](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPDFImageRep creates and returns a new initialized instance.
func NewPDFImageRep() PDFImageRep {
	return PDFImageRepClass.New()
}

// Init initializes the instance.
func (p_ PDFImageRep) Init() PDFImageRep {
	rv := objc.Send[PDFImageRep](p_.ID(), selInit)
	return rv
}
