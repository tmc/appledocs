// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNCustomDetectionTrack] class.
var (
	CNCustomDetectionTrackClass     _CNCustomDetectionTrackClass
	CNCustomDetectionTrackClassOnce sync.Once
)

func getCNCustomDetectionTrackClass() _CNCustomDetectionTrackClass {
	CNCustomDetectionTrackClassOnce.Do(func() {
		CNCustomDetectionTrackClass = _CNCustomDetectionTrackClass{objc.GetClass("CNCustomDetectionTrack")}
	})
	return CNCustomDetectionTrackClass
}

type _CNCustomDetectionTrackClass struct {
	class objc.Class
}

// An interface definition for the [CNCustomDetectionTrack] class.
type ICNCustomDetectionTrack interface {
	ICNDetectionTrack
}

// An object representing a discrete detection track composed of individual detections.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc
type CNCustomDetectionTrack struct {
	CNDetectionTrack
}

// CNCustomDetectionTrackFrom constructs a [CNCustomDetectionTrack] from an unsafe.Pointer.
//
// An object representing a discrete detection track composed of individual detections.
func CNCustomDetectionTrackFrom(ptr unsafe.Pointer) CNCustomDetectionTrack {
	return CNCustomDetectionTrack{
		CNDetectionTrack: CNDetectionTrackFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNCustomDetectionTrackClass) Alloc() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNCustomDetectionTrackClass) New() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNCustomDetectionTrack) Init() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNCustomDetectionTrack) Autorelease() CNCustomDetectionTrack {
	rv := objc.Send[CNCustomDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNCustomDetectionTrack creates a new CNCustomDetectionTrack instance.
func NewCNCustomDetectionTrack() CNCustomDetectionTrack {
	return getCNCustomDetectionTrackClass().New()
}




// Initializes a custom detection track with an array of detections, optionally applying smoothing.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc/initWithDetections:smooth:
func NewCNCustomDetectionTrackWithDetectionsSmooth(detections []CNDetection, applySmoothing bool) CNCustomDetectionTrack {
	instance := getCNCustomDetectionTrackClass().Alloc()
	rv := objc.Send[CNCustomDetectionTrack](instance.ID, objc.Sel("initWithDetections:smooth:"), detections, applySmoothing)
	rv.Autorelease()
	return rv
}


// All detected objects in the track.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCustomDetectionTrack-891hc/allDetections
func (c_ CNCustomDetectionTrack) AllDetections() []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("allDetections"))
	return rv
}


