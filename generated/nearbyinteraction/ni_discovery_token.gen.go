// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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


// A protocol object that describes the nearby interaction capabilities of a person’s device.
//
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NIDiscoveryToken/deviceCapabilities
func (n_ NIDiscoveryToken) DeviceCapabilities() objc.ID {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("deviceCapabilities"))
	return rv
}



