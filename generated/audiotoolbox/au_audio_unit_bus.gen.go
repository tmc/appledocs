// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUAudioUnitBus */


/* debug [class_header]: Header for AUAudioUnitBus */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitBus */
// An interface definition for the [AudioUnitBus] class.
type IAudioUnitBus interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnitBus */
	// properties:
	BusType() AudioUnitBusType
	ContextPresentationLatency() float64
	SetContextPresentationLatency(value float64)
	Format() avfaudio.AudioFormat
	Index() uint
	Enabled() bool
	SetEnabled(value bool)
	MaximumChannelCount() AudioChannelCount /* typedef */
	SetMaximumChannelCount(value AudioChannelCount /* typedef */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	OwnerAudioUnit() IAUAudioUnit
	ShouldAllocateBuffer() bool
	SetShouldAllocateBuffer(value bool)
	SupportedChannelCounts() []foundation.Number
	SetSupportedChannelCounts(value []foundation.Number)
	SupportedChannelLayoutTags() []foundation.Number
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitBus */
	// methods:
	SetFormatError(format avfaudio.AudioFormat, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitBus */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitBusClass) Alloc() AudioUnitBus {
	rv := objc.Send[AudioUnitBus](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitBus */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitBus */

// Initializes a bus object with a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/init(format:)
func NewAudioUnitBusWithFormatError(format avfaudio.AudioFormat, outError objectivec.IObject) AudioUnitBus {
	instance := getAudioUnitBusClass().Alloc()
	rv := objc.Send[AudioUnitBus](instance.ID, objc.Sel("initWithFormat:error:"), format, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAudioUnitBusWithFormatError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitBus */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitBus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitBus */

// Sets the bus’s audio format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/setFormat(_:)
func (a_ AudioUnitBus) SetFormatError(format avfaudio.AudioFormat, outError objectivec.IObject) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("setFormat:error:"), format, outError)
	return rv
}/* debug [instance_methods/method]: SetFormatError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitBus */

// The bus type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/busType
func (a_ AudioUnitBus) BusType() AudioUnitBusType {
	rv := objc.Send[AudioUnitBusType](a_.ID, objc.Sel("busType"))
	return rv
}/* debug [instance_properties/getter]: busType */


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/contextPresentationLatency
func (a_ AudioUnitBus) ContextPresentationLatency() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("contextPresentationLatency"))
	return rv
}/* debug [instance_properties/getter]: contextPresentationLatency */


// Information about latency in the audio unit’s processing context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/contextPresentationLatency
func (a_ AudioUnitBus) SetContextPresentationLatency(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContextPresentationLatency:"), value)
}/* debug [instance_properties/setter]: contextPresentationLatency */


// The audio format and channel layout of audio being transferred on the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/format
func (a_ AudioUnitBus) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](a_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// The index of this bus in its containing array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/index
func (a_ AudioUnitBus) Index() uint {
	rv := objc.Send[uint](a_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/isEnabled
func (a_ AudioUnitBus) Enabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/isEnabled
func (a_ AudioUnitBus) SetEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/maximumChannelCount
func (a_ AudioUnitBus) MaximumChannelCount() AudioChannelCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("maximumChannelCount"))
	return rv
}/* debug [instance_properties/getter]: maximumChannelCount */


// The maximum number of channels supported for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/maximumChannelCount
func (a_ AudioUnitBus) SetMaximumChannelCount(value AudioChannelCount /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumChannelCount:"), value)
}/* debug [instance_properties/setter]: maximumChannelCount */


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/name
func (a_ AudioUnitBus) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A name for the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/name
func (a_ AudioUnitBus) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// The audio unit that owns the bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/ownerAudioUnit
func (a_ AudioUnitBus) OwnerAudioUnit() IAUAudioUnit {
	rv := objc.Send[AudioUnit](a_.ID, objc.Sel("ownerAudioUnit"))
	return rv
}/* debug [instance_properties/getter]: ownerAudioUnit */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/shouldAllocateBuffer
func (a_ AudioUnitBus) ShouldAllocateBuffer() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldAllocateBuffer"))
	return rv
}/* debug [instance_properties/getter]: shouldAllocateBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/shouldAllocateBuffer
func (a_ AudioUnitBus) SetShouldAllocateBuffer(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldAllocateBuffer:"), value)
}/* debug [instance_properties/setter]: shouldAllocateBuffer */


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelCounts
func (a_ AudioUnitBus) SupportedChannelCounts() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("supportedChannelCounts"))
	return rv
}/* debug [instance_properties/getter]: supportedChannelCounts */


// An array of numbers indicating the supported number of channels for this bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelCounts
func (a_ AudioUnitBus) SetSupportedChannelCounts(value []foundation.Number) {
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
}/* debug [instance_properties/setter]: supportedChannelCounts */


// An array of audio channel layout tags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnitBus/supportedChannelLayoutTags
func (a_ AudioUnitBus) SupportedChannelLayoutTags() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("supportedChannelLayoutTags"))
	return rv
}/* debug [instance_properties/getter]: supportedChannelLayoutTags */


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/isenabled
func (a_ AudioUnitBus) IsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// Determines whether the bus is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auaudiounitbus/isenabled
func (a_ AudioUnitBus) SetIsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUAudioUnitBus */


