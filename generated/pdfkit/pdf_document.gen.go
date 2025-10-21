// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFDocument] class.
var (
	PDFDocumentClass     _PDFDocumentClass
	PDFDocumentClassOnce sync.Once
)

func getPDFDocumentClass() _PDFDocumentClass {
	PDFDocumentClassOnce.Do(func() {
		PDFDocumentClass = _PDFDocumentClass{objc.GetClass("PDFDocument")}
	})
	return PDFDocumentClass
}

type _PDFDocumentClass struct {
	class objc.Class
}

// An interface definition for the [PDFDocument] class.
type IPDFDocument interface {
	objectivec.IObject
	BeginFindStringWithOptions(string_ appkit.string, options unsafe.Pointer)
	FindStringWithOptions(string_ appkit.string, options unsafe.Pointer) []PDFSelection
	SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint foundation.IPoint, endPage IPDFPage, endPoint foundation.IPoint, granularity IPDFSelectionGranularity) PDFSelection
	UnlockWithPassword(password appkit.string) bool
}

// An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
//
// The other utility classes are either instantiated from methods in , as are and ; or support it, as do and . You initialize a object with PDF data or with a URL to a PDF file. You can then ask for the page count, add or delete pages, perform a find, or parse selected content into an object.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument
type PDFDocument struct {
	objectivec.Object
}

// PDFDocumentFrom constructs a [PDFDocument] from an unsafe.Pointer.
//
// An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
func PDFDocumentFrom(ptr unsafe.Pointer) PDFDocument {
	return PDFDocument{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFDocumentClass) Alloc() PDFDocument {
	rv := objc.Send[PDFDocument](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFDocumentClass) New() PDFDocument {
	rv := objc.Send[PDFDocument](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFDocument) Init() PDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFDocument) Autorelease() PDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDocument creates a new PDFDocument instance.
func NewPDFDocument() PDFDocument {
	return getPDFDocumentClass().New()
}




// Initializes a object with the passed-in data.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(data:)
func NewPDFDocumentWithData(data foundation.IData) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[PDFDocument](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}



// Initializes a object with the contents at the specified URL (if the URL is invalid, this method returns ).
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(url:)
func NewPDFDocumentWithURL(url foundation.IURL) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[PDFDocument](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}


// Asynchronously finds all instances of the specified string in the document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/beginFindString(_:withOptions:)
func (p_ PDFDocument) BeginFindStringWithOptions(string_ appkit.string, options unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginFindString:withOptions:"), string_, options)
}

// Synchronously finds all instances of the specified string in the document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/findString(_:withOptions:)
func (p_ PDFDocument) FindStringWithOptions(string_ appkit.string, options unsafe.Pointer) []PDFSelection {
	rv := objc.Send[[]PDFSelection](p_.ID, objc.Sel("findString:withOptions:"), string_, options)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:at:to:at:with:)
func (p_ PDFDocument) SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint foundation.IPoint, endPage IPDFPage, endPoint foundation.IPoint, granularity IPDFSelectionGranularity) PDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPage:atPoint:toPage:atPoint:withGranularity:"), startPage, startPoint, endPage, endPoint, granularity)
	return rv
}

// Attempts to unlock an encrypted document.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/unlock(withPassword:)
func (p_ PDFDocument) UnlockWithPassword(password appkit.string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("unlockWithPassword:"), password)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/accessPermissions
func (p_ PDFDocument) AccessPermissions() PDFAccessPermissions {
	rv := objc.Send[PDFAccessPermissions](p_.ID, objc.Sel("accessPermissions"))
	return rv
}

// The object acting as the delegate for the object.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/delegate
func (p_ PDFDocument) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object acting as the delegate for the object.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/delegate
func (p_ PDFDocument) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}


