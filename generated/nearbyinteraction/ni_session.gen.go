// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NISession] class.
var (
	NISessionClass     _NISessionClass
	NISessionClassOnce sync.Once
)

func getNISessionClass() _NISessionClass {
	NISessionClassOnce.Do(func() {
		NISessionClass = _NISessionClass{objc.GetClass("NISession")}
	})
	return NISessionClass
}

type _NISessionClass struct {
	class objc.Class
}

// An interface definition for the [NISession] class.
type INISession interface {
	objectivec.IObject
	// properties:
	Configuration() INIConfiguration
	SetConfiguration(value INIConfiguration)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DelegateQueue() unsafe.Pointer
	SetDelegateQueue(value unsafe.Pointer)
	// methods:
}

// An object that identifies a unique connection between two peer devices.
//
// This class represents the central mechanism to interact with nearby objects, for example, a peer Apple device or third-party accessory. After creating an for a nearby object, the app interacts with the object by receiving callbacks. One session represents an interaction between the user and a single nearby object. To interact with multiple nearby objects, create a separate session for each. For more information, see .


// An object that identifies a unique connection between two peer devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession
type NISession struct {
	objectivec.Object
}

// NISessionFrom constructs a [NISession] from an unsafe.Pointer.
//
// An object that identifies a unique connection between two peer devices.
func NISessionFrom(ptr unsafe.Pointer) NISession {
	return NISession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NISessionClass) Alloc() NISession {
	rv := objc.Send[NISession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NISessionClass) New() NISession {
	rv := objc.Send[NISession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NISession) Init() NISession {
	rv := objc.Send[NISession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NISession) Autorelease() NISession {
	rv := objc.Send[NISession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNISession creates a new NISession instance.
func NewNISession() NISession {
	return getNISessionClass().New()
}



// An object that communicates the device’s supported framework features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/deviceCapabilities
func (nc _NISessionClass) DeviceCapabilities() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("deviceCapabilities"))
	return rv
}

// A Boolean value that indicates whether the device supports basic interaction-session functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/isSupported
func (nc _NISessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("supported"))
	return rv
}

// The configuration run by the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NISession) Configuration() INIConfiguration {
	rv := objc.Send[NIConfiguration](n_.ID, objc.Sel("configuration"))
	return rv
}


// The configuration run by the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/configuration
func (n_ NISession) SetConfiguration(value INIConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setConfiguration:"), value)
}


// An object that the framework notifies of session events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NISession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}


// An object that the framework notifies of session events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegate
func (n_ NISession) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegate:"), value)
}


// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NISession) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegateQueue"))
	return rv
}


// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/nearbyinteraction/nisession/delegatequeue
func (n_ NISession) SetDelegateQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDelegateQueue:"), value)
}


