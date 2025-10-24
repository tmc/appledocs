//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCallUpdate


// iOS-only properties

// A Boolean value that indicates whether the call includes video in addition to audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/hasVideo
func (c_ CXCallUpdate) HasVideo() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasVideo"))
	return rv
}
func (c_ CXCallUpdate) SetHasVideo(value bool) {
	c_.ID.Send(objc.RegisterName("setHasVideo:"), value)
}

// The localized name of the caller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/localizedCallerName
func (c_ CXCallUpdate) LocalizedCallerName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("localizedCallerName"))
	return rv
}
func (c_ CXCallUpdate) SetLocalizedCallerName(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setLocalizedCallerName:"), value)
}

// The handle for the remote party (for an incoming call, this is the caller; for an outgoing call, this is the callee).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/remoteHandle
func (c_ CXCallUpdate) RemoteHandle() ICXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("remoteHandle"))
	return rv
}
func (c_ CXCallUpdate) SetRemoteHandle(value ICXHandle) {
	c_.ID.Send(objc.RegisterName("setRemoteHandle:"), value)
}

// A Boolean value that indicates whether the call can send DTMF (dual tone multifrequency) tones via hard pause digits or in-call keypad entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsDTMF
func (c_ CXCallUpdate) SupportsDTMF() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsDTMF"))
	return rv
}
func (c_ CXCallUpdate) SetSupportsDTMF(value bool) {
	c_.ID.Send(objc.RegisterName("setSupportsDTMF:"), value)
}

// A Boolean value that indicates whether the call can be grouped with other calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsGrouping
func (c_ CXCallUpdate) SupportsGrouping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsGrouping"))
	return rv
}
func (c_ CXCallUpdate) SetSupportsGrouping(value bool) {
	c_.ID.Send(objc.RegisterName("setSupportsGrouping:"), value)
}

// A Boolean value that indicates whether the call can be placed on hold or removed from hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsHolding
func (c_ CXCallUpdate) SupportsHolding() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsHolding"))
	return rv
}
func (c_ CXCallUpdate) SetSupportsHolding(value bool) {
	c_.ID.Send(objc.RegisterName("setSupportsHolding:"), value)
}

// A Boolean value that indicates whether the call can be ungrouped from other calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallUpdate/supportsUngrouping
func (c_ CXCallUpdate) SupportsUngrouping() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("supportsUngrouping"))
	return rv
}
func (c_ CXCallUpdate) SetSupportsUngrouping(value bool) {
	c_.ID.Send(objc.RegisterName("setSupportsUngrouping:"), value)
}





