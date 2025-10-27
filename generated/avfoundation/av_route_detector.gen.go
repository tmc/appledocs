// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RouteDetector] class.
var (
	RouteDetectorClass     _RouteDetectorClass
	RouteDetectorClassOnce sync.Once
)

func getRouteDetectorClass() _RouteDetectorClass {
	RouteDetectorClassOnce.Do(func() {
		RouteDetectorClass = _RouteDetectorClass{objc.GetClass("AVRouteDetector")}
	})
	return RouteDetectorClass
}

type _RouteDetectorClass struct {
	class objc.Class
}





// An interface definition for the [RouteDetector] class.
type IRouteDetector interface {
	objectivec.IObject
	

	// properties:
	RouteDetectionEnabled() bool
	SetRouteDetectionEnabled(value bool)
	MultipleRoutesDetected() bool
	IsRouteDetectionEnabled() bool
	SetIsRouteDetectionEnabled(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RouteDetectorClass) Alloc() RouteDetector {
	rv := objc.Send[RouteDetector](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RouteDetectorClass) New() RouteDetector {
	rv := objc.Send[RouteDetector](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RouteDetector) Init() RouteDetector {
	rv := objc.Send[RouteDetector](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RouteDetector) Autorelease() RouteDetector {
	rv := objc.Send[RouteDetector](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRouteDetector creates a new RouteDetector instance.
func NewRouteDetector() RouteDetector {
	return getRouteDetectorClass().New()
}





// An object that detects available media playback routes.
//
// If you enable route detection, the object reports whether it detects multiple playback routes. If it does, use to present the UI for the user to select an appropriate route.


// An object that detects available media playback routes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector
type RouteDetector struct {
	objectivec.Object
}

// RouteDetectorFrom constructs a [RouteDetector] from an unsafe.Pointer.
//
// An object that detects available media playback routes.
func RouteDetectorFrom(ptr unsafe.Pointer) RouteDetector {
	return RouteDetector{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/isRouteDetectionEnabled
func (r_ RouteDetector) RouteDetectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("routeDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/isRouteDetectionEnabled
func (r_ RouteDetector) SetRouteDetectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRouteDetectionEnabled:"), value)
}


// A Boolean value that indicates whether the object detects more than one playback route.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRouteDetector/multipleRoutesDetected
func (r_ RouteDetector) MultipleRoutesDetected() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("multipleRoutesDetected"))
	return rv
}


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/isroutedetectionenabled
func (r_ RouteDetector) IsRouteDetectionEnabled() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRouteDetectionEnabled"))
	return rv
}


// A Boolean value that indicates whether route detection is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avroutedetector/isroutedetectionenabled
func (r_ RouteDetector) SetIsRouteDetectionEnabled(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRouteDetectionEnabled:"), value)
}







