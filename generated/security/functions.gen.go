// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

/* debug [functions.gen.go]: Generating 396 functions for Security */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Security Functions (396 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SecACLCopySimpleContents func(SecACLRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLCreateFromSimpleContents func(SecAccessRef, ArrayRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLGetAuthorizations func(SecACLRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLSetAuthorizations func(SecACLRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLSetSimpleContents func(SecACLRef, ArrayRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_SecAccessCopySelectedACLList func(SecAccessRef, CSSM_ACL_AUTHORIZATION_TAG, unsafe.Pointer) unsafe.Pointer
	_SecAccessCreateFromOwnerAndACL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessGetOwnerAndACL func(SecAccessRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyNotValidAfterDate func(SecCertificateRef) DateRef
	_SecCertificateCopyPreference func(StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCertificateCreateFromData func(unsafe.Pointer, CSSM_CERT_TYPE, CSSM_CERT_ENCODING, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetAlgorithmID func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetCLHandle func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetData func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetIssuer func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetSubject func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateGetType func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateSetPreference func(SecCertificateRef, StringRef, unsafe.Pointer, DateRef) unsafe.Pointer
	_SecIdentityCopyPreference func(StringRef, CSSM_KEYUSE, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_SecIdentityCreate func(AllocatorRef, SecCertificateRef, SecKeyRef) SecIdentityRef
	_SecIdentitySearchCopyNext func(SecIdentitySearchRef, unsafe.Pointer) unsafe.Pointer
	_SecIdentitySearchCreate func(TypeRef, CSSM_KEYUSE, unsafe.Pointer) unsafe.Pointer
	_SecIdentitySearchGetTypeID func() TypeID
	_SecIdentitySetPreference func(SecIdentityRef, StringRef, CSSM_KEYUSE) unsafe.Pointer
	_SecKeyCreatePair func(SecKeychainRef, CSSM_ALGORITHMS, unsafe.Pointer, CSSM_CC_HANDLE, CSSM_KEYUSE, unsafe.Pointer, CSSM_KEYUSE, unsafe.Pointer, SecAccessRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyGenerate func(SecKeychainRef, CSSM_ALGORITHMS, unsafe.Pointer, CSSM_CC_HANDLE, CSSM_KEYUSE, unsafe.Pointer, SecAccessRef, unsafe.Pointer) unsafe.Pointer
	_SecKeyGetCSPHandle func(SecKeyRef, unsafe.Pointer) unsafe.Pointer
	_SecKeyGetCSSMKey func(SecKeyRef, unsafe.Pointer) unsafe.Pointer
	_SecKeyGetCredentials func(SecKeyRef, CSSM_ACL_AUTHORIZATION_TAG, SecCredentialType, unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetCSPHandle func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetDLDBHandle func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemExport func(TypeRef, SecExternalFormat, SecItemImportExportFlags, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemGetDLDBHandle func(SecKeychainItemRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemGetUniqueRecordID func(SecKeychainItemRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemImport func(DataRef, StringRef, unsafe.Pointer, unsafe.Pointer, SecItemImportExportFlags, unsafe.Pointer, SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainSearchCopyNext func(SecKeychainSearchRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainSearchCreateFromAttributes func(TypeRef, SecItemClass, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainSearchGetTypeID func() TypeID
	_SecPolicyCreateWithOID func(TypeRef) SecPolicyRef
	_SecPolicyGetOID func(SecPolicyRef, unsafe.Pointer) unsafe.Pointer
	_SecPolicyGetTPHandle func(SecPolicyRef, unsafe.Pointer) unsafe.Pointer
	_SecPolicyGetValue func(SecPolicyRef, unsafe.Pointer) unsafe.Pointer
	_SecPolicySearchCopyNext func(SecPolicySearchRef, unsafe.Pointer) unsafe.Pointer
	_SecPolicySearchCreate func(CSSM_CERT_TYPE, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecPolicySearchGetTypeID func() TypeID
	_SecPolicySetProperties func(SecPolicyRef, DictionaryRef) unsafe.Pointer
	_SecPolicySetValue func(SecPolicyRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetCssmResult func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetCssmResultCode func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetResult func(SecTrustRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetTPHandle func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetParameters func(SecTrustRef, CSSM_TP_ACTION, DataRef) unsafe.Pointer
	_SecureDownloadCopyCreationDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyInfo func(AuthorizationRef, AuthorizationString, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyRights func(AuthorizationRef, unsafe.Pointer, unsafe.Pointer, AuthorizationFlags, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyRightsAsync func(AuthorizationRef, unsafe.Pointer, unsafe.Pointer, AuthorizationFlags, unsafe.Pointer)
	_AuthorizationCreate func(unsafe.Pointer, unsafe.Pointer, AuthorizationFlags, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCreateFromExternalForm func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationFree func(AuthorizationRef, AuthorizationFlags) unsafe.Pointer
	_AuthorizationFreeItemSet func(unsafe.Pointer) unsafe.Pointer
	_AuthorizationMakeExternalForm func(AuthorizationRef, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightRemove func(AuthorizationRef, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightSet func(AuthorizationRef, unsafe.Pointer, TypeRef, StringRef, BundleRef, StringRef) unsafe.Pointer
	_CMSDecoderCopyAllCerts func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyContent func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyDetachedContent func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyEncapsulatedContentType func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerCert func(SDecoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerEmailAddress func(SDecoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerSigningTime func(SDecoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerStatus func(SDecoderRef, uintptr, TypeRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestamp func(SDecoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestampCertificates func(SDecoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestampWithPolicy func(SDecoderRef, TypeRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCreate func(unsafe.Pointer) unsafe.Pointer
	_CMSDecoderFinalizeMessage func(SDecoderRef) unsafe.Pointer
	_CMSDecoderGetNumSigners func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderGetTypeID func() TypeID
	_CMSDecoderIsContentEncrypted func(SDecoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderSetDetachedContent func(SDecoderRef, DataRef) unsafe.Pointer
	_CMSDecoderSetSearchKeychain func(SDecoderRef, TypeRef) unsafe.Pointer
	_CMSDecoderUpdateMessage func(SDecoderRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_CMSEncodeContent func(TypeRef, TypeRef, TypeRef, unsafe.Pointer, SSignedAttributes, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderAddRecipients func(SEncoderRef, TypeRef) unsafe.Pointer
	_CMSEncoderAddSignedAttributes func(SEncoderRef, SSignedAttributes) unsafe.Pointer
	_CMSEncoderAddSigners func(SEncoderRef, TypeRef) unsafe.Pointer
	_CMSEncoderAddSupportingCerts func(SEncoderRef, TypeRef) unsafe.Pointer
	_CMSEncoderCopyEncapsulatedContentType func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopyEncodedContent func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopyRecipients func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySigners func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySignerTimestamp func(SEncoderRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySignerTimestampWithPolicy func(SEncoderRef, TypeRef, uintptr, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySupportingCerts func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCreate func(unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetCertificateChainMode func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetHasDetachedContent func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetTypeID func() TypeID
	_CMSEncoderSetCertificateChainMode func(SEncoderRef, SCertificateChainMode) unsafe.Pointer
	_CMSEncoderSetEncapsulatedContentTypeOID func(SEncoderRef, TypeRef) unsafe.Pointer
	_CMSEncoderSetHasDetachedContent func(SEncoderRef, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderSetSignerAlgorithm func(SEncoderRef, StringRef) unsafe.Pointer
	_CMSEncoderUpdateContent func(SEncoderRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_cssmAlgToOid func(CSSM_ALGORITHMS) unsafe.Pointer
	_cssmOidToAlg func(unsafe.Pointer, unsafe.Pointer) bool
	_cssmPerror func(unsafe.Pointer, CSSM_RETURN)
	_sec_protocol_metadata_copy_negotiated_protocol func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_copy_server_name func(unsafe.Pointer) unsafe.Pointer
	_SecAccessControlCreateWithFlags func(AllocatorRef, TypeRef, SecAccessControlCreateFlags, unsafe.Pointer) SecAccessControlRef
	_SecAccessControlGetTypeID func() TypeID
	_SecAccessCopyACLList func(SecAccessRef, unsafe.Pointer) unsafe.Pointer
	_SecAccessCopyMatchingACLList func(SecAccessRef, TypeRef) ArrayRef
	_SecAccessCopyOwnerAndACL func(SecAccessRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessCreate func(StringRef, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_SecAccessCreateWithOwnerAndACL func(unsafe.Pointer, unsafe.Pointer, SecAccessOwnerType, ArrayRef, unsafe.Pointer) SecAccessRef
	_SecAccessGetTypeID func() TypeID
	_SecACLCopyAuthorizations func(SecACLRef) ArrayRef
	_SecACLCopyContents func(SecACLRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLCreateWithSimpleContents func(SecAccessRef, ArrayRef, StringRef, SecKeychainPromptSelector, unsafe.Pointer) unsafe.Pointer
	_SecACLGetTypeID func() TypeID
	_SecACLRemove func(SecACLRef) unsafe.Pointer
	_SecACLSetContents func(SecACLRef, ArrayRef, StringRef, SecKeychainPromptSelector) unsafe.Pointer
	_SecACLUpdateAuthorizations func(SecACLRef, ArrayRef) unsafe.Pointer
	_SecAddSharedWebCredential func(StringRef, StringRef, StringRef)
	_SecCertificateAddToKeychain func(SecCertificateRef, SecKeychainRef) unsafe.Pointer
	_SecCertificateCopyCommonName func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyData func(SecCertificateRef) DataRef
	_SecCertificateCopyEmailAddresses func(SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyKey func(SecCertificateRef) SecKeyRef
	_SecCertificateCopyLongDescription func(AllocatorRef, SecCertificateRef, unsafe.Pointer) StringRef
	_SecCertificateCopyNormalizedIssuerContent func(SecCertificateRef, unsafe.Pointer) DataRef
	_SecCertificateCopyNormalizedIssuerSequence func(SecCertificateRef) DataRef
	_SecCertificateCopyNormalizedSubjectContent func(SecCertificateRef, unsafe.Pointer) DataRef
	_SecCertificateCopyNormalizedSubjectSequence func(SecCertificateRef) DataRef
	_SecCertificateCopyPreferred func(StringRef, ArrayRef) SecCertificateRef
	_SecCertificateCopyPublicKey func(SecCertificateRef) SecKeyRef
	_SecCertificateCopySerialNumber func(SecCertificateRef) DataRef
	_SecCertificateCopySerialNumberData func(SecCertificateRef, unsafe.Pointer) DataRef
	_SecCertificateCopyShortDescription func(AllocatorRef, SecCertificateRef, unsafe.Pointer) StringRef
	_SecCertificateCopySubjectSummary func(SecCertificateRef) StringRef
	_SecCertificateCopyValues func(SecCertificateRef, ArrayRef, unsafe.Pointer) DictionaryRef
	_SecCertificateCreateWithData func(AllocatorRef, DataRef) SecCertificateRef
	_SecCertificateGetTypeID func() TypeID
	_SecCertificateSetPreferred func(SecCertificateRef, StringRef, ArrayRef) unsafe.Pointer
	_SecCodeCheckValidity func(SecCodeRef, SecCSFlags, SecRequirementRef) unsafe.Pointer
	_SecCodeCheckValidityWithErrors func(SecCodeRef, SecCSFlags, SecRequirementRef, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyDesignatedRequirement func(SecStaticCodeRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyGuestWithAttributes func(SecCodeRef, DictionaryRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyHost func(SecCodeRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyPath func(SecStaticCodeRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopySelf func(SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopySigningInformation func(SecStaticCodeRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyStaticCode func(SecCodeRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecCodeGetTypeID func() TypeID
	_SecCodeMapMemory func(SecStaticCodeRef, SecCSFlags) unsafe.Pointer
	_SecCopyErrorMessageString func(unsafe.Pointer, unsafe.Pointer) StringRef
	_SecCreateSharedWebCredentialPassword func() StringRef
	_SecDecodeTransformCreate func(TypeRef, unsafe.Pointer) SecTransformRef
	_SecDecryptTransformCreate func(SecKeyRef, unsafe.Pointer) SecTransformRef
	_SecDecryptTransformGetTypeID func() TypeID
	_SecDigestTransformCreate func(TypeRef, Index, unsafe.Pointer) SecTransformRef
	_SecDigestTransformGetTypeID func() TypeID
	_SecEncodeTransformCreate func(TypeRef, unsafe.Pointer) SecTransformRef
	_SecEncryptTransformCreate func(SecKeyRef, unsafe.Pointer) SecTransformRef
	_SecEncryptTransformGetTypeID func() TypeID
	_SecGroupTransformGetTypeID func() TypeID
	_SecIdentityCopyCertificate func(SecIdentityRef, unsafe.Pointer) unsafe.Pointer
	_SecIdentityCopyPreferred func(StringRef, ArrayRef, ArrayRef) SecIdentityRef
	_SecIdentityCopyPrivateKey func(SecIdentityRef, unsafe.Pointer) unsafe.Pointer
	_SecIdentityCopySystemIdentity func(StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecIdentityCreateWithCertificate func(TypeRef, SecCertificateRef, unsafe.Pointer) unsafe.Pointer
	_SecIdentityGetTypeID func() TypeID
	_SecIdentitySetPreferred func(SecIdentityRef, StringRef, ArrayRef) unsafe.Pointer
	_SecIdentitySetSystemIdentity func(StringRef, SecIdentityRef) unsafe.Pointer
	_SecItemAdd func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_SecItemCopyMatching func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_SecItemDelete func(DictionaryRef) unsafe.Pointer
	_SecItemExport func(TypeRef, SecExternalFormat, SecItemImportExportFlags, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecItemImport func(DataRef, StringRef, unsafe.Pointer, unsafe.Pointer, SecItemImportExportFlags, unsafe.Pointer, SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecItemUpdate func(DictionaryRef, DictionaryRef) unsafe.Pointer
	_SecKeychainAddCallback func(SecKeychainCallback, SecKeychainEventMask, unsafe.Pointer) unsafe.Pointer
	_SecKeychainAddGenericPassword func(SecKeychainRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainAddInternetPassword func(SecKeychainRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, SecProtocolType, SecAuthenticationType, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainAttributeInfoForItemID func(SecKeychainRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopyAccess func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopyDefault func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopyDomainDefault func(SecPreferencesDomain, unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopyDomainSearchList func(SecPreferencesDomain, unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopySearchList func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainCopySettings func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, SecAccessRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainDelete func(SecKeychainRef) unsafe.Pointer
	_SecKeychainFindGenericPassword func(TypeRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainFindInternetPassword func(TypeRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, SecProtocolType, SecAuthenticationType, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainFreeAttributeInfo func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetPath func(SecKeychainRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetPreferenceDomain func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetStatus func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetTypeID func() TypeID
	_SecKeychainGetUserInteractionAllowed func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainGetVersion func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyAccess func(SecKeychainItemRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyAttributesAndData func(SecKeychainItemRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyContent func(SecKeychainItemRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyFromPersistentReference func(DataRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyKeychain func(SecKeychainItemRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCreateCopy func(SecKeychainItemRef, SecKeychainRef, SecAccessRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCreateFromContent func(SecItemClass, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, SecKeychainRef, SecAccessRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCreatePersistentReference func(SecKeychainItemRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemDelete func(SecKeychainItemRef) unsafe.Pointer
	_SecKeychainItemFreeAttributesAndData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemFreeContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemGetTypeID func() TypeID
	_SecKeychainItemModifyAttributesAndData func(SecKeychainItemRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemModifyContent func(SecKeychainItemRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemSetAccess func(SecKeychainItemRef, SecAccessRef) unsafe.Pointer
	_SecKeychainLock func(SecKeychainRef) unsafe.Pointer
	_SecKeychainLockAll func() unsafe.Pointer
	_SecKeychainOpen func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainRemoveCallback func(SecKeychainCallback) unsafe.Pointer
	_SecKeychainSetAccess func(SecKeychainRef, SecAccessRef) unsafe.Pointer
	_SecKeychainSetDefault func(SecKeychainRef) unsafe.Pointer
	_SecKeychainSetDomainDefault func(SecPreferencesDomain, SecKeychainRef) unsafe.Pointer
	_SecKeychainSetDomainSearchList func(SecPreferencesDomain, ArrayRef) unsafe.Pointer
	_SecKeychainSetPreferenceDomain func(SecPreferencesDomain) unsafe.Pointer
	_SecKeychainSetSearchList func(ArrayRef) unsafe.Pointer
	_SecKeychainSetSettings func(SecKeychainRef, unsafe.Pointer) unsafe.Pointer
	_SecKeychainSetUserInteractionAllowed func(unsafe.Pointer) unsafe.Pointer
	_SecKeychainUnlock func(SecKeychainRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCopyAttributes func(SecKeyRef) DictionaryRef
	_SecKeyCopyExternalRepresentation func(SecKeyRef, unsafe.Pointer) DataRef
	_SecKeyCopyKeyExchangeResult func(SecKeyRef, SecKeyAlgorithm, SecKeyRef, DictionaryRef, unsafe.Pointer) DataRef
	_SecKeyCopyPublicKey func(SecKeyRef) SecKeyRef
	_SecKeyCreateDecryptedData func(SecKeyRef, SecKeyAlgorithm, DataRef, unsafe.Pointer) DataRef
	_SecKeyCreateEncryptedData func(SecKeyRef, SecKeyAlgorithm, DataRef, unsafe.Pointer) DataRef
	_SecKeyCreateFromData func(DictionaryRef, DataRef, unsafe.Pointer) SecKeyRef
	_SecKeyCreateRandomKey func(DictionaryRef, unsafe.Pointer) SecKeyRef
	_SecKeyCreateSignature func(SecKeyRef, SecKeyAlgorithm, DataRef, unsafe.Pointer) DataRef
	_SecKeyCreateWithData func(DataRef, DictionaryRef, unsafe.Pointer) SecKeyRef
	_SecKeyDecrypt func(SecKeyRef, SecPadding, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyDeriveFromPassword func(StringRef, DictionaryRef, unsafe.Pointer) SecKeyRef
	_SecKeyEncrypt func(SecKeyRef, SecPadding, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyGeneratePair func(DictionaryRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyGeneratePairAsync func(DictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_SecKeyGenerateSymmetric func(DictionaryRef, unsafe.Pointer) SecKeyRef
	_SecKeyGetBlockSize func(SecKeyRef) uintptr
	_SecKeyGetTypeID func() TypeID
	_SecKeyIsAlgorithmSupported func(SecKeyRef, SecKeyOperationType, SecKeyAlgorithm) unsafe.Pointer
	_SecKeyRawSign func(SecKeyRef, SecPadding, unsafe.Pointer, uintptr, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyRawVerify func(SecKeyRef, SecPadding, unsafe.Pointer, uintptr, unsafe.Pointer, uintptr) unsafe.Pointer
	_SecKeyUnwrapSymmetric func(unsafe.Pointer, SecKeyRef, DictionaryRef, unsafe.Pointer) SecKeyRef
	_SecKeyVerifySignature func(SecKeyRef, SecKeyAlgorithm, DataRef, DataRef, unsafe.Pointer) unsafe.Pointer
	_SecKeyWrapSymmetric func(SecKeyRef, SecKeyRef, DictionaryRef, unsafe.Pointer) DataRef
	_SecPKCS12Import func(DataRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_SecPolicyCopyProperties func(SecPolicyRef) DictionaryRef
	_SecPolicyCreateBasicX509 func() SecPolicyRef
	_SecPolicyCreateRevocation func(OptionFlags) SecPolicyRef
	_SecPolicyCreateSSL func(unsafe.Pointer, StringRef) SecPolicyRef
	_SecPolicyCreateWithProperties func(TypeRef, DictionaryRef) SecPolicyRef
	_SecPolicyGetTypeID func() TypeID
	_SecRandomCopyBytes func(SecRandomRef, uintptr, unsafe.Pointer) int
	_SecRequestSharedWebCredential func(StringRef, StringRef)
	_SecRequirementCopyData func(SecRequirementRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCopyString func(SecRequirementRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithData func(DataRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithString func(StringRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithStringAndErrors func(StringRef, SecCSFlags, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementGetTypeID func() TypeID
	_SecSignTransformCreate func(SecKeyRef, unsafe.Pointer) SecTransformRef
	_SecStaticCodeCheckValidity func(SecStaticCodeRef, SecCSFlags, SecRequirementRef) unsafe.Pointer
	_SecStaticCodeCheckValidityWithErrors func(SecStaticCodeRef, SecCSFlags, SecRequirementRef, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCreateWithPath func(URLRef, SecCSFlags, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCreateWithPathAndAttributes func(URLRef, SecCSFlags, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeGetTypeID func() TypeID
	_SecTaskCopySigningIdentifier func(SecTaskRef, unsafe.Pointer) StringRef
	_SecTaskCopyValueForEntitlement func(SecTaskRef, StringRef, unsafe.Pointer) TypeRef
	_SecTaskCopyValuesForEntitlements func(SecTaskRef, ArrayRef, unsafe.Pointer) DictionaryRef
	_SecTaskCreateFromSelf func(AllocatorRef) SecTaskRef
	_SecTaskCreateWithAuditToken func(AllocatorRef, unsafe.Pointer) SecTaskRef
	_SecTaskGetTypeID func() TypeID
	_SecTransformConnectTransforms func(SecTransformRef, StringRef, SecTransformRef, StringRef, SecGroupTransformRef, unsafe.Pointer) SecGroupTransformRef
	_SecTransformCopyExternalRepresentation func(SecTransformRef) DictionaryRef
	_SecTransformCreate func(StringRef, unsafe.Pointer) SecTransformRef
	_SecTransformCreateFromExternalRepresentation func(DictionaryRef, unsafe.Pointer) SecTransformRef
	_SecTransformCreateGroupTransform func() SecGroupTransformRef
	_SecTransformCreateReadTransformWithReadStream func(ReadStreamRef) SecTransformRef
	_SecTransformCustomGetAttribute func(SecTransformImplementationRef, SecTransformStringOrAttributeRef, SecTransformMetaAttributeType) TypeRef
	_SecTransformCustomSetAttribute func(SecTransformImplementationRef, SecTransformStringOrAttributeRef, SecTransformMetaAttributeType, TypeRef) TypeRef
	_SecTransformExecute func(SecTransformRef, unsafe.Pointer) TypeRef
	_SecTransformExecuteAsync func(SecTransformRef, unsafe.Pointer, unsafe.Pointer)
	_SecTransformFindByName func(SecGroupTransformRef, StringRef) SecTransformRef
	_SecTransformGetAttribute func(SecTransformRef, StringRef) TypeRef
	_SecTransformGetTypeID func() TypeID
	_SecTransformNoData func() TypeRef
	_SecTransformPushbackAttribute func(SecTransformImplementationRef, SecTransformStringOrAttributeRef, TypeRef) TypeRef
	_SecTransformRegister func(StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetAttribute func(SecTransformRef, StringRef, TypeRef, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetAttributeAction func(SecTransformImplementationRef, StringRef, SecTransformStringOrAttributeRef, unsafe.Pointer) ErrorRef
	_SecTransformSetDataAction func(SecTransformImplementationRef, StringRef, unsafe.Pointer) ErrorRef
	_SecTransformSetTransformAction func(SecTransformImplementationRef, StringRef, unsafe.Pointer) ErrorRef
	_SecTrustCopyAnchorCertificates func(unsafe.Pointer) unsafe.Pointer
	_SecTrustCopyCustomAnchorCertificates func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustCopyExceptions func(SecTrustRef) DataRef
	_SecTrustCopyPolicies func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustCopyProperties func(SecTrustRef) ArrayRef
	_SecTrustCopyPublicKey func(SecTrustRef) SecKeyRef
	_SecTrustCopyResult func(SecTrustRef) DictionaryRef
	_SecTrustCreateWithCertificates func(TypeRef, TypeRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustedApplicationCopyData func(SecTrustedApplicationRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustedApplicationCreateFromPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustedApplicationGetTypeID func() TypeID
	_SecTrustedApplicationSetData func(SecTrustedApplicationRef, DataRef) unsafe.Pointer
	_SecTrustEvaluate func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluateAsync func(SecTrustRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluateAsyncWithError func(SecTrustRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluateWithError func(SecTrustRef, unsafe.Pointer) bool
	_SecTrustGetCertificateAtIndex func(SecTrustRef, Index) SecCertificateRef
	_SecTrustGetCertificateCount func(SecTrustRef) Index
	_SecTrustGetNetworkFetchAllowed func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetTrustResult func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetTypeID func() TypeID
	_SecTrustGetVerifyTime func(SecTrustRef) AbsoluteTime
	_SecTrustSetAnchorCertificates func(SecTrustRef, ArrayRef) unsafe.Pointer
	_SecTrustSetAnchorCertificatesOnly func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetExceptions func(SecTrustRef, DataRef) bool
	_SecTrustSetKeychains func(SecTrustRef, TypeRef) unsafe.Pointer
	_SecTrustSetNetworkFetchAllowed func(SecTrustRef, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetOCSPResponse func(SecTrustRef, TypeRef) unsafe.Pointer
	_SecTrustSetOptions func(SecTrustRef, SecTrustOptionFlags) unsafe.Pointer
	_SecTrustSetPolicies func(SecTrustRef, TypeRef) unsafe.Pointer
	_SecTrustSetSignedCertificateTimestamps func(SecTrustRef, ArrayRef) unsafe.Pointer
	_SecTrustSettingsCopyCertificates func(SecTrustSettingsDomain, unsafe.Pointer) unsafe.Pointer
	_SecTrustSettingsCopyModificationDate func(SecCertificateRef, SecTrustSettingsDomain, unsafe.Pointer) unsafe.Pointer
	_SecTrustSettingsCopyTrustSettings func(SecCertificateRef, SecTrustSettingsDomain, unsafe.Pointer) unsafe.Pointer
	_SecTrustSettingsCreateExternalRepresentation func(SecTrustSettingsDomain, unsafe.Pointer) unsafe.Pointer
	_SecTrustSettingsImportExternalRepresentation func(SecTrustSettingsDomain, DataRef) unsafe.Pointer
	_SecTrustSettingsRemoveTrustSettings func(SecCertificateRef, SecTrustSettingsDomain) unsafe.Pointer
	_SecTrustSettingsSetTrustSettings func(SecCertificateRef, SecTrustSettingsDomain, TypeRef) unsafe.Pointer
	_SecTrustSetVerifyDate func(SecTrustRef, DateRef) unsafe.Pointer
	_SecVerifyTransformCreate func(SecKeyRef, DataRef, unsafe.Pointer) SecTransformRef
	_SessionCreate func(SessionCreationFlags, SessionAttributeBits) unsafe.Pointer
	_SessionGetInfo func(SecuritySessionId, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLAddDistinguishedName func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLClose func(SSLContextRef) unsafe.Pointer
	_SSLContextGetTypeID func() TypeID
	_SSLCopyALPNProtocols func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLCopyCertificateAuthorities func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLCopyDistinguishedNames func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLCopyPeerTrust func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLCopyRequestedPeerName func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyRequestedPeerNameLength func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLCreateContext func(AllocatorRef, SSLProtocolSide, SSLConnectionType) SSLContextRef
	_SSLGetBufferedReadSize func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetClientCertificateState func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetConnection func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetDatagramWriteSize func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetDiffieHellmanParams func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetEnabledCiphers func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetMaxDatagramRecordSize func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetNegotiatedCipher func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetNegotiatedProtocolVersion func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetNumberEnabledCiphers func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetNumberSupportedCiphers func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerDomainName func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerDomainNameLength func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerID func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersionMax func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersionMin func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetSessionOption func(SSLContextRef, SSLSessionOption, unsafe.Pointer) unsafe.Pointer
	_SSLGetSessionState func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLGetSupportedCiphers func(SSLContextRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLHandshake func(SSLContextRef) unsafe.Pointer
	_SSLRead func(SSLContextRef, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_SSLReHandshake func(SSLContextRef) unsafe.Pointer
	_SSLSetALPNProtocols func(SSLContextRef, ArrayRef) unsafe.Pointer
	_SSLSetCertificate func(SSLContextRef, ArrayRef) unsafe.Pointer
	_SSLSetCertificateAuthorities func(SSLContextRef, TypeRef, unsafe.Pointer) unsafe.Pointer
	_SSLSetClientSideAuthenticate func(SSLContextRef, SSLAuthenticate) unsafe.Pointer
	_SSLSetConnection func(SSLContextRef, SSLConnectionRef) unsafe.Pointer
	_SSLSetDatagramHelloCookie func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLSetDiffieHellmanParams func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLSetEnabledCiphers func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLSetEncryptionCertificate func(SSLContextRef, ArrayRef) unsafe.Pointer
	_SSLSetError func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLSetIOFuncs func(SSLContextRef, SSLReadFunc, SSLWriteFunc) unsafe.Pointer
	_SSLSetMaxDatagramRecordSize func(SSLContextRef, uintptr) unsafe.Pointer
	_SSLSetOCSPResponse func(SSLContextRef, DataRef) unsafe.Pointer
	_SSLSetPeerDomainName func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLSetPeerID func(SSLContextRef, unsafe.Pointer, uintptr) unsafe.Pointer
	_SSLSetProtocolVersionMax func(SSLContextRef, SSLProtocol) unsafe.Pointer
	_SSLSetProtocolVersionMin func(SSLContextRef, SSLProtocol) unsafe.Pointer
	_SSLSetSessionConfig func(SSLContextRef, StringRef) unsafe.Pointer
	_SSLSetSessionOption func(SSLContextRef, SSLSessionOption, unsafe.Pointer) unsafe.Pointer
	_SSLSetSessionTicketsEnabled func(SSLContextRef, unsafe.Pointer) unsafe.Pointer
	_SSLWrite func(SSLContextRef, unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SecACLCopySimpleContents, lib, "SecACLCopySimpleContents")
	tryRegister(&_SecACLCreateFromSimpleContents, lib, "SecACLCreateFromSimpleContents")
	tryRegister(&_SecACLGetAuthorizations, lib, "SecACLGetAuthorizations")
	tryRegister(&_SecACLSetAuthorizations, lib, "SecACLSetAuthorizations")
	tryRegister(&_SecACLSetSimpleContents, lib, "SecACLSetSimpleContents")
	tryRegister(&_SecAccessCopySelectedACLList, lib, "SecAccessCopySelectedACLList")
	tryRegister(&_SecAccessCreateFromOwnerAndACL, lib, "SecAccessCreateFromOwnerAndACL")
	tryRegister(&_SecAccessGetOwnerAndACL, lib, "SecAccessGetOwnerAndACL")
	tryRegister(&_SecCertificateCopyNotValidAfterDate, lib, "SecCertificateCopyNotValidAfterDate")
	tryRegister(&_SecCertificateCopyPreference, lib, "SecCertificateCopyPreference")
	tryRegister(&_SecCertificateCreateFromData, lib, "SecCertificateCreateFromData")
	tryRegister(&_SecCertificateGetAlgorithmID, lib, "SecCertificateGetAlgorithmID")
	tryRegister(&_SecCertificateGetCLHandle, lib, "SecCertificateGetCLHandle")
	tryRegister(&_SecCertificateGetData, lib, "SecCertificateGetData")
	tryRegister(&_SecCertificateGetIssuer, lib, "SecCertificateGetIssuer")
	tryRegister(&_SecCertificateGetSubject, lib, "SecCertificateGetSubject")
	tryRegister(&_SecCertificateGetType, lib, "SecCertificateGetType")
	tryRegister(&_SecCertificateSetPreference, lib, "SecCertificateSetPreference")
	tryRegister(&_SecIdentityCopyPreference, lib, "SecIdentityCopyPreference")
	tryRegister(&_SecIdentityCreate, lib, "SecIdentityCreate")
	tryRegister(&_SecIdentitySearchCopyNext, lib, "SecIdentitySearchCopyNext")
	tryRegister(&_SecIdentitySearchCreate, lib, "SecIdentitySearchCreate")
	tryRegister(&_SecIdentitySearchGetTypeID, lib, "SecIdentitySearchGetTypeID")
	tryRegister(&_SecIdentitySetPreference, lib, "SecIdentitySetPreference")
	tryRegister(&_SecKeyCreatePair, lib, "SecKeyCreatePair")
	tryRegister(&_SecKeyGenerate, lib, "SecKeyGenerate")
	tryRegister(&_SecKeyGetCSPHandle, lib, "SecKeyGetCSPHandle")
	tryRegister(&_SecKeyGetCSSMKey, lib, "SecKeyGetCSSMKey")
	tryRegister(&_SecKeyGetCredentials, lib, "SecKeyGetCredentials")
	tryRegister(&_SecKeychainGetCSPHandle, lib, "SecKeychainGetCSPHandle")
	tryRegister(&_SecKeychainGetDLDBHandle, lib, "SecKeychainGetDLDBHandle")
	tryRegister(&_SecKeychainItemExport, lib, "SecKeychainItemExport")
	tryRegister(&_SecKeychainItemGetDLDBHandle, lib, "SecKeychainItemGetDLDBHandle")
	tryRegister(&_SecKeychainItemGetUniqueRecordID, lib, "SecKeychainItemGetUniqueRecordID")
	tryRegister(&_SecKeychainItemImport, lib, "SecKeychainItemImport")
	tryRegister(&_SecKeychainSearchCopyNext, lib, "SecKeychainSearchCopyNext")
	tryRegister(&_SecKeychainSearchCreateFromAttributes, lib, "SecKeychainSearchCreateFromAttributes")
	tryRegister(&_SecKeychainSearchGetTypeID, lib, "SecKeychainSearchGetTypeID")
	tryRegister(&_SecPolicyCreateWithOID, lib, "SecPolicyCreateWithOID")
	tryRegister(&_SecPolicyGetOID, lib, "SecPolicyGetOID")
	tryRegister(&_SecPolicyGetTPHandle, lib, "SecPolicyGetTPHandle")
	tryRegister(&_SecPolicyGetValue, lib, "SecPolicyGetValue")
	tryRegister(&_SecPolicySearchCopyNext, lib, "SecPolicySearchCopyNext")
	tryRegister(&_SecPolicySearchCreate, lib, "SecPolicySearchCreate")
	tryRegister(&_SecPolicySearchGetTypeID, lib, "SecPolicySearchGetTypeID")
	tryRegister(&_SecPolicySetProperties, lib, "SecPolicySetProperties")
	tryRegister(&_SecPolicySetValue, lib, "SecPolicySetValue")
	tryRegister(&_SecTrustGetCssmResult, lib, "SecTrustGetCssmResult")
	tryRegister(&_SecTrustGetCssmResultCode, lib, "SecTrustGetCssmResultCode")
	tryRegister(&_SecTrustGetResult, lib, "SecTrustGetResult")
	tryRegister(&_SecTrustGetTPHandle, lib, "SecTrustGetTPHandle")
	tryRegister(&_SecTrustSetParameters, lib, "SecTrustSetParameters")
	tryRegister(&_SecureDownloadCopyCreationDate, lib, "SecureDownloadCopyCreationDate")
	tryRegister(&_AuthorizationCopyInfo, lib, "AuthorizationCopyInfo")
	tryRegister(&_AuthorizationCopyRights, lib, "AuthorizationCopyRights")
	tryRegister(&_AuthorizationCopyRightsAsync, lib, "AuthorizationCopyRightsAsync")
	tryRegister(&_AuthorizationCreate, lib, "AuthorizationCreate")
	tryRegister(&_AuthorizationCreateFromExternalForm, lib, "AuthorizationCreateFromExternalForm")
	tryRegister(&_AuthorizationFree, lib, "AuthorizationFree")
	tryRegister(&_AuthorizationFreeItemSet, lib, "AuthorizationFreeItemSet")
	tryRegister(&_AuthorizationMakeExternalForm, lib, "AuthorizationMakeExternalForm")
	tryRegister(&_AuthorizationRightGet, lib, "AuthorizationRightGet")
	tryRegister(&_AuthorizationRightRemove, lib, "AuthorizationRightRemove")
	tryRegister(&_AuthorizationRightSet, lib, "AuthorizationRightSet")
	tryRegister(&_CMSDecoderCopyAllCerts, lib, "CMSDecoderCopyAllCerts")
	tryRegister(&_CMSDecoderCopyContent, lib, "CMSDecoderCopyContent")
	tryRegister(&_CMSDecoderCopyDetachedContent, lib, "CMSDecoderCopyDetachedContent")
	tryRegister(&_CMSDecoderCopyEncapsulatedContentType, lib, "CMSDecoderCopyEncapsulatedContentType")
	tryRegister(&_CMSDecoderCopySignerCert, lib, "CMSDecoderCopySignerCert")
	tryRegister(&_CMSDecoderCopySignerEmailAddress, lib, "CMSDecoderCopySignerEmailAddress")
	tryRegister(&_CMSDecoderCopySignerSigningTime, lib, "CMSDecoderCopySignerSigningTime")
	tryRegister(&_CMSDecoderCopySignerStatus, lib, "CMSDecoderCopySignerStatus")
	tryRegister(&_CMSDecoderCopySignerTimestamp, lib, "CMSDecoderCopySignerTimestamp")
	tryRegister(&_CMSDecoderCopySignerTimestampCertificates, lib, "CMSDecoderCopySignerTimestampCertificates")
	tryRegister(&_CMSDecoderCopySignerTimestampWithPolicy, lib, "CMSDecoderCopySignerTimestampWithPolicy")
	tryRegister(&_CMSDecoderCreate, lib, "CMSDecoderCreate")
	tryRegister(&_CMSDecoderFinalizeMessage, lib, "CMSDecoderFinalizeMessage")
	tryRegister(&_CMSDecoderGetNumSigners, lib, "CMSDecoderGetNumSigners")
	tryRegister(&_CMSDecoderGetTypeID, lib, "CMSDecoderGetTypeID")
	tryRegister(&_CMSDecoderIsContentEncrypted, lib, "CMSDecoderIsContentEncrypted")
	tryRegister(&_CMSDecoderSetDetachedContent, lib, "CMSDecoderSetDetachedContent")
	tryRegister(&_CMSDecoderSetSearchKeychain, lib, "CMSDecoderSetSearchKeychain")
	tryRegister(&_CMSDecoderUpdateMessage, lib, "CMSDecoderUpdateMessage")
	tryRegister(&_CMSEncodeContent, lib, "CMSEncodeContent")
	tryRegister(&_CMSEncoderAddRecipients, lib, "CMSEncoderAddRecipients")
	tryRegister(&_CMSEncoderAddSignedAttributes, lib, "CMSEncoderAddSignedAttributes")
	tryRegister(&_CMSEncoderAddSigners, lib, "CMSEncoderAddSigners")
	tryRegister(&_CMSEncoderAddSupportingCerts, lib, "CMSEncoderAddSupportingCerts")
	tryRegister(&_CMSEncoderCopyEncapsulatedContentType, lib, "CMSEncoderCopyEncapsulatedContentType")
	tryRegister(&_CMSEncoderCopyEncodedContent, lib, "CMSEncoderCopyEncodedContent")
	tryRegister(&_CMSEncoderCopyRecipients, lib, "CMSEncoderCopyRecipients")
	tryRegister(&_CMSEncoderCopySigners, lib, "CMSEncoderCopySigners")
	tryRegister(&_CMSEncoderCopySignerTimestamp, lib, "CMSEncoderCopySignerTimestamp")
	tryRegister(&_CMSEncoderCopySignerTimestampWithPolicy, lib, "CMSEncoderCopySignerTimestampWithPolicy")
	tryRegister(&_CMSEncoderCopySupportingCerts, lib, "CMSEncoderCopySupportingCerts")
	tryRegister(&_CMSEncoderCreate, lib, "CMSEncoderCreate")
	tryRegister(&_CMSEncoderGetCertificateChainMode, lib, "CMSEncoderGetCertificateChainMode")
	tryRegister(&_CMSEncoderGetHasDetachedContent, lib, "CMSEncoderGetHasDetachedContent")
	tryRegister(&_CMSEncoderGetTypeID, lib, "CMSEncoderGetTypeID")
	tryRegister(&_CMSEncoderSetCertificateChainMode, lib, "CMSEncoderSetCertificateChainMode")
	tryRegister(&_CMSEncoderSetEncapsulatedContentTypeOID, lib, "CMSEncoderSetEncapsulatedContentTypeOID")
	tryRegister(&_CMSEncoderSetHasDetachedContent, lib, "CMSEncoderSetHasDetachedContent")
	tryRegister(&_CMSEncoderSetSignerAlgorithm, lib, "CMSEncoderSetSignerAlgorithm")
	tryRegister(&_CMSEncoderUpdateContent, lib, "CMSEncoderUpdateContent")
	tryRegister(&_cssmAlgToOid, lib, "cssmAlgToOid")
	tryRegister(&_cssmOidToAlg, lib, "cssmOidToAlg")
	tryRegister(&_cssmPerror, lib, "cssmPerror")
	tryRegister(&_sec_protocol_metadata_copy_negotiated_protocol, lib, "sec_protocol_metadata_copy_negotiated_protocol")
	tryRegister(&_sec_protocol_metadata_copy_server_name, lib, "sec_protocol_metadata_copy_server_name")
	tryRegister(&_SecAccessControlCreateWithFlags, lib, "SecAccessControlCreateWithFlags")
	tryRegister(&_SecAccessControlGetTypeID, lib, "SecAccessControlGetTypeID")
	tryRegister(&_SecAccessCopyACLList, lib, "SecAccessCopyACLList")
	tryRegister(&_SecAccessCopyMatchingACLList, lib, "SecAccessCopyMatchingACLList")
	tryRegister(&_SecAccessCopyOwnerAndACL, lib, "SecAccessCopyOwnerAndACL")
	tryRegister(&_SecAccessCreate, lib, "SecAccessCreate")
	tryRegister(&_SecAccessCreateWithOwnerAndACL, lib, "SecAccessCreateWithOwnerAndACL")
	tryRegister(&_SecAccessGetTypeID, lib, "SecAccessGetTypeID")
	tryRegister(&_SecACLCopyAuthorizations, lib, "SecACLCopyAuthorizations")
	tryRegister(&_SecACLCopyContents, lib, "SecACLCopyContents")
	tryRegister(&_SecACLCreateWithSimpleContents, lib, "SecACLCreateWithSimpleContents")
	tryRegister(&_SecACLGetTypeID, lib, "SecACLGetTypeID")
	tryRegister(&_SecACLRemove, lib, "SecACLRemove")
	tryRegister(&_SecACLSetContents, lib, "SecACLSetContents")
	tryRegister(&_SecACLUpdateAuthorizations, lib, "SecACLUpdateAuthorizations")
	tryRegister(&_SecAddSharedWebCredential, lib, "SecAddSharedWebCredential")
	tryRegister(&_SecCertificateAddToKeychain, lib, "SecCertificateAddToKeychain")
	tryRegister(&_SecCertificateCopyCommonName, lib, "SecCertificateCopyCommonName")
	tryRegister(&_SecCertificateCopyData, lib, "SecCertificateCopyData")
	tryRegister(&_SecCertificateCopyEmailAddresses, lib, "SecCertificateCopyEmailAddresses")
	tryRegister(&_SecCertificateCopyKey, lib, "SecCertificateCopyKey")
	tryRegister(&_SecCertificateCopyLongDescription, lib, "SecCertificateCopyLongDescription")
	tryRegister(&_SecCertificateCopyNormalizedIssuerContent, lib, "SecCertificateCopyNormalizedIssuerContent")
	tryRegister(&_SecCertificateCopyNormalizedIssuerSequence, lib, "SecCertificateCopyNormalizedIssuerSequence")
	tryRegister(&_SecCertificateCopyNormalizedSubjectContent, lib, "SecCertificateCopyNormalizedSubjectContent")
	tryRegister(&_SecCertificateCopyNormalizedSubjectSequence, lib, "SecCertificateCopyNormalizedSubjectSequence")
	tryRegister(&_SecCertificateCopyPreferred, lib, "SecCertificateCopyPreferred")
	tryRegister(&_SecCertificateCopyPublicKey, lib, "SecCertificateCopyPublicKey")
	tryRegister(&_SecCertificateCopySerialNumber, lib, "SecCertificateCopySerialNumber")
	tryRegister(&_SecCertificateCopySerialNumberData, lib, "SecCertificateCopySerialNumberData")
	tryRegister(&_SecCertificateCopyShortDescription, lib, "SecCertificateCopyShortDescription")
	tryRegister(&_SecCertificateCopySubjectSummary, lib, "SecCertificateCopySubjectSummary")
	tryRegister(&_SecCertificateCopyValues, lib, "SecCertificateCopyValues")
	tryRegister(&_SecCertificateCreateWithData, lib, "SecCertificateCreateWithData")
	tryRegister(&_SecCertificateGetTypeID, lib, "SecCertificateGetTypeID")
	tryRegister(&_SecCertificateSetPreferred, lib, "SecCertificateSetPreferred")
	tryRegister(&_SecCodeCheckValidity, lib, "SecCodeCheckValidity")
	tryRegister(&_SecCodeCheckValidityWithErrors, lib, "SecCodeCheckValidityWithErrors")
	tryRegister(&_SecCodeCopyDesignatedRequirement, lib, "SecCodeCopyDesignatedRequirement")
	tryRegister(&_SecCodeCopyGuestWithAttributes, lib, "SecCodeCopyGuestWithAttributes")
	tryRegister(&_SecCodeCopyHost, lib, "SecCodeCopyHost")
	tryRegister(&_SecCodeCopyPath, lib, "SecCodeCopyPath")
	tryRegister(&_SecCodeCopySelf, lib, "SecCodeCopySelf")
	tryRegister(&_SecCodeCopySigningInformation, lib, "SecCodeCopySigningInformation")
	tryRegister(&_SecCodeCopyStaticCode, lib, "SecCodeCopyStaticCode")
	tryRegister(&_SecCodeGetTypeID, lib, "SecCodeGetTypeID")
	tryRegister(&_SecCodeMapMemory, lib, "SecCodeMapMemory")
	tryRegister(&_SecCopyErrorMessageString, lib, "SecCopyErrorMessageString")
	tryRegister(&_SecCreateSharedWebCredentialPassword, lib, "SecCreateSharedWebCredentialPassword")
	tryRegister(&_SecDecodeTransformCreate, lib, "SecDecodeTransformCreate")
	tryRegister(&_SecDecryptTransformCreate, lib, "SecDecryptTransformCreate")
	tryRegister(&_SecDecryptTransformGetTypeID, lib, "SecDecryptTransformGetTypeID")
	tryRegister(&_SecDigestTransformCreate, lib, "SecDigestTransformCreate")
	tryRegister(&_SecDigestTransformGetTypeID, lib, "SecDigestTransformGetTypeID")
	tryRegister(&_SecEncodeTransformCreate, lib, "SecEncodeTransformCreate")
	tryRegister(&_SecEncryptTransformCreate, lib, "SecEncryptTransformCreate")
	tryRegister(&_SecEncryptTransformGetTypeID, lib, "SecEncryptTransformGetTypeID")
	tryRegister(&_SecGroupTransformGetTypeID, lib, "SecGroupTransformGetTypeID")
	tryRegister(&_SecIdentityCopyCertificate, lib, "SecIdentityCopyCertificate")
	tryRegister(&_SecIdentityCopyPreferred, lib, "SecIdentityCopyPreferred")
	tryRegister(&_SecIdentityCopyPrivateKey, lib, "SecIdentityCopyPrivateKey")
	tryRegister(&_SecIdentityCopySystemIdentity, lib, "SecIdentityCopySystemIdentity")
	tryRegister(&_SecIdentityCreateWithCertificate, lib, "SecIdentityCreateWithCertificate")
	tryRegister(&_SecIdentityGetTypeID, lib, "SecIdentityGetTypeID")
	tryRegister(&_SecIdentitySetPreferred, lib, "SecIdentitySetPreferred")
	tryRegister(&_SecIdentitySetSystemIdentity, lib, "SecIdentitySetSystemIdentity")
	tryRegister(&_SecItemAdd, lib, "SecItemAdd")
	tryRegister(&_SecItemCopyMatching, lib, "SecItemCopyMatching")
	tryRegister(&_SecItemDelete, lib, "SecItemDelete")
	tryRegister(&_SecItemExport, lib, "SecItemExport")
	tryRegister(&_SecItemImport, lib, "SecItemImport")
	tryRegister(&_SecItemUpdate, lib, "SecItemUpdate")
	tryRegister(&_SecKeychainAddCallback, lib, "SecKeychainAddCallback")
	tryRegister(&_SecKeychainAddGenericPassword, lib, "SecKeychainAddGenericPassword")
	tryRegister(&_SecKeychainAddInternetPassword, lib, "SecKeychainAddInternetPassword")
	tryRegister(&_SecKeychainAttributeInfoForItemID, lib, "SecKeychainAttributeInfoForItemID")
	tryRegister(&_SecKeychainCopyAccess, lib, "SecKeychainCopyAccess")
	tryRegister(&_SecKeychainCopyDefault, lib, "SecKeychainCopyDefault")
	tryRegister(&_SecKeychainCopyDomainDefault, lib, "SecKeychainCopyDomainDefault")
	tryRegister(&_SecKeychainCopyDomainSearchList, lib, "SecKeychainCopyDomainSearchList")
	tryRegister(&_SecKeychainCopySearchList, lib, "SecKeychainCopySearchList")
	tryRegister(&_SecKeychainCopySettings, lib, "SecKeychainCopySettings")
	tryRegister(&_SecKeychainCreate, lib, "SecKeychainCreate")
	tryRegister(&_SecKeychainDelete, lib, "SecKeychainDelete")
	tryRegister(&_SecKeychainFindGenericPassword, lib, "SecKeychainFindGenericPassword")
	tryRegister(&_SecKeychainFindInternetPassword, lib, "SecKeychainFindInternetPassword")
	tryRegister(&_SecKeychainFreeAttributeInfo, lib, "SecKeychainFreeAttributeInfo")
	tryRegister(&_SecKeychainGetPath, lib, "SecKeychainGetPath")
	tryRegister(&_SecKeychainGetPreferenceDomain, lib, "SecKeychainGetPreferenceDomain")
	tryRegister(&_SecKeychainGetStatus, lib, "SecKeychainGetStatus")
	tryRegister(&_SecKeychainGetTypeID, lib, "SecKeychainGetTypeID")
	tryRegister(&_SecKeychainGetUserInteractionAllowed, lib, "SecKeychainGetUserInteractionAllowed")
	tryRegister(&_SecKeychainGetVersion, lib, "SecKeychainGetVersion")
	tryRegister(&_SecKeychainItemCopyAccess, lib, "SecKeychainItemCopyAccess")
	tryRegister(&_SecKeychainItemCopyAttributesAndData, lib, "SecKeychainItemCopyAttributesAndData")
	tryRegister(&_SecKeychainItemCopyContent, lib, "SecKeychainItemCopyContent")
	tryRegister(&_SecKeychainItemCopyFromPersistentReference, lib, "SecKeychainItemCopyFromPersistentReference")
	tryRegister(&_SecKeychainItemCopyKeychain, lib, "SecKeychainItemCopyKeychain")
	tryRegister(&_SecKeychainItemCreateCopy, lib, "SecKeychainItemCreateCopy")
	tryRegister(&_SecKeychainItemCreateFromContent, lib, "SecKeychainItemCreateFromContent")
	tryRegister(&_SecKeychainItemCreatePersistentReference, lib, "SecKeychainItemCreatePersistentReference")
	tryRegister(&_SecKeychainItemDelete, lib, "SecKeychainItemDelete")
	tryRegister(&_SecKeychainItemFreeAttributesAndData, lib, "SecKeychainItemFreeAttributesAndData")
	tryRegister(&_SecKeychainItemFreeContent, lib, "SecKeychainItemFreeContent")
	tryRegister(&_SecKeychainItemGetTypeID, lib, "SecKeychainItemGetTypeID")
	tryRegister(&_SecKeychainItemModifyAttributesAndData, lib, "SecKeychainItemModifyAttributesAndData")
	tryRegister(&_SecKeychainItemModifyContent, lib, "SecKeychainItemModifyContent")
	tryRegister(&_SecKeychainItemSetAccess, lib, "SecKeychainItemSetAccess")
	tryRegister(&_SecKeychainLock, lib, "SecKeychainLock")
	tryRegister(&_SecKeychainLockAll, lib, "SecKeychainLockAll")
	tryRegister(&_SecKeychainOpen, lib, "SecKeychainOpen")
	tryRegister(&_SecKeychainRemoveCallback, lib, "SecKeychainRemoveCallback")
	tryRegister(&_SecKeychainSetAccess, lib, "SecKeychainSetAccess")
	tryRegister(&_SecKeychainSetDefault, lib, "SecKeychainSetDefault")
	tryRegister(&_SecKeychainSetDomainDefault, lib, "SecKeychainSetDomainDefault")
	tryRegister(&_SecKeychainSetDomainSearchList, lib, "SecKeychainSetDomainSearchList")
	tryRegister(&_SecKeychainSetPreferenceDomain, lib, "SecKeychainSetPreferenceDomain")
	tryRegister(&_SecKeychainSetSearchList, lib, "SecKeychainSetSearchList")
	tryRegister(&_SecKeychainSetSettings, lib, "SecKeychainSetSettings")
	tryRegister(&_SecKeychainSetUserInteractionAllowed, lib, "SecKeychainSetUserInteractionAllowed")
	tryRegister(&_SecKeychainUnlock, lib, "SecKeychainUnlock")
	tryRegister(&_SecKeyCopyAttributes, lib, "SecKeyCopyAttributes")
	tryRegister(&_SecKeyCopyExternalRepresentation, lib, "SecKeyCopyExternalRepresentation")
	tryRegister(&_SecKeyCopyKeyExchangeResult, lib, "SecKeyCopyKeyExchangeResult")
	tryRegister(&_SecKeyCopyPublicKey, lib, "SecKeyCopyPublicKey")
	tryRegister(&_SecKeyCreateDecryptedData, lib, "SecKeyCreateDecryptedData")
	tryRegister(&_SecKeyCreateEncryptedData, lib, "SecKeyCreateEncryptedData")
	tryRegister(&_SecKeyCreateFromData, lib, "SecKeyCreateFromData")
	tryRegister(&_SecKeyCreateRandomKey, lib, "SecKeyCreateRandomKey")
	tryRegister(&_SecKeyCreateSignature, lib, "SecKeyCreateSignature")
	tryRegister(&_SecKeyCreateWithData, lib, "SecKeyCreateWithData")
	tryRegister(&_SecKeyDecrypt, lib, "SecKeyDecrypt")
	tryRegister(&_SecKeyDeriveFromPassword, lib, "SecKeyDeriveFromPassword")
	tryRegister(&_SecKeyEncrypt, lib, "SecKeyEncrypt")
	tryRegister(&_SecKeyGeneratePair, lib, "SecKeyGeneratePair")
	tryRegister(&_SecKeyGeneratePairAsync, lib, "SecKeyGeneratePairAsync")
	tryRegister(&_SecKeyGenerateSymmetric, lib, "SecKeyGenerateSymmetric")
	tryRegister(&_SecKeyGetBlockSize, lib, "SecKeyGetBlockSize")
	tryRegister(&_SecKeyGetTypeID, lib, "SecKeyGetTypeID")
	tryRegister(&_SecKeyIsAlgorithmSupported, lib, "SecKeyIsAlgorithmSupported")
	tryRegister(&_SecKeyRawSign, lib, "SecKeyRawSign")
	tryRegister(&_SecKeyRawVerify, lib, "SecKeyRawVerify")
	tryRegister(&_SecKeyUnwrapSymmetric, lib, "SecKeyUnwrapSymmetric")
	tryRegister(&_SecKeyVerifySignature, lib, "SecKeyVerifySignature")
	tryRegister(&_SecKeyWrapSymmetric, lib, "SecKeyWrapSymmetric")
	tryRegister(&_SecPKCS12Import, lib, "SecPKCS12Import")
	tryRegister(&_SecPolicyCopyProperties, lib, "SecPolicyCopyProperties")
	tryRegister(&_SecPolicyCreateBasicX509, lib, "SecPolicyCreateBasicX509")
	tryRegister(&_SecPolicyCreateRevocation, lib, "SecPolicyCreateRevocation")
	tryRegister(&_SecPolicyCreateSSL, lib, "SecPolicyCreateSSL")
	tryRegister(&_SecPolicyCreateWithProperties, lib, "SecPolicyCreateWithProperties")
	tryRegister(&_SecPolicyGetTypeID, lib, "SecPolicyGetTypeID")
	tryRegister(&_SecRandomCopyBytes, lib, "SecRandomCopyBytes")
	tryRegister(&_SecRequestSharedWebCredential, lib, "SecRequestSharedWebCredential")
	tryRegister(&_SecRequirementCopyData, lib, "SecRequirementCopyData")
	tryRegister(&_SecRequirementCopyString, lib, "SecRequirementCopyString")
	tryRegister(&_SecRequirementCreateWithData, lib, "SecRequirementCreateWithData")
	tryRegister(&_SecRequirementCreateWithString, lib, "SecRequirementCreateWithString")
	tryRegister(&_SecRequirementCreateWithStringAndErrors, lib, "SecRequirementCreateWithStringAndErrors")
	tryRegister(&_SecRequirementGetTypeID, lib, "SecRequirementGetTypeID")
	tryRegister(&_SecSignTransformCreate, lib, "SecSignTransformCreate")
	tryRegister(&_SecStaticCodeCheckValidity, lib, "SecStaticCodeCheckValidity")
	tryRegister(&_SecStaticCodeCheckValidityWithErrors, lib, "SecStaticCodeCheckValidityWithErrors")
	tryRegister(&_SecStaticCodeCreateWithPath, lib, "SecStaticCodeCreateWithPath")
	tryRegister(&_SecStaticCodeCreateWithPathAndAttributes, lib, "SecStaticCodeCreateWithPathAndAttributes")
	tryRegister(&_SecStaticCodeGetTypeID, lib, "SecStaticCodeGetTypeID")
	tryRegister(&_SecTaskCopySigningIdentifier, lib, "SecTaskCopySigningIdentifier")
	tryRegister(&_SecTaskCopyValueForEntitlement, lib, "SecTaskCopyValueForEntitlement")
	tryRegister(&_SecTaskCopyValuesForEntitlements, lib, "SecTaskCopyValuesForEntitlements")
	tryRegister(&_SecTaskCreateFromSelf, lib, "SecTaskCreateFromSelf")
	tryRegister(&_SecTaskCreateWithAuditToken, lib, "SecTaskCreateWithAuditToken")
	tryRegister(&_SecTaskGetTypeID, lib, "SecTaskGetTypeID")
	tryRegister(&_SecTransformConnectTransforms, lib, "SecTransformConnectTransforms")
	tryRegister(&_SecTransformCopyExternalRepresentation, lib, "SecTransformCopyExternalRepresentation")
	tryRegister(&_SecTransformCreate, lib, "SecTransformCreate")
	tryRegister(&_SecTransformCreateFromExternalRepresentation, lib, "SecTransformCreateFromExternalRepresentation")
	tryRegister(&_SecTransformCreateGroupTransform, lib, "SecTransformCreateGroupTransform")
	tryRegister(&_SecTransformCreateReadTransformWithReadStream, lib, "SecTransformCreateReadTransformWithReadStream")
	tryRegister(&_SecTransformCustomGetAttribute, lib, "SecTransformCustomGetAttribute")
	tryRegister(&_SecTransformCustomSetAttribute, lib, "SecTransformCustomSetAttribute")
	tryRegister(&_SecTransformExecute, lib, "SecTransformExecute")
	tryRegister(&_SecTransformExecuteAsync, lib, "SecTransformExecuteAsync")
	tryRegister(&_SecTransformFindByName, lib, "SecTransformFindByName")
	tryRegister(&_SecTransformGetAttribute, lib, "SecTransformGetAttribute")
	tryRegister(&_SecTransformGetTypeID, lib, "SecTransformGetTypeID")
	tryRegister(&_SecTransformNoData, lib, "SecTransformNoData")
	tryRegister(&_SecTransformPushbackAttribute, lib, "SecTransformPushbackAttribute")
	tryRegister(&_SecTransformRegister, lib, "SecTransformRegister")
	tryRegister(&_SecTransformSetAttribute, lib, "SecTransformSetAttribute")
	tryRegister(&_SecTransformSetAttributeAction, lib, "SecTransformSetAttributeAction")
	tryRegister(&_SecTransformSetDataAction, lib, "SecTransformSetDataAction")
	tryRegister(&_SecTransformSetTransformAction, lib, "SecTransformSetTransformAction")
	tryRegister(&_SecTrustCopyAnchorCertificates, lib, "SecTrustCopyAnchorCertificates")
	tryRegister(&_SecTrustCopyCustomAnchorCertificates, lib, "SecTrustCopyCustomAnchorCertificates")
	tryRegister(&_SecTrustCopyExceptions, lib, "SecTrustCopyExceptions")
	tryRegister(&_SecTrustCopyPolicies, lib, "SecTrustCopyPolicies")
	tryRegister(&_SecTrustCopyProperties, lib, "SecTrustCopyProperties")
	tryRegister(&_SecTrustCopyPublicKey, lib, "SecTrustCopyPublicKey")
	tryRegister(&_SecTrustCopyResult, lib, "SecTrustCopyResult")
	tryRegister(&_SecTrustCreateWithCertificates, lib, "SecTrustCreateWithCertificates")
	tryRegister(&_SecTrustedApplicationCopyData, lib, "SecTrustedApplicationCopyData")
	tryRegister(&_SecTrustedApplicationCreateFromPath, lib, "SecTrustedApplicationCreateFromPath")
	tryRegister(&_SecTrustedApplicationGetTypeID, lib, "SecTrustedApplicationGetTypeID")
	tryRegister(&_SecTrustedApplicationSetData, lib, "SecTrustedApplicationSetData")
	tryRegister(&_SecTrustEvaluate, lib, "SecTrustEvaluate")
	tryRegister(&_SecTrustEvaluateAsync, lib, "SecTrustEvaluateAsync")
	tryRegister(&_SecTrustEvaluateAsyncWithError, lib, "SecTrustEvaluateAsyncWithError")
	tryRegister(&_SecTrustEvaluateWithError, lib, "SecTrustEvaluateWithError")
	tryRegister(&_SecTrustGetCertificateAtIndex, lib, "SecTrustGetCertificateAtIndex")
	tryRegister(&_SecTrustGetCertificateCount, lib, "SecTrustGetCertificateCount")
	tryRegister(&_SecTrustGetNetworkFetchAllowed, lib, "SecTrustGetNetworkFetchAllowed")
	tryRegister(&_SecTrustGetTrustResult, lib, "SecTrustGetTrustResult")
	tryRegister(&_SecTrustGetTypeID, lib, "SecTrustGetTypeID")
	tryRegister(&_SecTrustGetVerifyTime, lib, "SecTrustGetVerifyTime")
	tryRegister(&_SecTrustSetAnchorCertificates, lib, "SecTrustSetAnchorCertificates")
	tryRegister(&_SecTrustSetAnchorCertificatesOnly, lib, "SecTrustSetAnchorCertificatesOnly")
	tryRegister(&_SecTrustSetExceptions, lib, "SecTrustSetExceptions")
	tryRegister(&_SecTrustSetKeychains, lib, "SecTrustSetKeychains")
	tryRegister(&_SecTrustSetNetworkFetchAllowed, lib, "SecTrustSetNetworkFetchAllowed")
	tryRegister(&_SecTrustSetOCSPResponse, lib, "SecTrustSetOCSPResponse")
	tryRegister(&_SecTrustSetOptions, lib, "SecTrustSetOptions")
	tryRegister(&_SecTrustSetPolicies, lib, "SecTrustSetPolicies")
	tryRegister(&_SecTrustSetSignedCertificateTimestamps, lib, "SecTrustSetSignedCertificateTimestamps")
	tryRegister(&_SecTrustSettingsCopyCertificates, lib, "SecTrustSettingsCopyCertificates")
	tryRegister(&_SecTrustSettingsCopyModificationDate, lib, "SecTrustSettingsCopyModificationDate")
	tryRegister(&_SecTrustSettingsCopyTrustSettings, lib, "SecTrustSettingsCopyTrustSettings")
	tryRegister(&_SecTrustSettingsCreateExternalRepresentation, lib, "SecTrustSettingsCreateExternalRepresentation")
	tryRegister(&_SecTrustSettingsImportExternalRepresentation, lib, "SecTrustSettingsImportExternalRepresentation")
	tryRegister(&_SecTrustSettingsRemoveTrustSettings, lib, "SecTrustSettingsRemoveTrustSettings")
	tryRegister(&_SecTrustSettingsSetTrustSettings, lib, "SecTrustSettingsSetTrustSettings")
	tryRegister(&_SecTrustSetVerifyDate, lib, "SecTrustSetVerifyDate")
	tryRegister(&_SecVerifyTransformCreate, lib, "SecVerifyTransformCreate")
	tryRegister(&_SessionCreate, lib, "SessionCreate")
	tryRegister(&_SessionGetInfo, lib, "SessionGetInfo")
	tryRegister(&_SSLAddDistinguishedName, lib, "SSLAddDistinguishedName")
	tryRegister(&_SSLClose, lib, "SSLClose")
	tryRegister(&_SSLContextGetTypeID, lib, "SSLContextGetTypeID")
	tryRegister(&_SSLCopyALPNProtocols, lib, "SSLCopyALPNProtocols")
	tryRegister(&_SSLCopyCertificateAuthorities, lib, "SSLCopyCertificateAuthorities")
	tryRegister(&_SSLCopyDistinguishedNames, lib, "SSLCopyDistinguishedNames")
	tryRegister(&_SSLCopyPeerTrust, lib, "SSLCopyPeerTrust")
	tryRegister(&_SSLCopyRequestedPeerName, lib, "SSLCopyRequestedPeerName")
	tryRegister(&_SSLCopyRequestedPeerNameLength, lib, "SSLCopyRequestedPeerNameLength")
	tryRegister(&_SSLCreateContext, lib, "SSLCreateContext")
	tryRegister(&_SSLGetBufferedReadSize, lib, "SSLGetBufferedReadSize")
	tryRegister(&_SSLGetClientCertificateState, lib, "SSLGetClientCertificateState")
	tryRegister(&_SSLGetConnection, lib, "SSLGetConnection")
	tryRegister(&_SSLGetDatagramWriteSize, lib, "SSLGetDatagramWriteSize")
	tryRegister(&_SSLGetDiffieHellmanParams, lib, "SSLGetDiffieHellmanParams")
	tryRegister(&_SSLGetEnabledCiphers, lib, "SSLGetEnabledCiphers")
	tryRegister(&_SSLGetMaxDatagramRecordSize, lib, "SSLGetMaxDatagramRecordSize")
	tryRegister(&_SSLGetNegotiatedCipher, lib, "SSLGetNegotiatedCipher")
	tryRegister(&_SSLGetNegotiatedProtocolVersion, lib, "SSLGetNegotiatedProtocolVersion")
	tryRegister(&_SSLGetNumberEnabledCiphers, lib, "SSLGetNumberEnabledCiphers")
	tryRegister(&_SSLGetNumberSupportedCiphers, lib, "SSLGetNumberSupportedCiphers")
	tryRegister(&_SSLGetPeerDomainName, lib, "SSLGetPeerDomainName")
	tryRegister(&_SSLGetPeerDomainNameLength, lib, "SSLGetPeerDomainNameLength")
	tryRegister(&_SSLGetPeerID, lib, "SSLGetPeerID")
	tryRegister(&_SSLGetProtocolVersionMax, lib, "SSLGetProtocolVersionMax")
	tryRegister(&_SSLGetProtocolVersionMin, lib, "SSLGetProtocolVersionMin")
	tryRegister(&_SSLGetSessionOption, lib, "SSLGetSessionOption")
	tryRegister(&_SSLGetSessionState, lib, "SSLGetSessionState")
	tryRegister(&_SSLGetSupportedCiphers, lib, "SSLGetSupportedCiphers")
	tryRegister(&_SSLHandshake, lib, "SSLHandshake")
	tryRegister(&_SSLRead, lib, "SSLRead")
	tryRegister(&_SSLReHandshake, lib, "SSLReHandshake")
	tryRegister(&_SSLSetALPNProtocols, lib, "SSLSetALPNProtocols")
	tryRegister(&_SSLSetCertificate, lib, "SSLSetCertificate")
	tryRegister(&_SSLSetCertificateAuthorities, lib, "SSLSetCertificateAuthorities")
	tryRegister(&_SSLSetClientSideAuthenticate, lib, "SSLSetClientSideAuthenticate")
	tryRegister(&_SSLSetConnection, lib, "SSLSetConnection")
	tryRegister(&_SSLSetDatagramHelloCookie, lib, "SSLSetDatagramHelloCookie")
	tryRegister(&_SSLSetDiffieHellmanParams, lib, "SSLSetDiffieHellmanParams")
	tryRegister(&_SSLSetEnabledCiphers, lib, "SSLSetEnabledCiphers")
	tryRegister(&_SSLSetEncryptionCertificate, lib, "SSLSetEncryptionCertificate")
	tryRegister(&_SSLSetError, lib, "SSLSetError")
	tryRegister(&_SSLSetIOFuncs, lib, "SSLSetIOFuncs")
	tryRegister(&_SSLSetMaxDatagramRecordSize, lib, "SSLSetMaxDatagramRecordSize")
	tryRegister(&_SSLSetOCSPResponse, lib, "SSLSetOCSPResponse")
	tryRegister(&_SSLSetPeerDomainName, lib, "SSLSetPeerDomainName")
	tryRegister(&_SSLSetPeerID, lib, "SSLSetPeerID")
	tryRegister(&_SSLSetProtocolVersionMax, lib, "SSLSetProtocolVersionMax")
	tryRegister(&_SSLSetProtocolVersionMin, lib, "SSLSetProtocolVersionMin")
	tryRegister(&_SSLSetSessionConfig, lib, "SSLSetSessionConfig")
	tryRegister(&_SSLSetSessionOption, lib, "SSLSetSessionOption")
	tryRegister(&_SSLSetSessionTicketsEnabled, lib, "SSLSetSessionTicketsEnabled")
	tryRegister(&_SSLWrite, lib, "SSLWrite")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the application list, description, and CSSM prompt selector for a given access control list entry.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the application list, description, and CSSM prompt selector for a given access control list entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCopySimpleContents
func SecACLCopySimpleContents(acl SecACLRef, applicationList unsafe.Pointer, description unsafe.Pointer, promptSelector unsafe.Pointer) unsafe.Pointer {
	return _SecACLCopySimpleContents(acl, applicationList, description, promptSelector)
}/* debug [functions.gen.go/function]: SecACLCopySimpleContents */

// Creates a new access control list entry from the application list, description, and prompt selector provided and adds it to an item’s access object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a new access control list entry from the application list, description, and prompt selector provided and adds it to an item’s access object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCreateFromSimpleContents
func SecACLCreateFromSimpleContents(access SecAccessRef, applicationList ArrayRef, description StringRef, promptSelector unsafe.Pointer, newAcl unsafe.Pointer) unsafe.Pointer {
	return _SecACLCreateFromSimpleContents(access, applicationList, description, promptSelector, newAcl)
}/* debug [functions.gen.go/function]: SecACLCreateFromSimpleContents */

// Retrieves the CSSM authorization tags of a given access control list entry.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the CSSM authorization tags of a given access control list entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLGetAuthorizations
func SecACLGetAuthorizations(acl SecACLRef, tags unsafe.Pointer, tagCount unsafe.Pointer) unsafe.Pointer {
	return _SecACLGetAuthorizations(acl, tags, tagCount)
}/* debug [functions.gen.go/function]: SecACLGetAuthorizations */

// Sets the CSSM authorization tags for a given access control list entry.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Sets the CSSM authorization tags for a given access control list entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLSetAuthorizations
func SecACLSetAuthorizations(acl SecACLRef, tags unsafe.Pointer, tagCount unsafe.Pointer) unsafe.Pointer {
	return _SecACLSetAuthorizations(acl, tags, tagCount)
}/* debug [functions.gen.go/function]: SecACLSetAuthorizations */

// Sets the application list, description, and prompt selector for a given access control list entry.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Sets the application list, description, and prompt selector for a given access control list entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLSetSimpleContents
func SecACLSetSimpleContents(acl SecACLRef, applicationList ArrayRef, description StringRef, promptSelector unsafe.Pointer) unsafe.Pointer {
	return _SecACLSetSimpleContents(acl, applicationList, description, promptSelector)
}/* debug [functions.gen.go/function]: SecACLSetSimpleContents */

// Retrieves selected access control lists from a given access object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves selected access control lists from a given access object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopySelectedACLList
func SecAccessCopySelectedACLList(accessRef SecAccessRef, action CSSM_ACL_AUTHORIZATION_TAG, aclList unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopySelectedACLList(accessRef, action, aclList)
}/* debug [functions.gen.go/function]: SecAccessCopySelectedACLList */

// Creates a new access object using the owner and access control list you provide.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a new access object using the owner and access control list you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCreateFromOwnerAndACL
func SecAccessCreateFromOwnerAndACL(owner unsafe.Pointer, aclCount unsafe.Pointer, acls unsafe.Pointer, accessRef unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCreateFromOwnerAndACL(owner, aclCount, acls, accessRef)
}/* debug [functions.gen.go/function]: SecAccessCreateFromOwnerAndACL */

// Retrieves the owner and the access control list of a given access object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the owner and the access control list of a given access object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessGetOwnerAndACL
func SecAccessGetOwnerAndACL(accessRef SecAccessRef, owner unsafe.Pointer, aclCount unsafe.Pointer, acls unsafe.Pointer) unsafe.Pointer {
	return _SecAccessGetOwnerAndACL(accessRef, owner, aclCount, acls)
}/* debug [functions.gen.go/function]: SecAccessGetOwnerAndACL */

// SecCertificateCopyNotValidAfterDate is a Security function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidAfterDate(_:)
func SecCertificateCopyNotValidAfterDate(certificate SecCertificateRef) DateRef {
	return _SecCertificateCopyNotValidAfterDate(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyNotValidAfterDate */

// Retrieves the preferred certificate for the specified name and key use.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the preferred certificate for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyPreference
func SecCertificateCopyPreference(name StringRef, keyUsage unsafe.Pointer, certificate unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyPreference(name, keyUsage, certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyPreference */

// Creates a certificate object based on the specified data, type, and encoding.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a certificate object based on the specified data, type, and encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCreateFromData
func SecCertificateCreateFromData(data unsafe.Pointer, type_ CSSM_CERT_TYPE, encoding CSSM_CERT_ENCODING, certificate unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCreateFromData(data, type_, encoding, certificate)
}/* debug [functions.gen.go/function]: SecCertificateCreateFromData */

// Retrieves the algorithm identifier for a certificate.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the algorithm identifier for a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetAlgorithmID
func SecCertificateGetAlgorithmID(certificate SecCertificateRef, algid unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetAlgorithmID(certificate, algid)
}/* debug [functions.gen.go/function]: SecCertificateGetAlgorithmID */

// Retrieves the certificate library handle from a certificate object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the certificate library handle from a certificate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetCLHandle
func SecCertificateGetCLHandle(certificate SecCertificateRef, clHandle unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetCLHandle(certificate, clHandle)
}/* debug [functions.gen.go/function]: SecCertificateGetCLHandle */

// Retrieves the data for a certificate.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the data for a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetData
func SecCertificateGetData(certificate SecCertificateRef, data unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetData(certificate, data)
}/* debug [functions.gen.go/function]: SecCertificateGetData */

// Unsupported.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Unsupported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetIssuer
func SecCertificateGetIssuer(certificate SecCertificateRef, issuer unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetIssuer(certificate, issuer)
}/* debug [functions.gen.go/function]: SecCertificateGetIssuer */

// Unsupported.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Unsupported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetSubject
func SecCertificateGetSubject(certificate SecCertificateRef, subject unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetSubject(certificate, subject)
}/* debug [functions.gen.go/function]: SecCertificateGetSubject */

// Retrieves the type of a specified certificate.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves the type of a specified certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetType
func SecCertificateGetType(certificate SecCertificateRef, certificateType unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateGetType(certificate, certificateType)
}/* debug [functions.gen.go/function]: SecCertificateGetType */

// Sets the preferred certificate for a specified name, key use, and date.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Sets the preferred certificate for a specified name, key use, and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateSetPreference
func SecCertificateSetPreference(certificate SecCertificateRef, name StringRef, keyUsage unsafe.Pointer, date DateRef) unsafe.Pointer {
	return _SecCertificateSetPreference(certificate, name, keyUsage, date)
}/* debug [functions.gen.go/function]: SecCertificateSetPreference */

// Returns the preferred identity for the specified name and key use.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the preferred identity for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCopyPreference
func SecIdentityCopyPreference(name StringRef, keyUsage CSSM_KEYUSE, validIssuers ArrayRef, identity unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCopyPreference(name, keyUsage, validIssuers, identity)
}/* debug [functions.gen.go/function]: SecIdentityCopyPreference */

// SecIdentityCreate is a Security function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCreate(_:_:_:)
func SecIdentityCreate(allocator AllocatorRef, certificate SecCertificateRef, privateKey SecKeyRef) SecIdentityRef {
	return _SecIdentityCreate(allocator, certificate, privateKey)
}/* debug [functions.gen.go/function]: SecIdentityCreate */

// Finds the next identity matching specified search criteria
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Finds the next identity matching specified search criteria
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySearchCopyNext
func SecIdentitySearchCopyNext(searchRef SecIdentitySearchRef, identity unsafe.Pointer) unsafe.Pointer {
	return _SecIdentitySearchCopyNext(searchRef, identity)
}/* debug [functions.gen.go/function]: SecIdentitySearchCopyNext */

// Creates a search object for finding identities.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a search object for finding identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySearchCreate
func SecIdentitySearchCreate(keychainOrArray TypeRef, keyUsage CSSM_KEYUSE, searchRef unsafe.Pointer) unsafe.Pointer {
	return _SecIdentitySearchCreate(keychainOrArray, keyUsage, searchRef)
}/* debug [functions.gen.go/function]: SecIdentitySearchCreate */

// Returns the unique identifier of the opaque type to which a object belongs.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySearchGetTypeID
func SecIdentitySearchGetTypeID() TypeID {
	return _SecIdentitySearchGetTypeID()
}/* debug [functions.gen.go/function]: SecIdentitySearchGetTypeID */

// Sets the preferred identity for the specified name and key use.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Sets the preferred identity for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySetPreference
func SecIdentitySetPreference(identity SecIdentityRef, name StringRef, keyUsage CSSM_KEYUSE) unsafe.Pointer {
	return _SecIdentitySetPreference(identity, name, keyUsage)
}/* debug [functions.gen.go/function]: SecIdentitySetPreference */

// Creates an asymmetric key pair and stores it in a keychain.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates an asymmetric key pair and stores it in a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreatePair
func SecKeyCreatePair(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits unsafe.Pointer, contextHandle CSSM_CC_HANDLE, publicKeyUsage CSSM_KEYUSE, publicKeyAttr unsafe.Pointer, privateKeyUsage CSSM_KEYUSE, privateKeyAttr unsafe.Pointer, initialAccess SecAccessRef, publicKey unsafe.Pointer, privateKey unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCreatePair(keychainRef, algorithm, keySizeInBits, contextHandle, publicKeyUsage, publicKeyAttr, privateKeyUsage, privateKeyAttr, initialAccess, publicKey, privateKey)
}/* debug [functions.gen.go/function]: SecKeyCreatePair */

// Creates a symmetric key and optionally stores it in a keychain.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a symmetric key and optionally stores it in a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGenerate
func SecKeyGenerate(keychainRef SecKeychainRef, algorithm CSSM_ALGORITHMS, keySizeInBits unsafe.Pointer, contextHandle CSSM_CC_HANDLE, keyUsage CSSM_KEYUSE, keyAttr unsafe.Pointer, initialAccess SecAccessRef, keyRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeyGenerate(keychainRef, algorithm, keySizeInBits, contextHandle, keyUsage, keyAttr, initialAccess, keyRef)
}/* debug [functions.gen.go/function]: SecKeyGenerate */

// Returns the CSSM CSP handle for a key.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the CSSM CSP handle for a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGetCSPHandle
func SecKeyGetCSPHandle(keyRef SecKeyRef, cspHandle unsafe.Pointer) unsafe.Pointer {
	return _SecKeyGetCSPHandle(keyRef, cspHandle)
}/* debug [functions.gen.go/function]: SecKeyGetCSPHandle */

// Retrieves a pointer to the structure containing the key stored in a keychain item.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves a pointer to the structure containing the key stored in a keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGetCSSMKey
func SecKeyGetCSSMKey(key SecKeyRef, cssmKey unsafe.Pointer) unsafe.Pointer {
	return _SecKeyGetCSSMKey(key, cssmKey)
}/* debug [functions.gen.go/function]: SecKeyGetCSSMKey */

// Returns an access credential for a key.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns an access credential for a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGetCredentials
func SecKeyGetCredentials(keyRef SecKeyRef, operation CSSM_ACL_AUTHORIZATION_TAG, credentialType SecCredentialType, outCredentials unsafe.Pointer) unsafe.Pointer {
	return _SecKeyGetCredentials(keyRef, operation, credentialType, outCredentials)
}/* debug [functions.gen.go/function]: SecKeyGetCredentials */

// Returns the CSSM CSP handle for the given keychain object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the CSSM CSP handle for the given keychain object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetCSPHandle
func SecKeychainGetCSPHandle(keychain SecKeychainRef, cspHandle unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetCSPHandle(keychain, cspHandle)
}/* debug [functions.gen.go/function]: SecKeychainGetCSPHandle */

// Returns the CSSM database handle for a given keychain object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the CSSM database handle for a given keychain object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetDLDBHandle
func SecKeychainGetDLDBHandle(keychain SecKeychainRef, dldbHandle unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetDLDBHandle(keychain, dldbHandle)
}/* debug [functions.gen.go/function]: SecKeychainGetDLDBHandle */

// Exports one or more certificates, keys, or identities.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Exports one or more certificates, keys, or identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemExport
func SecKeychainItemExport(keychainItemOrArray TypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams unsafe.Pointer, exportedData unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemExport(keychainItemOrArray, outputFormat, flags, keyParams, exportedData)
}/* debug [functions.gen.go/function]: SecKeychainItemExport */

// Returns the CSSM database handle for a given keychain item object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the CSSM database handle for a given keychain item object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemGetDLDBHandle
func SecKeychainItemGetDLDBHandle(keyItemRef SecKeychainItemRef, dldbHandle unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemGetDLDBHandle(keyItemRef, dldbHandle)
}/* debug [functions.gen.go/function]: SecKeychainItemGetDLDBHandle */

// Returns a CSSM unique record for the given keychain item object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns a CSSM unique record for the given keychain item object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemGetUniqueRecordID
func SecKeychainItemGetUniqueRecordID(itemRef SecKeychainItemRef, uniqueRecordID unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemGetUniqueRecordID(itemRef, uniqueRecordID)
}/* debug [functions.gen.go/function]: SecKeychainItemGetUniqueRecordID */

// Imports one or more certificates, keys, or identities and adds them to a keychain.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Imports one or more certificates, keys, or identities and adds them to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemImport
func SecKeychainItemImport(importedData DataRef, fileNameOrExtension StringRef, inputFormat unsafe.Pointer, itemType unsafe.Pointer, flags SecItemImportExportFlags, keyParams unsafe.Pointer, importKeychain SecKeychainRef, outItems unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems)
}/* debug [functions.gen.go/function]: SecKeychainItemImport */

// Finds the next keychain item matching the given search criteria.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Finds the next keychain item matching the given search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSearchCopyNext
func SecKeychainSearchCopyNext(searchRef SecKeychainSearchRef, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainSearchCopyNext(searchRef, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainSearchCopyNext */

// Creates a search object matching a list of zero or more attributes.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a search object matching a list of zero or more attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSearchCreateFromAttributes
func SecKeychainSearchCreateFromAttributes(keychainOrArray TypeRef, itemClass SecItemClass, attrList unsafe.Pointer, searchRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainSearchCreateFromAttributes(keychainOrArray, itemClass, attrList, searchRef)
}/* debug [functions.gen.go/function]: SecKeychainSearchCreateFromAttributes */

// Returns the unique identifier of the opaque type to which a keychain search object belongs.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a keychain search object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSearchGetTypeID
func SecKeychainSearchGetTypeID() TypeID {
	return _SecKeychainSearchGetTypeID()
}/* debug [functions.gen.go/function]: SecKeychainSearchGetTypeID */

// Returns a policy object for the specified policy type object identifier.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.
// Returns a policy object for the specified policy type object identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateWithOID
func SecPolicyCreateWithOID(policyOID TypeRef) SecPolicyRef {
	return _SecPolicyCreateWithOID(policyOID)
}/* debug [functions.gen.go/function]: SecPolicyCreateWithOID */

// Retrieves a policy’s object identifier.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves a policy’s object identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyGetOID
func SecPolicyGetOID(policyRef SecPolicyRef, oid unsafe.Pointer) unsafe.Pointer {
	return _SecPolicyGetOID(policyRef, oid)
}/* debug [functions.gen.go/function]: SecPolicyGetOID */

// Retrieves the trust policy handle for a policy object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves the trust policy handle for a policy object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyGetTPHandle
func SecPolicyGetTPHandle(policyRef SecPolicyRef, tpHandle unsafe.Pointer) unsafe.Pointer {
	return _SecPolicyGetTPHandle(policyRef, tpHandle)
}/* debug [functions.gen.go/function]: SecPolicyGetTPHandle */

// Retrieves a policy’s value.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves a policy’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyGetValue
func SecPolicyGetValue(policyRef SecPolicyRef, value unsafe.Pointer) unsafe.Pointer {
	return _SecPolicyGetValue(policyRef, value)
}/* debug [functions.gen.go/function]: SecPolicyGetValue */

// Retrieves a policy object for the next policy matching specified search criteria.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Retrieves a policy object for the next policy matching specified search criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySearchCopyNext
func SecPolicySearchCopyNext(searchRef SecPolicySearchRef, policyRef unsafe.Pointer) unsafe.Pointer {
	return _SecPolicySearchCopyNext(searchRef, policyRef)
}/* debug [functions.gen.go/function]: SecPolicySearchCopyNext */

// Creates a search object for finding policies.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Creates a search object for finding policies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySearchCreate
func SecPolicySearchCreate(certType CSSM_CERT_TYPE, policyOID unsafe.Pointer, value unsafe.Pointer, searchRef unsafe.Pointer) unsafe.Pointer {
	return _SecPolicySearchCreate(certType, policyOID, value, searchRef)
}/* debug [functions.gen.go/function]: SecPolicySearchCreate */

// Returns the unique identifier of the opaque type to which a object belongs.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySearchGetTypeID
func SecPolicySearchGetTypeID() TypeID {
	return _SecPolicySearchGetTypeID()
}/* debug [functions.gen.go/function]: SecPolicySearchGetTypeID */

// Sets properties for a policy.
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.7.
// Sets properties for a policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySetProperties
func SecPolicySetProperties(policyRef SecPolicyRef, properties DictionaryRef) unsafe.Pointer {
	return _SecPolicySetProperties(policyRef, properties)
}/* debug [functions.gen.go/function]: SecPolicySetProperties */

// Sets a policy’s value.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Sets a policy’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySetValue
func SecPolicySetValue(policyRef SecPolicyRef, value unsafe.Pointer) unsafe.Pointer {
	return _SecPolicySetValue(policyRef, value)
}/* debug [functions.gen.go/function]: SecPolicySetValue */

// Retrieves the CSSM trust result.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves the CSSM trust result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetCssmResult
func SecTrustGetCssmResult(trust SecTrustRef, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetCssmResult(trust, result)
}/* debug [functions.gen.go/function]: SecTrustGetCssmResult */

// Retrieves the CSSM result code from the most recent trust evaluation for a trust management object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves the CSSM result code from the most recent trust evaluation for a trust management object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetCssmResultCode
func SecTrustGetCssmResultCode(trust SecTrustRef, resultCode unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetCssmResultCode(trust, resultCode)
}/* debug [functions.gen.go/function]: SecTrustGetCssmResultCode */

// Retrieves details on the outcome of a call to the function .
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves details on the outcome of a call to the function .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetResult
func SecTrustGetResult(trustRef SecTrustRef, result unsafe.Pointer, certChain unsafe.Pointer, statusChain unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetResult(trustRef, result, certChain, statusChain)
}/* debug [functions.gen.go/function]: SecTrustGetResult */

// Retrieves the trust policy handle.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Retrieves the trust policy handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetTPHandle
func SecTrustGetTPHandle(trust SecTrustRef, handle unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetTPHandle(trust, handle)
}/* debug [functions.gen.go/function]: SecTrustGetTPHandle */

// Sets the action and action data for a trust management object.
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
// Sets the action and action data for a trust management object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetParameters
func SecTrustSetParameters(trustRef SecTrustRef, action CSSM_TP_ACTION, actionData DataRef) unsafe.Pointer {
	return _SecTrustSetParameters(trustRef, action, actionData)
}/* debug [functions.gen.go/function]: SecTrustSetParameters */

// Returns download ticket’s creation date.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
// Returns download ticket’s creation date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCopyCreationDate
func SecureDownloadCopyCreationDate(downloadRef unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCopyCreationDate(downloadRef, date)
}/* debug [functions.gen.go/function]: SecureDownloadCopyCreationDate */

// Retrieves supporting data such as the user name and other information gathered during evaluation of authorization.
//
// Added in macOS 10.0.
// Retrieves supporting data such as the user name and other information gathered during evaluation of authorization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyInfo(_:_:_:)
func AuthorizationCopyInfo(authorization AuthorizationRef, tag AuthorizationString, info unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCopyInfo(authorization, tag, info)
}/* debug [functions.gen.go/function]: AuthorizationCopyInfo */

// Authorizes and preauthorizes rights synchronously.
//
// Added in macOS 10.0.
// Authorizes and preauthorizes rights synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyRights(_:_:_:_:_:)
func AuthorizationCopyRights(authorization AuthorizationRef, rights unsafe.Pointer, environment unsafe.Pointer, flags AuthorizationFlags, authorizedRights unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCopyRights(authorization, rights, environment, flags, authorizedRights)
}/* debug [functions.gen.go/function]: AuthorizationCopyRights */

// Authorizes and preauthorizes rights asynchronously.
//
// Added in macOS 10.7.
// Authorizes and preauthorizes rights asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyRightsAsync(_:_:_:_:_:)
func AuthorizationCopyRightsAsync(authorization AuthorizationRef, rights unsafe.Pointer, environment unsafe.Pointer, flags AuthorizationFlags, callbackBlock unsafe.Pointer) {
	_AuthorizationCopyRightsAsync(authorization, rights, environment, flags, callbackBlock)
}/* debug [functions.gen.go/function]: AuthorizationCopyRightsAsync */

// Creates a new authorization reference and provides an option to authorize or preauthorize rights.
//
// Added in macOS 10.0.
// Creates a new authorization reference and provides an option to authorize or preauthorize rights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCreate(_:_:_:_:)
func AuthorizationCreate(rights unsafe.Pointer, environment unsafe.Pointer, flags AuthorizationFlags, authorization unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCreate(rights, environment, flags, authorization)
}/* debug [functions.gen.go/function]: AuthorizationCreate */

// Internalizes the external representation of an authorization reference.
//
// Added in macOS 10.0.
// Internalizes the external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCreateFromExternalForm(_:_:)
func AuthorizationCreateFromExternalForm(extForm unsafe.Pointer, authorization unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCreateFromExternalForm(extForm, authorization)
}/* debug [functions.gen.go/function]: AuthorizationCreateFromExternalForm */

// Frees the memory associated with an authorization reference.
//
// Added in macOS 10.0.
// Frees the memory associated with an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFree(_:_:)
func AuthorizationFree(authorization AuthorizationRef, flags AuthorizationFlags) unsafe.Pointer {
	return _AuthorizationFree(authorization, flags)
}/* debug [functions.gen.go/function]: AuthorizationFree */

// Frees the memory associated with a set of authorization items.
//
// Added in macOS 10.0.
// Frees the memory associated with a set of authorization items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFreeItemSet(_:)
func AuthorizationFreeItemSet(set unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationFreeItemSet(set)
}/* debug [functions.gen.go/function]: AuthorizationFreeItemSet */

// Creates an external representation of an authorization reference.
//
// Added in macOS 10.0.
// Creates an external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationMakeExternalForm(_:_:)
func AuthorizationMakeExternalForm(authorization AuthorizationRef, extForm unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationMakeExternalForm(authorization, extForm)
}/* debug [functions.gen.go/function]: AuthorizationMakeExternalForm */

// Retrieves a right definition as a dictionary.
//
// Added in macOS 10.0.
// Retrieves a right definition as a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightGet(_:_:)
func AuthorizationRightGet(rightName unsafe.Pointer, rightDefinition unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationRightGet(rightName, rightDefinition)
}/* debug [functions.gen.go/function]: AuthorizationRightGet */

// Removes a right from the policy database.
//
// Added in macOS 10.0.
// Removes a right from the policy database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightRemove(_:_:)
func AuthorizationRightRemove(authRef AuthorizationRef, rightName unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationRightRemove(authRef, rightName)
}/* debug [functions.gen.go/function]: AuthorizationRightRemove */

// Creates or updates a right entry in the policy database.
//
// Added in macOS 10.0.
// Creates or updates a right entry in the policy database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightSet(_:_:_:_:_:_:)
func AuthorizationRightSet(authRef AuthorizationRef, rightName unsafe.Pointer, rightDefinition TypeRef, descriptionKey StringRef, bundle BundleRef, localeTableName StringRef) unsafe.Pointer {
	return _AuthorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName)
}/* debug [functions.gen.go/function]: AuthorizationRightSet */

// Obtains an array of all of the certificates in a message.
//
// Added in macOS 10.5.
// Obtains an array of all of the certificates in a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyAllCerts(_:_:)
func CMSDecoderCopyAllCerts(cmsDecoder SDecoderRef, certsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyAllCerts(cmsDecoder, certsOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopyAllCerts */

// Obtains the message content, if any.
//
// Added in macOS 10.5.
// Obtains the message content, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyContent(_:_:)
func CMSDecoderCopyContent(cmsDecoder SDecoderRef, contentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyContent(cmsDecoder, contentOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopyContent */

// Obtains the detached content specified with the function.
//
// Added in macOS 10.5.
// Obtains the detached content specified with the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyDetachedContent(_:_:)
func CMSDecoderCopyDetachedContent(cmsDecoder SDecoderRef, detachedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopyDetachedContent */

// Obtains the object identifier for the encapsulated data of a signed message.
//
// Added in macOS 10.5.
// Obtains the object identifier for the encapsulated data of a signed message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyEncapsulatedContentType(_:_:)
func CMSDecoderCopyEncapsulatedContentType(cmsDecoder SDecoderRef, eContentTypeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopyEncapsulatedContentType */

// Obtains the certificate of the specified signer of a CMS message.
//
// Added in macOS 10.5.
// Obtains the certificate of the specified signer of a CMS message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerCert(_:_:_:)
func CMSDecoderCopySignerCert(cmsDecoder SDecoderRef, signerIndex uintptr, signerCertOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerCert */

// Obtains the email address of the specified signer of a CMS message.
//
// Added in macOS 10.5.
// Obtains the email address of the specified signer of a CMS message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerEmailAddress(_:_:_:)
func CMSDecoderCopySignerEmailAddress(cmsDecoder SDecoderRef, signerIndex uintptr, signerEmailAddressOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerEmailAddress */

// Obtains the signing time of a CMS message, if present.
//
// Added in macOS 10.8.
// Obtains the signing time of a CMS message, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerSigningTime(_:_:_:)
func CMSDecoderCopySignerSigningTime(cmsDecoder SDecoderRef, signerIndex uintptr, signingTime unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerSigningTime */

// Obtains the status of a CMS message’s signature.
//
// Added in macOS 10.5.
// Obtains the status of a CMS message’s signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerStatus(_:_:_:_:_:_:_:)
func CMSDecoderCopySignerStatus(cmsDecoder SDecoderRef, signerIndex uintptr, policyOrArray TypeRef, evaluateSecTrust unsafe.Pointer, signerStatusOut unsafe.Pointer, secTrustOut unsafe.Pointer, certVerifyResultCodeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerStatus */

// Returns the timestamp of a signer of a CMS message, if present.
//
// Added in macOS 10.8.
// Returns the timestamp of a signer of a CMS message, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestamp(_:_:_:)
func CMSDecoderCopySignerTimestamp(cmsDecoder SDecoderRef, signerIndex uintptr, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerTimestamp */

// Returns an array containing the certificates from a timestamp response.
//
// Added in macOS 10.8.
// Returns an array containing the certificates from a timestamp response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampCertificates(_:_:_:)
func CMSDecoderCopySignerTimestampCertificates(cmsDecoder SDecoderRef, signerIndex uintptr, certificateRefs unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerTimestampCertificates */

// Returns the timestamp of a signer of a CMS message using a given policy, if present.
//
// Added in macOS 10.10.
// Returns the timestamp of a signer of a CMS message using a given policy, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder SDecoderRef, timeStampPolicy TypeRef, signerIndex uintptr, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp)
}/* debug [functions.gen.go/function]: CMSDecoderCopySignerTimestampWithPolicy */

// Creates a CMSDecoder reference.
//
// Added in macOS 10.5.
// Creates a CMSDecoder reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCreate(_:)
func CMSDecoderCreate(cmsDecoderOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCreate(cmsDecoderOut)
}/* debug [functions.gen.go/function]: CMSDecoderCreate */

// Indicates that there is no more data to decode.
//
// Added in macOS 10.5.
// Indicates that there is no more data to decode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderFinalizeMessage(_:)
func CMSDecoderFinalizeMessage(cmsDecoder SDecoderRef) unsafe.Pointer {
	return _CMSDecoderFinalizeMessage(cmsDecoder)
}/* debug [functions.gen.go/function]: CMSDecoderFinalizeMessage */

// Obtains the number of signers of a message.
//
// Added in macOS 10.5.
// Obtains the number of signers of a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderGetNumSigners(_:_:)
func CMSDecoderGetNumSigners(cmsDecoder SDecoderRef, numSignersOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderGetNumSigners(cmsDecoder, numSignersOut)
}/* debug [functions.gen.go/function]: CMSDecoderGetNumSigners */

// Returns the type identifier for the CMSDecoder opaque type.
//
// Added in macOS 10.0.
// Returns the type identifier for the CMSDecoder opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderGetTypeID()
func CMSDecoderGetTypeID() TypeID {
	return _CMSDecoderGetTypeID()
}/* debug [functions.gen.go/function]: CMSDecoderGetTypeID */

// Determines whether a CMS message was encrypted.
//
// Added in macOS 10.5.
// Determines whether a CMS message was encrypted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderIsContentEncrypted(_:_:)
func CMSDecoderIsContentEncrypted(cmsDecoder SDecoderRef, isEncryptedOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut)
}/* debug [functions.gen.go/function]: CMSDecoderIsContentEncrypted */

// Specifies the message’s detached content, if any.
//
// Added in macOS 10.5.
// Specifies the message’s detached content, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderSetDetachedContent(_:_:)
func CMSDecoderSetDetachedContent(cmsDecoder SDecoderRef, detachedContent DataRef) unsafe.Pointer {
	return _CMSDecoderSetDetachedContent(cmsDecoder, detachedContent)
}/* debug [functions.gen.go/function]: CMSDecoderSetDetachedContent */

// Specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
// Specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderSetSearchKeychain(_:_:)
func CMSDecoderSetSearchKeychain(cmsDecoder SDecoderRef, keychainOrArray TypeRef) unsafe.Pointer {
	return _CMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray)
}/* debug [functions.gen.go/function]: CMSDecoderSetSearchKeychain */

// Feeds raw bytes of the message to be decoded into the decoder.
//
// Added in macOS 10.5.
// Feeds raw bytes of the message to be decoded into the decoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderUpdateMessage(_:_:_:)
func CMSDecoderUpdateMessage(cmsDecoder SDecoderRef, msgBytes unsafe.Pointer, msgBytesLen uintptr) unsafe.Pointer {
	return _CMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen)
}/* debug [functions.gen.go/function]: CMSDecoderUpdateMessage */

// Encodes a message and obtains the result in one high-level function call.
//
// Added in macOS 10.7.
// Encodes a message and obtains the result in one high-level function call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncodeContent(_:_:_:_:_:_:_:_:)
func CMSEncodeContent(signers TypeRef, recipients TypeRef, eContentTypeOID TypeRef, detachedContent unsafe.Pointer, signedAttributes SSignedAttributes, content unsafe.Pointer, contentLen uintptr, encodedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
}/* debug [functions.gen.go/function]: CMSEncodeContent */

// Specifies a message is to be encrypted and specifies the recipients of the message.
//
// Added in macOS 10.5.
// Specifies a message is to be encrypted and specifies the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddRecipients(_:_:)
func CMSEncoderAddRecipients(cmsEncoder SEncoderRef, recipientOrArray TypeRef) unsafe.Pointer {
	return _CMSEncoderAddRecipients(cmsEncoder, recipientOrArray)
}/* debug [functions.gen.go/function]: CMSEncoderAddRecipients */

// Specifies attributes for a signed message.
//
// Added in macOS 10.5.
// Specifies attributes for a signed message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSignedAttributes(_:_:)
func CMSEncoderAddSignedAttributes(cmsEncoder SEncoderRef, signedAttributes SSignedAttributes) unsafe.Pointer {
	return _CMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes)
}/* debug [functions.gen.go/function]: CMSEncoderAddSignedAttributes */

// Specifies signers of the message.
//
// Added in macOS 10.5.
// Specifies signers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSigners(_:_:)
func CMSEncoderAddSigners(cmsEncoder SEncoderRef, signerOrArray TypeRef) unsafe.Pointer {
	return _CMSEncoderAddSigners(cmsEncoder, signerOrArray)
}/* debug [functions.gen.go/function]: CMSEncoderAddSigners */

// Adds certificates to a message.
//
// Added in macOS 10.5.
// Adds certificates to a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSupportingCerts(_:_:)
func CMSEncoderAddSupportingCerts(cmsEncoder SEncoderRef, certOrArray TypeRef) unsafe.Pointer {
	return _CMSEncoderAddSupportingCerts(cmsEncoder, certOrArray)
}/* debug [functions.gen.go/function]: CMSEncoderAddSupportingCerts */

// Obtains the object identifier for the encapsulated data of a signed message.
//
// Added in macOS 10.5.
// Obtains the object identifier for the encapsulated data of a signed message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncapsulatedContentType(_:_:)
func CMSEncoderCopyEncapsulatedContentType(cmsEncoder SEncoderRef, eContentTypeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut)
}/* debug [functions.gen.go/function]: CMSEncoderCopyEncapsulatedContentType */

// Finishes encoding the message and obtains the encoded result.
//
// Added in macOS 10.5.
// Finishes encoding the message and obtains the encoded result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncodedContent(_:_:)
func CMSEncoderCopyEncodedContent(cmsEncoder SEncoderRef, encodedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut)
}/* debug [functions.gen.go/function]: CMSEncoderCopyEncodedContent */

// Obtains the array of recipients specified with the function.
//
// Added in macOS 10.5.
// Obtains the array of recipients specified with the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyRecipients(_:_:)
func CMSEncoderCopyRecipients(cmsEncoder SEncoderRef, recipientsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyRecipients(cmsEncoder, recipientsOut)
}/* debug [functions.gen.go/function]: CMSEncoderCopyRecipients */

// Obtains the array of signers specified with the function.
//
// Added in macOS 10.5.
// Obtains the array of signers specified with the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySigners(_:_:)
func CMSEncoderCopySigners(cmsEncoder SEncoderRef, signersOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySigners(cmsEncoder, signersOut)
}/* debug [functions.gen.go/function]: CMSEncoderCopySigners */

// Returns the timestamp of a signer of a CMS message, if present.
//
// Added in macOS 10.8.
// Returns the timestamp of a signer of a CMS message, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestamp(_:_:_:)
func CMSEncoderCopySignerTimestamp(cmsEncoder SEncoderRef, signerIndex uintptr, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp)
}/* debug [functions.gen.go/function]: CMSEncoderCopySignerTimestamp */

// Returns the timestamp of a signer of a CMS message using a particular policy, if present.
//
// Added in macOS 10.10.
// Returns the timestamp of a signer of a CMS message using a particular policy, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder SEncoderRef, timeStampPolicy TypeRef, signerIndex uintptr, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp)
}/* debug [functions.gen.go/function]: CMSEncoderCopySignerTimestampWithPolicy */

// Obtains the certificates added to a message with .
//
// Added in macOS 10.5.
// Obtains the certificates added to a message with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySupportingCerts(_:_:)
func CMSEncoderCopySupportingCerts(cmsEncoder SEncoderRef, certsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySupportingCerts(cmsEncoder, certsOut)
}/* debug [functions.gen.go/function]: CMSEncoderCopySupportingCerts */

// Creates a CMSEncoder reference.
//
// Added in macOS 10.5.
// Creates a CMSEncoder reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCreate(_:)
func CMSEncoderCreate(cmsEncoderOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCreate(cmsEncoderOut)
}/* debug [functions.gen.go/function]: CMSEncoderCreate */

// Obtains a constant that indicates which certificates are to be included in a signed CMS message.
//
// Added in macOS 10.5.
// Obtains a constant that indicates which certificates are to be included in a signed CMS message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetCertificateChainMode(_:_:)
func CMSEncoderGetCertificateChainMode(cmsEncoder SEncoderRef, chainModeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut)
}/* debug [functions.gen.go/function]: CMSEncoderGetCertificateChainMode */

// Indicates whether the message is to have detached content.
//
// Added in macOS 10.5.
// Indicates whether the message is to have detached content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetHasDetachedContent(_:_:)
func CMSEncoderGetHasDetachedContent(cmsEncoder SEncoderRef, detachedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut)
}/* debug [functions.gen.go/function]: CMSEncoderGetHasDetachedContent */

// Returns the type identifier for the CMSEncoder opaque type.
//
// Added in macOS 10.5.
// Returns the type identifier for the CMSEncoder opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetTypeID()
func CMSEncoderGetTypeID() TypeID {
	return _CMSEncoderGetTypeID()
}/* debug [functions.gen.go/function]: CMSEncoderGetTypeID */

// Specifies which certificates to include in a signed CMS message.
//
// Added in macOS 10.5.
// Specifies which certificates to include in a signed CMS message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetCertificateChainMode(_:_:)
func CMSEncoderSetCertificateChainMode(cmsEncoder SEncoderRef, chainMode SCertificateChainMode) unsafe.Pointer {
	return _CMSEncoderSetCertificateChainMode(cmsEncoder, chainMode)
}/* debug [functions.gen.go/function]: CMSEncoderSetCertificateChainMode */

// Specifies an object identifier for the encapsulated data of a signed message.
//
// Added in macOS 10.7.
// Specifies an object identifier for the encapsulated data of a signed message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentTypeOID(_:_:)
func CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder SEncoderRef, eContentTypeOID TypeRef) unsafe.Pointer {
	return _CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID)
}/* debug [functions.gen.go/function]: CMSEncoderSetEncapsulatedContentTypeOID */

// Specifies whether the signed data is to be separate from the message.
//
// Added in macOS 10.5.
// Specifies whether the signed data is to be separate from the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetHasDetachedContent(_:_:)
func CMSEncoderSetHasDetachedContent(cmsEncoder SEncoderRef, detachedContent unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent)
}/* debug [functions.gen.go/function]: CMSEncoderSetHasDetachedContent */

// Sets the digest algorithm to use for the signer.
//
// Added in macOS 10.11.
// Sets the digest algorithm to use for the signer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetSignerAlgorithm(_:_:)
func CMSEncoderSetSignerAlgorithm(cmsEncoder SEncoderRef, digestAlgorithm StringRef) unsafe.Pointer {
	return _CMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm)
}/* debug [functions.gen.go/function]: CMSEncoderSetSignerAlgorithm */

// Feeds content bytes into the encoder.
//
// Added in macOS 10.5.
// Feeds content bytes into the encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderUpdateContent(_:_:_:)
func CMSEncoderUpdateContent(cmsEncoder SEncoderRef, content unsafe.Pointer, contentLen uintptr) unsafe.Pointer {
	return _CMSEncoderUpdateContent(cmsEncoder, content, contentLen)
}/* debug [functions.gen.go/function]: CMSEncoderUpdateContent */

// cssmAlgToOid is a Security function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmAlgToOid(_:)
func cssmAlgToOid(algId CSSM_ALGORITHMS) unsafe.Pointer {
	return _cssmAlgToOid(algId)
}/* debug [functions.gen.go/function]: cssmAlgToOid */

// cssmOidToAlg is a Security function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmOidToAlg(_:_:)
func cssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) bool {
	return _cssmOidToAlg(oid, alg)
}/* debug [functions.gen.go/function]: cssmOidToAlg */

// cssmPerror is a Security function.
//
// Added in macOS 10.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmPerror(_:_:)
func cssmPerror(how unsafe.Pointer, error_ CSSM_RETURN) {
	_cssmPerror(how, error_)
}/* debug [functions.gen.go/function]: cssmPerror */

// sec_protocol_metadata_copy_negotiated_protocol is a Security function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_negotiated_protocol(_:)
func sec_protocol_metadata_copy_negotiated_protocol(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_copy_negotiated_protocol(metadata)
}/* debug [functions.gen.go/function]: sec_protocol_metadata_copy_negotiated_protocol */

// sec_protocol_metadata_copy_server_name is a Security function.
//
// Added in macOS 15.5.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_server_name(_:)
func sec_protocol_metadata_copy_server_name(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_copy_server_name(metadata)
}/* debug [functions.gen.go/function]: sec_protocol_metadata_copy_server_name */

// Creates a new access control object with the specified protection type and flags.
//
// Added in macOS 10.10.
// Creates a new access control object with the specified protection type and flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateWithFlags(_:_:_:_:)
func SecAccessControlCreateWithFlags(allocator AllocatorRef, protection TypeRef, flags SecAccessControlCreateFlags, error_ unsafe.Pointer) SecAccessControlRef {
	return _SecAccessControlCreateWithFlags(allocator, protection, flags, error_)
}/* debug [functions.gen.go/function]: SecAccessControlCreateWithFlags */

// Returns the unique identifier of the opaque type to which a keychain item access control object belongs.
//
// Added in macOS 10.10.
// Returns the unique identifier of the opaque type to which a keychain item access control object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlGetTypeID()
func SecAccessControlGetTypeID() TypeID {
	return _SecAccessControlGetTypeID()
}/* debug [functions.gen.go/function]: SecAccessControlGetTypeID */

// Retrieves all the ACL entries of a given access instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves all the ACL entries of a given access instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopyACLList(_:_:)
func SecAccessCopyACLList(accessRef SecAccessRef, aclList unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopyACLList(accessRef, aclList)
}/* debug [functions.gen.go/function]: SecAccessCopyACLList */

// Retrieves selected ACL entries from a given access instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Retrieves selected ACL entries from a given access instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopyMatchingACLList(_:_:)
func SecAccessCopyMatchingACLList(accessRef SecAccessRef, authorizationTag TypeRef) ArrayRef {
	return _SecAccessCopyMatchingACLList(accessRef, authorizationTag)
}/* debug [functions.gen.go/function]: SecAccessCopyMatchingACLList */

// Retrieves the owner and the ACL entries of a given access instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Retrieves the owner and the ACL entries of a given access instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopyOwnerAndACL(_:_:_:_:_:)
func SecAccessCopyOwnerAndACL(accessRef SecAccessRef, userId unsafe.Pointer, groupId unsafe.Pointer, ownerType unsafe.Pointer, aclList unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopyOwnerAndACL(accessRef, userId, groupId, ownerType, aclList)
}/* debug [functions.gen.go/function]: SecAccessCopyOwnerAndACL */

// Creates a new access instance associated with a given protected keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Creates a new access instance associated with a given protected keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCreate(_:_:_:)
func SecAccessCreate(descriptor StringRef, trustedlist ArrayRef, accessRef unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCreate(descriptor, trustedlist, accessRef)
}/* debug [functions.gen.go/function]: SecAccessCreate */

// Creates a new access instance using the owner and ACL entries you provide.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Creates a new access instance using the owner and ACL entries you provide.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCreateWithOwnerAndACL(_:_:_:_:_:)
func SecAccessCreateWithOwnerAndACL(userId unsafe.Pointer, groupId unsafe.Pointer, ownerType SecAccessOwnerType, acls ArrayRef, error_ unsafe.Pointer) SecAccessRef {
	return _SecAccessCreateWithOwnerAndACL(userId, groupId, ownerType, acls, error_)
}/* debug [functions.gen.go/function]: SecAccessCreateWithOwnerAndACL */

// Returns the unique identifier of the opaque type to which an access instance belongs.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Returns the unique identifier of the opaque type to which an access instance belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessGetTypeID()
func SecAccessGetTypeID() TypeID {
	return _SecAccessGetTypeID()
}/* debug [functions.gen.go/function]: SecAccessGetTypeID */

// Retrieves the authorization tags of a given ACL entry.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Retrieves the authorization tags of a given ACL entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCopyAuthorizations(_:)
func SecACLCopyAuthorizations(acl SecACLRef) ArrayRef {
	return _SecACLCopyAuthorizations(acl)
}/* debug [functions.gen.go/function]: SecACLCopyAuthorizations */

// Returns the application list, description, and prompt selector for a given ACL entry.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Returns the application list, description, and prompt selector for a given ACL entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCopyContents(_:_:_:_:)
func SecACLCopyContents(acl SecACLRef, applicationList unsafe.Pointer, description unsafe.Pointer, promptSelector unsafe.Pointer) unsafe.Pointer {
	return _SecACLCopyContents(acl, applicationList, description, promptSelector)
}/* debug [functions.gen.go/function]: SecACLCopyContents */

// Creates a new ACL entry with the given characteristics, and adds it to an access instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Creates a new ACL entry with the given characteristics, and adds it to an access instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCreateWithSimpleContents(_:_:_:_:_:)
func SecACLCreateWithSimpleContents(access SecAccessRef, applicationList ArrayRef, description StringRef, promptSelector SecKeychainPromptSelector, newAcl unsafe.Pointer) unsafe.Pointer {
	return _SecACLCreateWithSimpleContents(access, applicationList, description, promptSelector, newAcl)
}/* debug [functions.gen.go/function]: SecACLCreateWithSimpleContents */

// Returns the unique identifier of the opaque type to which an ACL entry belongs.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which an ACL entry belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLGetTypeID()
func SecACLGetTypeID() TypeID {
	return _SecACLGetTypeID()
}/* debug [functions.gen.go/function]: SecACLGetTypeID */

// Removes the specified ACL entry from the access instance that contains it.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.3.
// Removes the specified ACL entry from the access instance that contains it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLRemove(_:)
func SecACLRemove(aclRef SecACLRef) unsafe.Pointer {
	return _SecACLRemove(aclRef)
}/* debug [functions.gen.go/function]: SecACLRemove */

// Sets the application list, description, and prompt selector for a given ACL entry.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Sets the application list, description, and prompt selector for a given ACL entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLSetContents(_:_:_:_:)
func SecACLSetContents(acl SecACLRef, applicationList ArrayRef, description StringRef, promptSelector SecKeychainPromptSelector) unsafe.Pointer {
	return _SecACLSetContents(acl, applicationList, description, promptSelector)
}/* debug [functions.gen.go/function]: SecACLSetContents */

// Sets the authorization tags for a given ACL.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
// Sets the authorization tags for a given ACL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLUpdateAuthorizations(_:_:)
func SecACLUpdateAuthorizations(acl SecACLRef, authorizations ArrayRef) unsafe.Pointer {
	return _SecACLUpdateAuthorizations(acl, authorizations)
}/* debug [functions.gen.go/function]: SecACLUpdateAuthorizations */

// Asynchronously stores (or updates) a shared password for a website.
//
// Added in macOS 11.0.
// Asynchronously stores (or updates) a shared password for a website.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAddSharedWebCredential(_:_:_:_:)
func SecAddSharedWebCredential(fqdn StringRef, account StringRef, password StringRef) {
	_SecAddSharedWebCredential(fqdn, account, password)
}/* debug [functions.gen.go/function]: SecAddSharedWebCredential */

// Adds a certificate to a keychain.
//
// Added in macOS 10.3.
// Adds a certificate to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateAddToKeychain(_:_:)
func SecCertificateAddToKeychain(certificate SecCertificateRef, keychain SecKeychainRef) unsafe.Pointer {
	return _SecCertificateAddToKeychain(certificate, keychain)
}/* debug [functions.gen.go/function]: SecCertificateAddToKeychain */

// Retrieves the common name of the subject of a certificate.
//
// Added in macOS 10.5.
// Retrieves the common name of the subject of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyCommonName(_:_:)
func SecCertificateCopyCommonName(certificate SecCertificateRef, commonName unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyCommonName(certificate, commonName)
}/* debug [functions.gen.go/function]: SecCertificateCopyCommonName */

// Returns a DER representation of a certificate given a certificate object.
//
// Added in macOS 10.6.
// Returns a DER representation of a certificate given a certificate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyData(_:)
func SecCertificateCopyData(certificate SecCertificateRef) DataRef {
	return _SecCertificateCopyData(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyData */

// Retrieves the email addresses for the subject of a certificate.
//
// Added in macOS 10.5.
// Retrieves the email addresses for the subject of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyEmailAddresses(_:_:)
func SecCertificateCopyEmailAddresses(certificate SecCertificateRef, emailAddresses unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyEmailAddresses(certificate, emailAddresses)
}/* debug [functions.gen.go/function]: SecCertificateCopyEmailAddresses */

// Retrieves the public key for a given certificate.
//
// Added in macOS 10.14.
// Retrieves the public key for a given certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyKey(_:)
func SecCertificateCopyKey(certificate SecCertificateRef) SecKeyRef {
	return _SecCertificateCopyKey(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyKey */

// Returns a copy of the long description of a certificate.
//
// Added in macOS 10.7.
// Returns a copy of the long description of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyLongDescription(_:_:_:)
func SecCertificateCopyLongDescription(alloc AllocatorRef, certificate SecCertificateRef, error_ unsafe.Pointer) StringRef {
	return _SecCertificateCopyLongDescription(alloc, certificate, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopyLongDescription */

// Returns a normalized copy of the distinguished name (DN) of the issuer of a certificate.
//
// Deprecated: This function was deprecated in macOS 10.12.4.
//
// Added in macOS 10.7.
// Returns a normalized copy of the distinguished name (DN) of the issuer of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedIssuerContent(_:_:)
func SecCertificateCopyNormalizedIssuerContent(certificate SecCertificateRef, error_ unsafe.Pointer) DataRef {
	return _SecCertificateCopyNormalizedIssuerContent(certificate, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopyNormalizedIssuerContent */

// Retrieves the normalized issuer sequence from a certificate.
//
// Added in macOS 10.12.4.
// Retrieves the normalized issuer sequence from a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedIssuerSequence(_:)
func SecCertificateCopyNormalizedIssuerSequence(certificate SecCertificateRef) DataRef {
	return _SecCertificateCopyNormalizedIssuerSequence(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyNormalizedIssuerSequence */

// Returns a normalized copy of the distinguished name (DN) of the subject of a certificate.
//
// Deprecated: This function was deprecated in macOS 10.12.4.
//
// Added in macOS 10.7.
// Returns a normalized copy of the distinguished name (DN) of the subject of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedSubjectContent(_:_:)
func SecCertificateCopyNormalizedSubjectContent(certificate SecCertificateRef, error_ unsafe.Pointer) DataRef {
	return _SecCertificateCopyNormalizedSubjectContent(certificate, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopyNormalizedSubjectContent */

// Retrieves the normalized subject sequence from a certificate.
//
// Added in macOS 10.12.4.
// Retrieves the normalized subject sequence from a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNormalizedSubjectSequence(_:)
func SecCertificateCopyNormalizedSubjectSequence(certificate SecCertificateRef) DataRef {
	return _SecCertificateCopyNormalizedSubjectSequence(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyNormalizedSubjectSequence */

// Returns the preferred certificate for the specified name and key usage.
//
// Added in macOS 10.7.
// Returns the preferred certificate for the specified name and key usage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyPreferred(_:_:)
func SecCertificateCopyPreferred(name StringRef, keyUsage ArrayRef) SecCertificateRef {
	return _SecCertificateCopyPreferred(name, keyUsage)
}/* debug [functions.gen.go/function]: SecCertificateCopyPreferred */

// Retrieves the public key from a certificate.
//
// Deprecated: This function was deprecated in macOS 10.14.
//
// Added in macOS 10.3.
// Retrieves the public key from a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyPublicKey(_:_:)
func SecCertificateCopyPublicKey(certificate SecCertificateRef) SecKeyRef {
	return _SecCertificateCopyPublicKey(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopyPublicKey */

// Returns a copy of a certificate’s serial number.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.7.
// Returns a copy of a certificate’s serial number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopySerialNumber(_:_:)
func SecCertificateCopySerialNumber(certificate SecCertificateRef) DataRef {
	return _SecCertificateCopySerialNumber(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopySerialNumber */

// Returns the certificate’s serial number.
//
// Added in macOS 10.13.
// Returns the certificate’s serial number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopySerialNumberData(_:_:)
func SecCertificateCopySerialNumberData(certificate SecCertificateRef, error_ unsafe.Pointer) DataRef {
	return _SecCertificateCopySerialNumberData(certificate, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopySerialNumberData */

// Returns a copy of the short description of a certificate.
//
// Added in macOS 10.7.
// Returns a copy of the short description of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyShortDescription(_:_:_:)
func SecCertificateCopyShortDescription(alloc AllocatorRef, certificate SecCertificateRef, error_ unsafe.Pointer) StringRef {
	return _SecCertificateCopyShortDescription(alloc, certificate, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopyShortDescription */

// Returns a human-readable summary of a certificate.
//
// Added in macOS 10.6.
// Returns a human-readable summary of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopySubjectSummary(_:)
func SecCertificateCopySubjectSummary(certificate SecCertificateRef) StringRef {
	return _SecCertificateCopySubjectSummary(certificate)
}/* debug [functions.gen.go/function]: SecCertificateCopySubjectSummary */

// Creates a dictionary that represents a certificate’s contents.
//
// Added in macOS 10.7.
// Creates a dictionary that represents a certificate’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyValues(_:_:_:)
func SecCertificateCopyValues(certificate SecCertificateRef, keys ArrayRef, error_ unsafe.Pointer) DictionaryRef {
	return _SecCertificateCopyValues(certificate, keys, error_)
}/* debug [functions.gen.go/function]: SecCertificateCopyValues */

// Creates a certificate object from a DER representation of a certificate.
//
// Added in macOS 10.6.
// Creates a certificate object from a DER representation of a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
func SecCertificateCreateWithData(allocator AllocatorRef, data DataRef) SecCertificateRef {
	return _SecCertificateCreateWithData(allocator, data)
}/* debug [functions.gen.go/function]: SecCertificateCreateWithData */

// Returns the unique identifier of the opaque type to which a certificate object belongs.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which a certificate object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateGetTypeID()
func SecCertificateGetTypeID() TypeID {
	return _SecCertificateGetTypeID()
}/* debug [functions.gen.go/function]: SecCertificateGetTypeID */

// Sets the certificate that should be preferred for the specified name and key use.
//
// Added in macOS 10.7.
// Sets the certificate that should be preferred for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateSetPreferred(_:_:_:)
func SecCertificateSetPreferred(certificate SecCertificateRef, name StringRef, keyUsage ArrayRef) unsafe.Pointer {
	return _SecCertificateSetPreferred(certificate, name, keyUsage)
}/* debug [functions.gen.go/function]: SecCertificateSetPreferred */

// Performs dynamic validation of signed code.
//
// Added in macOS 10.0.
// Performs dynamic validation of signed code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCheckValidity(_:_:_:)
func SecCodeCheckValidity(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef) unsafe.Pointer {
	return _SecCodeCheckValidity(code, flags, requirement)
}/* debug [functions.gen.go/function]: SecCodeCheckValidity */

// Performs dynamic validation of signed code and returns detailed error information in the case of failure.
//
// Added in macOS 10.0.
// Performs dynamic validation of signed code and returns detailed error information in the case of failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCheckValidityWithErrors(_:_:_:_:)
func SecCodeCheckValidityWithErrors(code SecCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCheckValidityWithErrors(code, flags, requirement, errors)
}/* debug [functions.gen.go/function]: SecCodeCheckValidityWithErrors */

// Retrieves the designated code requirement of signed code.
//
// Added in macOS 10.0.
// Retrieves the designated code requirement of signed code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)
func SecCodeCopyDesignatedRequirement(code SecStaticCodeRef, flags SecCSFlags, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyDesignatedRequirement(code, flags, requirement)
}/* debug [functions.gen.go/function]: SecCodeCopyDesignatedRequirement */

// Asks a code host to identify one of its guests given the type and value of specific attributes of the guest code.
//
// Added in macOS 10.0.
// Asks a code host to identify one of its guests given the type and value of specific attributes of the guest code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyGuestWithAttributes(_:_:_:_:)
func SecCodeCopyGuestWithAttributes(host SecCodeRef, attributes DictionaryRef, flags SecCSFlags, guest unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyGuestWithAttributes(host, attributes, flags, guest)
}/* debug [functions.gen.go/function]: SecCodeCopyGuestWithAttributes */

// Retrieves the code object for the host of specified guest code.
//
// Added in macOS 10.0.
// Retrieves the code object for the host of specified guest code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyHost(_:_:_:)
func SecCodeCopyHost(guest SecCodeRef, flags SecCSFlags, host unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyHost(guest, flags, host)
}/* debug [functions.gen.go/function]: SecCodeCopyHost */

// Retrieves the location on disk of signed code, given a code or static code object.
//
// Added in macOS 10.0.
// Retrieves the location on disk of signed code, given a code or static code object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyPath(_:_:_:)
func SecCodeCopyPath(staticCode SecStaticCodeRef, flags SecCSFlags, path unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyPath(staticCode, flags, path)
}/* debug [functions.gen.go/function]: SecCodeCopyPath */

// Retrieves the code object for the code making the call.
//
// Added in macOS 10.0.
// Retrieves the code object for the code making the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopySelf(_:_:)
func SecCodeCopySelf(flags SecCSFlags, self unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopySelf(flags, self)
}/* debug [functions.gen.go/function]: SecCodeCopySelf */

// Retrieves various pieces of information from a code signature.
//
// Added in macOS 10.0.
// Retrieves various pieces of information from a code signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopySigningInformation(_:_:_:)
func SecCodeCopySigningInformation(code SecStaticCodeRef, flags SecCSFlags, information unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopySigningInformation(code, flags, information)
}/* debug [functions.gen.go/function]: SecCodeCopySigningInformation */

// Returns a static code object representing the on-disk version of the given running code.
//
// Added in macOS 10.0.
// Returns a static code object representing the on-disk version of the given running code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyStaticCode(_:_:_:)
func SecCodeCopyStaticCode(code SecCodeRef, flags SecCSFlags, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyStaticCode(code, flags, staticCode)
}/* debug [functions.gen.go/function]: SecCodeCopyStaticCode */

// Returns the unique identifier of the opaque type to which a code object belongs.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a code object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeGetTypeID()
func SecCodeGetTypeID() TypeID {
	return _SecCodeGetTypeID()
}/* debug [functions.gen.go/function]: SecCodeGetTypeID */

// Asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins.
//
// Added in macOS 10.0.
// Asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeMapMemory(_:_:)
func SecCodeMapMemory(code SecStaticCodeRef, flags SecCSFlags) unsafe.Pointer {
	return _SecCodeMapMemory(code, flags)
}/* debug [functions.gen.go/function]: SecCodeMapMemory */

// Returns a string explaining the meaning of a security result code.
//
// Added in macOS 10.3.
// Returns a string explaining the meaning of a security result code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCopyErrorMessageString(_:_:)
func SecCopyErrorMessageString(status unsafe.Pointer, reserved unsafe.Pointer) StringRef {
	return _SecCopyErrorMessageString(status, reserved)
}/* debug [functions.gen.go/function]: SecCopyErrorMessageString */

// Returns a randomly generated password.
//
// Added in macOS 11.0.
// Returns a randomly generated password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCreateSharedWebCredentialPassword()
func SecCreateSharedWebCredentialPassword() StringRef {
	return _SecCreateSharedWebCredentialPassword()
}/* debug [functions.gen.go/function]: SecCreateSharedWebCredentialPassword */

// Creates a decode transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a decode transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecodeTransformCreate(_:_:)
func SecDecodeTransformCreate(DecodeType TypeRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecDecodeTransformCreate(DecodeType, error_)
}/* debug [functions.gen.go/function]: SecDecodeTransformCreate */

// Creates a decryption transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a decryption transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecryptTransformCreate(_:_:)
func SecDecryptTransformCreate(keyRef SecKeyRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecDecryptTransformCreate(keyRef, error_)
}/* debug [functions.gen.go/function]: SecDecryptTransformCreate */

// Returns the unique identifier of the opaque type to which a decryption transform belongs.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Returns the unique identifier of the opaque type to which a decryption transform belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecryptTransformGetTypeID()
func SecDecryptTransformGetTypeID() TypeID {
	return _SecDecryptTransformGetTypeID()
}/* debug [functions.gen.go/function]: SecDecryptTransformGetTypeID */

// Creates a digest transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a digest transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDigestTransformCreate(_:_:_:)
func SecDigestTransformCreate(digestType TypeRef, digestLength Index, error_ unsafe.Pointer) SecTransformRef {
	return _SecDigestTransformCreate(digestType, digestLength, error_)
}/* debug [functions.gen.go/function]: SecDigestTransformCreate */

// Returns the unique identifier of the opaque type to which a digest transform belongs.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Returns the unique identifier of the opaque type to which a digest transform belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDigestTransformGetTypeID()
func SecDigestTransformGetTypeID() TypeID {
	return _SecDigestTransformGetTypeID()
}/* debug [functions.gen.go/function]: SecDigestTransformGetTypeID */

// Creates an encode transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates an encode transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncodeTransformCreate(_:_:)
func SecEncodeTransformCreate(encodeType TypeRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecEncodeTransformCreate(encodeType, error_)
}/* debug [functions.gen.go/function]: SecEncodeTransformCreate */

// Creates an encryption transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates an encryption transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncryptTransformCreate(_:_:)
func SecEncryptTransformCreate(keyRef SecKeyRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecEncryptTransformCreate(keyRef, error_)
}/* debug [functions.gen.go/function]: SecEncryptTransformCreate */

// Returns the unique identifier of the opaque type to which an encryption transform belongs.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Returns the unique identifier of the opaque type to which an encryption transform belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncryptTransformGetTypeID()
func SecEncryptTransformGetTypeID() TypeID {
	return _SecEncryptTransformGetTypeID()
}/* debug [functions.gen.go/function]: SecEncryptTransformGetTypeID */

// Returns the Core Foundation type ID for a transform group container.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Returns the Core Foundation type ID for a transform group container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecGroupTransformGetTypeID()
func SecGroupTransformGetTypeID() TypeID {
	return _SecGroupTransformGetTypeID()
}/* debug [functions.gen.go/function]: SecGroupTransformGetTypeID */

// Retrieves a certificate associated with an identity.
//
// Added in macOS 10.3.
// Retrieves a certificate associated with an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCopyCertificate(_:_:)
func SecIdentityCopyCertificate(identityRef SecIdentityRef, certificateRef unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCopyCertificate(identityRef, certificateRef)
}/* debug [functions.gen.go/function]: SecIdentityCopyCertificate */

// Retrieves the preferred identity for the specified name and key use.
//
// Added in macOS 10.7.
// Retrieves the preferred identity for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCopyPreferred(_:_:_:)
func SecIdentityCopyPreferred(name StringRef, keyUsage ArrayRef, validIssuers ArrayRef) SecIdentityRef {
	return _SecIdentityCopyPreferred(name, keyUsage, validIssuers)
}/* debug [functions.gen.go/function]: SecIdentityCopyPreferred */

// Retrieves the private key associated with an identity.
//
// Added in macOS 10.3.
// Retrieves the private key associated with an identity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCopyPrivateKey(_:_:)
func SecIdentityCopyPrivateKey(identityRef SecIdentityRef, privateKeyRef unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCopyPrivateKey(identityRef, privateKeyRef)
}/* debug [functions.gen.go/function]: SecIdentityCopyPrivateKey */

// Obtains the system identity associated with a specified domain.
//
// Added in macOS 10.5.
// Obtains the system identity associated with a specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCopySystemIdentity(_:_:_:)
func SecIdentityCopySystemIdentity(domain StringRef, idRef unsafe.Pointer, actualDomain unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCopySystemIdentity(domain, idRef, actualDomain)
}/* debug [functions.gen.go/function]: SecIdentityCopySystemIdentity */

// Creates a new identity for a certificate and its associated private key.
//
// Added in macOS 10.5.
// Creates a new identity for a certificate and its associated private key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCreateWithCertificate(_:_:_:)
func SecIdentityCreateWithCertificate(keychainOrArray TypeRef, certificateRef SecCertificateRef, identityRef unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCreateWithCertificate(keychainOrArray, certificateRef, identityRef)
}/* debug [functions.gen.go/function]: SecIdentityCreateWithCertificate */

// Returns the unique identifier of the opaque type to which an identity object belongs.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which an identity object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityGetTypeID()
func SecIdentityGetTypeID() TypeID {
	return _SecIdentityGetTypeID()
}/* debug [functions.gen.go/function]: SecIdentityGetTypeID */

// Sets the identity that should be preferred for the specified name and key use.
//
// Added in macOS 10.7.
// Sets the identity that should be preferred for the specified name and key use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySetPreferred(_:_:_:)
func SecIdentitySetPreferred(identity SecIdentityRef, name StringRef, keyUsage ArrayRef) unsafe.Pointer {
	return _SecIdentitySetPreferred(identity, name, keyUsage)
}/* debug [functions.gen.go/function]: SecIdentitySetPreferred */

// Assigns the system identity to be associated with a specified domain.
//
// Added in macOS 10.5.
// Assigns the system identity to be associated with a specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySetSystemIdentity(_:_:)
func SecIdentitySetSystemIdentity(domain StringRef, idRef SecIdentityRef) unsafe.Pointer {
	return _SecIdentitySetSystemIdentity(domain, idRef)
}/* debug [functions.gen.go/function]: SecIdentitySetSystemIdentity */

// Adds one or more items to a keychain.
//
// Added in macOS 10.6.
// Adds one or more items to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAdd(_:_:)
func SecItemAdd(attributes DictionaryRef, result unsafe.Pointer) unsafe.Pointer {
	return _SecItemAdd(attributes, result)
}/* debug [functions.gen.go/function]: SecItemAdd */

// Returns one or more keychain items that match a search query, or copies attributes of specific keychain items.
//
// Added in macOS 10.6.
// Returns one or more keychain items that match a search query, or copies attributes of specific keychain items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemCopyMatching(_:_:)
func SecItemCopyMatching(query DictionaryRef, result unsafe.Pointer) unsafe.Pointer {
	return _SecItemCopyMatching(query, result)
}/* debug [functions.gen.go/function]: SecItemCopyMatching */

// Deletes items that match a search query.
//
// Added in macOS 10.6.
// Deletes items that match a search query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemDelete(_:)
func SecItemDelete(query DictionaryRef) unsafe.Pointer {
	return _SecItemDelete(query)
}/* debug [functions.gen.go/function]: SecItemDelete */

// Exports one or more certificates, keys, or identities.
//
// Added in macOS 10.7.
// Exports one or more certificates, keys, or identities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemExport(_:_:_:_:_:)
func SecItemExport(secItemOrArray TypeRef, outputFormat SecExternalFormat, flags SecItemImportExportFlags, keyParams unsafe.Pointer, exportedData unsafe.Pointer) unsafe.Pointer {
	return _SecItemExport(secItemOrArray, outputFormat, flags, keyParams, exportedData)
}/* debug [functions.gen.go/function]: SecItemExport */

// Imports one or more certificates, keys, or identities and optionally adds them to a keychain.
//
// Added in macOS 10.7.
// Imports one or more certificates, keys, or identities and optionally adds them to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImport(_:_:_:_:_:_:_:_:)
func SecItemImport(importedData DataRef, fileNameOrExtension StringRef, inputFormat unsafe.Pointer, itemType unsafe.Pointer, flags SecItemImportExportFlags, keyParams unsafe.Pointer, importKeychain SecKeychainRef, outItems unsafe.Pointer) unsafe.Pointer {
	return _SecItemImport(importedData, fileNameOrExtension, inputFormat, itemType, flags, keyParams, importKeychain, outItems)
}/* debug [functions.gen.go/function]: SecItemImport */

// Modifies items that match a search query.
//
// Added in macOS 10.6.
// Modifies items that match a search query.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemUpdate(_:_:)
func SecItemUpdate(query DictionaryRef, attributesToUpdate DictionaryRef) unsafe.Pointer {
	return _SecItemUpdate(query, attributesToUpdate)
}/* debug [functions.gen.go/function]: SecItemUpdate */

// Registers your keychain event callback function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Registers your keychain event callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAddCallback(_:_:_:)
func SecKeychainAddCallback(callbackFunction SecKeychainCallback, eventMask SecKeychainEventMask, userContext unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainAddCallback(callbackFunction, eventMask, userContext)
}/* debug [functions.gen.go/function]: SecKeychainAddCallback */

// Adds a new generic password to a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Adds a new generic password to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAddGenericPassword(_:_:_:_:_:_:_:_:)
func SecKeychainAddGenericPassword(keychain SecKeychainRef, serviceNameLength unsafe.Pointer, serviceName unsafe.Pointer, accountNameLength unsafe.Pointer, accountName unsafe.Pointer, passwordLength unsafe.Pointer, passwordData unsafe.Pointer, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainAddGenericPassword(keychain, serviceNameLength, serviceName, accountNameLength, accountName, passwordLength, passwordData, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainAddGenericPassword */

// Adds a new Internet password to a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Adds a new Internet password to a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAddInternetPassword(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func SecKeychainAddInternetPassword(keychain SecKeychainRef, serverNameLength unsafe.Pointer, serverName unsafe.Pointer, securityDomainLength unsafe.Pointer, securityDomain unsafe.Pointer, accountNameLength unsafe.Pointer, accountName unsafe.Pointer, pathLength unsafe.Pointer, path unsafe.Pointer, port unsafe.Pointer, protocol_ SecProtocolType, authenticationType SecAuthenticationType, passwordLength unsafe.Pointer, passwordData unsafe.Pointer, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainAddInternetPassword(keychain, serverNameLength, serverName, securityDomainLength, securityDomain, accountNameLength, accountName, pathLength, path, port, protocol_, authenticationType, passwordLength, passwordData, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainAddInternetPassword */

// Obtains tags for all possible attributes of a given item class.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Obtains tags for all possible attributes of a given item class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeInfoForItemID(_:_:_:)
func SecKeychainAttributeInfoForItemID(keychain SecKeychainRef, itemID unsafe.Pointer, info unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainAttributeInfoForItemID(keychain, itemID, info)
}/* debug [functions.gen.go/function]: SecKeychainAttributeInfoForItemID */

// Retrieves the application access of a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves the application access of a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopyAccess(_:_:)
func SecKeychainCopyAccess(keychain SecKeychainRef, access unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopyAccess(keychain, access)
}/* debug [functions.gen.go/function]: SecKeychainCopyAccess */

// Retrieves a pointer to the default keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves a pointer to the default keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopyDefault(_:)
func SecKeychainCopyDefault(keychain unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopyDefault(keychain)
}/* debug [functions.gen.go/function]: SecKeychainCopyDefault */

// Retrieves the default keychain from a specified preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves the default keychain from a specified preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopyDomainDefault(_:_:)
func SecKeychainCopyDomainDefault(domain SecPreferencesDomain, keychain unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopyDomainDefault(domain, keychain)
}/* debug [functions.gen.go/function]: SecKeychainCopyDomainDefault */

// Retrieves the keychain search list for a specified preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves the keychain search list for a specified preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopyDomainSearchList(_:_:)
func SecKeychainCopyDomainSearchList(domain SecPreferencesDomain, searchList unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopyDomainSearchList(domain, searchList)
}/* debug [functions.gen.go/function]: SecKeychainCopyDomainSearchList */

// Retrieves a keychain search list.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves a keychain search list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopySearchList(_:)
func SecKeychainCopySearchList(searchList unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopySearchList(searchList)
}/* debug [functions.gen.go/function]: SecKeychainCopySearchList */

// Obtains a keychain’s settings.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Obtains a keychain’s settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCopySettings(_:_:)
func SecKeychainCopySettings(keychain SecKeychainRef, outSettings unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCopySettings(keychain, outSettings)
}/* debug [functions.gen.go/function]: SecKeychainCopySettings */

// Creates an empty keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Creates an empty keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCreate(_:_:_:_:_:_:)
func SecKeychainCreate(pathName unsafe.Pointer, passwordLength unsafe.Pointer, password unsafe.Pointer, promptUser unsafe.Pointer, initialAccess SecAccessRef, keychain unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainCreate(pathName, passwordLength, password, promptUser, initialAccess, keychain)
}/* debug [functions.gen.go/function]: SecKeychainCreate */

// Deletes one or more keychains from the default keychain search list, and removes the keychain itself if it is a file.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Deletes one or more keychains from the default keychain search list, and removes the keychain itself if it is a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainDelete(_:)
func SecKeychainDelete(keychainOrArray SecKeychainRef) unsafe.Pointer {
	return _SecKeychainDelete(keychainOrArray)
}/* debug [functions.gen.go/function]: SecKeychainDelete */

// Finds the first generic password based on the attributes passed.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Finds the first generic password based on the attributes passed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainFindGenericPassword(_:_:_:_:_:_:_:_:)
func SecKeychainFindGenericPassword(keychainOrArray TypeRef, serviceNameLength unsafe.Pointer, serviceName unsafe.Pointer, accountNameLength unsafe.Pointer, accountName unsafe.Pointer, passwordLength unsafe.Pointer, passwordData unsafe.Pointer, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainFindGenericPassword(keychainOrArray, serviceNameLength, serviceName, accountNameLength, accountName, passwordLength, passwordData, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainFindGenericPassword */

// Finds the first Internet password based on the attributes passed.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Finds the first Internet password based on the attributes passed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainFindInternetPassword(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func SecKeychainFindInternetPassword(keychainOrArray TypeRef, serverNameLength unsafe.Pointer, serverName unsafe.Pointer, securityDomainLength unsafe.Pointer, securityDomain unsafe.Pointer, accountNameLength unsafe.Pointer, accountName unsafe.Pointer, pathLength unsafe.Pointer, path unsafe.Pointer, port unsafe.Pointer, protocol_ SecProtocolType, authenticationType SecAuthenticationType, passwordLength unsafe.Pointer, passwordData unsafe.Pointer, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainFindInternetPassword(keychainOrArray, serverNameLength, serverName, securityDomainLength, securityDomain, accountNameLength, accountName, pathLength, path, port, protocol_, authenticationType, passwordLength, passwordData, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainFindInternetPassword */

// Releases the memory acquired by calling the function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Releases the memory acquired by calling the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainFreeAttributeInfo(_:)
func SecKeychainFreeAttributeInfo(info unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainFreeAttributeInfo(info)
}/* debug [functions.gen.go/function]: SecKeychainFreeAttributeInfo */

// Determines the path of a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Determines the path of a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetPath(_:_:_:)
func SecKeychainGetPath(keychain SecKeychainRef, ioPathLength unsafe.Pointer, pathName unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetPath(keychain, ioPathLength, pathName)
}/* debug [functions.gen.go/function]: SecKeychainGetPath */

// Gets the current keychain preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Gets the current keychain preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetPreferenceDomain(_:)
func SecKeychainGetPreferenceDomain(domain unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetPreferenceDomain(domain)
}/* debug [functions.gen.go/function]: SecKeychainGetPreferenceDomain */

// Retrieves status information of a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves status information of a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetStatus(_:_:)
func SecKeychainGetStatus(keychain SecKeychainRef, keychainStatus unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetStatus(keychain, keychainStatus)
}/* debug [functions.gen.go/function]: SecKeychainGetStatus */

// Returns the unique identifier of the opaque type to which a keychain object belongs.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Returns the unique identifier of the opaque type to which a keychain object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetTypeID()
func SecKeychainGetTypeID() TypeID {
	return _SecKeychainGetTypeID()
}/* debug [functions.gen.go/function]: SecKeychainGetTypeID */

// Indicates whether keychain services functions that normally display a user interaction are allowed to do so.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Indicates whether keychain services functions that normally display a user interaction are allowed to do so.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetUserInteractionAllowed(_:)
func SecKeychainGetUserInteractionAllowed(state unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetUserInteractionAllowed(state)
}/* debug [functions.gen.go/function]: SecKeychainGetUserInteractionAllowed */

// Determines the version of keychain services installed on the user’s system.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Determines the version of keychain services installed on the user’s system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainGetVersion(_:)
func SecKeychainGetVersion(returnVers unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainGetVersion(returnVers)
}/* debug [functions.gen.go/function]: SecKeychainGetVersion */

// Retrieves the access of a given keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves the access of a given keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyAccess(_:_:)
func SecKeychainItemCopyAccess(itemRef SecKeychainItemRef, access unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyAccess(itemRef, access)
}/* debug [functions.gen.go/function]: SecKeychainItemCopyAccess */

// Retrieves the data and/or attributes stored in the given keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Retrieves the data and/or attributes stored in the given keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyAttributesAndData(_:_:_:_:_:_:)
func SecKeychainItemCopyAttributesAndData(itemRef SecKeychainItemRef, info unsafe.Pointer, itemClass unsafe.Pointer, attrList unsafe.Pointer, length unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyAttributesAndData(itemRef, info, itemClass, attrList, length, outData)
}/* debug [functions.gen.go/function]: SecKeychainItemCopyAttributesAndData */

// Copies the data and attributes stored in the given keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Copies the data and attributes stored in the given keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyContent(_:_:_:_:_:)
func SecKeychainItemCopyContent(itemRef SecKeychainItemRef, itemClass unsafe.Pointer, attrList unsafe.Pointer, length unsafe.Pointer, outData unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyContent(itemRef, itemClass, attrList, length, outData)
}/* debug [functions.gen.go/function]: SecKeychainItemCopyContent */

// Provides a keychain item reference, given a persistent reference.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Provides a keychain item reference, given a persistent reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyFromPersistentReference(_:_:)
func SecKeychainItemCopyFromPersistentReference(persistentItemRef DataRef, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyFromPersistentReference(persistentItemRef, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainItemCopyFromPersistentReference */

// Returns the keychain object of a given keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Returns the keychain object of a given keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyKeychain(_:_:)
func SecKeychainItemCopyKeychain(itemRef SecKeychainItemRef, keychainRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyKeychain(itemRef, keychainRef)
}/* debug [functions.gen.go/function]: SecKeychainItemCopyKeychain */

// Copies a keychain item from one keychain to another.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Copies a keychain item from one keychain to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCreateCopy(_:_:_:_:)
func SecKeychainItemCreateCopy(itemRef SecKeychainItemRef, destKeychainRef SecKeychainRef, initialAccess SecAccessRef, itemCopy unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCreateCopy(itemRef, destKeychainRef, initialAccess, itemCopy)
}/* debug [functions.gen.go/function]: SecKeychainItemCreateCopy */

// Creates a new keychain item from the supplied parameters.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Creates a new keychain item from the supplied parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCreateFromContent(_:_:_:_:_:_:_:)
func SecKeychainItemCreateFromContent(itemClass SecItemClass, attrList unsafe.Pointer, length unsafe.Pointer, data unsafe.Pointer, keychainRef SecKeychainRef, initialAccess SecAccessRef, itemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCreateFromContent(itemClass, attrList, length, data, keychainRef, initialAccess, itemRef)
}/* debug [functions.gen.go/function]: SecKeychainItemCreateFromContent */

// Creates a persistent reference for a keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Creates a persistent reference for a keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCreatePersistentReference(_:_:)
func SecKeychainItemCreatePersistentReference(itemRef SecKeychainItemRef, persistentItemRef unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCreatePersistentReference(itemRef, persistentItemRef)
}/* debug [functions.gen.go/function]: SecKeychainItemCreatePersistentReference */

// Deletes a keychain item from the default keychain’s permanent data store.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Deletes a keychain item from the default keychain’s permanent data store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemDelete(_:)
func SecKeychainItemDelete(itemRef SecKeychainItemRef) unsafe.Pointer {
	return _SecKeychainItemDelete(itemRef)
}/* debug [functions.gen.go/function]: SecKeychainItemDelete */

// Releases the memory used by the keychain attribute list and/or the keychain data retrieved in a call to .
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Releases the memory used by the keychain attribute list and/or the keychain data retrieved in a call to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemFreeAttributesAndData(_:_:)
func SecKeychainItemFreeAttributesAndData(attrList unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemFreeAttributesAndData(attrList, data)
}/* debug [functions.gen.go/function]: SecKeychainItemFreeAttributesAndData */

// Releases the memory used by the keychain attribute list and the keychain data retrieved in a call to the function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Releases the memory used by the keychain attribute list and the keychain data retrieved in a call to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemFreeContent(_:_:)
func SecKeychainItemFreeContent(attrList unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemFreeContent(attrList, data)
}/* debug [functions.gen.go/function]: SecKeychainItemFreeContent */

// Returns the unique identifier of the opaque type to which a keychain item object belongs.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Returns the unique identifier of the opaque type to which a keychain item object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemGetTypeID()
func SecKeychainItemGetTypeID() TypeID {
	return _SecKeychainItemGetTypeID()
}/* debug [functions.gen.go/function]: SecKeychainItemGetTypeID */

// Updates an existing keychain item after changing its attributes or data.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Updates an existing keychain item after changing its attributes or data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemModifyAttributesAndData(_:_:_:_:)
func SecKeychainItemModifyAttributesAndData(itemRef SecKeychainItemRef, attrList unsafe.Pointer, length unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemModifyAttributesAndData(itemRef, attrList, length, data)
}/* debug [functions.gen.go/function]: SecKeychainItemModifyAttributesAndData */

// Updates an existing keychain item after changing its attributes and/or data.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Updates an existing keychain item after changing its attributes and/or data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemModifyContent(_:_:_:_:)
func SecKeychainItemModifyContent(itemRef SecKeychainItemRef, attrList unsafe.Pointer, length unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemModifyContent(itemRef, attrList, length, data)
}/* debug [functions.gen.go/function]: SecKeychainItemModifyContent */

// Sets the access of a given keychain item.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the access of a given keychain item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemSetAccess(_:_:)
func SecKeychainItemSetAccess(itemRef SecKeychainItemRef, access SecAccessRef) unsafe.Pointer {
	return _SecKeychainItemSetAccess(itemRef, access)
}/* debug [functions.gen.go/function]: SecKeychainItemSetAccess */

// Locks a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Locks a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainLock(_:)
func SecKeychainLock(keychain SecKeychainRef) unsafe.Pointer {
	return _SecKeychainLock(keychain)
}/* debug [functions.gen.go/function]: SecKeychainLock */

// Locks all keychains belonging to the current user.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Locks all keychains belonging to the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainLockAll()
func SecKeychainLockAll() unsafe.Pointer {
	return _SecKeychainLockAll()
}/* debug [functions.gen.go/function]: SecKeychainLockAll */

// Opens a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Opens a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainOpen(_:_:)
func SecKeychainOpen(pathName unsafe.Pointer, keychain unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainOpen(pathName, keychain)
}/* debug [functions.gen.go/function]: SecKeychainOpen */

// Unregisters your keychain event callback function.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Unregisters your keychain event callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainRemoveCallback(_:)
func SecKeychainRemoveCallback(callbackFunction SecKeychainCallback) unsafe.Pointer {
	return _SecKeychainRemoveCallback(callbackFunction)
}/* debug [functions.gen.go/function]: SecKeychainRemoveCallback */

// Sets the application access for a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the application access for a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetAccess(_:_:)
func SecKeychainSetAccess(keychain SecKeychainRef, access SecAccessRef) unsafe.Pointer {
	return _SecKeychainSetAccess(keychain, access)
}/* debug [functions.gen.go/function]: SecKeychainSetAccess */

// Sets the default keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the default keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetDefault(_:)
func SecKeychainSetDefault(keychain SecKeychainRef) unsafe.Pointer {
	return _SecKeychainSetDefault(keychain)
}/* debug [functions.gen.go/function]: SecKeychainSetDefault */

// Sets the default keychain for a specified preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the default keychain for a specified preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetDomainDefault(_:_:)
func SecKeychainSetDomainDefault(domain SecPreferencesDomain, keychain SecKeychainRef) unsafe.Pointer {
	return _SecKeychainSetDomainDefault(domain, keychain)
}/* debug [functions.gen.go/function]: SecKeychainSetDomainDefault */

// Sets the keychain search list for a specified preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the keychain search list for a specified preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetDomainSearchList(_:_:)
func SecKeychainSetDomainSearchList(domain SecPreferencesDomain, searchList ArrayRef) unsafe.Pointer {
	return _SecKeychainSetDomainSearchList(domain, searchList)
}/* debug [functions.gen.go/function]: SecKeychainSetDomainSearchList */

// Sets the keychain preference domain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Sets the keychain preference domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetPreferenceDomain(_:)
func SecKeychainSetPreferenceDomain(domain SecPreferencesDomain) unsafe.Pointer {
	return _SecKeychainSetPreferenceDomain(domain)
}/* debug [functions.gen.go/function]: SecKeychainSetPreferenceDomain */

// Specifies the list of keychains to use in the default keychain search list.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Specifies the list of keychains to use in the default keychain search list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetSearchList(_:)
func SecKeychainSetSearchList(searchList ArrayRef) unsafe.Pointer {
	return _SecKeychainSetSearchList(searchList)
}/* debug [functions.gen.go/function]: SecKeychainSetSearchList */

// Changes the settings of a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Changes the settings of a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetSettings(_:_:)
func SecKeychainSetSettings(keychain SecKeychainRef, newSettings unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainSetSettings(keychain, newSettings)
}/* debug [functions.gen.go/function]: SecKeychainSetSettings */

// Enables or disables the user interface for keychain services functions that automatically display a user interface.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Enables or disables the user interface for keychain services functions that automatically display a user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSetUserInteractionAllowed(_:)
func SecKeychainSetUserInteractionAllowed(state unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainSetUserInteractionAllowed(state)
}/* debug [functions.gen.go/function]: SecKeychainSetUserInteractionAllowed */

// Unlocks a keychain.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Unlocks a keychain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainUnlock(_:_:_:_:)
func SecKeychainUnlock(keychain SecKeychainRef, passwordLength unsafe.Pointer, password unsafe.Pointer, usePassword unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainUnlock(keychain, passwordLength, password, usePassword)
}/* debug [functions.gen.go/function]: SecKeychainUnlock */

// Gets the attributes of a given key.
//
// Added in macOS 10.12.
// Gets the attributes of a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyAttributes(_:)
func SecKeyCopyAttributes(key SecKeyRef) DictionaryRef {
	return _SecKeyCopyAttributes(key)
}/* debug [functions.gen.go/function]: SecKeyCopyAttributes */

// Returns an external representation of the given key suitable for the key’s type.
//
// Added in macOS 10.12.
// Returns an external representation of the given key suitable for the key’s type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyExternalRepresentation(_:_:)
func SecKeyCopyExternalRepresentation(key SecKeyRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyCopyExternalRepresentation(key, error_)
}/* debug [functions.gen.go/function]: SecKeyCopyExternalRepresentation */

// Performs the Diffie-Hellman style of key exchange with optional key-derivation steps.
//
// Added in macOS 10.12.
// Performs the Diffie-Hellman style of key exchange with optional key-derivation steps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyKeyExchangeResult(_:_:_:_:_:)
func SecKeyCopyKeyExchangeResult(privateKey SecKeyRef, algorithm SecKeyAlgorithm, publicKey SecKeyRef, parameters DictionaryRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyCopyKeyExchangeResult(privateKey, algorithm, publicKey, parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyCopyKeyExchangeResult */

// Gets the public key associated with the given private key.
//
// Added in macOS 10.12.
// Gets the public key associated with the given private key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyPublicKey(_:)
func SecKeyCopyPublicKey(key SecKeyRef) SecKeyRef {
	return _SecKeyCopyPublicKey(key)
}/* debug [functions.gen.go/function]: SecKeyCopyPublicKey */

// Decrypts a block of data using a private key and specified algorithm.
//
// Added in macOS 10.12.
// Decrypts a block of data using a private key and specified algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateDecryptedData(_:_:_:_:)
func SecKeyCreateDecryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, ciphertext DataRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyCreateDecryptedData(key, algorithm, ciphertext, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateDecryptedData */

// Encrypts a block of data using a public key and specified algorithm.
//
// Added in macOS 10.12.
// Encrypts a block of data using a public key and specified algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateEncryptedData(_:_:_:_:)
func SecKeyCreateEncryptedData(key SecKeyRef, algorithm SecKeyAlgorithm, plaintext DataRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyCreateEncryptedData(key, algorithm, plaintext, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateEncryptedData */

// Constructs a SecKeyRef object for a symmetric key.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Constructs a SecKeyRef object for a symmetric key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateFromData(_:_:_:)
func SecKeyCreateFromData(parameters DictionaryRef, keyData DataRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyCreateFromData(parameters, keyData, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateFromData */

// Generates a new public-private key pair.
//
// Added in macOS 10.12.
// Generates a new public-private key pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateRandomKey(_:_:)
func SecKeyCreateRandomKey(parameters DictionaryRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyCreateRandomKey(parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateRandomKey */

// Creates the cryptographic signature for a block of data using a private key and specified algorithm.
//
// Added in macOS 10.12.
// Creates the cryptographic signature for a block of data using a private key and specified algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateSignature(_:_:_:_:)
func SecKeyCreateSignature(key SecKeyRef, algorithm SecKeyAlgorithm, dataToSign DataRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyCreateSignature(key, algorithm, dataToSign, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateSignature */

// Restores a key from an external representation of that key.
//
// Added in macOS 10.12.
// Restores a key from an external representation of that key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateWithData(_:_:_:)
func SecKeyCreateWithData(keyData DataRef, attributes DictionaryRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyCreateWithData(keyData, attributes, error_)
}/* debug [functions.gen.go/function]: SecKeyCreateWithData */

// Decrypts a block of ciphertext.

// Decrypts a block of ciphertext.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyDecrypt(_:_:_:_:_:_:)
func SecKeyDecrypt(key SecKeyRef, padding SecPadding, cipherText unsafe.Pointer, cipherTextLen uintptr, plainText unsafe.Pointer, plainTextLen unsafe.Pointer) unsafe.Pointer {
	return _SecKeyDecrypt(key, padding, cipherText, cipherTextLen, plainText, plainTextLen)
}/* debug [functions.gen.go/function]: SecKeyDecrypt */

// Returns a key object in which the key data is derived from a password.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Returns a key object in which the key data is derived from a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyDeriveFromPassword(_:_:_:)
func SecKeyDeriveFromPassword(password StringRef, parameters DictionaryRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyDeriveFromPassword(password, parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyDeriveFromPassword */

// Encrypts a block of plaintext.

// Encrypts a block of plaintext.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyEncrypt(_:_:_:_:_:_:)
func SecKeyEncrypt(key SecKeyRef, padding SecPadding, plainText unsafe.Pointer, plainTextLen uintptr, cipherText unsafe.Pointer, cipherTextLen unsafe.Pointer) unsafe.Pointer {
	return _SecKeyEncrypt(key, padding, plainText, plainTextLen, cipherText, cipherTextLen)
}/* debug [functions.gen.go/function]: SecKeyEncrypt */

// Creates an asymmetric key pair.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Creates an asymmetric key pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGeneratePair(_:_:_:)
func SecKeyGeneratePair(parameters DictionaryRef, publicKey unsafe.Pointer, privateKey unsafe.Pointer) unsafe.Pointer {
	return _SecKeyGeneratePair(parameters, publicKey, privateKey)
}/* debug [functions.gen.go/function]: SecKeyGeneratePair */

// Generates a public/private key pair.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Generates a public/private key pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGeneratePairAsync(_:_:_:)
func SecKeyGeneratePairAsync(parameters DictionaryRef, deliveryQueue unsafe.Pointer, result unsafe.Pointer) {
	_SecKeyGeneratePairAsync(parameters, deliveryQueue, result)
}/* debug [functions.gen.go/function]: SecKeyGeneratePairAsync */

// Generates a random symmetric key.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Generates a random symmetric key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGenerateSymmetric(_:_:)
func SecKeyGenerateSymmetric(parameters DictionaryRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyGenerateSymmetric(parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyGenerateSymmetric */

// Gets the block length associated with a cryptographic key.
//
// Added in macOS 10.6.
// Gets the block length associated with a cryptographic key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGetBlockSize(_:)
func SecKeyGetBlockSize(key SecKeyRef) uintptr {
	return _SecKeyGetBlockSize(key)
}/* debug [functions.gen.go/function]: SecKeyGetBlockSize */

// Returns the unique identifier of the opaque type to which a key object belongs.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which a key object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyGetTypeID()
func SecKeyGetTypeID() TypeID {
	return _SecKeyGetTypeID()
}/* debug [functions.gen.go/function]: SecKeyGetTypeID */

// Returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm.
//
// Added in macOS 10.12.
// Returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyIsAlgorithmSupported(_:_:_:)
func SecKeyIsAlgorithmSupported(key SecKeyRef, operation SecKeyOperationType, algorithm SecKeyAlgorithm) unsafe.Pointer {
	return _SecKeyIsAlgorithmSupported(key, operation, algorithm)
}/* debug [functions.gen.go/function]: SecKeyIsAlgorithmSupported */

// Generates a digital signature for a block of data.

// Generates a digital signature for a block of data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyRawSign(_:_:_:_:_:_:)
func SecKeyRawSign(key SecKeyRef, padding SecPadding, dataToSign unsafe.Pointer, dataToSignLen uintptr, sig unsafe.Pointer, sigLen unsafe.Pointer) unsafe.Pointer {
	return _SecKeyRawSign(key, padding, dataToSign, dataToSignLen, sig, sigLen)
}/* debug [functions.gen.go/function]: SecKeyRawSign */

// Verifies a digital signature.

// Verifies a digital signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyRawVerify(_:_:_:_:_:_:)
func SecKeyRawVerify(key SecKeyRef, padding SecPadding, signedData unsafe.Pointer, signedDataLen uintptr, sig unsafe.Pointer, sigLen uintptr) unsafe.Pointer {
	return _SecKeyRawVerify(key, padding, signedData, signedDataLen, sig, sigLen)
}/* debug [functions.gen.go/function]: SecKeyRawVerify */

// Unwraps a wrapped symmetric key.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Unwraps a wrapped symmetric key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUnwrapSymmetric(_:_:_:_:)
func SecKeyUnwrapSymmetric(keyToUnwrap unsafe.Pointer, unwrappingKey SecKeyRef, parameters DictionaryRef, error_ unsafe.Pointer) SecKeyRef {
	return _SecKeyUnwrapSymmetric(keyToUnwrap, unwrappingKey, parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyUnwrapSymmetric */

// Verifies the cryptographic signature of a block of data using a public key and specified algorithm.
//
// Added in macOS 10.12.
// Verifies the cryptographic signature of a block of data using a public key and specified algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyVerifySignature(_:_:_:_:_:)
func SecKeyVerifySignature(key SecKeyRef, algorithm SecKeyAlgorithm, signedData DataRef, signature DataRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyVerifySignature(key, algorithm, signedData, signature, error_)
}/* debug [functions.gen.go/function]: SecKeyVerifySignature */

// Wraps a symmetric key with another key.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Wraps a symmetric key with another key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyWrapSymmetric(_:_:_:_:)
func SecKeyWrapSymmetric(keyToWrap SecKeyRef, wrappingKey SecKeyRef, parameters DictionaryRef, error_ unsafe.Pointer) DataRef {
	return _SecKeyWrapSymmetric(keyToWrap, wrappingKey, parameters, error_)
}/* debug [functions.gen.go/function]: SecKeyWrapSymmetric */

// Returns the identities and certificates in a PKCS #12-formatted blob.
//
// Added in macOS 10.6.
// Returns the identities and certificates in a PKCS #12-formatted blob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPKCS12Import(_:_:_:)
func SecPKCS12Import(pkcs12_data DataRef, options DictionaryRef, items unsafe.Pointer) unsafe.Pointer {
	return _SecPKCS12Import(pkcs12_data, options, items)
}/* debug [functions.gen.go/function]: SecPKCS12Import */

// Returns a dictionary containing a policy’s properties.
//
// Added in macOS 10.7.
// Returns a dictionary containing a policy’s properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCopyProperties(_:)
func SecPolicyCopyProperties(policyRef SecPolicyRef) DictionaryRef {
	return _SecPolicyCopyProperties(policyRef)
}/* debug [functions.gen.go/function]: SecPolicyCopyProperties */

// Returns a policy object for the default X.509 policy.
//
// Added in macOS 10.6.
// Returns a policy object for the default X.509 policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateBasicX509()
func SecPolicyCreateBasicX509() SecPolicyRef {
	return _SecPolicyCreateBasicX509()
}/* debug [functions.gen.go/function]: SecPolicyCreateBasicX509 */

// Returns a policy object for checking revocation of certificates.
//
// Added in macOS 10.9.
// Returns a policy object for checking revocation of certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateRevocation(_:)
func SecPolicyCreateRevocation(revocationFlags OptionFlags) SecPolicyRef {
	return _SecPolicyCreateRevocation(revocationFlags)
}/* debug [functions.gen.go/function]: SecPolicyCreateRevocation */

// Returns a policy object for evaluating SSL certificate chains.
//
// Added in macOS 10.6.
// Returns a policy object for evaluating SSL certificate chains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateSSL(_:_:)
func SecPolicyCreateSSL(server unsafe.Pointer, hostname StringRef) SecPolicyRef {
	return _SecPolicyCreateSSL(server, hostname)
}/* debug [functions.gen.go/function]: SecPolicyCreateSSL */

// Returns a policy object based on an object identifier for the policy type.
//
// Added in macOS 10.9.
// Returns a policy object based on an object identifier for the policy type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateWithProperties(_:_:)
func SecPolicyCreateWithProperties(policyIdentifier TypeRef, properties DictionaryRef) SecPolicyRef {
	return _SecPolicyCreateWithProperties(policyIdentifier, properties)
}/* debug [functions.gen.go/function]: SecPolicyCreateWithProperties */

// Returns the unique identifier of the opaque type to which a policy object belongs.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which a policy object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyGetTypeID()
func SecPolicyGetTypeID() TypeID {
	return _SecPolicyGetTypeID()
}/* debug [functions.gen.go/function]: SecPolicyGetTypeID */

// Generates an array of cryptographically secure random bytes.
//
// Added in macOS 10.7.
// Generates an array of cryptographically secure random bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRandomCopyBytes(_:_:_:)
func SecRandomCopyBytes(rnd SecRandomRef, count uintptr, bytes unsafe.Pointer) int {
	return _SecRandomCopyBytes(rnd, count, bytes)
}/* debug [functions.gen.go/function]: SecRandomCopyBytes */

// Asynchronously obtains one or more shared passwords for a website.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 11.0.
// Asynchronously obtains one or more shared passwords for a website.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequestSharedWebCredential(_:_:_:)
func SecRequestSharedWebCredential(fqdn StringRef, account StringRef) {
	_SecRequestSharedWebCredential(fqdn, account)
}/* debug [functions.gen.go/function]: SecRequestSharedWebCredential */

// Extracts a binary form of a code requirement from a code requirement object.
//
// Added in macOS 10.0.
// Extracts a binary form of a code requirement from a code requirement object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCopyData(_:_:_:)
func SecRequirementCopyData(requirement SecRequirementRef, flags SecCSFlags, data unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCopyData(requirement, flags, data)
}/* debug [functions.gen.go/function]: SecRequirementCopyData */

// Converts a code requirement object into text form.
//
// Added in macOS 10.0.
// Converts a code requirement object into text form.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCopyString(_:_:_:)
func SecRequirementCopyString(requirement SecRequirementRef, flags SecCSFlags, text unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCopyString(requirement, flags, text)
}/* debug [functions.gen.go/function]: SecRequirementCopyString */

// Creates a code requirement object from the binary form of a code requirement.
//
// Added in macOS 10.0.
// Creates a code requirement object from the binary form of a code requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithData(_:_:_:)
func SecRequirementCreateWithData(data DataRef, flags SecCSFlags, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithData(data, flags, requirement)
}/* debug [functions.gen.go/function]: SecRequirementCreateWithData */

// Creates a code requirement object by compiling a valid text representation of a code requirement.
//
// Added in macOS 10.0.
// Creates a code requirement object by compiling a valid text representation of a code requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithString(_:_:_:)
func SecRequirementCreateWithString(text StringRef, flags SecCSFlags, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithString(text, flags, requirement)
}/* debug [functions.gen.go/function]: SecRequirementCreateWithString */

// Creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure.
//
// Added in macOS 10.0.
// Creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithStringAndErrors(_:_:_:_:)
func SecRequirementCreateWithStringAndErrors(text StringRef, flags SecCSFlags, errors unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithStringAndErrors(text, flags, errors, requirement)
}/* debug [functions.gen.go/function]: SecRequirementCreateWithStringAndErrors */

// Returns the unique identifier of the opaque type to which a code requirement object belongs.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a code requirement object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementGetTypeID()
func SecRequirementGetTypeID() TypeID {
	return _SecRequirementGetTypeID()
}/* debug [functions.gen.go/function]: SecRequirementGetTypeID */

// Creates a signing transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a signing transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecSignTransformCreate(_:_:)
func SecSignTransformCreate(key SecKeyRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecSignTransformCreate(key, error_)
}/* debug [functions.gen.go/function]: SecSignTransformCreate */

// Validates a static code object.
//
// Added in macOS 10.0.
// Validates a static code object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidity(_:_:_:)
func SecStaticCodeCheckValidity(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef) unsafe.Pointer {
	return _SecStaticCodeCheckValidity(staticCode, flags, requirement)
}/* debug [functions.gen.go/function]: SecStaticCodeCheckValidity */

// Performs static validation of static signed code and returns detailed error information in the case of failure.
//
// Added in macOS 10.0.
// Performs static validation of static signed code and returns detailed error information in the case of failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidityWithErrors(_:_:_:_:)
func SecStaticCodeCheckValidityWithErrors(staticCode SecStaticCodeRef, flags SecCSFlags, requirement SecRequirementRef, errors unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors)
}/* debug [functions.gen.go/function]: SecStaticCodeCheckValidityWithErrors */

// Creates a static code object representing the code at a specified file system path.
//
// Added in macOS 10.0.
// Creates a static code object representing the code at a specified file system path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPath(_:_:_:)
func SecStaticCodeCreateWithPath(path URLRef, flags SecCSFlags, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCreateWithPath(path, flags, staticCode)
}/* debug [functions.gen.go/function]: SecStaticCodeCreateWithPath */

// Creates a static code object representing the code at a specified file system path using an attributes dictionary.
//
// Added in macOS 10.0.
// Creates a static code object representing the code at a specified file system path using an attributes dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPathAndAttributes(_:_:_:_:)
func SecStaticCodeCreateWithPathAndAttributes(path URLRef, flags SecCSFlags, attributes DictionaryRef, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode)
}/* debug [functions.gen.go/function]: SecStaticCodeCreateWithPathAndAttributes */

// Returns the unique identifier of the opaque type to which a static code object belongs.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a static code object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeGetTypeID()
func SecStaticCodeGetTypeID() TypeID {
	return _SecStaticCodeGetTypeID()
}/* debug [functions.gen.go/function]: SecStaticCodeGetTypeID */

// Returns the value of the code signing identifier.
//
// Added in macOS 10.0.
// Returns the value of the code signing identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopySigningIdentifier(_:_:)
func SecTaskCopySigningIdentifier(task SecTaskRef, error_ unsafe.Pointer) StringRef {
	return _SecTaskCopySigningIdentifier(task, error_)
}/* debug [functions.gen.go/function]: SecTaskCopySigningIdentifier */

// Returns the value of a single entitlement for the represented task.
//
// Added in macOS 10.0.
// Returns the value of a single entitlement for the represented task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopyValueForEntitlement(_:_:_:)
func SecTaskCopyValueForEntitlement(task SecTaskRef, entitlement StringRef, error_ unsafe.Pointer) TypeRef {
	return _SecTaskCopyValueForEntitlement(task, entitlement, error_)
}/* debug [functions.gen.go/function]: SecTaskCopyValueForEntitlement */

// Returns the values of multiple entitlements for the represented task.
//
// Added in macOS 10.0.
// Returns the values of multiple entitlements for the represented task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopyValuesForEntitlements(_:_:_:)
func SecTaskCopyValuesForEntitlements(task SecTaskRef, entitlements ArrayRef, error_ unsafe.Pointer) DictionaryRef {
	return _SecTaskCopyValuesForEntitlements(task, entitlements, error_)
}/* debug [functions.gen.go/function]: SecTaskCopyValuesForEntitlements */

// Creates a task object for the current task.
//
// Added in macOS 10.0.
// Creates a task object for the current task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCreateFromSelf(_:)
func SecTaskCreateFromSelf(allocator AllocatorRef) SecTaskRef {
	return _SecTaskCreateFromSelf(allocator)
}/* debug [functions.gen.go/function]: SecTaskCreateFromSelf */

// Creates a task object for the task that sent the Mach message represented by the audit token.
//
// Added in macOS 10.0.
// Creates a task object for the task that sent the Mach message represented by the audit token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCreateWithAuditToken(_:_:)
func SecTaskCreateWithAuditToken(allocator AllocatorRef, token unsafe.Pointer) SecTaskRef {
	return _SecTaskCreateWithAuditToken(allocator, token)
}/* debug [functions.gen.go/function]: SecTaskCreateWithAuditToken */

// Returns the unique identifier of the opaque type to which a task object belongs.
//
// Added in macOS 10.0.
// Returns the unique identifier of the opaque type to which a task object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskGetTypeID()
func SecTaskGetTypeID() TypeID {
	return _SecTaskGetTypeID()
}/* debug [functions.gen.go/function]: SecTaskGetTypeID */

// Chains transforms together.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Chains transforms together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformConnectTransforms(_:_:_:_:_:_:)
func SecTransformConnectTransforms(sourceTransformRef SecTransformRef, sourceAttributeName StringRef, destinationTransformRef SecTransformRef, destinationAttributeName StringRef, group SecGroupTransformRef, error_ unsafe.Pointer) SecGroupTransformRef {
	return _SecTransformConnectTransforms(sourceTransformRef, sourceAttributeName, destinationTransformRef, destinationAttributeName, group, error_)
}/* debug [functions.gen.go/function]: SecTransformConnectTransforms */

// Creates a dictionary that contains enough information to be able to recreate a transform.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Creates a dictionary that contains enough information to be able to recreate a transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCopyExternalRepresentation(_:)
func SecTransformCopyExternalRepresentation(transformRef SecTransformRef) DictionaryRef {
	return _SecTransformCopyExternalRepresentation(transformRef)
}/* debug [functions.gen.go/function]: SecTransformCopyExternalRepresentation */

// Creates a transform computation object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a transform computation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreate(_:_:)
func SecTransformCreate(name StringRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecTransformCreate(name, error_)
}/* debug [functions.gen.go/function]: SecTransformCreate */

// Creates a transform instance from a dictionary of parameters.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Creates a transform instance from a dictionary of parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateFromExternalRepresentation(_:_:)
func SecTransformCreateFromExternalRepresentation(dictionary DictionaryRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecTransformCreateFromExternalRepresentation(dictionary, error_)
}/* debug [functions.gen.go/function]: SecTransformCreateFromExternalRepresentation */

// Creates an object that acts as a container for a set of connected transforms.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Creates an object that acts as a container for a set of connected transforms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateGroupTransform()
func SecTransformCreateGroupTransform() SecGroupTransformRef {
	return _SecTransformCreateGroupTransform()
}/* debug [functions.gen.go/function]: SecTransformCreateGroupTransform */

// Creates a read transform from a read stream reference.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a read transform from a read stream reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateReadTransformWithReadStream(_:)
func SecTransformCreateReadTransformWithReadStream(inputStream ReadStreamRef) SecTransformRef {
	return _SecTransformCreateReadTransformWithReadStream(inputStream)
}/* debug [functions.gen.go/function]: SecTransformCreateReadTransformWithReadStream */

// Gets an attribute value from a custom transform.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Gets an attribute value from a custom transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCustomGetAttribute(_:_:_:)
func SecTransformCustomGetAttribute(ref SecTransformImplementationRef, attribute SecTransformStringOrAttributeRef, type_ SecTransformMetaAttributeType) TypeRef {
	return _SecTransformCustomGetAttribute(ref, attribute, type_)
}/* debug [functions.gen.go/function]: SecTransformCustomGetAttribute */

// Sets an attribute value on a custom transform.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Sets an attribute value on a custom transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCustomSetAttribute(_:_:_:_:)
func SecTransformCustomSetAttribute(ref SecTransformImplementationRef, attribute SecTransformStringOrAttributeRef, type_ SecTransformMetaAttributeType, value TypeRef) TypeRef {
	return _SecTransformCustomSetAttribute(ref, attribute, type_, value)
}/* debug [functions.gen.go/function]: SecTransformCustomSetAttribute */

// Executes a transform or transform group synchronously.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Executes a transform or transform group synchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformExecute(_:_:)
func SecTransformExecute(transformRef SecTransformRef, errorRef unsafe.Pointer) TypeRef {
	return _SecTransformExecute(transformRef, errorRef)
}/* debug [functions.gen.go/function]: SecTransformExecute */

// Executes transform or transform group asynchronously.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Executes transform or transform group asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformExecuteAsync(_:_:_:)
func SecTransformExecuteAsync(transformRef SecTransformRef, deliveryQueue unsafe.Pointer, deliveryBlock unsafe.Pointer) {
	_SecTransformExecuteAsync(transformRef, deliveryQueue, deliveryBlock)
}/* debug [functions.gen.go/function]: SecTransformExecuteAsync */

// Finds a member of a transform group by its name.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Finds a member of a transform group by its name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformFindByName(_:_:)
func SecTransformFindByName(transform SecGroupTransformRef, name StringRef) SecTransformRef {
	return _SecTransformFindByName(transform, name)
}/* debug [functions.gen.go/function]: SecTransformFindByName */

// Gets the current value of a transform attribute.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Gets the current value of a transform attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformGetAttribute(_:_:)
func SecTransformGetAttribute(transformRef SecTransformRef, key StringRef) TypeRef {
	return _SecTransformGetAttribute(transformRef, key)
}/* debug [functions.gen.go/function]: SecTransformGetAttribute */

// Returns the unique identifier of the opaque type to which a security transform object belongs.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Returns the unique identifier of the opaque type to which a security transform object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformGetTypeID()
func SecTransformGetTypeID() TypeID {
	return _SecTransformGetTypeID()
}/* debug [functions.gen.go/function]: SecTransformGetTypeID */

// Returns an object from inside a ProcessData override that says that although no data is being returned the transform is still active and awaiting data.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Returns an object from inside a ProcessData override that says that although no data is being returned the transform is still active and awaiting data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformNoData()
func SecTransformNoData() TypeRef {
	return _SecTransformNoData()
}/* debug [functions.gen.go/function]: SecTransformNoData */

// Pushes a single value back for a specific attribute.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Pushes a single value back for a specific attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformPushbackAttribute(_:_:_:)
func SecTransformPushbackAttribute(ref SecTransformImplementationRef, attribute SecTransformStringOrAttributeRef, value TypeRef) TypeRef {
	return _SecTransformPushbackAttribute(ref, attribute, value)
}/* debug [functions.gen.go/function]: SecTransformPushbackAttribute */

// Registers a custom transform.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Registers a custom transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformRegister(_:_:_:)
func SecTransformRegister(uniqueName StringRef, createTransformFunction unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformRegister(uniqueName, createTransformFunction, error_)
}/* debug [functions.gen.go/function]: SecTransformRegister */

// Sets a static value for an attribute in a transform.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Sets a static value for an attribute in a transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetAttribute(_:_:_:_:)
func SecTransformSetAttribute(transformRef SecTransformRef, key StringRef, value TypeRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformSetAttribute(transformRef, key, value, error_)
}/* debug [functions.gen.go/function]: SecTransformSetAttribute */

// Requests a callback when an attribute is set.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Requests a callback when an attribute is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetAttributeAction(_:_:_:_:)
func SecTransformSetAttributeAction(ref SecTransformImplementationRef, action StringRef, attribute SecTransformStringOrAttributeRef, newAction unsafe.Pointer) ErrorRef {
	return _SecTransformSetAttributeAction(ref, action, attribute, newAction)
}/* debug [functions.gen.go/function]: SecTransformSetAttributeAction */

// Changes the way a custom transform processes data.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Changes the way a custom transform processes data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetDataAction(_:_:_:)
func SecTransformSetDataAction(ref SecTransformImplementationRef, action StringRef, newAction unsafe.Pointer) ErrorRef {
	return _SecTransformSetDataAction(ref, action, newAction)
}/* debug [functions.gen.go/function]: SecTransformSetDataAction */

// Changes the way that a transform deals with transform lifecycle behaviors.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Changes the way that a transform deals with transform lifecycle behaviors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetTransformAction(_:_:_:)
func SecTransformSetTransformAction(ref SecTransformImplementationRef, action StringRef, newAction unsafe.Pointer) ErrorRef {
	return _SecTransformSetTransformAction(ref, action, newAction)
}/* debug [functions.gen.go/function]: SecTransformSetTransformAction */

// Retrieves the anchor (root) certificates stored by macOS.
//
// Added in macOS 10.3.
// Retrieves the anchor (root) certificates stored by macOS.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyAnchorCertificates(_:)
func SecTrustCopyAnchorCertificates(anchors unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCopyAnchorCertificates(anchors)
}/* debug [functions.gen.go/function]: SecTrustCopyAnchorCertificates */

// Retrieves the custom anchor certificates, if any, used by a given trust.
//
// Added in macOS 10.5.
// Retrieves the custom anchor certificates, if any, used by a given trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyCustomAnchorCertificates(_:_:)
func SecTrustCopyCustomAnchorCertificates(trust SecTrustRef, anchors unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCopyCustomAnchorCertificates(trust, anchors)
}/* debug [functions.gen.go/function]: SecTrustCopyCustomAnchorCertificates */

// Returns an opaque cookie containing exceptions to trust policies that will allow future evaluations of the current certificate to succeed.
//
// Added in macOS 10.9.
// Returns an opaque cookie containing exceptions to trust policies that will allow future evaluations of the current certificate to succeed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyExceptions(_:)
func SecTrustCopyExceptions(trust SecTrustRef) DataRef {
	return _SecTrustCopyExceptions(trust)
}/* debug [functions.gen.go/function]: SecTrustCopyExceptions */

// Retrieves the policies used by a given trust management object.
//
// Added in macOS 10.3.
// Retrieves the policies used by a given trust management object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyPolicies(_:_:)
func SecTrustCopyPolicies(trust SecTrustRef, policies unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCopyPolicies(trust, policies)
}/* debug [functions.gen.go/function]: SecTrustCopyPolicies */

// Returns an array containing the properties of a trust object.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Returns an array containing the properties of a trust object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyProperties(_:)
func SecTrustCopyProperties(trust SecTrustRef) ArrayRef {
	return _SecTrustCopyProperties(trust)
}/* debug [functions.gen.go/function]: SecTrustCopyProperties */

// Returns the public key for a leaf certificate after it has been evaluated.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.7.
// Returns the public key for a leaf certificate after it has been evaluated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyPublicKey(_:)
func SecTrustCopyPublicKey(trust SecTrustRef) SecKeyRef {
	return _SecTrustCopyPublicKey(trust)
}/* debug [functions.gen.go/function]: SecTrustCopyPublicKey */

// Returns a dictionary containing information about an evaluated trust.
//
// Added in macOS 10.9.
// Returns a dictionary containing information about an evaluated trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyResult(_:)
func SecTrustCopyResult(trust SecTrustRef) DictionaryRef {
	return _SecTrustCopyResult(trust)
}/* debug [functions.gen.go/function]: SecTrustCopyResult */

// Creates a trust management object based on certificates and policies.
//
// Added in macOS 10.3.
// Creates a trust management object based on certificates and policies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCreateWithCertificates(_:_:_:)
func SecTrustCreateWithCertificates(certificates TypeRef, policies TypeRef, trust unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCreateWithCertificates(certificates, policies, trust)
}/* debug [functions.gen.go/function]: SecTrustCreateWithCertificates */

// Retrieves the data of a trusted app instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Retrieves the data of a trusted app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustedApplicationCopyData(_:_:)
func SecTrustedApplicationCopyData(appRef SecTrustedApplicationRef, data unsafe.Pointer) unsafe.Pointer {
	return _SecTrustedApplicationCopyData(appRef, data)
}/* debug [functions.gen.go/function]: SecTrustedApplicationCopyData */

// Creates a trusted app instance based on the app at the given path in the file system.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates a trusted app instance based on the app at the given path in the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustedApplicationCreateFromPath(_:_:)
func SecTrustedApplicationCreateFromPath(path unsafe.Pointer, app unsafe.Pointer) unsafe.Pointer {
	return _SecTrustedApplicationCreateFromPath(path, app)
}/* debug [functions.gen.go/function]: SecTrustedApplicationCreateFromPath */

// Returns the unique identifier of the opaque type to which a trusted app instance belongs.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Returns the unique identifier of the opaque type to which a trusted app instance belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustedApplicationGetTypeID()
func SecTrustedApplicationGetTypeID() TypeID {
	return _SecTrustedApplicationGetTypeID()
}/* debug [functions.gen.go/function]: SecTrustedApplicationGetTypeID */

// Sets the data of a given trusted app instance.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Sets the data of a given trusted app instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustedApplicationSetData(_:_:)
func SecTrustedApplicationSetData(appRef SecTrustedApplicationRef, data DataRef) unsafe.Pointer {
	return _SecTrustedApplicationSetData(appRef, data)
}/* debug [functions.gen.go/function]: SecTrustedApplicationSetData */

// Evaluates trust for the specified certificate and policies.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.3.
// Evaluates trust for the specified certificate and policies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluate(_:_:)
func SecTrustEvaluate(trust SecTrustRef, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluate(trust, result)
}/* debug [functions.gen.go/function]: SecTrustEvaluate */

// Evaluates a trust object asynchronously on the specified dispatch queue.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.7.
// Evaluates a trust object asynchronously on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsync(_:_:_:)
func SecTrustEvaluateAsync(trust SecTrustRef, queue unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluateAsync(trust, queue, result)
}/* debug [functions.gen.go/function]: SecTrustEvaluateAsync */

// Evaluates a trust object asynchronously on the specified dispatch queue.
//
// Added in macOS 10.15.
// Evaluates a trust object asynchronously on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsyncWithError(_:_:_:)
func SecTrustEvaluateAsyncWithError(trust SecTrustRef, queue unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluateAsyncWithError(trust, queue, result)
}/* debug [functions.gen.go/function]: SecTrustEvaluateAsyncWithError */

// Evaluates trust for the specified certificate and policies.
//
// Added in macOS 10.14.
// Evaluates trust for the specified certificate and policies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluateWithError(_:_:)
func SecTrustEvaluateWithError(trust SecTrustRef, error_ unsafe.Pointer) bool {
	return _SecTrustEvaluateWithError(trust, error_)
}/* debug [functions.gen.go/function]: SecTrustEvaluateWithError */

// Returns a specific certificate from the certificate chain used to evaluate trust.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
// Returns a specific certificate from the certificate chain used to evaluate trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetCertificateAtIndex(_:_:)
func SecTrustGetCertificateAtIndex(trust SecTrustRef, ix Index) SecCertificateRef {
	return _SecTrustGetCertificateAtIndex(trust, ix)
}/* debug [functions.gen.go/function]: SecTrustGetCertificateAtIndex */

// Returns the number of certificates in an evaluated certificate chain.
//
// Added in macOS 10.7.
// Returns the number of certificates in an evaluated certificate chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetCertificateCount(_:)
func SecTrustGetCertificateCount(trust SecTrustRef) Index {
	return _SecTrustGetCertificateCount(trust)
}/* debug [functions.gen.go/function]: SecTrustGetCertificateCount */

// Indicates whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// Added in macOS 10.9.
// Indicates whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetNetworkFetchAllowed(_:_:)
func SecTrustGetNetworkFetchAllowed(trust SecTrustRef, allowFetch unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetNetworkFetchAllowed(trust, allowFetch)
}/* debug [functions.gen.go/function]: SecTrustGetNetworkFetchAllowed */

// Returns the result code from the most recent trust evaluation.
//
// Added in macOS 10.7.
// Returns the result code from the most recent trust evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetTrustResult(_:_:)
func SecTrustGetTrustResult(trust SecTrustRef, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetTrustResult(trust, result)
}/* debug [functions.gen.go/function]: SecTrustGetTrustResult */

// Returns the unique identifier of the opaque type to which a trust object belongs.
//
// Added in macOS 10.3.
// Returns the unique identifier of the opaque type to which a trust object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetTypeID()
func SecTrustGetTypeID() TypeID {
	return _SecTrustGetTypeID()
}/* debug [functions.gen.go/function]: SecTrustGetTypeID */

// Gets the absolute time against which the certificates in a trust management object are verified.
//
// Added in macOS 10.6.
// Gets the absolute time against which the certificates in a trust management object are verified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetVerifyTime(_:)
func SecTrustGetVerifyTime(trust SecTrustRef) AbsoluteTime {
	return _SecTrustGetVerifyTime(trust)
}/* debug [functions.gen.go/function]: SecTrustGetVerifyTime */

// Sets the anchor certificates used when evaluating a trust management object.
//
// Added in macOS 10.3.
// Sets the anchor certificates used when evaluating a trust management object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificates(_:_:)
func SecTrustSetAnchorCertificates(trust SecTrustRef, anchorCertificates ArrayRef) unsafe.Pointer {
	return _SecTrustSetAnchorCertificates(trust, anchorCertificates)
}/* debug [functions.gen.go/function]: SecTrustSetAnchorCertificates */

// Reenables trusting built-in anchor certificates.
//
// Added in macOS 10.6.
// Reenables trusting built-in anchor certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificatesOnly(_:_:)
func SecTrustSetAnchorCertificatesOnly(trust SecTrustRef, anchorCertificatesOnly unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetAnchorCertificatesOnly(trust, anchorCertificatesOnly)
}/* debug [functions.gen.go/function]: SecTrustSetAnchorCertificatesOnly */

// Sets a list of exceptions that should be ignored when the certificate is evaluated.
//
// Added in macOS 10.9.
// Sets a list of exceptions that should be ignored when the certificate is evaluated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetExceptions(_:_:)
func SecTrustSetExceptions(trust SecTrustRef, exceptions DataRef) bool {
	return _SecTrustSetExceptions(trust, exceptions)
}/* debug [functions.gen.go/function]: SecTrustSetExceptions */

// Sets the keychains searched for intermediate certificates when evaluating a trust management object.
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.3.
// Sets the keychains searched for intermediate certificates when evaluating a trust management object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetKeychains(_:_:)
func SecTrustSetKeychains(trust SecTrustRef, keychainOrArray TypeRef) unsafe.Pointer {
	return _SecTrustSetKeychains(trust, keychainOrArray)
}/* debug [functions.gen.go/function]: SecTrustSetKeychains */

// Specifies whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// Added in macOS 10.9.
// Specifies whether a trust evaluation is permitted to fetch missing intermediate certificates from the network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetNetworkFetchAllowed(_:_:)
func SecTrustSetNetworkFetchAllowed(trust SecTrustRef, allowFetch unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetNetworkFetchAllowed(trust, allowFetch)
}/* debug [functions.gen.go/function]: SecTrustSetNetworkFetchAllowed */

// Attaches Online Certificate Status Protocol (OSCP) response data to a trust object.
//
// Added in macOS 10.9.
// Attaches Online Certificate Status Protocol (OSCP) response data to a trust object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetOCSPResponse(_:_:)
func SecTrustSetOCSPResponse(trust SecTrustRef, responseData TypeRef) unsafe.Pointer {
	return _SecTrustSetOCSPResponse(trust, responseData)
}/* debug [functions.gen.go/function]: SecTrustSetOCSPResponse */

// Sets option flags for customizing evaluation of a trust object.
//
// Added in macOS 10.7.
// Sets option flags for customizing evaluation of a trust object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetOptions(_:_:)
func SecTrustSetOptions(trustRef SecTrustRef, options SecTrustOptionFlags) unsafe.Pointer {
	return _SecTrustSetOptions(trustRef, options)
}/* debug [functions.gen.go/function]: SecTrustSetOptions */

// Sets the policies to use in an evaluation.
//
// Added in macOS 10.3.
// Sets the policies to use in an evaluation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetPolicies(_:_:)
func SecTrustSetPolicies(trust SecTrustRef, policies TypeRef) unsafe.Pointer {
	return _SecTrustSetPolicies(trust, policies)
}/* debug [functions.gen.go/function]: SecTrustSetPolicies */

// Attaches signed certificate timestamp data to a trust object.
//
// Added in macOS 10.14.2.
// Attaches signed certificate timestamp data to a trust object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetSignedCertificateTimestamps(_:_:)
func SecTrustSetSignedCertificateTimestamps(trust SecTrustRef, sctArray ArrayRef) unsafe.Pointer {
	return _SecTrustSetSignedCertificateTimestamps(trust, sctArray)
}/* debug [functions.gen.go/function]: SecTrustSetSignedCertificateTimestamps */

// Obtains an array of all certificates that have trust settings in a specific trust settings domain.
//
// Added in macOS 10.0.
// Obtains an array of all certificates that have trust settings in a specific trust settings domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyCertificates(_:_:)
func SecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain, certArray unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSettingsCopyCertificates(domain, certArray)
}/* debug [functions.gen.go/function]: SecTrustSettingsCopyCertificates */

// Obtains the date and time at which a certificate’s trust settings were last modified.
//
// Added in macOS 10.0.
// Obtains the date and time at which a certificate’s trust settings were last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyModificationDate(_:_:_:)
func SecTrustSettingsCopyModificationDate(certRef SecCertificateRef, domain SecTrustSettingsDomain, modificationDate unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSettingsCopyModificationDate(certRef, domain, modificationDate)
}/* debug [functions.gen.go/function]: SecTrustSettingsCopyModificationDate */

// Obtains the trust settings for a certificate.
//
// Added in macOS 10.0.
// Obtains the trust settings for a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsCopyTrustSettings(_:_:_:)
func SecTrustSettingsCopyTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettings unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSettingsCopyTrustSettings(certRef, domain, trustSettings)
}/* debug [functions.gen.go/function]: SecTrustSettingsCopyTrustSettings */

// Obtains an external, portable representation of the specified domain’s trust settings.
//
// Added in macOS 10.0.
// Obtains an external, portable representation of the specified domain’s trust settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsCreateExternalRepresentation(_:_:)
func SecTrustSettingsCreateExternalRepresentation(domain SecTrustSettingsDomain, trustSettings unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSettingsCreateExternalRepresentation(domain, trustSettings)
}/* debug [functions.gen.go/function]: SecTrustSettingsCreateExternalRepresentation */

// Imports trust settings into a trust domain.
//
// Added in macOS 10.0.
// Imports trust settings into a trust domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsImportExternalRepresentation(_:_:)
func SecTrustSettingsImportExternalRepresentation(domain SecTrustSettingsDomain, trustSettings DataRef) unsafe.Pointer {
	return _SecTrustSettingsImportExternalRepresentation(domain, trustSettings)
}/* debug [functions.gen.go/function]: SecTrustSettingsImportExternalRepresentation */

// Deletes the trust settings for a certificate.
//
// Added in macOS 10.0.
// Deletes the trust settings for a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsRemoveTrustSettings(_:_:)
func SecTrustSettingsRemoveTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain) unsafe.Pointer {
	return _SecTrustSettingsRemoveTrustSettings(certRef, domain)
}/* debug [functions.gen.go/function]: SecTrustSettingsRemoveTrustSettings */

// Specifies trust settings for a certificate.
//
// Added in macOS 10.0.
// Specifies trust settings for a certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSettingsSetTrustSettings(_:_:_:)
func SecTrustSettingsSetTrustSettings(certRef SecCertificateRef, domain SecTrustSettingsDomain, trustSettingsDictOrArray TypeRef) unsafe.Pointer {
	return _SecTrustSettingsSetTrustSettings(certRef, domain, trustSettingsDictOrArray)
}/* debug [functions.gen.go/function]: SecTrustSettingsSetTrustSettings */

// Sets the date and time against which the certificates in a trust management object are verified.
//
// Added in macOS 10.3.
// Sets the date and time against which the certificates in a trust management object are verified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetVerifyDate(_:_:)
func SecTrustSetVerifyDate(trust SecTrustRef, verifyDate DateRef) unsafe.Pointer {
	return _SecTrustSetVerifyDate(trust, verifyDate)
}/* debug [functions.gen.go/function]: SecTrustSetVerifyDate */

// Creates a verify transform object.
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
// Creates a verify transform object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecVerifyTransformCreate(_:_:_:)
func SecVerifyTransformCreate(key SecKeyRef, signature DataRef, error_ unsafe.Pointer) SecTransformRef {
	return _SecVerifyTransformCreate(key, signature, error_)
}/* debug [functions.gen.go/function]: SecVerifyTransformCreate */

// Creates a security session.
//
// Added in macOS 10.0.
// Creates a security session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionCreate(_:_:)
func SessionCreate(flags SessionCreationFlags, attributes SessionAttributeBits) unsafe.Pointer {
	return _SessionCreate(flags, attributes)
}/* debug [functions.gen.go/function]: SessionCreate */

// Obtains information about a security session.
//
// Added in macOS 10.0.
// Obtains information about a security session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionGetInfo(_:_:_:)
func SessionGetInfo(session SecuritySessionId, sessionId unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _SessionGetInfo(session, sessionId, attributes)
}/* debug [functions.gen.go/function]: SessionGetInfo */

// Adds a DER-encoded distinguished name to a list of acceptable names to be specified in requests for client certificates.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.4.
// Adds a DER-encoded distinguished name to a list of acceptable names to be specified in requests for client certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAddDistinguishedName(_:_:_:)
func SSLAddDistinguishedName(context SSLContextRef, derDN unsafe.Pointer, derDNLen uintptr) unsafe.Pointer {
	return _SSLAddDistinguishedName(context, derDN, derDNLen)
}/* debug [functions.gen.go/function]: SSLAddDistinguishedName */

// Terminates the current SSL session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Terminates the current SSL session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClose(_:)
func SSLClose(context SSLContextRef) unsafe.Pointer {
	return _SSLClose(context)
}/* debug [functions.gen.go/function]: SSLClose */

// Returns the Core Foundation type ID for context objects.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Returns the Core Foundation type ID for context objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLContextGetTypeID()
func SSLContextGetTypeID() TypeID {
	return _SSLContextGetTypeID()
}/* debug [functions.gen.go/function]: SSLContextGetTypeID */

// Gets the list of supported application layer protocols.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
// Gets the list of supported application layer protocols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyALPNProtocols(_:_:)
func SSLCopyALPNProtocols(context SSLContextRef, protocols unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyALPNProtocols(context, protocols)
}/* debug [functions.gen.go/function]: SSLCopyALPNProtocols */

// Retrieves the current list of certification authorities.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
// Retrieves the current list of certification authorities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyCertificateAuthorities(_:_:)
func SSLCopyCertificateAuthorities(context SSLContextRef, certificates unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyCertificateAuthorities(context, certificates)
}/* debug [functions.gen.go/function]: SSLCopyCertificateAuthorities */

// Retrieves the distinguished names of acceptable certification authorities.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
// Retrieves the distinguished names of acceptable certification authorities.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyDistinguishedNames(_:_:)
func SSLCopyDistinguishedNames(context SSLContextRef, names unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyDistinguishedNames(context, names)
}/* debug [functions.gen.go/function]: SSLCopyDistinguishedNames */

// Retrieves a trust management object for the certificate used by a session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Retrieves a trust management object for the certificate used by a session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyPeerTrust(_:_:)
func SSLCopyPeerTrust(context SSLContextRef, trust unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyPeerTrust(context, trust)
}/* debug [functions.gen.go/function]: SSLCopyPeerTrust */

// Determines the buffer size needed for the peer domain name.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
// Determines the buffer size needed for the peer domain name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyRequestedPeerName(_:_:_:)
func SSLCopyRequestedPeerName(context SSLContextRef, peerName unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyRequestedPeerName(context, peerName, peerNameLen)
}/* debug [functions.gen.go/function]: SSLCopyRequestedPeerName */

// Obtains the hostname specified by the client in the ServerName extension (SNI). Server only.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
// Obtains the hostname specified by the client in the ServerName extension (SNI). Server only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyRequestedPeerNameLength(_:_:)
func SSLCopyRequestedPeerNameLength(ctx SSLContextRef, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyRequestedPeerNameLength(ctx, peerNameLen)
}/* debug [functions.gen.go/function]: SSLCopyRequestedPeerNameLength */

// Allocates and returns a new context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Allocates and returns a new context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCreateContext(_:_:_:)
func SSLCreateContext(alloc AllocatorRef, protocolSide SSLProtocolSide, connectionType SSLConnectionType) SSLContextRef {
	return _SSLCreateContext(alloc, protocolSide, connectionType)
}/* debug [functions.gen.go/function]: SSLCreateContext */

// Determines how much data is available to be read.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines how much data is available to be read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetBufferedReadSize(_:_:)
func SSLGetBufferedReadSize(context SSLContextRef, bufferSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetBufferedReadSize(context, bufferSize)
}/* debug [functions.gen.go/function]: SSLGetBufferedReadSize */

// Retrieves the exchange status of the client certificate.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.3.
// Retrieves the exchange status of the client certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetClientCertificateState(_:_:)
func SSLGetClientCertificateState(context SSLContextRef, clientState unsafe.Pointer) unsafe.Pointer {
	return _SSLGetClientCertificateState(context, clientState)
}/* debug [functions.gen.go/function]: SSLGetClientCertificateState */

// Retrieves an I/O connection—such as a socket or endpoint—for a specific session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves an I/O connection—such as a socket or endpoint—for a specific session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetConnection(_:_:)
func SSLGetConnection(context SSLContextRef, connection unsafe.Pointer) unsafe.Pointer {
	return _SSLGetConnection(context, connection)
}/* debug [functions.gen.go/function]: SSLGetConnection */

// Provides the largest packet that the OS guarantees it can send without fragmentation.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Provides the largest packet that the OS guarantees it can send without fragmentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetDatagramWriteSize(_:_:)
func SSLGetDatagramWriteSize(dtlsContext SSLContextRef, bufSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetDatagramWriteSize(dtlsContext, bufSize)
}/* debug [functions.gen.go/function]: SSLGetDatagramWriteSize */

// Retrieves the Diffie-Hellman parameters for a given context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves the Diffie-Hellman parameters for a given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetDiffieHellmanParams(_:_:_:)
func SSLGetDiffieHellmanParams(context SSLContextRef, dhParams unsafe.Pointer, dhParamsLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetDiffieHellmanParams(context, dhParams, dhParamsLen)
}/* debug [functions.gen.go/function]: SSLGetDiffieHellmanParams */

// Determines which SSL cipher suites are currently enabled.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines which SSL cipher suites are currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetEnabledCiphers(_:_:_:)
func SSLGetEnabledCiphers(context SSLContextRef, ciphers unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetEnabledCiphers(context, ciphers, numCiphers)
}/* debug [functions.gen.go/function]: SSLGetEnabledCiphers */

// Obtains the maximum datagram record size allowed by the application for a given context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Obtains the maximum datagram record size allowed by the application for a given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetMaxDatagramRecordSize(_:_:)
func SSLGetMaxDatagramRecordSize(dtlsContext SSLContextRef, maxSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetMaxDatagramRecordSize(dtlsContext, maxSize)
}/* debug [functions.gen.go/function]: SSLGetMaxDatagramRecordSize */

// Retrieves the cipher suite negotiated for this session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves the cipher suite negotiated for this session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNegotiatedCipher(_:_:)
func SSLGetNegotiatedCipher(context SSLContextRef, cipherSuite unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNegotiatedCipher(context, cipherSuite)
}/* debug [functions.gen.go/function]: SSLGetNegotiatedCipher */

// Obtains the negotiated protocol version of the active session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Obtains the negotiated protocol version of the active session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNegotiatedProtocolVersion(_:_:)
func SSLGetNegotiatedProtocolVersion(context SSLContextRef, protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNegotiatedProtocolVersion(context, protocol_)
}/* debug [functions.gen.go/function]: SSLGetNegotiatedProtocolVersion */

// Determines the number of cipher suites currently enabled.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines the number of cipher suites currently enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNumberEnabledCiphers(_:_:)
func SSLGetNumberEnabledCiphers(context SSLContextRef, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNumberEnabledCiphers(context, numCiphers)
}/* debug [functions.gen.go/function]: SSLGetNumberEnabledCiphers */

// Determines the number of cipher suites supported.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines the number of cipher suites supported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNumberSupportedCiphers(_:_:)
func SSLGetNumberSupportedCiphers(context SSLContextRef, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNumberSupportedCiphers(context, numCiphers)
}/* debug [functions.gen.go/function]: SSLGetNumberSupportedCiphers */

// Retrieves the peer domain name specified previously.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves the peer domain name specified previously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerDomainName(_:_:_:)
func SSLGetPeerDomainName(context SSLContextRef, peerName unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerDomainName(context, peerName, peerNameLen)
}/* debug [functions.gen.go/function]: SSLGetPeerDomainName */

// Determines the length of a previously set peer domain name.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines the length of a previously set peer domain name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerDomainNameLength(_:_:)
func SSLGetPeerDomainNameLength(context SSLContextRef, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerDomainNameLength(context, peerNameLen)
}/* debug [functions.gen.go/function]: SSLGetPeerDomainNameLength */

// Retrieves the current peer ID data.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves the current peer ID data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerID(_:_:_:)
func SSLGetPeerID(context SSLContextRef, peerID unsafe.Pointer, peerIDLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerID(context, peerID, peerIDLen)
}/* debug [functions.gen.go/function]: SSLGetPeerID */

// Gets the maximum protocol version allowed by the application for a given SSL context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Gets the maximum protocol version allowed by the application for a given SSL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersionMax(_:_:)
func SSLGetProtocolVersionMax(context SSLContextRef, maxVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersionMax(context, maxVersion)
}/* debug [functions.gen.go/function]: SSLGetProtocolVersionMax */

// Gets the minimum protocol version allowed by the application for a given SSL context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Gets the minimum protocol version allowed by the application for a given SSL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersionMin(_:_:)
func SSLGetProtocolVersionMin(context SSLContextRef, minVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersionMin(context, minVersion)
}/* debug [functions.gen.go/function]: SSLGetProtocolVersionMin */

// Indicates the current setting of Secure Sockets Layer (SSL) session options.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Indicates the current setting of Secure Sockets Layer (SSL) session options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSessionOption(_:_:_:)
func SSLGetSessionOption(context SSLContextRef, option SSLSessionOption, value unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSessionOption(context, option, value)
}/* debug [functions.gen.go/function]: SSLGetSessionOption */

// Retrieves the state of an SSL session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Retrieves the state of an SSL session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSessionState(_:_:)
func SSLGetSessionState(context SSLContextRef, state unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSessionState(context, state)
}/* debug [functions.gen.go/function]: SSLGetSessionState */

// Determines the values of the supported cipher suites.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Determines the values of the supported cipher suites.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSupportedCiphers(_:_:_:)
func SSLGetSupportedCiphers(context SSLContextRef, ciphers unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSupportedCiphers(context, ciphers, numCiphers)
}/* debug [functions.gen.go/function]: SSLGetSupportedCiphers */

// Performs the SSL handshake.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Performs the SSL handshake.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLHandshake(_:)
func SSLHandshake(context SSLContextRef) unsafe.Pointer {
	return _SSLHandshake(context)
}/* debug [functions.gen.go/function]: SSLHandshake */

// Performs a normal application-level read operation.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Performs a normal application-level read operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLRead(_:_:_:_:)
func SSLRead(context SSLContextRef, data unsafe.Pointer, dataLength uintptr, processed unsafe.Pointer) unsafe.Pointer {
	return _SSLRead(context, data, dataLength, processed)
}/* debug [functions.gen.go/function]: SSLRead */

// Requests renegotiation of the SSL handshake. Server only.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.12.
// Requests renegotiation of the SSL handshake. Server only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLReHandshake(_:)
func SSLReHandshake(context SSLContextRef) unsafe.Pointer {
	return _SSLReHandshake(context)
}/* debug [functions.gen.go/function]: SSLReHandshake */

// Sets the list of supported applicaiton layer protocols.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
// Sets the list of supported applicaiton layer protocols.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetALPNProtocols(_:_:)
func SSLSetALPNProtocols(context SSLContextRef, protocols ArrayRef) unsafe.Pointer {
	return _SSLSetALPNProtocols(context, protocols)
}/* debug [functions.gen.go/function]: SSLSetALPNProtocols */

// Specifies this connection’s certificate or certificates.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies this connection’s certificate or certificates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetCertificate(_:_:)
func SSLSetCertificate(context SSLContextRef, certRefs ArrayRef) unsafe.Pointer {
	return _SSLSetCertificate(context, certRefs)
}/* debug [functions.gen.go/function]: SSLSetCertificate */

// Adds one or more certificates to a server’s list of certification authorities (CAs) acceptable for client authentication.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
// Adds one or more certificates to a server’s list of certification authorities (CAs) acceptable for client authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetCertificateAuthorities(_:_:_:)
func SSLSetCertificateAuthorities(context SSLContextRef, certificateOrArray TypeRef, replaceExisting unsafe.Pointer) unsafe.Pointer {
	return _SSLSetCertificateAuthorities(context, certificateOrArray, replaceExisting)
}/* debug [functions.gen.go/function]: SSLSetCertificateAuthorities */

// Specifies the requirements for client-side authentication.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies the requirements for client-side authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetClientSideAuthenticate(_:_:)
func SSLSetClientSideAuthenticate(context SSLContextRef, auth SSLAuthenticate) unsafe.Pointer {
	return _SSLSetClientSideAuthenticate(context, auth)
}/* debug [functions.gen.go/function]: SSLSetClientSideAuthenticate */

// Specifies an I/O connection for a specific session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies an I/O connection for a specific session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetConnection(_:_:)
func SSLSetConnection(context SSLContextRef, connection SSLConnectionRef) unsafe.Pointer {
	return _SSLSetConnection(context, connection)
}/* debug [functions.gen.go/function]: SSLSetConnection */

// Sets the cookie value used in the Datagram Transport Layer Security (DTLS) hello message.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Sets the cookie value used in the Datagram Transport Layer Security (DTLS) hello message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetDatagramHelloCookie(_:_:_:)
func SSLSetDatagramHelloCookie(dtlsContext SSLContextRef, cookie unsafe.Pointer, cookieLen uintptr) unsafe.Pointer {
	return _SSLSetDatagramHelloCookie(dtlsContext, cookie, cookieLen)
}/* debug [functions.gen.go/function]: SSLSetDatagramHelloCookie */

// Specifies Diffie-Hellman parameters for a given context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies Diffie-Hellman parameters for a given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetDiffieHellmanParams(_:_:_:)
func SSLSetDiffieHellmanParams(context SSLContextRef, dhParams unsafe.Pointer, dhParamsLen uintptr) unsafe.Pointer {
	return _SSLSetDiffieHellmanParams(context, dhParams, dhParamsLen)
}/* debug [functions.gen.go/function]: SSLSetDiffieHellmanParams */

// Specifies a restricted set of SSL cipher suites to be enabled by the current SSL session context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies a restricted set of SSL cipher suites to be enabled by the current SSL session context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetEnabledCiphers(_:_:_:)
func SSLSetEnabledCiphers(context SSLContextRef, ciphers unsafe.Pointer, numCiphers uintptr) unsafe.Pointer {
	return _SSLSetEnabledCiphers(context, ciphers, numCiphers)
}/* debug [functions.gen.go/function]: SSLSetEnabledCiphers */

// Specifies the encryption certificates used for this connection.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.2.
// Specifies the encryption certificates used for this connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetEncryptionCertificate(_:_:)
func SSLSetEncryptionCertificate(context SSLContextRef, certRefs ArrayRef) unsafe.Pointer {
	return _SSLSetEncryptionCertificate(context, certRefs)
}/* debug [functions.gen.go/function]: SSLSetEncryptionCertificate */

// Sets the status of a session context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
// Sets the status of a session context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetError(_:_:)
func SSLSetError(context SSLContextRef, status unsafe.Pointer) unsafe.Pointer {
	return _SSLSetError(context, status)
}/* debug [functions.gen.go/function]: SSLSetError */

// Specifies callback functions that perform the network I/O operations.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies callback functions that perform the network I/O operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetIOFuncs(_:_:_:)
func SSLSetIOFuncs(context SSLContextRef, readFunc SSLReadFunc, writeFunc SSLWriteFunc) unsafe.Pointer {
	return _SSLSetIOFuncs(context, readFunc, writeFunc)
}/* debug [functions.gen.go/function]: SSLSetIOFuncs */

// Sets the maximum datagram record size allowed by the application for a given context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Sets the maximum datagram record size allowed by the application for a given context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetMaxDatagramRecordSize(_:_:)
func SSLSetMaxDatagramRecordSize(dtlsContext SSLContextRef, maxSize uintptr) unsafe.Pointer {
	return _SSLSetMaxDatagramRecordSize(dtlsContext, maxSize)
}/* debug [functions.gen.go/function]: SSLSetMaxDatagramRecordSize */

// Sets the OCSP response for the given SSL session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
// Sets the OCSP response for the given SSL session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetOCSPResponse(_:_:)
func SSLSetOCSPResponse(context SSLContextRef, response DataRef) unsafe.Pointer {
	return _SSLSetOCSPResponse(context, response)
}/* debug [functions.gen.go/function]: SSLSetOCSPResponse */

// Specifies the fully qualified domain name of the peer.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies the fully qualified domain name of the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetPeerDomainName(_:_:_:)
func SSLSetPeerDomainName(context SSLContextRef, peerName unsafe.Pointer, peerNameLen uintptr) unsafe.Pointer {
	return _SSLSetPeerDomainName(context, peerName, peerNameLen)
}/* debug [functions.gen.go/function]: SSLSetPeerDomainName */

// Specifies data that is sufficient to uniquely identify the peer of the current session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Specifies data that is sufficient to uniquely identify the peer of the current session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetPeerID(_:_:_:)
func SSLSetPeerID(context SSLContextRef, peerID unsafe.Pointer, peerIDLen uintptr) unsafe.Pointer {
	return _SSLSetPeerID(context, peerID, peerIDLen)
}/* debug [functions.gen.go/function]: SSLSetPeerID */

// Sets the maximum protocol version allowed by the application for a given SSL context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Sets the maximum protocol version allowed by the application for a given SSL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersionMax(_:_:)
func SSLSetProtocolVersionMax(context SSLContextRef, maxVersion SSLProtocol) unsafe.Pointer {
	return _SSLSetProtocolVersionMax(context, maxVersion)
}/* debug [functions.gen.go/function]: SSLSetProtocolVersionMax */

// Sets the minimum protocol version allowed by the application for a given SSL context.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
// Sets the minimum protocol version allowed by the application for a given SSL context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersionMin(_:_:)
func SSLSetProtocolVersionMin(context SSLContextRef, minVersion SSLProtocol) unsafe.Pointer {
	return _SSLSetProtocolVersionMin(context, minVersion)
}/* debug [functions.gen.go/function]: SSLSetProtocolVersionMin */

// Sets a predefined configuration for the Secure Sockets Layer (SSL) session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.12.
// Sets a predefined configuration for the Secure Sockets Layer (SSL) session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionConfig(_:_:)
func SSLSetSessionConfig(context SSLContextRef, config StringRef) unsafe.Pointer {
	return _SSLSetSessionConfig(context, config)
}/* debug [functions.gen.go/function]: SSLSetSessionConfig */

// Specifies options for a specific session.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
// Specifies options for a specific session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionOption(_:_:_:)
func SSLSetSessionOption(context SSLContextRef, option SSLSessionOption, value unsafe.Pointer) unsafe.Pointer {
	return _SSLSetSessionOption(context, option, value)
}/* debug [functions.gen.go/function]: SSLSetSessionOption */

// Enables or disables session ticket resumption.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
// Enables or disables session ticket resumption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionTicketsEnabled(_:_:)
func SSLSetSessionTicketsEnabled(context SSLContextRef, enabled unsafe.Pointer) unsafe.Pointer {
	return _SSLSetSessionTicketsEnabled(context, enabled)
}/* debug [functions.gen.go/function]: SSLSetSessionTicketsEnabled */

// Performs a typical application-level write operation.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Performs a typical application-level write operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLWrite(_:_:_:_:)
func SSLWrite(context SSLContextRef, data unsafe.Pointer, dataLength uintptr, processed unsafe.Pointer) unsafe.Pointer {
	return _SSLWrite(context, data, dataLength, processed)
}/* debug [functions.gen.go/function]: SSLWrite */




