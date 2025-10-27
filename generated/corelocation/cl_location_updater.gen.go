// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [LocationUpdater] class.
var (
	LocationUpdaterClass     _LocationUpdaterClass
	LocationUpdaterClassOnce sync.Once
)

func getLocationUpdaterClass() _LocationUpdaterClass {
	LocationUpdaterClassOnce.Do(func() {
		LocationUpdaterClass = _LocationUpdaterClass{objc.GetClass("CLLocationUpdater")}
	})
	return LocationUpdaterClass
}

type _LocationUpdaterClass struct {
	class objc.Class
}





// An interface definition for the [LocationUpdater] class.
type ILocationUpdater interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	Invalidate()
	Pause()
	Resume()


}





// Alloc allocates a new instance without initialization.
func (lc _LocationUpdaterClass) Alloc() LocationUpdater {
	rv := objc.Send[LocationUpdater](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LocationUpdaterClass) New() LocationUpdater {
	rv := objc.Send[LocationUpdater](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocationUpdater) Init() LocationUpdater {
	rv := objc.Send[LocationUpdater](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocationUpdater) Autorelease() LocationUpdater {
	rv := objc.Send[LocationUpdater](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocationUpdater creates a new LocationUpdater instance.
func NewLocationUpdater() LocationUpdater {
	return getLocationUpdaterClass().New()
}





// An object that provides device location updates.


// An object that provides device location updates.
//
// [Full Topic]
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










// Creates a location updater with the configuration and queue that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithConfiguration:queue:handler:
func (lc _LocationUpdaterClass) LiveUpdaterWithConfigurationQueueHandler(configuration LiveUpdateConfiguration, queue objectivec.IObject, handler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("liveUpdaterWithConfiguration:queue:handler:"), configuration, queue, handler)
	return rv
}


// Creates a location updater on the queue you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/liveUpdaterWithQueue:handler:
func (lc _LocationUpdaterClass) LiveUpdaterWithQueueHandler(queue objectivec.IObject, handler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("liveUpdaterWithQueue:handler:"), queue, handler)
	return rv
}












// Invalidates the updater.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/invalidate
func (l_ LocationUpdater) Invalidate() {
	objc.Send[objc.ID](l_.ID, objc.Sel("invalidate"))
}


// Pauses the updater.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/pause
func (l_ LocationUpdater) Pause() {
	objc.Send[objc.ID](l_.ID, objc.Sel("pause"))
}


// Resumes the updater.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLLocationUpdater/resume
func (l_ LocationUpdater) Resume() {
	objc.Send[objc.ID](l_.ID, objc.Sel("resume"))
}













