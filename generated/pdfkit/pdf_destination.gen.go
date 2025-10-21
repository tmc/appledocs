// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PDFDestination] class.
var (
	PDFDestinationClass     _PDFDestinationClass
	PDFDestinationClassOnce sync.Once
)

func getPDFDestinationClass() _PDFDestinationClass {
	PDFDestinationClassOnce.Do(func() {
		PDFDestinationClass = _PDFDestinationClass{objc.GetClass("PDFDestination")}
	})
	return PDFDestinationClass
}

type _PDFDestinationClass struct {
	class objc.Class
}

// An interface definition for the [PDFDestination] class.
type IPDFDestination interface {
	objectivec.IObject
}

// A object describes a point on a PDF page.
//
// In typical usage, you do not initialize objects but rather get them as either attributes of or objects, or in response to the method .
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination
type PDFDestination struct {
	objectivec.Object
}

// PDFDestinationFrom constructs a [PDFDestination] from an unsafe.Pointer.
//
// A object describes a point on a PDF page.
func PDFDestinationFrom(ptr unsafe.Pointer) PDFDestination {
	return PDFDestination{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFDestinationClass) Alloc() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFDestinationClass) New() PDFDestination {
	rv := objc.Send[PDFDestination](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFDestination) Init() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFDestination) Autorelease() PDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDestination creates a new PDFDestination instance.
func NewPDFDestination() PDFDestination {
	return getPDFDestinationClass().New()
}


// Returns the page that the destination refers to.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDestination/page
func (p_ PDFDestination) Page() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("page"))
	return rv
}



