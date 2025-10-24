// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationPopup] class.
var (
	PDFAnnotationPopupClass     _PDFAnnotationPopupClass
	PDFAnnotationPopupClassOnce sync.Once
)

func getPDFAnnotationPopupClass() _PDFAnnotationPopupClass {
	PDFAnnotationPopupClassOnce.Do(func() {
		PDFAnnotationPopupClass = _PDFAnnotationPopupClass{objc.GetClass("PDFAnnotationPopup")}
	})
	return PDFAnnotationPopupClass
}

type _PDFAnnotationPopupClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationPopup] class.
type IPDFAnnotationPopup interface {
	IPDFAnnotation
	// properties:
	// methods:
}

// A object provides user interactivity on a PDF page in the form of a pop-up menu.


// A object provides user interactivity on a PDF page in the form of a pop-up menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationPopup
type PDFAnnotationPopup struct {
	PDFAnnotation
}

// PDFAnnotationPopupFrom constructs a [PDFAnnotationPopup] from an unsafe.Pointer.
//
// A object provides user interactivity on a PDF page in the form of a pop-up menu.
func PDFAnnotationPopupFrom(ptr unsafe.Pointer) PDFAnnotationPopup {
	return PDFAnnotationPopup{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationPopupClass) Alloc() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationPopupClass) New() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationPopup) Init() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationPopup) Autorelease() PDFAnnotationPopup {
	rv := objc.Send[PDFAnnotationPopup](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationPopup creates a new PDFAnnotationPopup instance.
func NewPDFAnnotationPopup() PDFAnnotationPopup {
	return getPDFAnnotationPopupClass().New()
}




