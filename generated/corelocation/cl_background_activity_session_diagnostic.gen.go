// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BackgroundActivitySessionDiagnostic] class.
var (
	BackgroundActivitySessionDiagnosticClass     _BackgroundActivitySessionDiagnosticClass
	BackgroundActivitySessionDiagnosticClassOnce sync.Once
)

func getBackgroundActivitySessionDiagnosticClass() _BackgroundActivitySessionDiagnosticClass {
	BackgroundActivitySessionDiagnosticClassOnce.Do(func() {
		BackgroundActivitySessionDiagnosticClass = _BackgroundActivitySessionDiagnosticClass{objc.GetClass("CLBackgroundActivitySessionDiagnostic")}
	})
	return BackgroundActivitySessionDiagnosticClass
}

type _BackgroundActivitySessionDiagnosticClass struct {
	class objc.Class
}

// An interface definition for the [BackgroundActivitySessionDiagnostic] class.
type IBackgroundActivitySessionDiagnostic interface {
	objectivec.IObject
	// properties:
	AuthorizationDenied() bool /* primitive/slice/pointer. */
	AuthorizationDeniedGlobally() bool /* primitive/slice/pointer. */
	AuthorizationRequestInProgress() bool /* primitive/slice/pointer. */
	AuthorizationRestricted() bool /* primitive/slice/pointer. */
	InsufficientlyInUse() bool /* primitive/slice/pointer. */
	ServiceSessionRequired() bool /* primitive/slice/pointer. */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic
type BackgroundActivitySessionDiagnostic struct {
	objectivec.Object
}

// BackgroundActivitySessionDiagnosticFrom constructs a [BackgroundActivitySessionDiagnostic] from an unsafe.Pointer.
func BackgroundActivitySessionDiagnosticFrom(ptr unsafe.Pointer) BackgroundActivitySessionDiagnostic {
	return BackgroundActivitySessionDiagnostic{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BackgroundActivitySessionDiagnosticClass) Alloc() BackgroundActivitySessionDiagnostic {
	rv := objc.Send[BackgroundActivitySessionDiagnostic](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BackgroundActivitySessionDiagnosticClass) New() BackgroundActivitySessionDiagnostic {
	rv := objc.Send[BackgroundActivitySessionDiagnostic](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BackgroundActivitySessionDiagnostic) Init() BackgroundActivitySessionDiagnostic {
	rv := objc.Send[BackgroundActivitySessionDiagnostic](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BackgroundActivitySessionDiagnostic) Autorelease() BackgroundActivitySessionDiagnostic {
	rv := objc.Send[BackgroundActivitySessionDiagnostic](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBackgroundActivitySessionDiagnostic creates a new BackgroundActivitySessionDiagnostic instance.
func NewBackgroundActivitySessionDiagnostic() BackgroundActivitySessionDiagnostic {
	return getBackgroundActivitySessionDiagnosticClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationDenied
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationDenied() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationDenied"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationDeniedGlobally
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationDeniedGlobally() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationRequestInProgress
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationRequestInProgress() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationRestricted
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationRestricted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationRestricted"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/insufficientlyInUse
func (b_ BackgroundActivitySessionDiagnostic) InsufficientlyInUse() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/serviceSessionRequired
func (b_ BackgroundActivitySessionDiagnostic) ServiceSessionRequired() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}



