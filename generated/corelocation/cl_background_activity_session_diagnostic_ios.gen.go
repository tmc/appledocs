//go:build darwin && ios

// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for BackgroundActivitySessionDiagnostic


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationDenied
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationDenied() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationDenied"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationDeniedGlobally
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationRequestInProgress
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/authorizationRestricted
func (b_ BackgroundActivitySessionDiagnostic) AuthorizationRestricted() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("authorizationRestricted"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/insufficientlyInUse
func (b_ BackgroundActivitySessionDiagnostic) InsufficientlyInUse() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLBackgroundActivitySessionDiagnostic/serviceSessionRequired
func (b_ BackgroundActivitySessionDiagnostic) ServiceSessionRequired() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}





