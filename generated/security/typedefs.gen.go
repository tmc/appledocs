// Code generated from Apple documentation for Security. DO NOT EDIT.

package security
import (
"unsafe"
)

// Type aliases and typedefs
// SecCECrlReason type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCECrlReason
// SecCECrlReason has base type: uint32_t
type SecCECrlReason uintptr
// AuthorizationEnvironment - An authorization item set designated to hold environment information relevant to authorization decisions.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationEnvironment
// AuthorizationEnvironment has base type: AuthorizationItemSet
type AuthorizationEnvironment uintptr
// AuthorizationRef - A pointer to an opaque authorization reference structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRef
// AuthorizationRef has base type: const struct AuthorizationOpaqueRef *
type AuthorizationRef uintptr
// AuthorizationRights - An authorization item set designated to represent a set of rights.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationRights
// AuthorizationRights has base type: AuthorizationItemSet
type AuthorizationRights uintptr
// AuthorizationString - A zero-terminated string in UTF-8 encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationString
// AuthorizationString has base type: const char *
type AuthorizationString uintptr
// CE_CrlNumber type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CrlNumber
type CE_CrlNumber = uint32
// CE_DeltaCrl type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DeltaCrl
type CE_DeltaCrl = uint32
// SDecoderRef - An opaque reference to a CMS decoder object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoder
// CMSDecoderRef has base type: struct _CMSDecoder *
type SDecoderRef uintptr
// SEncoderRef - Opaque reference to a CMS encoder object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoder
// CMSEncoderRef has base type: struct _CMSEncoder *
type SEncoderRef uintptr
// CSSM_AC_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_AC_HANDLE
// CSSM_AC_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_AC_HANDLE uintptr
// CSSM_ACL_AUTHORIZATION_TAG type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_AUTHORIZATION_TAG
// CSSM_ACL_AUTHORIZATION_TAG has base type: sint32
type CSSM_ACL_AUTHORIZATION_TAG uintptr
// CSSM_ACL_EDIT_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_EDIT_MODE
type CSSM_ACL_EDIT_MODE = uint32
// CSSM_ACL_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_HANDLE
// CSSM_ACL_HANDLE has base type: CSSM_HANDLE
type CSSM_ACL_HANDLE uintptr
// CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR-swift.typealias
// CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR has base type: struct cssm_acl_keychain_prompt_selector
type CSSM_ACL_KEYCHAIN_PROMPT_SELECTOR uintptr
// CSSM_ACL_PREAUTH_TRACKING_STATE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_PREAUTH_TRACKING_STATE
type CSSM_ACL_PREAUTH_TRACKING_STATE = uint32
// CSSM_ACL_PROCESS_SUBJECT_SELECTOR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_PROCESS_SUBJECT_SELECTOR-swift.typealias
// CSSM_ACL_PROCESS_SUBJECT_SELECTOR has base type: struct cssm_acl_process_subject_selector
type CSSM_ACL_PROCESS_SUBJECT_SELECTOR uintptr
// CSSM_ACL_SUBJECT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_SUBJECT_TYPE
// CSSM_ACL_SUBJECT_TYPE has base type: sint32
type CSSM_ACL_SUBJECT_TYPE uintptr
// CSSM_ALGORITHMS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ALGORITHMS
type CSSM_ALGORITHMS = uint32
// CSSM_APPLE_TP_ACTION_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_FLAGS
type CSSM_APPLE_TP_ACTION_FLAGS = uint32
// CSSM_APPLE_TP_CRL_OPT_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPT_FLAGS
type CSSM_APPLE_TP_CRL_OPT_FLAGS = uint32
// CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS-swift.typealias
// CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS has base type: struct cssm_applecspdl_db_change_password_parameters
type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS uintptr
// CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR
// CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR has base type: struct cssm_applecspdl_db_change_password_parameters *
type CSSM_APPLECSPDL_DB_CHANGE_PASSWORD_PARAMETERS_PTR uintptr
// CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS-swift.typealias
// CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS has base type: struct cssm_applecspdl_db_is_locked_parameters
type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS uintptr
// CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR
// CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR has base type: struct cssm_applecspdl_db_is_locked_parameters *
type CSSM_APPLECSPDL_DB_IS_LOCKED_PARAMETERS_PTR uintptr
// CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS-swift.typealias
// CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS has base type: struct cssm_applecspdl_db_settings_parameters
type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS uintptr
// CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR
// CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR has base type: struct cssm_applecspdl_db_settings_parameters *
type CSSM_APPLECSPDL_DB_SETTINGS_PARAMETERS_PTR uintptr
// CSSM_APPLEDL_OPEN_PARAMETERS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLEDL_OPEN_PARAMETERS-swift.typealias
// CSSM_APPLEDL_OPEN_PARAMETERS has base type: struct cssm_appledl_open_parameters
type CSSM_APPLEDL_OPEN_PARAMETERS uintptr
// CSSM_APPLEDL_OPEN_PARAMETERS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLEDL_OPEN_PARAMETERS_PTR
// CSSM_APPLEDL_OPEN_PARAMETERS_PTR has base type: struct cssm_appledl_open_parameters *
type CSSM_APPLEDL_OPEN_PARAMETERS_PTR uintptr
// CSSM_ATTACH_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ATTACH_FLAGS
type CSSM_ATTACH_FLAGS = uint32
// CSSM_ATTRIBUTE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ATTRIBUTE_TYPE
type CSSM_ATTRIBUTE_TYPE = uint32
// CSSM_BER_TAG type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BER_TAG
// CSSM_BER_TAG has base type: uint8
type CSSM_BER_TAG uintptr
// CSSM_BITMASK type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BITMASK
type CSSM_BITMASK = uint32
// CSSM_BOOL type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BOOL
// CSSM_BOOL has base type: sint32
type CSSM_BOOL uintptr
// CSSM_CALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CALLOC
// CSSM_CALLOC is a callback function
// C type: void *(*)(unsigned int, unsigned long, void *)
type CSSM_CALLOC = func(uint32, uint, unsafe.Pointer) unsafe.Pointer
// CSSM_CC_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CC_HANDLE
// CSSM_CC_HANDLE has base type: CSSM_LONG_HANDLE
type CSSM_CC_HANDLE uintptr
// CSSM_CERT_BUNDLE_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_ENCODING
type CSSM_CERT_BUNDLE_ENCODING = uint32
// CSSM_CERT_BUNDLE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_TYPE
type CSSM_CERT_BUNDLE_TYPE = uint32
// CSSM_CERT_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING
type CSSM_CERT_ENCODING = uint32
// CSSM_CERT_ENCODING_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING_PTR
// CSSM_CERT_ENCODING_PTR has base type: uint32 *
type CSSM_CERT_ENCODING_PTR uintptr
// CSSM_CERT_PARSE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT
type CSSM_CERT_PARSE_FORMAT = uint32
// CSSM_CERT_PARSE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT_PTR
// CSSM_CERT_PARSE_FORMAT_PTR has base type: uint32 *
type CSSM_CERT_PARSE_FORMAT_PTR uintptr
// CSSM_CERT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE
type CSSM_CERT_TYPE = uint32
// CSSM_CERT_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE_PTR
// CSSM_CERT_TYPE_PTR has base type: uint32 *
type CSSM_CERT_TYPE_PTR uintptr
// CSSM_CERTGROUP_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE
type CSSM_CERTGROUP_TYPE = uint32
// CSSM_CERTGROUP_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE_PTR
// CSSM_CERTGROUP_TYPE_PTR has base type: uint32 *
type CSSM_CERTGROUP_TYPE_PTR uintptr
// CSSM_CL_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_HANDLE
// CSSM_CL_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_CL_HANDLE uintptr
// CSSM_CL_TEMPLATE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_TEMPLATE_TYPE
type CSSM_CL_TEMPLATE_TYPE = uint32
// CSSM_CONTEXT_EVENT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_EVENT
type CSSM_CONTEXT_EVENT = uint32
// CSSM_CONTEXT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_TYPE
type CSSM_CONTEXT_TYPE = uint32
// CSSM_CRL_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING
type CSSM_CRL_ENCODING = uint32
// CSSM_CRL_ENCODING_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING_PTR
// CSSM_CRL_ENCODING_PTR has base type: uint32 *
type CSSM_CRL_ENCODING_PTR uintptr
// CSSM_CRL_PARSE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT
type CSSM_CRL_PARSE_FORMAT = uint32
// CSSM_CRL_PARSE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT_PTR
// CSSM_CRL_PARSE_FORMAT_PTR has base type: uint32 *
type CSSM_CRL_PARSE_FORMAT_PTR uintptr
// CSSM_CRL_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE
type CSSM_CRL_TYPE = uint32
// CSSM_CRL_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE_PTR
// CSSM_CRL_TYPE_PTR has base type: uint32 *
type CSSM_CRL_TYPE_PTR uintptr
// CSSM_CRLGROUP_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE
type CSSM_CRLGROUP_TYPE = uint32
// CSSM_CRLGROUP_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE_PTR
// CSSM_CRLGROUP_TYPE_PTR has base type: uint32 *
type CSSM_CRLGROUP_TYPE_PTR uintptr
// CSSM_CSP_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_FLAGS
type CSSM_CSP_FLAGS = uint32
// CSSM_CSP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_HANDLE
// CSSM_CSP_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_CSP_HANDLE uintptr
// CSSM_CSP_READER_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_READER_FLAGS
type CSSM_CSP_READER_FLAGS = uint32
// CSSM_CSPTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSPTYPE
type CSSM_CSPTYPE = uint32
// CSSM_DB_ACCESS_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE
type CSSM_DB_ACCESS_TYPE = uint32
// CSSM_DB_ACCESS_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE_PTR
// CSSM_DB_ACCESS_TYPE_PTR has base type: uint32 *
type CSSM_DB_ACCESS_TYPE_PTR uintptr
// CSSM_DB_ATTRIBUTE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT
type CSSM_DB_ATTRIBUTE_FORMAT = uint32
// CSSM_DB_ATTRIBUTE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT_PTR
// CSSM_DB_ATTRIBUTE_FORMAT_PTR has base type: uint32 *
type CSSM_DB_ATTRIBUTE_FORMAT_PTR uintptr
// CSSM_DB_ATTRIBUTE_NAME_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT
type CSSM_DB_ATTRIBUTE_NAME_FORMAT = uint32
// CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR
// CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR has base type: uint32 *
type CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR uintptr
// CSSM_DB_CONJUNCTIVE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_CONJUNCTIVE
type CSSM_DB_CONJUNCTIVE = uint32
// CSSM_DB_CONJUNCTIVE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_CONJUNCTIVE_PTR
// CSSM_DB_CONJUNCTIVE_PTR has base type: uint32 *
type CSSM_DB_CONJUNCTIVE_PTR uintptr
// CSSM_DB_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_HANDLE
// CSSM_DB_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_DB_HANDLE uintptr
// CSSM_DB_INDEX_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_INDEX_TYPE
type CSSM_DB_INDEX_TYPE = uint32
// CSSM_DB_INDEXED_DATA_LOCATION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_INDEXED_DATA_LOCATION
type CSSM_DB_INDEXED_DATA_LOCATION = uint32
// CSSM_DB_MODIFY_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_MODIFY_MODE
type CSSM_DB_MODIFY_MODE = uint32
// CSSM_DB_OPERATOR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR
type CSSM_DB_OPERATOR = uint32
// CSSM_DB_OPERATOR_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR_PTR
// CSSM_DB_OPERATOR_PTR has base type: uint32 *
type CSSM_DB_OPERATOR_PTR uintptr
// CSSM_DB_RECORDTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_RECORDTYPE
type CSSM_DB_RECORDTYPE = uint32
// CSSM_DB_RETRIEVAL_MODES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_RETRIEVAL_MODES
type CSSM_DB_RETRIEVAL_MODES = uint32
// CSSM_DL_CUSTOM_ATTRIBUTES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_CUSTOM_ATTRIBUTES
// CSSM_DL_CUSTOM_ATTRIBUTES has base type: void *
type CSSM_DL_CUSTOM_ATTRIBUTES uintptr
// CSSM_DL_FFS_ATTRIBUTES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_FFS_ATTRIBUTES
// CSSM_DL_FFS_ATTRIBUTES has base type: void *
type CSSM_DL_FFS_ATTRIBUTES uintptr
// CSSM_DL_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_HANDLE
// CSSM_DL_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_DL_HANDLE uintptr
// CSSM_DL_LDAP_ATTRIBUTES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_LDAP_ATTRIBUTES
// CSSM_DL_LDAP_ATTRIBUTES has base type: void *
type CSSM_DL_LDAP_ATTRIBUTES uintptr
// CSSM_DL_ODBC_ATTRIBUTES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_ODBC_ATTRIBUTES
// CSSM_DL_ODBC_ATTRIBUTES has base type: void *
type CSSM_DL_ODBC_ATTRIBUTES uintptr
// CSSM_DL_PKCS11_ATTRIBUTE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE
// CSSM_DL_PKCS11_ATTRIBUTE has base type: struct cssm_dl_pkcs11_attributes *
type CSSM_DL_PKCS11_ATTRIBUTE uintptr
// CSSM_DL_PKCS11_ATTRIBUTE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DL_PKCS11_ATTRIBUTE_PTR
// CSSM_DL_PKCS11_ATTRIBUTE_PTR has base type: struct cssm_dl_pkcs11_attributes *
type CSSM_DL_PKCS11_ATTRIBUTE_PTR uintptr
// CSSM_DLTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DLTYPE
type CSSM_DLTYPE = uint32
// CSSM_DLTYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DLTYPE_PTR
// CSSM_DLTYPE_PTR has base type: uint32 *
type CSSM_DLTYPE_PTR uintptr
// CSSM_ENCRYPT_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ENCRYPT_MODE
type CSSM_ENCRYPT_MODE = uint32
// CSSM_EVIDENCE_FORM type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EVIDENCE_FORM
type CSSM_EVIDENCE_FORM = uint32
// CSSM_FREE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_FREE
// CSSM_FREE is a callback function
// C type: void (*)(void *, void *)
type CSSM_FREE = func(unsafe.Pointer, unsafe.Pointer)
// CSSM_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_HANDLE
// CSSM_HANDLE has base type: CSSM_INTPTR
type CSSM_HANDLE uintptr
// CSSM_HANDLE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_HANDLE_PTR
// CSSM_HANDLE_PTR has base type: CSSM_INTPTR *
type CSSM_HANDLE_PTR uintptr
// CSSM_HEADERVERSION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_HEADERVERSION
type CSSM_HEADERVERSION = uint32
// CSSM_INTPTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_INTPTR
// CSSM_INTPTR has base type: intptr_t
type CSSM_INTPTR uintptr
// CSSM_KEY_HIERARCHY type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEY_HIERARCHY
// CSSM_KEY_HIERARCHY has base type: CSSM_BITMASK
type CSSM_KEY_HIERARCHY uintptr
// CSSM_KEY_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEY_TYPE
// CSSM_KEY_TYPE has base type: CSSM_ALGORITHMS
type CSSM_KEY_TYPE uintptr
// CSSM_KEYATTR_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYATTR_FLAGS
type CSSM_KEYATTR_FLAGS = uint32
// CSSM_KEYBLOB_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_FORMAT
type CSSM_KEYBLOB_FORMAT = uint32
// CSSM_KEYBLOB_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_TYPE
type CSSM_KEYBLOB_TYPE = uint32
// CSSM_KEYCLASS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYCLASS
type CSSM_KEYCLASS = uint32
// CSSM_KEYUSE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYUSE
type CSSM_KEYUSE = uint32
// CSSM_KR_POLICY_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_FLAGS
type CSSM_KR_POLICY_FLAGS = uint32
// CSSM_KR_POLICY_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_TYPE
type CSSM_KR_POLICY_TYPE = uint32
// CSSM_KRSP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KRSP_HANDLE
type CSSM_KRSP_HANDLE = uint32
// CSSM_LIST_ELEMENT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_PTR
// CSSM_LIST_ELEMENT_PTR has base type: struct cssm_list_element *
type CSSM_LIST_ELEMENT_PTR uintptr
// CSSM_LIST_ELEMENT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE
type CSSM_LIST_ELEMENT_TYPE = uint32
// CSSM_LIST_ELEMENT_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE_PTR
// CSSM_LIST_ELEMENT_TYPE_PTR has base type: uint32 *
type CSSM_LIST_ELEMENT_TYPE_PTR uintptr
// CSSM_LIST_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE
type CSSM_LIST_TYPE = uint32
// CSSM_LIST_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE_PTR
// CSSM_LIST_TYPE_PTR has base type: uint32 *
type CSSM_LIST_TYPE_PTR uintptr
// CSSM_LONG_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE
type CSSM_LONG_HANDLE = uint64
// CSSM_LONG_HANDLE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE_PTR
// CSSM_LONG_HANDLE_PTR has base type: uint64 *
type CSSM_LONG_HANDLE_PTR uintptr
// CSSM_MALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MALLOC
// CSSM_MALLOC is a callback function
// C type: void *(*)(unsigned long, void *)
type CSSM_MALLOC = func(uint, unsafe.Pointer) unsafe.Pointer
// CSSM_MANAGER_EVENT_TYPES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MANAGER_EVENT_TYPES
type CSSM_MANAGER_EVENT_TYPES = uint32
// CSSM_MODULE_EVENT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT
type CSSM_MODULE_EVENT = uint32
// CSSM_MODULE_EVENT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT_PTR
// CSSM_MODULE_EVENT_PTR has base type: uint32 *
type CSSM_MODULE_EVENT_PTR uintptr
// CSSM_MODULE_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE
// CSSM_MODULE_HANDLE has base type: CSSM_HANDLE
type CSSM_MODULE_HANDLE uintptr
// CSSM_MODULE_HANDLE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MODULE_HANDLE_PTR
// CSSM_MODULE_HANDLE_PTR has base type: CSSM_HANDLE *
type CSSM_MODULE_HANDLE_PTR uintptr
// CSSM_NET_ADDRESS_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_NET_ADDRESS_TYPE
type CSSM_NET_ADDRESS_TYPE = uint32
// CSSM_NET_PROTOCOL type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_NET_PROTOCOL
type CSSM_NET_PROTOCOL = uint32
// CSSM_PADDING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PADDING
type CSSM_PADDING = uint32
// CSSM_PKCS5_PBKDF2_PRF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS5_PBKDF2_PRF
type CSSM_PKCS5_PBKDF2_PRF = uint32
// CSSM_PKCS_OAEP_MGF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_MGF
type CSSM_PKCS_OAEP_MGF = uint32
// CSSM_PKCS_OAEP_PSOURCE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_PSOURCE
type CSSM_PKCS_OAEP_PSOURCE = uint32
// CSSM_PRIVILEGE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE
type CSSM_PRIVILEGE = uint64
// CSSM_PRIVILEGE_SCOPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE_SCOPE
type CSSM_PRIVILEGE_SCOPE = uint32
// CSSM_PROC_ADDR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PROC_ADDR
// CSSM_PROC_ADDR is a callback function
// C type: void (*)(void)
type CSSM_PROC_ADDR = func()
// CSSM_PROC_ADDR_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PROC_ADDR_PTR
// CSSM_PROC_ADDR_PTR has base type: CSSM_PROC_ADDR *
type CSSM_PROC_ADDR_PTR uintptr
// CSSM_PVC_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PVC_MODE
// CSSM_PVC_MODE has base type: CSSM_BITMASK
type CSSM_PVC_MODE uintptr
// CSSM_QUERY_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_QUERY_FLAGS
type CSSM_QUERY_FLAGS = uint32
// CSSM_REALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_REALLOC
// CSSM_REALLOC is a callback function
// C type: void *(*)(void *, unsigned long, void *)
type CSSM_REALLOC = func(unsafe.Pointer, uint, unsafe.Pointer) unsafe.Pointer
// CSSM_RETURN type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_RETURN
// CSSM_RETURN has base type: sint32
type CSSM_RETURN uintptr
// CSSM_SAMPLE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SAMPLE_TYPE
// CSSM_SAMPLE_TYPE has base type: CSSM_WORDID_TYPE
type CSSM_SAMPLE_TYPE uintptr
// CSSM_SC_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SC_FLAGS
type CSSM_SC_FLAGS = uint32
// CSSM_SERVICE_MASK type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SERVICE_MASK
type CSSM_SERVICE_MASK = uint32
// CSSM_SERVICE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SERVICE_TYPE
// CSSM_SERVICE_TYPE has base type: CSSM_SERVICE_MASK
type CSSM_SERVICE_TYPE uintptr
// CSSM_SIZE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SIZE
// CSSM_SIZE has base type: size_t
type CSSM_SIZE uintptr
// CSSM_STRING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_STRING
// CSSM_STRING has base type: char[68]
type CSSM_STRING uintptr
// CSSM_TIMESTRING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TIMESTRING
// CSSM_TIMESTRING has base type: char *
type CSSM_TIMESTRING uintptr
// CSSM_TP_ACTION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_ACTION
type CSSM_TP_ACTION = uint32
// CSSM_TP_APPLE_CERT_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_CERT_STATUS
type CSSM_TP_APPLE_CERT_STATUS = uint32
// CSSM_TP_AUTHORITY_REQUEST_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE
type CSSM_TP_AUTHORITY_REQUEST_TYPE = uint32
// CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR
// CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR has base type: uint32 *
type CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR uintptr
// CSSM_TP_CERTCHANGE_ACTION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_ACTION
type CSSM_TP_CERTCHANGE_ACTION = uint32
// CSSM_TP_CERTCHANGE_REASON type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_REASON
type CSSM_TP_CERTCHANGE_REASON = uint32
// CSSM_TP_CERTCHANGE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_STATUS
type CSSM_TP_CERTCHANGE_STATUS = uint32
// CSSM_TP_CERTISSUE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTISSUE_STATUS
type CSSM_TP_CERTISSUE_STATUS = uint32
// CSSM_TP_CERTNOTARIZE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTNOTARIZE_STATUS
type CSSM_TP_CERTNOTARIZE_STATUS = uint32
// CSSM_TP_CERTRECLAIM_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTRECLAIM_STATUS
type CSSM_TP_CERTRECLAIM_STATUS = uint32
// CSSM_TP_CERTVERIFY_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTVERIFY_STATUS
type CSSM_TP_CERTVERIFY_STATUS = uint32
// CSSM_TP_CONFIRM_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS
type CSSM_TP_CONFIRM_STATUS = uint32
// CSSM_TP_CONFIRM_STATUS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS_PTR
// CSSM_TP_CONFIRM_STATUS_PTR has base type: uint32 *
type CSSM_TP_CONFIRM_STATUS_PTR uintptr
// CSSM_TP_CRLISSUE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CRLISSUE_STATUS
type CSSM_TP_CRLISSUE_STATUS = uint32
// CSSM_TP_FORM_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_FORM_TYPE
type CSSM_TP_FORM_TYPE = uint32
// CSSM_TP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_HANDLE
// CSSM_TP_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_TP_HANDLE uintptr
// CSSM_TP_SERVICES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_SERVICES
type CSSM_TP_SERVICES = uint32
// CSSM_TP_STOP_ON type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_STOP_ON
type CSSM_TP_STOP_ON = uint32
// CSSM_USEE_TAG type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_USEE_TAG
// CSSM_USEE_TAG has base type: CSSM_PRIVILEGE
type CSSM_USEE_TAG uintptr
// CSSM_WORDID_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_WORDID_TYPE
// CSSM_WORDID_TYPE has base type: sint32
type CSSM_WORDID_TYPE uintptr
// CSSM_X509_OPTION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509_OPTION
// CSSM_X509_OPTION has base type: CSSM_BOOL
type CSSM_X509_OPTION uintptr
// CSSM_X509EXT_DATA_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509EXT_DATA_FORMAT
// CSSM_X509EXT_DATA_FORMAT has base type: enum extension_data_format
type CSSM_X509EXT_DATA_FORMAT uintptr
// SecAccessRef - An opaque type that identifies a keychain item’s access information.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccess
// SecAccessRef has base type: struct __SecAccess *
type SecAccessRef uintptr
// SecAccessControlRef - An opaque type that contains information about how a keychain item may be used.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessControl
// SecAccessControlRef has base type: struct __SecAccessControl *
type SecAccessControlRef uintptr
// SecAccessOwnerType - A type for flags that enable you to configure ACL ownership.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAccessOwnerType
// SecAccessOwnerType has base type: UInt32
type SecAccessOwnerType uintptr
// SecACLRef - An opaque type that represents information about an ACL entry.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecACL
// SecACLRef has base type: struct __SecACL *
type SecACLRef uintptr
// SecAFPServerSignature - Represents a 16-byte Apple File Protocol server signature block.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAFPServerSignature
// SecAFPServerSignature has base type: UInt8[16]
type SecAFPServerSignature uintptr
// SecAsn1Item - A structure holding DER encoded data.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Item
// SecAsn1Item has base type: struct cssm_data
type SecAsn1Item uintptr
// SecAsn1Oid - An object identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Oid
// SecAsn1Oid has base type: struct cssm_data
type SecAsn1Oid uintptr
// SecAsn1TemplateChooser - Dynamically provides the sub-template to use during encode or decode.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1TemplateChooser
// SecAsn1TemplateChooser has base type: const struct SecAsn1Template_struct *(void *, unsigned char, const char *, unsigned long, void *)
type SecAsn1TemplateChooser uintptr
// SecAsn1TemplateChooserPtr - A pointer to the template chooser function.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1TemplateChooserPtr
// SecAsn1TemplateChooserPtr is a callback function
// C type: const struct SecAsn1Template_struct *(*)(void *, unsigned char, const char *, unsigned long, void *)
type SecAsn1TemplateChooserPtr = func(unsafe.Pointer, uint8, string, uint, unsafe.Pointer) unsafe.Pointer
// SecCertificateRef - An abstract Core Foundation-type object representing an X.509 certificate.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCertificate
// SecCertificateRef has base type: struct __SecCertificate *
type SecCertificateRef uintptr
// SecCodeRef - A code object representing signed code running on the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCode
// SecCodeRef has base type: struct __SecCode *
type SecCodeRef uintptr
// SecGroupTransformRef - A Core Foundation type that represents a container holding a group of transforms.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecGroupTransform
// SecGroupTransformRef has base type: CFTypeRef
type SecGroupTransformRef uintptr
// SecGuestRef - A reference to a guest object, which identifies a particular block of guest code in the context of its code signing host.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecGuestRef
// SecGuestRef has base type: u_int32_t
type SecGuestRef uintptr
// SecIdentityRef - An abstract Core Foundation-type object representing an identity.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentity
// SecIdentityRef has base type: struct __SecIdentity *
type SecIdentityRef uintptr
// SecIdentitySearchRef - Contains information about an identity search.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecIdentitySearch
// SecIdentitySearchRef has base type: struct OpaqueSecIdentitySearchRef *
type SecIdentitySearchRef uintptr
// SecKeyRef - An object that represents a cryptographic key.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKey
// SecKeyRef has base type: struct __SecKey *
type SecKeyRef uintptr
// SecKeyAlgorithm - The algorithms that cryptographic keys enable.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyAlgorithm
// SecKeyAlgorithm has base type: CFStringRef
type SecKeyAlgorithm uintptr
// SecKeychainRef - An opaque type that represents a keychain.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychain
// SecKeychainRef has base type: struct __SecKeychain *
type SecKeychainRef uintptr
// SecKeychainAttributePtr - A pointer to a keychain attribute structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributePtr
// SecKeychainAttributePtr has base type: SecKeychainAttribute *
type SecKeychainAttributePtr uintptr
// SecKeychainAttrType - The keychain attribute type.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttrType
// SecKeychainAttrType has base type: OSType
type SecKeychainAttrType uintptr
// SecKeychainCallback - A customized callback function that keychain services call when a keychain event has occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainCallback
// SecKeychainCallback is a callback function
// C type: int (*)(enum SecKeychainEvent, struct SecKeychainCallbackInfo *, void *)
type SecKeychainCallback = func(SecKeychainEvent, unsafe.Pointer, unsafe.Pointer) int32
// SecKeychainItemRef - An opaque type that represents a keychain item.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItem
// SecKeychainItemRef has base type: struct __SecKeychainItem *
type SecKeychainItemRef uintptr
// SecKeychainSearchRef - An opaque type that contains information about a keychain search.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainSearch
// SecKeychainSearchRef has base type: struct __SecKeychainSearch *
type SecKeychainSearchRef uintptr
// SecKeychainStatus - A value that defines the current status of a keychain.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainStatus
// SecKeychainStatus has base type: UInt32
type SecKeychainStatus uintptr
// SecKeyKeyExchangeParameter - The dictionary keys used to specify Diffie-Hellman key exchange parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyKeyExchangeParameter
// SecKeyKeyExchangeParameter has base type: CFStringRef
type SecKeyKeyExchangeParameter uintptr
// SecPasswordRef - Contains information about a password.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPassword
// SecPasswordRef has base type: struct __SecPassword *
type SecPasswordRef uintptr
// SecPolicyRef - An object that represents a trust policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicy
// SecPolicyRef has base type: struct __SecPolicy *
type SecPolicyRef uintptr
// SecPolicySearchRef - An object that contains information about a policy search.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicySearch
// SecPolicySearchRef has base type: struct OpaquePolicySearchRef *
type SecPolicySearchRef uintptr
// SecPublicKeyHash - A container for a 20-byte public key hash.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPublicKeyHash
// SecPublicKeyHash has base type: UInt8[20]
type SecPublicKeyHash uintptr
// SecRandomRef - An abstract Core Foundation-type object containing information about a random number generator.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRandomRef
// SecRandomRef has base type: const struct __SecRandom *
type SecRandomRef uintptr
// SecRequirementRef - A code requirement object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecRequirement
// SecRequirementRef has base type: struct __SecRequirement *
type SecRequirementRef uintptr
// SecStaticCodeRef - A static code object representing signed code on disk.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecStaticCode
// SecStaticCodeRef has base type: const struct __SecCode *
type SecStaticCodeRef uintptr
// SecTaskRef - The Core Foundation type representing a task.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTask
// SecTaskRef has base type: struct __SecTask *
type SecTaskRef uintptr
// SecTransformRef - A Core Foundation type that represents a security transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransform
// SecTransformRef has base type: CFTypeRef
type SecTransformRef uintptr
// SecTransformAttributeRef - A direct reference to a security transform attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformAttribute
// SecTransformAttributeRef has base type: CFTypeRef
type SecTransformAttributeRef uintptr
// SKIPPED: (* - invalid Go identifier "(*"
// Original type: struct __CFError *(^(*)(const struct __CFString *, const void *, const struct OpaqueSecTransformImplementation *))(void) SecTransformCreateFP
// SecTransformImplementationRef - An opaque pointer to a block that implements an instance of a transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformImplementationRef
// SecTransformImplementationRef has base type: const struct OpaqueSecTransformImplementation *
type SecTransformImplementationRef uintptr
// SecTransformStringOrAttributeRef - A type that may be either a string or an attribute reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformStringOrAttribute
// SecTransformStringOrAttributeRef has base type: CFTypeRef
type SecTransformStringOrAttributeRef uintptr
// SecTrustRef - An object used to evaluate trust.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrust
// SecTrustRef has base type: struct __SecTrust *
type SecTrustRef uintptr
// SecTrustedApplicationRef - An opaque type that contains information about a trusted app.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTrustedApplication
// SecTrustedApplicationRef has base type: struct __SecTrustedApplication *
type SecTrustedApplicationRef uintptr
// SecuritySessionId - A type that contains an authorization session identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecuritySessionId
// SecuritySessionId has base type: UInt32
type SecuritySessionId uintptr
// SSLCipherSuite - A type for storing cipher suite values.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCipherSuite
// SSLCipherSuite has base type: uint16_t
type SSLCipherSuite uintptr
// SSLConnectionRef - A pointer to an opaque I/O connection object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLConnectionRef
// SSLConnectionRef has base type: const void *
type SSLConnectionRef uintptr
// SSLContextRef - An opaque type that represents an SSL session context object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLContext
// SSLContextRef has base type: struct SSLContext *
type SSLContextRef uintptr
// SSLReadFunc - A pointer to a customized read function that secure transport calls to read data from the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLReadFunc
// SSLReadFunc is a callback function
// C type: int (*)(const void *, void *, unsigned long *)
type SSLReadFunc = func(unsafe.Pointer, unsafe.Pointer, uint) int32
// SSLWriteFunc - A pointer to a customized write function that secure transport calls to write data to the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLWriteFunc
// SSLWriteFunc is a callback function
// C type: int (*)(const void *, const void *, unsigned long *)
type SSLWriteFunc = func(unsafe.Pointer, unsafe.Pointer, uint) int32

