// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNFixedDetectionTrack] class.
var (
	CNFixedDetectionTrackClass     _CNFixedDetectionTrackClass
	CNFixedDetectionTrackClassOnce sync.Once
)

func getCNFixedDetectionTrackClass() _CNFixedDetectionTrackClass {
	CNFixedDetectionTrackClassOnce.Do(func() {
		CNFixedDetectionTrackClass = _CNFixedDetectionTrackClass{objc.GetClass("CNFixedDetectionTrack")}
	})
	return CNFixedDetectionTrackClass
}

type _CNFixedDetectionTrackClass struct {
	class objc.Class
}

// An interface definition for the [CNFixedDetectionTrack] class.
type ICNFixedDetectionTrack interface {
	ICNDetectionTrack
	FocusDisparity() float32
	OriginalDetection() CNDetection
}

// An object representing the fixed detection track.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2
type CNFixedDetectionTrack struct {
	CNDetectionTrack
}

// CNFixedDetectionTrackFrom constructs a [CNFixedDetectionTrack] from an unsafe.Pointer.
//
// An object representing the fixed detection track.
func CNFixedDetectionTrackFrom(ptr unsafe.Pointer) CNFixedDetectionTrack {
	return CNFixedDetectionTrack{
		CNDetectionTrack: CNDetectionTrackFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNFixedDetectionTrackClass) Alloc() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNFixedDetectionTrackClass) New() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFixedDetectionTrack) Init() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFixedDetectionTrack) Autorelease() CNFixedDetectionTrack {
	rv := objc.Send[CNFixedDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFixedDetectionTrack creates a new CNFixedDetectionTrack instance.
func NewCNFixedDetectionTrack() CNFixedDetectionTrack {
	return getCNFixedDetectionTrackClass().New()
}




// Creates a detection track with fixed focus at the given disparity.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/initWithFocusDisparity:
func NewCNFixedDetectionTrackWithFocusDisparity(focusDisparity float32) CNFixedDetectionTrack {
	instance := getCNFixedDetectionTrackClass().Alloc()
	rv := objc.Send[CNFixedDetectionTrack](instance.ID, objc.Sel("initWithFocusDisparity:"), focusDisparity)
	rv.Autorelease()
	return rv
}



// Creates a detection track with fixed focus at the disparity of an existing detection.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/initWithOriginalDetection:
func NewCNFixedDetectionTrackWithOriginalDetection(originalDetection ICNDetection) CNFixedDetectionTrack {
	instance := getCNFixedDetectionTrackClass().Alloc()
	rv := objc.Send[CNFixedDetectionTrack](instance.ID, objc.Sel("initWithOriginalDetection:"), originalDetection)
	rv.Autorelease()
	return rv
}


// The disparity to use in order to focus on the object.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/focusDisparity
func (c_ CNFixedDetectionTrack) FocusDisparity() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("focusDisparity"))
	return rv
}

// The original detection based on the fixed detection track.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNFixedDetectionTrack-5aei2/originalDetection
func (c_ CNFixedDetectionTrack) OriginalDetection() CNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("originalDetection"))
	return rv
}


