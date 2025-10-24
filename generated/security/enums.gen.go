// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

// Enum types and constants
// AuthorizationContextFlags - The flags that specify whether authentication data should be made available to the authorization client.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationContextFlags
type AuthorizationContextFlags uint

const (
	// kAuthorizationContextFlagExtractable - It is possible for the authorization client to use the   function to obtain the value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationContextFlags/kAuthorizationContextFlagExtractable
	kAuthorizationContextFlagExtractable AuthorizationContextFlags = 0
	// kAuthorizationContextFlagSticky - This data persists through an interrupted or failed evaluation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationContextFlags/kAuthorizationContextFlagSticky
	kAuthorizationContextFlagSticky AuthorizationContextFlags = 0
	// kAuthorizationContextFlagVolatile - The value is not saved for the authorization client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationContextFlags/kAuthorizationContextFlagVolatile
	kAuthorizationContextFlagVolatile AuthorizationContextFlags = 0
)

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

// AuthorizationResult - The permissible values for an authorization evaluation result.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationResult
type AuthorizationResult uint

const (
	// kAuthorizationResultAllow - The authorization operation succeeded and authorization should be granted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationResult/kAuthorizationResultAllow
	kAuthorizationResultAllow AuthorizationResult = 0
	// kAuthorizationResultDeny - The authorization operation succeeded and authorization should be denied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationResult/kAuthorizationResultDeny
	kAuthorizationResultDeny AuthorizationResult = 0
	// kAuthorizationResultUndefined - The authorization operation failed and should not be retried for this session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationResult/kAuthorizationResultUndefined
	kAuthorizationResultUndefined AuthorizationResult = 0
	// kAuthorizationResultUserCanceled - The user has requested that the authorization evaluation be terminated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationResult/kAuthorizationResultUserCanceled
	kAuthorizationResultUserCanceled AuthorizationResult = 0
)

// __CE_CrlDistributionPointNameType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CrlDistributionPointNameType
type __CE_CrlDistributionPointNameType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CDNT_FullName
	CE_CDNT_FullName __CE_CrlDistributionPointNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CDNT_NameRelativeToCrlIssuer
	CE_CDNT_NameRelativeToCrlIssuer __CE_CrlDistributionPointNameType = 0
)

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

// __CE_GeneralNameType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralNameType-c.enum
type __CE_GeneralNameType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_DNSName
	GNT_DNSName __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_DirectoryName
	GNT_DirectoryName __CE_GeneralNameType = 0
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
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_RFC822Name
	GNT_RFC822Name __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_RegisteredID
	GNT_RegisteredID __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_URI
	GNT_URI __CE_GeneralNameType = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/GNT_X400Address
	GNT_X400Address __CE_GeneralNameType = 0
)

// CMSCertificateChainMode - Constants that can be set to specify what certificates to include in a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSCertificateChainMode
type CMSCertificateChainMode uint

// CMSSignedAttributes - Optional attributes you can add to a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignedAttributes
type CMSSignedAttributes uint

// CMSSignerStatus - The constants that indicate the status of the signature and signer information in a signed message.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSSignerStatus
type CMSSignerStatus uint

// SSLAuthenticate - The flags that represent the requirements for client-side authentication.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAuthenticate
type SSLAuthenticate uint

// SSLCiphersuiteGroup - A mechanism for grouping related cipher suites.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCiphersuiteGroup
type SSLCiphersuiteGroup uint

// SSLClientCertificateState - An enumeration of the states of client certificate exchange.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClientCertificateState
type SSLClientCertificateState uint

// SSLConnectionType - The flags that indicate whether a context is to be used for streaming or datagram-based communication.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLConnectionType
type SSLConnectionType uint

// SSLProtocol - An enumeration of valid SSL protocol versions.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocol
type SSLProtocol uint

// SSLProtocolSide - The flags that indicate whether a context is for the server or client side of a connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLProtocolSide
type SSLProtocolSide uint

// SSLSessionOption - The options that can be set for an SSL session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionOption
type SSLSessionOption uint

// SSLSessionState - The flags that represent the state of an SSL session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSessionState
type SSLSessionState uint

// SecAccessControlCreateFlags - Access control constants that dictate how a keychain item may be used.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateFlags
type SecAccessControlCreateFlags uint

// SecAuthenticationType - The authentication type to use for an Internet password.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAuthenticationType
type SecAuthenticationType uint

// SecCSDigestAlgorithm - The list of digest algorithms available for code signatures.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSDigestAlgorithm
type SecCSDigestAlgorithm uint

// SecCSFlags - Values that can be used in the 
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags
type SecCSFlags uint

const (
	// kSecCSDefaultFlags - No flags (use the default behavior).
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecCSFlags/kSecCSDefaultFlags
	kSecCSDefaultFlags SecCSFlags = 0
)

// SecCodeSignatureFlags - Specify option flags that can be embedded in a code signature during signing and that govern the use of the signature.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeSignatureFlags
type SecCodeSignatureFlags uint

// SecCodeStatus - Operational flags attached by code signing services to running code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeStatus
type SecCodeStatus uint

// SecExternalFormat - The external format of a keychain item.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalFormat
type SecExternalFormat uint

// SecExternalItemType - The import item type.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecExternalItemType
type SecExternalItemType uint

// SecItemAttr - Specifies a keychain item’s attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAttr
type SecItemAttr uint

// SecItemClass - Specifies a keychain item’s class code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass
type SecItemClass uint

const (
	// kSecGenericPasswordItemClass - Indicates that the item is a generic password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/genericPasswordItemClass
	kSecGenericPasswordItemClass SecItemClass = 0
	// kSecInternetPasswordItemClass - Indicates that the item is an Internet password.
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemClass/internetPasswordItemClass
	kSecInternetPasswordItemClass SecItemClass = 0
)

// SecItemImportExportFlags - The import and export function flags.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportFlags
type SecItemImportExportFlags uint

// SecKeyImportExportFlags - The import/export parameter structure flags.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportFlags
type SecKeyImportExportFlags uint

// SecProtocolType - The protocol type associated with an Internet password.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecProtocolType
type SecProtocolType uint

// SecRequirementType - An enumeration indicating different types of internal requirements for code.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementType
type SecRequirementType uint

// SecTransformMetaAttributeType - The keys that describe the metadata attributes of transform attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformMetaAttributeType
type SecTransformMetaAttributeType uint

// SecTrustResultType - Trust evaluation result codes.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustResultType
type SecTrustResultType uint

// _SecureDownloadTrustCallbackResult - A flag used to indicate whether or not a signer should be evaluated.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadTrustCallbackResult
type _SecureDownloadTrustCallbackResult uint

// SessionAttributeBits - The attributes of a security session.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionAttributeBits
type SessionAttributeBits uint

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

// cssm_appledl_open_parameters_mask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters_mask
type cssm_appledl_open_parameters_mask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/kCSSM_APPLEDL_MASK_MODE
	kCSSM_APPLEDL_MASK_MODE cssm_appledl_open_parameters_mask = 0
)

// extension_data_format enum type
//
// [Full Topic]: https://developer.apple.com/documentation/Security/extension_data_format
type extension_data_format uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509_DATAFORMAT_ENCODED
	CSSM_X509_DATAFORMAT_ENCODED extension_data_format = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509_DATAFORMAT_PAIR
	CSSM_X509_DATAFORMAT_PAIR extension_data_format = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509_DATAFORMAT_PARSED
	CSSM_X509_DATAFORMAT_PARSED extension_data_format = 0
)

// tls_ciphersuite_group_t - Groups that collect ciphersuites of comparable security properties.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_group_t
type tls_ciphersuite_group_t uint

// tls_ciphersuite_t - The collection of valid ciphersuites.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/tls_ciphersuite_t
type tls_ciphersuite_t uint

// tls_protocol_version_t - The collection of supported TLS and DTLS versions.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/tls_protocol_version_t
type tls_protocol_version_t uint


