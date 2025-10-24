// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [ExtendedTempoEvent] class.
var (
	ExtendedTempoEventClass     _ExtendedTempoEventClass
	ExtendedTempoEventClassOnce sync.Once
)

func getExtendedTempoEventClass() _ExtendedTempoEventClass {
	ExtendedTempoEventClassOnce.Do(func() {
		ExtendedTempoEventClass = _ExtendedTempoEventClass{objc.GetClass("AVExtendedTempoEvent")}
	})
	return ExtendedTempoEventClass
}

type _ExtendedTempoEventClass struct {
	class objc.Class
}





// An interface definition for the [ExtendedTempoEvent] class.
type IExtendedTempoEvent interface {
	IMusicEvent
	

	// properties:
	Tempo() float64
	SetTempo(value float64)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ec _ExtendedTempoEventClass) Alloc() ExtendedTempoEvent {
	rv := objc.Send[ExtendedTempoEvent](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExtendedTempoEventClass) New() ExtendedTempoEvent {
	rv := objc.Send[ExtendedTempoEvent](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtendedTempoEvent) Init() ExtendedTempoEvent {
	rv := objc.Send[ExtendedTempoEvent](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtendedTempoEvent) Autorelease() ExtendedTempoEvent {
	rv := objc.Send[ExtendedTempoEvent](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtendedTempoEvent creates a new ExtendedTempoEvent instance.
func NewExtendedTempoEvent() ExtendedTempoEvent {
	return getExtendedTempoEventClass().New()
}





// An object that represents a tempo change to a specific beats-per-minute value.


// An object that represents a tempo change to a specific beats-per-minute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent
type ExtendedTempoEvent struct {
	MusicEvent
}

// ExtendedTempoEventFrom constructs a [ExtendedTempoEvent] from an unsafe.Pointer.
//
// An object that represents a tempo change to a specific beats-per-minute value.
func ExtendedTempoEventFrom(ptr unsafe.Pointer) ExtendedTempoEvent {
	return ExtendedTempoEvent{
		MusicEvent: MusicEventFrom(ptr),
	}
}






// Creates an extended tempo event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/init(tempo:)
func NewExtendedTempoEventWithTempo(tempo float64) ExtendedTempoEvent {
	instance := getExtendedTempoEventClass().Alloc()
	rv := objc.Send[ExtendedTempoEvent](instance.ID, objc.Sel("initWithTempo:"), tempo)
	rv.Autorelease()
	return rv
}






















// The tempo in beats per minute as a positive value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/tempo
func (e_ ExtendedTempoEvent) Tempo() float64 {
	rv := objc.Send[float64](e_.ID, objc.Sel("tempo"))
	return rv
}


// The tempo in beats per minute as a positive value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/tempo
func (e_ ExtendedTempoEvent) SetTempo(value float64) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setTempo:"), value)
}







