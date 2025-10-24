// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureTimecodeGenerator */


/* debug [class_header]: Header for AVCaptureTimecodeGenerator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureTimecodeGenerator */
// An interface definition for the [CaptureTimecodeGenerator] class.
type ICaptureTimecodeGenerator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureTimecodeGenerator */
	// properties:
	AvailableSources() []CaptureTimecodeSource
	CurrentSource() IAVCaptureTimecodeSource
	Delegate() unsafe.Pointer
	DelegateCallbackQueue() objectivec.IObject
	SynchronizationTimeout() float64
	SetSynchronizationTimeout(value float64)
	TimecodeAlignmentOffset() float64
	SetTimecodeAlignmentOffset(value float64)
	TimecodeFrameDuration() objc.IObject /* cross-framework: Time */
	SetTimecodeFrameDuration(value objc.IObject /* cross-framework: Time */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureTimecodeGenerator */
	// methods:
	GenerateInitialTimecode() objc.IObject /* cross-framework: AVCaptureTimecode */
	SetDelegateQueue(delegate unsafe.Pointer, callbackQueue objectivec.IObject)
	StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureTimecodeGenerator */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureTimecodeGenerator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureTimecodeGenerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureTimecodeGenerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureTimecodeGenerator */

// A frame counter timecode source that operates independently of any internal or external synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/frameCountSource
func (cc _CaptureTimecodeGeneratorClass) FrameCountSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("frameCountSource"))
	return rv
}/* debug [class_properties_class/property]: frameCountSource */

// A predefined timecode source synchronized to the real-time system clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/realTimeClockSource
func (cc _CaptureTimecodeGeneratorClass) RealTimeClockSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](objc.ID(cc.class), objc.Sel("realTimeClockSource"))
	return rv
}/* debug [class_properties_class/property]: realTimeClockSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureTimecodeGenerator */

// Generates an initial timecode intended to be the first in a sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/generateInitialTimecode()
func (c_ CaptureTimecodeGenerator) GenerateInitialTimecode() objc.IObject /* cross-framework: AVCaptureTimecode */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("generateInitialTimecode"))
	return rv
}/* debug [instance_methods/method]: GenerateInitialTimecode */


// Assigns a delegate to receive real-time timecode updates and specifies a queue for callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/setDelegate(_:queue:)
func (c_ CaptureTimecodeGenerator) SetDelegateQueue(delegate unsafe.Pointer, callbackQueue objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, callbackQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */


// Synchronizes the generator with the specified timecode source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/startSynchronization(source:)
func (c_ CaptureTimecodeGenerator) StartSynchronizationWithTimecodeSource(source IAVCaptureTimecodeSource) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startSynchronizationWithTimecodeSource:"), source)
}/* debug [instance_methods/method]: StartSynchronizationWithTimecodeSource */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureTimecodeGenerator */

// An array of available timecode synchronization sources that can be used by the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/availableSources
func (c_ CaptureTimecodeGenerator) AvailableSources() []CaptureTimecodeSource {
	rv := objc.Send[[]CaptureTimecodeSource](c_.ID, objc.Sel("availableSources"))
	return rv
}/* debug [instance_properties/getter]: availableSources */


// The active timecode source used by to maintain clock synchronization for accurate timecode generation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/currentSource
func (c_ CaptureTimecodeGenerator) CurrentSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("currentSource"))
	return rv
}/* debug [instance_properties/getter]: currentSource */


// The delegate that receives timecode updates from the timecode generator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/delegate
func (c_ CaptureTimecodeGenerator) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue on which delegate callbacks are invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/delegateCallbackQueue
func (c_ CaptureTimecodeGenerator) DelegateCallbackQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("delegateCallbackQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateCallbackQueue */


// A frame counter timecode source that operates independently of any internal or external synchronization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/frameCountSource
func (c_ CaptureTimecodeGenerator) FrameCountSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("frameCountSource"))
	return rv
}/* debug [instance_properties/getter]: frameCountSource */


// A predefined timecode source synchronized to the real-time system clock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/realTimeClockSource
func (c_ CaptureTimecodeGenerator) RealTimeClockSource() IAVCaptureTimecodeSource {
	rv := objc.Send[CaptureTimecodeSource](c_.ID, objc.Sel("realTimeClockSource"))
	return rv
}/* debug [instance_properties/getter]: realTimeClockSource */


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/synchronizationTimeout
func (c_ CaptureTimecodeGenerator) SynchronizationTimeout() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("synchronizationTimeout"))
	return rv
}/* debug [instance_properties/getter]: synchronizationTimeout */


// The maximum time interval allowed for source synchronization attempts before timing out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/synchronizationTimeout
func (c_ CaptureTimecodeGenerator) SetSynchronizationTimeout(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSynchronizationTimeout:"), value)
}/* debug [instance_properties/setter]: synchronizationTimeout */


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeAlignmentOffset
func (c_ CaptureTimecodeGenerator) TimecodeAlignmentOffset() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("timecodeAlignmentOffset"))
	return rv
}/* debug [instance_properties/getter]: timecodeAlignmentOffset */


// The time offset, in seconds, applied to the generated timecode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeAlignmentOffset
func (c_ CaptureTimecodeGenerator) SetTimecodeAlignmentOffset(value float64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeAlignmentOffset:"), value)
}/* debug [instance_properties/setter]: timecodeAlignmentOffset */


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeFrameDuration
func (c_ CaptureTimecodeGenerator) TimecodeFrameDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("timecodeFrameDuration"))
	return rv
}/* debug [instance_properties/getter]: timecodeFrameDuration */


// The frame duration that the generator will use to generate timecodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/timecodeFrameDuration
func (c_ CaptureTimecodeGenerator) SetTimecodeFrameDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimecodeFrameDuration:"), value)
}/* debug [instance_properties/setter]: timecodeFrameDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureTimecodeGenerator */



