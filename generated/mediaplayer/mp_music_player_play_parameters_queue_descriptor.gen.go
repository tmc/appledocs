// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MusicPlayerPlayParametersQueueDescriptor] class.
var (
	MusicPlayerPlayParametersQueueDescriptorClass     _MusicPlayerPlayParametersQueueDescriptorClass
	MusicPlayerPlayParametersQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerPlayParametersQueueDescriptorClass() _MusicPlayerPlayParametersQueueDescriptorClass {
	MusicPlayerPlayParametersQueueDescriptorClassOnce.Do(func() {
		MusicPlayerPlayParametersQueueDescriptorClass = _MusicPlayerPlayParametersQueueDescriptorClass{objc.GetClass("MPMusicPlayerPlayParametersQueueDescriptor")}
	})
	return MusicPlayerPlayParametersQueueDescriptorClass
}

type _MusicPlayerPlayParametersQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerPlayParametersQueueDescriptor] class.
type IMusicPlayerPlayParametersQueueDescriptor interface {
	IMusicPlayerQueueDescriptor
	SetEndTimeForItemWithPlayParameters(endTime foundation.TimeInterval, playParameters unsafe.Pointer)
	SetStartTimeForItemWithPlayParameters(startTime foundation.TimeInterval, playParameters unsafe.Pointer)
}

// A set of properties and methods for modifying how to play items, based on play parameters the framework returns.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item for playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor
type MusicPlayerPlayParametersQueueDescriptor struct {
	MusicPlayerQueueDescriptor
}

// MusicPlayerPlayParametersQueueDescriptorFrom constructs a [MusicPlayerPlayParametersQueueDescriptor] from an unsafe.Pointer.
//
// A set of properties and methods for modifying how to play items, based on play parameters the framework returns.
func MusicPlayerPlayParametersQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerPlayParametersQueueDescriptor {
	return MusicPlayerPlayParametersQueueDescriptor{
		MusicPlayerQueueDescriptor: MusicPlayerQueueDescriptorFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerPlayParametersQueueDescriptorClass) Alloc() MusicPlayerPlayParametersQueueDescriptor {
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerPlayParametersQueueDescriptorClass) New() MusicPlayerPlayParametersQueueDescriptor {
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerPlayParametersQueueDescriptor) Init() MusicPlayerPlayParametersQueueDescriptor {
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerPlayParametersQueueDescriptor) Autorelease() MusicPlayerPlayParametersQueueDescriptor {
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerPlayParametersQueueDescriptor creates a new MusicPlayerPlayParametersQueueDescriptor instance.
func NewMusicPlayerPlayParametersQueueDescriptor() MusicPlayerPlayParametersQueueDescriptor {
	return getMusicPlayerPlayParametersQueueDescriptorClass().New()
}




// Creates a new queue descriptor using the designated queue of play parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/init(playParametersQueue:)
func NewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue(playParametersQueue unsafe.Pointer) MusicPlayerPlayParametersQueueDescriptor {
	instance := getMusicPlayerPlayParametersQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](instance.ID, objc.Sel("initWithPlayParametersQueue:"), playParametersQueue)
	rv.Autorelease()
	return rv
}


// Sets the time the item with the associated play parameters is to stop playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/setEndTime(_:forItemWith:)
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetEndTimeForItemWithPlayParameters(endTime foundation.TimeInterval, playParameters unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndTime:forItemWithPlayParameters:"), endTime, playParameters)
}

// Sets the time the item with the associated play parameters is to start playing.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/setStartTime(_:forItemWith:)
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetStartTimeForItemWithPlayParameters(startTime foundation.TimeInterval, playParameters unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:forItemWithPlayParameters:"), startTime, playParameters)
}

// An array containing the play parameters returned from querying MusicKit.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/playParametersQueue
func (m_ MusicPlayerPlayParametersQueueDescriptor) PlayParametersQueue() []MusicPlayerPlayParameters {
	rv := objc.Send[[]MusicPlayerPlayParameters](m_.ID, objc.Sel("playParametersQueue"))
	return rv
}


// SetPlayParametersQueue sets the value of the playParametersQueue property.
// An array containing the play parameters returned from querying MusicKit.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/playParametersQueue
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetPlayParametersQueue(value []MusicPlayerPlayParameters) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayParametersQueue:"), nsArray)
}

// The item identified by the play parameters to play first.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/startItemPlayParameters
func (m_ MusicPlayerPlayParametersQueueDescriptor) StartItemPlayParameters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("startItemPlayParameters"))
	return rv
}


// SetStartItemPlayParameters sets the value of the startItemPlayParameters property.
// The item identified by the play parameters to play first.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/startItemPlayParameters
func (m_ MusicPlayerPlayParametersQueueDescriptor) SetStartItemPlayParameters(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartItemPlayParameters:"), value)
}


