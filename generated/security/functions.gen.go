// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Security Functions (526 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_SecIdentitySearchCreate func(unsafe.Pointer) unsafe.Pointer
	_SecIdentitySearchGetTypeID func() unsafe.Pointer
	_SecIdentitySearchCopyNext func(unsafe.Pointer) unsafe.Pointer
	_SecTrustGetTPHandle func(unsafe.Pointer) unsafe.Pointer
	_SecTrustGetCssmResult func(unsafe.Pointer) unsafe.Pointer
	_SecTrustSetParameters func(unsafe.Pointer) unsafe.Pointer
	_SecTrustGetCssmResultCode func(unsafe.Pointer) unsafe.Pointer
	_SecTrustGetResult func(unsafe.Pointer) unsafe.Pointer
	_AuthorizationExecuteWithPrivileges func(unsafe.Pointer) unsafe.Pointer
	_AuthorizationPluginCreate func(unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyPrivilegedReference func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyRights func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCopyRightsAsync func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationCreateFromExternalForm func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationFree func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationFreeItemSet func(unsafe.Pointer) unsafe.Pointer
	_AuthorizationMakeExternalForm func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightRemove func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_AuthorizationRightSet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyAllCerts func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyDetachedContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopyEncapsulatedContentType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerCert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerEmailAddress func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerSigningTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerStatus func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestamp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestampCertificates func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCopySignerTimestampWithPolicy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderCreate func(unsafe.Pointer) unsafe.Pointer
	_CMSDecoderFinalizeMessage func(unsafe.Pointer) unsafe.Pointer
	_CMSDecoderGetNumSigners func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderGetTypeID func() unsafe.Pointer
	_CMSDecoderIsContentEncrypted func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderSetDetachedContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderSetSearchKeychain func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSDecoderUpdateMessage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncodeContent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderAddRecipients func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderAddSignedAttributes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderAddSigners func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderAddSupportingCerts func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopyEncapsulatedContentType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopyEncodedContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopyRecipients func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySignerTimestamp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySignerTimestampWithPolicy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySigners func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCopySupportingCerts func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderCreate func(unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetCertificateChainMode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetHasDetachedContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderGetTypeID func() unsafe.Pointer
	_CMSEncoderSetCertificateChainMode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderSetEncapsulatedContentType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderSetEncapsulatedContentTypeOID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderSetHasDetachedContent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderSetSignerAlgorithm func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CMSEncoderUpdateContent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_AC_AuthCompute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_AC_PassThrough func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertAbortCache func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertAbortQuery func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertCache func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertCreateTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertDescribeFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetAllFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetAllTemplateFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetFirstCachedFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetFirstFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetKeyInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetNextCachedFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGetNextFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGroupFromVerifiedBundle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertGroupToSignedBundle func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertSign func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertVerify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CertVerifyWithKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlAbortCache func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlAbortQuery func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlAddCert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlCache func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlCreateTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlDescribeFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetAllCachedRecordFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetAllFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetFirstCachedFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetFirstFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetNextCachedFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlGetNextFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlRemoveCert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlSetFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlSign func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlVerify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_CrlVerifyWithKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_FreeFieldValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_FreeFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_IsCertInCachedCrl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_IsCertInCrl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CL_PassThrough func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_ChangeLoginAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_ChangeLoginOwner func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateAsymmetricContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateDeriveKeyContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateDigestContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateKeyGenContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateMacContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreatePassThroughContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateRandomGenContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateSignatureContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_CreateSymmetricContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_GetLoginAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_GetLoginOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_GetOperationalStatistics func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_Login func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_Logout func(unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_ObtainPrivateKeyFromPublicKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_CSP_PassThrough func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ChangeKeyAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ChangeKeyOwner func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_Authenticate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_ChangeDbAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_ChangeDbOwner func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_CreateRelation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataAbortQuery func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataDelete func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataGetFirst func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataGetFromUniqueRecordId func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataGetNext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DataModify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DbClose func(unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DbCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DbDelete func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DbOpen func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_DestroyRelation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_FreeNameList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_FreeUniqueRecord func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_GetDbAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_GetDbNameFromHandle func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_GetDbNames func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_GetDbOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DL_PassThrough func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptDataFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptDataInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptDataInitP func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptDataP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DecryptDataUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DeleteContext func(unsafe.Pointer) unsafe.Pointer
	_CSSM_DeleteContextAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DeriveKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DigestData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DigestDataClone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DigestDataFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_DigestDataInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_DigestDataUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptDataFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptDataInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptDataInitP func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptDataP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_EncryptDataUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_FreeContext func(unsafe.Pointer) unsafe.Pointer
	_CSSM_FreeKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateAlgorithmParams func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateKeyP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateKeyPair func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateKeyPairP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateMac func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateMacFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateMacInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateMacUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GenerateRandom func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetAPIMemoryFunctions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetContextAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetKeyAcl func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetKeyOwner func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetModuleGUIDFromHandle func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetPrivilege func(unsafe.Pointer) unsafe.Pointer
	_CSSM_GetSubserviceUIDFromHandle func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_GetTimeValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_Init func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_Introduce func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ListAttachedModuleManagers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ModuleAttach func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ModuleDetach func(unsafe.Pointer) unsafe.Pointer
	_CSSM_ModuleLoad func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_ModuleUnload func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_QueryKeySizeInBits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_QuerySize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_RetrieveCounter func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_RetrieveUniqueId func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_SetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_SetPrivilege func(unsafe.Pointer) unsafe.Pointer
	_CSSM_SignData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_SignDataFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_SignDataInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_SignDataUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_ApplyCrlToDb func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertCreateTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertGetAllTemplateFields func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertGroupConstruct func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertGroupPrune func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertGroupToTupleGroup func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertGroupVerify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertReclaimAbort func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertReclaimKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertRemoveFromCrlTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertRevoke func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CertSign func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_ConfirmCredResult func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CrlCreateTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CrlSign func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_CrlVerify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_FormRequest func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_FormSubmit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_PassThrough func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_ReceiveConfirmation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_RetrieveCredResult func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_SubmitCredRequest func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_TP_TupleGroupToCertGroup func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_Terminate func() unsafe.Pointer
	_CSSM_Unintroduce func(unsafe.Pointer) unsafe.Pointer
	_CSSM_UnwrapKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_UnwrapKeyP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_UpdateContextAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyDataFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyDataInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyDataUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyDevice func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyMac func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyMacFinal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyMacInit func(unsafe.Pointer) unsafe.Pointer
	_CSSM_VerifyMacUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_WrapKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CSSM_WrapKeyP func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MDS_Initialize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_MDS_Install func(unsafe.Pointer) unsafe.Pointer
	_MDS_Terminate func(unsafe.Pointer) unsafe.Pointer
	_MDS_Uninstall func(unsafe.Pointer) unsafe.Pointer
	_SSLAddDistinguishedName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLClose func(unsafe.Pointer) unsafe.Pointer
	_SSLContextGetTypeID func() unsafe.Pointer
	_SSLCopyALPNProtocols func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyCertificateAuthorities func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyDistinguishedNames func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyPeerCertificates func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyPeerTrust func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyRequestedPeerName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyRequestedPeerNameLength func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCopyTrustedRoots func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLCreateContext func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLDisposeContext func(unsafe.Pointer) unsafe.Pointer
	_SSLGetAllowsAnyRoot func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetAllowsExpiredCerts func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetAllowsExpiredRoots func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetBufferedReadSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetClientCertificateState func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetConnection func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetDatagramWriteSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetDiffieHellmanParams func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetEnableCertVerify func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetEnabledCiphers func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetMaxDatagramRecordSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetNegotiatedCipher func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetNegotiatedProtocolVersion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetNumberEnabledCiphers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetNumberSupportedCiphers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerDomainName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerDomainNameLength func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetPeerID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersionEnabled func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersionMax func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetProtocolVersionMin func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetRsaBlinding func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetSessionOption func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetSessionState func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLGetSupportedCiphers func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLHandshake func(unsafe.Pointer) unsafe.Pointer
	_SSLNewContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLReHandshake func(unsafe.Pointer) unsafe.Pointer
	_SSLRead func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetALPNProtocols func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetAllowsAnyRoot func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetAllowsExpiredCerts func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetAllowsExpiredRoots func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetCertificate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetCertificateAuthorities func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetClientSideAuthenticate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetConnection func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetDatagramHelloCookie func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetDiffieHellmanParams func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetEnableCertVerify func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetEnabledCiphers func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetEncryptionCertificate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetIOFuncs func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetMaxDatagramRecordSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetOCSPResponse func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetPeerDomainName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetPeerID func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetProtocolVersion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetProtocolVersionEnabled func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetProtocolVersionMax func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetProtocolVersionMin func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetRsaBlinding func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetSessionConfig func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetSessionOption func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetSessionTicketsEnabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLSetTrustedRoots func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SSLWrite func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLCreateWithSimpleContents func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLSetContents func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecACLUpdateAuthorizations func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessControlCreateWithFlags func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessCopyACLList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessCopyMatchingACLList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessCopySelectedACLList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessGetOwnerAndACL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAccessGetTypeID func() unsafe.Pointer
	_SecAddSharedWebCredential func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1AllocCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1AllocCopyItem func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1AllocItem func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1CoderCreate func(unsafe.Pointer) unsafe.Pointer
	_SecAsn1CoderRelease func(unsafe.Pointer) unsafe.Pointer
	_SecAsn1Decode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1DecodeData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1EncodeItem func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1Malloc func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecAsn1OidCompare func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyData func(unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyNotValidAfterDate func(unsafe.Pointer) unsafe.Pointer
	_SecCertificateCopyNotValidBeforeDate func(unsafe.Pointer) unsafe.Pointer
	_SecCertificateCreateWithData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCheckValidity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCheckValidityWithErrors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyDesignatedRequirement func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyGuestWithAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyHost func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopySelf func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopySigningInformation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCopyStaticCode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeCreateWithXPCMessage func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeGetTypeID func() unsafe.Pointer
	_SecCodeMapMemory func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCodeValidateFileResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCopyErrorMessageString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecCreateSharedWebCredentialPassword func() unsafe.Pointer
	_SecDecodeTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecDecryptTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecDecryptTransformGetTypeID func() unsafe.Pointer
	_SecDigestTransformCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecDigestTransformGetTypeID func() unsafe.Pointer
	_SecEncodeTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecEncryptTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecEncryptTransformGetTypeID func() unsafe.Pointer
	_SecGroupTransformGetTypeID func() unsafe.Pointer
	_SecHostCreateGuest func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecHostRemoveGuest func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecHostSelectGuest func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecHostSelectedGuest func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecHostSetGuestStatus func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecHostSetHostingPort func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecIdentityCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecItemAdd func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecItemCopyMatching func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecItemDelete func(unsafe.Pointer) unsafe.Pointer
	_SecItemUpdate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCopyExternalRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCopyPublicKey func(unsafe.Pointer) unsafe.Pointer
	_SecKeyCreateDecryptedData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCreateEncryptedData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCreateRandomKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyCreateWithData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyIsAlgorithmSupported func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeyUnwrapSymmetric func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecKeychainItemCopyAccess func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecPolicyCreateBasicX509 func() unsafe.Pointer
	_SecPolicyCreateSSL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecPolicyCreateWithProperties func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRandomCopyBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequestSharedWebCredential func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCopyData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCopyString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementCreateWithStringAndErrors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecRequirementGetTypeID func() unsafe.Pointer
	_SecSignTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCheckValidity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCheckValidityWithErrors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCreateWithPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeCreateWithPathAndAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecStaticCodeGetTypeID func() unsafe.Pointer
	_SecTaskCopySigningIdentifier func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTaskCopyValueForEntitlement func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTaskCopyValuesForEntitlements func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTaskCreateFromSelf func(unsafe.Pointer) unsafe.Pointer
	_SecTaskCreateWithAuditToken func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTaskGetCodeSignStatus func(unsafe.Pointer) uint32
	_SecTaskGetTypeID func() unsafe.Pointer
	_SecTranformCustomGetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformConnectTransforms func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformCopyExternalRepresentation func(unsafe.Pointer) unsafe.Pointer
	_SecTransformCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformCreateFromExternalRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformCreateGroupTransform func() unsafe.Pointer
	_SecTransformCreateReadTransformWithReadStream func(unsafe.Pointer) unsafe.Pointer
	_SecTransformCustomGetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformCustomSetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformExecute func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformExecuteAsync func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformFindByName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformGetAttribute func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformGetTypeID func() unsafe.Pointer
	_SecTransformNoData func() unsafe.Pointer
	_SecTransformPushbackAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformRegister func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetAttributeAction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetDataAction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTransformSetTransformAction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustCopyCertificateChain func(unsafe.Pointer) unsafe.Pointer
	_SecTrustCopyKey func(unsafe.Pointer) unsafe.Pointer
	_SecTrustCreateWithCertificates func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluateWithError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustGetTrustResult func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetAnchorCertificates func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetVerifyDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecVerifyTransformCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadCopyCreationDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadCopyName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadCopyTicketLocation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadCopyURLs func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadCreateWithTicket func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadFinished func(unsafe.Pointer) unsafe.Pointer
	_SecureDownloadGetDownloadSize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecureDownloadRelease func(unsafe.Pointer) unsafe.Pointer
	_SecureDownloadUpdateWithData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SessionCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SessionGetInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_cssmAlgToOid func(unsafe.Pointer) unsafe.Pointer
	_cssmOidToAlg func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_cssmPerror func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_certificate_copy_ref func(unsafe.Pointer) unsafe.Pointer
	_sec_certificate_create func(unsafe.Pointer) unsafe.Pointer
	_sec_identity_access_certificates func(unsafe.Pointer) unsafe.Pointer
	_sec_identity_copy_certificates_ref func(unsafe.Pointer) unsafe.Pointer
	_sec_identity_copy_ref func(unsafe.Pointer) unsafe.Pointer
	_sec_identity_create func(unsafe.Pointer) unsafe.Pointer
	_sec_identity_create_with_certificates func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_access_distinguished_names func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_access_ocsp_response func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_access_peer_certificate_chain func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_access_pre_shared_keys func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_access_supported_signature_algorithms func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_challenge_parameters_are_equal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_copy_negotiated_protocol func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_copy_peer_public_key func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_copy_server_name func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_create_secret func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_create_secret_with_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_early_data_accepted func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_negotiated_ciphersuite func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_negotiated_protocol func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_negotiated_protocol_version func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_negotiated_tls_ciphersuite func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_negotiated_tls_protocol_version func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_get_server_name func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_metadata_peers_are_equal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_add_pre_shared_key func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_add_tls_application_protocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_add_tls_ciphersuite func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_add_tls_ciphersuite_group func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_append_tls_ciphersuite func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_append_tls_ciphersuite_group func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_are_equal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_get_default_max_dtls_protocol_version func() unsafe.Pointer
	_sec_protocol_options_get_default_max_tls_protocol_version func() unsafe.Pointer
	_sec_protocol_options_get_default_min_dtls_protocol_version func() unsafe.Pointer
	_sec_protocol_options_get_default_min_tls_protocol_version func() unsafe.Pointer
	_sec_protocol_options_get_enable_encrypted_client_hello func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_get_quic_use_legacy_codepoint func(unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_challenge_block func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_enable_encrypted_client_hello func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_key_update_block func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_local_identity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_max_tls_protocol_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_min_tls_protocol_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_peer_authentication_optional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_peer_authentication_required func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_pre_shared_key_selection_block func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_quic_use_legacy_codepoint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_diffie_hellman_parameters func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_false_start_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_is_fallback_attempt func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_max_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_min_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_ocsp_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_pre_shared_key_identity_hint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_renegotiation_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_resumption_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_sct_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_server_name func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_tls_tickets_enabled func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_protocol_options_set_verify_block func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_sec_release func(unsafe.Pointer) unsafe.Pointer
	_sec_retain func(unsafe.Pointer) unsafe.Pointer
	_sec_trust_copy_ref func(unsafe.Pointer) unsafe.Pointer
	_sec_trust_create func(unsafe.Pointer) unsafe.Pointer
	_SecTrustEvaluateAsync func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_SecTrustSetKeychains func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_SecIdentitySearchCreate, lib, "SecIdentitySearchCreate")
	tryRegister(&_SecIdentitySearchGetTypeID, lib, "SecIdentitySearchGetTypeID")
	tryRegister(&_SecIdentitySearchCopyNext, lib, "SecIdentitySearchCopyNext")
	tryRegister(&_SecTrustGetTPHandle, lib, "SecTrustGetTPHandle")
	tryRegister(&_SecTrustGetCssmResult, lib, "SecTrustGetCssmResult")
	tryRegister(&_SecTrustSetParameters, lib, "SecTrustSetParameters")
	tryRegister(&_SecTrustGetCssmResultCode, lib, "SecTrustGetCssmResultCode")
	tryRegister(&_SecTrustGetResult, lib, "SecTrustGetResult")
	tryRegister(&_AuthorizationExecuteWithPrivileges, lib, "AuthorizationExecuteWithPrivileges")
	tryRegister(&_AuthorizationPluginCreate, lib, "AuthorizationPluginCreate")
	tryRegister(&_AuthorizationCopyInfo, lib, "AuthorizationCopyInfo")
	tryRegister(&_AuthorizationCopyPrivilegedReference, lib, "AuthorizationCopyPrivilegedReference")
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
	tryRegister(&_CMSEncode, lib, "CMSEncode")
	tryRegister(&_CMSEncodeContent, lib, "CMSEncodeContent")
	tryRegister(&_CMSEncoderAddRecipients, lib, "CMSEncoderAddRecipients")
	tryRegister(&_CMSEncoderAddSignedAttributes, lib, "CMSEncoderAddSignedAttributes")
	tryRegister(&_CMSEncoderAddSigners, lib, "CMSEncoderAddSigners")
	tryRegister(&_CMSEncoderAddSupportingCerts, lib, "CMSEncoderAddSupportingCerts")
	tryRegister(&_CMSEncoderCopyEncapsulatedContentType, lib, "CMSEncoderCopyEncapsulatedContentType")
	tryRegister(&_CMSEncoderCopyEncodedContent, lib, "CMSEncoderCopyEncodedContent")
	tryRegister(&_CMSEncoderCopyRecipients, lib, "CMSEncoderCopyRecipients")
	tryRegister(&_CMSEncoderCopySignerTimestamp, lib, "CMSEncoderCopySignerTimestamp")
	tryRegister(&_CMSEncoderCopySignerTimestampWithPolicy, lib, "CMSEncoderCopySignerTimestampWithPolicy")
	tryRegister(&_CMSEncoderCopySigners, lib, "CMSEncoderCopySigners")
	tryRegister(&_CMSEncoderCopySupportingCerts, lib, "CMSEncoderCopySupportingCerts")
	tryRegister(&_CMSEncoderCreate, lib, "CMSEncoderCreate")
	tryRegister(&_CMSEncoderGetCertificateChainMode, lib, "CMSEncoderGetCertificateChainMode")
	tryRegister(&_CMSEncoderGetHasDetachedContent, lib, "CMSEncoderGetHasDetachedContent")
	tryRegister(&_CMSEncoderGetTypeID, lib, "CMSEncoderGetTypeID")
	tryRegister(&_CMSEncoderSetCertificateChainMode, lib, "CMSEncoderSetCertificateChainMode")
	tryRegister(&_CMSEncoderSetEncapsulatedContentType, lib, "CMSEncoderSetEncapsulatedContentType")
	tryRegister(&_CMSEncoderSetEncapsulatedContentTypeOID, lib, "CMSEncoderSetEncapsulatedContentTypeOID")
	tryRegister(&_CMSEncoderSetHasDetachedContent, lib, "CMSEncoderSetHasDetachedContent")
	tryRegister(&_CMSEncoderSetSignerAlgorithm, lib, "CMSEncoderSetSignerAlgorithm")
	tryRegister(&_CMSEncoderUpdateContent, lib, "CMSEncoderUpdateContent")
	tryRegister(&_CSSM_AC_AuthCompute, lib, "CSSM_AC_AuthCompute")
	tryRegister(&_CSSM_AC_PassThrough, lib, "CSSM_AC_PassThrough")
	tryRegister(&_CSSM_CL_CertAbortCache, lib, "CSSM_CL_CertAbortCache")
	tryRegister(&_CSSM_CL_CertAbortQuery, lib, "CSSM_CL_CertAbortQuery")
	tryRegister(&_CSSM_CL_CertCache, lib, "CSSM_CL_CertCache")
	tryRegister(&_CSSM_CL_CertCreateTemplate, lib, "CSSM_CL_CertCreateTemplate")
	tryRegister(&_CSSM_CL_CertDescribeFormat, lib, "CSSM_CL_CertDescribeFormat")
	tryRegister(&_CSSM_CL_CertGetAllFields, lib, "CSSM_CL_CertGetAllFields")
	tryRegister(&_CSSM_CL_CertGetAllTemplateFields, lib, "CSSM_CL_CertGetAllTemplateFields")
	tryRegister(&_CSSM_CL_CertGetFirstCachedFieldValue, lib, "CSSM_CL_CertGetFirstCachedFieldValue")
	tryRegister(&_CSSM_CL_CertGetFirstFieldValue, lib, "CSSM_CL_CertGetFirstFieldValue")
	tryRegister(&_CSSM_CL_CertGetKeyInfo, lib, "CSSM_CL_CertGetKeyInfo")
	tryRegister(&_CSSM_CL_CertGetNextCachedFieldValue, lib, "CSSM_CL_CertGetNextCachedFieldValue")
	tryRegister(&_CSSM_CL_CertGetNextFieldValue, lib, "CSSM_CL_CertGetNextFieldValue")
	tryRegister(&_CSSM_CL_CertGroupFromVerifiedBundle, lib, "CSSM_CL_CertGroupFromVerifiedBundle")
	tryRegister(&_CSSM_CL_CertGroupToSignedBundle, lib, "CSSM_CL_CertGroupToSignedBundle")
	tryRegister(&_CSSM_CL_CertSign, lib, "CSSM_CL_CertSign")
	tryRegister(&_CSSM_CL_CertVerify, lib, "CSSM_CL_CertVerify")
	tryRegister(&_CSSM_CL_CertVerifyWithKey, lib, "CSSM_CL_CertVerifyWithKey")
	tryRegister(&_CSSM_CL_CrlAbortCache, lib, "CSSM_CL_CrlAbortCache")
	tryRegister(&_CSSM_CL_CrlAbortQuery, lib, "CSSM_CL_CrlAbortQuery")
	tryRegister(&_CSSM_CL_CrlAddCert, lib, "CSSM_CL_CrlAddCert")
	tryRegister(&_CSSM_CL_CrlCache, lib, "CSSM_CL_CrlCache")
	tryRegister(&_CSSM_CL_CrlCreateTemplate, lib, "CSSM_CL_CrlCreateTemplate")
	tryRegister(&_CSSM_CL_CrlDescribeFormat, lib, "CSSM_CL_CrlDescribeFormat")
	tryRegister(&_CSSM_CL_CrlGetAllCachedRecordFields, lib, "CSSM_CL_CrlGetAllCachedRecordFields")
	tryRegister(&_CSSM_CL_CrlGetAllFields, lib, "CSSM_CL_CrlGetAllFields")
	tryRegister(&_CSSM_CL_CrlGetFirstCachedFieldValue, lib, "CSSM_CL_CrlGetFirstCachedFieldValue")
	tryRegister(&_CSSM_CL_CrlGetFirstFieldValue, lib, "CSSM_CL_CrlGetFirstFieldValue")
	tryRegister(&_CSSM_CL_CrlGetNextCachedFieldValue, lib, "CSSM_CL_CrlGetNextCachedFieldValue")
	tryRegister(&_CSSM_CL_CrlGetNextFieldValue, lib, "CSSM_CL_CrlGetNextFieldValue")
	tryRegister(&_CSSM_CL_CrlRemoveCert, lib, "CSSM_CL_CrlRemoveCert")
	tryRegister(&_CSSM_CL_CrlSetFields, lib, "CSSM_CL_CrlSetFields")
	tryRegister(&_CSSM_CL_CrlSign, lib, "CSSM_CL_CrlSign")
	tryRegister(&_CSSM_CL_CrlVerify, lib, "CSSM_CL_CrlVerify")
	tryRegister(&_CSSM_CL_CrlVerifyWithKey, lib, "CSSM_CL_CrlVerifyWithKey")
	tryRegister(&_CSSM_CL_FreeFieldValue, lib, "CSSM_CL_FreeFieldValue")
	tryRegister(&_CSSM_CL_FreeFields, lib, "CSSM_CL_FreeFields")
	tryRegister(&_CSSM_CL_IsCertInCachedCrl, lib, "CSSM_CL_IsCertInCachedCrl")
	tryRegister(&_CSSM_CL_IsCertInCrl, lib, "CSSM_CL_IsCertInCrl")
	tryRegister(&_CSSM_CL_PassThrough, lib, "CSSM_CL_PassThrough")
	tryRegister(&_CSSM_CSP_ChangeLoginAcl, lib, "CSSM_CSP_ChangeLoginAcl")
	tryRegister(&_CSSM_CSP_ChangeLoginOwner, lib, "CSSM_CSP_ChangeLoginOwner")
	tryRegister(&_CSSM_CSP_CreateAsymmetricContext, lib, "CSSM_CSP_CreateAsymmetricContext")
	tryRegister(&_CSSM_CSP_CreateDeriveKeyContext, lib, "CSSM_CSP_CreateDeriveKeyContext")
	tryRegister(&_CSSM_CSP_CreateDigestContext, lib, "CSSM_CSP_CreateDigestContext")
	tryRegister(&_CSSM_CSP_CreateKeyGenContext, lib, "CSSM_CSP_CreateKeyGenContext")
	tryRegister(&_CSSM_CSP_CreateMacContext, lib, "CSSM_CSP_CreateMacContext")
	tryRegister(&_CSSM_CSP_CreatePassThroughContext, lib, "CSSM_CSP_CreatePassThroughContext")
	tryRegister(&_CSSM_CSP_CreateRandomGenContext, lib, "CSSM_CSP_CreateRandomGenContext")
	tryRegister(&_CSSM_CSP_CreateSignatureContext, lib, "CSSM_CSP_CreateSignatureContext")
	tryRegister(&_CSSM_CSP_CreateSymmetricContext, lib, "CSSM_CSP_CreateSymmetricContext")
	tryRegister(&_CSSM_CSP_GetLoginAcl, lib, "CSSM_CSP_GetLoginAcl")
	tryRegister(&_CSSM_CSP_GetLoginOwner, lib, "CSSM_CSP_GetLoginOwner")
	tryRegister(&_CSSM_CSP_GetOperationalStatistics, lib, "CSSM_CSP_GetOperationalStatistics")
	tryRegister(&_CSSM_CSP_Login, lib, "CSSM_CSP_Login")
	tryRegister(&_CSSM_CSP_Logout, lib, "CSSM_CSP_Logout")
	tryRegister(&_CSSM_CSP_ObtainPrivateKeyFromPublicKey, lib, "CSSM_CSP_ObtainPrivateKeyFromPublicKey")
	tryRegister(&_CSSM_CSP_PassThrough, lib, "CSSM_CSP_PassThrough")
	tryRegister(&_CSSM_ChangeKeyAcl, lib, "CSSM_ChangeKeyAcl")
	tryRegister(&_CSSM_ChangeKeyOwner, lib, "CSSM_ChangeKeyOwner")
	tryRegister(&_CSSM_DL_Authenticate, lib, "CSSM_DL_Authenticate")
	tryRegister(&_CSSM_DL_ChangeDbAcl, lib, "CSSM_DL_ChangeDbAcl")
	tryRegister(&_CSSM_DL_ChangeDbOwner, lib, "CSSM_DL_ChangeDbOwner")
	tryRegister(&_CSSM_DL_CreateRelation, lib, "CSSM_DL_CreateRelation")
	tryRegister(&_CSSM_DL_DataAbortQuery, lib, "CSSM_DL_DataAbortQuery")
	tryRegister(&_CSSM_DL_DataDelete, lib, "CSSM_DL_DataDelete")
	tryRegister(&_CSSM_DL_DataGetFirst, lib, "CSSM_DL_DataGetFirst")
	tryRegister(&_CSSM_DL_DataGetFromUniqueRecordId, lib, "CSSM_DL_DataGetFromUniqueRecordId")
	tryRegister(&_CSSM_DL_DataGetNext, lib, "CSSM_DL_DataGetNext")
	tryRegister(&_CSSM_DL_DataInsert, lib, "CSSM_DL_DataInsert")
	tryRegister(&_CSSM_DL_DataModify, lib, "CSSM_DL_DataModify")
	tryRegister(&_CSSM_DL_DbClose, lib, "CSSM_DL_DbClose")
	tryRegister(&_CSSM_DL_DbCreate, lib, "CSSM_DL_DbCreate")
	tryRegister(&_CSSM_DL_DbDelete, lib, "CSSM_DL_DbDelete")
	tryRegister(&_CSSM_DL_DbOpen, lib, "CSSM_DL_DbOpen")
	tryRegister(&_CSSM_DL_DestroyRelation, lib, "CSSM_DL_DestroyRelation")
	tryRegister(&_CSSM_DL_FreeNameList, lib, "CSSM_DL_FreeNameList")
	tryRegister(&_CSSM_DL_FreeUniqueRecord, lib, "CSSM_DL_FreeUniqueRecord")
	tryRegister(&_CSSM_DL_GetDbAcl, lib, "CSSM_DL_GetDbAcl")
	tryRegister(&_CSSM_DL_GetDbNameFromHandle, lib, "CSSM_DL_GetDbNameFromHandle")
	tryRegister(&_CSSM_DL_GetDbNames, lib, "CSSM_DL_GetDbNames")
	tryRegister(&_CSSM_DL_GetDbOwner, lib, "CSSM_DL_GetDbOwner")
	tryRegister(&_CSSM_DL_PassThrough, lib, "CSSM_DL_PassThrough")
	tryRegister(&_CSSM_DecryptData, lib, "CSSM_DecryptData")
	tryRegister(&_CSSM_DecryptDataFinal, lib, "CSSM_DecryptDataFinal")
	tryRegister(&_CSSM_DecryptDataInit, lib, "CSSM_DecryptDataInit")
	tryRegister(&_CSSM_DecryptDataInitP, lib, "CSSM_DecryptDataInitP")
	tryRegister(&_CSSM_DecryptDataP, lib, "CSSM_DecryptDataP")
	tryRegister(&_CSSM_DecryptDataUpdate, lib, "CSSM_DecryptDataUpdate")
	tryRegister(&_CSSM_DeleteContext, lib, "CSSM_DeleteContext")
	tryRegister(&_CSSM_DeleteContextAttributes, lib, "CSSM_DeleteContextAttributes")
	tryRegister(&_CSSM_DeriveKey, lib, "CSSM_DeriveKey")
	tryRegister(&_CSSM_DigestData, lib, "CSSM_DigestData")
	tryRegister(&_CSSM_DigestDataClone, lib, "CSSM_DigestDataClone")
	tryRegister(&_CSSM_DigestDataFinal, lib, "CSSM_DigestDataFinal")
	tryRegister(&_CSSM_DigestDataInit, lib, "CSSM_DigestDataInit")
	tryRegister(&_CSSM_DigestDataUpdate, lib, "CSSM_DigestDataUpdate")
	tryRegister(&_CSSM_EncryptData, lib, "CSSM_EncryptData")
	tryRegister(&_CSSM_EncryptDataFinal, lib, "CSSM_EncryptDataFinal")
	tryRegister(&_CSSM_EncryptDataInit, lib, "CSSM_EncryptDataInit")
	tryRegister(&_CSSM_EncryptDataInitP, lib, "CSSM_EncryptDataInitP")
	tryRegister(&_CSSM_EncryptDataP, lib, "CSSM_EncryptDataP")
	tryRegister(&_CSSM_EncryptDataUpdate, lib, "CSSM_EncryptDataUpdate")
	tryRegister(&_CSSM_FreeContext, lib, "CSSM_FreeContext")
	tryRegister(&_CSSM_FreeKey, lib, "CSSM_FreeKey")
	tryRegister(&_CSSM_GenerateAlgorithmParams, lib, "CSSM_GenerateAlgorithmParams")
	tryRegister(&_CSSM_GenerateKey, lib, "CSSM_GenerateKey")
	tryRegister(&_CSSM_GenerateKeyP, lib, "CSSM_GenerateKeyP")
	tryRegister(&_CSSM_GenerateKeyPair, lib, "CSSM_GenerateKeyPair")
	tryRegister(&_CSSM_GenerateKeyPairP, lib, "CSSM_GenerateKeyPairP")
	tryRegister(&_CSSM_GenerateMac, lib, "CSSM_GenerateMac")
	tryRegister(&_CSSM_GenerateMacFinal, lib, "CSSM_GenerateMacFinal")
	tryRegister(&_CSSM_GenerateMacInit, lib, "CSSM_GenerateMacInit")
	tryRegister(&_CSSM_GenerateMacUpdate, lib, "CSSM_GenerateMacUpdate")
	tryRegister(&_CSSM_GenerateRandom, lib, "CSSM_GenerateRandom")
	tryRegister(&_CSSM_GetAPIMemoryFunctions, lib, "CSSM_GetAPIMemoryFunctions")
	tryRegister(&_CSSM_GetContext, lib, "CSSM_GetContext")
	tryRegister(&_CSSM_GetContextAttribute, lib, "CSSM_GetContextAttribute")
	tryRegister(&_CSSM_GetKeyAcl, lib, "CSSM_GetKeyAcl")
	tryRegister(&_CSSM_GetKeyOwner, lib, "CSSM_GetKeyOwner")
	tryRegister(&_CSSM_GetModuleGUIDFromHandle, lib, "CSSM_GetModuleGUIDFromHandle")
	tryRegister(&_CSSM_GetPrivilege, lib, "CSSM_GetPrivilege")
	tryRegister(&_CSSM_GetSubserviceUIDFromHandle, lib, "CSSM_GetSubserviceUIDFromHandle")
	tryRegister(&_CSSM_GetTimeValue, lib, "CSSM_GetTimeValue")
	tryRegister(&_CSSM_Init, lib, "CSSM_Init")
	tryRegister(&_CSSM_Introduce, lib, "CSSM_Introduce")
	tryRegister(&_CSSM_ListAttachedModuleManagers, lib, "CSSM_ListAttachedModuleManagers")
	tryRegister(&_CSSM_ModuleAttach, lib, "CSSM_ModuleAttach")
	tryRegister(&_CSSM_ModuleDetach, lib, "CSSM_ModuleDetach")
	tryRegister(&_CSSM_ModuleLoad, lib, "CSSM_ModuleLoad")
	tryRegister(&_CSSM_ModuleUnload, lib, "CSSM_ModuleUnload")
	tryRegister(&_CSSM_QueryKeySizeInBits, lib, "CSSM_QueryKeySizeInBits")
	tryRegister(&_CSSM_QuerySize, lib, "CSSM_QuerySize")
	tryRegister(&_CSSM_RetrieveCounter, lib, "CSSM_RetrieveCounter")
	tryRegister(&_CSSM_RetrieveUniqueId, lib, "CSSM_RetrieveUniqueId")
	tryRegister(&_CSSM_SetContext, lib, "CSSM_SetContext")
	tryRegister(&_CSSM_SetPrivilege, lib, "CSSM_SetPrivilege")
	tryRegister(&_CSSM_SignData, lib, "CSSM_SignData")
	tryRegister(&_CSSM_SignDataFinal, lib, "CSSM_SignDataFinal")
	tryRegister(&_CSSM_SignDataInit, lib, "CSSM_SignDataInit")
	tryRegister(&_CSSM_SignDataUpdate, lib, "CSSM_SignDataUpdate")
	tryRegister(&_CSSM_TP_ApplyCrlToDb, lib, "CSSM_TP_ApplyCrlToDb")
	tryRegister(&_CSSM_TP_CertCreateTemplate, lib, "CSSM_TP_CertCreateTemplate")
	tryRegister(&_CSSM_TP_CertGetAllTemplateFields, lib, "CSSM_TP_CertGetAllTemplateFields")
	tryRegister(&_CSSM_TP_CertGroupConstruct, lib, "CSSM_TP_CertGroupConstruct")
	tryRegister(&_CSSM_TP_CertGroupPrune, lib, "CSSM_TP_CertGroupPrune")
	tryRegister(&_CSSM_TP_CertGroupToTupleGroup, lib, "CSSM_TP_CertGroupToTupleGroup")
	tryRegister(&_CSSM_TP_CertGroupVerify, lib, "CSSM_TP_CertGroupVerify")
	tryRegister(&_CSSM_TP_CertReclaimAbort, lib, "CSSM_TP_CertReclaimAbort")
	tryRegister(&_CSSM_TP_CertReclaimKey, lib, "CSSM_TP_CertReclaimKey")
	tryRegister(&_CSSM_TP_CertRemoveFromCrlTemplate, lib, "CSSM_TP_CertRemoveFromCrlTemplate")
	tryRegister(&_CSSM_TP_CertRevoke, lib, "CSSM_TP_CertRevoke")
	tryRegister(&_CSSM_TP_CertSign, lib, "CSSM_TP_CertSign")
	tryRegister(&_CSSM_TP_ConfirmCredResult, lib, "CSSM_TP_ConfirmCredResult")
	tryRegister(&_CSSM_TP_CrlCreateTemplate, lib, "CSSM_TP_CrlCreateTemplate")
	tryRegister(&_CSSM_TP_CrlSign, lib, "CSSM_TP_CrlSign")
	tryRegister(&_CSSM_TP_CrlVerify, lib, "CSSM_TP_CrlVerify")
	tryRegister(&_CSSM_TP_FormRequest, lib, "CSSM_TP_FormRequest")
	tryRegister(&_CSSM_TP_FormSubmit, lib, "CSSM_TP_FormSubmit")
	tryRegister(&_CSSM_TP_PassThrough, lib, "CSSM_TP_PassThrough")
	tryRegister(&_CSSM_TP_ReceiveConfirmation, lib, "CSSM_TP_ReceiveConfirmation")
	tryRegister(&_CSSM_TP_RetrieveCredResult, lib, "CSSM_TP_RetrieveCredResult")
	tryRegister(&_CSSM_TP_SubmitCredRequest, lib, "CSSM_TP_SubmitCredRequest")
	tryRegister(&_CSSM_TP_TupleGroupToCertGroup, lib, "CSSM_TP_TupleGroupToCertGroup")
	tryRegister(&_CSSM_Terminate, lib, "CSSM_Terminate")
	tryRegister(&_CSSM_Unintroduce, lib, "CSSM_Unintroduce")
	tryRegister(&_CSSM_UnwrapKey, lib, "CSSM_UnwrapKey")
	tryRegister(&_CSSM_UnwrapKeyP, lib, "CSSM_UnwrapKeyP")
	tryRegister(&_CSSM_UpdateContextAttributes, lib, "CSSM_UpdateContextAttributes")
	tryRegister(&_CSSM_VerifyData, lib, "CSSM_VerifyData")
	tryRegister(&_CSSM_VerifyDataFinal, lib, "CSSM_VerifyDataFinal")
	tryRegister(&_CSSM_VerifyDataInit, lib, "CSSM_VerifyDataInit")
	tryRegister(&_CSSM_VerifyDataUpdate, lib, "CSSM_VerifyDataUpdate")
	tryRegister(&_CSSM_VerifyDevice, lib, "CSSM_VerifyDevice")
	tryRegister(&_CSSM_VerifyMac, lib, "CSSM_VerifyMac")
	tryRegister(&_CSSM_VerifyMacFinal, lib, "CSSM_VerifyMacFinal")
	tryRegister(&_CSSM_VerifyMacInit, lib, "CSSM_VerifyMacInit")
	tryRegister(&_CSSM_VerifyMacUpdate, lib, "CSSM_VerifyMacUpdate")
	tryRegister(&_CSSM_WrapKey, lib, "CSSM_WrapKey")
	tryRegister(&_CSSM_WrapKeyP, lib, "CSSM_WrapKeyP")
	tryRegister(&_MDS_Initialize, lib, "MDS_Initialize")
	tryRegister(&_MDS_Install, lib, "MDS_Install")
	tryRegister(&_MDS_Terminate, lib, "MDS_Terminate")
	tryRegister(&_MDS_Uninstall, lib, "MDS_Uninstall")
	tryRegister(&_SSLAddDistinguishedName, lib, "SSLAddDistinguishedName")
	tryRegister(&_SSLClose, lib, "SSLClose")
	tryRegister(&_SSLContextGetTypeID, lib, "SSLContextGetTypeID")
	tryRegister(&_SSLCopyALPNProtocols, lib, "SSLCopyALPNProtocols")
	tryRegister(&_SSLCopyCertificateAuthorities, lib, "SSLCopyCertificateAuthorities")
	tryRegister(&_SSLCopyDistinguishedNames, lib, "SSLCopyDistinguishedNames")
	tryRegister(&_SSLCopyPeerCertificates, lib, "SSLCopyPeerCertificates")
	tryRegister(&_SSLCopyPeerTrust, lib, "SSLCopyPeerTrust")
	tryRegister(&_SSLCopyRequestedPeerName, lib, "SSLCopyRequestedPeerName")
	tryRegister(&_SSLCopyRequestedPeerNameLength, lib, "SSLCopyRequestedPeerNameLength")
	tryRegister(&_SSLCopyTrustedRoots, lib, "SSLCopyTrustedRoots")
	tryRegister(&_SSLCreateContext, lib, "SSLCreateContext")
	tryRegister(&_SSLDisposeContext, lib, "SSLDisposeContext")
	tryRegister(&_SSLGetAllowsAnyRoot, lib, "SSLGetAllowsAnyRoot")
	tryRegister(&_SSLGetAllowsExpiredCerts, lib, "SSLGetAllowsExpiredCerts")
	tryRegister(&_SSLGetAllowsExpiredRoots, lib, "SSLGetAllowsExpiredRoots")
	tryRegister(&_SSLGetBufferedReadSize, lib, "SSLGetBufferedReadSize")
	tryRegister(&_SSLGetClientCertificateState, lib, "SSLGetClientCertificateState")
	tryRegister(&_SSLGetConnection, lib, "SSLGetConnection")
	tryRegister(&_SSLGetDatagramWriteSize, lib, "SSLGetDatagramWriteSize")
	tryRegister(&_SSLGetDiffieHellmanParams, lib, "SSLGetDiffieHellmanParams")
	tryRegister(&_SSLGetEnableCertVerify, lib, "SSLGetEnableCertVerify")
	tryRegister(&_SSLGetEnabledCiphers, lib, "SSLGetEnabledCiphers")
	tryRegister(&_SSLGetMaxDatagramRecordSize, lib, "SSLGetMaxDatagramRecordSize")
	tryRegister(&_SSLGetNegotiatedCipher, lib, "SSLGetNegotiatedCipher")
	tryRegister(&_SSLGetNegotiatedProtocolVersion, lib, "SSLGetNegotiatedProtocolVersion")
	tryRegister(&_SSLGetNumberEnabledCiphers, lib, "SSLGetNumberEnabledCiphers")
	tryRegister(&_SSLGetNumberSupportedCiphers, lib, "SSLGetNumberSupportedCiphers")
	tryRegister(&_SSLGetPeerDomainName, lib, "SSLGetPeerDomainName")
	tryRegister(&_SSLGetPeerDomainNameLength, lib, "SSLGetPeerDomainNameLength")
	tryRegister(&_SSLGetPeerID, lib, "SSLGetPeerID")
	tryRegister(&_SSLGetProtocolVersion, lib, "SSLGetProtocolVersion")
	tryRegister(&_SSLGetProtocolVersionEnabled, lib, "SSLGetProtocolVersionEnabled")
	tryRegister(&_SSLGetProtocolVersionMax, lib, "SSLGetProtocolVersionMax")
	tryRegister(&_SSLGetProtocolVersionMin, lib, "SSLGetProtocolVersionMin")
	tryRegister(&_SSLGetRsaBlinding, lib, "SSLGetRsaBlinding")
	tryRegister(&_SSLGetSessionOption, lib, "SSLGetSessionOption")
	tryRegister(&_SSLGetSessionState, lib, "SSLGetSessionState")
	tryRegister(&_SSLGetSupportedCiphers, lib, "SSLGetSupportedCiphers")
	tryRegister(&_SSLHandshake, lib, "SSLHandshake")
	tryRegister(&_SSLNewContext, lib, "SSLNewContext")
	tryRegister(&_SSLReHandshake, lib, "SSLReHandshake")
	tryRegister(&_SSLRead, lib, "SSLRead")
	tryRegister(&_SSLSetALPNProtocols, lib, "SSLSetALPNProtocols")
	tryRegister(&_SSLSetAllowsAnyRoot, lib, "SSLSetAllowsAnyRoot")
	tryRegister(&_SSLSetAllowsExpiredCerts, lib, "SSLSetAllowsExpiredCerts")
	tryRegister(&_SSLSetAllowsExpiredRoots, lib, "SSLSetAllowsExpiredRoots")
	tryRegister(&_SSLSetCertificate, lib, "SSLSetCertificate")
	tryRegister(&_SSLSetCertificateAuthorities, lib, "SSLSetCertificateAuthorities")
	tryRegister(&_SSLSetClientSideAuthenticate, lib, "SSLSetClientSideAuthenticate")
	tryRegister(&_SSLSetConnection, lib, "SSLSetConnection")
	tryRegister(&_SSLSetDatagramHelloCookie, lib, "SSLSetDatagramHelloCookie")
	tryRegister(&_SSLSetDiffieHellmanParams, lib, "SSLSetDiffieHellmanParams")
	tryRegister(&_SSLSetEnableCertVerify, lib, "SSLSetEnableCertVerify")
	tryRegister(&_SSLSetEnabledCiphers, lib, "SSLSetEnabledCiphers")
	tryRegister(&_SSLSetEncryptionCertificate, lib, "SSLSetEncryptionCertificate")
	tryRegister(&_SSLSetError, lib, "SSLSetError")
	tryRegister(&_SSLSetIOFuncs, lib, "SSLSetIOFuncs")
	tryRegister(&_SSLSetMaxDatagramRecordSize, lib, "SSLSetMaxDatagramRecordSize")
	tryRegister(&_SSLSetOCSPResponse, lib, "SSLSetOCSPResponse")
	tryRegister(&_SSLSetPeerDomainName, lib, "SSLSetPeerDomainName")
	tryRegister(&_SSLSetPeerID, lib, "SSLSetPeerID")
	tryRegister(&_SSLSetProtocolVersion, lib, "SSLSetProtocolVersion")
	tryRegister(&_SSLSetProtocolVersionEnabled, lib, "SSLSetProtocolVersionEnabled")
	tryRegister(&_SSLSetProtocolVersionMax, lib, "SSLSetProtocolVersionMax")
	tryRegister(&_SSLSetProtocolVersionMin, lib, "SSLSetProtocolVersionMin")
	tryRegister(&_SSLSetRsaBlinding, lib, "SSLSetRsaBlinding")
	tryRegister(&_SSLSetSessionConfig, lib, "SSLSetSessionConfig")
	tryRegister(&_SSLSetSessionOption, lib, "SSLSetSessionOption")
	tryRegister(&_SSLSetSessionTicketsEnabled, lib, "SSLSetSessionTicketsEnabled")
	tryRegister(&_SSLSetTrustedRoots, lib, "SSLSetTrustedRoots")
	tryRegister(&_SSLWrite, lib, "SSLWrite")
	tryRegister(&_SecACLCreateWithSimpleContents, lib, "SecACLCreateWithSimpleContents")
	tryRegister(&_SecACLSetContents, lib, "SecACLSetContents")
	tryRegister(&_SecACLUpdateAuthorizations, lib, "SecACLUpdateAuthorizations")
	tryRegister(&_SecAccessControlCreateWithFlags, lib, "SecAccessControlCreateWithFlags")
	tryRegister(&_SecAccessCopyACLList, lib, "SecAccessCopyACLList")
	tryRegister(&_SecAccessCopyMatchingACLList, lib, "SecAccessCopyMatchingACLList")
	tryRegister(&_SecAccessCopySelectedACLList, lib, "SecAccessCopySelectedACLList")
	tryRegister(&_SecAccessCreate, lib, "SecAccessCreate")
	tryRegister(&_SecAccessGetOwnerAndACL, lib, "SecAccessGetOwnerAndACL")
	tryRegister(&_SecAccessGetTypeID, lib, "SecAccessGetTypeID")
	tryRegister(&_SecAddSharedWebCredential, lib, "SecAddSharedWebCredential")
	tryRegister(&_SecAsn1AllocCopy, lib, "SecAsn1AllocCopy")
	tryRegister(&_SecAsn1AllocCopyItem, lib, "SecAsn1AllocCopyItem")
	tryRegister(&_SecAsn1AllocItem, lib, "SecAsn1AllocItem")
	tryRegister(&_SecAsn1CoderCreate, lib, "SecAsn1CoderCreate")
	tryRegister(&_SecAsn1CoderRelease, lib, "SecAsn1CoderRelease")
	tryRegister(&_SecAsn1Decode, lib, "SecAsn1Decode")
	tryRegister(&_SecAsn1DecodeData, lib, "SecAsn1DecodeData")
	tryRegister(&_SecAsn1EncodeItem, lib, "SecAsn1EncodeItem")
	tryRegister(&_SecAsn1Malloc, lib, "SecAsn1Malloc")
	tryRegister(&_SecAsn1OidCompare, lib, "SecAsn1OidCompare")
	tryRegister(&_SecCertificateCopyData, lib, "SecCertificateCopyData")
	tryRegister(&_SecCertificateCopyNotValidAfterDate, lib, "SecCertificateCopyNotValidAfterDate")
	tryRegister(&_SecCertificateCopyNotValidBeforeDate, lib, "SecCertificateCopyNotValidBeforeDate")
	tryRegister(&_SecCertificateCreateWithData, lib, "SecCertificateCreateWithData")
	tryRegister(&_SecCodeCheckValidity, lib, "SecCodeCheckValidity")
	tryRegister(&_SecCodeCheckValidityWithErrors, lib, "SecCodeCheckValidityWithErrors")
	tryRegister(&_SecCodeCopyDesignatedRequirement, lib, "SecCodeCopyDesignatedRequirement")
	tryRegister(&_SecCodeCopyGuestWithAttributes, lib, "SecCodeCopyGuestWithAttributes")
	tryRegister(&_SecCodeCopyHost, lib, "SecCodeCopyHost")
	tryRegister(&_SecCodeCopyPath, lib, "SecCodeCopyPath")
	tryRegister(&_SecCodeCopySelf, lib, "SecCodeCopySelf")
	tryRegister(&_SecCodeCopySigningInformation, lib, "SecCodeCopySigningInformation")
	tryRegister(&_SecCodeCopyStaticCode, lib, "SecCodeCopyStaticCode")
	tryRegister(&_SecCodeCreateWithXPCMessage, lib, "SecCodeCreateWithXPCMessage")
	tryRegister(&_SecCodeGetTypeID, lib, "SecCodeGetTypeID")
	tryRegister(&_SecCodeMapMemory, lib, "SecCodeMapMemory")
	tryRegister(&_SecCodeValidateFileResource, lib, "SecCodeValidateFileResource")
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
	tryRegister(&_SecHostCreateGuest, lib, "SecHostCreateGuest")
	tryRegister(&_SecHostRemoveGuest, lib, "SecHostRemoveGuest")
	tryRegister(&_SecHostSelectGuest, lib, "SecHostSelectGuest")
	tryRegister(&_SecHostSelectedGuest, lib, "SecHostSelectedGuest")
	tryRegister(&_SecHostSetGuestStatus, lib, "SecHostSetGuestStatus")
	tryRegister(&_SecHostSetHostingPort, lib, "SecHostSetHostingPort")
	tryRegister(&_SecIdentityCreate, lib, "SecIdentityCreate")
	tryRegister(&_SecItemAdd, lib, "SecItemAdd")
	tryRegister(&_SecItemCopyMatching, lib, "SecItemCopyMatching")
	tryRegister(&_SecItemDelete, lib, "SecItemDelete")
	tryRegister(&_SecItemUpdate, lib, "SecItemUpdate")
	tryRegister(&_SecKeyCopyExternalRepresentation, lib, "SecKeyCopyExternalRepresentation")
	tryRegister(&_SecKeyCopyPublicKey, lib, "SecKeyCopyPublicKey")
	tryRegister(&_SecKeyCreateDecryptedData, lib, "SecKeyCreateDecryptedData")
	tryRegister(&_SecKeyCreateEncryptedData, lib, "SecKeyCreateEncryptedData")
	tryRegister(&_SecKeyCreateRandomKey, lib, "SecKeyCreateRandomKey")
	tryRegister(&_SecKeyCreateWithData, lib, "SecKeyCreateWithData")
	tryRegister(&_SecKeyIsAlgorithmSupported, lib, "SecKeyIsAlgorithmSupported")
	tryRegister(&_SecKeyUnwrapSymmetric, lib, "SecKeyUnwrapSymmetric")
	tryRegister(&_SecKeychainItemCopyAccess, lib, "SecKeychainItemCopyAccess")
	tryRegister(&_SecPolicyCreateBasicX509, lib, "SecPolicyCreateBasicX509")
	tryRegister(&_SecPolicyCreateSSL, lib, "SecPolicyCreateSSL")
	tryRegister(&_SecPolicyCreateWithProperties, lib, "SecPolicyCreateWithProperties")
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
	tryRegister(&_SecTaskGetCodeSignStatus, lib, "SecTaskGetCodeSignStatus")
	tryRegister(&_SecTaskGetTypeID, lib, "SecTaskGetTypeID")
	tryRegister(&_SecTranformCustomGetAttribute, lib, "SecTranformCustomGetAttribute")
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
	tryRegister(&_SecTrustCopyCertificateChain, lib, "SecTrustCopyCertificateChain")
	tryRegister(&_SecTrustCopyKey, lib, "SecTrustCopyKey")
	tryRegister(&_SecTrustCreateWithCertificates, lib, "SecTrustCreateWithCertificates")
	tryRegister(&_SecTrustEvaluate, lib, "SecTrustEvaluate")
	tryRegister(&_SecTrustEvaluateWithError, lib, "SecTrustEvaluateWithError")
	tryRegister(&_SecTrustGetTrustResult, lib, "SecTrustGetTrustResult")
	tryRegister(&_SecTrustSetAnchorCertificates, lib, "SecTrustSetAnchorCertificates")
	tryRegister(&_SecTrustSetVerifyDate, lib, "SecTrustSetVerifyDate")
	tryRegister(&_SecVerifyTransformCreate, lib, "SecVerifyTransformCreate")
	tryRegister(&_SecureDownloadCopyCreationDate, lib, "SecureDownloadCopyCreationDate")
	tryRegister(&_SecureDownloadCopyName, lib, "SecureDownloadCopyName")
	tryRegister(&_SecureDownloadCopyTicketLocation, lib, "SecureDownloadCopyTicketLocation")
	tryRegister(&_SecureDownloadCopyURLs, lib, "SecureDownloadCopyURLs")
	tryRegister(&_SecureDownloadCreateWithTicket, lib, "SecureDownloadCreateWithTicket")
	tryRegister(&_SecureDownloadFinished, lib, "SecureDownloadFinished")
	tryRegister(&_SecureDownloadGetDownloadSize, lib, "SecureDownloadGetDownloadSize")
	tryRegister(&_SecureDownloadRelease, lib, "SecureDownloadRelease")
	tryRegister(&_SecureDownloadUpdateWithData, lib, "SecureDownloadUpdateWithData")
	tryRegister(&_SessionCreate, lib, "SessionCreate")
	tryRegister(&_SessionGetInfo, lib, "SessionGetInfo")
	tryRegister(&_cssmAlgToOid, lib, "cssmAlgToOid")
	tryRegister(&_cssmOidToAlg, lib, "cssmOidToAlg")
	tryRegister(&_cssmPerror, lib, "cssmPerror")
	tryRegister(&_sec_certificate_copy_ref, lib, "sec_certificate_copy_ref")
	tryRegister(&_sec_certificate_create, lib, "sec_certificate_create")
	tryRegister(&_sec_identity_access_certificates, lib, "sec_identity_access_certificates")
	tryRegister(&_sec_identity_copy_certificates_ref, lib, "sec_identity_copy_certificates_ref")
	tryRegister(&_sec_identity_copy_ref, lib, "sec_identity_copy_ref")
	tryRegister(&_sec_identity_create, lib, "sec_identity_create")
	tryRegister(&_sec_identity_create_with_certificates, lib, "sec_identity_create_with_certificates")
	tryRegister(&_sec_protocol_metadata_access_distinguished_names, lib, "sec_protocol_metadata_access_distinguished_names")
	tryRegister(&_sec_protocol_metadata_access_ocsp_response, lib, "sec_protocol_metadata_access_ocsp_response")
	tryRegister(&_sec_protocol_metadata_access_peer_certificate_chain, lib, "sec_protocol_metadata_access_peer_certificate_chain")
	tryRegister(&_sec_protocol_metadata_access_pre_shared_keys, lib, "sec_protocol_metadata_access_pre_shared_keys")
	tryRegister(&_sec_protocol_metadata_access_supported_signature_algorithms, lib, "sec_protocol_metadata_access_supported_signature_algorithms")
	tryRegister(&_sec_protocol_metadata_challenge_parameters_are_equal, lib, "sec_protocol_metadata_challenge_parameters_are_equal")
	tryRegister(&_sec_protocol_metadata_copy_negotiated_protocol, lib, "sec_protocol_metadata_copy_negotiated_protocol")
	tryRegister(&_sec_protocol_metadata_copy_peer_public_key, lib, "sec_protocol_metadata_copy_peer_public_key")
	tryRegister(&_sec_protocol_metadata_copy_server_name, lib, "sec_protocol_metadata_copy_server_name")
	tryRegister(&_sec_protocol_metadata_create_secret, lib, "sec_protocol_metadata_create_secret")
	tryRegister(&_sec_protocol_metadata_create_secret_with_context, lib, "sec_protocol_metadata_create_secret_with_context")
	tryRegister(&_sec_protocol_metadata_get_early_data_accepted, lib, "sec_protocol_metadata_get_early_data_accepted")
	tryRegister(&_sec_protocol_metadata_get_negotiated_ciphersuite, lib, "sec_protocol_metadata_get_negotiated_ciphersuite")
	tryRegister(&_sec_protocol_metadata_get_negotiated_protocol, lib, "sec_protocol_metadata_get_negotiated_protocol")
	tryRegister(&_sec_protocol_metadata_get_negotiated_protocol_version, lib, "sec_protocol_metadata_get_negotiated_protocol_version")
	tryRegister(&_sec_protocol_metadata_get_negotiated_tls_ciphersuite, lib, "sec_protocol_metadata_get_negotiated_tls_ciphersuite")
	tryRegister(&_sec_protocol_metadata_get_negotiated_tls_protocol_version, lib, "sec_protocol_metadata_get_negotiated_tls_protocol_version")
	tryRegister(&_sec_protocol_metadata_get_server_name, lib, "sec_protocol_metadata_get_server_name")
	tryRegister(&_sec_protocol_metadata_peers_are_equal, lib, "sec_protocol_metadata_peers_are_equal")
	tryRegister(&_sec_protocol_options_add_pre_shared_key, lib, "sec_protocol_options_add_pre_shared_key")
	tryRegister(&_sec_protocol_options_add_tls_application_protocol, lib, "sec_protocol_options_add_tls_application_protocol")
	tryRegister(&_sec_protocol_options_add_tls_ciphersuite, lib, "sec_protocol_options_add_tls_ciphersuite")
	tryRegister(&_sec_protocol_options_add_tls_ciphersuite_group, lib, "sec_protocol_options_add_tls_ciphersuite_group")
	tryRegister(&_sec_protocol_options_append_tls_ciphersuite, lib, "sec_protocol_options_append_tls_ciphersuite")
	tryRegister(&_sec_protocol_options_append_tls_ciphersuite_group, lib, "sec_protocol_options_append_tls_ciphersuite_group")
	tryRegister(&_sec_protocol_options_are_equal, lib, "sec_protocol_options_are_equal")
	tryRegister(&_sec_protocol_options_get_default_max_dtls_protocol_version, lib, "sec_protocol_options_get_default_max_dtls_protocol_version")
	tryRegister(&_sec_protocol_options_get_default_max_tls_protocol_version, lib, "sec_protocol_options_get_default_max_tls_protocol_version")
	tryRegister(&_sec_protocol_options_get_default_min_dtls_protocol_version, lib, "sec_protocol_options_get_default_min_dtls_protocol_version")
	tryRegister(&_sec_protocol_options_get_default_min_tls_protocol_version, lib, "sec_protocol_options_get_default_min_tls_protocol_version")
	tryRegister(&_sec_protocol_options_get_enable_encrypted_client_hello, lib, "sec_protocol_options_get_enable_encrypted_client_hello")
	tryRegister(&_sec_protocol_options_get_quic_use_legacy_codepoint, lib, "sec_protocol_options_get_quic_use_legacy_codepoint")
	tryRegister(&_sec_protocol_options_set_challenge_block, lib, "sec_protocol_options_set_challenge_block")
	tryRegister(&_sec_protocol_options_set_enable_encrypted_client_hello, lib, "sec_protocol_options_set_enable_encrypted_client_hello")
	tryRegister(&_sec_protocol_options_set_key_update_block, lib, "sec_protocol_options_set_key_update_block")
	tryRegister(&_sec_protocol_options_set_local_identity, lib, "sec_protocol_options_set_local_identity")
	tryRegister(&_sec_protocol_options_set_max_tls_protocol_version, lib, "sec_protocol_options_set_max_tls_protocol_version")
	tryRegister(&_sec_protocol_options_set_min_tls_protocol_version, lib, "sec_protocol_options_set_min_tls_protocol_version")
	tryRegister(&_sec_protocol_options_set_peer_authentication_optional, lib, "sec_protocol_options_set_peer_authentication_optional")
	tryRegister(&_sec_protocol_options_set_peer_authentication_required, lib, "sec_protocol_options_set_peer_authentication_required")
	tryRegister(&_sec_protocol_options_set_pre_shared_key_selection_block, lib, "sec_protocol_options_set_pre_shared_key_selection_block")
	tryRegister(&_sec_protocol_options_set_quic_use_legacy_codepoint, lib, "sec_protocol_options_set_quic_use_legacy_codepoint")
	tryRegister(&_sec_protocol_options_set_tls_diffie_hellman_parameters, lib, "sec_protocol_options_set_tls_diffie_hellman_parameters")
	tryRegister(&_sec_protocol_options_set_tls_false_start_enabled, lib, "sec_protocol_options_set_tls_false_start_enabled")
	tryRegister(&_sec_protocol_options_set_tls_is_fallback_attempt, lib, "sec_protocol_options_set_tls_is_fallback_attempt")
	tryRegister(&_sec_protocol_options_set_tls_max_version, lib, "sec_protocol_options_set_tls_max_version")
	tryRegister(&_sec_protocol_options_set_tls_min_version, lib, "sec_protocol_options_set_tls_min_version")
	tryRegister(&_sec_protocol_options_set_tls_ocsp_enabled, lib, "sec_protocol_options_set_tls_ocsp_enabled")
	tryRegister(&_sec_protocol_options_set_tls_pre_shared_key_identity_hint, lib, "sec_protocol_options_set_tls_pre_shared_key_identity_hint")
	tryRegister(&_sec_protocol_options_set_tls_renegotiation_enabled, lib, "sec_protocol_options_set_tls_renegotiation_enabled")
	tryRegister(&_sec_protocol_options_set_tls_resumption_enabled, lib, "sec_protocol_options_set_tls_resumption_enabled")
	tryRegister(&_sec_protocol_options_set_tls_sct_enabled, lib, "sec_protocol_options_set_tls_sct_enabled")
	tryRegister(&_sec_protocol_options_set_tls_server_name, lib, "sec_protocol_options_set_tls_server_name")
	tryRegister(&_sec_protocol_options_set_tls_tickets_enabled, lib, "sec_protocol_options_set_tls_tickets_enabled")
	tryRegister(&_sec_protocol_options_set_verify_block, lib, "sec_protocol_options_set_verify_block")
	tryRegister(&_sec_release, lib, "sec_release")
	tryRegister(&_sec_retain, lib, "sec_retain")
	tryRegister(&_sec_trust_copy_ref, lib, "sec_trust_copy_ref")
	tryRegister(&_sec_trust_create, lib, "sec_trust_create")
	tryRegister(&_SecTrustEvaluateAsync, lib, "SecTrustEvaluateAsync")
	tryRegister(&_SecTrustSetKeychains, lib, "SecTrustSetKeychains")
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



// Creates a search object for finding identities. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1396821-secidentitysearchcreate
func SecIdentitySearchCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecIdentitySearchCreate(p0)
	}


// Returns the unique identifier of the opaque type to which a object belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1396823-secidentitysearchgettypeid
func SecIdentitySearchGetTypeID() unsafe.Pointer {
	return _SecIdentitySearchGetTypeID()
	}


// Finds the next identity matching specified search criteria [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1396825-secidentitysearchcopynext
func SecIdentitySearchCopyNext(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecIdentitySearchCopyNext(p0)
	}


// Retrieves the trust policy handle. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1524309-sectrustgettphandle
func SecTrustGetTPHandle(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetTPHandle(p0)
	}


// Retrieves the CSSM trust result. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1524311-sectrustgetcssmresult
func SecTrustGetCssmResult(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetCssmResult(p0)
	}


// Sets the action and action data for a trust management object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1524326-sectrustsetparameters
func SecTrustSetParameters(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetParameters(p0)
	}


// Retrieves the CSSM result code from the most recent trust evaluation for a trust management object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1524327-sectrustgetcssmresultcode
func SecTrustGetCssmResultCode(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetCssmResultCode(p0)
	}


// Retrieves details on the outcome of a call to the function . [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1524331-sectrustgetresult
func SecTrustGetResult(p0 unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetResult(p0)
	}


// Runs an executable tool with root privileges. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1540038-authorizationexecutewithprivileg
func AuthorizationExecuteWithPrivileges(p0 unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationExecuteWithPrivileges(p0)
	}


// Initializes the plug-in and exchanges interfaces with the authorization engine. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/security/1543160-authorizationplugincreate
func AuthorizationPluginCreate(p0 unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationPluginCreate(p0)
	}


// Retrieves supporting data such as the user name and other information gathered during evaluation of authorization. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyInfo(_:_:_:)
func AuthorizationCopyInfo(authorization unsafe.Pointer, tag unsafe.Pointer, info unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCopyInfo(authorization, tag, info)
	}


// Retrieves the authorization reference passed by the AuthorizationExecuteWithPrivileges function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyPrivilegedReference
func AuthorizationCopyPrivilegedReference(authorization unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCopyPrivilegedReference(authorization, flags)
	}


// Authorizes and preauthorizes rights synchronously. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyRights(_:_:_:_:_:)
func AuthorizationCopyRights(authorization unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer, flags unsafe.Pointer, authorizedRights unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCopyRights(authorization, rights, environment, flags, authorizedRights)
	}


// Authorizes and preauthorizes rights asynchronously. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCopyRightsAsync(_:_:_:_:_:)
func AuthorizationCopyRightsAsync(authorization unsafe.Pointer, rights unsafe.Pointer, environment unsafe.Pointer, flags unsafe.Pointer, callbackBlock unsafe.Pointer) {
	_AuthorizationCopyRightsAsync(authorization, rights, environment, flags, callbackBlock)
	}


// Creates a new authorization reference and provides an option to authorize or preauthorize rights. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCreate(_:_:_:_:)
func AuthorizationCreate(rights unsafe.Pointer, environment unsafe.Pointer, flags unsafe.Pointer, authorization unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCreate(rights, environment, flags, authorization)
	}


// Internalizes the external representation of an authorization reference. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCreateFromExternalForm(_:_:)
func AuthorizationCreateFromExternalForm(extForm unsafe.Pointer, authorization unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationCreateFromExternalForm(extForm, authorization)
	}


// Frees the memory associated with an authorization reference. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFree(_:_:)
func AuthorizationFree(authorization unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationFree(authorization, flags)
	}


// Frees the memory associated with a set of authorization items. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationFreeItemSet(_:)
func AuthorizationFreeItemSet(set unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationFreeItemSet(set)
	}


// Creates an external representation of an authorization reference. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationMakeExternalForm(_:_:)
func AuthorizationMakeExternalForm(authorization unsafe.Pointer, extForm unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationMakeExternalForm(authorization, extForm)
	}


// Retrieves a right definition as a dictionary. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightGet(_:_:)
func AuthorizationRightGet(rightName unsafe.Pointer, rightDefinition unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationRightGet(rightName, rightDefinition)
	}


// Removes a right from the policy database. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightRemove(_:_:)
func AuthorizationRightRemove(authRef unsafe.Pointer, rightName unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationRightRemove(authRef, rightName)
	}


// Creates or updates a right entry in the policy database. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRightSet(_:_:_:_:_:_:)
func AuthorizationRightSet(authRef unsafe.Pointer, rightName unsafe.Pointer, rightDefinition unsafe.Pointer, descriptionKey unsafe.Pointer, bundle unsafe.Pointer, localeTableName unsafe.Pointer) unsafe.Pointer {
	return _AuthorizationRightSet(authRef, rightName, rightDefinition, descriptionKey, bundle, localeTableName)
	}


// Obtains an array of all of the certificates in a message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyAllCerts(_:_:)
func CMSDecoderCopyAllCerts(cmsDecoder unsafe.Pointer, certsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyAllCerts(cmsDecoder, certsOut)
	}


// Obtains the message content, if any. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyContent(_:_:)
func CMSDecoderCopyContent(cmsDecoder unsafe.Pointer, contentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyContent(cmsDecoder, contentOut)
	}


// Obtains the detached content specified with the function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyDetachedContent(_:_:)
func CMSDecoderCopyDetachedContent(cmsDecoder unsafe.Pointer, detachedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyDetachedContent(cmsDecoder, detachedContentOut)
	}


// Obtains the object identifier for the encapsulated data of a signed message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopyEncapsulatedContentType(_:_:)
func CMSDecoderCopyEncapsulatedContentType(cmsDecoder unsafe.Pointer, eContentTypeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopyEncapsulatedContentType(cmsDecoder, eContentTypeOut)
	}


// Obtains the certificate of the specified signer of a CMS message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerCert(_:_:_:)
func CMSDecoderCopySignerCert(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, signerCertOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerCert(cmsDecoder, signerIndex, signerCertOut)
	}


// Obtains the email address of the specified signer of a CMS message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerEmailAddress(_:_:_:)
func CMSDecoderCopySignerEmailAddress(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, signerEmailAddressOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerEmailAddress(cmsDecoder, signerIndex, signerEmailAddressOut)
	}


// Obtains the signing time of a CMS message, if present. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerSigningTime(_:_:_:)
func CMSDecoderCopySignerSigningTime(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, signingTime unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerSigningTime(cmsDecoder, signerIndex, signingTime)
	}


// Obtains the status of a CMS message’s signature. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerStatus(_:_:_:_:_:_:_:)
func CMSDecoderCopySignerStatus(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, policyOrArray unsafe.Pointer, evaluateSecTrust unsafe.Pointer, signerStatusOut unsafe.Pointer, secTrustOut unsafe.Pointer, certVerifyResultCodeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerStatus(cmsDecoder, signerIndex, policyOrArray, evaluateSecTrust, signerStatusOut, secTrustOut, certVerifyResultCodeOut)
	}


// Returns the timestamp of a signer of a CMS message, if present. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestamp(_:_:_:)
func CMSDecoderCopySignerTimestamp(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestamp(cmsDecoder, signerIndex, timestamp)
	}


// Returns an array containing the certificates from a timestamp response. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampCertificates(_:_:_:)
func CMSDecoderCopySignerTimestampCertificates(cmsDecoder unsafe.Pointer, signerIndex unsafe.Pointer, certificateRefs unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestampCertificates(cmsDecoder, signerIndex, certificateRefs)
	}


// Returns the timestamp of a signer of a CMS message using a given policy, if present. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder unsafe.Pointer, timeStampPolicy unsafe.Pointer, signerIndex unsafe.Pointer, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCopySignerTimestampWithPolicy(cmsDecoder, timeStampPolicy, signerIndex, timestamp)
	}


// Creates a CMSDecoder reference. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderCreate(_:)
func CMSDecoderCreate(cmsDecoderOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderCreate(cmsDecoderOut)
	}


// Indicates that there is no more data to decode. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderFinalizeMessage(_:)
func CMSDecoderFinalizeMessage(cmsDecoder unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderFinalizeMessage(cmsDecoder)
	}


// Obtains the number of signers of a message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderGetNumSigners(_:_:)
func CMSDecoderGetNumSigners(cmsDecoder unsafe.Pointer, numSignersOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderGetNumSigners(cmsDecoder, numSignersOut)
	}


// Returns the type identifier for the CMSDecoder opaque type. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderGetTypeID()
func CMSDecoderGetTypeID() unsafe.Pointer {
	return _CMSDecoderGetTypeID()
	}


// Determines whether a CMS message was encrypted. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderIsContentEncrypted(_:_:)
func CMSDecoderIsContentEncrypted(cmsDecoder unsafe.Pointer, isEncryptedOut unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderIsContentEncrypted(cmsDecoder, isEncryptedOut)
	}


// Specifies the message’s detached content, if any. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderSetDetachedContent(_:_:)
func CMSDecoderSetDetachedContent(cmsDecoder unsafe.Pointer, detachedContent unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderSetDetachedContent(cmsDecoder, detachedContent)
	}


// Specifies the keychains to search for intermediate certificates to be used in verifying a signed message’s signer certificates. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderSetSearchKeychain(_:_:)
func CMSDecoderSetSearchKeychain(cmsDecoder unsafe.Pointer, keychainOrArray unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderSetSearchKeychain(cmsDecoder, keychainOrArray)
	}


// Feeds raw bytes of the message to be decoded into the decoder. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoderUpdateMessage(_:_:_:)
func CMSDecoderUpdateMessage(cmsDecoder unsafe.Pointer, msgBytes unsafe.Pointer, msgBytesLen unsafe.Pointer) unsafe.Pointer {
	return _CMSDecoderUpdateMessage(cmsDecoder, msgBytes, msgBytesLen)
	}


// Encodes a message and obtains the result in one high-level function call. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncode
func CMSEncode(signers unsafe.Pointer, recipients unsafe.Pointer, eContentType unsafe.Pointer, detachedContent unsafe.Pointer, signedAttributes unsafe.Pointer, content unsafe.Pointer, contentLen unsafe.Pointer, encodedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncode(signers, recipients, eContentType, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	}


// Encodes a message and obtains the result in one high-level function call. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncodeContent(_:_:_:_:_:_:_:_:)
func CMSEncodeContent(signers unsafe.Pointer, recipients unsafe.Pointer, eContentTypeOID unsafe.Pointer, detachedContent unsafe.Pointer, signedAttributes unsafe.Pointer, content unsafe.Pointer, contentLen unsafe.Pointer, encodedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncodeContent(signers, recipients, eContentTypeOID, detachedContent, signedAttributes, content, contentLen, encodedContentOut)
	}


// Specifies a message is to be encrypted and specifies the recipients of the message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddRecipients(_:_:)
func CMSEncoderAddRecipients(cmsEncoder unsafe.Pointer, recipientOrArray unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderAddRecipients(cmsEncoder, recipientOrArray)
	}


// Specifies attributes for a signed message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSignedAttributes(_:_:)
func CMSEncoderAddSignedAttributes(cmsEncoder unsafe.Pointer, signedAttributes unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderAddSignedAttributes(cmsEncoder, signedAttributes)
	}


// Specifies signers of the message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSigners(_:_:)
func CMSEncoderAddSigners(cmsEncoder unsafe.Pointer, signerOrArray unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderAddSigners(cmsEncoder, signerOrArray)
	}


// Adds certificates to a message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderAddSupportingCerts(_:_:)
func CMSEncoderAddSupportingCerts(cmsEncoder unsafe.Pointer, certOrArray unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderAddSupportingCerts(cmsEncoder, certOrArray)
	}


// Obtains the object identifier for the encapsulated data of a signed message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncapsulatedContentType(_:_:)
func CMSEncoderCopyEncapsulatedContentType(cmsEncoder unsafe.Pointer, eContentTypeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyEncapsulatedContentType(cmsEncoder, eContentTypeOut)
	}


// Finishes encoding the message and obtains the encoded result. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyEncodedContent(_:_:)
func CMSEncoderCopyEncodedContent(cmsEncoder unsafe.Pointer, encodedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyEncodedContent(cmsEncoder, encodedContentOut)
	}


// Obtains the array of recipients specified with the function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopyRecipients(_:_:)
func CMSEncoderCopyRecipients(cmsEncoder unsafe.Pointer, recipientsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopyRecipients(cmsEncoder, recipientsOut)
	}


// Returns the timestamp of a signer of a CMS message, if present. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestamp(_:_:_:)
func CMSEncoderCopySignerTimestamp(cmsEncoder unsafe.Pointer, signerIndex unsafe.Pointer, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySignerTimestamp(cmsEncoder, signerIndex, timestamp)
	}


// Returns the timestamp of a signer of a CMS message using a particular policy, if present. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySignerTimestampWithPolicy(_:_:_:_:)
func CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder unsafe.Pointer, timeStampPolicy unsafe.Pointer, signerIndex unsafe.Pointer, timestamp unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySignerTimestampWithPolicy(cmsEncoder, timeStampPolicy, signerIndex, timestamp)
	}


// Obtains the array of signers specified with the function. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySigners(_:_:)
func CMSEncoderCopySigners(cmsEncoder unsafe.Pointer, signersOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySigners(cmsEncoder, signersOut)
	}


// Obtains the certificates added to a message with . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCopySupportingCerts(_:_:)
func CMSEncoderCopySupportingCerts(cmsEncoder unsafe.Pointer, certsOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCopySupportingCerts(cmsEncoder, certsOut)
	}


// Creates a CMSEncoder reference. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderCreate(_:)
func CMSEncoderCreate(cmsEncoderOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderCreate(cmsEncoderOut)
	}


// Obtains a constant that indicates which certificates are to be included in a signed CMS message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetCertificateChainMode(_:_:)
func CMSEncoderGetCertificateChainMode(cmsEncoder unsafe.Pointer, chainModeOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderGetCertificateChainMode(cmsEncoder, chainModeOut)
	}


// Indicates whether the message is to have detached content. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetHasDetachedContent(_:_:)
func CMSEncoderGetHasDetachedContent(cmsEncoder unsafe.Pointer, detachedContentOut unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderGetHasDetachedContent(cmsEncoder, detachedContentOut)
	}


// Returns the type identifier for the CMSEncoder opaque type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderGetTypeID()
func CMSEncoderGetTypeID() unsafe.Pointer {
	return _CMSEncoderGetTypeID()
	}


// Specifies which certificates to include in a signed CMS message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetCertificateChainMode(_:_:)
func CMSEncoderSetCertificateChainMode(cmsEncoder unsafe.Pointer, chainMode unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetCertificateChainMode(cmsEncoder, chainMode)
	}


// Specifies an object identifier for the encapsulated data of a signed message. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentType
func CMSEncoderSetEncapsulatedContentType(cmsEncoder unsafe.Pointer, eContentType unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetEncapsulatedContentType(cmsEncoder, eContentType)
	}


// Specifies an object identifier for the encapsulated data of a signed message. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetEncapsulatedContentTypeOID(_:_:)
func CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder unsafe.Pointer, eContentTypeOID unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetEncapsulatedContentTypeOID(cmsEncoder, eContentTypeOID)
	}


// Specifies whether the signed data is to be separate from the message. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetHasDetachedContent(_:_:)
func CMSEncoderSetHasDetachedContent(cmsEncoder unsafe.Pointer, detachedContent unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetHasDetachedContent(cmsEncoder, detachedContent)
	}


// Sets the digest algorithm to use for the signer. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderSetSignerAlgorithm(_:_:)
func CMSEncoderSetSignerAlgorithm(cmsEncoder unsafe.Pointer, digestAlgorithm unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderSetSignerAlgorithm(cmsEncoder, digestAlgorithm)
	}


// Feeds content bytes into the encoder. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoderUpdateContent(_:_:_:)
func CMSEncoderUpdateContent(cmsEncoder unsafe.Pointer, content unsafe.Pointer, contentLen unsafe.Pointer) unsafe.Pointer {
	return _CMSEncoderUpdateContent(cmsEncoder, content, contentLen)
	}


// CSSM_AC_AuthCompute is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_AC_AuthCompute
func CSSM_AC_AuthCompute(ACHandle unsafe.Pointer, BaseAuthorizations unsafe.Pointer, Credentials unsafe.Pointer, NumberOfRequestors unsafe.Pointer, Requestors unsafe.Pointer, RequestedAuthorizationPeriod unsafe.Pointer, RequestedAuthorization unsafe.Pointer, AuthorizationResult unsafe.Pointer) unsafe.Pointer {
	return _CSSM_AC_AuthCompute(ACHandle, BaseAuthorizations, Credentials, NumberOfRequestors, Requestors, RequestedAuthorizationPeriod, RequestedAuthorization, AuthorizationResult)
	}


// CSSM_AC_PassThrough is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_AC_PassThrough
func CSSM_AC_PassThrough(ACHandle unsafe.Pointer, TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, DBList unsafe.Pointer, PassThroughId unsafe.Pointer, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) unsafe.Pointer {
	return _CSSM_AC_PassThrough(ACHandle, TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	}


// CSSM_CL_CertAbortCache is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortCache
func CSSM_CL_CertAbortCache(CLHandle unsafe.Pointer, CertHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertAbortCache(CLHandle, CertHandle)
	}


// CSSM_CL_CertAbortQuery is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertAbortQuery
func CSSM_CL_CertAbortQuery(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertAbortQuery(CLHandle, ResultsHandle)
	}


// CSSM_CL_CertCache is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertCache
func CSSM_CL_CertCache(CLHandle unsafe.Pointer, Cert unsafe.Pointer, CertHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertCache(CLHandle, Cert, CertHandle)
	}


// CSSM_CL_CertCreateTemplate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertCreateTemplate
func CSSM_CL_CertCreateTemplate(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertCreateTemplate(CLHandle, NumberOfFields, CertFields, CertTemplate)
	}


// CSSM_CL_CertDescribeFormat is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertDescribeFormat
func CSSM_CL_CertDescribeFormat(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, OidList unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertDescribeFormat(CLHandle, NumberOfFields, OidList)
	}


// CSSM_CL_CertGetAllFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllFields
func CSSM_CL_CertGetAllFields(CLHandle unsafe.Pointer, Cert unsafe.Pointer, NumberOfFields unsafe.Pointer, CertFields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetAllFields(CLHandle, Cert, NumberOfFields, CertFields)
	}


// CSSM_CL_CertGetAllTemplateFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetAllTemplateFields
func CSSM_CL_CertGetAllTemplateFields(CLHandle unsafe.Pointer, CertTemplate unsafe.Pointer, NumberOfFields unsafe.Pointer, CertFields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetAllTemplateFields(CLHandle, CertTemplate, NumberOfFields, CertFields)
	}


// CSSM_CL_CertGetFirstCachedFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstCachedFieldValue
func CSSM_CL_CertGetFirstCachedFieldValue(CLHandle unsafe.Pointer, CertHandle unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle unsafe.Pointer, NumberOfMatchedFields unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetFirstCachedFieldValue(CLHandle, CertHandle, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	}


// CSSM_CL_CertGetFirstFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetFirstFieldValue
func CSSM_CL_CertGetFirstFieldValue(CLHandle unsafe.Pointer, Cert unsafe.Pointer, CertField unsafe.Pointer, ResultsHandle unsafe.Pointer, NumberOfMatchedFields unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetFirstFieldValue(CLHandle, Cert, CertField, ResultsHandle, NumberOfMatchedFields, Value)
	}


// CSSM_CL_CertGetKeyInfo is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetKeyInfo
func CSSM_CL_CertGetKeyInfo(CLHandle unsafe.Pointer, Cert unsafe.Pointer, Key unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetKeyInfo(CLHandle, Cert, Key)
	}


// CSSM_CL_CertGetNextCachedFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextCachedFieldValue
func CSSM_CL_CertGetNextCachedFieldValue(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	}


// CSSM_CL_CertGetNextFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGetNextFieldValue
func CSSM_CL_CertGetNextFieldValue(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGetNextFieldValue(CLHandle, ResultsHandle, Value)
	}


// CSSM_CL_CertGroupFromVerifiedBundle is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupFromVerifiedBundle
func CSSM_CL_CertGroupFromVerifiedBundle(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertBundle unsafe.Pointer, SignerCert unsafe.Pointer, CertGroup unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGroupFromVerifiedBundle(CLHandle, CCHandle, CertBundle, SignerCert, CertGroup)
	}


// CSSM_CL_CertGroupToSignedBundle is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertGroupToSignedBundle
func CSSM_CL_CertGroupToSignedBundle(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertGroupToBundle unsafe.Pointer, BundleInfo unsafe.Pointer, SignedBundle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertGroupToSignedBundle(CLHandle, CCHandle, CertGroupToBundle, BundleInfo, SignedBundle)
	}


// CSSM_CL_CertSign is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertSign
func CSSM_CL_CertSign(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertTemplate unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize unsafe.Pointer, SignedCert unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertSign(CLHandle, CCHandle, CertTemplate, SignScope, ScopeSize, SignedCert)
	}


// CSSM_CL_CertVerify is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerify
func CSSM_CL_CertVerify(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertVerify(CLHandle, CCHandle, CertToBeVerified, SignerCert, VerifyScope, ScopeSize)
	}


// CSSM_CL_CertVerifyWithKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CertVerifyWithKey
func CSSM_CL_CertVerifyWithKey(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertToBeVerified unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CertVerifyWithKey(CLHandle, CCHandle, CertToBeVerified)
	}


// CSSM_CL_CrlAbortCache is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortCache
func CSSM_CL_CrlAbortCache(CLHandle unsafe.Pointer, CrlHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlAbortCache(CLHandle, CrlHandle)
	}


// CSSM_CL_CrlAbortQuery is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAbortQuery
func CSSM_CL_CrlAbortQuery(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlAbortQuery(CLHandle, ResultsHandle)
	}


// CSSM_CL_CrlAddCert is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlAddCert
func CSSM_CL_CrlAddCert(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, Cert unsafe.Pointer, NumberOfFields unsafe.Pointer, CrlEntryFields unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlAddCert(CLHandle, CCHandle, Cert, NumberOfFields, CrlEntryFields, OldCrl, NewCrl)
	}


// CSSM_CL_CrlCache is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCache
func CSSM_CL_CrlCache(CLHandle unsafe.Pointer, Crl unsafe.Pointer, CrlHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlCache(CLHandle, Crl, CrlHandle)
	}


// CSSM_CL_CrlCreateTemplate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlCreateTemplate
func CSSM_CL_CrlCreateTemplate(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, CrlTemplate unsafe.Pointer, NewCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlCreateTemplate(CLHandle, NumberOfFields, CrlTemplate, NewCrl)
	}


// CSSM_CL_CrlDescribeFormat is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlDescribeFormat
func CSSM_CL_CrlDescribeFormat(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, OidList unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlDescribeFormat(CLHandle, NumberOfFields, OidList)
	}


// CSSM_CL_CrlGetAllCachedRecordFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllCachedRecordFields
func CSSM_CL_CrlGetAllCachedRecordFields(CLHandle unsafe.Pointer, CrlHandle unsafe.Pointer, CrlRecordIndex unsafe.Pointer, NumberOfFields unsafe.Pointer, CrlFields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetAllCachedRecordFields(CLHandle, CrlHandle, CrlRecordIndex, NumberOfFields, CrlFields)
	}


// CSSM_CL_CrlGetAllFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetAllFields
func CSSM_CL_CrlGetAllFields(CLHandle unsafe.Pointer, Crl unsafe.Pointer, NumberOfCrlFields unsafe.Pointer, CrlFields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetAllFields(CLHandle, Crl, NumberOfCrlFields, CrlFields)
	}


// CSSM_CL_CrlGetFirstCachedFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstCachedFieldValue
func CSSM_CL_CrlGetFirstCachedFieldValue(CLHandle unsafe.Pointer, CrlHandle unsafe.Pointer, CrlRecordIndex unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle unsafe.Pointer, NumberOfMatchedFields unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetFirstCachedFieldValue(CLHandle, CrlHandle, CrlRecordIndex, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	}


// CSSM_CL_CrlGetFirstFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetFirstFieldValue
func CSSM_CL_CrlGetFirstFieldValue(CLHandle unsafe.Pointer, Crl unsafe.Pointer, CrlField unsafe.Pointer, ResultsHandle unsafe.Pointer, NumberOfMatchedFields unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetFirstFieldValue(CLHandle, Crl, CrlField, ResultsHandle, NumberOfMatchedFields, Value)
	}


// CSSM_CL_CrlGetNextCachedFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextCachedFieldValue
func CSSM_CL_CrlGetNextCachedFieldValue(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetNextCachedFieldValue(CLHandle, ResultsHandle, Value)
	}


// CSSM_CL_CrlGetNextFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlGetNextFieldValue
func CSSM_CL_CrlGetNextFieldValue(CLHandle unsafe.Pointer, ResultsHandle unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlGetNextFieldValue(CLHandle, ResultsHandle, Value)
	}


// CSSM_CL_CrlRemoveCert is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlRemoveCert
func CSSM_CL_CrlRemoveCert(CLHandle unsafe.Pointer, Cert unsafe.Pointer, OldCrl unsafe.Pointer, NewCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlRemoveCert(CLHandle, Cert, OldCrl, NewCrl)
	}


// CSSM_CL_CrlSetFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSetFields
func CSSM_CL_CrlSetFields(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, CrlTemplate unsafe.Pointer, OldCrl unsafe.Pointer, ModifiedCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlSetFields(CLHandle, NumberOfFields, CrlTemplate, OldCrl, ModifiedCrl)
	}


// CSSM_CL_CrlSign is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlSign
func CSSM_CL_CrlSign(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, UnsignedCrl unsafe.Pointer, SignScope unsafe.Pointer, ScopeSize unsafe.Pointer, SignedCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlSign(CLHandle, CCHandle, UnsignedCrl, SignScope, ScopeSize, SignedCrl)
	}


// CSSM_CL_CrlVerify is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerify
func CSSM_CL_CrlVerify(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CrlToBeVerified unsafe.Pointer, SignerCert unsafe.Pointer, VerifyScope unsafe.Pointer, ScopeSize unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlVerify(CLHandle, CCHandle, CrlToBeVerified, SignerCert, VerifyScope, ScopeSize)
	}


// CSSM_CL_CrlVerifyWithKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_CrlVerifyWithKey
func CSSM_CL_CrlVerifyWithKey(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CrlToBeVerified unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_CrlVerifyWithKey(CLHandle, CCHandle, CrlToBeVerified)
	}


// CSSM_CL_FreeFieldValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFieldValue
func CSSM_CL_FreeFieldValue(CLHandle unsafe.Pointer, CertOrCrlOid unsafe.Pointer, Value unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_FreeFieldValue(CLHandle, CertOrCrlOid, Value)
	}


// CSSM_CL_FreeFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_FreeFields
func CSSM_CL_FreeFields(CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, Fields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_FreeFields(CLHandle, NumberOfFields, Fields)
	}


// CSSM_CL_IsCertInCachedCrl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCachedCrl
func CSSM_CL_IsCertInCachedCrl(CLHandle unsafe.Pointer, Cert unsafe.Pointer, CrlHandle unsafe.Pointer, CertFound unsafe.Pointer, CrlRecordIndex unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_IsCertInCachedCrl(CLHandle, Cert, CrlHandle, CertFound, CrlRecordIndex)
	}


// CSSM_CL_IsCertInCrl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_IsCertInCrl
func CSSM_CL_IsCertInCrl(CLHandle unsafe.Pointer, Cert unsafe.Pointer, Crl unsafe.Pointer, CertFound unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_IsCertInCrl(CLHandle, Cert, Crl, CertFound)
	}


// CSSM_CL_PassThrough is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_PassThrough
func CSSM_CL_PassThrough(CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, PassThroughId unsafe.Pointer, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CL_PassThrough(CLHandle, CCHandle, PassThroughId, InputParams, OutputParams)
	}


// CSSM_CSP_ChangeLoginAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginAcl
func CSSM_CSP_ChangeLoginAcl(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_ChangeLoginAcl(CSPHandle, AccessCred, AclEdit)
	}


// CSSM_CSP_ChangeLoginOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_ChangeLoginOwner
func CSSM_CSP_ChangeLoginOwner(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_ChangeLoginOwner(CSPHandle, AccessCred, NewOwner)
	}


// CSSM_CSP_CreateAsymmetricContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateAsymmetricContext
func CSSM_CSP_CreateAsymmetricContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, Padding unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateAsymmetricContext(CSPHandle, AlgorithmID, AccessCred, Key, Padding, NewContextHandle)
	}


// CSSM_CSP_CreateDeriveKeyContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDeriveKeyContext
func CSSM_CSP_CreateDeriveKeyContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, DeriveKeyType unsafe.Pointer, DeriveKeyLengthInBits unsafe.Pointer, AccessCred unsafe.Pointer, BaseKey unsafe.Pointer, IterationCount unsafe.Pointer, Salt unsafe.Pointer, Seed unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateDeriveKeyContext(CSPHandle, AlgorithmID, DeriveKeyType, DeriveKeyLengthInBits, AccessCred, BaseKey, IterationCount, Salt, Seed, NewContextHandle)
	}


// CSSM_CSP_CreateDigestContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateDigestContext
func CSSM_CSP_CreateDigestContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateDigestContext(CSPHandle, AlgorithmID, NewContextHandle)
	}


// CSSM_CSP_CreateKeyGenContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateKeyGenContext
func CSSM_CSP_CreateKeyGenContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, KeySizeInBits unsafe.Pointer, Seed unsafe.Pointer, Salt unsafe.Pointer, StartDate unsafe.Pointer, EndDate unsafe.Pointer, Params unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateKeyGenContext(CSPHandle, AlgorithmID, KeySizeInBits, Seed, Salt, StartDate, EndDate, Params, NewContextHandle)
	}


// CSSM_CSP_CreateMacContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateMacContext
func CSSM_CSP_CreateMacContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateMacContext(CSPHandle, AlgorithmID, Key, NewContextHandle)
	}


// CSSM_CSP_CreatePassThroughContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreatePassThroughContext
func CSSM_CSP_CreatePassThroughContext(CSPHandle unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreatePassThroughContext(CSPHandle, Key, NewContextHandle)
	}


// CSSM_CSP_CreateRandomGenContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateRandomGenContext
func CSSM_CSP_CreateRandomGenContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, Seed unsafe.Pointer, Length unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateRandomGenContext(CSPHandle, AlgorithmID, Seed, Length, NewContextHandle)
	}


// CSSM_CSP_CreateSignatureContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSignatureContext
func CSSM_CSP_CreateSignatureContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateSignatureContext(CSPHandle, AlgorithmID, AccessCred, Key, NewContextHandle)
	}


// CSSM_CSP_CreateSymmetricContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_CreateSymmetricContext
func CSSM_CSP_CreateSymmetricContext(CSPHandle unsafe.Pointer, AlgorithmID unsafe.Pointer, Mode unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, InitVector unsafe.Pointer, Padding unsafe.Pointer, Reserved unsafe.Pointer, NewContextHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_CreateSymmetricContext(CSPHandle, AlgorithmID, Mode, AccessCred, Key, InitVector, Padding, Reserved, NewContextHandle)
	}


// CSSM_CSP_GetLoginAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginAcl
func CSSM_CSP_GetLoginAcl(CSPHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos unsafe.Pointer, AclInfos unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_GetLoginAcl(CSPHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	}


// CSSM_CSP_GetLoginOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_GetLoginOwner
func CSSM_CSP_GetLoginOwner(CSPHandle unsafe.Pointer, Owner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_GetLoginOwner(CSPHandle, Owner)
	}


// CSSM_CSP_GetOperationalStatistics is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_GetOperationalStatistics
func CSSM_CSP_GetOperationalStatistics(CSPHandle unsafe.Pointer, Statistics unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_GetOperationalStatistics(CSPHandle, Statistics)
	}


// CSSM_CSP_Login is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_Login
func CSSM_CSP_Login(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, LoginName unsafe.Pointer, Reserved unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_Login(CSPHandle, AccessCred, LoginName, Reserved)
	}


// CSSM_CSP_Logout is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_Logout
func CSSM_CSP_Logout(CSPHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_Logout(CSPHandle)
	}


// CSSM_CSP_ObtainPrivateKeyFromPublicKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_ObtainPrivateKeyFromPublicKey
func CSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKey unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_ObtainPrivateKeyFromPublicKey(CSPHandle, PublicKey, PrivateKey)
	}


// CSSM_CSP_PassThrough is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_PassThrough
func CSSM_CSP_PassThrough(CCHandle unsafe.Pointer, PassThroughId unsafe.Pointer, InData unsafe.Pointer, OutData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_CSP_PassThrough(CCHandle, PassThroughId, InData, OutData)
	}


// CSSM_ChangeKeyAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyAcl
func CSSM_ChangeKeyAcl(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer, Key unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ChangeKeyAcl(CSPHandle, AccessCred, AclEdit, Key)
	}


// CSSM_ChangeKeyOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ChangeKeyOwner
func CSSM_ChangeKeyOwner(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, NewOwner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ChangeKeyOwner(CSPHandle, AccessCred, Key, NewOwner)
	}


// CSSM_DL_Authenticate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_Authenticate
func CSSM_DL_Authenticate(DLDBHandle unsafe.Pointer, AccessRequest unsafe.Pointer, AccessCred unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_Authenticate(DLDBHandle, AccessRequest, AccessCred)
	}


// CSSM_DL_ChangeDbAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbAcl
func CSSM_DL_ChangeDbAcl(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, AclEdit unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_ChangeDbAcl(DLDBHandle, AccessCred, AclEdit)
	}


// CSSM_DL_ChangeDbOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_ChangeDbOwner
func CSSM_DL_ChangeDbOwner(DLDBHandle unsafe.Pointer, AccessCred unsafe.Pointer, NewOwner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_ChangeDbOwner(DLDBHandle, AccessCred, NewOwner)
	}


// CSSM_DL_CreateRelation is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_CreateRelation
func CSSM_DL_CreateRelation(DLDBHandle unsafe.Pointer, RelationID unsafe.Pointer, RelationName unsafe.Pointer, NumberOfAttributes unsafe.Pointer, pAttributeInfo unsafe.Pointer, NumberOfIndexes unsafe.Pointer, pIndexInfo unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_CreateRelation(DLDBHandle, RelationID, RelationName, NumberOfAttributes, pAttributeInfo, NumberOfIndexes, pIndexInfo)
	}


// CSSM_DL_DataAbortQuery is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataAbortQuery
func CSSM_DL_DataAbortQuery(DLDBHandle unsafe.Pointer, ResultsHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataAbortQuery(DLDBHandle, ResultsHandle)
	}


// CSSM_DL_DataDelete is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataDelete
func CSSM_DL_DataDelete(DLDBHandle unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataDelete(DLDBHandle, UniqueRecordIdentifier)
	}


// CSSM_DL_DataGetFirst is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFirst
func CSSM_DL_DataGetFirst(DLDBHandle unsafe.Pointer, Query unsafe.Pointer, ResultsHandle unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataGetFirst(DLDBHandle, Query, ResultsHandle, Attributes, Data, UniqueId)
	}


// CSSM_DL_DataGetFromUniqueRecordId is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetFromUniqueRecordId
func CSSM_DL_DataGetFromUniqueRecordId(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataGetFromUniqueRecordId(DLDBHandle, UniqueRecord, Attributes, Data)
	}


// CSSM_DL_DataGetNext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataGetNext
func CSSM_DL_DataGetNext(DLDBHandle unsafe.Pointer, ResultsHandle unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataGetNext(DLDBHandle, ResultsHandle, Attributes, Data, UniqueId)
	}


// CSSM_DL_DataInsert is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataInsert
func CSSM_DL_DataInsert(DLDBHandle unsafe.Pointer, RecordType unsafe.Pointer, Attributes unsafe.Pointer, Data unsafe.Pointer, UniqueId unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataInsert(DLDBHandle, RecordType, Attributes, Data, UniqueId)
	}


// CSSM_DL_DataModify is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DataModify
func CSSM_DL_DataModify(DLDBHandle unsafe.Pointer, RecordType unsafe.Pointer, UniqueRecordIdentifier unsafe.Pointer, AttributesToBeModified unsafe.Pointer, DataToBeModified unsafe.Pointer, ModifyMode unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DataModify(DLDBHandle, RecordType, UniqueRecordIdentifier, AttributesToBeModified, DataToBeModified, ModifyMode)
	}


// CSSM_DL_DbClose is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DbClose
func CSSM_DL_DbClose(DLDBHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DbClose(DLDBHandle)
	}


// CSSM_DL_DbCreate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DbCreate
func CSSM_DL_DbCreate(DLHandle unsafe.Pointer, DbName unsafe.Pointer, DbLocation unsafe.Pointer, DBInfo unsafe.Pointer, AccessRequest unsafe.Pointer, CredAndAclEntry unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DbCreate(DLHandle, DbName, DbLocation, DBInfo, AccessRequest, CredAndAclEntry, OpenParameters, DbHandle)
	}


// CSSM_DL_DbDelete is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DbDelete
func CSSM_DL_DbDelete(DLHandle unsafe.Pointer, DbName unsafe.Pointer, DbLocation unsafe.Pointer, AccessCred unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DbDelete(DLHandle, DbName, DbLocation, AccessCred)
	}


// CSSM_DL_DbOpen is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DbOpen
func CSSM_DL_DbOpen(DLHandle unsafe.Pointer, DbName unsafe.Pointer, DbLocation unsafe.Pointer, AccessRequest unsafe.Pointer, AccessCred unsafe.Pointer, OpenParameters unsafe.Pointer, DbHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DbOpen(DLHandle, DbName, DbLocation, AccessRequest, AccessCred, OpenParameters, DbHandle)
	}


// CSSM_DL_DestroyRelation is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_DestroyRelation
func CSSM_DL_DestroyRelation(DLDBHandle unsafe.Pointer, RelationID unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_DestroyRelation(DLDBHandle, RelationID)
	}


// CSSM_DL_FreeNameList is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_FreeNameList
func CSSM_DL_FreeNameList(DLHandle unsafe.Pointer, NameList unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_FreeNameList(DLHandle, NameList)
	}


// CSSM_DL_FreeUniqueRecord is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_FreeUniqueRecord
func CSSM_DL_FreeUniqueRecord(DLDBHandle unsafe.Pointer, UniqueRecord unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_FreeUniqueRecord(DLDBHandle, UniqueRecord)
	}


// CSSM_DL_GetDbAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbAcl
func CSSM_DL_GetDbAcl(DLDBHandle unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos unsafe.Pointer, AclInfos unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_GetDbAcl(DLDBHandle, SelectionTag, NumberOfAclInfos, AclInfos)
	}


// CSSM_DL_GetDbNameFromHandle is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNameFromHandle
func CSSM_DL_GetDbNameFromHandle(DLDBHandle unsafe.Pointer, DbName unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_GetDbNameFromHandle(DLDBHandle, DbName)
	}


// CSSM_DL_GetDbNames is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbNames
func CSSM_DL_GetDbNames(DLHandle unsafe.Pointer, NameList unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_GetDbNames(DLHandle, NameList)
	}


// CSSM_DL_GetDbOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_GetDbOwner
func CSSM_DL_GetDbOwner(DLDBHandle unsafe.Pointer, Owner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_GetDbOwner(DLDBHandle, Owner)
	}


// CSSM_DL_PassThrough is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_PassThrough
func CSSM_DL_PassThrough(DLDBHandle unsafe.Pointer, PassThroughId unsafe.Pointer, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DL_PassThrough(DLDBHandle, PassThroughId, InputParams, OutputParams)
	}


// CSSM_DecryptData is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptData
func CSSM_DecryptData(CCHandle unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptData(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData)
	}


// CSSM_DecryptDataFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptDataFinal
func CSSM_DecryptDataFinal(CCHandle unsafe.Pointer, RemData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptDataFinal(CCHandle, RemData)
	}


// CSSM_DecryptDataInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInit
func CSSM_DecryptDataInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptDataInit(CCHandle)
	}


// CSSM_DecryptDataInitP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptDataInitP
func CSSM_DecryptDataInitP(CCHandle unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptDataInitP(CCHandle, Privilege)
	}


// CSSM_DecryptDataP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptDataP
func CSSM_DecryptDataP(CCHandle unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, bytesDecrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptDataP(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted, RemData, Privilege)
	}


// CSSM_DecryptDataUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DecryptDataUpdate
func CSSM_DecryptDataUpdate(CCHandle unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, bytesDecrypted unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DecryptDataUpdate(CCHandle, CipherBufs, CipherBufCount, ClearBufs, ClearBufCount, bytesDecrypted)
	}


// CSSM_DeleteContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DeleteContext
func CSSM_DeleteContext(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DeleteContext(CCHandle)
	}


// CSSM_DeleteContextAttributes is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DeleteContextAttributes
func CSSM_DeleteContextAttributes(CCHandle unsafe.Pointer, NumberOfAttributes unsafe.Pointer, ContextAttributes unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DeleteContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	}


// CSSM_DeriveKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DeriveKey
func CSSM_DeriveKey(CCHandle unsafe.Pointer, Param unsafe.Pointer, KeyUsage unsafe.Pointer, KeyAttr unsafe.Pointer, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, DerivedKey unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DeriveKey(CCHandle, Param, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, DerivedKey)
	}


// CSSM_DigestData is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DigestData
func CSSM_DigestData(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer, Digest unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DigestData(CCHandle, DataBufs, DataBufCount, Digest)
	}


// CSSM_DigestDataClone is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DigestDataClone
func CSSM_DigestDataClone(CCHandle unsafe.Pointer, ClonednewCCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DigestDataClone(CCHandle, ClonednewCCHandle)
	}


// CSSM_DigestDataFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DigestDataFinal
func CSSM_DigestDataFinal(CCHandle unsafe.Pointer, Digest unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DigestDataFinal(CCHandle, Digest)
	}


// CSSM_DigestDataInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DigestDataInit
func CSSM_DigestDataInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DigestDataInit(CCHandle)
	}


// CSSM_DigestDataUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DigestDataUpdate
func CSSM_DigestDataUpdate(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer) unsafe.Pointer {
	return _CSSM_DigestDataUpdate(CCHandle, DataBufs, DataBufCount)
	}


// CSSM_EncryptData is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptData
func CSSM_EncryptData(CCHandle unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptData(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData)
	}


// CSSM_EncryptDataFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptDataFinal
func CSSM_EncryptDataFinal(CCHandle unsafe.Pointer, RemData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptDataFinal(CCHandle, RemData)
	}


// CSSM_EncryptDataInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInit
func CSSM_EncryptDataInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptDataInit(CCHandle)
	}


// CSSM_EncryptDataInitP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptDataInitP
func CSSM_EncryptDataInitP(CCHandle unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptDataInitP(CCHandle, Privilege)
	}


// CSSM_EncryptDataP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptDataP
func CSSM_EncryptDataP(CCHandle unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, bytesEncrypted unsafe.Pointer, RemData unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptDataP(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted, RemData, Privilege)
	}


// CSSM_EncryptDataUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EncryptDataUpdate
func CSSM_EncryptDataUpdate(CCHandle unsafe.Pointer, ClearBufs unsafe.Pointer, ClearBufCount unsafe.Pointer, CipherBufs unsafe.Pointer, CipherBufCount unsafe.Pointer, bytesEncrypted unsafe.Pointer) unsafe.Pointer {
	return _CSSM_EncryptDataUpdate(CCHandle, ClearBufs, ClearBufCount, CipherBufs, CipherBufCount, bytesEncrypted)
	}


// CSSM_FreeContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_FreeContext
func CSSM_FreeContext(Context unsafe.Pointer) unsafe.Pointer {
	return _CSSM_FreeContext(Context)
	}


// CSSM_FreeKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_FreeKey
func CSSM_FreeKey(CSPHandle unsafe.Pointer, AccessCred unsafe.Pointer, KeyPtr unsafe.Pointer, Delete unsafe.Pointer) unsafe.Pointer {
	return _CSSM_FreeKey(CSPHandle, AccessCred, KeyPtr, Delete)
	}


// CSSM_GenerateAlgorithmParams is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateAlgorithmParams
func CSSM_GenerateAlgorithmParams(CCHandle unsafe.Pointer, ParamBits unsafe.Pointer, Param unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateAlgorithmParams(CCHandle, ParamBits, Param)
	}


// CSSM_GenerateKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateKey
func CSSM_GenerateKey(CCHandle unsafe.Pointer, KeyUsage unsafe.Pointer, KeyAttr unsafe.Pointer, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateKey(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key)
	}


// CSSM_GenerateKeyP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyP
func CSSM_GenerateKeyP(CCHandle unsafe.Pointer, KeyUsage unsafe.Pointer, KeyAttr unsafe.Pointer, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, Key unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateKeyP(CCHandle, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, Key, Privilege)
	}


// CSSM_GenerateKeyPair is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPair
func CSSM_GenerateKeyPair(CCHandle unsafe.Pointer, PublicKeyUsage unsafe.Pointer, PublicKeyAttr unsafe.Pointer, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage unsafe.Pointer, PrivateKeyAttr unsafe.Pointer, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateKeyPair(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey)
	}


// CSSM_GenerateKeyPairP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateKeyPairP
func CSSM_GenerateKeyPairP(CCHandle unsafe.Pointer, PublicKeyUsage unsafe.Pointer, PublicKeyAttr unsafe.Pointer, PublicKeyLabel unsafe.Pointer, PublicKey unsafe.Pointer, PrivateKeyUsage unsafe.Pointer, PrivateKeyAttr unsafe.Pointer, PrivateKeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, PrivateKey unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateKeyPairP(CCHandle, PublicKeyUsage, PublicKeyAttr, PublicKeyLabel, PublicKey, PrivateKeyUsage, PrivateKeyAttr, PrivateKeyLabel, CredAndAclEntry, PrivateKey, Privilege)
	}


// CSSM_GenerateMac is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateMac
func CSSM_GenerateMac(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer, Mac unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateMac(CCHandle, DataBufs, DataBufCount, Mac)
	}


// CSSM_GenerateMacFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateMacFinal
func CSSM_GenerateMacFinal(CCHandle unsafe.Pointer, Mac unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateMacFinal(CCHandle, Mac)
	}


// CSSM_GenerateMacInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateMacInit
func CSSM_GenerateMacInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateMacInit(CCHandle)
	}


// CSSM_GenerateMacUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateMacUpdate
func CSSM_GenerateMacUpdate(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateMacUpdate(CCHandle, DataBufs, DataBufCount)
	}


// CSSM_GenerateRandom is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GenerateRandom
func CSSM_GenerateRandom(CCHandle unsafe.Pointer, RandomNumber unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GenerateRandom(CCHandle, RandomNumber)
	}


// CSSM_GetAPIMemoryFunctions is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetAPIMemoryFunctions
func CSSM_GetAPIMemoryFunctions(AddInHandle unsafe.Pointer, AppMemoryFuncs unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetAPIMemoryFunctions(AddInHandle, AppMemoryFuncs)
	}


// CSSM_GetContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetContext
func CSSM_GetContext(CCHandle unsafe.Pointer, Context unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetContext(CCHandle, Context)
	}


// CSSM_GetContextAttribute is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetContextAttribute
func CSSM_GetContextAttribute(Context unsafe.Pointer, AttributeType unsafe.Pointer, ContextAttribute unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetContextAttribute(Context, AttributeType, ContextAttribute)
	}


// CSSM_GetKeyAcl is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetKeyAcl
func CSSM_GetKeyAcl(CSPHandle unsafe.Pointer, Key unsafe.Pointer, SelectionTag unsafe.Pointer, NumberOfAclInfos unsafe.Pointer, AclInfos unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetKeyAcl(CSPHandle, Key, SelectionTag, NumberOfAclInfos, AclInfos)
	}


// CSSM_GetKeyOwner is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetKeyOwner
func CSSM_GetKeyOwner(CSPHandle unsafe.Pointer, Key unsafe.Pointer, Owner unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetKeyOwner(CSPHandle, Key, Owner)
	}


// CSSM_GetModuleGUIDFromHandle is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetModuleGUIDFromHandle
func CSSM_GetModuleGUIDFromHandle(ModuleHandle unsafe.Pointer, ModuleGUID unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetModuleGUIDFromHandle(ModuleHandle, ModuleGUID)
	}


// CSSM_GetPrivilege is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetPrivilege
func CSSM_GetPrivilege(Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetPrivilege(Privilege)
	}


// CSSM_GetSubserviceUIDFromHandle is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetSubserviceUIDFromHandle
func CSSM_GetSubserviceUIDFromHandle(ModuleHandle unsafe.Pointer, SubserviceUID unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetSubserviceUIDFromHandle(ModuleHandle, SubserviceUID)
	}


// CSSM_GetTimeValue is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_GetTimeValue
func CSSM_GetTimeValue(CSPHandle unsafe.Pointer, TimeAlgorithm unsafe.Pointer, TimeData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_GetTimeValue(CSPHandle, TimeAlgorithm, TimeData)
	}


// CSSM_Init is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_Init
func CSSM_Init(Version unsafe.Pointer, Scope unsafe.Pointer, CallerGuid unsafe.Pointer, KeyHierarchy unsafe.Pointer, PvcPolicy unsafe.Pointer, Reserved unsafe.Pointer) unsafe.Pointer {
	return _CSSM_Init(Version, Scope, CallerGuid, KeyHierarchy, PvcPolicy, Reserved)
	}


// CSSM_Introduce is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_Introduce
func CSSM_Introduce(ModuleID unsafe.Pointer, KeyHierarchy unsafe.Pointer) unsafe.Pointer {
	return _CSSM_Introduce(ModuleID, KeyHierarchy)
	}


// CSSM_ListAttachedModuleManagers is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ListAttachedModuleManagers
func CSSM_ListAttachedModuleManagers(NumberOfModuleManagers unsafe.Pointer, ModuleManagerGuids unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ListAttachedModuleManagers(NumberOfModuleManagers, ModuleManagerGuids)
	}


// CSSM_ModuleAttach is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ModuleAttach
func CSSM_ModuleAttach(ModuleGuid unsafe.Pointer, Version unsafe.Pointer, MemoryFuncs unsafe.Pointer, SubserviceID unsafe.Pointer, SubServiceType unsafe.Pointer, AttachFlags unsafe.Pointer, KeyHierarchy unsafe.Pointer, FunctionTable unsafe.Pointer, NumFunctionTable unsafe.Pointer, Reserved unsafe.Pointer, NewModuleHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ModuleAttach(ModuleGuid, Version, MemoryFuncs, SubserviceID, SubServiceType, AttachFlags, KeyHierarchy, FunctionTable, NumFunctionTable, Reserved, NewModuleHandle)
	}


// CSSM_ModuleDetach is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ModuleDetach
func CSSM_ModuleDetach(ModuleHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ModuleDetach(ModuleHandle)
	}


// CSSM_ModuleLoad is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ModuleLoad
func CSSM_ModuleLoad(ModuleGuid unsafe.Pointer, KeyHierarchy unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ModuleLoad(ModuleGuid, KeyHierarchy, AppNotifyCallback, AppNotifyCallbackCtx)
	}


// CSSM_ModuleUnload is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ModuleUnload
func CSSM_ModuleUnload(ModuleGuid unsafe.Pointer, AppNotifyCallback unsafe.Pointer, AppNotifyCallbackCtx unsafe.Pointer) unsafe.Pointer {
	return _CSSM_ModuleUnload(ModuleGuid, AppNotifyCallback, AppNotifyCallbackCtx)
	}


// CSSM_QueryKeySizeInBits is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_QueryKeySizeInBits
func CSSM_QueryKeySizeInBits(CSPHandle unsafe.Pointer, CCHandle unsafe.Pointer, Key unsafe.Pointer, KeySize unsafe.Pointer) unsafe.Pointer {
	return _CSSM_QueryKeySizeInBits(CSPHandle, CCHandle, Key, KeySize)
	}


// CSSM_QuerySize is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_QuerySize
func CSSM_QuerySize(CCHandle unsafe.Pointer, Encrypt unsafe.Pointer, QuerySizeCount unsafe.Pointer, DataBlockSizes unsafe.Pointer) unsafe.Pointer {
	return _CSSM_QuerySize(CCHandle, Encrypt, QuerySizeCount, DataBlockSizes)
	}


// CSSM_RetrieveCounter is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_RetrieveCounter
func CSSM_RetrieveCounter(CSPHandle unsafe.Pointer, Counter unsafe.Pointer) unsafe.Pointer {
	return _CSSM_RetrieveCounter(CSPHandle, Counter)
	}


// CSSM_RetrieveUniqueId is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_RetrieveUniqueId
func CSSM_RetrieveUniqueId(CSPHandle unsafe.Pointer, UniqueID unsafe.Pointer) unsafe.Pointer {
	return _CSSM_RetrieveUniqueId(CSPHandle, UniqueID)
	}


// CSSM_SetContext is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SetContext
func CSSM_SetContext(CCHandle unsafe.Pointer, Context unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SetContext(CCHandle, Context)
	}


// CSSM_SetPrivilege is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SetPrivilege
func CSSM_SetPrivilege(Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SetPrivilege(Privilege)
	}


// CSSM_SignData is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SignData
func CSSM_SignData(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer, DigestAlgorithm unsafe.Pointer, Signature unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SignData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	}


// CSSM_SignDataFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SignDataFinal
func CSSM_SignDataFinal(CCHandle unsafe.Pointer, Signature unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SignDataFinal(CCHandle, Signature)
	}


// CSSM_SignDataInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SignDataInit
func CSSM_SignDataInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SignDataInit(CCHandle)
	}


// CSSM_SignDataUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SignDataUpdate
func CSSM_SignDataUpdate(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer) unsafe.Pointer {
	return _CSSM_SignDataUpdate(CCHandle, DataBufs, DataBufCount)
	}


// CSSM_TP_ApplyCrlToDb is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_ApplyCrlToDb
func CSSM_TP_ApplyCrlToDb(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, CrlToBeApplied unsafe.Pointer, SignerCertGroup unsafe.Pointer, ApplyCrlVerifyContext unsafe.Pointer, ApplyCrlVerifyResult unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_ApplyCrlToDb(TPHandle, CLHandle, CSPHandle, CrlToBeApplied, SignerCertGroup, ApplyCrlVerifyContext, ApplyCrlVerifyResult)
	}


// CSSM_TP_CertCreateTemplate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertCreateTemplate
func CSSM_TP_CertCreateTemplate(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, CertFields unsafe.Pointer, CertTemplate unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertCreateTemplate(TPHandle, CLHandle, NumberOfFields, CertFields, CertTemplate)
	}


// CSSM_TP_CertGetAllTemplateFields is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertGetAllTemplateFields
func CSSM_TP_CertGetAllTemplateFields(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CertTemplate unsafe.Pointer, NumberOfFields unsafe.Pointer, CertFields unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertGetAllTemplateFields(TPHandle, CLHandle, CertTemplate, NumberOfFields, CertFields)
	}


// CSSM_TP_CertGroupConstruct is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupConstruct
func CSSM_TP_CertGroupConstruct(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, DBList unsafe.Pointer, ConstructParams unsafe.Pointer, CertGroupFrag unsafe.Pointer, CertGroup unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertGroupConstruct(TPHandle, CLHandle, CSPHandle, DBList, ConstructParams, CertGroupFrag, CertGroup)
	}


// CSSM_TP_CertGroupPrune is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupPrune
func CSSM_TP_CertGroupPrune(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, DBList unsafe.Pointer, OrderedCertGroup unsafe.Pointer, PrunedCertGroup unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertGroupPrune(TPHandle, CLHandle, DBList, OrderedCertGroup, PrunedCertGroup)
	}


// CSSM_TP_CertGroupToTupleGroup is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupToTupleGroup
func CSSM_TP_CertGroupToTupleGroup(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CertGroup unsafe.Pointer, TupleGroup unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertGroupToTupleGroup(TPHandle, CLHandle, CertGroup, TupleGroup)
	}


// CSSM_TP_CertGroupVerify is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertGroupVerify
func CSSM_TP_CertGroupVerify(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, CertGroupToBeVerified unsafe.Pointer, VerifyContext unsafe.Pointer, VerifyContextResult unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertGroupVerify(TPHandle, CLHandle, CSPHandle, CertGroupToBeVerified, VerifyContext, VerifyContextResult)
	}


// CSSM_TP_CertReclaimAbort is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimAbort
func CSSM_TP_CertReclaimAbort(TPHandle unsafe.Pointer, KeyCacheHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertReclaimAbort(TPHandle, KeyCacheHandle)
	}


// CSSM_TP_CertReclaimKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertReclaimKey
func CSSM_TP_CertReclaimKey(TPHandle unsafe.Pointer, CertGroup unsafe.Pointer, CertIndex unsafe.Pointer, KeyCacheHandle unsafe.Pointer, CSPHandle unsafe.Pointer, CredAndAclEntry unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertReclaimKey(TPHandle, CertGroup, CertIndex, KeyCacheHandle, CSPHandle, CredAndAclEntry)
	}


// CSSM_TP_CertRemoveFromCrlTemplate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertRemoveFromCrlTemplate
func CSSM_TP_CertRemoveFromCrlTemplate(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, OldCrlTemplate unsafe.Pointer, CertGroupToBeRemoved unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, NewCrlTemplate unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertRemoveFromCrlTemplate(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRemoved, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, NewCrlTemplate)
	}


// CSSM_TP_CertRevoke is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertRevoke
func CSSM_TP_CertRevoke(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, OldCrlTemplate unsafe.Pointer, CertGroupToBeRevoked unsafe.Pointer, RevokerCertGroup unsafe.Pointer, RevokerVerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer, Reason unsafe.Pointer, NewCrlTemplate unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertRevoke(TPHandle, CLHandle, CSPHandle, OldCrlTemplate, CertGroupToBeRevoked, RevokerCertGroup, RevokerVerifyContext, RevokerVerifyResult, Reason, NewCrlTemplate)
	}


// CSSM_TP_CertSign is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CertSign
func CSSM_TP_CertSign(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CertTemplateToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCert unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CertSign(TPHandle, CLHandle, CCHandle, CertTemplateToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCert)
	}


// CSSM_TP_ConfirmCredResult is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_ConfirmCredResult
func CSSM_TP_ConfirmCredResult(TPHandle unsafe.Pointer, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, Responses unsafe.Pointer, PreferredAuthority unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_ConfirmCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, Responses, PreferredAuthority)
	}


// CSSM_TP_CrlCreateTemplate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CrlCreateTemplate
func CSSM_TP_CrlCreateTemplate(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, NumberOfFields unsafe.Pointer, CrlFields unsafe.Pointer, NewCrlTemplate unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CrlCreateTemplate(TPHandle, CLHandle, NumberOfFields, CrlFields, NewCrlTemplate)
	}


// CSSM_TP_CrlSign is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CrlSign
func CSSM_TP_CrlSign(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, CrlToBeSigned unsafe.Pointer, SignerCertGroup unsafe.Pointer, SignerVerifyContext unsafe.Pointer, SignerVerifyResult unsafe.Pointer, SignedCrl unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CrlSign(TPHandle, CLHandle, CCHandle, CrlToBeSigned, SignerCertGroup, SignerVerifyContext, SignerVerifyResult, SignedCrl)
	}


// CSSM_TP_CrlVerify is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CrlVerify
func CSSM_TP_CrlVerify(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CSPHandle unsafe.Pointer, CrlToBeVerified unsafe.Pointer, SignerCertGroup unsafe.Pointer, VerifyContext unsafe.Pointer, RevokerVerifyResult unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_CrlVerify(TPHandle, CLHandle, CSPHandle, CrlToBeVerified, SignerCertGroup, VerifyContext, RevokerVerifyResult)
	}


// CSSM_TP_FormRequest is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_FormRequest
func CSSM_TP_FormRequest(TPHandle unsafe.Pointer, PreferredAuthority unsafe.Pointer, FormType unsafe.Pointer, BlankForm unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_FormRequest(TPHandle, PreferredAuthority, FormType, BlankForm)
	}


// CSSM_TP_FormSubmit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_FormSubmit
func CSSM_TP_FormSubmit(TPHandle unsafe.Pointer, FormType unsafe.Pointer, Form unsafe.Pointer, ClearanceAuthority unsafe.Pointer, RepresentedAuthority unsafe.Pointer, Credentials unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_FormSubmit(TPHandle, FormType, Form, ClearanceAuthority, RepresentedAuthority, Credentials)
	}


// CSSM_TP_PassThrough is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_PassThrough
func CSSM_TP_PassThrough(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, CCHandle unsafe.Pointer, DBList unsafe.Pointer, PassThroughId unsafe.Pointer, InputParams unsafe.Pointer, OutputParams unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_PassThrough(TPHandle, CLHandle, CCHandle, DBList, PassThroughId, InputParams, OutputParams)
	}


// CSSM_TP_ReceiveConfirmation is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_ReceiveConfirmation
func CSSM_TP_ReceiveConfirmation(TPHandle unsafe.Pointer, ReferenceIdentifier unsafe.Pointer, Responses unsafe.Pointer, ElapsedTime unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_ReceiveConfirmation(TPHandle, ReferenceIdentifier, Responses, ElapsedTime)
	}


// CSSM_TP_RetrieveCredResult is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_RetrieveCredResult
func CSSM_TP_RetrieveCredResult(TPHandle unsafe.Pointer, ReferenceIdentifier unsafe.Pointer, CallerAuthCredentials unsafe.Pointer, EstimatedTime unsafe.Pointer, ConfirmationRequired unsafe.Pointer, RetrieveOutput unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_RetrieveCredResult(TPHandle, ReferenceIdentifier, CallerAuthCredentials, EstimatedTime, ConfirmationRequired, RetrieveOutput)
	}


// CSSM_TP_SubmitCredRequest is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_SubmitCredRequest
func CSSM_TP_SubmitCredRequest(TPHandle unsafe.Pointer, PreferredAuthority unsafe.Pointer, RequestType unsafe.Pointer, RequestInput unsafe.Pointer, CallerAuthContext unsafe.Pointer, EstimatedTime unsafe.Pointer, ReferenceIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_SubmitCredRequest(TPHandle, PreferredAuthority, RequestType, RequestInput, CallerAuthContext, EstimatedTime, ReferenceIdentifier)
	}


// CSSM_TP_TupleGroupToCertGroup is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_TupleGroupToCertGroup
func CSSM_TP_TupleGroupToCertGroup(TPHandle unsafe.Pointer, CLHandle unsafe.Pointer, TupleGroup unsafe.Pointer, CertTemplates unsafe.Pointer) unsafe.Pointer {
	return _CSSM_TP_TupleGroupToCertGroup(TPHandle, CLHandle, TupleGroup, CertTemplates)
	}


// CSSM_Terminate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_Terminate
func CSSM_Terminate() unsafe.Pointer {
	return _CSSM_Terminate()
	}


// CSSM_Unintroduce is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_Unintroduce
func CSSM_Unintroduce(ModuleID unsafe.Pointer) unsafe.Pointer {
	return _CSSM_Unintroduce(ModuleID)
	}


// CSSM_UnwrapKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_UnwrapKey
func CSSM_UnwrapKey(CCHandle unsafe.Pointer, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage unsafe.Pointer, KeyAttr unsafe.Pointer, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer) unsafe.Pointer {
	return _CSSM_UnwrapKey(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData)
	}


// CSSM_UnwrapKeyP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_UnwrapKeyP
func CSSM_UnwrapKeyP(CCHandle unsafe.Pointer, PublicKey unsafe.Pointer, WrappedKey unsafe.Pointer, KeyUsage unsafe.Pointer, KeyAttr unsafe.Pointer, KeyLabel unsafe.Pointer, CredAndAclEntry unsafe.Pointer, UnwrappedKey unsafe.Pointer, DescriptiveData unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_UnwrapKeyP(CCHandle, PublicKey, WrappedKey, KeyUsage, KeyAttr, KeyLabel, CredAndAclEntry, UnwrappedKey, DescriptiveData, Privilege)
	}


// CSSM_UpdateContextAttributes is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_UpdateContextAttributes
func CSSM_UpdateContextAttributes(CCHandle unsafe.Pointer, NumberOfAttributes unsafe.Pointer, ContextAttributes unsafe.Pointer) unsafe.Pointer {
	return _CSSM_UpdateContextAttributes(CCHandle, NumberOfAttributes, ContextAttributes)
	}


// CSSM_VerifyData is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyData
func CSSM_VerifyData(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer, DigestAlgorithm unsafe.Pointer, Signature unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyData(CCHandle, DataBufs, DataBufCount, DigestAlgorithm, Signature)
	}


// CSSM_VerifyDataFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyDataFinal
func CSSM_VerifyDataFinal(CCHandle unsafe.Pointer, Signature unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyDataFinal(CCHandle, Signature)
	}


// CSSM_VerifyDataInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyDataInit
func CSSM_VerifyDataInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyDataInit(CCHandle)
	}


// CSSM_VerifyDataUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyDataUpdate
func CSSM_VerifyDataUpdate(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyDataUpdate(CCHandle, DataBufs, DataBufCount)
	}


// CSSM_VerifyDevice is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyDevice
func CSSM_VerifyDevice(CSPHandle unsafe.Pointer, DeviceCert unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyDevice(CSPHandle, DeviceCert)
	}


// CSSM_VerifyMac is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyMac
func CSSM_VerifyMac(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer, Mac unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyMac(CCHandle, DataBufs, DataBufCount, Mac)
	}


// CSSM_VerifyMacFinal is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyMacFinal
func CSSM_VerifyMacFinal(CCHandle unsafe.Pointer, Mac unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyMacFinal(CCHandle, Mac)
	}


// CSSM_VerifyMacInit is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyMacInit
func CSSM_VerifyMacInit(CCHandle unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyMacInit(CCHandle)
	}


// CSSM_VerifyMacUpdate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_VerifyMacUpdate
func CSSM_VerifyMacUpdate(CCHandle unsafe.Pointer, DataBufs unsafe.Pointer, DataBufCount unsafe.Pointer) unsafe.Pointer {
	return _CSSM_VerifyMacUpdate(CCHandle, DataBufs, DataBufCount)
	}


// CSSM_WrapKey is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_WrapKey
func CSSM_WrapKey(CCHandle unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer) unsafe.Pointer {
	return _CSSM_WrapKey(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey)
	}


// CSSM_WrapKeyP is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_WrapKeyP
func CSSM_WrapKeyP(CCHandle unsafe.Pointer, AccessCred unsafe.Pointer, Key unsafe.Pointer, DescriptiveData unsafe.Pointer, WrappedKey unsafe.Pointer, Privilege unsafe.Pointer) unsafe.Pointer {
	return _CSSM_WrapKeyP(CCHandle, AccessCred, Key, DescriptiveData, WrappedKey, Privilege)
	}


// MDS_Initialize is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/MDS_Initialize
func MDS_Initialize(pCallerGuid unsafe.Pointer, pMemoryFunctions unsafe.Pointer, pDlFunctions unsafe.Pointer, hMds unsafe.Pointer) unsafe.Pointer {
	return _MDS_Initialize(pCallerGuid, pMemoryFunctions, pDlFunctions, hMds)
	}


// MDS_Install is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/MDS_Install
func MDS_Install(MdsHandle unsafe.Pointer) unsafe.Pointer {
	return _MDS_Install(MdsHandle)
	}


// MDS_Terminate is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/MDS_Terminate
func MDS_Terminate(MdsHandle unsafe.Pointer) unsafe.Pointer {
	return _MDS_Terminate(MdsHandle)
	}


// MDS_Uninstall is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/MDS_Uninstall
func MDS_Uninstall(MdsHandle unsafe.Pointer) unsafe.Pointer {
	return _MDS_Uninstall(MdsHandle)
	}


// Adds a DER-encoded distinguished name to a list of acceptable names to be specified in requests for client certificates. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLAddDistinguishedName(_:_:_:)
func SSLAddDistinguishedName(context unsafe.Pointer, derDN unsafe.Pointer, derDNLen unsafe.Pointer) unsafe.Pointer {
	return _SSLAddDistinguishedName(context, derDN, derDNLen)
	}


// Terminates the current SSL session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLClose(_:)
func SSLClose(context unsafe.Pointer) unsafe.Pointer {
	return _SSLClose(context)
	}


// Returns the Core Foundation type ID for context objects. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLContextGetTypeID()
func SSLContextGetTypeID() unsafe.Pointer {
	return _SSLContextGetTypeID()
	}


// Gets the list of supported application layer protocols. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyALPNProtocols(_:_:)
func SSLCopyALPNProtocols(context unsafe.Pointer, protocols unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyALPNProtocols(context, protocols)
	}


// Retrieves the current list of certification authorities. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyCertificateAuthorities(_:_:)
func SSLCopyCertificateAuthorities(context unsafe.Pointer, certificates unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyCertificateAuthorities(context, certificates)
	}


// Retrieves the distinguished names of acceptable certification authorities. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyDistinguishedNames(_:_:)
func SSLCopyDistinguishedNames(context unsafe.Pointer, names unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyDistinguishedNames(context, names)
	}


// Retrieves a peer certificate and its certificate chain. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyPeerCertificates
func SSLCopyPeerCertificates(context unsafe.Pointer, certs unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyPeerCertificates(context, certs)
	}


// Retrieves a trust management object for the certificate used by a session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyPeerTrust(_:_:)
func SSLCopyPeerTrust(context unsafe.Pointer, trust unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyPeerTrust(context, trust)
	}


// Determines the buffer size needed for the peer domain name. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyRequestedPeerName(_:_:_:)
func SSLCopyRequestedPeerName(context unsafe.Pointer, peerName unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyRequestedPeerName(context, peerName, peerNameLen)
	}


// Obtains the hostname specified by the client in the ServerName extension (SNI). Server only. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.11.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyRequestedPeerNameLength(_:_:)
func SSLCopyRequestedPeerNameLength(ctx unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyRequestedPeerNameLength(ctx, peerNameLen)
	}


// Retrieves the current list of trusted root certificates. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCopyTrustedRoots
func SSLCopyTrustedRoots(context unsafe.Pointer, trustedRoots unsafe.Pointer) unsafe.Pointer {
	return _SSLCopyTrustedRoots(context, trustedRoots)
	}


// Allocates and returns a new context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCreateContext(_:_:_:)
func SSLCreateContext(alloc unsafe.Pointer, protocolSide unsafe.Pointer, connectionType unsafe.Pointer) unsafe.Pointer {
	return _SSLCreateContext(alloc, protocolSide, connectionType)
	}


// Disposes of a Secure Sockets Layer (SSL) session context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLDisposeContext
func SSLDisposeContext(context unsafe.Pointer) unsafe.Pointer {
	return _SSLDisposeContext(context)
	}


// Obtains a value specifying whether an unknown root is allowed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetAllowsAnyRoot
func SSLGetAllowsAnyRoot(context unsafe.Pointer, anyRoot unsafe.Pointer) unsafe.Pointer {
	return _SSLGetAllowsAnyRoot(context, anyRoot)
	}


// Retrieves the value specifying whether expired certificates are allowed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetAllowsExpiredCerts
func SSLGetAllowsExpiredCerts(context unsafe.Pointer, allowsExpired unsafe.Pointer) unsafe.Pointer {
	return _SSLGetAllowsExpiredCerts(context, allowsExpired)
	}


// Retrieves the value indicating whether expired roots are allowed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetAllowsExpiredRoots
func SSLGetAllowsExpiredRoots(context unsafe.Pointer, allowsExpired unsafe.Pointer) unsafe.Pointer {
	return _SSLGetAllowsExpiredRoots(context, allowsExpired)
	}


// Determines how much data is available to be read. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetBufferedReadSize(_:_:)
func SSLGetBufferedReadSize(context unsafe.Pointer, bufferSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetBufferedReadSize(context, bufferSize)
	}


// Retrieves the exchange status of the client certificate. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetClientCertificateState(_:_:)
func SSLGetClientCertificateState(context unsafe.Pointer, clientState unsafe.Pointer) unsafe.Pointer {
	return _SSLGetClientCertificateState(context, clientState)
	}


// Retrieves an I/O connection—such as a socket or endpoint—for a specific session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetConnection(_:_:)
func SSLGetConnection(context unsafe.Pointer, connection unsafe.Pointer) unsafe.Pointer {
	return _SSLGetConnection(context, connection)
	}


// Provides the largest packet that the OS guarantees it can send without fragmentation. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetDatagramWriteSize(_:_:)
func SSLGetDatagramWriteSize(dtlsContext unsafe.Pointer, bufSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetDatagramWriteSize(dtlsContext, bufSize)
	}


// Retrieves the Diffie-Hellman parameters for a given context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetDiffieHellmanParams(_:_:_:)
func SSLGetDiffieHellmanParams(context unsafe.Pointer, dhParams unsafe.Pointer, dhParamsLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetDiffieHellmanParams(context, dhParams, dhParamsLen)
	}


// Determines whether peer certificate chain validation is currently enabled. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetEnableCertVerify
func SSLGetEnableCertVerify(context unsafe.Pointer, enableVerify unsafe.Pointer) unsafe.Pointer {
	return _SSLGetEnableCertVerify(context, enableVerify)
	}


// Determines which SSL cipher suites are currently enabled. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetEnabledCiphers(_:_:_:)
func SSLGetEnabledCiphers(context unsafe.Pointer, ciphers unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetEnabledCiphers(context, ciphers, numCiphers)
	}


// Obtains the maximum datagram record size allowed by the application for a given context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetMaxDatagramRecordSize(_:_:)
func SSLGetMaxDatagramRecordSize(dtlsContext unsafe.Pointer, maxSize unsafe.Pointer) unsafe.Pointer {
	return _SSLGetMaxDatagramRecordSize(dtlsContext, maxSize)
	}


// Retrieves the cipher suite negotiated for this session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNegotiatedCipher(_:_:)
func SSLGetNegotiatedCipher(context unsafe.Pointer, cipherSuite unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNegotiatedCipher(context, cipherSuite)
	}


// Obtains the negotiated protocol version of the active session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNegotiatedProtocolVersion(_:_:)
func SSLGetNegotiatedProtocolVersion(context unsafe.Pointer, protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNegotiatedProtocolVersion(context, protocol_)
	}


// Determines the number of cipher suites currently enabled. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNumberEnabledCiphers(_:_:)
func SSLGetNumberEnabledCiphers(context unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNumberEnabledCiphers(context, numCiphers)
	}


// Determines the number of cipher suites supported. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetNumberSupportedCiphers(_:_:)
func SSLGetNumberSupportedCiphers(context unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetNumberSupportedCiphers(context, numCiphers)
	}


// Retrieves the peer domain name specified previously. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerDomainName(_:_:_:)
func SSLGetPeerDomainName(context unsafe.Pointer, peerName unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerDomainName(context, peerName, peerNameLen)
	}


// Determines the length of a previously set peer domain name. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerDomainNameLength(_:_:)
func SSLGetPeerDomainNameLength(context unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerDomainNameLength(context, peerNameLen)
	}


// Retrieves the current peer ID data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetPeerID(_:_:_:)
func SSLGetPeerID(context unsafe.Pointer, peerID unsafe.Pointer, peerIDLen unsafe.Pointer) unsafe.Pointer {
	return _SSLGetPeerID(context, peerID, peerIDLen)
	}


// Gets the SSL protocol version. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersion
func SSLGetProtocolVersion(context unsafe.Pointer, protocol_ unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersion(context, protocol_)
	}


// Retrieves the enabled status of a given protocol. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersionEnabled
func SSLGetProtocolVersionEnabled(context unsafe.Pointer, protocol_ unsafe.Pointer, enable unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersionEnabled(context, protocol_, enable)
	}


// Gets the maximum protocol version allowed by the application for a given SSL context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersionMax(_:_:)
func SSLGetProtocolVersionMax(context unsafe.Pointer, maxVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersionMax(context, maxVersion)
	}


// Gets the minimum protocol version allowed by the application for a given SSL context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetProtocolVersionMin(_:_:)
func SSLGetProtocolVersionMin(context unsafe.Pointer, minVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLGetProtocolVersionMin(context, minVersion)
	}


// Obtains a value indicating whether RSA blinding is enabled. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetRsaBlinding
func SSLGetRsaBlinding(context unsafe.Pointer, blinding unsafe.Pointer) unsafe.Pointer {
	return _SSLGetRsaBlinding(context, blinding)
	}


// Indicates the current setting of Secure Sockets Layer (SSL) session options. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSessionOption(_:_:_:)
func SSLGetSessionOption(context unsafe.Pointer, option unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSessionOption(context, option, value)
	}


// Retrieves the state of an SSL session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSessionState(_:_:)
func SSLGetSessionState(context unsafe.Pointer, state unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSessionState(context, state)
	}


// Determines the values of the supported cipher suites. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLGetSupportedCiphers(_:_:_:)
func SSLGetSupportedCiphers(context unsafe.Pointer, ciphers unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLGetSupportedCiphers(context, ciphers, numCiphers)
	}


// Performs the SSL handshake. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLHandshake(_:)
func SSLHandshake(context unsafe.Pointer) unsafe.Pointer {
	return _SSLHandshake(context)
	}


// Creates a new Secure Sockets Layer (SSL) session context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLNewContext
func SSLNewContext(isServer unsafe.Pointer, contextPtr unsafe.Pointer) unsafe.Pointer {
	return _SSLNewContext(isServer, contextPtr)
	}


// Requests renegotiation of the SSL handshake. Server only. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLReHandshake(_:)
func SSLReHandshake(context unsafe.Pointer) unsafe.Pointer {
	return _SSLReHandshake(context)
	}


// Performs a normal application-level read operation. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLRead(_:_:_:_:)
func SSLRead(context unsafe.Pointer, data unsafe.Pointer, dataLength unsafe.Pointer, processed unsafe.Pointer) unsafe.Pointer {
	return _SSLRead(context, data, dataLength, processed)
	}


// Sets the list of supported applicaiton layer protocols. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetALPNProtocols(_:_:)
func SSLSetALPNProtocols(context unsafe.Pointer, protocols unsafe.Pointer) unsafe.Pointer {
	return _SSLSetALPNProtocols(context, protocols)
	}


// Specifies whether root certificates from unrecognized certification authorities are allowed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetAllowsAnyRoot
func SSLSetAllowsAnyRoot(context unsafe.Pointer, anyRoot unsafe.Pointer) unsafe.Pointer {
	return _SSLSetAllowsAnyRoot(context, anyRoot)
	}


// Specifies whether certificate expiration times are ignored. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetAllowsExpiredCerts
func SSLSetAllowsExpiredCerts(context unsafe.Pointer, allowsExpired unsafe.Pointer) unsafe.Pointer {
	return _SSLSetAllowsExpiredCerts(context, allowsExpired)
	}


// Specifies whether expired root certificates are allowed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetAllowsExpiredRoots
func SSLSetAllowsExpiredRoots(context unsafe.Pointer, allowsExpired unsafe.Pointer) unsafe.Pointer {
	return _SSLSetAllowsExpiredRoots(context, allowsExpired)
	}


// Specifies this connection’s certificate or certificates. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetCertificate(_:_:)
func SSLSetCertificate(context unsafe.Pointer, certRefs unsafe.Pointer) unsafe.Pointer {
	return _SSLSetCertificate(context, certRefs)
	}


// Adds one or more certificates to a server’s list of certification authorities (CAs) acceptable for client authentication. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetCertificateAuthorities(_:_:_:)
func SSLSetCertificateAuthorities(context unsafe.Pointer, certificateOrArray unsafe.Pointer, replaceExisting unsafe.Pointer) unsafe.Pointer {
	return _SSLSetCertificateAuthorities(context, certificateOrArray, replaceExisting)
	}


// Specifies the requirements for client-side authentication. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetClientSideAuthenticate(_:_:)
func SSLSetClientSideAuthenticate(context unsafe.Pointer, auth unsafe.Pointer) unsafe.Pointer {
	return _SSLSetClientSideAuthenticate(context, auth)
	}


// Specifies an I/O connection for a specific session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetConnection(_:_:)
func SSLSetConnection(context unsafe.Pointer, connection unsafe.Pointer) unsafe.Pointer {
	return _SSLSetConnection(context, connection)
	}


// Sets the cookie value used in the Datagram Transport Layer Security (DTLS) hello message. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetDatagramHelloCookie(_:_:_:)
func SSLSetDatagramHelloCookie(dtlsContext unsafe.Pointer, cookie unsafe.Pointer, cookieLen unsafe.Pointer) unsafe.Pointer {
	return _SSLSetDatagramHelloCookie(dtlsContext, cookie, cookieLen)
	}


// Specifies Diffie-Hellman parameters for a given context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetDiffieHellmanParams(_:_:_:)
func SSLSetDiffieHellmanParams(context unsafe.Pointer, dhParams unsafe.Pointer, dhParamsLen unsafe.Pointer) unsafe.Pointer {
	return _SSLSetDiffieHellmanParams(context, dhParams, dhParamsLen)
	}


// Enables or disables peer certificate chain validation. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetEnableCertVerify
func SSLSetEnableCertVerify(context unsafe.Pointer, enableVerify unsafe.Pointer) unsafe.Pointer {
	return _SSLSetEnableCertVerify(context, enableVerify)
	}


// Specifies a restricted set of SSL cipher suites to be enabled by the current SSL session context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetEnabledCiphers(_:_:_:)
func SSLSetEnabledCiphers(context unsafe.Pointer, ciphers unsafe.Pointer, numCiphers unsafe.Pointer) unsafe.Pointer {
	return _SSLSetEnabledCiphers(context, ciphers, numCiphers)
	}


// Specifies the encryption certificates used for this connection. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetEncryptionCertificate(_:_:)
func SSLSetEncryptionCertificate(context unsafe.Pointer, certRefs unsafe.Pointer) unsafe.Pointer {
	return _SSLSetEncryptionCertificate(context, certRefs)
	}


// Sets the status of a session context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetError(_:_:)
func SSLSetError(context unsafe.Pointer, status unsafe.Pointer) unsafe.Pointer {
	return _SSLSetError(context, status)
	}


// Specifies callback functions that perform the network I/O operations. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetIOFuncs(_:_:_:)
func SSLSetIOFuncs(context unsafe.Pointer, readFunc unsafe.Pointer, writeFunc unsafe.Pointer) unsafe.Pointer {
	return _SSLSetIOFuncs(context, readFunc, writeFunc)
	}


// Sets the maximum datagram record size allowed by the application for a given context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetMaxDatagramRecordSize(_:_:)
func SSLSetMaxDatagramRecordSize(dtlsContext unsafe.Pointer, maxSize unsafe.Pointer) unsafe.Pointer {
	return _SSLSetMaxDatagramRecordSize(dtlsContext, maxSize)
	}


// Sets the OCSP response for the given SSL session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetOCSPResponse(_:_:)
func SSLSetOCSPResponse(context unsafe.Pointer, response unsafe.Pointer) unsafe.Pointer {
	return _SSLSetOCSPResponse(context, response)
	}


// Specifies the fully qualified domain name of the peer. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetPeerDomainName(_:_:_:)
func SSLSetPeerDomainName(context unsafe.Pointer, peerName unsafe.Pointer, peerNameLen unsafe.Pointer) unsafe.Pointer {
	return _SSLSetPeerDomainName(context, peerName, peerNameLen)
	}


// Specifies data that is sufficient to uniquely identify the peer of the current session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetPeerID(_:_:_:)
func SSLSetPeerID(context unsafe.Pointer, peerID unsafe.Pointer, peerIDLen unsafe.Pointer) unsafe.Pointer {
	return _SSLSetPeerID(context, peerID, peerIDLen)
	}


// Sets the SSL protocol version. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersion
func SSLSetProtocolVersion(context unsafe.Pointer, version unsafe.Pointer) unsafe.Pointer {
	return _SSLSetProtocolVersion(context, version)
	}


// Sets the allowed Secure Sockets Layer (SSL) protocol versions. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersionEnabled
func SSLSetProtocolVersionEnabled(context unsafe.Pointer, protocol_ unsafe.Pointer, enable unsafe.Pointer) unsafe.Pointer {
	return _SSLSetProtocolVersionEnabled(context, protocol_, enable)
	}


// Sets the maximum protocol version allowed by the application for a given SSL context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersionMax(_:_:)
func SSLSetProtocolVersionMax(context unsafe.Pointer, maxVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLSetProtocolVersionMax(context, maxVersion)
	}


// Sets the minimum protocol version allowed by the application for a given SSL context. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetProtocolVersionMin(_:_:)
func SSLSetProtocolVersionMin(context unsafe.Pointer, minVersion unsafe.Pointer) unsafe.Pointer {
	return _SSLSetProtocolVersionMin(context, minVersion)
	}


// Enables or disables RSA blinding. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetRsaBlinding
func SSLSetRsaBlinding(context unsafe.Pointer, blinding unsafe.Pointer) unsafe.Pointer {
	return _SSLSetRsaBlinding(context, blinding)
	}


// Sets a predefined configuration for the Secure Sockets Layer (SSL) session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionConfig(_:_:)
func SSLSetSessionConfig(context unsafe.Pointer, config unsafe.Pointer) unsafe.Pointer {
	return _SSLSetSessionConfig(context, config)
	}


// Specifies options for a specific session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionOption(_:_:_:)
func SSLSetSessionOption(context unsafe.Pointer, option unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SSLSetSessionOption(context, option, value)
	}


// Enables or disables session ticket resumption. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetSessionTicketsEnabled(_:_:)
func SSLSetSessionTicketsEnabled(context unsafe.Pointer, enabled unsafe.Pointer) unsafe.Pointer {
	return _SSLSetSessionTicketsEnabled(context, enabled)
	}


// Augments or replaces the default set of trusted root certificates for this session. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.9.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLSetTrustedRoots
func SSLSetTrustedRoots(context unsafe.Pointer, trustedRoots unsafe.Pointer, replaceExisting unsafe.Pointer) unsafe.Pointer {
	return _SSLSetTrustedRoots(context, trustedRoots, replaceExisting)
	}


// Performs a typical application-level write operation. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLWrite(_:_:_:_:)
func SSLWrite(context unsafe.Pointer, data unsafe.Pointer, dataLength unsafe.Pointer, processed unsafe.Pointer) unsafe.Pointer {
	return _SSLWrite(context, data, dataLength, processed)
	}


// Creates a new ACL entry with the given characteristics, and adds it to an access instance. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLCreateWithSimpleContents(_:_:_:_:_:)
func SecACLCreateWithSimpleContents(access unsafe.Pointer, applicationList unsafe.Pointer, description unsafe.Pointer, promptSelector unsafe.Pointer, newAcl unsafe.Pointer) unsafe.Pointer {
	return _SecACLCreateWithSimpleContents(access, applicationList, description, promptSelector, newAcl)
	}


// Sets the application list, description, and prompt selector for a given ACL entry. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLSetContents(_:_:_:_:)
func SecACLSetContents(acl unsafe.Pointer, applicationList unsafe.Pointer, description unsafe.Pointer, promptSelector unsafe.Pointer) unsafe.Pointer {
	return _SecACLSetContents(acl, applicationList, description, promptSelector)
	}


// Sets the authorization tags for a given ACL. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACLUpdateAuthorizations(_:_:)
func SecACLUpdateAuthorizations(acl unsafe.Pointer, authorizations unsafe.Pointer) unsafe.Pointer {
	return _SecACLUpdateAuthorizations(acl, authorizations)
	}


// Creates a new access control object with the specified protection type and flags. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControlCreateWithFlags(_:_:_:_:)
func SecAccessControlCreateWithFlags(allocator unsafe.Pointer, protection unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecAccessControlCreateWithFlags(allocator, protection, flags, error_)
	}


// Retrieves all the ACL entries of a given access instance. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopyACLList(_:_:)
func SecAccessCopyACLList(accessRef unsafe.Pointer, aclList unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopyACLList(accessRef, aclList)
	}


// Retrieves selected ACL entries from a given access instance. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopyMatchingACLList(_:_:)
func SecAccessCopyMatchingACLList(accessRef unsafe.Pointer, authorizationTag unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopyMatchingACLList(accessRef, authorizationTag)
	}


// Retrieves selected access control lists from a given access object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCopySelectedACLList
func SecAccessCopySelectedACLList(accessRef unsafe.Pointer, action unsafe.Pointer, aclList unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCopySelectedACLList(accessRef, action, aclList)
	}


// Creates a new access instance associated with a given protected keychain item. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessCreate(_:_:_:)
func SecAccessCreate(descriptor unsafe.Pointer, trustedlist unsafe.Pointer, accessRef unsafe.Pointer) unsafe.Pointer {
	return _SecAccessCreate(descriptor, trustedlist, accessRef)
	}


// Retrieves the owner and the access control list of a given access object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.7.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessGetOwnerAndACL
func SecAccessGetOwnerAndACL(accessRef unsafe.Pointer, owner unsafe.Pointer, aclCount unsafe.Pointer, acls unsafe.Pointer) unsafe.Pointer {
	return _SecAccessGetOwnerAndACL(accessRef, owner, aclCount, acls)
	}


// Returns the unique identifier of the opaque type to which an access instance belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessGetTypeID()
func SecAccessGetTypeID() unsafe.Pointer {
	return _SecAccessGetTypeID()
	}


// Asynchronously stores (or updates) a shared password for a website. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAddSharedWebCredential(_:_:_:_:)
func SecAddSharedWebCredential(fqdn unsafe.Pointer, account unsafe.Pointer, password unsafe.Pointer) {
	_SecAddSharedWebCredential(fqdn, account, password)
	}


// Allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AllocCopy
func SecAsn1AllocCopy(coder unsafe.Pointer, src unsafe.Pointer, len_ unsafe.Pointer, dest unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1AllocCopy(coder, src, len_, dest)
	}


// Allocates memory for an item’s data field in the coder object’s memory pool and copies in a block of data from another item. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AllocCopyItem
func SecAsn1AllocCopyItem(coder unsafe.Pointer, src unsafe.Pointer, dest unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1AllocCopyItem(coder, src, dest)
	}


// Allocates memory for an item’s data field in the coder object’s memory pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AllocItem
func SecAsn1AllocItem(coder unsafe.Pointer, item unsafe.Pointer, len_ unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1AllocItem(coder, item, len_)
	}


// Creates an ASN.1 coder object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1CoderCreate
func SecAsn1CoderCreate(coder unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1CoderCreate(coder)
	}


// Destroys an ASN.1 coder object and releases all of its memory. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1CoderRelease
func SecAsn1CoderRelease(coder unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1CoderRelease(coder)
	}


// Decodes untyped DER data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Decode
func SecAsn1Decode(coder unsafe.Pointer, src unsafe.Pointer, len_ unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1Decode(coder, src, len_, templates, dest)
	}


// Decodes an ASN.1 item in DER format. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1DecodeData
func SecAsn1DecodeData(coder unsafe.Pointer, src unsafe.Pointer, templ unsafe.Pointer, dest unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1DecodeData(coder, src, templ, dest)
	}


// Encodes data in DER format. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1EncodeItem
func SecAsn1EncodeItem(coder unsafe.Pointer, src unsafe.Pointer, templates unsafe.Pointer, dest unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1EncodeItem(coder, src, templates, dest)
	}


// Allocates memory in the coder object’s memory pool. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Malloc
func SecAsn1Malloc(coder unsafe.Pointer, len_ unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1Malloc(coder, len_)
	}


// Compares two decoded object identifiers. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1OidCompare
func SecAsn1OidCompare(oid1 unsafe.Pointer, oid2 unsafe.Pointer) unsafe.Pointer {
	return _SecAsn1OidCompare(oid1, oid2)
	}


// Returns a DER representation of a certificate given a certificate object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyData(_:)
func SecCertificateCopyData(certificate unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyData(certificate)
	}


// SecCertificateCopyNotValidAfterDate is a Security function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidAfterDate(_:)
func SecCertificateCopyNotValidAfterDate(certificate unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyNotValidAfterDate(certificate)
	}


// SecCertificateCopyNotValidBeforeDate is a Security function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCopyNotValidBeforeDate(_:)
func SecCertificateCopyNotValidBeforeDate(certificate unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCopyNotValidBeforeDate(certificate)
	}


// Creates a certificate object from a DER representation of a certificate. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificateCreateWithData(_:_:)
func SecCertificateCreateWithData(allocator unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecCertificateCreateWithData(allocator, data)
	}


// Performs dynamic validation of signed code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCheckValidity(_:_:_:)
func SecCodeCheckValidity(code unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCheckValidity(code, flags, requirement)
	}


// Performs dynamic validation of signed code and returns detailed error information in the case of failure. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCheckValidityWithErrors(_:_:_:_:)
func SecCodeCheckValidityWithErrors(code unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer, errors unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCheckValidityWithErrors(code, flags, requirement, errors)
	}


// Retrieves the designated code requirement of signed code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyDesignatedRequirement(_:_:_:)
func SecCodeCopyDesignatedRequirement(code unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyDesignatedRequirement(code, flags, requirement)
	}


// Asks a code host to identify one of its guests given the type and value of specific attributes of the guest code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyGuestWithAttributes(_:_:_:_:)
func SecCodeCopyGuestWithAttributes(host unsafe.Pointer, attributes unsafe.Pointer, flags unsafe.Pointer, guest unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyGuestWithAttributes(host, attributes, flags, guest)
	}


// Retrieves the code object for the host of specified guest code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyHost(_:_:_:)
func SecCodeCopyHost(guest unsafe.Pointer, flags unsafe.Pointer, host unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyHost(guest, flags, host)
	}


// Retrieves the location on disk of signed code, given a code or static code object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyPath(_:_:_:)
func SecCodeCopyPath(staticCode unsafe.Pointer, flags unsafe.Pointer, path unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyPath(staticCode, flags, path)
	}


// Retrieves the code object for the code making the call. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopySelf(_:_:)
func SecCodeCopySelf(flags unsafe.Pointer, self unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopySelf(flags, self)
	}


// Retrieves various pieces of information from a code signature. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopySigningInformation(_:_:_:)
func SecCodeCopySigningInformation(code unsafe.Pointer, flags unsafe.Pointer, information unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopySigningInformation(code, flags, information)
	}


// Returns a static code object representing the on-disk version of the given running code. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCopyStaticCode(_:_:_:)
func SecCodeCopyStaticCode(code unsafe.Pointer, flags unsafe.Pointer, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCopyStaticCode(code, flags, staticCode)
	}


// SecCodeCreateWithXPCMessage is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeCreateWithXPCMessage(_:_:_:)
func SecCodeCreateWithXPCMessage(message unsafe.Pointer, flags unsafe.Pointer, target unsafe.Pointer) unsafe.Pointer {
	return _SecCodeCreateWithXPCMessage(message, flags, target)
	}


// Returns the unique identifier of the opaque type to which a code object belongs. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeGetTypeID()
func SecCodeGetTypeID() unsafe.Pointer {
	return _SecCodeGetTypeID()
	}


// Asks the kernel to accept the signing information currently attached to a code object and uses it to validate memory page-ins. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeMapMemory(_:_:)
func SecCodeMapMemory(code unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecCodeMapMemory(code, flags)
	}


// SecCodeValidateFileResource is a Security function. [Full Topic]
//
// Added in macOS 10.13.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCodeValidateFileResource(_:_:_:_:)
func SecCodeValidateFileResource(code unsafe.Pointer, relativePath unsafe.Pointer, fileData unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecCodeValidateFileResource(code, relativePath, fileData, flags)
	}


// Returns a string explaining the meaning of a security result code. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCopyErrorMessageString(_:_:)
func SecCopyErrorMessageString(status unsafe.Pointer, reserved unsafe.Pointer) unsafe.Pointer {
	return _SecCopyErrorMessageString(status, reserved)
	}


// Returns a randomly generated password. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCreateSharedWebCredentialPassword()
func SecCreateSharedWebCredentialPassword() unsafe.Pointer {
	return _SecCreateSharedWebCredentialPassword()
	}


// Creates a decode transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecodeTransformCreate(_:_:)
func SecDecodeTransformCreate(DecodeType unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecDecodeTransformCreate(DecodeType, error_)
	}


// Creates a decryption transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecryptTransformCreate(_:_:)
func SecDecryptTransformCreate(keyRef unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecDecryptTransformCreate(keyRef, error_)
	}


// Returns the unique identifier of the opaque type to which a decryption transform belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDecryptTransformGetTypeID()
func SecDecryptTransformGetTypeID() unsafe.Pointer {
	return _SecDecryptTransformGetTypeID()
	}


// Creates a digest transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDigestTransformCreate(_:_:_:)
func SecDigestTransformCreate(digestType unsafe.Pointer, digestLength unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecDigestTransformCreate(digestType, digestLength, error_)
	}


// Returns the unique identifier of the opaque type to which a digest transform belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecDigestTransformGetTypeID()
func SecDigestTransformGetTypeID() unsafe.Pointer {
	return _SecDigestTransformGetTypeID()
	}


// Creates an encode transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncodeTransformCreate(_:_:)
func SecEncodeTransformCreate(encodeType unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecEncodeTransformCreate(encodeType, error_)
	}


// Creates an encryption transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncryptTransformCreate(_:_:)
func SecEncryptTransformCreate(keyRef unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecEncryptTransformCreate(keyRef, error_)
	}


// Returns the unique identifier of the opaque type to which an encryption transform belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecEncryptTransformGetTypeID()
func SecEncryptTransformGetTypeID() unsafe.Pointer {
	return _SecEncryptTransformGetTypeID()
	}


// Returns the Core Foundation type ID for a transform group container. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecGroupTransformGetTypeID()
func SecGroupTransformGetTypeID() unsafe.Pointer {
	return _SecGroupTransformGetTypeID()
	}


// Creates a new guest and describes its initial properties. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostCreateGuest
func SecHostCreateGuest(host unsafe.Pointer, status uint32, path unsafe.Pointer, attributes unsafe.Pointer, flags unsafe.Pointer, newGuest unsafe.Pointer) unsafe.Pointer {
	return _SecHostCreateGuest(host, status, path, attributes, flags, newGuest)
	}


// Removes a guest from a host. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostRemoveGuest
func SecHostRemoveGuest(host unsafe.Pointer, guest unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecHostRemoveGuest(host, guest, flags)
	}


// Makes the calling thread the proxy for a specified guest. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostSelectGuest
func SecHostSelectGuest(guestRef unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecHostSelectGuest(guestRef, flags)
	}


// Retrieves the handle for the guest currently selected for the calling thread. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostSelectedGuest
func SecHostSelectedGuest(flags unsafe.Pointer, guestRef unsafe.Pointer) unsafe.Pointer {
	return _SecHostSelectedGuest(flags, guestRef)
	}


// Updates the status and attributes of a particular guest. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostSetGuestStatus
func SecHostSetGuestStatus(guestRef unsafe.Pointer, status uint32, attributes unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecHostSetGuestStatus(guestRef, status, attributes, flags)
	}


// Tells code signing services that the calling code will directly respond to hosting inquiries over the given port. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.6.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecHostSetHostingPort
func SecHostSetHostingPort(hostingPort unsafe.Pointer, flags unsafe.Pointer) unsafe.Pointer {
	return _SecHostSetHostingPort(hostingPort, flags)
	}


// SecIdentityCreate is a Security function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentityCreate(_:_:_:)
func SecIdentityCreate(allocator unsafe.Pointer, certificate unsafe.Pointer, privateKey unsafe.Pointer) unsafe.Pointer {
	return _SecIdentityCreate(allocator, certificate, privateKey)
	}


// Adds one or more items to a keychain. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemAdd(_:_:)
func SecItemAdd(attributes unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecItemAdd(attributes, result)
	}


// Returns one or more keychain items that match a search query, or copies attributes of specific keychain items. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemCopyMatching(_:_:)
func SecItemCopyMatching(query unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecItemCopyMatching(query, result)
	}


// Deletes items that match a search query. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemDelete(_:)
func SecItemDelete(query unsafe.Pointer) unsafe.Pointer {
	return _SecItemDelete(query)
	}


// Modifies items that match a search query. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemUpdate(_:_:)
func SecItemUpdate(query unsafe.Pointer, attributesToUpdate unsafe.Pointer) unsafe.Pointer {
	return _SecItemUpdate(query, attributesToUpdate)
	}


// Returns an external representation of the given key suitable for the key’s type. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyExternalRepresentation(_:_:)
func SecKeyCopyExternalRepresentation(key unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCopyExternalRepresentation(key, error_)
	}


// Gets the public key associated with the given private key. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCopyPublicKey(_:)
func SecKeyCopyPublicKey(key unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCopyPublicKey(key)
	}


// Decrypts a block of data using a private key and specified algorithm. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateDecryptedData(_:_:_:_:)
func SecKeyCreateDecryptedData(key unsafe.Pointer, algorithm unsafe.Pointer, ciphertext unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCreateDecryptedData(key, algorithm, ciphertext, error_)
	}


// Encrypts a block of data using a public key and specified algorithm. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateEncryptedData(_:_:_:_:)
func SecKeyCreateEncryptedData(key unsafe.Pointer, algorithm unsafe.Pointer, plaintext unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCreateEncryptedData(key, algorithm, plaintext, error_)
	}


// Generates a new public-private key pair. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateRandomKey(_:_:)
func SecKeyCreateRandomKey(parameters unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCreateRandomKey(parameters, error_)
	}


// Restores a key from an external representation of that key. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyCreateWithData(_:_:_:)
func SecKeyCreateWithData(keyData unsafe.Pointer, attributes unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyCreateWithData(keyData, attributes, error_)
	}


// Returns a Boolean indicating whether a key is suitable for an operation using a certain algorithm. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyIsAlgorithmSupported(_:_:_:)
func SecKeyIsAlgorithmSupported(key unsafe.Pointer, operation unsafe.Pointer, algorithm unsafe.Pointer) unsafe.Pointer {
	return _SecKeyIsAlgorithmSupported(key, operation, algorithm)
	}


// Unwraps a wrapped symmetric key. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyUnwrapSymmetric(_:_:_:_:)
func SecKeyUnwrapSymmetric(keyToUnwrap unsafe.Pointer, unwrappingKey unsafe.Pointer, parameters unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecKeyUnwrapSymmetric(keyToUnwrap, unwrappingKey, parameters, error_)
	}


// Retrieves the access of a given keychain item. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItemCopyAccess(_:_:)
func SecKeychainItemCopyAccess(itemRef unsafe.Pointer, access unsafe.Pointer) unsafe.Pointer {
	return _SecKeychainItemCopyAccess(itemRef, access)
	}


// Returns a policy object for the default X.509 policy. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateBasicX509()
func SecPolicyCreateBasicX509() unsafe.Pointer {
	return _SecPolicyCreateBasicX509()
	}


// Returns a policy object for evaluating SSL certificate chains. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateSSL(_:_:)
func SecPolicyCreateSSL(server unsafe.Pointer, hostname unsafe.Pointer) unsafe.Pointer {
	return _SecPolicyCreateSSL(server, hostname)
	}


// Returns a policy object based on an object identifier for the policy type. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicyCreateWithProperties(_:_:)
func SecPolicyCreateWithProperties(policyIdentifier unsafe.Pointer, properties unsafe.Pointer) unsafe.Pointer {
	return _SecPolicyCreateWithProperties(policyIdentifier, properties)
	}


// Generates an array of cryptographically secure random bytes. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRandomCopyBytes(_:_:_:)
func SecRandomCopyBytes(rnd unsafe.Pointer, count unsafe.Pointer, bytes unsafe.Pointer) unsafe.Pointer {
	return _SecRandomCopyBytes(rnd, count, bytes)
	}


// Asynchronously obtains one or more shared passwords for a website. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequestSharedWebCredential(_:_:_:)
func SecRequestSharedWebCredential(fqdn unsafe.Pointer, account unsafe.Pointer) {
	_SecRequestSharedWebCredential(fqdn, account)
	}


// Extracts a binary form of a code requirement from a code requirement object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCopyData(_:_:_:)
func SecRequirementCopyData(requirement unsafe.Pointer, flags unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCopyData(requirement, flags, data)
	}


// Converts a code requirement object into text form. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCopyString(_:_:_:)
func SecRequirementCopyString(requirement unsafe.Pointer, flags unsafe.Pointer, text unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCopyString(requirement, flags, text)
	}


// Creates a code requirement object from the binary form of a code requirement. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithData(_:_:_:)
func SecRequirementCreateWithData(data unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithData(data, flags, requirement)
	}


// Creates a code requirement object by compiling a valid text representation of a code requirement. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithString(_:_:_:)
func SecRequirementCreateWithString(text unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithString(text, flags, requirement)
	}


// Creates a code requirement object by compiling a valid text representation of a code requirement and returns detailed error information in the case of failure. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementCreateWithStringAndErrors(_:_:_:_:)
func SecRequirementCreateWithStringAndErrors(text unsafe.Pointer, flags unsafe.Pointer, errors unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecRequirementCreateWithStringAndErrors(text, flags, errors, requirement)
	}


// Returns the unique identifier of the opaque type to which a code requirement object belongs. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirementGetTypeID()
func SecRequirementGetTypeID() unsafe.Pointer {
	return _SecRequirementGetTypeID()
	}


// Creates a signing transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecSignTransformCreate(_:_:)
func SecSignTransformCreate(key unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecSignTransformCreate(key, error_)
	}


// Validates a static code object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidity(_:_:_:)
func SecStaticCodeCheckValidity(staticCode unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCheckValidity(staticCode, flags, requirement)
	}


// Performs static validation of static signed code and returns detailed error information in the case of failure. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCheckValidityWithErrors(_:_:_:_:)
func SecStaticCodeCheckValidityWithErrors(staticCode unsafe.Pointer, flags unsafe.Pointer, requirement unsafe.Pointer, errors unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCheckValidityWithErrors(staticCode, flags, requirement, errors)
	}


// Creates a static code object representing the code at a specified file system path. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPath(_:_:_:)
func SecStaticCodeCreateWithPath(path unsafe.Pointer, flags unsafe.Pointer, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCreateWithPath(path, flags, staticCode)
	}


// Creates a static code object representing the code at a specified file system path using an attributes dictionary. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeCreateWithPathAndAttributes(_:_:_:_:)
func SecStaticCodeCreateWithPathAndAttributes(path unsafe.Pointer, flags unsafe.Pointer, attributes unsafe.Pointer, staticCode unsafe.Pointer) unsafe.Pointer {
	return _SecStaticCodeCreateWithPathAndAttributes(path, flags, attributes, staticCode)
	}


// Returns the unique identifier of the opaque type to which a static code object belongs. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCodeGetTypeID()
func SecStaticCodeGetTypeID() unsafe.Pointer {
	return _SecStaticCodeGetTypeID()
	}


// Returns the value of the code signing identifier. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopySigningIdentifier(_:_:)
func SecTaskCopySigningIdentifier(task unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTaskCopySigningIdentifier(task, error_)
	}


// Returns the value of a single entitlement for the represented task. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopyValueForEntitlement(_:_:_:)
func SecTaskCopyValueForEntitlement(task unsafe.Pointer, entitlement unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTaskCopyValueForEntitlement(task, entitlement, error_)
	}


// Returns the values of multiple entitlements for the represented task. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCopyValuesForEntitlements(_:_:_:)
func SecTaskCopyValuesForEntitlements(task unsafe.Pointer, entitlements unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTaskCopyValuesForEntitlements(task, entitlements, error_)
	}


// Creates a task object for the current task. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCreateFromSelf(_:)
func SecTaskCreateFromSelf(allocator unsafe.Pointer) unsafe.Pointer {
	return _SecTaskCreateFromSelf(allocator)
	}


// Creates a task object for the task that sent the Mach message represented by the audit token. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskCreateWithAuditToken(_:_:)
func SecTaskCreateWithAuditToken(allocator unsafe.Pointer, token unsafe.Pointer) unsafe.Pointer {
	return _SecTaskCreateWithAuditToken(allocator, token)
	}


// SecTaskGetCodeSignStatus is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskGetCodeSignStatus(_:)
func SecTaskGetCodeSignStatus(task unsafe.Pointer) uint32 {
	return _SecTaskGetCodeSignStatus(task)
	}


// Returns the unique identifier of the opaque type to which a task object belongs. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTaskGetTypeID()
func SecTaskGetTypeID() unsafe.Pointer {
	return _SecTaskGetTypeID()
	}


// Gets an attribute value from a custom transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTranformCustomGetAttribute
func SecTranformCustomGetAttribute(ref unsafe.Pointer, attribute unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _SecTranformCustomGetAttribute(ref, attribute, type_)
	}


// Chains transforms together. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformConnectTransforms(_:_:_:_:_:_:)
func SecTransformConnectTransforms(sourceTransformRef unsafe.Pointer, sourceAttributeName unsafe.Pointer, destinationTransformRef unsafe.Pointer, destinationAttributeName unsafe.Pointer, group unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformConnectTransforms(sourceTransformRef, sourceAttributeName, destinationTransformRef, destinationAttributeName, group, error_)
	}


// Creates a dictionary that contains enough information to be able to recreate a transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCopyExternalRepresentation(_:)
func SecTransformCopyExternalRepresentation(transformRef unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCopyExternalRepresentation(transformRef)
	}


// Creates a transform computation object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreate(_:_:)
func SecTransformCreate(name unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCreate(name, error_)
	}


// Creates a transform instance from a dictionary of parameters. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateFromExternalRepresentation(_:_:)
func SecTransformCreateFromExternalRepresentation(dictionary unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCreateFromExternalRepresentation(dictionary, error_)
	}


// Creates an object that acts as a container for a set of connected transforms. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateGroupTransform()
func SecTransformCreateGroupTransform() unsafe.Pointer {
	return _SecTransformCreateGroupTransform()
	}


// Creates a read transform from a read stream reference. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateReadTransformWithReadStream(_:)
func SecTransformCreateReadTransformWithReadStream(inputStream unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCreateReadTransformWithReadStream(inputStream)
	}


// Gets an attribute value from a custom transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCustomGetAttribute(_:_:_:)
func SecTransformCustomGetAttribute(ref unsafe.Pointer, attribute unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCustomGetAttribute(ref, attribute, type_)
	}


// Sets an attribute value on a custom transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCustomSetAttribute(_:_:_:_:)
func SecTransformCustomSetAttribute(ref unsafe.Pointer, attribute unsafe.Pointer, type_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SecTransformCustomSetAttribute(ref, attribute, type_, value)
	}


// Executes a transform or transform group synchronously. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformExecute(_:_:)
func SecTransformExecute(transformRef unsafe.Pointer, errorRef unsafe.Pointer) unsafe.Pointer {
	return _SecTransformExecute(transformRef, errorRef)
	}


// Executes transform or transform group asynchronously. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformExecuteAsync(_:_:_:)
func SecTransformExecuteAsync(transformRef unsafe.Pointer, deliveryQueue unsafe.Pointer, deliveryBlock unsafe.Pointer) {
	_SecTransformExecuteAsync(transformRef, deliveryQueue, deliveryBlock)
	}


// Finds a member of a transform group by its name. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformFindByName(_:_:)
func SecTransformFindByName(transform unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _SecTransformFindByName(transform, name)
	}


// Gets the current value of a transform attribute. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformGetAttribute(_:_:)
func SecTransformGetAttribute(transformRef unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _SecTransformGetAttribute(transformRef, key)
	}


// Returns the unique identifier of the opaque type to which a security transform object belongs. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformGetTypeID()
func SecTransformGetTypeID() unsafe.Pointer {
	return _SecTransformGetTypeID()
	}


// Returns an object from inside a ProcessData override that says that although no data is being returned the transform is still active and awaiting data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformNoData()
func SecTransformNoData() unsafe.Pointer {
	return _SecTransformNoData()
	}


// Pushes a single value back for a specific attribute. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformPushbackAttribute(_:_:_:)
func SecTransformPushbackAttribute(ref unsafe.Pointer, attribute unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _SecTransformPushbackAttribute(ref, attribute, value)
	}


// Registers a custom transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformRegister(_:_:_:)
func SecTransformRegister(uniqueName unsafe.Pointer, createTransformFunction unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformRegister(uniqueName, createTransformFunction, error_)
	}


// Sets a static value for an attribute in a transform. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetAttribute(_:_:_:_:)
func SecTransformSetAttribute(transformRef unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTransformSetAttribute(transformRef, key, value, error_)
	}


// Requests a callback when an attribute is set. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetAttributeAction(_:_:_:_:)
func SecTransformSetAttributeAction(ref unsafe.Pointer, action unsafe.Pointer, attribute unsafe.Pointer, newAction unsafe.Pointer) unsafe.Pointer {
	return _SecTransformSetAttributeAction(ref, action, attribute, newAction)
	}


// Changes the way a custom transform processes data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetDataAction(_:_:_:)
func SecTransformSetDataAction(ref unsafe.Pointer, action unsafe.Pointer, newAction unsafe.Pointer) unsafe.Pointer {
	return _SecTransformSetDataAction(ref, action, newAction)
	}


// Changes the way that a transform deals with transform lifecycle behaviors. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformSetTransformAction(_:_:_:)
func SecTransformSetTransformAction(ref unsafe.Pointer, action unsafe.Pointer, newAction unsafe.Pointer) unsafe.Pointer {
	return _SecTransformSetTransformAction(ref, action, newAction)
	}


// SecTrustCopyCertificateChain is a Security function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyCertificateChain(_:)
func SecTrustCopyCertificateChain(trust unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCopyCertificateChain(trust)
	}


// SecTrustCopyKey is a Security function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCopyKey(_:)
func SecTrustCopyKey(trust unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCopyKey(trust)
	}


// Creates a trust management object based on certificates and policies. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustCreateWithCertificates(_:_:_:)
func SecTrustCreateWithCertificates(certificates unsafe.Pointer, policies unsafe.Pointer, trust unsafe.Pointer) unsafe.Pointer {
	return _SecTrustCreateWithCertificates(certificates, policies, trust)
	}


// Evaluates trust for the specified certificate and policies. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluate(_:_:)
func SecTrustEvaluate(trust unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluate(trust, result)
	}


// Evaluates trust for the specified certificate and policies. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluateWithError(_:_:)
func SecTrustEvaluateWithError(trust unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluateWithError(trust, error_)
	}


// Returns the result code from the most recent trust evaluation. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustGetTrustResult(_:_:)
func SecTrustGetTrustResult(trust unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustGetTrustResult(trust, result)
	}


// Sets the anchor certificates used when evaluating a trust management object. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetAnchorCertificates(_:_:)
func SecTrustSetAnchorCertificates(trust unsafe.Pointer, anchorCertificates unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetAnchorCertificates(trust, anchorCertificates)
	}


// Sets the date and time against which the certificates in a trust management object are verified. [Full Topic]
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetVerifyDate(_:_:)
func SecTrustSetVerifyDate(trust unsafe.Pointer, verifyDate unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetVerifyDate(trust, verifyDate)
	}


// Creates a verify transform object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 13.0.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecVerifyTransformCreate(_:_:_:)
func SecVerifyTransformCreate(key unsafe.Pointer, signature unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _SecVerifyTransformCreate(key, signature, error_)
	}


// Returns download ticket’s creation date. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCopyCreationDate
func SecureDownloadCopyCreationDate(downloadRef unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCopyCreationDate(downloadRef, date)
	}


// Returns the printable name of the download ticket. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCopyName
func SecureDownloadCopyName(downloadRef unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCopyName(downloadRef, name)
	}


// Copies the ticket location from a secure download URL. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCopyTicketLocation
func SecureDownloadCopyTicketLocation(url unsafe.Pointer, ticketLocation unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCopyTicketLocation(url, ticketLocation)
	}


// Returns a list of URLs from which the data can be downloaded. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCopyURLs
func SecureDownloadCopyURLs(downloadRef unsafe.Pointer, urls unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCopyURLs(downloadRef, urls)
	}


// Creates a secure download object for use during the download process. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadCreateWithTicket
func SecureDownloadCreateWithTicket(ticket unsafe.Pointer, setup unsafe.Pointer, setupContext unsafe.Pointer, evaluate unsafe.Pointer, evaluateContext unsafe.Pointer, downloadRef unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadCreateWithTicket(ticket, setup, setupContext, evaluate, evaluateContext, downloadRef)
	}


// Concludes the secure download process. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadFinished
func SecureDownloadFinished(downloadRef unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadFinished(downloadRef)
	}


// Returns the size of the expected download. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadGetDownloadSize
func SecureDownloadGetDownloadSize(downloadRef unsafe.Pointer, downloadSize unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadGetDownloadSize(downloadRef, downloadSize)
	}


// Releases the memory associated with a secure download object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadRelease
func SecureDownloadRelease(downloadRef unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadRelease(downloadRef)
	}


// Checks data received during download for validity. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecureDownloadUpdateWithData
func SecureDownloadUpdateWithData(downloadRef unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _SecureDownloadUpdateWithData(downloadRef, data)
	}


// Creates a security session. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionCreate(_:_:)
func SessionCreate(flags unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _SessionCreate(flags, attributes)
	}


// Obtains information about a security session. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SessionGetInfo(_:_:_:)
func SessionGetInfo(session unsafe.Pointer, sessionId unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _SessionGetInfo(session, sessionId, attributes)
	}


// cssmAlgToOid is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmAlgToOid(_:)
func cssmAlgToOid(algId unsafe.Pointer) unsafe.Pointer {
	return _cssmAlgToOid(algId)
	}


// cssmOidToAlg is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmOidToAlg(_:_:)
func cssmOidToAlg(oid unsafe.Pointer, alg unsafe.Pointer) unsafe.Pointer {
	return _cssmOidToAlg(oid, alg)
	}


// cssmPerror is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/cssmPerror(_:_:)
func cssmPerror(how unsafe.Pointer, error_ unsafe.Pointer) {
	_cssmPerror(how, error_)
	}


// sec_certificate_copy_ref is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_certificate_copy_ref(_:)
func sec_certificate_copy_ref(certificate unsafe.Pointer) unsafe.Pointer {
	return _sec_certificate_copy_ref(certificate)
	}


// sec_certificate_create is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_certificate_create(_:)
func sec_certificate_create(certificate unsafe.Pointer) unsafe.Pointer {
	return _sec_certificate_create(certificate)
	}


// sec_identity_access_certificates is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_access_certificates(_:_:)
func sec_identity_access_certificates(identity unsafe.Pointer) unsafe.Pointer {
	return _sec_identity_access_certificates(identity)
	}


// sec_identity_copy_certificates_ref is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_copy_certificates_ref(_:)
func sec_identity_copy_certificates_ref(identity unsafe.Pointer) unsafe.Pointer {
	return _sec_identity_copy_certificates_ref(identity)
	}


// sec_identity_copy_ref is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_copy_ref(_:)
func sec_identity_copy_ref(identity unsafe.Pointer) unsafe.Pointer {
	return _sec_identity_copy_ref(identity)
	}


// sec_identity_create is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_create(_:)
func sec_identity_create(identity unsafe.Pointer) unsafe.Pointer {
	return _sec_identity_create(identity)
	}


// sec_identity_create_with_certificates is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_create_with_certificates(_:_:)
func sec_identity_create_with_certificates(identity unsafe.Pointer, certificates unsafe.Pointer) unsafe.Pointer {
	return _sec_identity_create_with_certificates(identity, certificates)
	}


// sec_protocol_metadata_access_distinguished_names is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_distinguished_names(_:_:)
func sec_protocol_metadata_access_distinguished_names(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_access_distinguished_names(metadata)
	}


// sec_protocol_metadata_access_ocsp_response is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_ocsp_response(_:_:)
func sec_protocol_metadata_access_ocsp_response(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_access_ocsp_response(metadata)
	}


// sec_protocol_metadata_access_peer_certificate_chain is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_peer_certificate_chain(_:_:)
func sec_protocol_metadata_access_peer_certificate_chain(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_access_peer_certificate_chain(metadata)
	}


// sec_protocol_metadata_access_pre_shared_keys is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_pre_shared_keys(_:_:)
func sec_protocol_metadata_access_pre_shared_keys(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_access_pre_shared_keys(metadata)
	}


// sec_protocol_metadata_access_supported_signature_algorithms is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_access_supported_signature_algorithms(_:_:)
func sec_protocol_metadata_access_supported_signature_algorithms(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_access_supported_signature_algorithms(metadata)
	}


// sec_protocol_metadata_challenge_parameters_are_equal is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_challenge_parameters_are_equal(_:_:)
func sec_protocol_metadata_challenge_parameters_are_equal(metadataA unsafe.Pointer, metadataB unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_challenge_parameters_are_equal(metadataA, metadataB)
	}


// sec_protocol_metadata_copy_negotiated_protocol is a Security function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_negotiated_protocol(_:)
func sec_protocol_metadata_copy_negotiated_protocol(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_copy_negotiated_protocol(metadata)
	}


// sec_protocol_metadata_copy_peer_public_key is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_peer_public_key(_:)
func sec_protocol_metadata_copy_peer_public_key(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_copy_peer_public_key(metadata)
	}


// sec_protocol_metadata_copy_server_name is a Security function. [Full Topic]
//
// Added in macOS 15.5.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_copy_server_name(_:)
func sec_protocol_metadata_copy_server_name(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_copy_server_name(metadata)
	}


// sec_protocol_metadata_create_secret is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret(_:_:_:_:)
func sec_protocol_metadata_create_secret(metadata unsafe.Pointer, label_len unsafe.Pointer, label unsafe.Pointer, exporter_length unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_create_secret(metadata, label_len, label, exporter_length)
	}


// sec_protocol_metadata_create_secret_with_context is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_create_secret_with_context(_:_:_:_:_:_:)
func sec_protocol_metadata_create_secret_with_context(metadata unsafe.Pointer, label_len unsafe.Pointer, label unsafe.Pointer, context_len unsafe.Pointer, context unsafe.Pointer, exporter_length unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_create_secret_with_context(metadata, label_len, label, context_len, context, exporter_length)
	}


// sec_protocol_metadata_get_early_data_accepted is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_early_data_accepted(_:)
func sec_protocol_metadata_get_early_data_accepted(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_early_data_accepted(metadata)
	}


// sec_protocol_metadata_get_negotiated_ciphersuite is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_ciphersuite(_:)
func sec_protocol_metadata_get_negotiated_ciphersuite(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_negotiated_ciphersuite(metadata)
	}


// sec_protocol_metadata_get_negotiated_protocol is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.5.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol(_:)
func sec_protocol_metadata_get_negotiated_protocol(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_negotiated_protocol(metadata)
	}


// sec_protocol_metadata_get_negotiated_protocol_version is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_protocol_version(_:)
func sec_protocol_metadata_get_negotiated_protocol_version(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_negotiated_protocol_version(metadata)
	}


// sec_protocol_metadata_get_negotiated_tls_ciphersuite is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_ciphersuite(_:)
func sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_negotiated_tls_ciphersuite(metadata)
	}


// sec_protocol_metadata_get_negotiated_tls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_negotiated_tls_protocol_version(_:)
func sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_negotiated_tls_protocol_version(metadata)
	}


// sec_protocol_metadata_get_server_name is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 15.5.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_get_server_name(_:)
func sec_protocol_metadata_get_server_name(metadata unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_get_server_name(metadata)
	}


// sec_protocol_metadata_peers_are_equal is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_peers_are_equal(_:_:)
func sec_protocol_metadata_peers_are_equal(metadataA unsafe.Pointer, metadataB unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_metadata_peers_are_equal(metadataA, metadataB)
	}


// sec_protocol_options_add_pre_shared_key is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_add_pre_shared_key(_:_:_:)
func sec_protocol_options_add_pre_shared_key(options unsafe.Pointer, psk unsafe.Pointer, psk_identity unsafe.Pointer) {
	_sec_protocol_options_add_pre_shared_key(options, psk, psk_identity)
	}


// sec_protocol_options_add_tls_application_protocol is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_application_protocol(_:_:)
func sec_protocol_options_add_tls_application_protocol(options unsafe.Pointer, application_protocol unsafe.Pointer) {
	_sec_protocol_options_add_tls_application_protocol(options, application_protocol)
	}


// sec_protocol_options_add_tls_ciphersuite is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_ciphersuite(_:_:)
func sec_protocol_options_add_tls_ciphersuite(options unsafe.Pointer, ciphersuite unsafe.Pointer) {
	_sec_protocol_options_add_tls_ciphersuite(options, ciphersuite)
	}


// sec_protocol_options_add_tls_ciphersuite_group is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_add_tls_ciphersuite_group(_:_:)
func sec_protocol_options_add_tls_ciphersuite_group(options unsafe.Pointer, group unsafe.Pointer) {
	_sec_protocol_options_add_tls_ciphersuite_group(options, group)
	}


// sec_protocol_options_append_tls_ciphersuite is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite(_:_:)
func sec_protocol_options_append_tls_ciphersuite(options unsafe.Pointer, ciphersuite unsafe.Pointer) {
	_sec_protocol_options_append_tls_ciphersuite(options, ciphersuite)
	}


// sec_protocol_options_append_tls_ciphersuite_group is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_append_tls_ciphersuite_group(_:_:)
func sec_protocol_options_append_tls_ciphersuite_group(options unsafe.Pointer, group unsafe.Pointer) {
	_sec_protocol_options_append_tls_ciphersuite_group(options, group)
	}


// sec_protocol_options_are_equal is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_are_equal(_:_:)
func sec_protocol_options_are_equal(optionsA unsafe.Pointer, optionsB unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_options_are_equal(optionsA, optionsB)
	}


// sec_protocol_options_get_default_max_dtls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_dtls_protocol_version()
func sec_protocol_options_get_default_max_dtls_protocol_version() unsafe.Pointer {
	return _sec_protocol_options_get_default_max_dtls_protocol_version()
	}


// sec_protocol_options_get_default_max_tls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_max_tls_protocol_version()
func sec_protocol_options_get_default_max_tls_protocol_version() unsafe.Pointer {
	return _sec_protocol_options_get_default_max_tls_protocol_version()
	}


// sec_protocol_options_get_default_min_dtls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_dtls_protocol_version()
func sec_protocol_options_get_default_min_dtls_protocol_version() unsafe.Pointer {
	return _sec_protocol_options_get_default_min_dtls_protocol_version()
	}


// sec_protocol_options_get_default_min_tls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_default_min_tls_protocol_version()
func sec_protocol_options_get_default_min_tls_protocol_version() unsafe.Pointer {
	return _sec_protocol_options_get_default_min_tls_protocol_version()
	}


// sec_protocol_options_get_enable_encrypted_client_hello is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_enable_encrypted_client_hello
func sec_protocol_options_get_enable_encrypted_client_hello(options unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_options_get_enable_encrypted_client_hello(options)
	}


// sec_protocol_options_get_quic_use_legacy_codepoint is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_get_quic_use_legacy_codepoint
func sec_protocol_options_get_quic_use_legacy_codepoint(options unsafe.Pointer) unsafe.Pointer {
	return _sec_protocol_options_get_quic_use_legacy_codepoint(options)
	}


// sec_protocol_options_set_challenge_block is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_challenge_block(_:_:_:)
func sec_protocol_options_set_challenge_block(options unsafe.Pointer, challenge_block unsafe.Pointer, challenge_queue unsafe.Pointer) {
	_sec_protocol_options_set_challenge_block(options, challenge_block, challenge_queue)
	}


// sec_protocol_options_set_enable_encrypted_client_hello is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_enable_encrypted_client_hello
func sec_protocol_options_set_enable_encrypted_client_hello(options unsafe.Pointer, enable_encrypted_client_hello unsafe.Pointer) {
	_sec_protocol_options_set_enable_encrypted_client_hello(options, enable_encrypted_client_hello)
	}


// sec_protocol_options_set_key_update_block is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_key_update_block(_:_:_:)
func sec_protocol_options_set_key_update_block(options unsafe.Pointer, key_update_block unsafe.Pointer, key_update_queue unsafe.Pointer) {
	_sec_protocol_options_set_key_update_block(options, key_update_block, key_update_queue)
	}


// sec_protocol_options_set_local_identity is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_local_identity(_:_:)
func sec_protocol_options_set_local_identity(options unsafe.Pointer, identity unsafe.Pointer) {
	_sec_protocol_options_set_local_identity(options, identity)
	}


// sec_protocol_options_set_max_tls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_max_tls_protocol_version(_:_:)
func sec_protocol_options_set_max_tls_protocol_version(options unsafe.Pointer, version unsafe.Pointer) {
	_sec_protocol_options_set_max_tls_protocol_version(options, version)
	}


// sec_protocol_options_set_min_tls_protocol_version is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_min_tls_protocol_version(_:_:)
func sec_protocol_options_set_min_tls_protocol_version(options unsafe.Pointer, version unsafe.Pointer) {
	_sec_protocol_options_set_min_tls_protocol_version(options, version)
	}


// sec_protocol_options_set_peer_authentication_optional is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_optional
func sec_protocol_options_set_peer_authentication_optional(options unsafe.Pointer, peer_authentication_optional unsafe.Pointer) {
	_sec_protocol_options_set_peer_authentication_optional(options, peer_authentication_optional)
	}


// sec_protocol_options_set_peer_authentication_required is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_peer_authentication_required(_:_:)
func sec_protocol_options_set_peer_authentication_required(options unsafe.Pointer, peer_authentication_required unsafe.Pointer) {
	_sec_protocol_options_set_peer_authentication_required(options, peer_authentication_required)
	}


// sec_protocol_options_set_pre_shared_key_selection_block is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_pre_shared_key_selection_block(_:_:_:)
func sec_protocol_options_set_pre_shared_key_selection_block(options unsafe.Pointer, psk_selection_block unsafe.Pointer, psk_selection_queue unsafe.Pointer) {
	_sec_protocol_options_set_pre_shared_key_selection_block(options, psk_selection_block, psk_selection_queue)
	}


// sec_protocol_options_set_quic_use_legacy_codepoint is a Security function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_quic_use_legacy_codepoint
func sec_protocol_options_set_quic_use_legacy_codepoint(options unsafe.Pointer, quic_use_legacy_codepoint unsafe.Pointer) {
	_sec_protocol_options_set_quic_use_legacy_codepoint(options, quic_use_legacy_codepoint)
	}


// sec_protocol_options_set_tls_diffie_hellman_parameters is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_diffie_hellman_parameters(_:_:)
func sec_protocol_options_set_tls_diffie_hellman_parameters(options unsafe.Pointer, params unsafe.Pointer) {
	_sec_protocol_options_set_tls_diffie_hellman_parameters(options, params)
	}


// sec_protocol_options_set_tls_false_start_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_false_start_enabled(_:_:)
func sec_protocol_options_set_tls_false_start_enabled(options unsafe.Pointer, false_start_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_false_start_enabled(options, false_start_enabled)
	}


// sec_protocol_options_set_tls_is_fallback_attempt is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_is_fallback_attempt(_:_:)
func sec_protocol_options_set_tls_is_fallback_attempt(options unsafe.Pointer, is_fallback_attempt unsafe.Pointer) {
	_sec_protocol_options_set_tls_is_fallback_attempt(options, is_fallback_attempt)
	}


// sec_protocol_options_set_tls_max_version is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_max_version(_:_:)
func sec_protocol_options_set_tls_max_version(options unsafe.Pointer, version unsafe.Pointer) {
	_sec_protocol_options_set_tls_max_version(options, version)
	}


// sec_protocol_options_set_tls_min_version is a Security function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_min_version(_:_:)
func sec_protocol_options_set_tls_min_version(options unsafe.Pointer, version unsafe.Pointer) {
	_sec_protocol_options_set_tls_min_version(options, version)
	}


// sec_protocol_options_set_tls_ocsp_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_ocsp_enabled(_:_:)
func sec_protocol_options_set_tls_ocsp_enabled(options unsafe.Pointer, ocsp_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_ocsp_enabled(options, ocsp_enabled)
	}


// sec_protocol_options_set_tls_pre_shared_key_identity_hint is a Security function. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_pre_shared_key_identity_hint(_:_:)
func sec_protocol_options_set_tls_pre_shared_key_identity_hint(options unsafe.Pointer, psk_identity_hint unsafe.Pointer) {
	_sec_protocol_options_set_tls_pre_shared_key_identity_hint(options, psk_identity_hint)
	}


// sec_protocol_options_set_tls_renegotiation_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_renegotiation_enabled(_:_:)
func sec_protocol_options_set_tls_renegotiation_enabled(options unsafe.Pointer, renegotiation_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_renegotiation_enabled(options, renegotiation_enabled)
	}


// sec_protocol_options_set_tls_resumption_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_resumption_enabled(_:_:)
func sec_protocol_options_set_tls_resumption_enabled(options unsafe.Pointer, resumption_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_resumption_enabled(options, resumption_enabled)
	}


// sec_protocol_options_set_tls_sct_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_sct_enabled(_:_:)
func sec_protocol_options_set_tls_sct_enabled(options unsafe.Pointer, sct_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_sct_enabled(options, sct_enabled)
	}


// sec_protocol_options_set_tls_server_name is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_server_name(_:_:)
func sec_protocol_options_set_tls_server_name(options unsafe.Pointer, server_name unsafe.Pointer) {
	_sec_protocol_options_set_tls_server_name(options, server_name)
	}


// sec_protocol_options_set_tls_tickets_enabled is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_tls_tickets_enabled(_:_:)
func sec_protocol_options_set_tls_tickets_enabled(options unsafe.Pointer, tickets_enabled unsafe.Pointer) {
	_sec_protocol_options_set_tls_tickets_enabled(options, tickets_enabled)
	}


// sec_protocol_options_set_verify_block is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_set_verify_block(_:_:_:)
func sec_protocol_options_set_verify_block(options unsafe.Pointer, verify_block unsafe.Pointer, verify_block_queue unsafe.Pointer) {
	_sec_protocol_options_set_verify_block(options, verify_block, verify_block_queue)
	}


// sec_release is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_release(_:)
func sec_release(obj unsafe.Pointer) {
	_sec_release(obj)
	}


// sec_retain is a Security function. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_retain(_:)
func sec_retain(obj unsafe.Pointer) unsafe.Pointer {
	return _sec_retain(obj)
	}


// sec_trust_copy_ref is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_trust_copy_ref(_:)
func sec_trust_copy_ref(trust unsafe.Pointer) unsafe.Pointer {
	return _sec_trust_copy_ref(trust)
	}


// sec_trust_create is a Security function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_trust_create(_:)
func sec_trust_create(trust unsafe.Pointer) unsafe.Pointer {
	return _sec_trust_create(trust)
	}


// Evaluates a trust object asynchronously on the specified dispatch queue. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustEvaluateAsync(_:_:_:)
func SecTrustEvaluateAsync(trust unsafe.Pointer, queue unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _SecTrustEvaluateAsync(trust, queue, result)
	}


// Sets the keychains searched for intermediate certificates when evaluating a trust management object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.13.
//
// Added in macOS 10.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustSetKeychains(_:_:)
func SecTrustSetKeychains(trust unsafe.Pointer, keychainOrArray unsafe.Pointer) unsafe.Pointer {
	return _SecTrustSetKeychains(trust, keychainOrArray)
	}




