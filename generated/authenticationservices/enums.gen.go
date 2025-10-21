// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// Enum types and constants
// ASAuthorizationAppleIDButtonType - A type for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType
type AuthorizationAppleIDButtonType uint

const (
// AuthorizationAppleIDButtonTypeContinue - A button type that continues the Sign in with Apple authorization process.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/continue
AuthorizationAppleIDButtonTypeContinue AuthorizationAppleIDButtonType = 0
// AuthorizationAppleIDButtonTypeDefault - A default button type for the Sign in with Apple authorization process.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/default
AuthorizationAppleIDButtonTypeDefault AuthorizationAppleIDButtonType = 0
// AuthorizationAppleIDButtonTypeSignIn - A button type that performs authorization using Sign in with Apple.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/signIn
AuthorizationAppleIDButtonTypeSignIn AuthorizationAppleIDButtonType = 0
// AuthorizationAppleIDButtonTypeSignUp - A button type that allows the user to sign up for Sign in with Apple.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/signUp
AuthorizationAppleIDButtonTypeSignUp AuthorizationAppleIDButtonType = 0
)

// ASAuthorizationAppleIDButtonStyle - A style for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style
type AuthorizationAppleIDButtonStyle uint

const (
// AuthorizationAppleIDButtonStyleBlack - A black button.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/black
AuthorizationAppleIDButtonStyleBlack AuthorizationAppleIDButtonStyle = 0
// AuthorizationAppleIDButtonStyleWhite - A white button.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/white
AuthorizationAppleIDButtonStyleWhite AuthorizationAppleIDButtonStyle = 0
// AuthorizationAppleIDButtonStyleWhiteOutline - A button with a white outline.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/whiteOutline
AuthorizationAppleIDButtonStyleWhiteOutline AuthorizationAppleIDButtonStyle = 0
)

// ASAuthorizationControllerRequestOptions - Options that modify how a controller performs authorization requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/RequestOptions
type AuthorizationControllerRequestOptions uint

const (
// AuthorizationControllerRequestOptionPreferImmediatelyAvailableCredentials - Tells the authorization controller to prefer credentials that are immediately available on the local device.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/RequestOptions/preferImmediatelyAvailableCredentials
AuthorizationControllerRequestOptionPreferImmediatelyAvailableCredentials AuthorizationControllerRequestOptions = 0
)

// ASAuthorizationError - Codes that authorization errors can have.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code
type AuthorizationError uint

// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum
type AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle uint

const (
// AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional - Perform a conditional request. This style of request is meant to opportunistically add passkeys to existing   password-based accounts, at the discretion of the user’s credential manager. It should be performed   shortly after a user has signed in with a password. If the user is using a password and passkey manager,   and certain internal conditions of that credential manager are met (e.g. the user signed in recently with a   matching password-based account and does not yet have a passkey for this account), then this request   may proceed automatically, without further user interaction. If any of the internal conditions are not met,   this request will return an error without showing any UI to the user, and may be retried the next time they   sign in.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum/conditional
AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle = 0
// AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleStandard - Perform a request using the standard presentation style. This is the default style.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum/standard
AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleStandard AuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle = 0
)

// ASAuthorizationProviderExtensionFederationType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum
type AuthorizationProviderExtensionFederationType uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum/dynamicWSTrust
AuthorizationProviderExtensionFederationTypeDynamicWSTrust AuthorizationProviderExtensionFederationType = 0
)

// ASAuthorizationProviderExtensionPlatformSSOProtocolVersion enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionPlatformSSOProtocolVersion
type AuthorizationProviderExtensionPlatformSSOProtocolVersion uint

// ASAuthorizationPublicKeyCredentialAttachment enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialAttachment
type AuthorizationPublicKeyCredentialAttachment uint

// ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation
type AuthorizationPublicKeyCredentialLargeBlobAssertionOperation uint

// ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement
type AuthorizationPublicKeyCredentialLargeBlobSupportRequirement uint

// ASCredentialIdentityTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes
type CredentialIdentityTypes uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/oneTimeCode
CredentialIdentityTypesOneTimeCode CredentialIdentityTypes = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/passkey
CredentialIdentityTypesPasskey CredentialIdentityTypes = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/password
CredentialIdentityTypesPassword CredentialIdentityTypes = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityTypes/ASCredentialIdentityTypesAll
CredentialIdentityTypesAll CredentialIdentityTypes = 0
)

// ASPublicKeyCredentialClientDataCrossOriginValue enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientDataCrossOriginValue
type PublicKeyCredentialClientDataCrossOriginValue uint

// ASUserAgeRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange
type UserAgeRange uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/child
UserAgeRangeChild UserAgeRange = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/notChild
UserAgeRangeNotChild UserAgeRange = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/unknown
UserAgeRangeUnknown UserAgeRange = 0
)

// ASUserDetectionStatus - Possible values for the real user indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus
type UserDetectionStatus uint

const (
// UserDetectionStatusLikelyReal - The user appears to be a real person.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/likelyReal
UserDetectionStatusLikelyReal UserDetectionStatus = 0
// UserDetectionStatusUnknown - The system hasn’t determined whether the user might be a real person.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/unknown
UserDetectionStatusUnknown UserDetectionStatus = 0
// UserDetectionStatusUnsupported - The system can’t determine this user’s status as a real person.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/unsupported
UserDetectionStatusUnsupported UserDetectionStatus = 0
)

// ASWebAuthenticationSessionErrorCode - The error code for a web authentication session error.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code
type WebAuthenticationSessionErrorCode uint

const (
// WebAuthenticationSessionErrorCodeCanceledLogin - The login has been canceled.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/canceledLogin
WebAuthenticationSessionErrorCodeCanceledLogin WebAuthenticationSessionErrorCode = 0
// WebAuthenticationSessionErrorCodePresentationContextInvalid - The context was invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/presentationContextInvalid
WebAuthenticationSessionErrorCodePresentationContextInvalid WebAuthenticationSessionErrorCode = 0
// WebAuthenticationSessionErrorCodePresentationContextNotProvided - A context wasn’t provided.
//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/presentationContextNotProvided
WebAuthenticationSessionErrorCodePresentationContextNotProvided WebAuthenticationSessionErrorCode = 0
)


