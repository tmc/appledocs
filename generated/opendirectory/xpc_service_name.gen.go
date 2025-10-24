// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class xpcServiceName */


/* debug [class_header]: Header for xpcServiceName */
// The class instance for the [xpcServiceName] class.
var (
	XpcServiceNameClass     _xpcServiceNameClass
	XpcServiceNameClassOnce sync.Once
)

func getxpcServiceNameClass() _xpcServiceNameClass {
	XpcServiceNameClassOnce.Do(func() {
		XpcServiceNameClass = _xpcServiceNameClass{objc.GetClass("xpcServiceName")}
	})
	return XpcServiceNameClass
}

type _xpcServiceNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for xpcServiceName */
// An interface definition for the [xpcServiceName] class.
type IxpcServiceName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for xpcServiceName */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for xpcServiceName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for xpcServiceName */
// Alloc allocates a new instance without initialization.
func (xc _xpcServiceNameClass) Alloc() xpcServiceName {
	rv := objc.Send[xpcServiceName](objc.ID(xc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (xc _xpcServiceNameClass) New() xpcServiceName {
	rv := objc.Send[xpcServiceName](objc.ID(xc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (x_ xpcServiceName) Init() xpcServiceName {
	rv := objc.Send[xpcServiceName](x_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x_ xpcServiceName) Autorelease() xpcServiceName {
	rv := objc.Send[xpcServiceName](x_.ID, objc.Sel("autorelease"))
	return rv
}

// NewxpcServiceName creates a new xpcServiceName instance.
func NewxpcServiceName() xpcServiceName {
	return getxpcServiceNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for xpcServiceName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/xpcServiceName-c.ivar
type xpcServiceName struct {
	objectivec.Object
}

// xpcServiceNameFrom constructs a [xpcServiceName] from an unsafe.Pointer.
func xpcServiceNameFrom(ptr unsafe.Pointer) xpcServiceName {
	return xpcServiceName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for xpcServiceName *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for xpcServiceName */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for xpcServiceName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for xpcServiceName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for xpcServiceName */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class xpcServiceName */



