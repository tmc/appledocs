// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocationUpdater] class.
var locationUpdaterClass = _LocationUpdaterClass{objc.GetClass("CLLocationUpdater")}

type _LocationUpdaterClass struct {
	class objc.Class
}

// An object that provides device location updates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater

type LocationUpdater struct {
	objectivec.Object
}

// LocationUpdaterFrom constructs a [LocationUpdater] from an unsafe.Pointer.
//
// An object that provides device location updates.
func LocationUpdaterFrom(ptr unsafe.Pointer) LocationUpdater {
	return LocationUpdater{objectivec.Object{objc.ID(ptr)}}
}

// Creates a location updater with the configuration and queue that you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithConfiguration:queue:handler:
func (lc _LocationUpdaterClass) LiveUpdaterWithConfigurationQueueHandler(configuration unsafe.Pointer, queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("liveUpdaterWithConfiguration:queue:handler:"), configuration, queue, handler)
	return rv
}
// Creates a location updater on the queue you specify. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithQueue:handler:
func (lc _LocationUpdaterClass) LiveUpdaterWithQueueHandler(queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("liveUpdaterWithQueue:handler:"), queue, handler)
	return rv
}
// Invalidates the updater. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/invalidate
func (l_ LocationUpdater) Invalidate() {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidate"))
}
// Pauses the updater. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/pause
func (l_ LocationUpdater) Pause() {
	objc.Send[objc.ID](l_.ID, objc.Sel("pause"))
}
// Resumes the updater. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/resume
func (l_ LocationUpdater) Resume() {
	objc.Send[objc.ID](l_.ID, objc.Sel("resume"))
}


