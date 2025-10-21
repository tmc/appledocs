// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PushRegistry] class.
var (
	PushRegistryClass     _PushRegistryClass
	PushRegistryClassOnce sync.Once
)

func getPushRegistryClass() _PushRegistryClass {
	PushRegistryClassOnce.Do(func() {
		PushRegistryClass = _PushRegistryClass{objc.GetClass("PKPushRegistry")}
	})
	return PushRegistryClass
}

type _PushRegistryClass struct {
	class objc.Class
}

// An interface definition for the [PushRegistry] class.
type IPushRegistry interface {
	objectivec.IObject
	PushTokenForType(type_ unsafe.Pointer) unsafe.Pointer
}

// An object that requests the delivery and handles the receipt of PushKit notifications.
//
// A object manages only certain types of notifications, such as high-priority notifications needed by a VoIP app. PushKit wakes up your app as needed to deliver incoming notifications and delivers the notifications directly to the push registry object that requested them. Every time your app launches, whether in the foreground or in the background, create a push registry object and configure it. Typically, you keep the push registry object running for the duration of your app. Each push registry object delivers incoming notifications to its object, which also handles the responses for registration requests. Listing 1 shows how to create a push registry object and request VoIP notifications. Always assign an appropriate delegate object before modifying the property. Listing 1. Creating and configuring a push registry object Assigning a new value to the property registers the push registry object with the PushKit servers. The server reports the success or failure of your registration attempts asynchronously to the push registry, which then reports those results to its delegate object. The push registry also delivers all received notifications to the delegate object. For more information about the delegate methods, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry
type PushRegistry struct {
	objectivec.Object
}

// PushRegistryFrom constructs a [PushRegistry] from an unsafe.Pointer.
//
// An object that requests the delivery and handles the receipt of PushKit notifications.
func PushRegistryFrom(ptr unsafe.Pointer) PushRegistry {
	return PushRegistry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PushRegistryClass) Alloc() PushRegistry {
	rv := objc.Send[PushRegistry](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PushRegistryClass) New() PushRegistry {
	rv := objc.Send[PushRegistry](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PushRegistry) Init() PushRegistry {
	rv := objc.Send[PushRegistry](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PushRegistry) Autorelease() PushRegistry {
	rv := objc.Send[PushRegistry](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPushRegistry creates a new PushRegistry instance.
func NewPushRegistry() PushRegistry {
	return getPushRegistryClass().New()
}


// Creates a push registry with the specified dispatch queue.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/init(queue:)
func NewPushRegistryWithQueue(queue unsafe.Pointer) PushRegistry {
	instance := getPushRegistryClass().Alloc()
	rv := objc.Send[PushRegistry](instance.ID, objc.Sel("initWithQueue:"), queue)
	rv.Autorelease()
	return rv
}


// Retrieves the locally cached push token for the specified push type.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/pushToken(for:)
func (p_ PushRegistry) PushTokenForType(type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pushTokenForType:"), type_)
	return rv
}

// The delegate object that receives notifications coming from the push registry object.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/delegate
func (p_ PushRegistry) Delegate() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that receives notifications coming from the push registry object.

//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/delegate
func (p_ PushRegistry) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}
// Registers the push types for this push registry object.
//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/desiredPushTypes
func (p_ PushRegistry) DesiredPushTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("desiredPushTypes"))
	return rv
}


// SetDesiredPushTypes sets the value of the desiredPushTypes property.
// Registers the push types for this push registry object.

//
// [Full Topic]: https://developer.apple.com/documentation/PushKit/PKPushRegistry/desiredPushTypes
func (p_ PushRegistry) SetDesiredPushTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDesiredPushTypes:"), value)
}

