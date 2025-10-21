// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MusicPlayerControllerQueue] class.
var (
	MusicPlayerControllerQueueClass     _MusicPlayerControllerQueueClass
	MusicPlayerControllerQueueClassOnce sync.Once
)

func getMusicPlayerControllerQueueClass() _MusicPlayerControllerQueueClass {
	MusicPlayerControllerQueueClassOnce.Do(func() {
		MusicPlayerControllerQueueClass = _MusicPlayerControllerQueueClass{objc.GetClass("MPMusicPlayerControllerQueue")}
	})
	return MusicPlayerControllerQueueClass
}

type _MusicPlayerControllerQueueClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerControllerQueue] class.
type IMusicPlayerControllerQueue interface {
	objectivec.IObject
}

// An immutable queue containing the media items to play.
//
// An object contains the current queue for an application queue music player. To add or remove media items from a playing queue, use . The results of the method is an object that updates the playing queue. You don’t create your own instance of this class.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerQueue
type MusicPlayerControllerQueue struct {
	objectivec.Object
}

// MusicPlayerControllerQueueFrom constructs a [MusicPlayerControllerQueue] from an unsafe.Pointer.
//
// An immutable queue containing the media items to play.
func MusicPlayerControllerQueueFrom(ptr unsafe.Pointer) MusicPlayerControllerQueue {
	return MusicPlayerControllerQueue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerControllerQueueClass) Alloc() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerControllerQueueClass) New() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerControllerQueue) Init() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerControllerQueue) Autorelease() MusicPlayerControllerQueue {
	rv := objc.Send[MusicPlayerControllerQueue](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerControllerQueue creates a new MusicPlayerControllerQueue instance.
func NewMusicPlayerControllerQueue() MusicPlayerControllerQueue {
	return getMusicPlayerControllerQueueClass().New()
}


// The media items in the queue.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerControllerQueue/items
func (m_ MusicPlayerControllerQueue) Items() []MediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}



