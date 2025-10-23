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
	CurrentAppleEvent() IAppleEventDescriptor
	CurrentReplyAppleEvent() IAppleEventDescriptor
	AppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) IAppleEventDescriptor
	DispatchRawAppleEventWithRawReplyHandlerRefCon(theAppleEvent unsafe.Pointer, theReply unsafe.Pointer, handlerRefCon unsafe.Pointer) unsafe.Pointer
	RemoveEventHandlerForEventClassAndEventID(eventClass unsafe.Pointer, eventID unsafe.Pointer)
	ReplyAppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) IAppleEventDescriptor
	ResumeWithSuspensionID(suspensionID AppleEventManagerSuspensionID)
	SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID AppleEventManagerSuspensionID)
	SetEventHandlerAndSelectorForEventClassAndEventID(handler objectivec.IObject, handleEventSelector objc.SEL, eventClass unsafe.Pointer, eventID unsafe.Pointer)
	SuspendCurrentAppleEvent() AppleEventManagerSuspensionID
}

// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers.
//
// Cocoa provides built-in scriptability support that uses scriptability information supplied by an application to automatically convert Apple events into script command objects that perform the desired operation. However, some applications may want to perform more basic Apple event handling, in which an application registers handlers for the Apple events it can process, then calls on the Apple Event Manager to dispatch received Apple events to the appropriate handler. supports these mechanisms by providing methods to register and remove handlers and to dispatch Apple events to the appropriate handler, if one exists. For related information, see Each application has at most one instance of . To obtain a reference to it, you call the class method , which creates the instance if it doesn’t already exist. For information about the Apple Event Manager, see and Apple Events Programming Guide.


// A mechanism for registering handler routines for specific types of Apple events and dispatching events to those handlers.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/shared()
func (ac _AppleEventManagerClass) SharedAppleEventManager() IAppleEventManager {
	rv := objc.Send[AppleEventManager](objc.ID(ac.class), objc.Sel("sharedAppleEventManager"))
	return rv
}


// Given a nonzero returned by an invocation of , returns the descriptor for the event whose handling was suspended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/appleEvent(forSuspensionID:)
func (a_ AppleEventManager) AppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("appleEventForSuspensionID:"), suspensionID)
	return rv
}


// Causes the Apple event specified by to be dispatched to the appropriate Apple event handler, if one has been registered by calling .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/dispatchRawAppleEvent(_:withRawReply:handlerRefCon:)
func (a_ AppleEventManager) DispatchRawAppleEventWithRawReplyHandlerRefCon(theAppleEvent unsafe.Pointer, theReply unsafe.Pointer, handlerRefCon unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("dispatchRawAppleEvent:withRawReply:handlerRefCon:"), theAppleEvent, theReply, handlerRefCon)
	return rv
}


// If an Apple event handler has been registered for the event specified by and , removes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/removeEventHandler(forEventClass:andEventID:)
func (a_ AppleEventManager) RemoveEventHandlerForEventClassAndEventID(eventClass unsafe.Pointer, eventID unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeEventHandlerForEventClass:andEventID:"), eventClass, eventID)
}


// Given a nonzero returned by an invocation of , returns the corresponding reply event descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/replyAppleEvent(forSuspensionID:)
func (a_ AppleEventManager) ReplyAppleEventForSuspensionID(suspensionID AppleEventManagerSuspensionID) IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("replyAppleEventForSuspensionID:"), suspensionID)
	return rv
}


// Given a nonzero returned by an invocation of , signal that handling of the suspended event may now continue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/resume(withSuspensionID:)
func (a_ AppleEventManager) ResumeWithSuspensionID(suspensionID AppleEventManagerSuspensionID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("resumeWithSuspensionID:"), suspensionID)
}


// Given a nonzero returned by an invocation of , sets the values that will be returned by subsequent invocations of and to be the event whose handling was suspended and its corresponding reply event, respectively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/setCurrentAppleEventAndReplyEventWithSuspensionID(_:)
func (a_ AppleEventManager) SetCurrentAppleEventAndReplyEventWithSuspensionID(suspensionID AppleEventManagerSuspensionID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentAppleEventAndReplyEventWithSuspensionID:"), suspensionID)
}


// Registers the Apple event handler specified by for the event specified by and .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/setEventHandler(_:andSelector:forEventClass:andEventID:)
func (a_ AppleEventManager) SetEventHandlerAndSelectorForEventClassAndEventID(handler objectivec.IObject, handleEventSelector objc.SEL, eventClass unsafe.Pointer, eventID unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEventHandler:andSelector:forEventClass:andEventID:"), handler, handleEventSelector, eventClass, eventID)
}


// Suspends the handling of the current event and returns an ID that must be used to resume the handling of the event if an Apple event is being handled on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/suspendCurrentAppleEvent()
func (a_ AppleEventManager) SuspendCurrentAppleEvent() AppleEventManagerSuspensionID {
	rv := objc.Send[AppleEventManagerSuspensionID](a_.ID, objc.Sel("suspendCurrentAppleEvent"))
	return rv
}


// Returns the descriptor for if an Apple event is being handled on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/currentAppleEvent
func (a_ AppleEventManager) CurrentAppleEvent() IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("currentAppleEvent"))
	return rv
}


// Returns the corresponding reply event descriptor if an Apple event is being handled on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager/currentReplyAppleEvent
func (a_ AppleEventManager) CurrentReplyAppleEvent() IAppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("currentReplyAppleEvent"))
	return rv
}



