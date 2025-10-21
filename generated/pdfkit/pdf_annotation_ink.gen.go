// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFAnnotationInk] class.
var (
	PDFAnnotationInkClass     _PDFAnnotationInkClass
	PDFAnnotationInkClassOnce sync.Once
)

func getPDFAnnotationInkClass() _PDFAnnotationInkClass {
	PDFAnnotationInkClassOnce.Do(func() {
		PDFAnnotationInkClass = _PDFAnnotationInkClass{objc.GetClass("PDFAnnotationInk")}
	})
	return PDFAnnotationInkClass
}

type _PDFAnnotationInkClass struct {
	class objc.Class
}

// An interface definition for the [PDFAnnotationInk] class.
type IPDFAnnotationInk interface {
	IPDFAnnotation
	AddBezierPath(path unsafe.Pointer)
	Paths() unsafe.Pointer
	RemoveBezierPath(path unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk
type PDFAnnotationInk struct {
	PDFAnnotation
}

// PDFAnnotationInkFrom constructs a [PDFAnnotationInk] from an unsafe.Pointer.
func PDFAnnotationInkFrom(ptr unsafe.Pointer) PDFAnnotationInk {
	return PDFAnnotationInk{
		PDFAnnotation: PDFAnnotationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAnnotationInkClass) Alloc() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAnnotationInkClass) New() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAnnotationInk) Init() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAnnotationInk) Autorelease() PDFAnnotationInk {
	rv := objc.Send[PDFAnnotationInk](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAnnotationInk creates a new PDFAnnotationInk instance.
func NewPDFAnnotationInk() PDFAnnotationInk {
	return getPDFAnnotationInkClass().New()
}


// Adds a Bezier path to an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/add(_:)
func (p_ PDFAnnotationInk) AddBezierPath(path unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addBezierPath:"), path)
}

// Returns an array containing the Bezier paths that make up an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/paths()
func (p_ PDFAnnotationInk) Paths() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("paths"))
	return rv
}

// Removes a Bezier path from an annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAnnotationInk/remove(_:)
func (p_ PDFAnnotationInk) RemoveBezierPath(path unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeBezierPath:"), path)
}



