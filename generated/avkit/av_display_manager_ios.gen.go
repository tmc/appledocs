//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DisplayManager


// iOS-only properties

// A Boolean value that indicates whether the user has enabled display critera matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/isDisplayCriteriaMatchingEnabled
func (d_ DisplayManager) DisplayCriteriaMatchingEnabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("displayCriteriaMatchingEnabled"))
	return rv
}

// A Boolean value that indicates whether a display mode switch is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/isDisplayModeSwitchInProgress
func (d_ DisplayManager) DisplayModeSwitchInProgress() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("displayModeSwitchInProgress"))
	return rv
}

// A hint for the TV to set the display mode to best match the currently playing content’s display criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/preferredDisplayCriteria
func (d_ DisplayManager) PreferredDisplayCriteria() objc.IObject /* cross-framework: DisplayCriteria */ {
	rv := objc.Send[avfoundation.DisplayCriteria](d_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}
func (d_ DisplayManager) SetPreferredDisplayCriteria(value objc.IObject /* cross-framework: DisplayCriteria */) {
	d_.ID.Send(objc.RegisterName("setPreferredDisplayCriteria:"), value)
}





