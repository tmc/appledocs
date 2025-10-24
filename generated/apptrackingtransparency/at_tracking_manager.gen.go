// Code generated from Apple documentation for AppTrackingTransparency. DO NOT EDIT.

package apptrackingtransparency

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ATTrackingManager */


/* debug [class_header]: Header for ATTrackingManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ATTrackingManager */
// An interface definition for the [ATTrackingManager] class.
type IATTrackingManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ATTrackingManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ATTrackingManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ATTrackingManager */
// Alloc allocates a new instance without initialization.
func (ac _ATTrackingManagerClass) Alloc() ATTrackingManager {
	rv := objc.Send[ATTrackingManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ATTrackingManager */
// A class that provides a tracking authorization request and the tracking authorization status of the app.


// A class that provides a tracking authorization request and the tracking authorization status of the app.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ATTrackingManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ATTrackingManager */

// The request for user authorization to access app-related data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/requestTrackingAuthorization(completionHandler:)
func (ac _ATTrackingManagerClass) RequestTrackingAuthorizationWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("requestTrackingAuthorizationWithCompletionHandler:"), completion)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestTrackingAuthorizationWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ATTrackingManager */

// The authorization status that is current for the calling application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/trackingAuthorizationStatus
func (ac _ATTrackingManagerClass) TrackingAuthorizationStatus() ATTrackingManagerAuthorizationStatus {
	rv := objc.Send[ATTrackingManagerAuthorizationStatus](objc.ID(ac.class), objc.Sel("trackingAuthorizationStatus"))
	return rv
}/* debug [class_properties_class/property]: trackingAuthorizationStatus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ATTrackingManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ATTrackingManager */

// The authorization status that is current for the calling application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/trackingAuthorizationStatus
func (a_ ATTrackingManager) TrackingAuthorizationStatus() ATTrackingManagerAuthorizationStatus {
	rv := objc.Send[ATTrackingManagerAuthorizationStatus](a_.ID, objc.Sel("trackingAuthorizationStatus"))
	return rv
}/* debug [instance_properties/getter]: trackingAuthorizationStatus */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ATTrackingManager */






