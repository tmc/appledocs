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

// The class instance for the [PDFAction] class.
var (
	PDFActionClass     _PDFActionClass
	PDFActionClassOnce sync.Once
)

func getPDFActionClass() _PDFActionClass {
	PDFActionClassOnce.Do(func() {
		PDFActionClass = _PDFActionClass{objc.GetClass("PDFAction")}
	})
	return PDFActionClass
}

type _PDFActionClass struct {
	class objc.Class
}

// An interface definition for the [PDFAction] class.
type IPDFAction interface {
	objectivec.IObject
}

// An action that is performed when, for example, a PDF annotation is activated or an outline item is clicked.
//
// A object represents an action associated with a PDF element, such as an annotation or a link, that the viewer application can perform. See the Adobe PDF Specification for more about actions and action types. is an abstract superclass of the following concrete classes:
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAction
type PDFAction struct {
	objectivec.Object
}

// PDFActionFrom constructs a [PDFAction] from an unsafe.Pointer.
//
// An action that is performed when, for example, a PDF annotation is activated or an outline item is clicked.
func PDFActionFrom(ptr unsafe.Pointer) PDFAction {
	return PDFAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionClass) Alloc() PDFAction {
	rv := objc.Send[PDFAction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionClass) New() PDFAction {
	rv := objc.Send[PDFAction](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAction) Init() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAction) Autorelease() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAction creates a new PDFAction instance.
func NewPDFAction() PDFAction {
	return getPDFActionClass().New()
}


// Returns the type of the action.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAction/type
func (p_ PDFAction) Type() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("type"))
	return rv
}

// An object that represents an action for a PDF element, such as a link annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFAction) Action() PDFAction {
	rv := objc.Send[PDFAction](p_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
// An object that represents an action for a PDF element, such as a link annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/action
func (p_ PDFAction) SetAction(value IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}

// Returns the modification date of the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFAction) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// SetModificationDate sets the value of the modificationDate property.
// Returns the modification date of the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/modificationdate
func (p_ PDFAction) SetModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}

// Returns the page that the annotation is associated with.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFAction) Page() PDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("page"))
	return rv
}


// SetPage sets the value of the page property.
// Returns the page that the annotation is associated with.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/page
func (p_ PDFAction) SetPage(value IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPage:"), value)
}

// Returns the name of the user who created the annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFAction) UserName() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("userName"))
	return rv
}


// SetUserName sets the value of the userName property.
// Returns the name of the user who created the annotation.

//
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfannotation/username
func (p_ PDFAction) SetUserName(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserName:"), value)
}



