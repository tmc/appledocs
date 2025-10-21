// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PDFActionURL] class.
var (
	PDFActionURLClass     _PDFActionURLClass
	PDFActionURLClassOnce sync.Once
)

func getPDFActionURLClass() _PDFActionURLClass {
	PDFActionURLClassOnce.Do(func() {
		PDFActionURLClass = _PDFActionURLClass{objc.GetClass("PDFActionURL")}
	})
	return PDFActionURLClass
}

type _PDFActionURLClass struct {
	class objc.Class
}

// An interface definition for the [PDFActionURL] class.
type IPDFActionURL interface {
	IPDFAction
}

// , a subclass of , defines methods for getting and setting the URL associated with a URL action.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionURL
type PDFActionURL struct {
	PDFAction
}

// PDFActionURLFrom constructs a [PDFActionURL] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and setting the URL associated with a URL action.
func PDFActionURLFrom(ptr unsafe.Pointer) PDFActionURL {
	return PDFActionURL{
		PDFAction: PDFActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionURLClass) Alloc() PDFActionURL {
	rv := objc.Send[PDFActionURL](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionURLClass) New() PDFActionURL {
	rv := objc.Send[PDFActionURL](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionURL) Init() PDFActionURL {
	rv := objc.Send[PDFActionURL](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionURL) Autorelease() PDFActionURL {
	rv := objc.Send[PDFActionURL](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionURL creates a new PDFActionURL instance.
func NewPDFActionURL() PDFActionURL {
	return getPDFActionURLClass().New()
}


// Returns the URL associated with the URL action.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionurl/url
func (p_ PDFActionURL) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// Returns the URL associated with the URL action.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfactionurl/url
func (p_ PDFActionURL) SetUrl(value foundation.IURL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



