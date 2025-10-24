// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INExtension] class.
var (
	INExtensionClass     _INExtensionClass
	INExtensionClassOnce sync.Once
)

func getINExtensionClass() _INExtensionClass {
	INExtensionClassOnce.Do(func() {
		INExtensionClass = _INExtensionClass{objc.GetClass("INExtension")}
	})
	return INExtensionClass
}

type _INExtensionClass struct {
	class objc.Class
}

// An interface definition for the [INExtension] class.
type IINExtension interface {
	objectivec.IObject
}

// The entry point for an Intents extension.
//
// The class is the principal class of your Intents extension, which provides information to SiriKit. Your extension object acts like a dispatcher, providing SiriKit with the objects it needs to resolve and handle requests. This class defines no methods of its own, but it adopts the protocol, which contains the main method you must implement. You do not create instances of this class directly. When the system needs your Intents extension to handle a request, it automatically instantiates the class specified in the key of the extension’s file. All you have to do is provide the implementation for that class. Your extension object must be able to return handler objects for all of the intents that your app supports. The Intents framework supports the following types of intents: VoIP calling intents support audio and video calls to another user of your app. Workout intents support the starting and stopping of workouts. Message intents support the sending of message data to specific users. Payment intents support financial transactions between users. Photo intents support the searching and displaying of photos. Ride-booking intents support the booking and management of user transportation from Siri or Maps. CarPlay intents support the changing of settings in automobiles that support CarPlay. Restaurant reservation intents support the creating and viewing of restaurant reservations in Maps. For more information about implementing your extension’s main dispatching method, see . For information about how to implement this class in your Intents extension, see .

// The entry point for an Intents extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INExtension
type INExtension struct {
	objectivec.Object
}

// INExtensionFrom constructs a [INExtension] from an unsafe.Pointer.
//
// The entry point for an Intents extension.
func INExtensionFrom(ptr unsafe.Pointer) INExtension {
	return INExtension{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INExtensionClass) Alloc() INExtension {
	rv := objc.Send[INExtension](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INExtensionClass) New() INExtension {
	rv := objc.Send[INExtension](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INExtension) Init() INExtension {
	rv := objc.Send[INExtension](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INExtension) Autorelease() INExtension {
	rv := objc.Send[INExtension](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINExtension creates a new INExtension instance.
func NewINExtension() INExtension {
	return getINExtensionClass().New()
}
