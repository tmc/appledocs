// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ServiceSession] class.
var (
	ServiceSessionClass     _ServiceSessionClass
	ServiceSessionClassOnce sync.Once
)

func getServiceSessionClass() _ServiceSessionClass {
	ServiceSessionClassOnce.Do(func() {
		ServiceSessionClass = _ServiceSessionClass{objc.GetClass("CLServiceSession")}
	})
	return ServiceSessionClass
}

type _ServiceSessionClass struct {
	class objc.Class
}

// An interface definition for the [ServiceSession] class.
type IServiceSession interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd
type ServiceSession struct {
	objectivec.Object
}

// ServiceSessionFrom constructs a [ServiceSession] from an unsafe.Pointer.
func ServiceSessionFrom(ptr unsafe.Pointer) ServiceSession {
	return ServiceSession{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ServiceSessionClass) Alloc() ServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ServiceSessionClass) New() ServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ServiceSession) Init() ServiceSession {
	rv := objc.Send[ServiceSession](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ServiceSession) Autorelease() ServiceSession {
	rv := objc.Send[ServiceSession](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewServiceSession creates a new ServiceSession instance.
func NewServiceSession() ServiceSession {
	return getServiceSessionClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:
func (sc _ServiceSessionClass) SessionRequiringAuthorization(authorizationRequirement ServiceSessionAuthorizationRequirement) IServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:"), authorizationRequirement)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:fullAccuracyPurposeKey:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationFullAccuracyPurposeKey(authorizationRequirement ServiceSessionAuthorizationRequirement, purposeKey objc.IObject /* cross-framework: NSString */) IServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:fullAccuracyPurposeKey:"), authorizationRequirement, purposeKey)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:fullAccuracyPurposeKey:queue:handler:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationFullAccuracyPurposeKeyQueueHandler(authorizationRequirement ServiceSessionAuthorizationRequirement, purposeKey objc.IObject /* cross-framework: NSString */, queue unsafe.Pointer, handler unsafe.Pointer) IServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:fullAccuracyPurposeKey:queue:handler:"), authorizationRequirement, purposeKey, queue, handler)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSession-2ddhd/sessionRequiringAuthorization:queue:handler:
func (sc _ServiceSessionClass) SessionRequiringAuthorizationQueueHandler(authorizationRequirement ServiceSessionAuthorizationRequirement, queue unsafe.Pointer, handler unsafe.Pointer) IServiceSession {
	rv := objc.Send[ServiceSession](objc.ID(sc.class), objc.Sel("sessionRequiringAuthorization:queue:handler:"), authorizationRequirement, queue, handler)
	return rv
}


