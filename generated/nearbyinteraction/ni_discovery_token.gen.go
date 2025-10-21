// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NIDiscoveryToken] class.
var (
	NIDiscoveryTokenClass     _NIDiscoveryTokenClass
	NIDiscoveryTokenClassOnce sync.Once
)

func getNIDiscoveryTokenClass() _NIDiscoveryTokenClass {
	NIDiscoveryTokenClassOnce.Do(func() {
		NIDiscoveryTokenClass = _NIDiscoveryTokenClass{objc.GetClass("NIDiscoveryToken")}
	})
	return NIDiscoveryTokenClass
}

type _NIDiscoveryTokenClass struct {
	class objc.Class
}

// An interface definition for the [NIDiscoveryToken] class.
type INIDiscoveryToken interface {
	objectivec.IObject
}

// An object that uniquely identifies a peer that participates in an interaction session.
//
// Use to determine the peer device’s nearby interaction capabilities by examining the that describes the available capabilities on a person’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDiscoveryToken
type NIDiscoveryToken struct {
	objectivec.Object
}

// NIDiscoveryTokenFrom constructs a [NIDiscoveryToken] from an unsafe.Pointer.
//
// An object that uniquely identifies a peer that participates in an interaction session.
func NIDiscoveryTokenFrom(ptr unsafe.Pointer) NIDiscoveryToken {
	return NIDiscoveryToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NIDiscoveryTokenClass) Alloc() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NIDiscoveryTokenClass) New() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NIDiscoveryToken) Init() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NIDiscoveryToken) Autorelease() NIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNIDiscoveryToken creates a new NIDiscoveryToken instance.
func NewNIDiscoveryToken() NIDiscoveryToken {
	return getNIDiscoveryTokenClass().New()
}


// The configuration run by the session.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NIDiscoveryToken) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// The configuration run by the session.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NIDiscoveryToken) SetConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConfiguration:"), value)
}

// A temporary, random identifier for a device.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/discoverytoken
func (n_ NIDiscoveryToken) DiscoveryToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("discoveryToken"))
	return rv
}


// SetDiscoveryToken sets the value of the discoveryToken property.
// A temporary, random identifier for a device.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/discoverytoken
func (n_ NIDiscoveryToken) SetDiscoveryToken(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDiscoveryToken:"), value)
}

// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NIDiscoveryToken) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegateQueue"))
	return rv
}


// SetDelegateQueue sets the value of the delegateQueue property.
// The dispatch queue on which the session invokes delegate callbacks.

//
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NIDiscoveryToken) SetDelegateQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegateQueue:"), value)
}

// A protocol object that describes the nearby interaction capabilities of a person’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDiscoveryToken/deviceCapabilities
func (n_ NIDiscoveryToken) DeviceCapabilities() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("deviceCapabilities"))
	return rv
}



