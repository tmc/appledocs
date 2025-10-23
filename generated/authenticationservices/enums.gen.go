// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

// Enum types and constants
// ASAuthorizationAppleIDButtonType - A type for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType
type ASAuthorizationAppleIDButtonType uint

// ASAuthorizationAppleIDButtonStyle - A style for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style
type ASAuthorizationAppleIDButtonStyle uint

// ASAuthorizationAppleIDProviderCredentialState - Possible values for the credential state of a user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState
type ASAuthorizationAppleIDProviderCredentialState uint

const (
	// ASAuthorizationAppleIDProviderCredentialNotFound - The user hasn’t established a relationship with Sign in with Apple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState/notFound
	ASAuthorizationAppleIDProviderCredentialNotFound ASAuthorizationAppleIDProviderCredentialState = 0
	// ASAuthorizationAppleIDProviderCredentialRevoked - The given user’s authorization has been revoked and they should be signed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState/revoked
	ASAuthorizationAppleIDProviderCredentialRevoked ASAuthorizationAppleIDProviderCredentialState = 0
	// ASAuthorizationAppleIDProviderCredentialTransferred - The app has been transferred to a different team, and you need to migrate the user’s identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState/transferred
	ASAuthorizationAppleIDProviderCredentialTransferred ASAuthorizationAppleIDProviderCredentialState = 0
)

// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum
type ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle uint

const (
	// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional - Perform a conditional request. This style of request is meant to opportunistically add passkeys to existing   password-based accounts, at the discretion of the user’s credential manager. It should be performed   shortly after a user has signed in with a password. If the user is using a password and passkey manager,   and certain internal conditions of that credential manager are met (e.g. the user signed in recently with a   matching password-based account and does not yet have a passkey for this account), then this request   may proceed automatically, without further user interaction. If any of the internal conditions are not met,   this request will return an error without showing any UI to the user, and may be retried the next time they   sign in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum/conditional
	ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle = 0
)

// ASAuthorizationPublicKeyCredentialAttachment enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialAttachment
type ASAuthorizationPublicKeyCredentialAttachment uint


