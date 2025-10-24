// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioConnectionPoint] class.
var (
	AudioConnectionPointClass     _AudioConnectionPointClass
	AudioConnectionPointClassOnce sync.Once
)

func getAudioConnectionPointClass() _AudioConnectionPointClass {
	AudioConnectionPointClassOnce.Do(func() {
		AudioConnectionPointClass = _AudioConnectionPointClass{objc.GetClass("AVAudioConnectionPoint")}
	})
	return AudioConnectionPointClass
}

type _AudioConnectionPointClass struct {
	class objc.Class
}





// An interface definition for the [AudioConnectionPoint] class.
type IAudioConnectionPoint interface {
	objectivec.IObject
	

	// properties:
	Bus() AudioNodeBus /* typedef */
	Node() IAVAudioNode


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioConnectionPointClass) Alloc() AudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioConnectionPointClass) New() AudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioConnectionPoint) Init() AudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioConnectionPoint) Autorelease() AudioConnectionPoint {
	rv := objc.Send[AudioConnectionPoint](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioConnectionPoint creates a new AudioConnectionPoint instance.
func NewAudioConnectionPoint() AudioConnectionPoint {
	return getAudioConnectionPointClass().New()
}





// A representation of either a source or destination connection point in the audio engine.
//
// Instances of this class are immutable.


// A representation of either a source or destination connection point in the audio engine.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint
type AudioConnectionPoint struct {
	objectivec.Object
}

// AudioConnectionPointFrom constructs a [AudioConnectionPoint] from an unsafe.Pointer.
//
// A representation of either a source or destination connection point in the audio engine.
func AudioConnectionPointFrom(ptr unsafe.Pointer) AudioConnectionPoint {
	return AudioConnectionPoint{objectivec.Object{objc.ID(ptr)}}
}






// Creates a connection point object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/init(node:bus:)
func NewAudioConnectionPointWithNodeBus(node IAVAudioNode, bus AudioNodeBus /* typedef */) AudioConnectionPoint {
	instance := getAudioConnectionPointClass().Alloc()
	rv := objc.Send[AudioConnectionPoint](instance.ID, objc.Sel("initWithNode:bus:"), node, bus)
	rv.Autorelease()
	return rv
}






















// The bus on the node in the connection point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/bus
func (a_ AudioConnectionPoint) Bus() AudioNodeBus /* typedef */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("bus"))
	return rv
}


// The node in the connection point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioConnectionPoint/node
func (a_ AudioConnectionPoint) Node() IAVAudioNode {
	rv := objc.Send[AudioNode](a_.ID, objc.Sel("node"))
	return rv
}







