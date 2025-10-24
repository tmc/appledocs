// Code generated from Apple documentation for AppTrackingTransparency. DO NOT EDIT.

package apptrackingtransparency

/* debug [enums.gen.go]: Generating 1 enums for AppTrackingTransparency */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ATTrackingManagerAuthorizationStatus (4 cases) */
// ATTrackingManagerAuthorizationStatus - The status values for app tracking authorization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/AuthorizationStatus
type ATTrackingManagerAuthorizationStatus uint

const (
	// ATTrackingManagerAuthorizationStatusAuthorized - The value that returns if the user authorizes access to app-related data for   tracking the user or the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/AuthorizationStatus/authorized
	ATTrackingManagerAuthorizationStatusAuthorized ATTrackingManagerAuthorizationStatus = 0
	// ATTrackingManagerAuthorizationStatusDenied - The value that returns if the user denies authorization to access   app-related data for tracking the user or the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/AuthorizationStatus/denied
	ATTrackingManagerAuthorizationStatusDenied ATTrackingManagerAuthorizationStatus = 0
	// ATTrackingManagerAuthorizationStatusNotDetermined - The value that returns when the app can’t determine the user’s   authorization status for access to app-related data for tracking the   user or the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/AuthorizationStatus/notDetermined
	ATTrackingManagerAuthorizationStatusNotDetermined ATTrackingManagerAuthorizationStatus = 0
	// ATTrackingManagerAuthorizationStatusRestricted - The value that returns if authorization to access app-related data for   tracking the user or the device has a restricted status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppTrackingTransparency/ATTrackingManager/AuthorizationStatus/restricted
	ATTrackingManagerAuthorizationStatusRestricted ATTrackingManagerAuthorizationStatus = 0
)


