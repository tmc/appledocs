// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PDFInfo] class.
var (
	PDFInfoClass     _PDFInfoClass
	PDFInfoClassOnce sync.Once
)

func getPDFInfoClass() _PDFInfoClass {
	PDFInfoClassOnce.Do(func() {
		PDFInfoClass = _PDFInfoClass{objc.GetClass("NSPDFInfo")}
	})
	return PDFInfoClass
}

type _PDFInfoClass struct {
	class objc.Class
}

// An interface definition for the [PDFInfo] class.
type IPDFInfo interface {
	objectivec.IObject
	// properties:
	Attributes() MutableDictionary /* not a class type */
	SetAttributes(value MutableDictionary /* not a class type */)
	IsFileExtensionHidden() bool /* primitive/slice/pointer. */
	SetIsFileExtensionHidden(value bool /* primitive/slice/pointer. */)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperSize() coregraphics.CGSize
	SetPaperSize(value coregraphics.CGSize)
	TagNames() string /* primitive/slice/pointer. */
	SetTagNames(value string /* primitive/slice/pointer. */)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size.
//
// Typically, a PDF panel—that is, a panel created by an object—displays the information supplied by an object when the user wants to export content as a PDF file. A PDF panel can also update a PDF info object with information it receives from the user.


// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A dictionary of additional attributes that describe how to export content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/attributes
func (p_ PDFInfo) Attributes() MutableDictionary /* not a class type */ {
	rv := objc.Send[MutableDictionary](p_.ID, objc.Sel("attributes"))
	return rv
}


// A dictionary of additional attributes that describe how to export content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/attributes
func (p_ PDFInfo) SetAttributes(value MutableDictionary /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttributes:"), value)
}


// A Boolean value that indicates whether the file extension should appear after the filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/isfileextensionhidden
func (p_ PDFInfo) IsFileExtensionHidden() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFileExtensionHidden"))
	return rv
}


// A Boolean value that indicates whether the file extension should appear after the filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/isfileextensionhidden
func (p_ PDFInfo) SetIsFileExtensionHidden(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFileExtensionHidden:"), value)
}


// The paper orientation to use when exporting content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/orientation
func (p_ PDFInfo) Orientation() PaperOrientation {
	rv := objc.Send[PaperOrientation](p_.ID, objc.Sel("orientation"))
	return rv
}


// The paper orientation to use when exporting content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/orientation
func (p_ PDFInfo) SetOrientation(value PaperOrientation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}


// The paper size to use when exporting content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/papersize
func (p_ PDFInfo) PaperSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("paperSize"))
	return rv
}


// The paper size to use when exporting content as a PDF file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/papersize
func (p_ PDFInfo) SetPaperSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperSize:"), value)
}


// An array of tag names that should be applied to the PDF file after it’s created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/tagnames
func (p_ PDFInfo) TagNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("tagNames"))
	return rv
}


// An array of tag names that should be applied to the PDF file after it’s created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/tagnames
func (p_ PDFInfo) SetTagNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTagNames:"), objc.String(value))
}


// The URL identifying the location at which the PDF file will be created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/url
func (p_ PDFInfo) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// The URL identifying the location at which the PDF file will be created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/url
func (p_ PDFInfo) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



