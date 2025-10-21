// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MusicPlayerQueueDescriptor] class.
var (
	MusicPlayerQueueDescriptorClass     _MusicPlayerQueueDescriptorClass
	MusicPlayerQueueDescriptorClassOnce sync.Once
)

func getMusicPlayerQueueDescriptorClass() _MusicPlayerQueueDescriptorClass {
	MusicPlayerQueueDescriptorClassOnce.Do(func() {
		MusicPlayerQueueDescriptorClass = _MusicPlayerQueueDescriptorClass{objc.GetClass("MPMusicPlayerQueueDescriptor")}
	})
	return MusicPlayerQueueDescriptorClass
}

type _MusicPlayerQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerQueueDescriptor] class.
type IMusicPlayerQueueDescriptor interface {
	objectivec.IObject
}

// The abstract base class for audio media item and store queue descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerQueueDescriptor
type MusicPlayerQueueDescriptor struct {
	objectivec.Object
}

// MusicPlayerQueueDescriptorFrom constructs a [MusicPlayerQueueDescriptor] from an unsafe.Pointer.
//
// The abstract base class for audio media item and store queue descriptors.
func MusicPlayerQueueDescriptorFrom(ptr unsafe.Pointer) MusicPlayerQueueDescriptor {
	return MusicPlayerQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerQueueDescriptorClass) Alloc() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerQueueDescriptorClass) New() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerQueueDescriptor) Init() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerQueueDescriptor) Autorelease() MusicPlayerQueueDescriptor {
	rv := objc.Send[MusicPlayerQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerQueueDescriptor creates a new MusicPlayerQueueDescriptor instance.
func NewMusicPlayerQueueDescriptor() MusicPlayerQueueDescriptor {
	return getMusicPlayerQueueDescriptorClass().New()
}




