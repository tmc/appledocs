//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MusicPlayerMediaItemQueueDescriptor


// The time the designated media item is to stop playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/setEndTime(_:for:)
func (m_ MusicPlayerMediaItemQueueDescriptor) SetEndTimeForItem(endTime float64, mediaItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItem:"), endTime, mediaItem)
}

// The time the designated media item is to start playing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/setStartTime(_:for:)
func (m_ MusicPlayerMediaItemQueueDescriptor) SetStartTimeForItem(startTime float64, mediaItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItem:"), startTime, mediaItem)
}

// iOS-only properties

// Contains the media item collection used to create the queue descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/itemCollection
func (m_ MusicPlayerMediaItemQueueDescriptor) ItemCollection() IMPMediaItemCollection {
	rv := objc.Send[MediaItemCollection](m_.ID, objc.Sel("itemCollection"))
	return rv
}

// Contains the media items found by the query used to create the queue descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/query
func (m_ MusicPlayerMediaItemQueueDescriptor) Query() IMPMediaQuery {
	rv := objc.Send[MediaQuery](m_.ID, objc.Sel("query"))
	return rv
}

// Designates the media item to play first.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/startItem
func (m_ MusicPlayerMediaItemQueueDescriptor) StartItem() IMPMediaItem {
	rv := objc.Send[MediaItem](m_.ID, objc.Sel("startItem"))
	return rv
}
func (m_ MusicPlayerMediaItemQueueDescriptor) SetStartItem(value IMPMediaItem) {
	m_.ID.Send(objc.RegisterName("setStartItem:"), value)
}




