// Code generated from Apple documentation for AppTrackingTransparency. DO NOT EDIT.

package apptrackingtransparency

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ATTrackingManager] class.
var (
	ATTrackingManagerClass     _ATTrackingManagerClass
	ATTrackingManagerClassOnce sync.Once
)

func getATTrackingManagerClass() _ATTrackingManagerClass {
	ATTrackingManagerClassOnce.Do(func() {
		ATTrackingManagerClass = _ATTrackingManagerClass{objc.GetClass("ATTrackingManager")}
	})
	return ATTrackingManagerClass
}

type _ATTrackingManagerClass struct {
	class objc.Class
}

// An interface definition for the [ATTrackingManager] class.
type IATTrackingManager interface {
	objectivec.IObject
}

// A class that provides a tracking authorization request and the tracking authorization status of the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager
type ATTrackingManager struct {
	objectivec.Object
}

// ATTrackingManagerFrom constructs a [ATTrackingManager] from an unsafe.Pointer.
//
// A class that provides a tracking authorization request and the tracking authorization status of the app.
func ATTrackingManagerFrom(ptr unsafe.Pointer) ATTrackingManager {
	return ATTrackingManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ATTrackingManagerClass) Alloc() ATTrackingManager {
	rv := objc.Send[ATTrackingManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ATTrackingManagerClass) New() ATTrackingManager {
	rv := objc.Send[ATTrackingManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ATTrackingManager) Init() ATTrackingManager {
	rv := objc.Send[ATTrackingManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ATTrackingManager) Autorelease() ATTrackingManager {
	rv := objc.Send[ATTrackingManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewATTrackingManager creates a new ATTrackingManager instance.
func NewATTrackingManager() ATTrackingManager {
	return getATTrackingManagerClass().New()
}


// The request for user authorization to access app-related data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/requestTrackingAuthorization(completionHandler:)
func (ac _ATTrackingManagerClass) RequestTrackingAuthorizationWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestTrackingAuthorizationWithCompletionHandler:"), completion)
}

// The authorization status that is current for the calling application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/trackingAuthorizationStatus
func (ac _ATTrackingManagerClass) TrackingAuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("trackingAuthorizationStatus"))
	return rv
}
// The authorization status that is current for the calling application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/trackingAuthorizationStatus
func (a_ ATTrackingManager) TrackingAuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackingAuthorizationStatus"))
	return rv
}


