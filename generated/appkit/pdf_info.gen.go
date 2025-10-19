// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFInfo] class.
var (
	pDFInfoClass     _PDFInfoClass
	pDFInfoClassOnce sync.Once
)

func getPDFInfoClass() _PDFInfoClass {
	pDFInfoClassOnce.Do(func() {
		pDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}
	})
	return pDFInfoClass
}

type _PDFInfoClass struct {
	class objc.Class
}

// An interface definition for the [PDFInfo] class.
type IPDFInfo interface {
	objectivec.IObject
}

// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo

type PDFInfo struct {
	objectivec.Object
}

// PDFInfoFrom constructs a [PDFInfo] from an unsafe.Pointer.
//
// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size.
func PDFInfoFrom(ptr unsafe.Pointer) PDFInfo {
	return PDFInfo{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PDFInfoClass) Alloc() PDFInfo {
	rv := objc.Send[PDFInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PDFInfoClass) New() PDFInfo {
	rv := objc.Send[PDFInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFInfo) Init() PDFInfo {
	rv := objc.Send[PDFInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFInfo) Autorelease() PDFInfo {
	rv := objc.Send[PDFInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFInfo creates a new PDFInfo instance.
func NewPDFInfo() PDFInfo {
	return getPDFInfoClass().New()
}




