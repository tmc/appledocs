//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NavigationMarkersGroup


// iOS-only properties

// The array of date range navigation markers for which the group provides navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/dateRangeNavigationMarkers
func (n_ NavigationMarkersGroup) DateRangeNavigationMarkers() []objc.IObject /* cross-framework: DateRangeMetadataGroup */ {
	rv := objc.Send[[]avfoundation.DateRangeMetadataGroup](n_.ID, objc.Sel("dateRangeNavigationMarkers"))
	return rv
}

// The array of timed navigation markers for which the group provides navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/timedNavigationMarkers
func (n_ NavigationMarkersGroup) TimedNavigationMarkers() []objc.IObject /* cross-framework: TimedMetadataGroup */ {
	rv := objc.Send[[]avfoundation.TimedMetadataGroup](n_.ID, objc.Sel("timedNavigationMarkers"))
	return rv
}

// The title of the marker group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVNavigationMarkersGroup/title
func (n_ NavigationMarkersGroup) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("title"))
	return rv
}




