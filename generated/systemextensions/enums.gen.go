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
	// OSSystemExtensionErrorDuplicateExtensionIdentifer - An error code that indicates the extension identifier duplicates an existing identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/duplicateExtensionIdentifer
	OSSystemExtensionErrorDuplicateExtensionIdentifer OSSystemExtensionErrorCode = 0
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
	// OSSystemExtensionErrorRequestSuperseded - An error code that indicates the system extension request failed because the system already has a pending request for the same identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/requestSuperseded
	OSSystemExtensionErrorRequestSuperseded OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorUnknownExtensionCategory - An error code that indicates the extension manager can’t recognize the extension’s category identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/unknownExtensionCategory
	OSSystemExtensionErrorUnknownExtensionCategory OSSystemExtensionErrorCode = 0
	// OSSystemExtensionErrorUnsupportedParentBundleLocation - An error code that indicates the extension’s parent app isn’t in a valid location for activation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionError/Code/unsupportedParentBundleLocation
	OSSystemExtensionErrorUnsupportedParentBundleLocation OSSystemExtensionErrorCode = 0
)


