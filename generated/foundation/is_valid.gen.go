// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class isValid */


/* debug [class_header]: Header for isValid */
// The class instance for the [isValid] class.
var (
	IsValidClass     _isValidClass
	IsValidClassOnce sync.Once
)

func getisValidClass() _isValidClass {
	IsValidClassOnce.Do(func() {
		IsValidClass = _isValidClass{objc.GetClass("isValid")}
	})
	return IsValidClass
}

type _isValidClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for isValid */
// An interface definition for the [isValid] class.
type IisValid interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for isValid */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for isValid */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for isValid */
// Alloc allocates a new instance without initialization.
func (ic _isValidClass) Alloc() isValid {
	rv := objc.Send[isValid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isValidClass) New() isValid {
	rv := objc.Send[isValid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isValid) Init() isValid {
	rv := objc.Send[isValid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isValid) Autorelease() isValid {
	rv := objc.Send[isValid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisValid creates a new isValid instance.
func NewisValid() isValid {
	return getisValidClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for isValid */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isValid
type isValid struct {
	objectivec.Object
}

// isValidFrom constructs a [isValid] from an unsafe.Pointer.
func isValidFrom(ptr unsafe.Pointer) isValid {
	return isValid{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for isValid *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for isValid */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for isValid */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for isValid */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for isValid */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isValid */



