// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

// Enum types and constants
// ABAuthorizationStatus - Different possible values for the authorization status of an app with respect to address book data.
//
// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAuthorizationStatus
type ABAuthorizationStatus uint

const (
// kABAuthorizationStatusAuthorized - The app is authorized to access address book data.
//
	// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAuthorizationStatus/authorized
kABAuthorizationStatusAuthorized ABAuthorizationStatus = 0
// kABAuthorizationStatusDenied - The user explicitly denied access to address book data for this app.
//
	// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAuthorizationStatus/denied
kABAuthorizationStatusDenied ABAuthorizationStatus = 0
// kABAuthorizationStatusNotDetermined - No authorization status could be determined.
//
	// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAuthorizationStatus/notDetermined
kABAuthorizationStatusNotDetermined ABAuthorizationStatus = 0
// kABAuthorizationStatusRestricted - The app is not authorized to access address book data. The user cannot change this access, possibly due to restrictions such as parental controls.
//
	// [Full Topic]: https://developer.apple.com/documentation/AddressBook/ABAuthorizationStatus/restricted
kABAuthorizationStatusRestricted ABAuthorizationStatus = 0
)


