// Code generated from Apple documentation for MediaAccessibility. DO NOT EDIT.

package mediaaccessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MAFlashingLightsProcessor] class.
var (
	MAFlashingLightsProcessorClass     _MAFlashingLightsProcessorClass
	MAFlashingLightsProcessorClassOnce sync.Once
)

func getMAFlashingLightsProcessorClass() _MAFlashingLightsProcessorClass {
	MAFlashingLightsProcessorClassOnce.Do(func() {
		MAFlashingLightsProcessorClass = _MAFlashingLightsProcessorClass{objc.GetClass("MAFlashingLightsProcessor")}
	})
	return MAFlashingLightsProcessorClass
}

type _MAFlashingLightsProcessorClass struct {
	class objc.Class
}

// An interface definition for the [MAFlashingLightsProcessor] class.
type IMAFlashingLightsProcessor interface {
	objectivec.IObject
	CanProcessSurface(surface unsafe.Pointer) bool
	ProcessSurfaceOutSurfaceTimestampOptions(inSurface unsafe.Pointer, outSurface unsafe.Pointer, timestamp unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer
}

// A class that processes a framebuffer object to detect and dim sequences of flashing lights.
//
// A device with the Dim Flashing Lights setting on automatically dims the brightness of flashing effect sequences when it detects them in video content. If your app performs custom video drawing instead of using APIs, you can use the class to detect and mitigate sequences of flashing effects in your video content. The following example shows how you might incorporate into code that uses APIs. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor
type MAFlashingLightsProcessor struct {
	objectivec.Object
}

// MAFlashingLightsProcessorFrom constructs a [MAFlashingLightsProcessor] from an unsafe.Pointer.
//
// A class that processes a framebuffer object to detect and dim sequences of flashing lights.
func MAFlashingLightsProcessorFrom(ptr unsafe.Pointer) MAFlashingLightsProcessor {
	return MAFlashingLightsProcessor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MAFlashingLightsProcessorClass) Alloc() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MAFlashingLightsProcessorClass) New() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MAFlashingLightsProcessor) Init() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MAFlashingLightsProcessor) Autorelease() MAFlashingLightsProcessor {
	rv := objc.Send[MAFlashingLightsProcessor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMAFlashingLightsProcessor creates a new MAFlashingLightsProcessor instance.
func NewMAFlashingLightsProcessor() MAFlashingLightsProcessor {
	return getMAFlashingLightsProcessorClass().New()
}


// Returns a Boolean value that indicates whether the flashing lights processor can process the content in the surface for sequences of flashing lights.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor/canProcessSurface(_:)
func (m_ MAFlashingLightsProcessor) CanProcessSurface(surface unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canProcessSurface:"), surface)
	return rv
}

// Processes a surface by analyzing pixels for sequences of flashing lights and mitigates them by dimming the content.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaAccessibility/MAFlashingLightsProcessor/processSurface:outSurface:timestamp:options:
func (m_ MAFlashingLightsProcessor) ProcessSurfaceOutSurfaceTimestampOptions(inSurface unsafe.Pointer, outSurface unsafe.Pointer, timestamp unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("processSurface:outSurface:timestamp:options:"), inSurface, outSurface, timestamp, options)
	return rv
}



