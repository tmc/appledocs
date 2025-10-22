// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXStartCallAction] class.
var (
	CXStartCallActionClass     _CXStartCallActionClass
	CXStartCallActionClassOnce sync.Once
)

func getCXStartCallActionClass() _CXStartCallActionClass {
	CXStartCallActionClassOnce.Do(func() {
		CXStartCallActionClass = _CXStartCallActionClass{objc.GetClass("CXStartCallAction")}
	})
	return CXStartCallActionClass
}

type _CXStartCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXStartCallAction] class.
type ICXStartCallAction interface {
	ICXCallAction
	FulfillWithDateStarted(dateStarted foundation.IDate)
	ContactIdentifier() string
	SetContactIdentifier(value string)
	Handle() CXHandle
	SetHandle(value ICXHandle)
	Video() bool
	SetVideo(value bool)
	IsVideo() bool
	SetIsVideo(value bool)
}

// An encapsulation of the act of initiating an outgoing call.
//
// is a concrete subclass of . When the user initiates an outgoing call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call started at a time other than the current time, you can instead call the .


// An encapsulation of the act of initiating an outgoing call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction

type CXStartCallAction struct {
	CXCallAction
}

// CXStartCallActionFrom constructs a [CXStartCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of initiating an outgoing call.
func CXStartCallActionFrom(ptr unsafe.Pointer) CXStartCallAction {
	return CXStartCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXStartCallActionClass) Alloc() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXStartCallActionClass) New() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXStartCallAction) Init() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXStartCallAction) Autorelease() CXStartCallAction {
	rv := objc.Send[CXStartCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXStartCallAction creates a new CXStartCallAction instance.
func NewCXStartCallAction() CXStartCallAction {
	return getCXStartCallActionClass().New()
}




// Initializes a new action to start a call with the specified UUID to a recipient with the specified handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/init(call:handle:)

func NewCXStartCallActionWithCallUUIDHandle(callUUID foundation.IUUID, handle ICXHandle) CXStartCallAction {
	instance := getCXStartCallActionClass().Alloc()
	rv := objc.Send[CXStartCallAction](instance.ID, objc.Sel("initWithCallUUID:handle:"), callUUID, handle)
	rv.Autorelease()
	return rv
}



// Creates a new action to start a call with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/init(coder:)

func NewCXStartCallActionWithCoder(aDecoder foundation.ICoder) CXStartCallAction {
	instance := getCXStartCallActionClass().Alloc()
	rv := objc.Send[CXStartCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}




// Reports the successful execution of the action at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/fulfill(withDateStarted:)

func (c_ CXStartCallAction) FulfillWithDateStarted(dateStarted foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateStarted:"), dateStarted)
}


// The identifier for the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/contactIdentifier

func (c_ CXStartCallAction) ContactIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}


// The identifier for the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/contactIdentifier

func (c_ CXStartCallAction) SetContactIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactIdentifier:"), objc.String(value))
}


// The handle of the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/handle

func (c_ CXStartCallAction) Handle() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("handle"))
	return rv
}


// The handle of the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/handle

func (c_ CXStartCallAction) SetHandle(value ICXHandle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHandle:"), value)
}


// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/isVideo

func (c_ CXStartCallAction) Video() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("video"))
	return rv
}


// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/isVideo

func (c_ CXStartCallAction) SetVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideo:"), value)
}


// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxstartcallaction/isvideo

func (c_ CXStartCallAction) IsVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isVideo"))
	return rv
}


// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxstartcallaction/isvideo

func (c_ CXStartCallAction) SetIsVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsVideo:"), value)
}


