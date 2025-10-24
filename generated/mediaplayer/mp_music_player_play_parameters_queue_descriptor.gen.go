// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	// methods:
}

// A set of properties and methods for modifying how to play items, based on play parameters the framework returns.
//
// Use this class to modify the player queue created by a query before the queue begins to play. You can modify when individual items start and stop playing, along with setting the first item for playing.


// A set of properties and methods for modifying how to play items, based on play parameters the framework returns.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParametersQueueDescriptor/init(playParametersQueue:)
func NewMusicPlayerPlayParametersQueueDescriptorWithPlayParametersQueue(playParametersQueue []IMusicPlayerPlayParameters) MusicPlayerPlayParametersQueueDescriptor {
	instance := getMusicPlayerPlayParametersQueueDescriptorClass().Alloc()
	rv := objc.Send[MusicPlayerPlayParametersQueueDescriptor](instance.ID, objc.Sel("initWithPlayParametersQueue:"), playParametersQueue)
	rv.Autorelease()
	return rv
}



