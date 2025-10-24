//go:build darwin && ios

// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for ServiceSessionDiagnostic

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/alwaysAuthorizationDenied
func (s_ ServiceSessionDiagnostic) AlwaysAuthorizationDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("alwaysAuthorizationDenied"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationDenied
func (s_ ServiceSessionDiagnostic) AuthorizationDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationDenied"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationDeniedGlobally
func (s_ ServiceSessionDiagnostic) AuthorizationDeniedGlobally() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationDeniedGlobally"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationRequestInProgress
func (s_ ServiceSessionDiagnostic) AuthorizationRequestInProgress() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationRequestInProgress"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/authorizationRestricted
func (s_ ServiceSessionDiagnostic) AuthorizationRestricted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("authorizationRestricted"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/fullAccuracyDenied
func (s_ ServiceSessionDiagnostic) FullAccuracyDenied() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("fullAccuracyDenied"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/insufficientlyInUse
func (s_ ServiceSessionDiagnostic) InsufficientlyInUse() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("insufficientlyInUse"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLServiceSessionDiagnostic/serviceSessionRequired
func (s_ ServiceSessionDiagnostic) ServiceSessionRequired() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("serviceSessionRequired"))
	return rv
}
