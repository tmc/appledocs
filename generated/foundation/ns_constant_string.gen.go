// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSConstantString */


/* debug [class_header]: Header for NSConstantString */
// The class instance for the [ConstantString] class.
var (
	ConstantStringClass     _ConstantStringClass
	ConstantStringClassOnce sync.Once
)

func getConstantStringClass() _ConstantStringClass {
	ConstantStringClassOnce.Do(func() {
		ConstantStringClass = _ConstantStringClass{objc.GetClass("NSConstantString")}
	})
	return ConstantStringClass
}

type _ConstantStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ConstantString */
// An interface definition for the [ConstantString] class.
type IConstantString interface {
	ISimpleCString
	
/* debug [class_interface_properties]: Properties for ConstantString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ConstantString */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ConstantString */
// Alloc allocates a new instance without initialization.
func (cc _ConstantStringClass) Alloc() ConstantString {
	rv := objc.Send[ConstantString](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ConstantStringClass) New() ConstantString {
	rv := objc.Send[ConstantString](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstantString) Init() ConstantString {
	rv := objc.Send[ConstantString](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstantString) Autorelease() ConstantString {
	rv := objc.Send[ConstantString](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstantString creates a new ConstantString instance.
func NewConstantString() ConstantString {
	return getConstantStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ConstantString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConstantString
type ConstantString struct {
	SimpleCString
}

// ConstantStringFrom constructs a [ConstantString] from an unsafe.Pointer.
func ConstantStringFrom(ptr unsafe.Pointer) ConstantString {
	return ConstantString{
		SimpleCString: SimpleCStringFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ConstantString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ConstantString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ConstantString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ConstantString */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ConstantString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSConstantString */



