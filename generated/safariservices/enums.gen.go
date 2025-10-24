// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

/* debug [enums.gen.go]: Generating 6 enums for SafariServices */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum SFErrorCode (5 cases) */
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

/* debug [enums.gen.go]: Processing enum SFSafariServicesVersion (6 cases) */
// SFSafariServicesVersion - The version of Safari services.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion
type SFSafariServicesVersion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version10_0
	SFSafariServicesVersion10_0 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version10_1
	SFSafariServicesVersion10_1 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version11_0
	SFSafariServicesVersion11_0 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version12_0
	SFSafariServicesVersion12_0 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version12_1
	SFSafariServicesVersion12_1 SFSafariServicesVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariServicesVersion/version13_0
	SFSafariServicesVersion13_0 SFSafariServicesVersion = 0
)

/* debug [enums.gen.go]: Processing enum SFSafariViewControllerDismissButtonStyle (3 cases) */
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

/* debug [enums.gen.go]: Processing enum SSReadingListErrorCode (1 cases) */
// SSReadingListErrorCode - Messages that describe a Safari Reading List error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingListError/Code
type SSReadingListErrorCode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingListErrorCode/urlSchemeNotAllowed
	SSReadingListErrorURLSchemeNotAllowed SSReadingListErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum SFAuthenticationError (1 cases) */
// SFAuthenticationError - Messages that describe an authentication error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAuthenticationError-swift.struct/Code
type SFAuthenticationError uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAuthenticationError-swift.enum/canceledLogin
	SFAuthenticationErrorCanceledLogin SFAuthenticationError = 0
)

/* debug [enums.gen.go]: Processing enum SFContentBlockerErrorCode (3 cases) */
// SFContentBlockerErrorCode - Messages that describe a content blocker error.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerErrorCode
type SFContentBlockerErrorCode uint

const (
	// SFContentBlockerLoadingInterrupted - There was an error loading the content blocker extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerErrorCode/loadingInterrupted
	SFContentBlockerLoadingInterrupted SFContentBlockerErrorCode = 0
	// SFContentBlockerNoAttachmentFound - The Content Blocker extension returned an   that did not include an attachment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerErrorCode/noAttachmentFound
	SFContentBlockerNoAttachmentFound SFContentBlockerErrorCode = 0
	// SFContentBlockerNoExtensionFound - A Content Blocker extension with the specified bundle identifier was not found, or the bundle identifier specified an extension that was not owned by you.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerErrorCode/noExtensionFound
	SFContentBlockerNoExtensionFound SFContentBlockerErrorCode = 0
)


