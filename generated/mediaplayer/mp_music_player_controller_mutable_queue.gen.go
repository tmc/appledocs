// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MusicPlayerControllerMutableQueue] class.
var (
	MusicPlayerControllerMutableQueueClass     _MusicPlayerControllerMutableQueueClass
	MusicPlayerControllerMutableQueueClassOnce sync.Once
)

func getMusicPlayerControllerMutableQueueClass() _MusicPlayerControllerMutableQueueClass {
	MusicPlayerControllerMutableQueueClassOnce.Do(func() {
		MusicPlayerControllerMutableQueueClass = _MusicPlayerControllerMutableQueueClass{objc.GetClass("MPMusicPlayerControllerMutableQueue")}
	})
	return MusicPlayerControllerMutableQueueClass
}

type _MusicPlayerControllerMutableQueueClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerControllerMutableQueue] class.
type IMusicPlayerControllerMutableQueue interface {
	IMusicPlayerControllerQueue
	InsertQueueDescriptorAfterItem(queueDescriptor IMPMusicPlayerQueueDescriptor, afterItem IMPMediaItem)
	RemoveItem(item IMPMediaItem)
}

// A mutable queue containing the media items to play.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue
type MusicPlayerControllerMutableQueue struct {
	MusicPlayerControllerQueue
}

// MusicPlayerControllerMutableQueueFrom constructs a [MusicPlayerControllerMutableQueue] from an unsafe.Pointer.
//
// A mutable queue containing the media items to play.
func MusicPlayerControllerMutableQueueFrom(ptr unsafe.Pointer) MusicPlayerControllerMutableQueue {
	return MusicPlayerControllerMutableQueue{
		MusicPlayerControllerQueue: MusicPlayerControllerQueueFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerMutableQueueClass) Alloc() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerControllerMutableQueueClass) New() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerControllerMutableQueue) Init() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerControllerMutableQueue) Autorelease() MusicPlayerControllerMutableQueue {
	rv := objc.Send[MusicPlayerControllerMutableQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerControllerMutableQueue creates a new MusicPlayerControllerMutableQueue instance.
func NewMusicPlayerControllerMutableQueue() MusicPlayerControllerMutableQueue {
	return getMusicPlayerControllerMutableQueueClass().New()
}


// Inserts a modified queue after the designated media item.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue/insert(_:after:)
func (m_ MusicPlayerControllerMutableQueue) InsertQueueDescriptorAfterItem(queueDescriptor IMPMusicPlayerQueueDescriptor, afterItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertQueueDescriptor:afterItem:"), queueDescriptor, afterItem)
}

// Removes a media item from the music player’s queue.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerMutableQueue/remove(_:)
func (m_ MusicPlayerControllerMutableQueue) RemoveItem(item IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeItem:"), item)
}



