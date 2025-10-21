// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

// Enum types and constants
// OSSystemExtensionErrorCode - Error codes for system extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code
type OSSystemExtensionErrorCode uint

const (
	// OSSystemExtensionErrorAuthorizationRequired - An error code that indicates the system was unable to obtain the proper authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/authorizationRequired
	OSSystemExtensionErrorAuthorizationRequired OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorCodeSignatureInvalid - An error code that indicates the extension’s signature is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/codeSignatureInvalid
	OSSystemExtensionErrorCodeSignatureInvalid OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorDuplicateExtensionIdentifer - An error code that indicates the extension identifier duplicates an existing identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/duplicateExtensionIdentifer
	OSSystemExtensionErrorDuplicateExtensionIdentifer OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorExtensionMissingIdentifier - An error code that indicates the extension identifier is missing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/extensionMissingIdentifier
	OSSystemExtensionErrorExtensionMissingIdentifier OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorExtensionNotFound - An error code that indicates the manager can’t find the system extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/extensionNotFound
	OSSystemExtensionErrorExtensionNotFound OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorForbiddenBySystemPolicy - An error code that indicates the system policy prohibits activating the system extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/forbiddenBySystemPolicy
	OSSystemExtensionErrorForbiddenBySystemPolicy OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorMissingEntitlement - An error code that indicates the system extension lacks a required entitlement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/missingEntitlement
	OSSystemExtensionErrorMissingEntitlement OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorRequestCanceled - An error code that indicates the system extension manager request was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/requestCanceled
	OSSystemExtensionErrorRequestCanceled OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorRequestSuperseded - An error code that indicates the system extension request failed because the system already has a pending request for the same identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/requestSuperseded
	OSSystemExtensionErrorRequestSuperseded OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorUnknown - An error code that indicates an unknown error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/unknown
	OSSystemExtensionErrorUnknown OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorUnknownExtensionCategory - An error code that indicates the extension manager can’t recognize the extension’s category identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/unknownExtensionCategory
	OSSystemExtensionErrorUnknownExtensionCategory OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorUnsupportedParentBundleLocation - An error code that indicates the extension’s parent app isn’t in a valid location for activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/unsupportedParentBundleLocation
	OSSystemExtensionErrorUnsupportedParentBundleLocation OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorValidationFailed - An error code that indicates the manager can’t validate the extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/validationFailed
	OSSystemExtensionErrorValidationFailed OSSystemExtensionErrorCode = 0
)

// OSSystemExtensionReplacementAction - Actions for describing how the extension manager should resolve a version conflict.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/ReplacementAction
type OSSystemExtensionReplacementAction uint

const (
	// OSSystemExtensionReplacementActionCancel - An action that tells the manager to cancel replacement of a system extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/ReplacementAction/cancel
	OSSystemExtensionReplacementActionCancel OSSystemExtensionReplacementAction = 0
	// OSSystemExtensionReplacementActionReplace - An action that tells the manager to replace an existing system extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/ReplacementAction/replace
	OSSystemExtensionReplacementActionReplace OSSystemExtensionReplacementAction = 0
)

// OSSystemExtensionRequestResult - The result of a completed request, possibly including additional information about the extension’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/Result
type OSSystemExtensionRequestResult uint

const (
	// OSSystemExtensionRequestCompleted - The request completed successfully.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionRequest/Result/completed
	OSSystemExtensionRequestCompleted OSSystemExtensionRequestResult = 0
)


