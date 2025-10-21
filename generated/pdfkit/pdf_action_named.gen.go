// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFActionNamed] class.
var (
	PDFActionNamedClass     _PDFActionNamedClass
	PDFActionNamedClassOnce sync.Once
)

func getPDFActionNamedClass() _PDFActionNamedClass {
	PDFActionNamedClassOnce.Do(func() {
		PDFActionNamedClass = _PDFActionNamedClass{objc.GetClass("PDFActionNamed")}
	})
	return PDFActionNamedClass
}

type _PDFActionNamedClass struct {
	class objc.Class
}

// An interface definition for the [PDFActionNamed] class.
type IPDFActionNamed interface {
	IPDFAction
}

// defines methods used to work with actions in PDF documents, some of which are named in the Adobe PDF Specification.
//
// A object represents an action with a defined name, such as “Go back” or “Zoom in.”
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionNamed
type PDFActionNamed struct {
	PDFAction
}

// PDFActionNamedFrom constructs a [PDFActionNamed] from an unsafe.Pointer.
//
// defines methods used to work with actions in PDF documents, some of which are named in the Adobe PDF Specification.
func PDFActionNamedFrom(ptr unsafe.Pointer) PDFActionNamed {
	return PDFActionNamed{
		PDFAction: PDFActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionNamedClass) Alloc() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionNamedClass) New() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionNamed) Init() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionNamed) Autorelease() PDFActionNamed {
	rv := objc.Send[PDFActionNamed](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionNamed creates a new PDFActionNamed instance.
func NewPDFActionNamed() PDFActionNamed {
	return getPDFActionNamedClass().New()
}




