// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AudioUnitBus] class.
var (
	AudioUnitBusClass     _AudioUnitBusClass
	AudioUnitBusClassOnce sync.Once
)

func getAudioUnitBusClass() _AudioUnitBusClass {
	AudioUnitBusClassOnce.Do(func() {
		AudioUnitBusClass = _AudioUnitBusClass{objc.GetClass("AUAudioUnitBus")}
	})
	return AudioUnitBusClass
}

type _AudioUnitBusClass struct {
	class objc.Class
}

// An interface definition for the [AudioUnitBus] class.
type IAudioUnitBus interface {
	objectivec.IObject
	SetFormatError(format avfaudio.IAudioFormat, outError unsafe.Pointer) bool
	BusType() AudioUnitBusType
	ContextPresentationLatency() foundation.TimeInterval
	SetContextPresentationLatency(value foundation.ITimeInterval)
	Format() avfaudio.AudioFormat
	Index() uint
	Enabled() bool
	SetEnabled(value bool)
	MaximumChannelCount() AudioChannelCount
	SetMaximumChannelCount(value IAudioChannelCount)
	Name() string
	SetName(value string)
	OwnerAudioUnit() AUAudioUnit
	ShouldAllocateBuffer() bool
	SetShouldAllocateBuffer(value bool)
	SupportedChannelCounts() []foundation.Number
	SetSupportedChannelCounts(value []foundation.INumber)
	SupportedChannelLayoutTags() []foundation.Number
	IsEnabled() bool
	SetIsEnabled(value bool)
}

// A class that defines an input or output connection point on an audio unit.


// A class that defines an input or output connection point on an audio unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus
type AudioUnitBus struct {
	objectivec.Object
}

// AudioUnitBusFrom constructs a [AudioUnitBus] from an unsafe.Pointer.
//
// A class that defines an input or output connection point on an audio unit.
func AudioUnitBusFrom(ptr unsafe.Pointer) AudioUnitBus {
	return AudioUnitBus{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AudioUnitBusClass) Alloc() AudioUnitBus {
	rv := objc.Send[AudioUnitBus](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AudioUnitBusClass) New() AudioUnitBus {
	rv := objc.Send[AudioUnitBus](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitBus) Init() AudioUnitBus {
	rv := objc.Send[AudioUnitBus](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitBus) Autorelease() AudioUnitBus {
	rv := objc.Send[AudioUnitBus](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitBus creates a new AudioUnitBus instance.
func NewAudioUnitBus() AudioUnitBus {
	return getAudioUnitBusClass().New()
}



// Initializes a bus object with a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/init(format:)
func NewAudioUnitBusWithFormatError(format avfaudio.IAudioFormat, outError unsafe.Pointer) AudioUnitBus {
	instance := getAudioUnitBusClass().Alloc()
	rv := objc.Send[AudioUnitBus](instance.ID, objc.Sel("initWithFormat:error:"), format, outError)
	rv.Autorelease()
	return rv
}



// Sets the bus’s audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/setFormat(_:)
func (a_ AudioUnitBus) SetFormatError(format avfaudio.IAudioFormat, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setFormat:error:"), format, outError)
	return rv
}


// The bus type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/busType
func (a_ AudioUnitBus) BusType() AudioUnitBusType {
	rv := objc.Send[AudioUnitBusType](a_.ID, objc.Sel("busType"))
	return rv
}


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/contextPresentationLatency
func (a_ AudioUnitBus) ContextPresentationLatency() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](a_.ID, objc.Sel("contextPresentationLatency"))
	return rv
}


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/contextPresentationLatency
func (a_ AudioUnitBus) SetContextPresentationLatency(value foundation.ITimeInterval) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContextPresentationLatency:"), value)
}


// The audio format and channel layout of audio being transferred on the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/format
func (a_ AudioUnitBus) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}


// The index of this bus in its containing array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/index
func (a_ AudioUnitBus) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/isEnabled
func (a_ AudioUnitBus) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/isEnabled
func (a_ AudioUnitBus) SetEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnabled:"), value)
}


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/maximumChannelCount
func (a_ AudioUnitBus) MaximumChannelCount() AudioChannelCount {
	rv := objc.Send[AudioChannelCount](a_.ID, objc.Sel("maximumChannelCount"))
	return rv
}


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/maximumChannelCount
func (a_ AudioUnitBus) SetMaximumChannelCount(value IAudioChannelCount) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumChannelCount:"), value)
}


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/name
func (a_ AudioUnitBus) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/name
func (a_ AudioUnitBus) SetName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}


// The audio unit that owns the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/ownerAudioUnit
func (a_ AudioUnitBus) OwnerAudioUnit() AUAudioUnit {
	rv := objc.Send[AUAudioUnit](a_.ID, objc.Sel("ownerAudioUnit"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/shouldAllocateBuffer
func (a_ AudioUnitBus) ShouldAllocateBuffer() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldAllocateBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/shouldAllocateBuffer
func (a_ AudioUnitBus) SetShouldAllocateBuffer(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldAllocateBuffer:"), value)
}


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelCounts
func (a_ AudioUnitBus) SupportedChannelCounts() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("supportedChannelCounts"))
	return rv
}


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelCounts
func (a_ AudioUnitBus) SetSupportedChannelCounts(value []foundation.INumber) {
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
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportedChannelCounts:"), nsArray)
}


// An array of audio channel layout tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelLayoutTags
func (a_ AudioUnitBus) SupportedChannelLayoutTags() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("supportedChannelLayoutTags"))
	return rv
}


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/isenabled
func (a_ AudioUnitBus) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/isenabled
func (a_ AudioUnitBus) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}


