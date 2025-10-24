// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

// Enum types and constants
// SFAuthenticationError enum type
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAuthenticationError-swift.enum
type SFAuthenticationError uint

// SFContentBlockerErrorCode - Messages that describe a content blocker error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerErrorCode
type SFContentBlockerErrorCode uint

// SFErrorCode - Messages that describe a content blocker or Safari app extension error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFError/Code
type SFErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFErrorCode/internalError
	SFErrorInternalError SFErrorCode = 0
	// SFErrorLoadingInterrupted - There was an error loading the content blocker extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFErrorCode/loadingInterrupted
	SFErrorLoadingInterrupted SFErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFErrorCode/missingEntitlement
	SFErrorMissingEntitlement SFErrorCode = 0
	// SFErrorNoAttachmentFound - The Content Blocker extension returned an   that did not include an attachment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFErrorCode/noAttachmentFound
	SFErrorNoAttachmentFound SFErrorCode = 0
	// SFErrorNoExtensionFound - A Content Blocker or Safari app extension with the specified bundle identifier was not found, or the bundle identifier specified an extension that was not owned by you.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFErrorCode/noExtensionFound
	SFErrorNoExtensionFound SFErrorCode = 0
)

// SFSafariServicesVersion - The version of Safari services.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion
type SFSafariServicesVersion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version10_0
	SFSafariServicesVersion10_0 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version12_0
	SFSafariServicesVersion12_0 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version12_1
	SFSafariServicesVersion12_1 SFSafariServicesVersion = 0
)

// SFSafariViewControllerDismissButtonStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DismissButtonStyle-swift.enum
type SFSafariViewControllerDismissButtonStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DismissButtonStyle-swift.enum/cancel
	SFSafariViewControllerDismissButtonStyleCancel SFSafariViewControllerDismissButtonStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DismissButtonStyle-swift.enum/close
	SFSafariViewControllerDismissButtonStyleClose SFSafariViewControllerDismissButtonStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/DismissButtonStyle-swift.enum/done
	SFSafariViewControllerDismissButtonStyleDone SFSafariViewControllerDismissButtonStyle = 0
)

// SSReadingListErrorCode - Messages that describe a Safari Reading List error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingListError/Code
type SSReadingListErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingListErrorCode/urlSchemeNotAllowed
	SSReadingListErrorURLSchemeNotAllowed SSReadingListErrorCode = 0
)


