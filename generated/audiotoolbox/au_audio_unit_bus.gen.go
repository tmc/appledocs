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
	Index() uint
	BusType() AUAudioUnitBusType
	SetBusType(value AUAudioUnitBusType)
	ContextPresentationLatency() unsafe.Pointer
	SetContextPresentationLatency(value unsafe.Pointer)
	Format() avfaudio.AudioFormat
	SetFormat(value avfaudio.AudioFormat)
	IsEnabled() bool
	SetIsEnabled(value bool)
	MaximumChannelCount() AudioChannelCount
	SetMaximumChannelCount(value AudioChannelCount)
	Name() string
	SetName(value string)
	OwnerAudioUnit() IAUAudioUnit
	SetOwnerAudioUnit(value IAUAudioUnit)
	ShouldAllocateBuffer() bool
	SetShouldAllocateBuffer(value bool)
	SupportedChannelCounts() foundation.Number
	SetSupportedChannelCounts(value foundation.Number)
	SupportedChannelLayoutTags() foundation.Number
	SetSupportedChannelLayoutTags(value foundation.Number)
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



// The index of this bus in its containing array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/index
func (a_ AudioUnitBus) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}


// The bus type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/bustype
func (a_ AudioUnitBus) BusType() AUAudioUnitBusType {
	rv := objc.Send[AUAudioUnitBusType](a_.ID, objc.Sel("busType"))
	return rv
}


// The bus type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/bustype
func (a_ AudioUnitBus) SetBusType(value AUAudioUnitBusType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBusType:"), value)
}


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/contextpresentationlatency
func (a_ AudioUnitBus) ContextPresentationLatency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("contextPresentationLatency"))
	return rv
}


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/contextpresentationlatency
func (a_ AudioUnitBus) SetContextPresentationLatency(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContextPresentationLatency:"), value)
}


// The audio format and channel layout of audio being transferred on the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/format
func (a_ AudioUnitBus) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}


// The audio format and channel layout of audio being transferred on the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/format
func (a_ AudioUnitBus) SetFormat(value avfaudio.AudioFormat) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormat:"), value)
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


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/maximumchannelcount
func (a_ AudioUnitBus) MaximumChannelCount() AudioChannelCount {
	rv := objc.Send[AudioChannelCount](a_.ID, objc.Sel("maximumChannelCount"))
	return rv
}


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/maximumchannelcount
func (a_ AudioUnitBus) SetMaximumChannelCount(value AudioChannelCount) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumChannelCount:"), value)
}


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/name
func (a_ AudioUnitBus) Name() string {
	rv := objc.Send[string](a_.ID, objc.Sel("name"))
	return rv
}


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/name
func (a_ AudioUnitBus) SetName(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), objc.String(value))
}


// The audio unit that owns the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/owneraudiounit
func (a_ AudioUnitBus) OwnerAudioUnit() IAUAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("ownerAudioUnit"))
	return rv
}


// The audio unit that owns the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/owneraudiounit
func (a_ AudioUnitBus) SetOwnerAudioUnit(value IAUAudioUnit) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOwnerAudioUnit:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/shouldallocatebuffer
func (a_ AudioUnitBus) ShouldAllocateBuffer() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldAllocateBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/shouldallocatebuffer
func (a_ AudioUnitBus) SetShouldAllocateBuffer(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldAllocateBuffer:"), value)
}


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/supportedchannelcounts
func (a_ AudioUnitBus) SupportedChannelCounts() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("supportedChannelCounts"))
	return rv
}


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/supportedchannelcounts
func (a_ AudioUnitBus) SetSupportedChannelCounts(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportedChannelCounts:"), value)
}


// An array of audio channel layout tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/supportedchannellayouttags
func (a_ AudioUnitBus) SupportedChannelLayoutTags() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("supportedChannelLayoutTags"))
	return rv
}


// An array of audio channel layout tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/supportedchannellayouttags
func (a_ AudioUnitBus) SetSupportedChannelLayoutTags(value foundation.Number) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportedChannelLayoutTags:"), value)
}



