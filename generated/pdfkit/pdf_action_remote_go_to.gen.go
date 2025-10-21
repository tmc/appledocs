// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
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


// Returns the zero-based page index referenced by the remote go-to action.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/pageindex
func (p_ PDFActionRemoteGoTo) PageIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pageIndex"))
	return rv
}


// SetPageIndex sets the value of the pageIndex property.
// Returns the zero-based page index referenced by the remote go-to action.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/pageindex
func (p_ PDFActionRemoteGoTo) SetPageIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageIndex:"), value)
}

// Sets the point, in page space, on the page referenced by the remote go-to action.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/point
func (p_ PDFActionRemoteGoTo) Point() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](p_.ID, objc.Sel("point"))
	return rv
}


// SetPoint sets the value of the point property.
// Sets the point, in page space, on the page referenced by the remote go-to action.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/point
func (p_ PDFActionRemoteGoTo) SetPoint(value coregraphics.CGPoint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPoint:"), value)
}

// Returns the URL of the document referenced by the remote go-to action.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/url
func (p_ PDFActionRemoteGoTo) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// Returns the URL of the document referenced by the remote go-to action.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionremotegoto/url
func (p_ PDFActionRemoteGoTo) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



