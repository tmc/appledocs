// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventManager] class.
var (
	AppleEventManagerClass     _AppleEventManagerClass
	AppleEventManagerClassOnce sync.Once
)

func getAppleEventManagerClass() _AppleEventManagerClass {
	AppleEventManagerClassOnce.Do(func() {
		AppleEventManagerClass = _AppleEventManagerClass{objc.GetClass("NSAppleEventManager")}
	})
	return AppleEventManagerClass
}

type _AppleEventManagerClass struct {
	class objc.Class
}

// An interface definition for the [AppleEventManager] class.
type IAppleEventManager interface {
	objectivec.IObject
	CurrentAppleEvent() NSAppleEventDescriptor
	SetCurrentAppleEvent(value IAppleEventDescriptor)
	CurrentReplyAppleEvent() NSAppleEventDescriptor
	SetCurrentReplyAppleEvent(value IAppleEventDescriptor)
}

// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers.
//
// Cocoa provides built-in scriptability support that uses scriptability information supplied by an application to automatically convert Apple events into script command objects that perform the desired operation. However, some applications may want to perform more basic Apple event handling, in which an application registers handlers for the Apple events it can process, then calls on the Apple Event Manager to dispatch received Apple events to the appropriate handler. supports these mechanisms by providing methods to register and remove handlers and to dispatch Apple events to the appropriate handler, if one exists. For related information, see Each application has at most one instance of . To obtain a reference to it, you call the class method , which creates the instance if it doesn’t already exist. For information about the Apple Event Manager, see and Apple Events Programming Guide.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager
type AppleEventManager struct {
	objectivec.Object
}

// AppleEventManagerFrom constructs a [AppleEventManager] from an unsafe.Pointer.
//
// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers.
func AppleEventManagerFrom(ptr unsafe.Pointer) AppleEventManager {
	return AppleEventManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleEventManagerClass) Alloc() AppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleEventManagerClass) New() AppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleEventManager) Init() AppleEventManager {
	rv := objc.Send[AppleEventManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleEventManager) Autorelease() AppleEventManager {
	rv := objc.Send[AppleEventManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleEventManager creates a new AppleEventManager instance.
func NewAppleEventManager() AppleEventManager {
	return getAppleEventManagerClass().New()
}


// Returns the single instance of , creating it first if it doesn’t exist.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/shared()
func (ac _AppleEventManagerClass) SharedAppleEventManager() AppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("sharedAppleEventManager"))
	return rv
}

// Returns the descriptor for
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/currentappleevent
func (a_ AppleEventManager) CurrentAppleEvent() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a_.ID, objc.Sel("currentAppleEvent"))
	return rv
}


// SetCurrentAppleEvent sets the value of the currentAppleEvent property.
// Returns the descriptor for

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/currentappleevent
func (a_ AppleEventManager) SetCurrentAppleEvent(value IAppleEventDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentAppleEvent:"), value)
}

// Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/currentreplyappleevent
func (a_ AppleEventManager) CurrentReplyAppleEvent() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](a_.ID, objc.Sel("currentReplyAppleEvent"))
	return rv
}


// SetCurrentReplyAppleEvent sets the value of the currentReplyAppleEvent property.
// Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventmanager/currentreplyappleevent
func (a_ AppleEventManager) SetCurrentReplyAppleEvent(value IAppleEventDescriptor) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentReplyAppleEvent:"), value)
}



