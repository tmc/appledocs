// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	SetEndTimeForItem(endTime foundation.ITimeInterval, mediaItem IMPMediaItem)
	SetStartTimeForItem(startTime foundation.ITimeInterval, mediaItem IMPMediaItem)
	ItemCollection() MPMediaItemCollection
	Query() MPMediaQuery
	StartItem() MPMediaItem
	SetStartItem(value IMPMediaItem)
}

// A set of properties and methods for modifying audio media items in the player’s media queue.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item to play.
//
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
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/init(itemCollection:)
func NewMusicPlayerMediaItemQueueDescriptorWithItemCollection(itemCollection IMPMediaItemCollection) MusicPlayerMediaItemQueueDescriptor {
	instance := getMusicPlayerMediaItemQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](instance.ID, objc.Sel("initWithItemCollection:"), itemCollection)
	rv.Autorelease()
	return rv
}



// Creates a new queue descriptor using the designated query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/init(query:)
func NewMusicPlayerMediaItemQueueDescriptorWithQuery(query IMPMediaQuery) MusicPlayerMediaItemQueueDescriptor {
	instance := getMusicPlayerMediaItemQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerMediaItemQueueDescriptor](instance.ID, objc.Sel("initWithQuery:"), query)
	rv.Autorelease()
	return rv
}


// The time the designated media item is to stop playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/setEndTime(_:for:)
func (m_ MusicPlayerMediaItemQueueDescriptor) SetEndTimeForItem(endTime foundation.ITimeInterval, mediaItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItem:"), endTime, mediaItem)
}

// The time the designated media item is to start playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/setStartTime(_:for:)
func (m_ MusicPlayerMediaItemQueueDescriptor) SetStartTimeForItem(startTime foundation.ITimeInterval, mediaItem IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItem:"), startTime, mediaItem)
}

// Contains the media item collection used to create the queue descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/itemCollection
func (m_ MusicPlayerMediaItemQueueDescriptor) ItemCollection() MPMediaItemCollection {
	rv := objc.Send[MPMediaItemCollection](m_.ID, objc.Sel("itemCollection"))
	return rv
}

// Contains the media items found by the query used to create the queue descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/query
func (m_ MusicPlayerMediaItemQueueDescriptor) Query() MPMediaQuery {
	rv := objc.Send[MPMediaQuery](m_.ID, objc.Sel("query"))
	return rv
}

// Designates the media item to play first.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/startItem
func (m_ MusicPlayerMediaItemQueueDescriptor) StartItem() MPMediaItem {
	rv := objc.Send[MPMediaItem](m_.ID, objc.Sel("startItem"))
	return rv
}


// SetStartItem sets the value of the startItem property.
// Designates the media item to play first.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerMediaItemQueueDescriptor/startItem
func (m_ MusicPlayerMediaItemQueueDescriptor) SetStartItem(value IMPMediaItem) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartItem:"), value)
}


