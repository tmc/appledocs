// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class functionAttributes */


/* debug [class_header]: Header for functionAttributes */
// The class instance for the [functionAttributes] class.
var (
	FunctionAttributesClass     _functionAttributesClass
	FunctionAttributesClassOnce sync.Once
)

func getfunctionAttributesClass() _functionAttributesClass {
	FunctionAttributesClassOnce.Do(func() {
		FunctionAttributesClass = _functionAttributesClass{objc.GetClass("functionAttributes")}
	})
	return FunctionAttributesClass
}

type _functionAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for functionAttributes */
// An interface definition for the [functionAttributes] class.
type IfunctionAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for functionAttributes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for functionAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for functionAttributes */
// Alloc allocates a new instance without initialization.
func (fc _functionAttributesClass) Alloc() functionAttributes {
	rv := objc.Send[functionAttributes](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _functionAttributesClass) New() functionAttributes {
	rv := objc.Send[functionAttributes](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ functionAttributes) Init() functionAttributes {
	rv := objc.Send[functionAttributes](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ functionAttributes) Autorelease() functionAttributes {
	rv := objc.Send[functionAttributes](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewfunctionAttributes creates a new functionAttributes instance.
func NewfunctionAttributes() functionAttributes {
	return getfunctionAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for functionAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/functionAttributes-c.ivar
type functionAttributes struct {
	objectivec.Object
}

// functionAttributesFrom constructs a [functionAttributes] from an unsafe.Pointer.
func functionAttributesFrom(ptr unsafe.Pointer) functionAttributes {
	return functionAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for functionAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for functionAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for functionAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for functionAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for functionAttributes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class functionAttributes */



