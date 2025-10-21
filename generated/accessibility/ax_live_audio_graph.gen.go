// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXLiveAudioGraph] class.
var (
	AXLiveAudioGraphClass     _AXLiveAudioGraphClass
	AXLiveAudioGraphClassOnce sync.Once
)

func getAXLiveAudioGraphClass() _AXLiveAudioGraphClass {
	AXLiveAudioGraphClassOnce.Do(func() {
		AXLiveAudioGraphClass = _AXLiveAudioGraphClass{objc.GetClass("AXLiveAudioGraph")}
	})
	return AXLiveAudioGraphClass
}

type _AXLiveAudioGraphClass struct {
	class objc.Class
}

// An interface definition for the [AXLiveAudioGraph] class.
type IAXLiveAudioGraph interface {
	objectivec.IObject
}

// An object that represents an audio graph for a live-updating, continuous data series for VoiceOver.
//
// Use to interact with an ongoing, continuous stream of data that updates with new data in real time.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph
type AXLiveAudioGraph struct {
	objectivec.Object
}

// AXLiveAudioGraphFrom constructs a [AXLiveAudioGraph] from an unsafe.Pointer.
//
// An object that represents an audio graph for a live-updating, continuous data series for VoiceOver.
func AXLiveAudioGraphFrom(ptr unsafe.Pointer) AXLiveAudioGraph {
	return AXLiveAudioGraph{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXLiveAudioGraphClass) Alloc() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXLiveAudioGraphClass) New() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXLiveAudioGraph) Init() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXLiveAudioGraph) Autorelease() AXLiveAudioGraph {
	rv := objc.Send[AXLiveAudioGraph](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXLiveAudioGraph creates a new AXLiveAudioGraph instance.
func NewAXLiveAudioGraph() AXLiveAudioGraph {
	return getAXLiveAudioGraphClass().New()
}


// Begins the live audio graph session.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/start()
func (ac _AXLiveAudioGraphClass) Start() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("start"))
}

// Ends the live audio graph session.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/stop()
func (ac _AXLiveAudioGraphClass) Stop() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("stop"))
}

// Sets the pitch of the audio graph’s tone.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXLiveAudioGraph/updateValue(_:)
func (ac _AXLiveAudioGraphClass) UpdateValue(value unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("updateValue:"), value)
}



