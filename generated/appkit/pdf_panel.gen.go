
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFPanel] class.
var PDFPanelClass _PDFPanelClass

func init() {
	PDFPanelClass = _PDFPanelClass{objc.GetClass("NSPDFPanel")}
}

type _PDFPanelClass struct {
	objc.Class
}

// An interface definition for the [PDFPanel] class.
type IPDFPanel interface {
	ID() objc.ID
}

type PDFPanel struct {
	id objc.ID
}

func PDFPanelFrom(ptr unsafe.Pointer) PDFPanel {
	return PDFPanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PDFPanel) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PDFPanelClass) Alloc() PDFPanel {
	rv := objc.Send[PDFPanel](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PDFPanelClass) New() PDFPanel {
	rv := objc.Send[PDFPanel](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPDFPanel creates and returns a new initialized instance.
func NewPDFPanel() PDFPanel {
	return PDFPanelClass.New()
}

// Init initializes the instance.
func (p_ PDFPanel) Init() PDFPanel {
	rv := objc.Send[PDFPanel](p_.ID(), selInit)
	return rv
}
