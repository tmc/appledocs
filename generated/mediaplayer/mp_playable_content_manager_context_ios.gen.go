//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PlayableContentManagerContext


// iOS-only properties

// A Boolean value that indicates whether the content server enables content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext/contentLimitsEnabled
func (p_ PlayableContentManagerContext) ContentLimitsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentLimitsEnabled"))
	return rv
}

// A Boolean value that indicates whether the content server enforces content limits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext/contentLimitsEnforced
func (p_ PlayableContentManagerContext) ContentLimitsEnforced() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("contentLimitsEnforced"))
	return rv
}

// Returns a Boolean that indicates whether the content server is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext/endpointAvailable
func (p_ PlayableContentManagerContext) EndpointAvailable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("endpointAvailable"))
	return rv
}

// Returns the number of content items to display during content limiting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext/enforcedContentItemsCount
func (p_ PlayableContentManagerContext) EnforcedContentItemsCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("enforcedContentItemsCount"))
	return rv
}

// The maximum depth of the navigation hierarchy allowed by the content server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext/enforcedContentTreeDepth
func (p_ PlayableContentManagerContext) EnforcedContentTreeDepth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("enforcedContentTreeDepth"))
	return rv
}





