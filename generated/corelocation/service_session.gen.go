// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ServiceSession] class.
var serviceSessionClass = _ServiceSessionClass{objc.GetClass("CLServiceSession")}

type _ServiceSessionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd

type ServiceSession struct {
	objectivec.Object
}

// ServiceSessionFrom constructs a [ServiceSession] from an unsafe.Pointer.
func ServiceSessionFrom(ptr unsafe.Pointer) ServiceSession {
	return ServiceSession{objectivec.Object{objc.ID(ptr)}}
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:
func (sc _ServiceSessionClass) SessionRequiringAuthorization(authorizationRequirement unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:"), authorizationRequirement)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:fullAccuracyPurposeKey:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationFullAccuracyPurposeKey(authorizationRequirement unsafe.Pointer, purposeKey string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:fullAccuracyPurposeKey:"), authorizationRequirement, purposeKey)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:fullAccuracyPurposeKey:queue:handler:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationFullAccuracyPurposeKeyQueueHandler(authorizationRequirement unsafe.Pointer, purposeKey string, queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:fullAccuracyPurposeKey:queue:handler:"), authorizationRequirement, purposeKey, queue, handler)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:queue:handler:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationQueueHandler(authorizationRequirement unsafe.Pointer, queue unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:queue:handler:"), authorizationRequirement, queue, handler)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/invalidate
func (s_ ServiceSession) Invalidate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidate"))
}


