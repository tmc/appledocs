// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CaptureTimecodeGenerator] class.
var (
	CaptureTimecodeGeneratorClass     _CaptureTimecodeGeneratorClass
	CaptureTimecodeGeneratorClassOnce sync.Once
)

func getCaptureTimecodeGeneratorClass() _CaptureTimecodeGeneratorClass {
	CaptureTimecodeGeneratorClassOnce.Do(func() {
		CaptureTimecodeGeneratorClass = _CaptureTimecodeGeneratorClass{objc.GetClass("AVCaptureTimecodeGenerator")}
	})
	return CaptureTimecodeGeneratorClass
}

type _CaptureTimecodeGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [CaptureTimecodeGenerator] class.
type ICaptureTimecodeGenerator interface {
	objectivec.IObject
	StartSynchronizationWithTimecodeSource(source unsafe.Pointer)
	AvailableSources() unsafe.Pointer
	SetAvailableSources(value unsafe.Pointer)
	CurrentSource() unsafe.Pointer
	SetCurrentSource(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DelegateCallbackQueue() unsafe.Pointer
	SetDelegateCallbackQueue(value unsafe.Pointer)
	SynchronizationTimeout() unsafe.Pointer
	SetSynchronizationTimeout(value unsafe.Pointer)
	TimecodeAlignmentOffset() unsafe.Pointer
	SetTimecodeAlignmentOffset(value unsafe.Pointer)
	TimecodeFrameDuration() unsafe.Pointer
	SetTimecodeFrameDuration(value unsafe.Pointer)
}

// Generates and synchronizes timecode data from various sources for precise video and audio synchronization.
//
// The class supports multiple timecode sources, including frame counting, system clock synchronization, and MIDI timecode input (MTC). Suitable for playback, recording, or other time-sensitive operations where precise timecode metadata is required. Use the method to set up the desired timecode source.


// Generates and synchronizes timecode data from various sources for precise video and audio synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator

type CaptureTimecodeGenerator struct {
	objectivec.Object
}

// CaptureTimecodeGeneratorFrom constructs a [CaptureTimecodeGenerator] from an unsafe.Pointer.
//
// Generates and synchronizes timecode data from various sources for precise video and audio synchronization.
func CaptureTimecodeGeneratorFrom(ptr unsafe.Pointer) CaptureTimecodeGenerator {
	return CaptureTimecodeGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureTimecodeGeneratorClass) Alloc() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureTimecodeGeneratorClass) New() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureTimecodeGenerator) Init() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureTimecodeGenerator) Autorelease() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureTimecodeGenerator creates a new CaptureTimecodeGenerator instance.
func NewCaptureTimecodeGenerator() CaptureTimecodeGenerator {
	return getCaptureTimecodeGeneratorClass().New()
}




// Synchronizes the generator with the specified timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/startSynchronization(source:)

func (c_ CaptureTimecodeGenerator) StartSynchronizationWithTimecodeSource(source unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startSynchronizationWithTimecodeSource:"), source)
}


// An array of available timecode synchronization sources that can be used by the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/availablesources

func (c_ CaptureTimecodeGenerator) AvailableSources() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("availableSources"))
	return rv
}


// An array of available timecode synchronization sources that can be used by the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/availablesources

func (c_ CaptureTimecodeGenerator) SetAvailableSources(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableSources:"), value)
}


// The active timecode source used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/currentsource

func (c_ CaptureTimecodeGenerator) CurrentSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("currentSource"))
	return rv
}


// The active timecode source used by
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/currentsource

func (c_ CaptureTimecodeGenerator) SetCurrentSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCurrentSource:"), value)
}


// The delegate that receives timecode updates from the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/delegate

func (c_ CaptureTimecodeGenerator) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate that receives timecode updates from the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/delegate

func (c_ CaptureTimecodeGenerator) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}


// The dispatch queue on which delegate callbacks are invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/delegatecallbackqueue

func (c_ CaptureTimecodeGenerator) DelegateCallbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}


// The dispatch queue on which delegate callbacks are invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/delegatecallbackqueue

func (c_ CaptureTimecodeGenerator) SetDelegateCallbackQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegateCallbackQueue:"), value)
}


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/synchronizationtimeout

func (c_ CaptureTimecodeGenerator) SynchronizationTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("synchronizationTimeout"))
	return rv
}


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/synchronizationtimeout

func (c_ CaptureTimecodeGenerator) SetSynchronizationTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSynchronizationTimeout:"), value)
}


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/timecodealignmentoffset

func (c_ CaptureTimecodeGenerator) TimecodeAlignmentOffset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timecodeAlignmentOffset"))
	return rv
}


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/timecodealignmentoffset

func (c_ CaptureTimecodeGenerator) SetTimecodeAlignmentOffset(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeAlignmentOffset:"), value)
}


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/timecodeframeduration

func (c_ CaptureTimecodeGenerator) TimecodeFrameDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timecodeFrameDuration"))
	return rv
}


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturetimecodegenerator/timecodeframeduration

func (c_ CaptureTimecodeGenerator) SetTimecodeFrameDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeFrameDuration:"), value)
}



