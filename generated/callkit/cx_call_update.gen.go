// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallUpdate] class.
var (
	CXCallUpdateClass     _CXCallUpdateClass
	CXCallUpdateClassOnce sync.Once
)

func getCXCallUpdateClass() _CXCallUpdateClass {
	CXCallUpdateClassOnce.Do(func() {
		CXCallUpdateClass = _CXCallUpdateClass{objc.GetClass("CXCallUpdate")}
	})
	return CXCallUpdateClass
}

type _CXCallUpdateClass struct {
	class objc.Class
}

// An interface definition for the [CXCallUpdate] class.
type ICXCallUpdate interface {
	objectivec.IObject
	HasVideo() bool
	SetHasVideo(value bool)
	LocalizedCallerName() string
	SetLocalizedCallerName(value string)
	RemoteHandle() CXHandle
	SetRemoteHandle(value ICXHandle)
	SupportsDTMF() bool
	SetSupportsDTMF(value bool)
	SupportsGrouping() bool
	SetSupportsGrouping(value bool)
	SupportsHolding() bool
	SetSupportsHolding(value bool)
	SupportsUngrouping() bool
	SetSupportsUngrouping(value bool)
}

// An encapsulation of new and changed information about a call.
//
// objects are used by the system to communicate changes to calls over time. Not every property on a object must be set each time, as each object includes only new and changed information. For example, when a call is started, only some properties may be known and included in the first object sent to the system, such as . Later in the same call, other properties may change; for example, a call may be upgraded from audio only to audio and video, which would be reflected by a new object with its property set to . When an incoming call is received, you construct a object specifying a and pass that to the method to notify the telephony provider. When an active call is updated, you construct a object specifying any updated information and pass that to the method. For example, if a user changes their contact information during a call, you could notify the telephony provider of this change using a new object with the new value set to its property.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate
type CXCallUpdate struct {
	objectivec.Object
}

// CXCallUpdateFrom constructs a [CXCallUpdate] from an unsafe.Pointer.
//
// An encapsulation of new and changed information about a call.
func CXCallUpdateFrom(ptr unsafe.Pointer) CXCallUpdate {
	return CXCallUpdate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallUpdateClass) Alloc() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallUpdateClass) New() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallUpdate) Init() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallUpdate) Autorelease() CXCallUpdate {
	rv := objc.Send[CXCallUpdate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallUpdate creates a new CXCallUpdate instance.
func NewCXCallUpdate() CXCallUpdate {
	return getCXCallUpdateClass().New()
}


// A Boolean value that indicates whether the call includes video in addition to audio.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/hasVideo
func (c_ CXCallUpdate) HasVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVideo"))
	return rv
}


// SetHasVideo sets the value of the hasVideo property.
// A Boolean value that indicates whether the call includes video in addition to audio.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/hasVideo
func (c_ CXCallUpdate) SetHasVideo(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasVideo:"), value)
}

// The localized name of the caller.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/localizedCallerName
func (c_ CXCallUpdate) LocalizedCallerName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedCallerName"))
	return rv
}


// SetLocalizedCallerName sets the value of the localizedCallerName property.
// The localized name of the caller.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/localizedCallerName
func (c_ CXCallUpdate) SetLocalizedCallerName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedCallerName:"), objc.String(value))
}

// The handle for the remote party (for an incoming call, this is the caller; for an outgoing call, this is the callee).
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/remoteHandle
func (c_ CXCallUpdate) RemoteHandle() CXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("remoteHandle"))
	return rv
}


// SetRemoteHandle sets the value of the remoteHandle property.
// The handle for the remote party (for an incoming call, this is the caller; for an outgoing call, this is the callee).

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/remoteHandle
func (c_ CXCallUpdate) SetRemoteHandle(value ICXHandle) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRemoteHandle:"), value)
}

// A Boolean value that indicates whether the call can send DTMF (dual tone multifrequency) tones via hard pause digits or in-call keypad entries.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsDTMF
func (c_ CXCallUpdate) SupportsDTMF() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsDTMF"))
	return rv
}


// SetSupportsDTMF sets the value of the supportsDTMF property.
// A Boolean value that indicates whether the call can send DTMF (dual tone multifrequency) tones via hard pause digits or in-call keypad entries.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsDTMF
func (c_ CXCallUpdate) SetSupportsDTMF(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsDTMF:"), value)
}

// A Boolean value that indicates whether the call can be grouped with other calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsGrouping
func (c_ CXCallUpdate) SupportsGrouping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsGrouping"))
	return rv
}


// SetSupportsGrouping sets the value of the supportsGrouping property.
// A Boolean value that indicates whether the call can be grouped with other calls.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsGrouping
func (c_ CXCallUpdate) SetSupportsGrouping(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsGrouping:"), value)
}

// A Boolean value that indicates whether the call can be placed on hold or removed from hold.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsHolding
func (c_ CXCallUpdate) SupportsHolding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsHolding"))
	return rv
}


// SetSupportsHolding sets the value of the supportsHolding property.
// A Boolean value that indicates whether the call can be placed on hold or removed from hold.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsHolding
func (c_ CXCallUpdate) SetSupportsHolding(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsHolding:"), value)
}

// A Boolean value that indicates whether the call can be ungrouped from other calls.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsUngrouping
func (c_ CXCallUpdate) SupportsUngrouping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsUngrouping"))
	return rv
}


// SetSupportsUngrouping sets the value of the supportsUngrouping property.
// A Boolean value that indicates whether the call can be ungrouped from other calls.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsUngrouping
func (c_ CXCallUpdate) SetSupportsUngrouping(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsUngrouping:"), value)
}



