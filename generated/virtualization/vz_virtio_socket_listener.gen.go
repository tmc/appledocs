// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VZVirtioSocketListener */


/* debug [class_header]: Header for VZVirtioSocketListener */
// The class instance for the [VZVirtioSocketListener] class.
var (
	VZVirtioSocketListenerClass     _VZVirtioSocketListenerClass
	VZVirtioSocketListenerClassOnce sync.Once
)

func getVZVirtioSocketListenerClass() _VZVirtioSocketListenerClass {
	VZVirtioSocketListenerClassOnce.Do(func() {
		VZVirtioSocketListenerClass = _VZVirtioSocketListenerClass{objc.GetClass("VZVirtioSocketListener")}
	})
	return VZVirtioSocketListenerClass
}

type _VZVirtioSocketListenerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for VZVirtioSocketListener */
// An interface definition for the [VZVirtioSocketListener] class.
type IVZVirtioSocketListener interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for VZVirtioSocketListener */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for VZVirtioSocketListener */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for VZVirtioSocketListener */
// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketListenerClass) Alloc() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VZVirtioSocketListenerClass) New() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioSocketListener) Init() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioSocketListener) Autorelease() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketListener creates a new VZVirtioSocketListener instance.
func NewVZVirtioSocketListener() VZVirtioSocketListener {
	return getVZVirtioSocketListenerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for VZVirtioSocketListener */
// An object that listens for port-based connection requests from the guest operating system.
//
// Use a object to route connection requests to your associated delegate object. The socket listener object handles incoming connection requests from the guest operating system and directs them to the methods of its associated object. You may use the same listener object to monitor connections on multiple ports. After creating a object, assign a custom object to its property. The delegate must implement the protocol. To connect the listener to a port, call the method of your virtual machine’s object.


// An object that listens for port-based connection requests from the guest operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener
type VZVirtioSocketListener struct {
	objectivec.Object
}

// VZVirtioSocketListenerFrom constructs a [VZVirtioSocketListener] from an unsafe.Pointer.
//
// An object that listens for port-based connection requests from the guest operating system.
func VZVirtioSocketListenerFrom(ptr unsafe.Pointer) VZVirtioSocketListener {
	return VZVirtioSocketListener{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for VZVirtioSocketListener *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for VZVirtioSocketListener */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for VZVirtioSocketListener */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for VZVirtioSocketListener */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for VZVirtioSocketListener */

// The custom object you use to respond to port-based connection attempts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener/delegate
func (v_ VZVirtioSocketListener) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The custom object you use to respond to port-based connection attempts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener/delegate
func (v_ VZVirtioSocketListener) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VZVirtioSocketListener */



