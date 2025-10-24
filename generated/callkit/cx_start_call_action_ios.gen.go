//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXStartCallAction


// Reports the successful execution of the action at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/fulfill(withDateStarted:)
func (c_ CXStartCallAction) FulfillWithDateStarted(dateStarted objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateStarted:"), dateStarted)
}

// iOS-only properties

// The identifier for the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/contactIdentifier
func (c_ CXStartCallAction) ContactIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contactIdentifier"))
	return rv
}
func (c_ CXStartCallAction) SetContactIdentifier(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setContactIdentifier:"), value)
}

// The handle of the call recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/handle
func (c_ CXStartCallAction) Handle() ICXHandle {
	rv := objc.Send[CXHandle](c_.ID, objc.Sel("handle"))
	return rv
}
func (c_ CXStartCallAction) SetHandle(value ICXHandle) {
	c_.ID.Send(objc.RegisterName("setHandle:"), value)
}

// A Boolean value that indicates whether the call is a video call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXStartCallAction/isVideo
func (c_ CXStartCallAction) Video() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("video"))
	return rv
}
func (c_ CXStartCallAction) SetVideo(value bool) {
	c_.ID.Send(objc.RegisterName("setVideo:"), value)
}




