//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for MusicPlayerControllerQueue


// iOS-only properties

// The media items in the queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerQueue/items
func (m_ MusicPlayerControllerQueue) Items() []IMediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}





