// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CXSetMutedCallAction] class.
var (
	CXSetMutedCallActionClass     _CXSetMutedCallActionClass
	CXSetMutedCallActionClassOnce sync.Once
)

func getCXSetMutedCallActionClass() _CXSetMutedCallActionClass {
	CXSetMutedCallActionClassOnce.Do(func() {
		CXSetMutedCallActionClass = _CXSetMutedCallActionClass{objc.GetClass("CXSetMutedCallAction")}
	})
	return CXSetMutedCallActionClass
}

type _CXSetMutedCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXSetMutedCallAction] class.
type ICXSetMutedCallAction interface {
	ICXCallAction
}

// An encapsulation of the act of muting or unmuting a call.
//
// is a concrete subclass of . When the user or the system mutes a call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. When a caller mutes a call, that caller is unable to communicate with other callers until they unmute the call. A muted caller still receives communication from other unmuted callers.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction
type CXSetMutedCallAction struct {
	CXCallAction
}

// CXSetMutedCallActionFrom constructs a [CXSetMutedCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of muting or unmuting a call.
func CXSetMutedCallActionFrom(ptr unsafe.Pointer) CXSetMutedCallAction {
	return CXSetMutedCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXSetMutedCallActionClass) Alloc() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXSetMutedCallActionClass) New() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetMutedCallAction) Init() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetMutedCallAction) Autorelease() CXSetMutedCallAction {
	rv := objc.Send[CXSetMutedCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetMutedCallAction creates a new CXSetMutedCallAction instance.
func NewCXSetMutedCallAction() CXSetMutedCallAction {
	return getCXSetMutedCallActionClass().New()
}




// Initializes a new action for a call identified by a given UUID, as well as whether the call is muted.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/init(call:muted:)
func NewCXSetMutedCallActionWithCallUUIDMuted(callUUID unsafe.Pointer, muted bool) CXSetMutedCallAction {
	instance := getCXSetMutedCallActionClass().Alloc()
	rv := objc.Send[CXSetMutedCallAction](instance.ID, objc.Sel("initWithCallUUID:muted:"), callUUID, muted)
	rv.Autorelease()
	return rv
}



// Creates a new action for a call with data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/init(coder:)
func NewCXSetMutedCallActionWithCoder(aDecoder unsafe.Pointer) CXSetMutedCallAction {
	instance := getCXSetMutedCallActionClass().Alloc()
	rv := objc.Send[CXSetMutedCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether the call is muted.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/isMuted
func (c_ CXSetMutedCallAction) Muted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("muted"))
	return rv
}


// SetMuted sets the value of the muted property.
// A Boolean value that indicates whether the call is muted.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/isMuted
func (c_ CXSetMutedCallAction) SetMuted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMuted:"), value)
}

// A Boolean value that indicates whether the call is muted.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetmutedcallaction/ismuted
func (c_ CXSetMutedCallAction) IsMuted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isMuted"))
	return rv
}


// SetIsMuted sets the value of the isMuted property.
// A Boolean value that indicates whether the call is muted.

//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetmutedcallaction/ismuted
func (c_ CXSetMutedCallAction) SetIsMuted(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsMuted:"), value)
}


