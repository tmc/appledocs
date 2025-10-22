// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PDFActionResetForm] class.
var (
	PDFActionResetFormClass     _PDFActionResetFormClass
	PDFActionResetFormClassOnce sync.Once
)

func getPDFActionResetFormClass() _PDFActionResetFormClass {
	PDFActionResetFormClassOnce.Do(func() {
		PDFActionResetFormClass = _PDFActionResetFormClass{objc.GetClass("PDFActionResetForm")}
	})
	return PDFActionResetFormClass
}

type _PDFActionResetFormClass struct {
	class objc.Class
}

// An interface definition for the [PDFActionResetForm] class.
type IPDFActionResetForm interface {
	IPDFAction
	Fields() []string
	SetFields(value []string)
	FieldsIncludedAreCleared() bool
	SetFieldsIncludedAreCleared(value bool)
}

// , a subclass of , defines methods for getting and clearing fields in a PDF form.
//
// A object represents an action associated with a PDF form.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm
type PDFActionResetForm struct {
	PDFAction
}

// PDFActionResetFormFrom constructs a [PDFActionResetForm] from an unsafe.Pointer.
//
// , a subclass of , defines methods for getting and clearing fields in a PDF form.
func PDFActionResetFormFrom(ptr unsafe.Pointer) PDFActionResetForm {
	return PDFActionResetForm{
		PDFAction: PDFActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFActionResetFormClass) Alloc() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFActionResetFormClass) New() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFActionResetForm) Init() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFActionResetForm) Autorelease() PDFActionResetForm {
	rv := objc.Send[PDFActionResetForm](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFActionResetForm creates a new PDFActionResetForm instance.
func NewPDFActionResetForm() PDFActionResetForm {
	return getPDFActionResetFormClass().New()
}



// Returns an array of fields associated with the reset action.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fields
func (p_ PDFActionResetForm) Fields() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("fields"))
	return rv
}


// SetFields sets the value of the fields property.
// Returns an array of fields associated with the reset action.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fields
func (p_ PDFActionResetForm) SetFields(value []string) {
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
	objc.Send[objc.ID](p_.ID, objc.Sel("setFields:"), nsArray)
}

// Sets whether the fields associated with the reset action are cleared when the action is performed.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fieldsIncludedAreCleared
func (p_ PDFActionResetForm) FieldsIncludedAreCleared() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("fieldsIncludedAreCleared"))
	return rv
}


// SetFieldsIncludedAreCleared sets the value of the fieldsIncludedAreCleared property.
// Sets whether the fields associated with the reset action are cleared when the action is performed.

//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFActionResetForm/fieldsIncludedAreCleared
func (p_ PDFActionResetForm) SetFieldsIncludedAreCleared(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFieldsIncludedAreCleared:"), value)
}


