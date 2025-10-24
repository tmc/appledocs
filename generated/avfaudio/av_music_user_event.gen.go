// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MusicUserEvent] class.
var (
	MusicUserEventClass     _MusicUserEventClass
	MusicUserEventClassOnce sync.Once
)

func getMusicUserEventClass() _MusicUserEventClass {
	MusicUserEventClassOnce.Do(func() {
		MusicUserEventClass = _MusicUserEventClass{objc.GetClass("AVMusicUserEvent")}
	})
	return MusicUserEventClass
}

type _MusicUserEventClass struct {
	class objc.Class
}





// An interface definition for the [MusicUserEvent] class.
type IMusicUserEvent interface {
	IMusicEvent
	

	// properties:
	SizeInBytes() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MusicUserEventClass) Alloc() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MusicUserEventClass) New() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicUserEvent) Init() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicUserEvent) Autorelease() MusicUserEvent {
	rv := objc.Send[MusicUserEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicUserEvent creates a new MusicUserEvent instance.
func NewMusicUserEvent() MusicUserEvent {
	return getMusicUserEventClass().New()
}





// An object that represents a custom user message.
//
// When playback of an reaches this event, the system calls the track’s callback. You can’t modify the size and contents of an once you create it.


// An object that represents a custom user message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent
type MusicUserEvent struct {
	MusicEvent
}

// MusicUserEventFrom constructs a [MusicUserEvent] from an unsafe.Pointer.
//
// An object that represents a custom user message.
func MusicUserEventFrom(ptr unsafe.Pointer) MusicUserEvent {
	return MusicUserEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates a user event with the data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/init(data:)
func NewMusicUserEventWithData(data objc.IObject /* cross-framework: NSData */) MusicUserEvent {
	instance := getMusicUserEventClass().Alloc()
	rv := objc.Send[MusicUserEvent](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}






















// The size of the data that the user event represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/sizeInBytes
func (m_ MusicUserEvent) SizeInBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sizeInBytes"))
	return rv
}







