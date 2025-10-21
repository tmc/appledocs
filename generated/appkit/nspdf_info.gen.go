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
}

// An object that stores information associated with the creation of a PDF file, such as its URL, tag names, page orientation, and paper size.
//
// Typically, a PDF panel—that is, a panel created by an object—displays the information supplied by an object when the user wants to export content as a PDF file. A PDF panel can also update a PDF info object with information it receives from the user.
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


// A Boolean value that indicates whether the file extension should appear after the filename.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/isFileExtensionHidden
func (p_ PDFInfo) FileExtensionHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("fileExtensionHidden"))
	return rv
}


// SetFileExtensionHidden sets the value of the fileExtensionHidden property.
// A Boolean value that indicates whether the file extension should appear after the filename.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/isFileExtensionHidden
func (p_ PDFInfo) SetFileExtensionHidden(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFileExtensionHidden:"), value)
}

// The paper orientation to use when exporting content as a PDF file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/orientation
func (p_ PDFInfo) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// The paper orientation to use when exporting content as a PDF file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/orientation
func (p_ PDFInfo) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}

// An array of tag names that should be applied to the PDF file after it’s created.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/tagNames
func (p_ PDFInfo) TagNames() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("tagNames"))
	return rv
}


// SetTagNames sets the value of the tagNames property.
// An array of tag names that should be applied to the PDF file after it’s created.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPDFInfo/tagNames
func (p_ PDFInfo) SetTagNames(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setTagNames:"), nsArray)
}

// A dictionary of additional attributes that describe how to export content as a PDF file.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/attributes
func (p_ PDFInfo) Attributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// A dictionary of additional attributes that describe how to export content as a PDF file.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/attributes
func (p_ PDFInfo) SetAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttributes:"), value)
}

// A Boolean value that indicates whether the file extension should appear after the filename.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/isfileextensionhidden
func (p_ PDFInfo) IsFileExtensionHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFileExtensionHidden"))
	return rv
}


// SetIsFileExtensionHidden sets the value of the isFileExtensionHidden property.
// A Boolean value that indicates whether the file extension should appear after the filename.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/isfileextensionhidden
func (p_ PDFInfo) SetIsFileExtensionHidden(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFileExtensionHidden:"), value)
}

// The paper size to use when exporting content as a PDF file.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/papersize
func (p_ PDFInfo) PaperSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("paperSize"))
	return rv
}


// SetPaperSize sets the value of the paperSize property.
// The paper size to use when exporting content as a PDF file.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/papersize
func (p_ PDFInfo) SetPaperSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperSize:"), value)
}

// The URL identifying the location at which the PDF file will be created.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/url
func (p_ PDFInfo) Url() foundation.URL {
	rv := objc.Send[foundation.URL](p_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The URL identifying the location at which the PDF file will be created.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspdfinfo/url
func (p_ PDFInfo) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUrl:"), value)
}



