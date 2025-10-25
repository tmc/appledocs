// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSimpleCString */


/* debug [class_header]: Header for NSSimpleCString */
// The class instance for the [SimpleCString] class.
var (
	SimpleCStringClass     _SimpleCStringClass
	SimpleCStringClassOnce sync.Once
)

func getSimpleCStringClass() _SimpleCStringClass {
	SimpleCStringClassOnce.Do(func() {
		SimpleCStringClass = _SimpleCStringClass{objc.GetClass("NSSimpleCString")}
	})
	return SimpleCStringClass
}

type _SimpleCStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SimpleCString */
// An interface definition for the [SimpleCString] class.
type ISimpleCString interface {
	IString
	
/* debug [class_interface_properties]: Properties for SimpleCString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SimpleCString */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SimpleCString */
// Alloc allocates a new instance without initialization.
func (sc _SimpleCStringClass) Alloc() SimpleCString {
	rv := objc.Send[SimpleCString](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SimpleCStringClass) New() SimpleCString {
	rv := objc.Send[SimpleCString](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SimpleCString) Init() SimpleCString {
	rv := objc.Send[SimpleCString](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SimpleCString) Autorelease() SimpleCString {
	rv := objc.Send[SimpleCString](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSimpleCString creates a new SimpleCString instance.
func NewSimpleCString() SimpleCString {
	return getSimpleCStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SimpleCString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString
type SimpleCString struct {
	String
}

// SimpleCStringFrom constructs a [SimpleCString] from an unsafe.Pointer.
func SimpleCStringFrom(ptr unsafe.Pointer) SimpleCString {
	return SimpleCString{
		String: StringFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SimpleCString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SimpleCString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SimpleCString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SimpleCString */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SimpleCString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSimpleCString */



