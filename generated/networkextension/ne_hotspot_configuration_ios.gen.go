//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotConfiguration


// iOS-only properties

// Restricts the lifetime of a configuration to the operating status of the app that created it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotConfiguration/joinOnce
func (n_ NEHotspotConfiguration) JoinOnce() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("joinOnce"))
	return rv
}
func (n_ NEHotspotConfiguration) SetJoinOnce(value bool) {
	n_.ID.Send(objc.RegisterName("setJoinOnce:"), value)
}




