// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNull */


/* debug [class_header]: Header for NSNull */
// The class instance for the [Null] class.
var (
	NullClass     _NullClass
	NullClassOnce sync.Once
)

func getNullClass() _NullClass {
	NullClassOnce.Do(func() {
		NullClass = _NullClass{objc.GetClass("NSNull")}
	})
	return NullClass
}

type _NullClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Null */
// An interface definition for the [Null] class.
type INull interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Null */
	// properties:
	NSNotFound() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Null */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Null */
// Alloc allocates a new instance without initialization.
func (nc _NullClass) Alloc() Null {
	rv := objc.Send[Null](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NullClass) New() Null {
	rv := objc.Send[Null](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Null) Init() Null {
	rv := objc.Send[Null](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Null) Autorelease() Null {
	rv := objc.Send[Null](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNull creates a new Null instance.
func NewNull() Null {
	return getNullClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Null */
// A singleton object used to represent null values in collection objects that don’t allow values.
//
// is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A singleton object used to represent null values in collection objects that don’t allow values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNull
type Null struct {
	objectivec.Object
}

// NullFrom constructs a [Null] from an unsafe.Pointer.
//
// A singleton object used to represent null values in collection objects that don’t allow values.
func NullFrom(ptr unsafe.Pointer) Null {
	return Null{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Null *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Null */

// Returns the singleton instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNull/null
func (nc _NullClass) Null() INull {
	rv := objc.Send[Null](objc.ID(nc.class), objc.Sel("null"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Null) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Null */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Null */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Null */

// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (n_ Null) NSNotFound() int {
	rv := objc.Send[int](n_.ID, objc.Sel("NSNotFound"))
	return rv
}/* debug [instance_properties/getter]: NSNotFound */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNull */



