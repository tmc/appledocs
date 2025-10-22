// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNDetectionTrack] class.
var (
	CNDetectionTrackClass     _CNDetectionTrackClass
	CNDetectionTrackClassOnce sync.Once
)

func getCNDetectionTrackClass() _CNDetectionTrackClass {
	CNDetectionTrackClassOnce.Do(func() {
		CNDetectionTrackClass = _CNDetectionTrackClass{objc.GetClass("CNDetectionTrack")}
	})
	return CNDetectionTrackClass
}

type _CNDetectionTrackClass struct {
	class objc.Class
}

// An interface definition for the [CNDetectionTrack] class.
type ICNDetectionTrack interface {
	objectivec.IObject
	DetectionAtOrBeforeTime(time unsafe.Pointer) CNDetection
	DetectionNearestTime(time unsafe.Pointer) CNDetection
	DetectionsInTimeRange(timeRange unsafe.Pointer) []CNDetection
	DetectionGroupID() CNDetectionGroupID
	DetectionType() CNDetectionType
	UserCreated() bool
}

// An object representing a series of detections of the same subject over time.


// An object representing a series of detections of the same subject over time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g

type CNDetectionTrack struct {
	objectivec.Object
}

// CNDetectionTrackFrom constructs a [CNDetectionTrack] from an unsafe.Pointer.
//
// An object representing a series of detections of the same subject over time.
func CNDetectionTrackFrom(ptr unsafe.Pointer) CNDetectionTrack {
	return CNDetectionTrack{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNDetectionTrackClass) Alloc() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNDetectionTrackClass) New() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNDetectionTrack) Init() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNDetectionTrack) Autorelease() CNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNDetectionTrack creates a new CNDetectionTrack instance.
func NewCNDetectionTrack() CNDetectionTrack {
	return getCNDetectionTrackClass().New()
}




// Returns the array of detections in the detection track before a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionAtOrBeforeTime:

func (c_ CNDetectionTrack) DetectionAtOrBeforeTime(time unsafe.Pointer) CNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("detectionAtOrBeforeTime:"), time)
	return rv
}



// Returns the array of detections in the detection track nearest a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionNearestTime:

func (c_ CNDetectionTrack) DetectionNearestTime(time unsafe.Pointer) CNDetection {
	rv := objc.Send[CNDetection](c_.ID, objc.Sel("detectionNearestTime:"), time)
	return rv
}



// Returns the array of detections in the detection track within the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionsInTimeRange:

func (c_ CNDetectionTrack) DetectionsInTimeRange(timeRange unsafe.Pointer) []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("detectionsInTimeRange:"), timeRange)
	return rv
}


// The detection group ID of the subject detected by the track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionGroupID

func (c_ CNDetectionTrack) DetectionGroupID() CNDetectionGroupID {
	rv := objc.Send[CNDetectionGroupID](c_.ID, objc.Sel("detectionGroupID"))
	return rv
}


// The type of object that’s detected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/detectionType

func (c_ CNDetectionTrack) DetectionType() CNDetectionType {
	rv := objc.Send[CNDetectionType](c_.ID, objc.Sel("detectionType"))
	return rv
}


// A flag indicating if the client created the detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDetectionTrack-61x7g/userCreated

func (c_ CNDetectionTrack) UserCreated() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("userCreated"))
	return rv
}



