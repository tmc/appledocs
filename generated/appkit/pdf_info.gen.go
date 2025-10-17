// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFInfo] class.
var pDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}

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



