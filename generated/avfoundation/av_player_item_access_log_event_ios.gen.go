//go:build darwin && ios

// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for PlayerItemAccessLogEvent


// iOS-only properties

// A count of the media segments downloaded from the server to this client.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemAccessLogEvent/numberOfSegmentsDownloaded
func (p_ PlayerItemAccessLogEvent) NumberOfSegmentsDownloaded() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfSegmentsDownloaded"))
	return rv
}





