// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFActionGoTo] class.
var (
	PDFActionGoToClass     _PDFActionGoToClass
	PDFActionGoToClassOnce sync.Once
)

func getPDFActionGoToClass() _PDFActionGoToClass {
	PDFActionGoToClassOnce.Do(func() {
		PDFActionGoToClass = _PDFActionGoToClass{objc.GetClass("PDFActionGoTo")}
	})
	return PDFActionGoToClass
}

type _PDFActionGoToClass struct {
	class objc.Class
}

// An interface definition for the [PDFActionGoTo] class.
type IPDFActionGoTo interface {
	IPDFAction
}

// , a subclass of , defines methods for getting and setting the destination of a go-to action.
//
// A object represents the action of going to a specific location within the PDF document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionGoTo
type PDFActionGoTo struct {
	PDFAction
}

// PDFActionGoToFrom constructs a [PDFActionGoTo] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the destination of a go-to action.
func PDFActionGoToFrom(ptr unsafe.Pointer) PDFActionGoTo {
	return PDFActionGoTo{
		PDFAction: PDFActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionGoToClass) Alloc() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionGoToClass) New() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionGoTo) Init() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionGoTo) Autorelease() PDFActionGoTo {
	rv := objc.Send[PDFActionGoTo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionGoTo creates a new PDFActionGoTo instance.
func NewPDFActionGoTo() PDFActionGoTo {
	return getPDFActionGoToClass().New()
}




