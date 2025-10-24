// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class value */


/* debug [class_header]: Header for value */
// The class instance for the [value] class.
var (
	ValueClass     _valueClass
	ValueClassOnce sync.Once
)

func getvalueClass() _valueClass {
	ValueClassOnce.Do(func() {
		ValueClass = _valueClass{objc.GetClass("value")}
	})
	return ValueClass
}

type _valueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for value */
// An interface definition for the [value] class.
type Ivalue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for value */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for value */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for value */
// Alloc allocates a new instance without initialization.
func (vc _valueClass) Alloc() value {
	rv := objc.Send[value](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _valueClass) New() value {
	rv := objc.Send[value](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ value) Init() value {
	rv := objc.Send[value](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ value) Autorelease() value {
	rv := objc.Send[value](v_.ID, objc.Sel("autorelease"))
	return rv
}

// Newvalue creates a new value instance.
func Newvalue() value {
	return getvalueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for value */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODAttributeMap/value-c.ivar
type value struct {
	objectivec.Object
}

// valueFrom constructs a [value] from an unsafe.Pointer.
func valueFrom(ptr unsafe.Pointer) value {
	return value{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for value *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for value */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for value */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for value */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for value */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class value */



