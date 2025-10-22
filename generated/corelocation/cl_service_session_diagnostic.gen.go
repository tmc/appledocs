// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ServiceSessionDiagnostic] class.
var (
	ServiceSessionDiagnosticClass     _ServiceSessionDiagnosticClass
	ServiceSessionDiagnosticClassOnce sync.Once
)

func getServiceSessionDiagnosticClass() _ServiceSessionDiagnosticClass {
	ServiceSessionDiagnosticClassOnce.Do(func() {
		ServiceSessionDiagnosticClass = _ServiceSessionDiagnosticClass{objc.GetClass("CLServiceSessionDiagnostic")}
	})
	return ServiceSessionDiagnosticClass
}

type _ServiceSessionDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [ServiceSessionDiagnostic] class.
type IServiceSessionDiagnostic interface {
	objectivec.IObject
	AlwaysAuthorizationDenied() bool
	AuthorizationDenied() bool
	AuthorizationDeniedGlobally() bool
	AuthorizationRequestInProgress() bool
	AuthorizationRestricted() bool
	FullAccuracyDenied() bool
	InsufficientlyInUse() bool
	ServiceSessionRequired() bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic
type ServiceSessionDiagnostic struct {
	objectivec.Object
}

// ServiceSessionDiagnosticFrom constructs a [ServiceSessionDiagnostic] from an unsafe.Pointer.
func ServiceSessionDiagnosticFrom(ptr unsafe.Pointer) ServiceSessionDiagnostic {
	return ServiceSessionDiagnostic{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ServiceSessionDiagnosticClass) Alloc() ServiceSessionDiagnostic {
	rv := objc.Send[ServiceSessionDiagnostic](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ServiceSessionDiagnosticClass) New() ServiceSessionDiagnostic {
	rv := objc.Send[ServiceSessionDiagnostic](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ServiceSessionDiagnostic) Init() ServiceSessionDiagnostic {
	rv := objc.Send[ServiceSessionDiagnostic](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ServiceSessionDiagnostic) Autorelease() ServiceSessionDiagnostic {
	rv := objc.Send[ServiceSessionDiagnostic](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewServiceSessionDiagnostic creates a new ServiceSessionDiagnostic instance.
func NewServiceSessionDiagnostic() ServiceSessionDiagnostic {
	return getServiceSessionDiagnosticClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/alwaysAuthorizationDenied
func (s_ ServiceSessionDiagnostic) AlwaysAuthorizationDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("alwaysAuthorizationDenied"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationDenied
func (s_ ServiceSessionDiagnostic) AuthorizationDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationDenied"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationDeniedGlobally
func (s_ ServiceSessionDiagnostic) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationRequestInProgress
func (s_ ServiceSessionDiagnostic) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationRestricted
func (s_ ServiceSessionDiagnostic) AuthorizationRestricted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationRestricted"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/fullAccuracyDenied
func (s_ ServiceSessionDiagnostic) FullAccuracyDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("fullAccuracyDenied"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/insufficientlyInUse
func (s_ ServiceSessionDiagnostic) InsufficientlyInUse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/serviceSessionRequired
func (s_ ServiceSessionDiagnostic) ServiceSessionRequired() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}



