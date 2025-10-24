//go:build darwin && ios

// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for INSearchForMediaIntent


// iOS-only properties

// The media items for which to search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForMediaIntent/mediaItems
func (i_ INSearchForMediaIntent) MediaItems() []INMediaItem {
	rv := objc.Send[[]INMediaItem](i_.ID, objc.Sel("mediaItems"))
	return rv
}





