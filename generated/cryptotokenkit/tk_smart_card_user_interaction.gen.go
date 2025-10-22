// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TKSmartCardUserInteraction] class.
var (
	TKSmartCardUserInteractionClass     _TKSmartCardUserInteractionClass
	TKSmartCardUserInteractionClassOnce sync.Once
)

func getTKSmartCardUserInteractionClass() _TKSmartCardUserInteractionClass {
	TKSmartCardUserInteractionClassOnce.Do(func() {
		TKSmartCardUserInteractionClass = _TKSmartCardUserInteractionClass{objc.GetClass("TKSmartCardUserInteraction")}
	})
	return TKSmartCardUserInteractionClass
}

type _TKSmartCardUserInteractionClass struct {
	class objc.Class
}

// An interface definition for the [TKSmartCardUserInteraction] class.
type ITKSmartCardUserInteraction interface {
	objectivec.IObject
	Cancel() bool
	RunWithReply(reply unsafe.Pointer)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	InitialTimeout() foundation.TimeInterval
	SetInitialTimeout(value foundation.ITimeInterval)
	InteractionTimeout() foundation.TimeInterval
	SetInteractionTimeout(value foundation.ITimeInterval)
}

// The base class for encapsulating user interaction with a Smart Card reader.
//
// There are two types of user interactions: those for secure PIN change and those for secure PIN validation. These interactions are instances of the , or subclasses of , respectively. is a subclass of . You interact with instances of one of the subclasses of when calling the and methods on an object.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction
type TKSmartCardUserInteraction struct {
	objectivec.Object
}

// TKSmartCardUserInteractionFrom constructs a [TKSmartCardUserInteraction] from an unsafe.Pointer.
//
// The base class for encapsulating user interaction with a Smart Card reader.
func TKSmartCardUserInteractionFrom(ptr unsafe.Pointer) TKSmartCardUserInteraction {
	return TKSmartCardUserInteraction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardUserInteractionClass) Alloc() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TKSmartCardUserInteractionClass) New() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardUserInteraction) Init() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardUserInteraction) Autorelease() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteraction creates a new TKSmartCardUserInteraction instance.
func NewTKSmartCardUserInteraction() TKSmartCardUserInteraction {
	return getTKSmartCardUserInteractionClass().New()
}


// Attempts to cancel an interaction started by calling . For certain interactions, cancellation may not be available.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/cancel()
func (t_ TKSmartCardUserInteraction) Cancel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("cancel"))
	return rv
}

// Runs the user interaction and asynchronously receives a reply.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)
func (t_ TKSmartCardUserInteraction) RunWithReply(reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("runWithReply:"), reply)
}

// The delegate for observing events that occur during the user interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/delegate
func (t_ TKSmartCardUserInteraction) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for observing events that occur during the user interaction.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/delegate
func (t_ TKSmartCardUserInteraction) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// The timeout, in seconds, for initial interaction. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/initialTimeout
func (t_ TKSmartCardUserInteraction) InitialTimeout() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](t_.ID, objc.Sel("initialTimeout"))
	return rv
}


// SetInitialTimeout sets the value of the initialTimeout property.
// The timeout, in seconds, for initial interaction. If set to , the reader-defined default timeout is used. by default.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/initialTimeout
func (t_ TKSmartCardUserInteraction) SetInitialTimeout(value foundation.ITimeInterval) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInitialTimeout:"), value)
}

// The timeout, in seconds, after the first key stroke. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/interactionTimeout
func (t_ TKSmartCardUserInteraction) InteractionTimeout() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](t_.ID, objc.Sel("interactionTimeout"))
	return rv
}


// SetInteractionTimeout sets the value of the interactionTimeout property.
// The timeout, in seconds, after the first key stroke. If set to , the reader-defined default timeout is used. by default.

//
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/interactionTimeout
func (t_ TKSmartCardUserInteraction) SetInteractionTimeout(value foundation.ITimeInterval) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInteractionTimeout:"), value)
}



