// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CXPlayDTMFCallAction] class.
var (
	CXPlayDTMFCallActionClass     _CXPlayDTMFCallActionClass
	CXPlayDTMFCallActionClassOnce sync.Once
)

func getCXPlayDTMFCallActionClass() _CXPlayDTMFCallActionClass {
	CXPlayDTMFCallActionClassOnce.Do(func() {
		CXPlayDTMFCallActionClass = _CXPlayDTMFCallActionClass{objc.GetClass("CXPlayDTMFCallAction")}
	})
	return CXPlayDTMFCallActionClass
}

type _CXPlayDTMFCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXPlayDTMFCallAction] class.
type ICXPlayDTMFCallAction interface {
	ICXCallAction
}

// An encapsulation of the act of playing a dual tone multifrequency (DTMF) sequence.
//
// is a concrete subclass of . Whenever digits are transmitted during a call, whether from a user interacting with a number pad or following a hard or soft pause, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. The provider sends for successive actions only after the current action is fulfilled. When interacting with the number pad, each entered digit constitutes its own action. Digits following a hard or soft pause, however, are passed to as a single string of digits. For example, if a user taps the 4 button on the number pad, followed by the 2 button, the delegate is sent for the digit and waits for the action to be fulfilled; after the action is fulfilled, the delegate is sent for the digit . CallKit automatically plays the corresponding DTMF frequencies for any digits transmitted over a call. The app is responsible for managing the timing and handling of digits as part of fulfilling the action.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction
type CXPlayDTMFCallAction struct {
	CXCallAction
}

// CXPlayDTMFCallActionFrom constructs a [CXPlayDTMFCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of playing a dual tone multifrequency (DTMF) sequence.
func CXPlayDTMFCallActionFrom(ptr unsafe.Pointer) CXPlayDTMFCallAction {
	return CXPlayDTMFCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXPlayDTMFCallActionClass) Alloc() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXPlayDTMFCallActionClass) New() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXPlayDTMFCallAction) Init() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXPlayDTMFCallAction) Autorelease() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXPlayDTMFCallAction creates a new CXPlayDTMFCallAction instance.
func NewCXPlayDTMFCallAction() CXPlayDTMFCallAction {
	return getCXPlayDTMFCallActionClass().New()
}


// Initializes a new action for a call identified by a given UUID, as well as a specified type and sequence of digits.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/init(call:digits:type:)
func NewCXPlayDTMFCallActionWithCallUUIDDigitsType(callUUID unsafe.Pointer, digits string, type_ unsafe.Pointer) CXPlayDTMFCallAction {
	instance := getCXPlayDTMFCallActionClass().Alloc()
	rv := objc.Send[CXPlayDTMFCallAction](instance.ID, objc.Sel("initWithCallUUID:digits:type:"), callUUID, objc.String(digits), type_)
	rv.Autorelease()
	return rv
}

// Creates a new action to play dual-tone multifrequency (DTMF) tones with data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/init(coder:)
func NewCXPlayDTMFCallActionWithCoder(aDecoder unsafe.Pointer) CXPlayDTMFCallAction {
	instance := getCXPlayDTMFCallActionClass().Alloc()
	rv := objc.Send[CXPlayDTMFCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


// The digits tapped by the user into the in-call keypad or included in the dial string.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/digits
func (c_ CXPlayDTMFCallAction) Digits() string {
	rv := objc.Send[string](c_.ID, objc.Sel("digits"))
	return rv
}


// SetDigits sets the value of the digits property.
// The digits tapped by the user into the in-call keypad or included in the dial string.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/digits
func (c_ CXPlayDTMFCallAction) SetDigits(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDigits:"), objc.String(value))
}
// The type of the call action.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/type
func (c_ CXPlayDTMFCallAction) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
// The type of the call action.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/type
func (c_ CXPlayDTMFCallAction) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}

