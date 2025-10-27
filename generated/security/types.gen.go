// Code generated from Apple documentation for Security. DO NOT EDIT.

package security
import (
	"unsafe"
)


// C struct types
// _CE_ExtendedKeyUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_ExtendedKeyUsage-c.struct
type _CE_ExtendedKeyUsage struct {
	NumPurposes unsafe.Pointer
	Purposes unsafe.Pointer
}

// SecCEBasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEBasicConstraints
type SecCEBasicConstraints struct {
	Critical bool
	IsCA bool
	PathLenConstraint uint32
	PathLenConstraintPresent bool
	Present bool
}

// AuthorizationExternalForm - The external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationExternalForm
type AuthorizationExternalForm struct {
	Bytes unsafe.Pointer // An array of characters representing the external form of an authorization reference.
}

// AuthorizationItem - A structure containing information about an authorization right or the authorization environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItem
type AuthorizationItem struct {
	Flags unsafe.Pointer // Reserved option bits.
	Name AuthorizationString // The required name of the authorization right or environment data.
	Value unsafe.Pointer // A pointer to information pertaining to the name field.
	ValueLength uintptr // The number of bytes in the value field.
}

// AuthorizationItemSet - A structure containing a set of authorization items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItemSet
type AuthorizationItemSet struct {
	Count unsafe.Pointer // The number of elements in the   array.
	Items AuthorizationItem // A pointer to an array of authorization items.
}

// cssm_acl_keychain_prompt_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_keychain_prompt_selector-swift.struct
type cssm_acl_keychain_prompt_selector struct {
	Flags unsafe.Pointer
	Version unsafe.Pointer
}

// cssm_acl_process_subject_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_process_subject_selector-swift.struct
type cssm_acl_process_subject_selector struct {
	Gid unsafe.Pointer
	Mask unsafe.Pointer
	Uid unsafe.Pointer
	Version unsafe.Pointer
}

// CSSM_APPLE_CL_CSR_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_CL_CSR_REQUEST
type CSSM_APPLE_CL_CSR_REQUEST struct {
	ChallengeString unsafe.Pointer
	CspHand CSSM_CSP_HANDLE
	SignatureAlg CSSM_ALGORITHMS
	SignatureOid SecAsn1Oid
	SubjectNameX509 unsafe.Pointer
	SubjectPrivateKey unsafe.Pointer
	SubjectPublicKey unsafe.Pointer
}

// CSSM_APPLE_TP_ACTION_DATA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_DATA
type CSSM_APPLE_TP_ACTION_DATA struct {
	ActionFlags CSSM_APPLE_TP_ACTION_FLAGS
	Version unsafe.Pointer
}

// CSSM_APPLE_TP_CERT_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CERT_REQUEST
type CSSM_APPLE_TP_CERT_REQUEST struct {
	CertPublicKey unsafe.Pointer
	ChallengeString unsafe.Pointer
	ClHand CSSM_CL_HANDLE
	CspHand CSSM_CSP_HANDLE
	Extensions unsafe.Pointer
	IssuerNames CSSM_APPLE_TP_NAME_OID
	IssuerNameX509 unsafe.Pointer
	IssuerPrivateKey unsafe.Pointer
	NotAfter unsafe.Pointer
	NotBefore unsafe.Pointer
	NumExtensions unsafe.Pointer
	NumIssuerNames unsafe.Pointer
	NumSubjectNames unsafe.Pointer
	SerialNumber unsafe.Pointer
	SignatureAlg CSSM_ALGORITHMS
	SignatureOid SecAsn1Oid
	SubjectNames CSSM_APPLE_TP_NAME_OID
}

// CSSM_APPLE_TP_CRL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPTIONS
type CSSM_APPLE_TP_CRL_OPTIONS struct {
	CrlFlags CSSM_APPLE_TP_CRL_OPT_FLAGS
	CrlStore unsafe.Pointer
	Version unsafe.Pointer
}

// CSSM_APPLE_TP_NAME_OID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_NAME_OID
type CSSM_APPLE_TP_NAME_OID struct {
	Oid SecAsn1Oid
	String unsafe.Pointer
}

// CSSM_APPLE_TP_SMIME_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SMIME_OPTIONS
type CSSM_APPLE_TP_SMIME_OPTIONS struct {
	IntendedUsage unsafe.Pointer
	SenderEmail unsafe.Pointer
	SenderEmailLen unsafe.Pointer
	Version unsafe.Pointer
}

// CSSM_APPLE_TP_SSL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SSL_OPTIONS
type CSSM_APPLE_TP_SSL_OPTIONS struct {
	Flags unsafe.Pointer
	ServerName unsafe.Pointer
	ServerNameLen unsafe.Pointer
	Version unsafe.Pointer
}

// cssm_applecspdl_db_change_password_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_change_password_parameters-swift.struct
type cssm_applecspdl_db_change_password_parameters struct {
	AccessCredentials unsafe.Pointer
}

// cssm_applecspdl_db_is_locked_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_is_locked_parameters-swift.struct
type cssm_applecspdl_db_is_locked_parameters struct {
	IsLocked unsafe.Pointer
}

// cssm_applecspdl_db_settings_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_settings_parameters-swift.struct
type cssm_applecspdl_db_settings_parameters struct {
	IdleTimeout unsafe.Pointer
	LockOnSleep unsafe.Pointer
}

// cssm_appledl_open_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters-swift.struct
type cssm_appledl_open_parameters struct {
	AutoCommit CSSM_BOOL
	Length unsafe.Pointer
	Mask unsafe.Pointer
	Mode unsafe.Pointer
	Version unsafe.Pointer
}

// cssm_authorizationgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_authorizationgroup-swift.struct
type cssm_authorizationgroup struct {
	AuthTags CSSM_ACL_AUTHORIZATION_TAG
	NumberOfAuthTags unsafe.Pointer
}

// cssm_csp_operational_statistics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_csp_operational_statistics-swift.struct
type cssm_csp_operational_statistics struct {
	DeviceFlags CSSM_CSP_FLAGS
	TokenFreePrivateMem unsafe.Pointer
	TokenFreePublicMem unsafe.Pointer
	TokenMaxRWSessionCount unsafe.Pointer
	TokenMaxSessionCount unsafe.Pointer
	TokenOpenedRWSessionCount unsafe.Pointer
	TokenOpenedSessionCount unsafe.Pointer
	TokenTotalPrivateMem unsafe.Pointer
	TokenTotalPublicMem unsafe.Pointer
	UserAuthenticated CSSM_BOOL
}

// cssm_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_data-swift.struct
type cssm_data struct {
	Data unsafe.Pointer
	Length uintptr
}

// cssm_date
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_date-swift.struct
type cssm_date struct {
	Day unsafe.Pointer
	Month unsafe.Pointer
	Year unsafe.Pointer
}

// cssm_db_schema_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_index_info-swift.struct
type cssm_db_schema_index_info struct {
	AttributeId unsafe.Pointer
	IndexedDataLocation CSSM_DB_INDEXED_DATA_LOCATION
	IndexId unsafe.Pointer
	IndexType CSSM_DB_INDEX_TYPE
}

// cssm_dl_db_handle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_db_handle-swift.struct
type cssm_dl_db_handle struct {
	DBHandle CSSM_DB_HANDLE
	DLHandle CSSM_DL_HANDLE
}

// cssm_dl_pkcs11_attributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_pkcs11_attributes
type cssm_dl_pkcs11_attributes struct {
	DeviceAccessFlags unsafe.Pointer
}

// cssm_func_name_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_func_name_addr-swift.struct
type cssm_func_name_addr struct {
	Address CSSM_PROC_ADDR
	Name CSSM_STRING
}

// cssm_guid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_guid-swift.struct
type cssm_guid struct {
	Data1 unsafe.Pointer
	Data2 unsafe.Pointer
	Data3 unsafe.Pointer
	Data4 unsafe.Pointer
}

// cssm_key_size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key_size-swift.struct
type cssm_key_size struct {
	EffectiveKeySizeInBits unsafe.Pointer
	LogicalKeySizeInBits unsafe.Pointer
}

// cssm_kr_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_name-swift.struct
type cssm_kr_name struct {
	Length unsafe.Pointer
	Name unsafe.Pointer
	Type unsafe.Pointer
}

// cssm_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list-swift.struct
type cssm_list struct {
	Head CSSM_LIST_ELEMENT_PTR
	ListType CSSM_LIST_TYPE
	Tail CSSM_LIST_ELEMENT_PTR
}

// cssm_memory_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_memory_funcs-swift.struct
type cssm_memory_funcs struct {
	AllocRef unsafe.Pointer
	Calloc_func CSSM_CALLOC
	Free_func CSSM_FREE
	Malloc_func CSSM_MALLOC
	Realloc_func CSSM_REALLOC
}

// cssm_name_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_name_list-swift.struct
type cssm_name_list struct {
	NumStrings unsafe.Pointer
	String unsafe.Pointer
}

// cssm_parsed_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_cert-swift.struct
type cssm_parsed_cert struct {
	CertType CSSM_CERT_TYPE
	ParsedCert unsafe.Pointer
	ParsedCertFormat CSSM_CERT_PARSE_FORMAT
}

// cssm_parsed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_crl-swift.struct
type cssm_parsed_crl struct {
	CrlType CSSM_CRL_TYPE
	ParsedCrl unsafe.Pointer
	ParsedCrlFormat CSSM_CRL_PARSE_FORMAT
}

// cssm_query_size_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_size_data-swift.struct
type cssm_query_size_data struct {
	SizeInputBlock unsafe.Pointer
	SizeOutputBlock unsafe.Pointer
}

// cssm_range
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_range-swift.struct
type cssm_range struct {
	Max unsafe.Pointer
	Min unsafe.Pointer
}

// CSSM_TP_APPLE_EVIDENCE_HEADER
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_EVIDENCE_HEADER
type CSSM_TP_APPLE_EVIDENCE_HEADER struct {
	Version unsafe.Pointer
}

// cssm_tp_result_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_result_set-swift.struct
type cssm_tp_result_set struct {
	NumberOfResults unsafe.Pointer
	Results unsafe.Pointer
}

// SecAsn1AlgId - A structure identifying an ASN.1 algorithm by its OID, and its corresponding parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AlgId
type SecAsn1AlgId struct {
	Algorithm SecAsn1Oid
	Parameters SecAsn1Item
}

// SecAsn1PubKeyInfo - A structure containing a public key and its associated algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1PubKeyInfo
type SecAsn1PubKeyInfo struct {
	Algorithm SecAsn1AlgId
	SubjectPublicKey SecAsn1Item
}

// SecAsn1Template_struct - A structure that defines one element of a BER or DER encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Template_struct
type SecAsn1Template_struct struct {
	Kind uint32
	Offset uint32
	Size uint32
	Sub unsafe.Pointer
}

// SecItemImportExportKeyParameters - The import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportKeyParameters
type SecItemImportExportKeyParameters struct {
	AccessRef SecAccessRef // Specifies the initial access controls of imported private keys.
	AlertPrompt StringRef // The prompt to display in the secure passphrase alert panel.
	AlertTitle StringRef // The title to display in the secure passphrase alert panel.
	Flags SecKeyImportExportFlags // The bitwise   of zero or more key import/export flags.
	KeyAttributes ArrayRef // An array containing zero or more key attributes for an imported key.
	KeyUsage ArrayRef // An array containing usage attributes applied to a key on import.
	Passphrase TypeRef // The password to use during key import or export.
	Version uint32 // The version of this structure.
}

// SecKeychainAttribute - A structure that holds a single keychain attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttribute
type SecKeychainAttribute struct {
	Data unsafe.Pointer // A pointer to the attribute data.
	Length unsafe.Pointer // The length of the buffer pointed to by data.
	Tag SecKeychainAttrType // A 4-byte attribute tag.
}

// SecKeychainAttributeInfo - A structure that represents an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeInfo
type SecKeychainAttributeInfo struct {
	Count unsafe.Pointer // The number of tag-format pairs in the respective arrays.
	Format unsafe.Pointer // A pointer to the first attribute format in the array.
	Tag unsafe.Pointer // A pointer to the first attribute tag in the array.
}

// SecKeychainAttributeList - A list of keychain attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeList
type SecKeychainAttributeList struct {
	Attr SecKeychainAttribute // A pointer to the first keychain attribute in the array.
	Count unsafe.Pointer // The number of keychain attributes in the array.
}

// SecKeychainCallbackInfo - Information about a keychain event that keychain services deliver to your app via a callback function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCallbackInfo
type SecKeychainCallbackInfo struct {
	Item SecKeychainItemRef // A reference to the keychain item in which the event occurred. If the event did not involve an item, this field is not valid.
	Keychain SecKeychainRef // A reference to the keychain in which the event occurred. If the event did not involve a keychain, this field is not valid.
	Pid unsafe.Pointer // The ID of the process that generated this event.
	Version unsafe.Pointer // The version of this structure.
}

// SecKeychainSettings - A structure that contains information about keychain settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSettings
type SecKeychainSettings struct {
	LockInterval unsafe.Pointer // The number of seconds to wait before the keychain locks.
	LockOnSleep unsafe.Pointer // A Boolean value indicating whether the keychain locks when the system sleeps.
	UseLockInterval unsafe.Pointer // A Boolean value indicating whether the keychain automatically locks after a certain period of time.
	Version unsafe.Pointer // The keychain version.
}

// SecKeyImportExportParameters - The legacy import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportParameters
type SecKeyImportExportParameters struct {
	AccessRef SecAccessRef // Specifies the initial access controls of imported private keys.
	AlertPrompt StringRef // The prompt to display in the secure passphrase alert panel.
	AlertTitle StringRef // The title to display in the secure passphrase alert panel.
	Flags SecKeyImportExportFlags // The bitwise   of zero or more key import/export flags.
	KeyAttributes CSSM_KEYATTR_FLAGS // A word of bits constituting the low-level attribute flags for imported keys.
	KeyUsage CSSM_KEYUSE // A word of bits constituting the low-level use flags for imported keys.
	Passphrase TypeRef // The password to use during key import or export.
	Version uint32 // The version of this structure.
}





