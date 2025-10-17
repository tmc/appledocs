
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PDFInfo] class.
var PDFInfoClass _PDFInfoClass

func init() {
	PDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}
}

type _PDFInfoClass struct {
	objc.Class
}

// An interface definition for the [PDFInfo] class.
type IPDFInfo interface {
	ID() objc.ID
}

type PDFInfo struct {
	id objc.ID
}

func PDFInfoFrom(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PDFInfo) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PDFInfoClass) Alloc() PDFInfo {
	rv := objc.Send[PDFInfo](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PDFInfoClass) New() PDFInfo {
	rv := objc.Send[PDFInfo](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPDFInfo creates and returns a new initialized instance.
func NewPDFInfo() PDFInfo {
	return PDFInfoClass.New()
}

// Init initializes the instance.
func (p_ PDFInfo) Init() PDFInfo {
	rv := objc.Send[PDFInfo](p_.ID(), selInit)
	return rv
}
