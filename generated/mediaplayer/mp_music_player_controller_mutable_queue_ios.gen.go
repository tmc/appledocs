//go:build darwin && ios

// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MusicPlayerControllerMutableQueue


// Inserts a modified queue after the designated media item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue/insert(_:after:)
func (m_ MusicPlayerControllerMutableQueue) InsertQueueDescriptorAfterItem(queueDescriptor IMPMusicPlayerQueueDescriptor, afterItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertQueueDescriptor:afterItem:"), queueDescriptor, afterItem)
}

// Removes a media item from the music player’s queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue/remove(_:)
func (m_ MusicPlayerControllerMutableQueue) RemoveItem(item IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItem:"), item)
}

// iOS-only properties





