// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class arrayLength */


/* debug [class_header]: Header for arrayLength */
// The class instance for the [arrayLength] class.
var (
	ArrayLengthClass     _arrayLengthClass
	ArrayLengthClassOnce sync.Once
)

func getarrayLengthClass() _arrayLengthClass {
	ArrayLengthClassOnce.Do(func() {
		ArrayLengthClass = _arrayLengthClass{objc.GetClass("arrayLength")}
	})
	return ArrayLengthClass
}

type _arrayLengthClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for arrayLength */
// An interface definition for the [arrayLength] class.
type IarrayLength interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for arrayLength */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for arrayLength */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for arrayLength */
// Alloc allocates a new instance without initialization.
func (ac _arrayLengthClass) Alloc() arrayLength {
	rv := objc.Send[arrayLength](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _arrayLengthClass) New() arrayLength {
	rv := objc.Send[arrayLength](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ arrayLength) Init() arrayLength {
	rv := objc.Send[arrayLength](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ arrayLength) Autorelease() arrayLength {
	rv := objc.Send[arrayLength](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewarrayLength creates a new arrayLength instance.
func NewarrayLength() arrayLength {
	return getarrayLengthClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for arrayLength */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/arrayLength-c.ivar
type arrayLength struct {
	objectivec.Object
}

// arrayLengthFrom constructs a [arrayLength] from an unsafe.Pointer.
func arrayLengthFrom(ptr unsafe.Pointer) arrayLength {
	return arrayLength{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for arrayLength *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for arrayLength */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for arrayLength */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for arrayLength */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for arrayLength */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class arrayLength */



