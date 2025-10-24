// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Form] class.
var (
	FormClass     _FormClass
	FormClassOnce sync.Once
)

func getFormClass() _FormClass {
	FormClassOnce.Do(func() {
		FormClass = _FormClass{objc.GetClass("NSForm")}
	})
	return FormClass
}

type _FormClass struct {
	class objc.Class
}

// An interface definition for the [Form] class.
type IForm interface {
	IMatrix
	// properties:
	// methods:
}

// An object is a vertical matrix of objects to implement the fields.


// An object is a vertical matrix of objects to implement the fields.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSForm
type Form struct {
	Matrix
}

// FormFrom constructs a [Form] from an unsafe.Pointer.
//
// An object is a vertical matrix of objects to implement the fields.
func FormFrom(ptr unsafe.Pointer) Form {
	return Form{
		Matrix: MatrixFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (fc _FormClass) Alloc() Form {
	rv := objc.Send[Form](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FormClass) New() Form {
	rv := objc.Send[Form](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Form) Init() Form {
	rv := objc.Send[Form](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Form) Autorelease() Form {
	rv := objc.Send[Form](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewForm creates a new Form instance.
func NewForm() Form {
	return getFormClass().New()
}




