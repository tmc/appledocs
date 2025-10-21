// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [VZVirtioSocketListener] class.
type IVZVirtioSocketListener interface {
	objectivec.IObject
}

// An object that listens for port-based connection requests from the guest operating system.
//
// Use a object to route connection requests to your associated delegate object. The socket listener object handles incoming connection requests from the guest operating system and directs them to the methods of its associated object. You may use the same listener object to monitor connections on multiple ports. After creating a object, assign a custom object to its property. The delegate must implement the protocol. To connect the listener to a port, call the method of your virtual machine’s object.
//
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

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioSocketListenerClass) Alloc() VZVirtioSocketListener {
	rv := objc.Send[VZVirtioSocketListener](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The custom object you use to respond to port-based connection attempts.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener/delegate
func (v_ VZVirtioSocketListener) Delegate() objc.ID {
	rv := objc.Send[objc.ID](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The custom object you use to respond to port-based connection attempts.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketListener/delegate
func (v_ VZVirtioSocketListener) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}


