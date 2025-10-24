// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class remotePort */


/* debug [class_header]: Header for remotePort */
// The class instance for the [remotePort] class.
var (
	RemotePortClass     _remotePortClass
	RemotePortClassOnce sync.Once
)

func getremotePortClass() _remotePortClass {
	RemotePortClassOnce.Do(func() {
		RemotePortClass = _remotePortClass{objc.GetClass("remotePort")}
	})
	return RemotePortClass
}

type _remotePortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for remotePort */
// An interface definition for the [remotePort] class.
type IremotePort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for remotePort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for remotePort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for remotePort */
// Alloc allocates a new instance without initialization.
func (rc _remotePortClass) Alloc() remotePort {
	rv := objc.Send[remotePort](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _remotePortClass) New() remotePort {
	rv := objc.Send[remotePort](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ remotePort) Init() remotePort {
	rv := objc.Send[remotePort](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ remotePort) Autorelease() remotePort {
	rv := objc.Send[remotePort](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewremotePort creates a new remotePort instance.
func NewremotePort() remotePort {
	return getremotePortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for remotePort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/remotePort
type remotePort struct {
	objectivec.Object
}

// remotePortFrom constructs a [remotePort] from an unsafe.Pointer.
func remotePortFrom(ptr unsafe.Pointer) remotePort {
	return remotePort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for remotePort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for remotePort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for remotePort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for remotePort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for remotePort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class remotePort */



