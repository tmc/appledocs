// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNObjectTracker] class.
var (
	CNObjectTrackerClass     _CNObjectTrackerClass
	CNObjectTrackerClassOnce sync.Once
)

func getCNObjectTrackerClass() _CNObjectTrackerClass {
	CNObjectTrackerClassOnce.Do(func() {
		CNObjectTrackerClass = _CNObjectTrackerClass{objc.GetClass("CNObjectTracker")}
	})
	return CNObjectTrackerClass
}

type _CNObjectTrackerClass struct {
	class objc.Class
}

// An interface definition for the [CNObjectTracker] class.
type ICNObjectTracker interface {
	objectivec.IObject
	// properties:
	// methods:
	ContinueTrackingAtSourceImageSourceDisparity(time Time /* not a class type */, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */) ICNBoundsPrediction
	FindObjectAtPointSourceImage(point coregraphics.CGPoint, sourceImage PixelBufferRef /* not a class type */) ICNBoundsPrediction
	FinishDetectionTrack() ICNDetectionTrack
	ResetDetectionTrack()
	StartTrackingAtWithinSourceImageSourceDisparity(time Time /* not a class type */, normalizedBounds coregraphics.CGRect, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */) bool /* primitive/slice/pointer. */
}

// An object that converts a normalized point or rectangle into a detection track that tracks an object over time.


// An object that converts a normalized point or rectangle into a detection track that tracks an object over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq
type CNObjectTracker struct {
	objectivec.Object
}

// CNObjectTrackerFrom constructs a [CNObjectTracker] from an unsafe.Pointer.
//
// An object that converts a normalized point or rectangle into a detection track that tracks an object over time.
func CNObjectTrackerFrom(ptr unsafe.Pointer) CNObjectTracker {
	return CNObjectTracker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNObjectTrackerClass) Alloc() CNObjectTracker {
	rv := objc.Send[CNObjectTracker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNObjectTrackerClass) New() CNObjectTracker {
	rv := objc.Send[CNObjectTracker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNObjectTracker) Init() CNObjectTracker {
	rv := objc.Send[CNObjectTracker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNObjectTracker) Autorelease() CNObjectTracker {
	rv := objc.Send[CNObjectTracker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNObjectTracker creates a new CNObjectTracker instance.
func NewCNObjectTracker() CNObjectTracker {
	return getCNObjectTrackerClass().New()
}



// Creates a new detection track builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/initWithCommandQueue:
func NewCNObjectTrackerWithCommandQueue(commandQueue objectivec.IObject) CNObjectTracker {
	instance := getCNObjectTrackerClass().Alloc()
	rv := objc.Send[CNObjectTracker](instance.ID, objc.Sel("initWithCommandQueue:"), commandQueue)
	rv.Autorelease()
	return rv
}



// Indicates whether the current device supports object detection and tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/isSupported
func (cc _CNObjectTrackerClass) IsSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(cc.class), objc.Sel("isSupported"))
	return rv
}

// Continues to track an object that you’ve started tracking, and adds a new detection to the detection track you’re building.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/continueTrackingAt:sourceImage:sourceDisparity:
func (c_ CNObjectTracker) ContinueTrackingAtSourceImageSourceDisparity(time Time /* not a class type */, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */) ICNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("continueTrackingAt:sourceImage:sourceDisparity:"), time, sourceImage, sourceDisparity)
	return rv
}


// Finds the bounds of an object at the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/findObjectAtPoint:sourceImage:
func (c_ CNObjectTracker) FindObjectAtPointSourceImage(point coregraphics.CGPoint, sourceImage PixelBufferRef /* not a class type */) ICNBoundsPrediction {
	rv := objc.Send[CNBoundsPrediction](c_.ID, objc.Sel("findObjectAtPoint:sourceImage:"), point, sourceImage)
	return rv
}


// Finish constructing the detection track and return it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/finishDetectionTrack
func (c_ CNObjectTracker) FinishDetectionTrack() ICNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("finishDetectionTrack"))
	return rv
}


// Resets the builder to construct a new detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/resetDetectionTrack
func (c_ CNObjectTracker) ResetDetectionTrack() {
	objc.Send[objc.ID](c_.ID, objc.Sel("resetDetectionTrack"))
}


// Starts creating a detection track to track an object within the given bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/startTrackingAt:within:sourceImage:sourceDisparity:
func (c_ CNObjectTracker) StartTrackingAtWithinSourceImageSourceDisparity(time Time /* not a class type */, normalizedBounds coregraphics.CGRect, sourceImage PixelBufferRef /* not a class type */, sourceDisparity PixelBufferRef /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("startTrackingAt:within:sourceImage:sourceDisparity:"), time, normalizedBounds, sourceImage, sourceDisparity)
	return rv
}


// Indicates whether the current device supports object detection and tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNObjectTracker-7aliq/isSupported
func (c_ CNObjectTracker) IsSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSupported"))
	return rv
}


