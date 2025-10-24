// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

/* debug [enums.gen.go]: Generating 48 enums for Security */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum __CE_DataType (21 cases) */
// __CE_DataType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DataType-c.enum
type __CE_DataType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_AuthorityInfoAccess
	DT_AuthorityInfoAccess __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_AuthorityKeyID
	DT_AuthorityKeyID __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_BasicConstraints
	DT_BasicConstraints __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_CertPolicies
	DT_CertPolicies __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_CrlDistributionPoints
	DT_CrlDistributionPoints __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_CrlNumber
	DT_CrlNumber __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_CrlReason
	DT_CrlReason __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_DeltaCrl
	DT_DeltaCrl __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_ExtendedKeyUsage
	DT_ExtendedKeyUsage __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_InhibitAnyPolicy
	DT_InhibitAnyPolicy __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_IssuerAltName
	DT_IssuerAltName __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_IssuingDistributionPoint
	DT_IssuingDistributionPoint __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_KeyUsage
	DT_KeyUsage __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_NameConstraints
	DT_NameConstraints __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_NetscapeCertType
	DT_NetscapeCertType __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_Other
	DT_Other __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_PolicyConstraints
	DT_PolicyConstraints __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_PolicyMappings
	DT_PolicyMappings __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_QC_Statements
	DT_QC_Statements __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_SubjectAltName
	DT_SubjectAltName __CE_DataType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/DT_SubjectKeyID
	DT_SubjectKeyID __CE_DataType = 0
)

/* debug [enums.gen.go]: Processing enum __CE_GeneralNameType (9 cases) */
// __CE_GeneralNameType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralNameType-c.enum
type __CE_GeneralNameType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_DirectoryName
	GNT_DirectoryName __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_DNSName
	GNT_DNSName __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_EdiPartyName
	GNT_EdiPartyName __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_IPAddress
	GNT_IPAddress __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_OtherName
	GNT_OtherName __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_RegisteredID
	GNT_RegisteredID __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_RFC822Name
	GNT_RFC822Name __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_URI
	GNT_URI __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_X400Address
	GNT_X400Address __CE_GeneralNameType = 0
)

/* debug [enums.gen.go]: Processing enum AuthorizationFlags (8 cases) */
// AuthorizationFlags - The flags used to specify authorization options.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags
type AuthorizationFlags uint

const (
	// kAuthorizationFlagDestroyRights - A flag that instructs the Security Server to revoke authorization.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/destroyRights
	kAuthorizationFlagDestroyRights AuthorizationFlags = 0
	// kAuthorizationFlagExtendRights - A flag that permits the Security Server to attempt to grant the rights requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/extendRights
	kAuthorizationFlagExtendRights AuthorizationFlags = 0
	// kAuthorizationFlagInteractionAllowed - A flag that permits user interaction as needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/interactionAllowed
	kAuthorizationFlagInteractionAllowed AuthorizationFlags = 0
	// kAuthorizationFlagDefaults - An empty flag set that you use as a placeholder when you don’t want any of the other flags.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/kAuthorizationFlagDefaults
	kAuthorizationFlagDefaults AuthorizationFlags = 0
	// kAuthorizationFlagNoData - Private flag. Do not use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/noData
	kAuthorizationFlagNoData AuthorizationFlags = 0
	// kAuthorizationFlagPartialRights - A flag that permits the Security Server to grant rights on an individual basis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/partialRights
	kAuthorizationFlagPartialRights AuthorizationFlags = 0
	// kAuthorizationFlagPreAuthorize - A flag that instructs the Security Server to preauthorize the rights requested.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/preAuthorize
	kAuthorizationFlagPreAuthorize AuthorizationFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFlags/skipInternalAuth
	kAuthorizationFlagSkipInternalAuth AuthorizationFlags = 0
)

/* debug [enums.gen.go]: Processing enum CMSCertificateChainMode (5 cases) */
// CMSCertificateChainMode - Constants that can be set to specify what certificates to include in a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode
type CMSCertificateChainMode uint

const (
	// kCMSCertificateChain - Include the signer certificate chain up to but not including the root certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode/chain
	kCMSCertificateChain CMSCertificateChainMode = 0
	// kCMSCertificateChainWithRoot - Include the entire signer certificate chain, including the root certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode/chainWithRoot
	kCMSCertificateChainWithRoot CMSCertificateChainMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode/chainWithRootOrFail
	kCMSCertificateChainWithRootOrFail CMSCertificateChainMode = 0
	// kCMSCertificateNone - Don’t include any certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode/none
	kCMSCertificateNone CMSCertificateChainMode = 0
	// kCMSCertificateSignerOnly - Only include signer certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode/signerOnly
	kCMSCertificateSignerOnly CMSCertificateChainMode = 0
)

/* debug [enums.gen.go]: Processing enum CMSSignedAttributes (8 cases) */
// CMSSignedAttributes - Optional attributes you can add to a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes
type CMSSignedAttributes uint

const (
	// kCMSAttrAppleCodesigningHashAgility - Include Apple codesigning hash agility.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrAppleCodesigningHashAgility
	kCMSAttrAppleCodesigningHashAgility CMSSignedAttributes = 0
	// kCMSAttrAppleCodesigningHashAgilityV2 - Include Apple codesigning hash agility, version 2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrAppleCodesigningHashAgilityV2
	kCMSAttrAppleCodesigningHashAgilityV2 CMSSignedAttributes = 0
	// kCMSAttrAppleExpirationTime - Include the expiration time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrAppleExpirationTime
	kCMSAttrAppleExpirationTime CMSSignedAttributes = 0
	// kCMSAttrSigningTime - Include the signing time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrSigningTime
	kCMSAttrSigningTime CMSSignedAttributes = 0
	// kCMSAttrSmimeCapabilities - Identify signature, encryption, and digest algorithms supported by the encoder.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrSmimeCapabilities
	kCMSAttrSmimeCapabilities CMSSignedAttributes = 0
	// kCMSAttrSmimeEncryptionKeyPrefs - Indicate that the signing certificate included with the message is the preferred one for S/MIME encryption.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrSmimeEncryptionKeyPrefs
	kCMSAttrSmimeEncryptionKeyPrefs CMSSignedAttributes = 0
	// kCMSAttrSmimeMSEncryptionKeyPrefs - Indicate that the signing certificate included with the message is the preferred one for S/MIME encryption, but using an attribute object identifier (OID) preferred by Microsoft.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/attrSmimeMSEncryptionKeyPrefs
	kCMSAttrSmimeMSEncryptionKeyPrefs CMSSignedAttributes = 0
	// kCMSAttrNone - No attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes/kCMSAttrNone
	kCMSAttrNone CMSSignedAttributes = 0
)

/* debug [enums.gen.go]: Processing enum CMSSignerStatus (6 cases) */
// CMSSignerStatus - The constants that indicate the status of the signature and signer information in a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus
type CMSSignerStatus uint

const (
	// kCMSSignerInvalidCert - The message was signed but the signer’s certificate could not be verified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/invalidCert
	kCMSSignerInvalidCert CMSSignerStatus = 0
	// kCMSSignerInvalidIndex - The specified value for the signer index (  parameter) is greater than the number of signers of the message minus one ( ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/invalidIndex
	kCMSSignerInvalidIndex CMSSignerStatus = 0
	// kCMSSignerInvalidSignature - The message was signed but the signature is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/invalidSignature
	kCMSSignerInvalidSignature CMSSignerStatus = 0
	// kCMSSignerNeedsDetachedContent - The message was signed but has detached content. You must call the   function before ascertaining the signature status.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/needsDetachedContent
	kCMSSignerNeedsDetachedContent CMSSignerStatus = 0
	// kCMSSignerUnsigned - The message was not signed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/unsigned
	kCMSSignerUnsigned CMSSignerStatus = 0
	// kCMSSignerValid - The message was signed and both the signature and the signer certificate have been verified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus/valid
	kCMSSignerValid CMSSignerStatus = 0
)

/* debug [enums.gen.go]: Processing enum cssm_appledl_open_parameters_mask (1 cases) */
// cssm_appledl_open_parameters_mask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters_mask
type cssm_appledl_open_parameters_mask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/kCSSM_APPLEDL_MASK_MODE
	kCSSM_APPLEDL_MASK_MODE cssm_appledl_open_parameters_mask = 0
)

/* debug [enums.gen.go]: Processing enum SecAccessControlCreateFlags (12 cases) */
// SecAccessControlCreateFlags - Access control constants that dictate how a keychain item may be used.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags
type SecAccessControlCreateFlags uint

const (
	// kSecAccessControlAnd - Indicates that all constraints must be satisfied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/and
	kSecAccessControlAnd SecAccessControlCreateFlags = 0
	// kSecAccessControlApplicationPassword - Option to use an application-provided password for data encryption key generation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/applicationPassword
	kSecAccessControlApplicationPassword SecAccessControlCreateFlags = 0
	// kSecAccessControlBiometryAny - Constraint to access an item with Touch ID for any enrolled fingers, or Face ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/biometryAny
	kSecAccessControlBiometryAny SecAccessControlCreateFlags = 0
	// kSecAccessControlBiometryCurrentSet - Constraint to access an item with Touch ID for currently enrolled fingers, or from Face ID with the currently enrolled user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/biometryCurrentSet
	kSecAccessControlBiometryCurrentSet SecAccessControlCreateFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/companion
	kSecAccessControlCompanion SecAccessControlCreateFlags = 0
	// kSecAccessControlDevicePasscode - Constraint to access an item with a passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/devicePasscode
	kSecAccessControlDevicePasscode SecAccessControlCreateFlags = 0
	// kSecAccessControlOr - Indicates that at least one constraint must be satisfied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/or
	kSecAccessControlOr SecAccessControlCreateFlags = 0
	// kSecAccessControlPrivateKeyUsage - Enable a private key to be used in signing a block of data or verifying a signed block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/privateKeyUsage
	kSecAccessControlPrivateKeyUsage SecAccessControlCreateFlags = 0
	// kSecAccessControlTouchIDAny - Constraint to access an item with Touch ID for any enrolled fingers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/touchIDAny
	kSecAccessControlTouchIDAny SecAccessControlCreateFlags = 0
	// kSecAccessControlTouchIDCurrentSet - Constraint to access an item with Touch ID for currently enrolled fingers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/touchIDCurrentSet
	kSecAccessControlTouchIDCurrentSet SecAccessControlCreateFlags = 0
	// kSecAccessControlUserPresence - Constraint to access an item with either biometry or passcode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/userPresence
	kSecAccessControlUserPresence SecAccessControlCreateFlags = 0
	// kSecAccessControlWatch - Constraint to access an item with a watch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags/watch
	kSecAccessControlWatch SecAccessControlCreateFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecAuthenticationType (9 cases) */
// SecAuthenticationType - The authentication type to use for an Internet password.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType
type SecAuthenticationType uint

const (
	// kSecAuthenticationTypeAny - Specifies that any authentication type is acceptable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/any
	kSecAuthenticationTypeAny SecAuthenticationType = 0
	// kSecAuthenticationTypeDefault - Specifies the default authentication type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/default
	kSecAuthenticationTypeDefault SecAuthenticationType = 0
	// kSecAuthenticationTypeDPA - Specifies Distributed Password authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/DPA
	kSecAuthenticationTypeDPA SecAuthenticationType = 0
	// kSecAuthenticationTypeHTMLForm - Specifies HTML form based authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/htmlForm
	kSecAuthenticationTypeHTMLForm SecAuthenticationType = 0
	// kSecAuthenticationTypeHTTPBasic - Specifies HTTP Basic authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/httpBasic
	kSecAuthenticationTypeHTTPBasic SecAuthenticationType = 0
	// kSecAuthenticationTypeHTTPDigest - Specifies HTTP Digest Access authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/httpDigest
	kSecAuthenticationTypeHTTPDigest SecAuthenticationType = 0
	// kSecAuthenticationTypeMSN - Specifies Microsoft Network default authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/MSN
	kSecAuthenticationTypeMSN SecAuthenticationType = 0
	// kSecAuthenticationTypeNTLM - Specifies Windows NT LAN Manager authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/NTLM
	kSecAuthenticationTypeNTLM SecAuthenticationType = 0
	// kSecAuthenticationTypeRPA - Specifies Remote Password authentication.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType/RPA
	kSecAuthenticationTypeRPA SecAuthenticationType = 0
)

/* debug [enums.gen.go]: Processing enum SecCodeSignatureFlags (10 cases) */
// SecCodeSignatureFlags - Specify option flags that can be embedded in a code signature during signing and that govern the use of the signature.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags
type SecCodeSignatureFlags uint

const (
	// kSecCodeSignatureAdhoc - Must be used without a signing identity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/adhoc
	kSecCodeSignatureAdhoc SecCodeSignatureFlags = 0
	// kSecCodeSignatureEnforcement - Enforce code signing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/enforcement
	kSecCodeSignatureEnforcement SecCodeSignatureFlags = 0
	// kSecCodeSignatureForceExpiration - Always set the   flag when validating the code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/forceExpiration
	kSecCodeSignatureForceExpiration SecCodeSignatureFlags = 0
	// kSecCodeSignatureForceHard - Always set the   status flag on launch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/forceHard
	kSecCodeSignatureForceHard SecCodeSignatureFlags = 0
	// kSecCodeSignatureForceKill - Always set the termination status flag on launch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/forceKill
	kSecCodeSignatureForceKill SecCodeSignatureFlags = 0
	// kSecCodeSignatureHost - May host guest code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/host
	kSecCodeSignatureHost SecCodeSignatureFlags = 0
	// kSecCodeSignatureLibraryValidation - Require library validation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/libraryValidation
	kSecCodeSignatureLibraryValidation SecCodeSignatureFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/linkerSigned
	kSecCodeSignatureLinkerSigned SecCodeSignatureFlags = 0
	// kSecCodeSignatureRestrict - Restrict dyld loading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/restrict
	kSecCodeSignatureRestrict SecCodeSignatureFlags = 0
	// kSecCodeSignatureRuntime - Apply runtime hardening policies as required by the hardened runtime version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags/runtime
	kSecCodeSignatureRuntime SecCodeSignatureFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecCodeStatus (5 cases) */
// SecCodeStatus - Operational flags attached by code signing services to running code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus
type SecCodeStatus uint

const (
	// kSecCodeStatusDebugged - The code has been debugged by another process that was allowed to do so.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus/debugged
	kSecCodeStatusDebugged SecCodeStatus = 0
	// kSecCodeStatusHard - The code prefers to be denied access to resources if gaining access would invalidate it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus/hard
	kSecCodeStatusHard SecCodeStatus = 0
	// kSecCodeStatusKill - The code wants to be terminated if it ever loses its validity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus/kill
	kSecCodeStatusKill SecCodeStatus = 0
	// kSecCodeStatusPlatform - The code ships with the operating system and is signed by Apple.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus/platform
	kSecCodeStatusPlatform SecCodeStatus = 0
	// kSecCodeStatusValid - The code is dynamically valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus/valid
	kSecCodeStatusValid SecCodeStatus = 0
)

/* debug [enums.gen.go]: Processing enum SecCredentialType (3 cases) */
// SecCredentialType - The credential type to be returned by 
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCredentialType
type SecCredentialType uint

const (
	// kSecCredentialTypeDefault - The default setting for determining whether to present UI is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCredentialType/default
	kSecCredentialTypeDefault SecCredentialType = 0
	// kSecCredentialTypeNoUI - Keychain operations on keys that have this credential are not allowed to present UI, and will fail if UI is required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCredentialType/noUI
	kSecCredentialTypeNoUI SecCredentialType = 0
	// kSecCredentialTypeWithUI - Keychain operations on keys that have this credential are allowed to present UI if required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCredentialType/withUI
	kSecCredentialTypeWithUI SecCredentialType = 0
)

/* debug [enums.gen.go]: Processing enum SecCSDigestAlgorithm (6 cases) */
// SecCSDigestAlgorithm - The list of digest algorithms available for code signatures.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm
type SecCSDigestAlgorithm uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureHashSHA1
	kSecCodeSignatureHashSHA1 SecCSDigestAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureHashSHA256
	kSecCodeSignatureHashSHA256 SecCSDigestAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureHashSHA256Truncated
	kSecCodeSignatureHashSHA256Truncated SecCSDigestAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureHashSHA384
	kSecCodeSignatureHashSHA384 SecCSDigestAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureHashSHA512
	kSecCodeSignatureHashSHA512 SecCSDigestAlgorithm = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm/codeSignatureNoHash
	kSecCodeSignatureNoHash SecCSDigestAlgorithm = 0
)

/* debug [enums.gen.go]: Processing enum SecCSFlags (10 cases) */
// SecCSFlags - Values that can be used in the 
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags
type SecCSFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/applyEmbeddedPolicy
	kSecCSApplyEmbeddedPolicy SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/checkTrustedAnchors
	kSecCSCheckTrustedAnchors SecCSFlags = 0
	// kSecCSConsiderExpiration - Consider expired certificates invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/considerExpiration
	kSecCSConsiderExpiration SecCSFlags = 0
	// kSecCSEnforceRevocationChecks - Forces checking of certificates against revocation lists or OCSP (online certificate status protocol) regardless of preference settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/enforceRevocationChecks
	kSecCSEnforceRevocationChecks SecCSFlags = 0
	// kSecCSDefaultFlags - No flags (use the default behavior).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/kSecCSDefaultFlags
	kSecCSDefaultFlags SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/matchGuestRequirementInKernel
	kSecCSMatchGuestRequirementInKernel SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/noNetworkAccess
	kSecCSNoNetworkAccess SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/quickCheck
	kSecCSQuickCheck SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/reportProgress
	kSecCSReportProgress SecCSFlags = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/stripDisallowedXattrs
	kSecCSStripDisallowedXattrs SecCSFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecExternalFormat (15 cases) */
// SecExternalFormat - The external format of a keychain item.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat
type SecExternalFormat uint

const (
	// kSecFormatBSAFE - Format for asymmetric keys. BSAFE is a standard from RSA Security for encryption, digital signatures, and privacy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatBSAFE
	kSecFormatBSAFE SecExternalFormat = 0
	// kSecFormatNetscapeCertSequence - Set of certificates in the Netscape Certificate Sequence format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatNetscapeCertSequence
	kSecFormatNetscapeCertSequence SecExternalFormat = 0
	// kSecFormatOpenSSL - Format for asymmetric (public/private) keys. OpenSSL is an open source toolkit for Secure Sockets Layer (SSL) and Transport Layer Security (TLS). Also known as X.509 for public keys.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatOpenSSL
	kSecFormatOpenSSL SecExternalFormat = 0
	// kSecFormatPEMSequence - Sequence of certificates and keys with PEM armor. PEM armor refers to a way of expressing binary data as an ASCII string so that it can be transferred over text-only channels such as email. This is the default format for multiple items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatPEMSequence
	kSecFormatPEMSequence SecExternalFormat = 0
	// kSecFormatPKCS12 - Set of certificates and private keys. PKCS12 is the Personal Information Exchange Syntax from RSA Security, Inc.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatPKCS12
	kSecFormatPKCS12 SecExternalFormat = 0
	// kSecFormatPKCS7 - Sequence of certificates, no PEM armor. PKCS7 is the Cryptographic Message Syntax Standard from RSA Security, Inc.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatPKCS7
	kSecFormatPKCS7 SecExternalFormat = 0
	// kSecFormatRawKey - Format for symmetric keys. Raw, unformatted key bits. This is the default for symmetric keys.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatRawKey
	kSecFormatRawKey SecExternalFormat = 0
	// kSecFormatSSH - OpenSSH 1 format for asymmetric (public/private) keys. OpenSSH is an OpenBSD implementation of the Secure Shell (SSH) protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatSSH
	kSecFormatSSH SecExternalFormat = 0
	// kSecFormatSSHv2 - OpenSSH 2 format for public keys. OpenSSH version 2 private keys are in format   or  . OpenSSH is an OpenBSD implementation of the Secure Shell (SSH) protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatSSHv2
	kSecFormatSSHv2 SecExternalFormat = 0
	// kSecFormatUnknown - When importing, indicates the format is unknown. When exporting, use the default format for the item. For asymmetric keys, the default is  . For symmetric keys, the default is  . For certificates, the default is  . For multiple items, the default is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatUnknown
	kSecFormatUnknown SecExternalFormat = 0
	// kSecFormatWrappedLSH - Not supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatWrappedLSH
	kSecFormatWrappedLSH SecExternalFormat = 0
	// kSecFormatWrappedOpenSSL - Format for wrapped symmetric and private keys. OpenSSL is an open-source toolkit for Secure Sockets Layer (SSL) and Transport Layer Security (TLS).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatWrappedOpenSSL
	kSecFormatWrappedOpenSSL SecExternalFormat = 0
	// kSecFormatWrappedPKCS8 - Format for wrapped symmetric and private keys. PKCS8 is the Private-Key Information Syntax Standard from RSA Security.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatWrappedPKCS8
	kSecFormatWrappedPKCS8 SecExternalFormat = 0
	// kSecFormatWrappedSSH - OpenSSH 1 format for wrapped symmetric and private keys.  OpenSSH is an OpenBSD implementation of the Secure Shell (SSH) protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatWrappedSSH
	kSecFormatWrappedSSH SecExternalFormat = 0
	// kSecFormatX509Cert - Format for certificates. DER (distinguished encoding rules) encoded. X.509 is a standard for digital certificates from the International Telecommunication Union (ITU). This is the default for certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat/formatX509Cert
	kSecFormatX509Cert SecExternalFormat = 0
)

/* debug [enums.gen.go]: Processing enum SecExternalItemType (6 cases) */
// SecExternalItemType - The import item type.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType
type SecExternalItemType uint

const (
	// kSecItemTypeAggregate - Indicates a set of certificates or certificates and private keys.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypeAggregate
	kSecItemTypeAggregate SecExternalItemType = 0
	// kSecItemTypeCertificate - Indicates a certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypeCertificate
	kSecItemTypeCertificate SecExternalItemType = 0
	// kSecItemTypePrivateKey - Indicates a private key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypePrivateKey
	kSecItemTypePrivateKey SecExternalItemType = 0
	// kSecItemTypePublicKey - Indicates a public key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypePublicKey
	kSecItemTypePublicKey SecExternalItemType = 0
	// kSecItemTypeSessionKey - Indicates a session key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypeSessionKey
	kSecItemTypeSessionKey SecExternalItemType = 0
	// kSecItemTypeUnknown - Indicates that the caller does not know the type of information being imported or exported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType/itemTypeUnknown
	kSecItemTypeUnknown SecExternalItemType = 0
)

/* debug [enums.gen.go]: Processing enum SecItemAttr (28 cases) */
// SecItemAttr - Specifies a keychain item’s attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr
type SecItemAttr uint

const (
	// kSecAccountItemAttr - Identifies the account attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/accountItemAttr
	kSecAccountItemAttr SecItemAttr = 0
	// kSecAddressItemAttr - Identifies the address attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/addressItemAttr
	kSecAddressItemAttr SecItemAttr = 0
	// kSecAlias - Indicates an alias.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/alias
	kSecAlias SecItemAttr = 0
	// kSecAuthenticationTypeItemAttr - Identifies the authentication type attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/authenticationTypeItemAttr
	kSecAuthenticationTypeItemAttr SecItemAttr = 0
	// kSecCertificateEncoding - Indicates a   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/certificateEncoding
	kSecCertificateEncoding SecItemAttr = 0
	// kSecCertificateType - Indicates a   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/certificateType
	kSecCertificateType SecItemAttr = 0
	// kSecCommentItemAttr - Identifies the comment attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/commentItemAttr
	kSecCommentItemAttr SecItemAttr = 0
	// kSecCreationDateItemAttr - Identifies the creation date attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/creationDateItemAttr
	kSecCreationDateItemAttr SecItemAttr = 0
	// kSecCreatorItemAttr - Identifies the creator attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/creatorItemAttr
	kSecCreatorItemAttr SecItemAttr = 0
	// kSecCrlEncoding - Indicates a   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/crlEncoding
	kSecCrlEncoding SecItemAttr = 0
	// kSecCrlType - Indicates a   type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/crlType
	kSecCrlType SecItemAttr = 0
	// kSecCustomIconItemAttr - Identifies the custom icon attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/customIconItemAttr
	kSecCustomIconItemAttr SecItemAttr = 0
	// kSecDescriptionItemAttr - Identifies the description attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/descriptionItemAttr
	kSecDescriptionItemAttr SecItemAttr = 0
	// kSecGenericItemAttr - Identifies the generic attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/genericItemAttr
	kSecGenericItemAttr SecItemAttr = 0
	// kSecInvisibleItemAttr - Identifies the invisible attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/invisibleItemAttr
	kSecInvisibleItemAttr SecItemAttr = 0
	// kSecLabelItemAttr - Identifies the label attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/labelItemAttr
	kSecLabelItemAttr SecItemAttr = 0
	// kSecModDateItemAttr - Identifies the modification date attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/modDateItemAttr
	kSecModDateItemAttr SecItemAttr = 0
	// kSecNegativeItemAttr - Identifies the negative attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/negativeItemAttr
	kSecNegativeItemAttr SecItemAttr = 0
	// kSecPathItemAttr - Identifies the path attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/pathItemAttr
	kSecPathItemAttr SecItemAttr = 0
	// kSecPortItemAttr - Identifies the port attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/portItemAttr
	kSecPortItemAttr SecItemAttr = 0
	// kSecProtocolItemAttr - Identifies the protocol attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/protocolItemAttr
	kSecProtocolItemAttr SecItemAttr = 0
	// kSecScriptCodeItemAttr - Identifies the script code attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/scriptCodeItemAttr
	kSecScriptCodeItemAttr SecItemAttr = 0
	// kSecSecurityDomainItemAttr - Identifies the security domain attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/securityDomainItemAttr
	kSecSecurityDomainItemAttr SecItemAttr = 0
	// kSecServerItemAttr - Identifies the server attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/serverItemAttr
	kSecServerItemAttr SecItemAttr = 0
	// kSecServiceItemAttr - Identifies the service attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/serviceItemAttr
	kSecServiceItemAttr SecItemAttr = 0
	// kSecSignatureItemAttr - Identifies the server signature attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/signatureItemAttr
	kSecSignatureItemAttr SecItemAttr = 0
	// kSecTypeItemAttr - Identifies the type attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/typeItemAttr
	kSecTypeItemAttr SecItemAttr = 0
	// kSecVolumeItemAttr - Identifies the volume attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr/volumeItemAttr
	kSecVolumeItemAttr SecItemAttr = 0
)

/* debug [enums.gen.go]: Processing enum SecItemClass (7 cases) */
// SecItemClass - Specifies a keychain item’s class code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass
type SecItemClass uint

const (
	// kSecCertificateItemClass - Indicates that the item is an X509 certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/certificateItemClass
	kSecCertificateItemClass SecItemClass = 0
	// kSecGenericPasswordItemClass - Indicates that the item is a generic password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/genericPasswordItemClass
	kSecGenericPasswordItemClass SecItemClass = 0
	// kSecInternetPasswordItemClass - Indicates that the item is an Internet password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/internetPasswordItemClass
	kSecInternetPasswordItemClass SecItemClass = 0
	// kSecAppleSharePasswordItemClass - Indicates that the item is an AppleShare password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/kSecAppleSharePasswordItemClass
	kSecAppleSharePasswordItemClass SecItemClass = 0
	// kSecPrivateKeyItemClass - Indicates that the item is a private key of a public-private pair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/privateKeyItemClass
	kSecPrivateKeyItemClass SecItemClass = 0
	// kSecPublicKeyItemClass - Indicates that the item is a public key of a public-private pair.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/publicKeyItemClass
	kSecPublicKeyItemClass SecItemClass = 0
	// kSecSymmetricKeyItemClass - Indicates that the item is a private key used for symmetric-key encryption.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/symmetricKeyItemClass
	kSecSymmetricKeyItemClass SecItemClass = 0
)

/* debug [enums.gen.go]: Processing enum SecItemImportExportFlags (1 cases) */
// SecItemImportExportFlags - The import and export function flags.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportFlags
type SecItemImportExportFlags uint

const (
	// kSecItemPemArmour - A flag that indicates the exported data should have PEM armor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportFlags/pemArmour
	kSecItemPemArmour SecItemImportExportFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecKeychainEvent (10 cases) */
// SecKeychainEvent - The list of keychain events that can trigger a callback.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent
type SecKeychainEvent uint

const (
	// kSecAddEvent - Indicates an item was added to a keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/addEvent
	kSecAddEvent SecKeychainEvent = 0
	// kSecDataAccessEvent - Indicates a process has accessed a keychain item’s data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/dataAccessEvent
	kSecDataAccessEvent SecKeychainEvent = 0
	// kSecDefaultChangedEvent - Indicates that a different keychain was specified as the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/defaultChangedEvent
	kSecDefaultChangedEvent SecKeychainEvent = 0
	// kSecDeleteEvent - Indicates an item was deleted from a keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/deleteEvent
	kSecDeleteEvent SecKeychainEvent = 0
	// kSecKeychainListChangedEvent - Indicates the list of keychains has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/keychainListChangedEvent
	kSecKeychainListChangedEvent SecKeychainEvent = 0
	// kSecLockEvent - Indicates a keychain was locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/lockEvent
	kSecLockEvent SecKeychainEvent = 0
	// kSecPasswordChangedEvent - Indicates the keychain password was changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/passwordChangedEvent
	kSecPasswordChangedEvent SecKeychainEvent = 0
	// kSecTrustSettingsChangedEvent - Indicates trust settings have changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/trustSettingsChangedEvent
	kSecTrustSettingsChangedEvent SecKeychainEvent = 0
	// kSecUnlockEvent - Indicates a keychain was successfully unlocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/unlockEvent
	kSecUnlockEvent SecKeychainEvent = 0
	// kSecUpdateEvent - Indicates a keychain item was updated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEvent/updateEvent
	kSecUpdateEvent SecKeychainEvent = 0
)

/* debug [enums.gen.go]: Processing enum SecKeychainEventMask (11 cases) */
// SecKeychainEventMask - Bit masks corresponding to the events that can trigger a keychain callback.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask
type SecKeychainEventMask uint

const (
	// kSecAddEventMask - If the bit specified by this mask is set, your callback function is invoked when an item is added to a keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/addEventMask
	kSecAddEventMask SecKeychainEventMask = 0
	// kSecDataAccessEventMask - If the bit specified by this mask is set, your callback function is invoked when a process accesses a keychain item’s data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/dataAccessEventMask
	kSecDataAccessEventMask SecKeychainEventMask = 0
	// kSecDefaultChangedEventMask - If the bit specified by this mask is set, your callback function is invoked when a different keychain is specified as the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/defaultChangedEventMask
	kSecDefaultChangedEventMask SecKeychainEventMask = 0
	// kSecDeleteEventMask - If the bit specified by this mask is set, your callback function is invoked when an item is deleted from a keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/deleteEventMask
	kSecDeleteEventMask SecKeychainEventMask = 0
	// kSecEveryEventMask - If all the bits are set, your callback function is invoked whenever any event occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/everyEventMask
	kSecEveryEventMask SecKeychainEventMask = 0
	// kSecKeychainListChangedMask - If the bit specified by this mask is set, your callback function is invoked when a keychain list is changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/keychainListChangedMask
	kSecKeychainListChangedMask SecKeychainEventMask = 0
	// kSecLockEventMask - If the bit specified by this mask is set, your callback function is invoked when a keychain is locked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/lockEventMask
	kSecLockEventMask SecKeychainEventMask = 0
	// kSecPasswordChangedEventMask - If the bit specified by this mask is set, your callback function is invoked when the keychain password is changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/passwordChangedEventMask
	kSecPasswordChangedEventMask SecKeychainEventMask = 0
	// kSecTrustSettingsChangedEventMask - If the bit specified by this mask is set, your callback function is invoked when there is a change in certificate trust settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/trustSettingsChangedEventMask
	kSecTrustSettingsChangedEventMask SecKeychainEventMask = 0
	// kSecUnlockEventMask - If the bit specified by this mask is set, your callback function is invoked when a keychain is unlocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/unlockEventMask
	kSecUnlockEventMask SecKeychainEventMask = 0
	// kSecUpdateEventMask - If the bit specified by this mask is set, your callback function is invoked when a keychain item is updated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainEventMask/updateEventMask
	kSecUpdateEventMask SecKeychainEventMask = 0
)

/* debug [enums.gen.go]: Processing enum SecKeychainPromptSelector (5 cases) */
// SecKeychainPromptSelector - Bits that define when a keychain should require a passphrase.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector
type SecKeychainPromptSelector uint

const (
	// kSecKeychainPromptInvalid - Indicates that a passphrase should be required when an application with an invalid signature attempts to use the keychain, overriding the system default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector/invalid
	kSecKeychainPromptInvalid SecKeychainPromptSelector = 0
	// kSecKeychainPromptInvalidAct - Indicates that a passphrase should be required when an application with an invalid signature attempts to use the keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector/invalidAct
	kSecKeychainPromptInvalidAct SecKeychainPromptSelector = 0
	// kSecKeychainPromptRequirePassphase - Indicates that a passphrase should be required for every access.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector/requirePassphase
	kSecKeychainPromptRequirePassphase SecKeychainPromptSelector = 0
	// kSecKeychainPromptUnsigned - Indicates that a passphrase should be required when an unsigned application attempts to use the keychain, overriding the system default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector/unsigned
	kSecKeychainPromptUnsigned SecKeychainPromptSelector = 0
	// kSecKeychainPromptUnsignedAct - Indicates that a passphrase should be required when an unsigned application attempts to use the keychain.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainPromptSelector/unsignedAct
	kSecKeychainPromptUnsignedAct SecKeychainPromptSelector = 0
)

/* debug [enums.gen.go]: Processing enum SecKeyImportExportFlags (3 cases) */
// SecKeyImportExportFlags - The import/export parameter structure flags.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags
type SecKeyImportExportFlags uint

const (
	// kSecKeyImportOnlyOne - A flag that you set to prevent importing more than one private key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags/importOnlyOne
	kSecKeyImportOnlyOne SecKeyImportExportFlags = 0
	// kSecKeyNoAccessControl - A flag that indicates imported private keys have no access object attached to them.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags/noAccessControl
	kSecKeyNoAccessControl SecKeyImportExportFlags = 0
	// kSecKeySecurePassphrase - A flag that indicates the user should be prompted for a passphrase on import or export.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags/securePassphrase
	kSecKeySecurePassphrase SecKeyImportExportFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecKeyOperationType (5 cases) */
// SecKeyOperationType - The types of operations that you can use a cryptographic key to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType
type SecKeyOperationType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType/decrypt
	kSecKeyOperationTypeDecrypt SecKeyOperationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType/encrypt
	kSecKeyOperationTypeEncrypt SecKeyOperationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType/keyExchange
	kSecKeyOperationTypeKeyExchange SecKeyOperationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType/sign
	kSecKeyOperationTypeSign SecKeyOperationType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyOperationType/verify
	kSecKeyOperationTypeVerify SecKeyOperationType = 0
)

/* debug [enums.gen.go]: Processing enum SecKeySizes (11 cases) */
// SecKeySizes - The supported sizes for keys of various common types.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes
type SecKeySizes uint

const (
	// kSec3DES192 - 192-bit DES.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/sec3DES192
	kSec3DES192 SecKeySizes = 0
	// kSecAES128 - 128-bit AES.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secAES128
	kSecAES128 SecKeySizes = 0
	// kSecAES192 - 192-bit AES.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secAES192
	kSecAES192 SecKeySizes = 0
	// kSecAES256 - 256-bit AES.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secAES256
	kSecAES256 SecKeySizes = 0
	// kSecDefaultKeySize - The default key size for the specified type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secDefaultKeySize
	kSecDefaultKeySize SecKeySizes = 0
	// kSecp192r1 - 192-bit ECC Keys for Suite-B from RFC 4492 section 5.1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secp192r1
	kSecp192r1 SecKeySizes = 0
	// kSecp256r1 - 256-bit ECC Keys for Suite-B from RFC 4492 section 5.1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secp256r1
	kSecp256r1 SecKeySizes = 0
	// kSecp384r1 - 384-bit ECC Keys for Suite-B from RFC 4492 section 5.1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secp384r1
	kSecp384r1 SecKeySizes = 0
	// kSecp521r1 - 521-bit ECC Keys for Suite-B from RFC 4492 section 5.1.1.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secp521r1
	kSecp521r1 SecKeySizes = 0
	// kSecRSAMax - 4096 bits is the maximum size for an RSA key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secRSAMax
	kSecRSAMax SecKeySizes = 0
	// kSecRSAMin - 1024 bits is the minimum size for an RSA key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeySizes/secRSAMin
	kSecRSAMin SecKeySizes = 0
)

/* debug [enums.gen.go]: Processing enum SecKeyUsage (13 cases) */
// SecKeyUsage - The flags that indicate key usage in the 
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage
type SecKeyUsage uint

const (
	// kSecKeyUsageAll - All flags set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/all
	kSecKeyUsageAll SecKeyUsage = 0
	// kSecKeyUsageContentCommitment - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/contentCommitment
	kSecKeyUsageContentCommitment SecKeyUsage = 0
	// kSecKeyUsageCritical - The KeyUsage extension is marked critical.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/critical
	kSecKeyUsageCritical SecKeyUsage = 0
	// kSecKeyUsageCRLSign - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/crlSign
	kSecKeyUsageCRLSign SecKeyUsage = 0
	// kSecKeyUsageDataEncipherment - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/dataEncipherment
	kSecKeyUsageDataEncipherment SecKeyUsage = 0
	// kSecKeyUsageDecipherOnly - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/decipherOnly
	kSecKeyUsageDecipherOnly SecKeyUsage = 0
	// kSecKeyUsageDigitalSignature - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/digitalSignature
	kSecKeyUsageDigitalSignature SecKeyUsage = 0
	// kSecKeyUsageEncipherOnly - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/encipherOnly
	kSecKeyUsageEncipherOnly SecKeyUsage = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/kSecKeyUsageUnspecified
	kSecKeyUsageUnspecified SecKeyUsage = 0
	// kSecKeyUsageKeyAgreement - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/keyAgreement
	kSecKeyUsageKeyAgreement SecKeyUsage = 0
	// kSecKeyUsageKeyCertSign - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/keyCertSign
	kSecKeyUsageKeyCertSign SecKeyUsage = 0
	// kSecKeyUsageKeyEncipherment - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/keyEncipherment
	kSecKeyUsageKeyEncipherment SecKeyUsage = 0
	// kSecKeyUsageNonRepudiation - The   bit is set in KeyUsage extension.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUsage/nonRepudiation
	kSecKeyUsageNonRepudiation SecKeyUsage = 0
)

/* debug [enums.gen.go]: Processing enum SecPadding (11 cases) */
// SecPadding - The types of padding to use when you create or verify a digital signature.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding
type SecPadding uint

const (
	// kSecPaddingNone - No padding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/kSecPaddingNone
	kSecPaddingNone SecPadding = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/OAEP
	kSecPaddingOAEP SecPadding = 0
	// kSecPaddingPKCS1 - PKCS1 padding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1
	kSecPaddingPKCS1 SecPadding = 0
	// kSecPaddingPKCS1MD2 - Data to be signed is an MD2 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1MD2
	kSecPaddingPKCS1MD2 SecPadding = 0
	// kSecPaddingPKCS1MD5 - Data to be signed is an MD5 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1MD5
	kSecPaddingPKCS1MD5 SecPadding = 0
	// kSecPaddingPKCS1SHA1 - Data to be signed is a SHA1 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1SHA1
	kSecPaddingPKCS1SHA1 SecPadding = 0
	// kSecPaddingPKCS1SHA224 - Data to be signed is a SHA224 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1SHA224
	kSecPaddingPKCS1SHA224 SecPadding = 0
	// kSecPaddingPKCS1SHA256 - Data to be signed is a SHA256 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1SHA256
	kSecPaddingPKCS1SHA256 SecPadding = 0
	// kSecPaddingPKCS1SHA384 - Data to be signed is a SHA384 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1SHA384
	kSecPaddingPKCS1SHA384 SecPadding = 0
	// kSecPaddingPKCS1SHA512 - Data to be signed is a SHA512 hash.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/PKCS1SHA512
	kSecPaddingPKCS1SHA512 SecPadding = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPadding/sigRaw
	kSecPaddingSigRaw SecPadding = 0
)

/* debug [enums.gen.go]: Processing enum SecPreferencesDomain (4 cases) */
// SecPreferencesDomain - The keychain preference domains.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPreferencesDomain
type SecPreferencesDomain uint

const (
	// kSecPreferencesDomainCommon - Indicates the preferences are common to everyone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPreferencesDomain/common
	kSecPreferencesDomainCommon SecPreferencesDomain = 0
	// kSecPreferencesDomainDynamic - Indicates a dynamic search list (typically provided by removable keychains such as smart cards).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPreferencesDomain/dynamic
	kSecPreferencesDomainDynamic SecPreferencesDomain = 0
	// kSecPreferencesDomainSystem - Indicates the system or daemon preference domain preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPreferencesDomain/system
	kSecPreferencesDomainSystem SecPreferencesDomain = 0
	// kSecPreferencesDomainUser - Indicates the user preference domain preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecPreferencesDomain/user
	kSecPreferencesDomainUser SecPreferencesDomain = 0
)

/* debug [enums.gen.go]: Processing enum SecProtocolType (35 cases) */
// SecProtocolType - The protocol type associated with an Internet password.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType
type SecProtocolType uint

const (
	// kSecProtocolTypeAFP - Indicates AFP over TCP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/AFP
	kSecProtocolTypeAFP SecProtocolType = 0
	// kSecProtocolTypeAny - Indicates that any protocol is acceptable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/any
	kSecProtocolTypeAny SecProtocolType = 0
	// kSecProtocolTypeAppleTalk - Indicates AFP over AppleTalk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/appleTalk
	kSecProtocolTypeAppleTalk SecProtocolType = 0
	// kSecProtocolTypeCIFS - Indicates CIFS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/CIFS
	kSecProtocolTypeCIFS SecProtocolType = 0
	// kSecProtocolTypeCVSpserver - Indicates CVS pserver.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/cvSpserver
	kSecProtocolTypeCVSpserver SecProtocolType = 0
	// kSecProtocolTypeDAAP - Indicates DAAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/DAAP
	kSecProtocolTypeDAAP SecProtocolType = 0
	// kSecProtocolTypeEPPC - Indicates Remote Apple Events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/EPPC
	kSecProtocolTypeEPPC SecProtocolType = 0
	// kSecProtocolTypeFTP - Indicates FTP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/FTP
	kSecProtocolTypeFTP SecProtocolType = 0
	// kSecProtocolTypeFTPAccount - Indicates a client side FTP account. The usage of this constant is deprecated as of macOS 10.3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/ftpAccount
	kSecProtocolTypeFTPAccount SecProtocolType = 0
	// kSecProtocolTypeFTPProxy - Indicates FTP proxy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/ftpProxy
	kSecProtocolTypeFTPProxy SecProtocolType = 0
	// kSecProtocolTypeFTPS - Indicates FTP over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/FTPS
	kSecProtocolTypeFTPS SecProtocolType = 0
	// kSecProtocolTypeHTTP - Indicates HTTP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/HTTP
	kSecProtocolTypeHTTP SecProtocolType = 0
	// kSecProtocolTypeHTTPProxy - Indicates HTTP proxy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/httpProxy
	kSecProtocolTypeHTTPProxy SecProtocolType = 0
	// kSecProtocolTypeHTTPS - Indicates HTTP over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/HTTPS
	kSecProtocolTypeHTTPS SecProtocolType = 0
	// kSecProtocolTypeHTTPSProxy - Indicates HTTPS proxy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/httpsProxy
	kSecProtocolTypeHTTPSProxy SecProtocolType = 0
	// kSecProtocolTypeIMAP - Indicates IMAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/IMAP
	kSecProtocolTypeIMAP SecProtocolType = 0
	// kSecProtocolTypeIMAPS - Indicates IMAP4 over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/IMAPS
	kSecProtocolTypeIMAPS SecProtocolType = 0
	// kSecProtocolTypeIPP - Indicates IPP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/IPP
	kSecProtocolTypeIPP SecProtocolType = 0
	// kSecProtocolTypeIRC - Indicates IRC.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/IRC
	kSecProtocolTypeIRC SecProtocolType = 0
	// kSecProtocolTypeIRCS - Indicates IRC over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/IRCS
	kSecProtocolTypeIRCS SecProtocolType = 0
	// kSecProtocolTypeLDAP - Indicates LDAP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/LDAP
	kSecProtocolTypeLDAP SecProtocolType = 0
	// kSecProtocolTypeLDAPS - Indicates LDAP over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/LDAPS
	kSecProtocolTypeLDAPS SecProtocolType = 0
	// kSecProtocolTypeNNTP - Indicates NNTP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/NNTP
	kSecProtocolTypeNNTP SecProtocolType = 0
	// kSecProtocolTypeNNTPS - Indicates NNTP over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/NNTPS
	kSecProtocolTypeNNTPS SecProtocolType = 0
	// kSecProtocolTypePOP3 - Indicates POP3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/POP3
	kSecProtocolTypePOP3 SecProtocolType = 0
	// kSecProtocolTypePOP3S - Indicates POP3 over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/POP3S
	kSecProtocolTypePOP3S SecProtocolType = 0
	// kSecProtocolTypeRTSP - Indicates RTSP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/RTSP
	kSecProtocolTypeRTSP SecProtocolType = 0
	// kSecProtocolTypeRTSPProxy - Indicates RTSP proxy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/rtspProxy
	kSecProtocolTypeRTSPProxy SecProtocolType = 0
	// kSecProtocolTypeSMB - Indicates SMB.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/SMB
	kSecProtocolTypeSMB SecProtocolType = 0
	// kSecProtocolTypeSMTP - Indicates SMTP.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/SMTP
	kSecProtocolTypeSMTP SecProtocolType = 0
	// kSecProtocolTypeSOCKS - Indicates SOCKS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/SOCKS
	kSecProtocolTypeSOCKS SecProtocolType = 0
	// kSecProtocolTypeSSH - Indicates SSH.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/SSH
	kSecProtocolTypeSSH SecProtocolType = 0
	// kSecProtocolTypeSVN - Indicates Subversion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/SVN
	kSecProtocolTypeSVN SecProtocolType = 0
	// kSecProtocolTypeTelnet - Indicates Telnet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/telnet
	kSecProtocolTypeTelnet SecProtocolType = 0
	// kSecProtocolTypeTelnetS - Indicates Telnet over TLS/SSL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType/telnetS
	kSecProtocolTypeTelnetS SecProtocolType = 0
)

/* debug [enums.gen.go]: Processing enum SecRequirementType (7 cases) */
// SecRequirementType - An enumeration indicating different types of internal requirements for code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType
type SecRequirementType uint

const (
	// kSecDesignatedRequirementType - A designated requirement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/designatedRequirementType
	kSecDesignatedRequirementType SecRequirementType = 0
	// kSecGuestRequirementType - What guests this code may run.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/guestRequirementType
	kSecGuestRequirementType SecRequirementType = 0
	// kSecHostRequirementType - What hosts may run this code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/hostRequirementType
	kSecHostRequirementType SecRequirementType = 0
	// kSecInvalidRequirementType - Invalid type of requirement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/invalidRequirementType
	kSecInvalidRequirementType SecRequirementType = 0
	// kSecLibraryRequirementType - What libraries this code may link against.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/libraryRequirementType
	kSecLibraryRequirementType SecRequirementType = 0
	// kSecPluginRequirementType - What plug-ins this code may load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/pluginRequirementType
	kSecPluginRequirementType SecRequirementType = 0
	// kSecRequirementTypeCount - The number of valid requirement types.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType/requirementTypeCount
	kSecRequirementTypeCount SecRequirementType = 0
)

/* debug [enums.gen.go]: Processing enum SecTransformMetaAttributeType (11 cases) */
// SecTransformMetaAttributeType - The keys that describe the metadata attributes of transform attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType
type SecTransformMetaAttributeType uint

const (
	// kSecTransformMetaAttributeCanCycle - The transform allows cyclic behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/canCycle
	kSecTransformMetaAttributeCanCycle SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeDeferred - The attribute defers notifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/deferred
	kSecTransformMetaAttributeDeferred SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeExternalize - The attribute is exportable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/externalize
	kSecTransformMetaAttributeExternalize SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeHasInboundConnection - The attribute has an inbound connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/hasInboundConnection
	kSecTransformMetaAttributeHasInboundConnection SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeHasOutboundConnections - The attribute has an outbound connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/hasOutboundConnections
	kSecTransformMetaAttributeHasOutboundConnections SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeName - The name of the attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/name
	kSecTransformMetaAttributeName SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeRef - A direct reference to an attribute’s value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/ref
	kSecTransformMetaAttributeRef SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeRequired - Indicates whether the attribute value is optional.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/required
	kSecTransformMetaAttributeRequired SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeRequiresOutboundConnection - The attribute requires an outbound connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/requiresOutboundConnection
	kSecTransformMetaAttributeRequiresOutboundConnection SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeStream - The attribute expects stream operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/stream
	kSecTransformMetaAttributeStream SecTransformMetaAttributeType = 0
	// kSecTransformMetaAttributeValue - The actual value of the attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType/value
	kSecTransformMetaAttributeValue SecTransformMetaAttributeType = 0
)

/* debug [enums.gen.go]: Processing enum SecTrustOptionFlags (7 cases) */
// SecTrustOptionFlags - The option flags used to condition a trust evaluation.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags
type SecTrustOptionFlags uint

const (
	// kSecTrustOptionAllowExpired - Allow expired certificates (except for the root certificate).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/allowExpired
	kSecTrustOptionAllowExpired SecTrustOptionFlags = 0
	// kSecTrustOptionAllowExpiredRoot - Allow expired root certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/allowExpiredRoot
	kSecTrustOptionAllowExpiredRoot SecTrustOptionFlags = 0
	// kSecTrustOptionFetchIssuerFromNet - Allow network downloads of CA certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/fetchIssuerFromNet
	kSecTrustOptionFetchIssuerFromNet SecTrustOptionFlags = 0
	// kSecTrustOptionImplicitAnchors - Treat properly self-signed certificates as anchors implicitly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/implicitAnchors
	kSecTrustOptionImplicitAnchors SecTrustOptionFlags = 0
	// kSecTrustOptionLeafIsCA - Allow CA certificates as leaf certificates.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/leafIsCA
	kSecTrustOptionLeafIsCA SecTrustOptionFlags = 0
	// kSecTrustOptionRequireRevPerCert - Require a positive revocation check for each certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/requireRevPerCert
	kSecTrustOptionRequireRevPerCert SecTrustOptionFlags = 0
	// kSecTrustOptionUseTrustSettings - Use TrustSettings instead of anchors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustOptionFlags/useTrustSettings
	kSecTrustOptionUseTrustSettings SecTrustOptionFlags = 0
)

/* debug [enums.gen.go]: Processing enum SecTrustResultType (8 cases) */
// SecTrustResultType - Trust evaluation result codes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType
type SecTrustResultType uint

const (
	// kSecTrustResultConfirm - User confirmation is required before proceeding.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/confirm
	kSecTrustResultConfirm SecTrustResultType = 0
	// kSecTrustResultDeny - The user specified that the certificate should not be trusted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/deny
	kSecTrustResultDeny SecTrustResultType = 0
	// kSecTrustResultFatalTrustFailure - Trust is denied and no simple fix is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/fatalTrustFailure
	kSecTrustResultFatalTrustFailure SecTrustResultType = 0
	// kSecTrustResultInvalid - An indication of an invalid setting or result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/invalid
	kSecTrustResultInvalid SecTrustResultType = 0
	// kSecTrustResultOtherError - A value that indicates a failure other than trust evaluation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/otherError
	kSecTrustResultOtherError SecTrustResultType = 0
	// kSecTrustResultProceed - The user granted permission to trust the certificate for the purposes designated in the specified policies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/proceed
	kSecTrustResultProceed SecTrustResultType = 0
	// kSecTrustResultRecoverableTrustFailure - Trust is denied, but recovery may be possible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/recoverableTrustFailure
	kSecTrustResultRecoverableTrustFailure SecTrustResultType = 0
	// kSecTrustResultUnspecified - The user did not specify a trust setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType/unspecified
	kSecTrustResultUnspecified SecTrustResultType = 0
)

/* debug [enums.gen.go]: Processing enum SecTrustSettingsDomain (3 cases) */
// SecTrustSettingsDomain - The trust settings domains.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsDomain
type SecTrustSettingsDomain uint

const (
	// kSecTrustSettingsDomainAdmin - Locally administered, system-wide trust settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsDomain/admin
	kSecTrustSettingsDomainAdmin SecTrustSettingsDomain = 0
	// kSecTrustSettingsDomainSystem - System trust settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsDomain/system
	kSecTrustSettingsDomainSystem SecTrustSettingsDomain = 0
	// kSecTrustSettingsDomainUser - Per-user trust settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsDomain/user
	kSecTrustSettingsDomainUser SecTrustSettingsDomain = 0
)

/* debug [enums.gen.go]: Processing enum SecTrustSettingsKeyUsage (7 cases) */
// SecTrustSettingsKeyUsage - Allowed uses for the encryption key in a certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage
type SecTrustSettingsKeyUsage uint

const (
	// kSecTrustSettingsKeyUseAny - The key can be used for any purpose.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useAny
	kSecTrustSettingsKeyUseAny SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseEnDecryptData - The key can be used to encrypt or decrypt data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useEnDecryptData
	kSecTrustSettingsKeyUseEnDecryptData SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseEnDecryptKey - The key can be used to encrypt or decrypt (wrap or unwrap) a key.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useEnDecryptKey
	kSecTrustSettingsKeyUseEnDecryptKey SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseKeyExchange - The key is a private key that has been shared using a key exchange protocol, such as Diffie-Hellman key exchange.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useKeyExchange
	kSecTrustSettingsKeyUseKeyExchange SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseSignature - The key can be used to sign data or verify a signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useSignature
	kSecTrustSettingsKeyUseSignature SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseSignCert - The key can be used to sign a certificate or verify a signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useSignCert
	kSecTrustSettingsKeyUseSignCert SecTrustSettingsKeyUsage = 0
	// kSecTrustSettingsKeyUseSignRevocation - The key can be used to sign an OCSP (online certificate status protocol) message or CRL (certificate verification list), or to verify a signature.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsKeyUsage/useSignRevocation
	kSecTrustSettingsKeyUseSignRevocation SecTrustSettingsKeyUsage = 0
)

/* debug [enums.gen.go]: Processing enum SecTrustSettingsResult (5 cases) */
// SecTrustSettingsResult - Trust settings returned in usage constraints dictionaries.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult
type SecTrustSettingsResult uint

const (
	// kSecTrustSettingsResultDeny - This certificate is explicitly distrusted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult/deny
	kSecTrustSettingsResultDeny SecTrustSettingsResult = 0
	// kSecTrustSettingsResultInvalid - Never valid in a trust settings array or in an API call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult/invalid
	kSecTrustSettingsResultInvalid SecTrustSettingsResult = 0
	// kSecTrustSettingsResultTrustAsRoot - This non-root certificate is explicitly trusted as if it were a trusted root.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult/trustAsRoot
	kSecTrustSettingsResultTrustAsRoot SecTrustSettingsResult = 0
	// kSecTrustSettingsResultTrustRoot - This root certificate is explicitly trusted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult/trustRoot
	kSecTrustSettingsResultTrustRoot SecTrustSettingsResult = 0
	// kSecTrustSettingsResultUnspecified - This certificate is neither trusted nor distrusted. This value can be used to specify an “allowed error” without assigning trust to a specific certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsResult/unspecified
	kSecTrustSettingsResultUnspecified SecTrustSettingsResult = 0
)

/* debug [enums.gen.go]: Processing enum SessionAttributeBits (4 cases) */
// SessionAttributeBits - The attributes of a security session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits
type SessionAttributeBits uint

const (
	// sessionHasGraphicAccess - A bit that indicates a graphic subsystem is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits/sessionHasGraphicAccess
	sessionHasGraphicAccess SessionAttributeBits = 0
	// sessionHasTTY - A bit that indicates   is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits/sessionHasTTY
	sessionHasTTY SessionAttributeBits = 0
	// sessionIsRemote - A bit that indicates the session was initiated over the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits/sessionIsRemote
	sessionIsRemote SessionAttributeBits = 0
	// sessionIsRoot - A bit that indicates the session is the root session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits/sessionIsRoot
	sessionIsRoot SessionAttributeBits = 0
)

/* debug [enums.gen.go]: Processing enum SessionCreationFlags (1 cases) */
// SessionCreationFlags - The flags that affect the creation of a security session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionCreationFlags
type SessionCreationFlags uint

const (
	// sessionKeepCurrentBootstrap - The caller has allocated sub-bootstrap.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SessionCreationFlags/sessionKeepCurrentBootstrap
	sessionKeepCurrentBootstrap SessionCreationFlags = 0
)

/* debug [enums.gen.go]: Processing enum SSLAuthenticate (3 cases) */
// SSLAuthenticate - The flags that represent the requirements for client-side authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAuthenticate
type SSLAuthenticate uint

const (
	// kAlwaysAuthenticate - Indicates that client-side authentication is required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAuthenticate/alwaysAuthenticate
	kAlwaysAuthenticate SSLAuthenticate = 0
	// kNeverAuthenticate - Indicates that client-side authentication is not required. (Default.)
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAuthenticate/neverAuthenticate
	kNeverAuthenticate SSLAuthenticate = 0
	// kTryAuthenticate - Indicates that client-side authentication should be attempted. There is no error if the client doesn’t have a certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAuthenticate/tryAuthenticate
	kTryAuthenticate SSLAuthenticate = 0
)

/* debug [enums.gen.go]: Processing enum SSLCiphersuiteGroup (5 cases) */
// SSLCiphersuiteGroup - A mechanism for grouping related cipher suites.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup
type SSLCiphersuiteGroup uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup/ATS
	kSSLCiphersuiteGroupATS SSLCiphersuiteGroup = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup/atsCompatibility
	kSSLCiphersuiteGroupATSCompatibility SSLCiphersuiteGroup = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup/compatibility
	kSSLCiphersuiteGroupCompatibility SSLCiphersuiteGroup = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup/default
	kSSLCiphersuiteGroupDefault SSLCiphersuiteGroup = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup/legacy
	kSSLCiphersuiteGroupLegacy SSLCiphersuiteGroup = 0
)

/* debug [enums.gen.go]: Processing enum SSLClientCertificateState (4 cases) */
// SSLClientCertificateState - An enumeration of the states of client certificate exchange.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState
type SSLClientCertificateState uint

const (
	// kSSLClientCertNone - Indicates that the server hasn’t asked for a certificate and that the client hasn’t sent one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState/certNone
	kSSLClientCertNone SSLClientCertificateState = 0
	// kSSLClientCertRejected - Indicates that the client sent a certificate but the certificate failed validation. This value is seen only on the server side. The server application can inspect the certificate using the function  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState/certRejected
	kSSLClientCertRejected SSLClientCertificateState = 0
	// kSSLClientCertRequested - Indicates that the server has asked for a certificate, but the client has not sent it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState/certRequested
	kSSLClientCertRequested SSLClientCertificateState = 0
	// kSSLClientCertSent - Indicates that the server asked for a certificate, the client sent one, and the server validated it. The application can inspect the certificate using the function  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState/certSent
	kSSLClientCertSent SSLClientCertificateState = 0
)

/* debug [enums.gen.go]: Processing enum SSLConnectionType (2 cases) */
// SSLConnectionType - The flags that indicate whether a context is to be used for streaming or datagram-based communication.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLConnectionType
type SSLConnectionType uint

const (
	// kSSLDatagramType - Datagram-based communication (UDP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLConnectionType/datagramType
	kSSLDatagramType SSLConnectionType = 0
	// kSSLStreamType - Stream-based communication (TCP).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLConnectionType/streamType
	kSSLStreamType SSLConnectionType = 0
)

/* debug [enums.gen.go]: Processing enum SSLProtocol (13 cases) */
// SSLProtocol - An enumeration of valid SSL protocol versions.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol
type SSLProtocol uint

const (
	// kDTLSProtocol1 - Specifies the DTLS 1.0 protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/dtlsProtocol1
	kDTLSProtocol1 SSLProtocol = 0
	// kDTLSProtocol12 - Specifies the DTLS 1.2 protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/dtlsProtocol12
	kDTLSProtocol12 SSLProtocol = 0
	// kSSLProtocol2 - Specifies that only the SSL 2.0 protocol may be negotiated. Deprecated in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/sslProtocol2
	kSSLProtocol2 SSLProtocol = 0
	// kSSLProtocol3 - Specifies that the SSL 3.0 protocol is preferred; the SSL 2.0 protocol may be negotiated if the peer cannot use the SSL 3.0 protocol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/sslProtocol3
	kSSLProtocol3 SSLProtocol = 0
	// kSSLProtocol3Only - Specifies that only the SSL 3.0 protocol may be negotiated; fails if the peer tries to negotiate the SSL 2.0 protocol. Deprecated in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/sslProtocol3Only
	kSSLProtocol3Only SSLProtocol = 0
	// kSSLProtocolAll - Specifies all supported versions. Deprecated in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/sslProtocolAll
	kSSLProtocolAll SSLProtocol = 0
	// kSSLProtocolUnknown - Specifies that no protocol has been or should be negotiated or specified; use default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/sslProtocolUnknown
	kSSLProtocolUnknown SSLProtocol = 0
	// kTLSProtocol1 - Specifies that the TLS 1.0 protocol is preferred but lower versions may be negotiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocol1
	kTLSProtocol1 SSLProtocol = 0
	// kTLSProtocol11 - Specifies that the TLS 1.1 protocol is preferred but lower versions may be negotiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocol11
	kTLSProtocol11 SSLProtocol = 0
	// kTLSProtocol12 - Specifies that the TLS 1.2 protocol is preferred but lower versions may be negotiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocol12
	kTLSProtocol12 SSLProtocol = 0
	// kTLSProtocol13 - Specifies that the TLS 1.3 protocol is preferred but lower versions may be negotiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocol13
	kTLSProtocol13 SSLProtocol = 0
	// kTLSProtocol1Only - Specifies that only the TLS 1.0 protocol may be negotiated. Deprecated in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocol1Only
	kTLSProtocol1Only SSLProtocol = 0
	// kTLSProtocolMaxSupported - The maximum system supported version.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol/tlsProtocolMaxSupported
	kTLSProtocolMaxSupported SSLProtocol = 0
)

/* debug [enums.gen.go]: Processing enum SSLProtocolSide (2 cases) */
// SSLProtocolSide - The flags that indicate whether a context is for the server or client side of a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocolSide
type SSLProtocolSide uint

const (
	// kSSLClientSide - Client side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocolSide/clientSide
	kSSLClientSide SSLProtocolSide = 0
	// kSSLServerSide - Server side.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocolSide/serverSide
	kSSLServerSide SSLProtocolSide = 0
)

/* debug [enums.gen.go]: Processing enum SSLSessionOption (10 cases) */
// SSLSessionOption - The options that can be set for an SSL session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption
type SSLSessionOption uint

const (
	// kSSLSessionOptionAllowRenegotiation - Allow renegotiation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/allowRenegotiation
	kSSLSessionOptionAllowRenegotiation SSLSessionOption = 0
	// kSSLSessionOptionAllowServerIdentityChange - Allow server identity change on renegotiation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/allowServerIdentityChange
	kSSLSessionOptionAllowServerIdentityChange SSLSessionOption = 0
	// kSSLSessionOptionBreakOnCertRequested - Enables returning from   (with a result of  ) when the server requests a client certificate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/breakOnCertRequested
	kSSLSessionOptionBreakOnCertRequested SSLSessionOption = 0
	// kSSLSessionOptionBreakOnClientAuth - Enables returning from   (with a result of  ) when the client authentication portion of the handshake is complete to allow your application to perform its own certificate verification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/breakOnClientAuth
	kSSLSessionOptionBreakOnClientAuth SSLSessionOption = 0
	// kSSLSessionOptionBreakOnClientHello - Break from a client hello in order to check for SNI.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/breakOnClientHello
	kSSLSessionOptionBreakOnClientHello SSLSessionOption = 0
	// kSSLSessionOptionBreakOnServerAuth - Enables returning from   (with a result of  ) when the server authentication portion of the handshake is complete to allow your application to perform its own certificate verification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/breakOnServerAuth
	kSSLSessionOptionBreakOnServerAuth SSLSessionOption = 0
	// kSSLSessionOptionEnableSessionTickets - Enable session tickets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/enableSessionTickets
	kSSLSessionOptionEnableSessionTickets SSLSessionOption = 0
	// kSSLSessionOptionFallback - Enable fallback countermeasures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/fallback
	kSSLSessionOptionFallback SSLSessionOption = 0
	// kSSLSessionOptionFalseStart - When enabled, TLS False Start is used if an adequate cipher-suite is negotiated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/falseStart
	kSSLSessionOptionFalseStart SSLSessionOption = 0
	// kSSLSessionOptionSendOneByteRecord - Enables   record splitting for BEAST attack mitigation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption/sendOneByteRecord
	kSSLSessionOptionSendOneByteRecord SSLSessionOption = 0
)

/* debug [enums.gen.go]: Processing enum SSLSessionState (5 cases) */
// SSLSessionState - The flags that represent the state of an SSL session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState
type SSLSessionState uint

const (
	// kSSLAborted - The connection aborted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState/aborted
	kSSLAborted SSLSessionState = 0
	// kSSLClosed - The connection closed normally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState/closed
	kSSLClosed SSLSessionState = 0
	// kSSLConnected - The SSL handshake is complete; the connection is ready for normal I/O.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState/connected
	kSSLConnected SSLSessionState = 0
	// kSSLHandshake - The SSL handshake is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState/handshake
	kSSLHandshake SSLSessionState = 0
	// kSSLIdle - No I/O has been performed yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState/idle
	kSSLIdle SSLSessionState = 0
)

/* debug [enums.gen.go]: Processing enum tls_ciphersuite_group_t (5 cases) */
// tls_ciphersuite_group_t - Groups that collect ciphersuites of comparable security properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t
type tls_ciphersuite_group_t uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t/ats
	tls_ciphersuite_group_ats tls_ciphersuite_group_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t/ats_compatibility
	tls_ciphersuite_group_ats_compatibility tls_ciphersuite_group_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t/compatibility
	tls_ciphersuite_group_compatibility tls_ciphersuite_group_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t/default
	tls_ciphersuite_group_default tls_ciphersuite_group_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t/legacy
	tls_ciphersuite_group_legacy tls_ciphersuite_group_t = 0
)

/* debug [enums.gen.go]: Processing enum tls_ciphersuite_t (26 cases) */
// tls_ciphersuite_t - The collection of valid ciphersuites.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t
type tls_ciphersuite_t uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/AES_128_GCM_SHA256
	tls_ciphersuite_AES_128_GCM_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/AES_256_GCM_SHA384
	tls_ciphersuite_AES_256_GCM_SHA384 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/CHACHA20_POLY1305_SHA256
	tls_ciphersuite_CHACHA20_POLY1305_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA
	tls_ciphersuite_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_128_CBC_SHA
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_128_CBC_SHA256
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_256_CBC_SHA
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_256_CBC_SHA384
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
	tls_ciphersuite_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	tls_ciphersuite_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_3DES_EDE_CBC_SHA
	tls_ciphersuite_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_128_CBC_SHA
	tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_128_CBC_SHA256
	tls_ciphersuite_ECDHE_RSA_WITH_AES_128_CBC_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_128_GCM_SHA256
	tls_ciphersuite_ECDHE_RSA_WITH_AES_128_GCM_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_256_CBC_SHA
	tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_256_CBC_SHA384
	tls_ciphersuite_ECDHE_RSA_WITH_AES_256_CBC_SHA384 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_AES_256_GCM_SHA384
	tls_ciphersuite_ECDHE_RSA_WITH_AES_256_GCM_SHA384 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
	tls_ciphersuite_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_3DES_EDE_CBC_SHA
	tls_ciphersuite_RSA_WITH_3DES_EDE_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_128_CBC_SHA
	tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_128_CBC_SHA256
	tls_ciphersuite_RSA_WITH_AES_128_CBC_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_128_GCM_SHA256
	tls_ciphersuite_RSA_WITH_AES_128_GCM_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_256_CBC_SHA
	tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_256_CBC_SHA256
	tls_ciphersuite_RSA_WITH_AES_256_CBC_SHA256 tls_ciphersuite_t = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t/RSA_WITH_AES_256_GCM_SHA384
	tls_ciphersuite_RSA_WITH_AES_256_GCM_SHA384 tls_ciphersuite_t = 0
)


