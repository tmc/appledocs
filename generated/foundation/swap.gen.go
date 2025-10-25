// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class swap */


/* debug [class_header]: Header for swap */
// The class instance for the [swap] class.
var (
	SwapClass     _swapClass
	SwapClassOnce sync.Once
)

func getswapClass() _swapClass {
	SwapClassOnce.Do(func() {
		SwapClass = _swapClass{objc.GetClass("swap")}
	})
	return SwapClass
}

type _swapClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for swap */
// An interface definition for the [swap] class.
type Iswap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for swap */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for swap */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for swap */
// Alloc allocates a new instance without initialization.
func (sc _swapClass) Alloc() swap {
	rv := objc.Send[swap](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _swapClass) New() swap {
	rv := objc.Send[swap](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ swap) Init() swap {
	rv := objc.Send[swap](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ swap) Autorelease() swap {
	rv := objc.Send[swap](s_.ID, objc.Sel("autorelease"))
	return rv
}

// Newswap creates a new swap instance.
func Newswap() swap {
	return getswapClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for swap */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/swap
type swap struct {
	objectivec.Object
}

// swapFrom constructs a [swap] from an unsafe.Pointer.
func swapFrom(ptr unsafe.Pointer) swap {
	return swap{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for swap *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for swap */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for swap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for swap */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for swap */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class swap */



