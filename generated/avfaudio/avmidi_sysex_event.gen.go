// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MIDISysexEvent] class.
var (
	MIDISysexEventClass     _MIDISysexEventClass
	MIDISysexEventClassOnce sync.Once
)

func getMIDISysexEventClass() _MIDISysexEventClass {
	MIDISysexEventClassOnce.Do(func() {
		MIDISysexEventClass = _MIDISysexEventClass{objc.GetClass("AVMIDISysexEvent")}
	})
	return MIDISysexEventClass
}

type _MIDISysexEventClass struct {
	class objc.Class
}





// An interface definition for the [MIDISysexEvent] class.
type IMIDISysexEvent interface {
	IMusicEvent
	

	// properties:
	SizeInBytes() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MIDISysexEventClass) Alloc() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MIDISysexEventClass) New() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDISysexEvent) Init() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDISysexEvent) Autorelease() MIDISysexEvent {
	rv := objc.Send[MIDISysexEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDISysexEvent creates a new MIDISysexEvent instance.
func NewMIDISysexEvent() MIDISysexEvent {
	return getMIDISysexEventClass().New()
}





// An object that represents a MIDI system exclusive message.
//
// You can’t modify the size and contents of this event once you create it.


// An object that represents a MIDI system exclusive message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent
type MIDISysexEvent struct {
	MusicEvent
}

// MIDISysexEventFrom constructs a [MIDISysexEvent] from an unsafe.Pointer.
//
// An object that represents a MIDI system exclusive message.
func MIDISysexEventFrom(ptr unsafe.Pointer) MIDISysexEvent {
	return MIDISysexEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates a system event with the data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/init(data:)
func NewMIDISysexEventWithData(data objc.IObject /* cross-framework: NSData */) MIDISysexEvent {
	instance := getMIDISysexEventClass().Alloc()
	rv := objc.Send[MIDISysexEvent](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}






















// The size of the data that this event contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/sizeInBytes
func (m_ MIDISysexEvent) SizeInBytes() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("sizeInBytes"))
	return rv
}







