// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSForm */


/* debug [class_header]: Header for NSForm */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Form */
// An interface definition for the [Form] class.
type IForm interface {
	IMatrix
	
/* debug [class_interface_properties]: Properties for Form */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Form */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Form */
// Alloc allocates a new instance without initialization.
func (fc _FormClass) Alloc() Form {
	rv := objc.Send[Form](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Form */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Form *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Form */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Form */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Form */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Form */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSForm */



