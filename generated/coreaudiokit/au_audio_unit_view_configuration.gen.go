// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitViewConfiguration] class.
var (
	AudioUnitViewConfigurationClass     _AudioUnitViewConfigurationClass
	AudioUnitViewConfigurationClassOnce sync.Once
)

func getAudioUnitViewConfigurationClass() _AudioUnitViewConfigurationClass {
	AudioUnitViewConfigurationClassOnce.Do(func() {
		AudioUnitViewConfigurationClass = _AudioUnitViewConfigurationClass{objc.GetClass("AUAudioUnitViewConfiguration")}
	})
	return AudioUnitViewConfigurationClass
}

type _AudioUnitViewConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitViewConfiguration] class.
type IAudioUnitViewConfiguration interface {
	objectivec.IObject
}

// A configuration object that describes how to present the audio unit’s user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUAudioUnitViewConfiguration
type AudioUnitViewConfiguration struct {
	objectivec.Object
}

// AudioUnitViewConfigurationFrom constructs a [AudioUnitViewConfiguration] from an unsafe.Pointer.
//
// A configuration object that describes how to present the audio unit’s user interface.
func AudioUnitViewConfigurationFrom(ptr unsafe.Pointer) AudioUnitViewConfiguration {
	return AudioUnitViewConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitViewConfigurationClass) Alloc() AudioUnitViewConfiguration {
	rv := objc.Send[AudioUnitViewConfiguration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitViewConfigurationClass) New() AudioUnitViewConfiguration {
	rv := objc.Send[AudioUnitViewConfiguration](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitViewConfiguration) Init() AudioUnitViewConfiguration {
	rv := objc.Send[AudioUnitViewConfiguration](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitViewConfiguration) Autorelease() AudioUnitViewConfiguration {
	rv := objc.Send[AudioUnitViewConfiguration](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitViewConfiguration creates a new AudioUnitViewConfiguration instance.
func NewAudioUnitViewConfiguration() AudioUnitViewConfiguration {
	return getAudioUnitViewConfigurationClass().New()
}




// Creates a new configuration object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUAudioUnitViewConfiguration/init(width:height:hostHasController:)
func NewAudioUnitViewConfigurationWithWidthHeightHostHasController(width float64, height float64, hostHasController bool) AudioUnitViewConfiguration {
	instance := getAudioUnitViewConfigurationClass().Alloc()
	rv := objc.Send[AudioUnitViewConfiguration](instance.ID, objc.Sel("initWithWidth:height:hostHasController:"), width, height, hostHasController)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether the host shows its own control surface in this view configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/auaudiounitviewconfiguration/hosthascontroller
func (a_ AudioUnitViewConfiguration) HostHasController() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hostHasController"))
	return rv
}


// SetHostHasController sets the value of the hostHasController property.
// A Boolean value that indicates whether the host shows its own control surface in this view configuration.

//
// [Full Topic]: https://developer.apple.com/documentation/coreaudiokit/auaudiounitviewconfiguration/hosthascontroller
func (a_ AudioUnitViewConfiguration) SetHostHasController(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHostHasController:"), value)
}

// The configured height.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUAudioUnitViewConfiguration/height
func (a_ AudioUnitViewConfiguration) Height() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("height"))
	return rv
}

// The configured width.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioKit/AUAudioUnitViewConfiguration/width
func (a_ AudioUnitViewConfiguration) Width() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("width"))
	return rv
}


