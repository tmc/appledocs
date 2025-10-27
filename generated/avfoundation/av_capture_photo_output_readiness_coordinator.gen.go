// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CapturePhotoOutputReadinessCoordinator] class.
var (
	CapturePhotoOutputReadinessCoordinatorClass     _CapturePhotoOutputReadinessCoordinatorClass
	CapturePhotoOutputReadinessCoordinatorClassOnce sync.Once
)

func getCapturePhotoOutputReadinessCoordinatorClass() _CapturePhotoOutputReadinessCoordinatorClass {
	CapturePhotoOutputReadinessCoordinatorClassOnce.Do(func() {
		CapturePhotoOutputReadinessCoordinatorClass = _CapturePhotoOutputReadinessCoordinatorClass{objc.GetClass("AVCapturePhotoOutputReadinessCoordinator")}
	})
	return CapturePhotoOutputReadinessCoordinatorClass
}

type _CapturePhotoOutputReadinessCoordinatorClass struct {
	class objc.Class
}





// An interface definition for the [CapturePhotoOutputReadinessCoordinator] class.
type ICapturePhotoOutputReadinessCoordinator interface {
	objectivec.IObject
	

	// properties:
	CaptureReadiness() CapturePhotoOutputCaptureReadiness


	

	// methods:
	StartTrackingCaptureRequestUsingPhotoSettings(settings IAVCapturePhotoSettings)
	StopTrackingCaptureRequestUsingPhotoSettingsUniqueID(settingsUniqueID int64)


}





// Alloc allocates a new instance without initialization.
func (cc _CapturePhotoOutputReadinessCoordinatorClass) Alloc() CapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[CapturePhotoOutputReadinessCoordinator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CapturePhotoOutputReadinessCoordinatorClass) New() CapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[CapturePhotoOutputReadinessCoordinator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CapturePhotoOutputReadinessCoordinator) Init() CapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[CapturePhotoOutputReadinessCoordinator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CapturePhotoOutputReadinessCoordinator) Autorelease() CapturePhotoOutputReadinessCoordinator {
	rv := objc.Send[CapturePhotoOutputReadinessCoordinator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCapturePhotoOutputReadinessCoordinator creates a new CapturePhotoOutputReadinessCoordinator instance.
func NewCapturePhotoOutputReadinessCoordinator() CapturePhotoOutputReadinessCoordinator {
	return getCapturePhotoOutputReadinessCoordinatorClass().New()
}





// An object that monitors changes to a photo output’s capture readiness.
//
// Use this object to coordinate user interface updates on the main queue with a that runs on a background queue. Adopt the protocol in your app and set its implementation as the coordinator’s delegate object to receive callbacks as the associated photo output’s state changes. You can track additional capture requests with this object by calling its method. You can use it to synchronously update shutter button availability and appearance and on the main thread while calling the photo output’s method asynchronously on a background queue.


// An object that monitors changes to a photo output’s capture readiness.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator
type CapturePhotoOutputReadinessCoordinator struct {
	objectivec.Object
}

// CapturePhotoOutputReadinessCoordinatorFrom constructs a [CapturePhotoOutputReadinessCoordinator] from an unsafe.Pointer.
//
// An object that monitors changes to a photo output’s capture readiness.
func CapturePhotoOutputReadinessCoordinatorFrom(ptr unsafe.Pointer) CapturePhotoOutputReadinessCoordinator {
	return CapturePhotoOutputReadinessCoordinator{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that helps coordinate user interface changes with a photo output that runs on a background queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/init(photoOutput:)
func NewCapturePhotoOutputReadinessCoordinatorWithPhotoOutput(photoOutput IAVCapturePhotoOutput) CapturePhotoOutputReadinessCoordinator {
	instance := getCapturePhotoOutputReadinessCoordinatorClass().Alloc()
	rv := objc.Send[CapturePhotoOutputReadinessCoordinator](instance.ID, objc.Sel("initWithPhotoOutput:"), photoOutput)
	rv.Autorelease()
	return rv
}

















// Tracks a capture request that uses the specified photo settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/startTrackingCaptureRequest(using:)
func (c_ CapturePhotoOutputReadinessCoordinator) StartTrackingCaptureRequestUsingPhotoSettings(settings IAVCapturePhotoSettings) {
	objc.Send[objc.ID](c_.ID, objc.Sel("startTrackingCaptureRequestUsingPhotoSettings:"), settings)
}


// Stop tracking the capture request represented by the specified photo setting’s unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/stopTrackingCaptureRequest(using:)
func (c_ CapturePhotoOutputReadinessCoordinator) StopTrackingCaptureRequestUsingPhotoSettingsUniqueID(settingsUniqueID int64) {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopTrackingCaptureRequestUsingPhotoSettingsUniqueID:"), settingsUniqueID)
}







// A value that indicates whether the coordinator’s photo output is ready to respond to new capture requests in a timely manner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutputReadinessCoordinator/captureReadiness
func (c_ CapturePhotoOutputReadinessCoordinator) CaptureReadiness() CapturePhotoOutputCaptureReadiness {
	rv := objc.Send[CapturePhotoOutputCaptureReadiness](c_.ID, objc.Sel("captureReadiness"))
	return rv
}







