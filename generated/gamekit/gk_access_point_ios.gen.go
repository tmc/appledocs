//go:build darwin && ios

// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for AccessPoint


// iOS-only properties

// A Boolean value that indicates whether the access point is in focus on tvOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/isFocused
func (a_ AccessPoint) Focused() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("focused"))
	return rv
}
func (a_ AccessPoint) SetFocused(value bool) {
	a_.ID.Send(objc.RegisterName("setFocused:"), value)
}





