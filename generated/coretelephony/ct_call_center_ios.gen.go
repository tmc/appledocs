//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CallCenter


// iOS-only properties

// A closure dispatched when a call changes state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCallCenter/callEventHandler
func (c_ CallCenter) CallEventHandler() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](c_.ID, objc.Sel("callEventHandler"))
	return rv
}
func (c_ CallCenter) SetCallEventHandler(value func(unsafe.Pointer)) {
	c_.ID.Send(objc.RegisterName("setCallEventHandler:"), value)
}

// An array representing the cellular calls in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCallCenter/currentCalls
func (c_ CallCenter) CurrentCalls() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentCalls"))
	return rv
}





