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
	

	// properties:
	AvailableSources() []CaptureTimecodeSource
	CurrentSource() IAVCaptureTimecodeSource
	DelegateCallbackQueue() objectivec.IObject
	SynchronizationTimeout() float64
	SetSynchronizationTimeout(value float64)
	TimecodeAlignmentOffset() float64
	SetTimecodeAlignmentOffset(value float64)
	TimecodeFrameDuration() objectivec.IObject
	SetTimecodeFrameDuration(value objectivec.IObject)


	

	// methods:
	GenerateInitialTimecode() AVCaptureTimecode
	SetDelegateQueue(delegate unsafe.Pointer, callbackQueue objectivec.IObject)
	StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureTimecodeGeneratorClass) Alloc() CaptureTimecodeGenerator {
	rv := objc.Send[CaptureTimecodeGenerator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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















// A frame counter timecode source that operates independently of any internal or external synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/frameCountSource
func (cc _CaptureTimecodeGeneratorClass) FrameCountSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("frameCountSource"))
	return rv
}

// A predefined timecode source synchronized to the real-time system clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/realTimeClockSource
func (cc _CaptureTimecodeGeneratorClass) RealTimeClockSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("realTimeClockSource"))
	return rv
}






// Generates an initial timecode intended to be the first in a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/generateInitialTimecode()
func (c_ CaptureTimecodeGenerator) GenerateInitialTimecode() AVCaptureTimecode {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("generateInitialTimecode"))
	return rv
}


// Assigns a delegate to receive real-time timecode updates and specifies a queue for callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/setDelegate(_:queue:)
func (c_ CaptureTimecodeGenerator) SetDelegateQueue(delegate unsafe.Pointer, callbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, callbackQueue)
}


// Synchronizes the generator with the specified timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/startSynchronization(source:)
func (c_ CaptureTimecodeGenerator) StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startSynchronizationWithTimecodeSource:"), source)
}







// An array of available timecode synchronization sources that can be used by the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/availableSources
func (c_ CaptureTimecodeGenerator) AvailableSources() []CaptureTimecodeSource {
	rv := objc.Send[[]CaptureTimecodeSource](c_.ID, objc.Sel("availableSources"))
	return rv
}


// The active timecode source used by to maintain clock synchronization for accurate timecode generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/currentSource
func (c_ CaptureTimecodeGenerator) CurrentSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("currentSource"))
	return rv
}


// The dispatch queue on which delegate callbacks are invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/delegateCallbackQueue
func (c_ CaptureTimecodeGenerator) DelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}


// A frame counter timecode source that operates independently of any internal or external synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/frameCountSource
func (c_ CaptureTimecodeGenerator) FrameCountSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("frameCountSource"))
	return rv
}


// A predefined timecode source synchronized to the real-time system clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/realTimeClockSource
func (c_ CaptureTimecodeGenerator) RealTimeClockSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("realTimeClockSource"))
	return rv
}


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/synchronizationTimeout
func (c_ CaptureTimecodeGenerator) SynchronizationTimeout() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("synchronizationTimeout"))
	return rv
}


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/synchronizationTimeout
func (c_ CaptureTimecodeGenerator) SetSynchronizationTimeout(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSynchronizationTimeout:"), value)
}


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeAlignmentOffset
func (c_ CaptureTimecodeGenerator) TimecodeAlignmentOffset() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("timecodeAlignmentOffset"))
	return rv
}


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeAlignmentOffset
func (c_ CaptureTimecodeGenerator) SetTimecodeAlignmentOffset(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeAlignmentOffset:"), value)
}


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeFrameDuration
func (c_ CaptureTimecodeGenerator) TimecodeFrameDuration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("timecodeFrameDuration"))
	return rv
}


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeFrameDuration
func (c_ CaptureTimecodeGenerator) SetTimecodeFrameDuration(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeFrameDuration:"), value)
}








