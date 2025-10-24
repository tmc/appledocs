// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MusicPlayerMediaItemQueueDescriptor] class.
var (
	MusicPlayerMediaItemQueueDescriptorClass     _MusicPlayerMediaItemQueueDescriptorClass
	MusicPlayerMediaItemQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerMediaItemQueueDescriptorClass() _MusicPlayerMediaItemQueueDescriptorClass {
	MusicPlayerMediaItemQueueDescriptorClassOnce.Do(func() {
		MusicPlayerMediaItemQueueDescriptorClass = _MusicPlayerMediaItemQueueDescriptorClass{objc.GetClass("MPMusicPlayerMediaItemQueueDescriptor")}
	})
	return MusicPlayerMediaItemQueueDescriptorClass
}

type _MusicPlayerMediaItemQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerMediaItemQueueDescriptor] class.
type IMusicPlayerMediaItemQueueDescriptor interface {
	IMusicPlayerQueueDescriptor
	// properties:
	// methods:
}

// A set of properties and methods for modifying audio media items in the player’s media queue.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item to play.


// A set of properties and methods for modifying audio media items in the player’s media queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor
type MusicPlayerMediaItemQueueDescriptor struct {
	MusicPlayerQueueDescriptor
}

// MusicPlayerMediaItemQueueDescriptorFrom constructs a [MusicPlayerMediaItemQueueDescriptor] from an unsafe.Pointer.
//
// A set of properties and methods for modifying audio media items in the player’s media queue.
func MusicPlayerMediaItemQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerMediaItemQueueDescriptor {
	return MusicPlayerMediaItemQueueDescriptor{
		MusicPlayerQueueDescriptor: MusicPlayerQueueDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerMediaItemQueueDescriptorClass) Alloc() MusicPlayerMediaItemQueueDescriptor {
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerMediaItemQueueDescriptorClass) New() MusicPlayerMediaItemQueueDescriptor {
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerMediaItemQueueDescriptor) Init() MusicPlayerMediaItemQueueDescriptor {
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerMediaItemQueueDescriptor) Autorelease() MusicPlayerMediaItemQueueDescriptor {
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerMediaItemQueueDescriptor creates a new MusicPlayerMediaItemQueueDescriptor instance.
func NewMusicPlayerMediaItemQueueDescriptor() MusicPlayerMediaItemQueueDescriptor {
	return getMusicPlayerMediaItemQueueDescriptorClass().New()
}



// Creates a new queue descriptor using the designated collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/init(itemCollection:)
func NewMusicPlayerMediaItemQueueDescriptorWithItemCollection(itemCollection IMPMediaItemCollection) MusicPlayerMediaItemQueueDescriptor {
	instance := getMusicPlayerMediaItemQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](instance.ID, objc.Sel("initWithItemCollection:"), itemCollection)
	rv.Autorelease()
	return rv
}


// Creates a new queue descriptor using the designated query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/init(query:)
func NewMusicPlayerMediaItemQueueDescriptorWithQuery(query IMPMediaQuery) MusicPlayerMediaItemQueueDescriptor {
	instance := getMusicPlayerMediaItemQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](instance.ID, objc.Sel("initWithQuery:"), query)
	rv.Autorelease()
	return rv
}



