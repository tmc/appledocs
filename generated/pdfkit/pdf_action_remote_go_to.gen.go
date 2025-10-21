// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFActionRemoteGoTo] class.
var (
	PDFActionRemoteGoToClass     _PDFActionRemoteGoToClass
	PDFActionRemoteGoToClassOnce sync.Once
)

func getPDFActionRemoteGoToClass() _PDFActionRemoteGoToClass {
	PDFActionRemoteGoToClassOnce.Do(func() {
		PDFActionRemoteGoToClass = _PDFActionRemoteGoToClass{objc.GetClass("PDFActionRemoteGoTo")}
	})
	return PDFActionRemoteGoToClass
}

type _PDFActionRemoteGoToClass struct {
	class objc.Class
}

// An interface definition for the [PDFActionRemoteGoTo] class.
type IPDFActionRemoteGoTo interface {
	IPDFAction
}

// , a subclass of , defines methods for getting and setting the destination of a go-to action that targets another document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionRemoteGoTo
type PDFActionRemoteGoTo struct {
	PDFAction
}

// PDFActionRemoteGoToFrom constructs a [PDFActionRemoteGoTo] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the destination of a go-to action that targets another document.
func PDFActionRemoteGoToFrom(ptr unsafe.Pointer) PDFActionRemoteGoTo {
	return PDFActionRemoteGoTo{
		PDFAction: PDFActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionRemoteGoToClass) Alloc() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionRemoteGoToClass) New() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionRemoteGoTo) Init() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionRemoteGoTo) Autorelease() PDFActionRemoteGoTo {
	rv := objc.Send[PDFActionRemoteGoTo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionRemoteGoTo creates a new PDFActionRemoteGoTo instance.
func NewPDFActionRemoteGoTo() PDFActionRemoteGoTo {
	return getPDFActionRemoteGoToClass().New()
}




