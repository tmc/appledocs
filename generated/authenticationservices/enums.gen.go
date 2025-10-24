// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

/* debug [enums.gen.go]: Generating 27 enums for AuthenticationServices */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum ASAuthorizationAppleIDProviderCredentialState (4 cases) */
// ASAuthorizationAppleIDProviderCredentialState - Possible values for the credential state of a user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState
type ASAuthorizationAppleIDProviderCredentialState uint

const (
	// ASAuthorizationAppleIDProviderCredentialAuthorized - The user is authorized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDProvider/CredentialState/authorized
	ASAuthorizationAppleIDProviderCredentialAuthorized ASAuthorizationAppleIDProviderCredentialState = 0
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

/* debug [enums.gen.go]: Processing enum ASAuthorizationControllerRequestOptions (1 cases) */
// ASAuthorizationControllerRequestOptions - Options that modify how a controller performs authorization requests.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/RequestOptions
type ASAuthorizationControllerRequestOptions uint

const (
	// ASAuthorizationControllerRequestOptionPreferImmediatelyAvailableCredentials - Tells the authorization controller to prefer credentials that are immediately available on the local device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationController/RequestOptions/preferImmediatelyAvailableCredentials
	ASAuthorizationControllerRequestOptionPreferImmediatelyAvailableCredentials ASAuthorizationControllerRequestOptions = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle (2 cases) */
// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum
type ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle uint

const (
	// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional - Perform a conditional request. This style of request is meant to opportunistically add passkeys to existing   password-based accounts, at the discretion of the user’s credential manager. It should be performed   shortly after a user has signed in with a password. If the user is using a password and passkey manager,   and certain internal conditions of that credential manager are met (e.g. the user signed in recently with a   matching password-based account and does not yet have a passkey for this account), then this request   may proceed automatically, without further user interaction. If any of the internal conditions are not met,   this request will return an error without showing any UI to the user, and may be retried the next time they   sign in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum/conditional
	ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleConditional ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle = 0
	// ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleStandard - Perform a request using the standard presentation style. This is the default style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialRegistrationRequest/RequestStyle-swift.enum/standard
	ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyleStandard ASAuthorizationPlatformPublicKeyCredentialRegistrationRequestStyle = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationAppleIDButtonType (4 cases) */
// ASAuthorizationAppleIDButtonType - A type for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType
type ASAuthorizationAppleIDButtonType uint

const (
	// ASAuthorizationAppleIDButtonTypeContinue - A button type that continues the Sign in with Apple authorization process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/continue
	ASAuthorizationAppleIDButtonTypeContinue ASAuthorizationAppleIDButtonType = 0
	// ASAuthorizationAppleIDButtonTypeDefault - A default button type for the Sign in with Apple authorization process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/default
	ASAuthorizationAppleIDButtonTypeDefault ASAuthorizationAppleIDButtonType = 0
	// ASAuthorizationAppleIDButtonTypeSignIn - A button type that performs authorization using Sign in with Apple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/signIn
	ASAuthorizationAppleIDButtonTypeSignIn ASAuthorizationAppleIDButtonType = 0
	// ASAuthorizationAppleIDButtonTypeSignUp - A button type that allows the user to sign up for Sign in with Apple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/ButtonType/signUp
	ASAuthorizationAppleIDButtonTypeSignUp ASAuthorizationAppleIDButtonType = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationAppleIDButtonStyle (3 cases) */
// ASAuthorizationAppleIDButtonStyle - A style for the authorization button.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style
type ASAuthorizationAppleIDButtonStyle uint

const (
	// ASAuthorizationAppleIDButtonStyleBlack - A black button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/black
	ASAuthorizationAppleIDButtonStyleBlack ASAuthorizationAppleIDButtonStyle = 0
	// ASAuthorizationAppleIDButtonStyleWhite - A white button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/white
	ASAuthorizationAppleIDButtonStyleWhite ASAuthorizationAppleIDButtonStyle = 0
	// ASAuthorizationAppleIDButtonStyleWhiteOutline - A button with a white outline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationAppleIDButton/Style/whiteOutline
	ASAuthorizationAppleIDButtonStyleWhiteOutline ASAuthorizationAppleIDButtonStyle = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationError (11 cases) */
// ASAuthorizationError - Codes that authorization errors can have.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code
type ASAuthorizationError uint

const (
	// ASAuthorizationErrorCanceled - The user canceled the authorization attempt.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/canceled
	ASAuthorizationErrorCanceled ASAuthorizationError = 0
	// ASAuthorizationErrorCredentialExport - The credential export request failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/credentialExport
	ASAuthorizationErrorCredentialExport ASAuthorizationError = 0
	// ASAuthorizationErrorCredentialImport - The credential import request failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/credentialImport
	ASAuthorizationErrorCredentialImport ASAuthorizationError = 0
	// ASAuthorizationErrorDeviceNotConfiguredForPasskeyCreation - This error signals that the device is not currently set up to create passkeys.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/deviceNotConfiguredForPasskeyCreation
	ASAuthorizationErrorDeviceNotConfiguredForPasskeyCreation ASAuthorizationError = 0
	// ASAuthorizationErrorFailed - The authorization attempt failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/failed
	ASAuthorizationErrorFailed ASAuthorizationError = 0
	// ASAuthorizationErrorInvalidResponse - The authorization request received an invalid response.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/invalidResponse
	ASAuthorizationErrorInvalidResponse ASAuthorizationError = 0
	// ASAuthorizationErrorMatchedExcludedCredential - This error should only be returned when specifying @c excludedCredentials on a public key credential registration request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/matchedExcludedCredential
	ASAuthorizationErrorMatchedExcludedCredential ASAuthorizationError = 0
	// ASAuthorizationErrorNotHandled - The authorization request wasn’t handled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/notHandled
	ASAuthorizationErrorNotHandled ASAuthorizationError = 0
	// ASAuthorizationErrorNotInteractive - The authorization request isn’t interactive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/notInteractive
	ASAuthorizationErrorNotInteractive ASAuthorizationError = 0
	// ASAuthorizationErrorPreferSignInWithApple - This error signals the user has an existing Sign in with Apple account that they would prefer to use instead of continuing the current request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/preferSignInWithApple
	ASAuthorizationErrorPreferSignInWithApple ASAuthorizationError = 0
	// ASAuthorizationErrorUnknown - The authorization attempt failed for an unknown reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationError-swift.struct/Code/unknown
	ASAuthorizationErrorUnknown ASAuthorizationError = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionAuthenticationMethod (3 cases) */
// ASAuthorizationProviderExtensionAuthenticationMethod - The platform single sign-on method for the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthenticationMethod
type ASAuthorizationProviderExtensionAuthenticationMethod uint

const (
	// ASAuthorizationProviderExtensionAuthenticationMethodPassword - Password authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthenticationMethod/password
	ASAuthorizationProviderExtensionAuthenticationMethodPassword ASAuthorizationProviderExtensionAuthenticationMethod = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthenticationMethod/smartCard
	ASAuthorizationProviderExtensionAuthenticationMethodSmartCard ASAuthorizationProviderExtensionAuthenticationMethod = 0
	// ASAuthorizationProviderExtensionAuthenticationMethodUserSecureEnclaveKey - Secure Enclave key authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionAuthenticationMethod/userSecureEnclaveKey
	ASAuthorizationProviderExtensionAuthenticationMethodUserSecureEnclaveKey ASAuthorizationProviderExtensionAuthenticationMethod = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionKeyType (8 cases) */
// ASAuthorizationProviderExtensionKeyType - The key types for platform single sign-on.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType
type ASAuthorizationProviderExtensionKeyType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/currentDeviceEncryption
	ASAuthorizationProviderExtensionKeyTypeCurrentDeviceEncryption ASAuthorizationProviderExtensionKeyType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/currentDeviceSigning
	ASAuthorizationProviderExtensionKeyTypeCurrentDeviceSigning ASAuthorizationProviderExtensionKeyType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/sharedDeviceEncryption
	ASAuthorizationProviderExtensionKeyTypeSharedDeviceEncryption ASAuthorizationProviderExtensionKeyType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/sharedDeviceSigning
	ASAuthorizationProviderExtensionKeyTypeSharedDeviceSigning ASAuthorizationProviderExtensionKeyType = 0
	// ASAuthorizationProviderExtensionKeyTypeUserDeviceEncryption - The user device encryption key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/userDeviceEncryption
	ASAuthorizationProviderExtensionKeyTypeUserDeviceEncryption ASAuthorizationProviderExtensionKeyType = 0
	// ASAuthorizationProviderExtensionKeyTypeUserDeviceSigning - The user device signing key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/userDeviceSigning
	ASAuthorizationProviderExtensionKeyTypeUserDeviceSigning ASAuthorizationProviderExtensionKeyType = 0
	// ASAuthorizationProviderExtensionKeyTypeUserSecureEnclaveKey - The user Secure Enclave key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/userSecureEnclaveKey
	ASAuthorizationProviderExtensionKeyTypeUserSecureEnclaveKey ASAuthorizationProviderExtensionKeyType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionKeyType/userSmartCard
	ASAuthorizationProviderExtensionKeyTypeUserSmartCard ASAuthorizationProviderExtensionKeyType = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionFederationType (3 cases) */
// ASAuthorizationProviderExtensionFederationType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum
type ASAuthorizationProviderExtensionFederationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum/dynamicWSTrust
	ASAuthorizationProviderExtensionFederationTypeDynamicWSTrust ASAuthorizationProviderExtensionFederationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum/none
	ASAuthorizationProviderExtensionFederationTypeNone ASAuthorizationProviderExtensionFederationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/FederationType-swift.enum/wsTrust
	ASAuthorizationProviderExtensionFederationTypeWSTrust ASAuthorizationProviderExtensionFederationType = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy (5 cases) */
// ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/UserSecureEnclaveKeyBiometricPolicy-swift.struct
type ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/UserSecureEnclaveKeyBiometricPolicy-swift.struct/passwordFallback
	ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyPasswordFallback ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/UserSecureEnclaveKeyBiometricPolicy-swift.struct/reuseDuringUnlock
	ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyReuseDuringUnlock ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/UserSecureEnclaveKeyBiometricPolicy-swift.struct/touchIDOrWatchAny
	ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyTouchIDOrWatchAny ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginConfiguration/UserSecureEnclaveKeyBiometricPolicy-swift.struct/touchIDOrWatchCurrentSet
	ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyTouchIDOrWatchCurrentSet ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy/ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyNone
	ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicyNone ASAuthorizationProviderExtensionUserSecureEnclaveKeyBiometricPolicy = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionPlatformSSOProtocolVersion (2 cases) */
// ASAuthorizationProviderExtensionPlatformSSOProtocolVersion enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionPlatformSSOProtocolVersion
type ASAuthorizationProviderExtensionPlatformSSOProtocolVersion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionPlatformSSOProtocolVersion/version1_0
	ASAuthorizationProviderExtensionPlatformSSOProtocolVersion1_0 ASAuthorizationProviderExtensionPlatformSSOProtocolVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionPlatformSSOProtocolVersion/version2_0
	ASAuthorizationProviderExtensionPlatformSSOProtocolVersion2_0 ASAuthorizationProviderExtensionPlatformSSOProtocolVersion = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionRegistrationResult (4 cases) */
// ASAuthorizationProviderExtensionRegistrationResult - The registration result.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationResult
type ASAuthorizationProviderExtensionRegistrationResult uint

const (
	// ASAuthorizationProviderExtensionRegistrationResultFailed - The registration fails to complete and the system retries later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationResult/failed
	ASAuthorizationProviderExtensionRegistrationResultFailed ASAuthorizationProviderExtensionRegistrationResult = 0
	// ASAuthorizationProviderExtensionRegistrationResultFailedNoRetry - The registration fails to complete and the system doesn’t retry later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationResult/failedNoRetry
	ASAuthorizationProviderExtensionRegistrationResultFailedNoRetry ASAuthorizationProviderExtensionRegistrationResult = 0
	// ASAuthorizationProviderExtensionRegistrationResultSuccess - The registration succeeds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationResult/success
	ASAuthorizationProviderExtensionRegistrationResultSuccess ASAuthorizationProviderExtensionRegistrationResult = 0
	// ASAuthorizationProviderExtensionRegistrationResultUserInterfaceRequired - The user interface is required to complete registration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRegistrationResult/userInterfaceRequired
	ASAuthorizationProviderExtensionRegistrationResultUserInterfaceRequired ASAuthorizationProviderExtensionRegistrationResult = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionRequestOptions (8 cases) */
// ASAuthorizationProviderExtensionRequestOptions - The options for the extension to obtain the status of the registration.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions
type ASAuthorizationProviderExtensionRequestOptions uint

const (
	// ASAuthorizationProviderExtensionRequestOptionsNone - Options aren’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/ASAuthorizationProviderExtensionRequestOptionsNone
	ASAuthorizationProviderExtensionRequestOptionsNone ASAuthorizationProviderExtensionRequestOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/registrationDeviceKeyMigration
	ASAuthorizationProviderExtensionRequestOptionsRegistrationDeviceKeyMigration ASAuthorizationProviderExtensionRequestOptions = 0
	// ASAuthorizationProviderExtensionRequestOptionsRegistrationRepair - Indicates that the registration is undergoing repair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/registrationRepair
	ASAuthorizationProviderExtensionRequestOptionsRegistrationRepair ASAuthorizationProviderExtensionRequestOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/registrationSharedDeviceKeys
	ASAuthorizationProviderExtensionRequestOptionsRegistrationSharedDeviceKeys ASAuthorizationProviderExtensionRequestOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/setupAssistant
	ASAuthorizationProviderExtensionRequestOptionsSetupAssistant ASAuthorizationProviderExtensionRequestOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/strongerKeyAvailable
	ASAuthorizationProviderExtensionRequestOptionsStrongerKeyAvailable ASAuthorizationProviderExtensionRequestOptions = 0
	// ASAuthorizationProviderExtensionRequestOptionsUserInteractionEnabled - Indicates that the user interface is in an enabled state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/userInteractionEnabled
	ASAuthorizationProviderExtensionRequestOptionsUserInteractionEnabled ASAuthorizationProviderExtensionRequestOptions = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionRequestOptions/userKeyInvalid
	ASAuthorizationProviderExtensionRequestOptionsUserKeyInvalid ASAuthorizationProviderExtensionRequestOptions = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationProviderExtensionSupportedGrantTypes (5 cases) */
// ASAuthorizationProviderExtensionSupportedGrantTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes
type ASAuthorizationProviderExtensionSupportedGrantTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes/ASAuthorizationProviderExtensionSupportedGrantTypesNone
	ASAuthorizationProviderExtensionSupportedGrantTypesNone ASAuthorizationProviderExtensionSupportedGrantTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes/jwtBearer
	ASAuthorizationProviderExtensionSupportedGrantTypesJWTBearer ASAuthorizationProviderExtensionSupportedGrantTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes/password
	ASAuthorizationProviderExtensionSupportedGrantTypesPassword ASAuthorizationProviderExtensionSupportedGrantTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes/saml1_1
	ASAuthorizationProviderExtensionSupportedGrantTypesSAML1_1 ASAuthorizationProviderExtensionSupportedGrantTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionSupportedGrantTypes/saml2_0
	ASAuthorizationProviderExtensionSupportedGrantTypesSAML2_0 ASAuthorizationProviderExtensionSupportedGrantTypes = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationPublicKeyCredentialAttachment (2 cases) */
// ASAuthorizationPublicKeyCredentialAttachment enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialAttachment
type ASAuthorizationPublicKeyCredentialAttachment uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialAttachment/crossPlatform
	ASAuthorizationPublicKeyCredentialAttachmentCrossPlatform ASAuthorizationPublicKeyCredentialAttachment = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialAttachment/platform
	ASAuthorizationPublicKeyCredentialAttachmentPlatform ASAuthorizationPublicKeyCredentialAttachment = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation (2 cases) */
// ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation
type ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperationRead
	ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperationRead ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperationWrite
	ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperationWrite ASAuthorizationPublicKeyCredentialLargeBlobAssertionOperation = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement (2 cases) */
// ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement
type ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirementPreferred
	ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirementPreferred ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement/ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirementRequired
	ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirementRequired ASAuthorizationPublicKeyCredentialLargeBlobSupportRequirement = 0
)

/* debug [enums.gen.go]: Processing enum ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState (3 cases) */
// ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState - An enumeration of values that indicate whether the browser app has access to a person’s passkeys.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/AuthorizationState
type ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState uint

const (
	// ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateAuthorized - Someone allows the browser app to use passkeys stored in the keychain, and managed by third-party credential manager apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/AuthorizationState/authorized
	ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateAuthorized ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState = 0
	// ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateDenied - Someone forbids the browser app to use passkeys stored in the keychain, and managed by third-party credential manager apps.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/AuthorizationState/denied
	ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateDenied ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState = 0
	// ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateNotDetermined - The person has yet to choose whether to allow the browser app to access passkeys stored on the keychain, or managed by third-party credential managers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/AuthorizationState/notDetermined
	ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationStateNotDetermined ASAuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState = 0
)

/* debug [enums.gen.go]: Processing enum ASCredentialIdentityTypes (4 cases) */
// ASCredentialIdentityTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes
type ASCredentialIdentityTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/oneTimeCode
	ASCredentialIdentityTypesOneTimeCode ASCredentialIdentityTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/passkey
	ASCredentialIdentityTypesPasskey ASCredentialIdentityTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStore/IdentityTypes/password
	ASCredentialIdentityTypesPassword ASCredentialIdentityTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityTypes/ASCredentialIdentityTypesAll
	ASCredentialIdentityTypesAll ASCredentialIdentityTypes = 0
)

/* debug [enums.gen.go]: Processing enum ASCredentialIdentityStoreErrorCode (3 cases) */
// ASCredentialIdentityStoreErrorCode - Constants that represent credential identity store error codes.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreError/Code
type ASCredentialIdentityStoreErrorCode uint

const (
	// ASCredentialIdentityStoreErrorCodeInternalError - The operation failed due to an internal error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreError/Code/internalError
	ASCredentialIdentityStoreErrorCodeInternalError ASCredentialIdentityStoreErrorCode = 0
	// ASCredentialIdentityStoreErrorCodeStoreBusy - The operation failed because the credential identity store is busy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreError/Code/storeBusy
	ASCredentialIdentityStoreErrorCodeStoreBusy ASCredentialIdentityStoreErrorCode = 0
	// ASCredentialIdentityStoreErrorCodeStoreDisabled - The operation failed because the credential identity store is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialIdentityStoreError/Code/storeDisabled
	ASCredentialIdentityStoreErrorCodeStoreDisabled ASCredentialIdentityStoreErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ASCredentialRequestType (4 cases) */
// ASCredentialRequestType - An enumeration that identifies different types of credentials that apps and websites can request.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialRequestType
type ASCredentialRequestType uint

const (
	// ASCredentialRequestTypeOneTimeCode - The app or website is requesting a one-time passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialRequestType/oneTimeCode
	ASCredentialRequestTypeOneTimeCode ASCredentialRequestType = 0
	// ASCredentialRequestTypePasskeyAssertion - The app or website is requesting a passkey assertion credential.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialRequestType/passkeyAssertion
	ASCredentialRequestTypePasskeyAssertion ASCredentialRequestType = 0
	// ASCredentialRequestTypePasskeyRegistration - The app or website is requesting a passkey registration credential.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialRequestType/passkeyRegistration
	ASCredentialRequestTypePasskeyRegistration ASCredentialRequestType = 0
	// ASCredentialRequestTypePassword - The app or website is requesting a password credential.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialRequestType/password
	ASCredentialRequestTypePassword ASCredentialRequestType = 0
)

/* debug [enums.gen.go]: Processing enum ASCredentialServiceIdentifierType (2 cases) */
// ASCredentialServiceIdentifierType - Possible values for the service identifier type.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/IdentifierType
type ASCredentialServiceIdentifierType uint

const (
	// ASCredentialServiceIdentifierTypeDomain - A domain service identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/IdentifierType/domain
	ASCredentialServiceIdentifierTypeDomain ASCredentialServiceIdentifierType = 0
	// ASCredentialServiceIdentifierTypeURL - A URL service identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASCredentialServiceIdentifier/IdentifierType/URL
	ASCredentialServiceIdentifierTypeURL ASCredentialServiceIdentifierType = 0
)

/* debug [enums.gen.go]: Processing enum ASExtensionErrorCode (5 cases) */
// ASExtensionErrorCode - The codes for a credential provider extension error.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code
type ASExtensionErrorCode uint

const (
	// ASExtensionErrorCodeCredentialIdentityNotFound - The credential identity was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code/credentialIdentityNotFound
	ASExtensionErrorCodeCredentialIdentityNotFound ASExtensionErrorCode = 0
	// ASExtensionErrorCodeFailed - The operation failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code/failed
	ASExtensionErrorCodeFailed ASExtensionErrorCode = 0
	// ASExtensionErrorCodeMatchedExcludedCredential - This error should only be used for a passkey registration request, if the @c excludedCredentials property matches a known passkey.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code/matchedExcludedCredential
	ASExtensionErrorCodeMatchedExcludedCredential ASExtensionErrorCode = 0
	// ASExtensionErrorCodeUserCanceled - The user canceled the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code/userCanceled
	ASExtensionErrorCodeUserCanceled ASExtensionErrorCode = 0
	// ASExtensionErrorCodeUserInteractionRequired - User interaction is required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASExtensionError/Code/userInteractionRequired
	ASExtensionErrorCodeUserInteractionRequired ASExtensionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum ASPublicKeyCredentialClientDataCrossOriginValue (3 cases) */
// ASPublicKeyCredentialClientDataCrossOriginValue enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientDataCrossOriginValue
type ASPublicKeyCredentialClientDataCrossOriginValue uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientDataCrossOriginValue/crossOrigin
	ASPublicKeyCredentialClientDataCrossOriginValueCrossOrigin ASPublicKeyCredentialClientDataCrossOriginValue = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientDataCrossOriginValue/notSet
	ASPublicKeyCredentialClientDataCrossOriginValueNotSet ASPublicKeyCredentialClientDataCrossOriginValue = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPublicKeyCredentialClientDataCrossOriginValue/sameOriginWithAncestors
	ASPublicKeyCredentialClientDataCrossOriginValueSameOriginWithAncestors ASPublicKeyCredentialClientDataCrossOriginValue = 0
)

/* debug [enums.gen.go]: Processing enum ASUserAgeRange (3 cases) */
// ASUserAgeRange enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange
type ASUserAgeRange uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/child
	ASUserAgeRangeChild ASUserAgeRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/notChild
	ASUserAgeRangeNotChild ASUserAgeRange = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserAgeRange/unknown
	ASUserAgeRangeUnknown ASUserAgeRange = 0
)

/* debug [enums.gen.go]: Processing enum ASUserDetectionStatus (3 cases) */
// ASUserDetectionStatus - Possible values for the real user indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus
type ASUserDetectionStatus uint

const (
	// ASUserDetectionStatusLikelyReal - The user appears to be a real person.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/likelyReal
	ASUserDetectionStatusLikelyReal ASUserDetectionStatus = 0
	// ASUserDetectionStatusUnknown - The system hasn’t determined whether the user might be a real person.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/unknown
	ASUserDetectionStatusUnknown ASUserDetectionStatus = 0
	// ASUserDetectionStatusUnsupported - The system can’t determine this user’s status as a real person.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASUserDetectionStatus/unsupported
	ASUserDetectionStatusUnsupported ASUserDetectionStatus = 0
)

/* debug [enums.gen.go]: Processing enum ASWebAuthenticationSessionErrorCode (3 cases) */
// ASWebAuthenticationSessionErrorCode - The error code for a web authentication session error.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code
type ASWebAuthenticationSessionErrorCode uint

const (
	// ASWebAuthenticationSessionErrorCodeCanceledLogin - The login has been canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/canceledLogin
	ASWebAuthenticationSessionErrorCodeCanceledLogin ASWebAuthenticationSessionErrorCode = 0
	// ASWebAuthenticationSessionErrorCodePresentationContextInvalid - The context was invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/presentationContextInvalid
	ASWebAuthenticationSessionErrorCodePresentationContextInvalid ASWebAuthenticationSessionErrorCode = 0
	// ASWebAuthenticationSessionErrorCodePresentationContextNotProvided - A context wasn’t provided.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASWebAuthenticationSessionError/Code/presentationContextNotProvided
	ASWebAuthenticationSessionErrorCodePresentationContextNotProvided ASWebAuthenticationSessionErrorCode = 0
)


