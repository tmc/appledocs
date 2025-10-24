//go:build darwin && ios

// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for GCGameControllerActivationContext


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGameControllerActivationContext/previousApplicationBundleID
func (g_ GCGameControllerActivationContext) PreviousApplicationBundleID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("previousApplicationBundleID"))
	return rv
}





