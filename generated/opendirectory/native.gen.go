// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class native */


/* debug [class_header]: Header for native */
// The class instance for the [native] class.
var (
	NativeClass     _nativeClass
	NativeClassOnce sync.Once
)

func getnativeClass() _nativeClass {
	NativeClassOnce.Do(func() {
		NativeClass = _nativeClass{objc.GetClass("native")}
	})
	return NativeClass
}

type _nativeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for native */
// An interface definition for the [native] class.
type Inative interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for native */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for native */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for native */
// Alloc allocates a new instance without initialization.
func (nc _nativeClass) Alloc() native {
	rv := objc.Send[native](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _nativeClass) New() native {
	rv := objc.Send[native](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ native) Init() native {
	rv := objc.Send[native](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ native) Autorelease() native {
	rv := objc.Send[native](n_.ID, objc.Sel("autorelease"))
	return rv
}

// Newnative creates a new native instance.
func Newnative() native {
	return getnativeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for native */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/native-c.ivar
type native struct {
	objectivec.Object
}

// nativeFrom constructs a [native] from an unsafe.Pointer.
func nativeFrom(ptr unsafe.Pointer) native {
	return native{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for native *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for native */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for native */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for native */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for native */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class native */



