// Code generated from Apple documentation for Security. DO NOT EDIT.

package security

// Type aliases and typedefs
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
// CE_CrlNumber has base type: uint32
type CE_CrlNumber uintptr
// CE_DeltaCrl type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DeltaCrl
// CE_DeltaCrl has base type: uint32
type CE_DeltaCrl uintptr
// CMSDecoderRef - An opaque reference to a CMS decoder object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSDecoder
// CMSDecoderRef has base type: struct _CMSDecoder *
type CMSDecoderRef uintptr
// CMSEncoderRef - Opaque reference to a CMS encoder object.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CMSEncoder
// CMSEncoderRef has base type: struct _CMSEncoder *
type CMSEncoderRef uintptr
// CSSM_ACL_AUTHORIZATION_TAG type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_AUTHORIZATION_TAG
// CSSM_ACL_AUTHORIZATION_TAG has base type: sint32
type CSSM_ACL_AUTHORIZATION_TAG uintptr
// CSSM_ACL_EDIT_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ACL_EDIT_MODE
// CSSM_ACL_EDIT_MODE has base type: uint32
type CSSM_ACL_EDIT_MODE uintptr
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
// CSSM_ACL_PREAUTH_TRACKING_STATE has base type: uint32
type CSSM_ACL_PREAUTH_TRACKING_STATE uintptr
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
// CSSM_AC_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_AC_HANDLE
// CSSM_AC_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_AC_HANDLE uintptr
// CSSM_ALGORITHMS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ALGORITHMS
// CSSM_ALGORITHMS has base type: uint32
type CSSM_ALGORITHMS uintptr
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
// CSSM_APPLE_TP_ACTION_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_FLAGS
// CSSM_APPLE_TP_ACTION_FLAGS has base type: uint32
type CSSM_APPLE_TP_ACTION_FLAGS uintptr
// CSSM_APPLE_TP_CRL_OPT_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPT_FLAGS
// CSSM_APPLE_TP_CRL_OPT_FLAGS has base type: uint32
type CSSM_APPLE_TP_CRL_OPT_FLAGS uintptr
// CSSM_ATTACH_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ATTACH_FLAGS
// CSSM_ATTACH_FLAGS has base type: uint32
type CSSM_ATTACH_FLAGS uintptr
// CSSM_ATTRIBUTE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ATTRIBUTE_TYPE
// CSSM_ATTRIBUTE_TYPE has base type: uint32
type CSSM_ATTRIBUTE_TYPE uintptr
// CSSM_BER_TAG type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BER_TAG
// CSSM_BER_TAG has base type: uint8
type CSSM_BER_TAG uintptr
// CSSM_BITMASK type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BITMASK
// CSSM_BITMASK has base type: uint32
type CSSM_BITMASK uintptr
// CSSM_BOOL type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_BOOL
// CSSM_BOOL has base type: sint32
type CSSM_BOOL uintptr
// CSSM_CALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CALLOC
// CSSM_CALLOC has base type: void *(*)(unsigned int, unsigned long, void *)
type CSSM_CALLOC uintptr
// CSSM_CC_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CC_HANDLE
// CSSM_CC_HANDLE has base type: CSSM_LONG_HANDLE
type CSSM_CC_HANDLE uintptr
// CSSM_CERTGROUP_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE
// CSSM_CERTGROUP_TYPE has base type: uint32
type CSSM_CERTGROUP_TYPE uintptr
// CSSM_CERTGROUP_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERTGROUP_TYPE_PTR
// CSSM_CERTGROUP_TYPE_PTR has base type: uint32 *
type CSSM_CERTGROUP_TYPE_PTR uintptr
// CSSM_CERT_BUNDLE_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_ENCODING
// CSSM_CERT_BUNDLE_ENCODING has base type: uint32
type CSSM_CERT_BUNDLE_ENCODING uintptr
// CSSM_CERT_BUNDLE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_BUNDLE_TYPE
// CSSM_CERT_BUNDLE_TYPE has base type: uint32
type CSSM_CERT_BUNDLE_TYPE uintptr
// CSSM_CERT_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING
// CSSM_CERT_ENCODING has base type: uint32
type CSSM_CERT_ENCODING uintptr
// CSSM_CERT_ENCODING_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_ENCODING_PTR
// CSSM_CERT_ENCODING_PTR has base type: uint32 *
type CSSM_CERT_ENCODING_PTR uintptr
// CSSM_CERT_PARSE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT
// CSSM_CERT_PARSE_FORMAT has base type: uint32
type CSSM_CERT_PARSE_FORMAT uintptr
// CSSM_CERT_PARSE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_PARSE_FORMAT_PTR
// CSSM_CERT_PARSE_FORMAT_PTR has base type: uint32 *
type CSSM_CERT_PARSE_FORMAT_PTR uintptr
// CSSM_CERT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE
// CSSM_CERT_TYPE has base type: uint32
type CSSM_CERT_TYPE uintptr
// CSSM_CERT_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CERT_TYPE_PTR
// CSSM_CERT_TYPE_PTR has base type: uint32 *
type CSSM_CERT_TYPE_PTR uintptr
// CSSM_CL_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_HANDLE
// CSSM_CL_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_CL_HANDLE uintptr
// CSSM_CL_TEMPLATE_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CL_TEMPLATE_TYPE
// CSSM_CL_TEMPLATE_TYPE has base type: uint32
type CSSM_CL_TEMPLATE_TYPE uintptr
// CSSM_CONTEXT_EVENT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_EVENT
// CSSM_CONTEXT_EVENT has base type: uint32
type CSSM_CONTEXT_EVENT uintptr
// CSSM_CONTEXT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CONTEXT_TYPE
// CSSM_CONTEXT_TYPE has base type: uint32
type CSSM_CONTEXT_TYPE uintptr
// CSSM_CRLGROUP_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE
// CSSM_CRLGROUP_TYPE has base type: uint32
type CSSM_CRLGROUP_TYPE uintptr
// CSSM_CRLGROUP_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRLGROUP_TYPE_PTR
// CSSM_CRLGROUP_TYPE_PTR has base type: uint32 *
type CSSM_CRLGROUP_TYPE_PTR uintptr
// CSSM_CRL_ENCODING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING
// CSSM_CRL_ENCODING has base type: uint32
type CSSM_CRL_ENCODING uintptr
// CSSM_CRL_ENCODING_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_ENCODING_PTR
// CSSM_CRL_ENCODING_PTR has base type: uint32 *
type CSSM_CRL_ENCODING_PTR uintptr
// CSSM_CRL_PARSE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT
// CSSM_CRL_PARSE_FORMAT has base type: uint32
type CSSM_CRL_PARSE_FORMAT uintptr
// CSSM_CRL_PARSE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_PARSE_FORMAT_PTR
// CSSM_CRL_PARSE_FORMAT_PTR has base type: uint32 *
type CSSM_CRL_PARSE_FORMAT_PTR uintptr
// CSSM_CRL_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE
// CSSM_CRL_TYPE has base type: uint32
type CSSM_CRL_TYPE uintptr
// CSSM_CRL_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CRL_TYPE_PTR
// CSSM_CRL_TYPE_PTR has base type: uint32 *
type CSSM_CRL_TYPE_PTR uintptr
// CSSM_CSPTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSPTYPE
// CSSM_CSPTYPE has base type: uint32
type CSSM_CSPTYPE uintptr
// CSSM_CSP_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_FLAGS
// CSSM_CSP_FLAGS has base type: uint32
type CSSM_CSP_FLAGS uintptr
// CSSM_CSP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_HANDLE
// CSSM_CSP_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_CSP_HANDLE uintptr
// CSSM_CSP_READER_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_CSP_READER_FLAGS
// CSSM_CSP_READER_FLAGS has base type: uint32
type CSSM_CSP_READER_FLAGS uintptr
// CSSM_DB_ACCESS_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE
// CSSM_DB_ACCESS_TYPE has base type: uint32
type CSSM_DB_ACCESS_TYPE uintptr
// CSSM_DB_ACCESS_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ACCESS_TYPE_PTR
// CSSM_DB_ACCESS_TYPE_PTR has base type: uint32 *
type CSSM_DB_ACCESS_TYPE_PTR uintptr
// CSSM_DB_ATTRIBUTE_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT
// CSSM_DB_ATTRIBUTE_FORMAT has base type: uint32
type CSSM_DB_ATTRIBUTE_FORMAT uintptr
// CSSM_DB_ATTRIBUTE_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_FORMAT_PTR
// CSSM_DB_ATTRIBUTE_FORMAT_PTR has base type: uint32 *
type CSSM_DB_ATTRIBUTE_FORMAT_PTR uintptr
// CSSM_DB_ATTRIBUTE_NAME_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT
// CSSM_DB_ATTRIBUTE_NAME_FORMAT has base type: uint32
type CSSM_DB_ATTRIBUTE_NAME_FORMAT uintptr
// CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR
// CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR has base type: uint32 *
type CSSM_DB_ATTRIBUTE_NAME_FORMAT_PTR uintptr
// CSSM_DB_CONJUNCTIVE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_CONJUNCTIVE
// CSSM_DB_CONJUNCTIVE has base type: uint32
type CSSM_DB_CONJUNCTIVE uintptr
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
// CSSM_DB_INDEXED_DATA_LOCATION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_INDEXED_DATA_LOCATION
// CSSM_DB_INDEXED_DATA_LOCATION has base type: uint32
type CSSM_DB_INDEXED_DATA_LOCATION uintptr
// CSSM_DB_INDEX_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_INDEX_TYPE
// CSSM_DB_INDEX_TYPE has base type: uint32
type CSSM_DB_INDEX_TYPE uintptr
// CSSM_DB_MODIFY_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_MODIFY_MODE
// CSSM_DB_MODIFY_MODE has base type: uint32
type CSSM_DB_MODIFY_MODE uintptr
// CSSM_DB_OPERATOR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR
// CSSM_DB_OPERATOR has base type: uint32
type CSSM_DB_OPERATOR uintptr
// CSSM_DB_OPERATOR_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_OPERATOR_PTR
// CSSM_DB_OPERATOR_PTR has base type: uint32 *
type CSSM_DB_OPERATOR_PTR uintptr
// CSSM_DB_RECORDTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_RECORDTYPE
// CSSM_DB_RECORDTYPE has base type: uint32
type CSSM_DB_RECORDTYPE uintptr
// CSSM_DB_RETRIEVAL_MODES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DB_RETRIEVAL_MODES
// CSSM_DB_RETRIEVAL_MODES has base type: uint32
type CSSM_DB_RETRIEVAL_MODES uintptr
// CSSM_DLTYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DLTYPE
// CSSM_DLTYPE has base type: uint32
type CSSM_DLTYPE uintptr
// CSSM_DLTYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_DLTYPE_PTR
// CSSM_DLTYPE_PTR has base type: uint32 *
type CSSM_DLTYPE_PTR uintptr
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
// CSSM_ENCRYPT_MODE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_ENCRYPT_MODE
// CSSM_ENCRYPT_MODE has base type: uint32
type CSSM_ENCRYPT_MODE uintptr
// CSSM_EVIDENCE_FORM type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_EVIDENCE_FORM
// CSSM_EVIDENCE_FORM has base type: uint32
type CSSM_EVIDENCE_FORM uintptr
// CSSM_FREE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_FREE
// CSSM_FREE has base type: void (*)(void *, void *)
type CSSM_FREE uintptr
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
// CSSM_HEADERVERSION has base type: uint32
type CSSM_HEADERVERSION uintptr
// CSSM_INTPTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_INTPTR
// CSSM_INTPTR has base type: intptr_t
type CSSM_INTPTR uintptr
// CSSM_KEYATTR_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYATTR_FLAGS
// CSSM_KEYATTR_FLAGS has base type: uint32
type CSSM_KEYATTR_FLAGS uintptr
// CSSM_KEYBLOB_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_FORMAT
// CSSM_KEYBLOB_FORMAT has base type: uint32
type CSSM_KEYBLOB_FORMAT uintptr
// CSSM_KEYBLOB_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYBLOB_TYPE
// CSSM_KEYBLOB_TYPE has base type: uint32
type CSSM_KEYBLOB_TYPE uintptr
// CSSM_KEYCLASS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYCLASS
// CSSM_KEYCLASS has base type: uint32
type CSSM_KEYCLASS uintptr
// CSSM_KEYUSE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KEYUSE
// CSSM_KEYUSE has base type: uint32
type CSSM_KEYUSE uintptr
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
// CSSM_KRSP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KRSP_HANDLE
// CSSM_KRSP_HANDLE has base type: uint32
type CSSM_KRSP_HANDLE uintptr
// CSSM_KR_POLICY_FLAGS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_FLAGS
// CSSM_KR_POLICY_FLAGS has base type: uint32
type CSSM_KR_POLICY_FLAGS uintptr
// CSSM_KR_POLICY_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_KR_POLICY_TYPE
// CSSM_KR_POLICY_TYPE has base type: uint32
type CSSM_KR_POLICY_TYPE uintptr
// CSSM_LIST_ELEMENT_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_PTR
// CSSM_LIST_ELEMENT_PTR has base type: struct cssm_list_element *
type CSSM_LIST_ELEMENT_PTR uintptr
// CSSM_LIST_ELEMENT_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE
// CSSM_LIST_ELEMENT_TYPE has base type: uint32
type CSSM_LIST_ELEMENT_TYPE uintptr
// CSSM_LIST_ELEMENT_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_ELEMENT_TYPE_PTR
// CSSM_LIST_ELEMENT_TYPE_PTR has base type: uint32 *
type CSSM_LIST_ELEMENT_TYPE_PTR uintptr
// CSSM_LIST_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE
// CSSM_LIST_TYPE has base type: uint32
type CSSM_LIST_TYPE uintptr
// CSSM_LIST_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LIST_TYPE_PTR
// CSSM_LIST_TYPE_PTR has base type: uint32 *
type CSSM_LIST_TYPE_PTR uintptr
// CSSM_LONG_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE
// CSSM_LONG_HANDLE has base type: uint64
type CSSM_LONG_HANDLE uintptr
// CSSM_LONG_HANDLE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_LONG_HANDLE_PTR
// CSSM_LONG_HANDLE_PTR has base type: uint64 *
type CSSM_LONG_HANDLE_PTR uintptr
// CSSM_MALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MALLOC
// CSSM_MALLOC has base type: void *(*)(unsigned long, void *)
type CSSM_MALLOC uintptr
// CSSM_MANAGER_EVENT_TYPES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MANAGER_EVENT_TYPES
// CSSM_MANAGER_EVENT_TYPES has base type: uint32
type CSSM_MANAGER_EVENT_TYPES uintptr
// CSSM_MODULE_EVENT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_MODULE_EVENT
// CSSM_MODULE_EVENT has base type: uint32
type CSSM_MODULE_EVENT uintptr
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
// CSSM_NET_ADDRESS_TYPE has base type: uint32
type CSSM_NET_ADDRESS_TYPE uintptr
// CSSM_NET_PROTOCOL type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_NET_PROTOCOL
// CSSM_NET_PROTOCOL has base type: uint32
type CSSM_NET_PROTOCOL uintptr
// CSSM_PADDING type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PADDING
// CSSM_PADDING has base type: uint32
type CSSM_PADDING uintptr
// CSSM_PKCS5_PBKDF2_PRF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS5_PBKDF2_PRF
// CSSM_PKCS5_PBKDF2_PRF has base type: uint32
type CSSM_PKCS5_PBKDF2_PRF uintptr
// CSSM_PKCS_OAEP_MGF type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_MGF
// CSSM_PKCS_OAEP_MGF has base type: uint32
type CSSM_PKCS_OAEP_MGF uintptr
// CSSM_PKCS_OAEP_PSOURCE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PKCS_OAEP_PSOURCE
// CSSM_PKCS_OAEP_PSOURCE has base type: uint32
type CSSM_PKCS_OAEP_PSOURCE uintptr
// CSSM_PRIVILEGE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE
// CSSM_PRIVILEGE has base type: uint64
type CSSM_PRIVILEGE uintptr
// CSSM_PRIVILEGE_SCOPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PRIVILEGE_SCOPE
// CSSM_PRIVILEGE_SCOPE has base type: uint32
type CSSM_PRIVILEGE_SCOPE uintptr
// CSSM_PROC_ADDR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_PROC_ADDR
// CSSM_PROC_ADDR has base type: void (*)(void)
type CSSM_PROC_ADDR uintptr
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
// CSSM_QUERY_FLAGS has base type: uint32
type CSSM_QUERY_FLAGS uintptr
// CSSM_REALLOC type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_REALLOC
// CSSM_REALLOC has base type: void *(*)(void *, unsigned long, void *)
type CSSM_REALLOC uintptr
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
// CSSM_SC_FLAGS has base type: uint32
type CSSM_SC_FLAGS uintptr
// CSSM_SERVICE_MASK type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_SERVICE_MASK
// CSSM_SERVICE_MASK has base type: uint32
type CSSM_SERVICE_MASK uintptr
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
// CSSM_TP_ACTION has base type: uint32
type CSSM_TP_ACTION uintptr
// CSSM_TP_APPLE_CERT_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_CERT_STATUS
// CSSM_TP_APPLE_CERT_STATUS has base type: uint32
type CSSM_TP_APPLE_CERT_STATUS uintptr
// CSSM_TP_AUTHORITY_REQUEST_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE
// CSSM_TP_AUTHORITY_REQUEST_TYPE has base type: uint32
type CSSM_TP_AUTHORITY_REQUEST_TYPE uintptr
// CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR
// CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR has base type: uint32 *
type CSSM_TP_AUTHORITY_REQUEST_TYPE_PTR uintptr
// CSSM_TP_CERTCHANGE_ACTION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_ACTION
// CSSM_TP_CERTCHANGE_ACTION has base type: uint32
type CSSM_TP_CERTCHANGE_ACTION uintptr
// CSSM_TP_CERTCHANGE_REASON type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_REASON
// CSSM_TP_CERTCHANGE_REASON has base type: uint32
type CSSM_TP_CERTCHANGE_REASON uintptr
// CSSM_TP_CERTCHANGE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTCHANGE_STATUS
// CSSM_TP_CERTCHANGE_STATUS has base type: uint32
type CSSM_TP_CERTCHANGE_STATUS uintptr
// CSSM_TP_CERTISSUE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTISSUE_STATUS
// CSSM_TP_CERTISSUE_STATUS has base type: uint32
type CSSM_TP_CERTISSUE_STATUS uintptr
// CSSM_TP_CERTNOTARIZE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTNOTARIZE_STATUS
// CSSM_TP_CERTNOTARIZE_STATUS has base type: uint32
type CSSM_TP_CERTNOTARIZE_STATUS uintptr
// CSSM_TP_CERTRECLAIM_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTRECLAIM_STATUS
// CSSM_TP_CERTRECLAIM_STATUS has base type: uint32
type CSSM_TP_CERTRECLAIM_STATUS uintptr
// CSSM_TP_CERTVERIFY_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CERTVERIFY_STATUS
// CSSM_TP_CERTVERIFY_STATUS has base type: uint32
type CSSM_TP_CERTVERIFY_STATUS uintptr
// CSSM_TP_CONFIRM_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS
// CSSM_TP_CONFIRM_STATUS has base type: uint32
type CSSM_TP_CONFIRM_STATUS uintptr
// CSSM_TP_CONFIRM_STATUS_PTR type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CONFIRM_STATUS_PTR
// CSSM_TP_CONFIRM_STATUS_PTR has base type: uint32 *
type CSSM_TP_CONFIRM_STATUS_PTR uintptr
// CSSM_TP_CRLISSUE_STATUS type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_CRLISSUE_STATUS
// CSSM_TP_CRLISSUE_STATUS has base type: uint32
type CSSM_TP_CRLISSUE_STATUS uintptr
// CSSM_TP_FORM_TYPE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_FORM_TYPE
// CSSM_TP_FORM_TYPE has base type: uint32
type CSSM_TP_FORM_TYPE uintptr
// CSSM_TP_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_HANDLE
// CSSM_TP_HANDLE has base type: CSSM_MODULE_HANDLE
type CSSM_TP_HANDLE uintptr
// CSSM_TP_SERVICES type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_SERVICES
// CSSM_TP_SERVICES has base type: uint32
type CSSM_TP_SERVICES uintptr
// CSSM_TP_STOP_ON type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_STOP_ON
// CSSM_TP_STOP_ON has base type: uint32
type CSSM_TP_STOP_ON uintptr
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
// CSSM_X509EXT_DATA_FORMAT type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509EXT_DATA_FORMAT
// CSSM_X509EXT_DATA_FORMAT has base type: enum extension_data_format
type CSSM_X509EXT_DATA_FORMAT uintptr
// CSSM_X509_OPTION type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_X509_OPTION
// CSSM_X509_OPTION has base type: CSSM_BOOL
type CSSM_X509_OPTION uintptr
// MDS_HANDLE type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/MDS_HANDLE
// MDS_HANDLE has base type: CSSM_DL_HANDLE
type MDS_HANDLE uintptr
// SSLCipherSuite - A type for storing cipher suite values.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLCipherSuite
// SSLCipherSuite has base type: uint32_t
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
// SSLReadFunc has base type: int (*)(const void *, void *, unsigned long *)
type SSLReadFunc uintptr
// SSLWriteFunc - A pointer to a customized write function that secure transport calls to write data to the connection.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SSLWriteFunc
// SSLWriteFunc has base type: int (*)(const void *, const void *, unsigned long *)
type SSLWriteFunc uintptr
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
// SecAsn1Template - A structure that defines one element of a BER or DER encoding.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Template
// SecAsn1Template has base type: struct SecAsn1Template_struct
type SecAsn1Template uintptr
// SecAsn1TemplateChooser - Dynamically provides the sub-template to use during encode or decode.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1TemplateChooser
// SecAsn1TemplateChooser has base type: const struct SecAsn1Template_struct *(void *, unsigned char, const char *, unsigned long, void *)
type SecAsn1TemplateChooser uintptr
// SecAsn1TemplateChooserPtr - A pointer to the template chooser function.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1TemplateChooserPtr
// SecAsn1TemplateChooserPtr has base type: const struct SecAsn1Template_struct *(*)(void *, unsigned char, const char *, unsigned long, void *)
type SecAsn1TemplateChooserPtr uintptr
// SecCECrlReason type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCECrlReason
// SecCECrlReason has base type: uint32_t
type SecCECrlReason uintptr
// SecCEKeyUsage type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEKeyUsage
// SecCEKeyUsage has base type: uint16_t
type SecCEKeyUsage uintptr
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
// SecKeychainItemRef - An opaque type that represents a keychain item.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainItem
// SecKeychainItemRef has base type: struct __SecKeychainItem *
type SecKeychainItemRef uintptr
// SecPolicyRef - An object that represents a trust policy.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecPolicy
// SecPolicyRef has base type: struct __SecPolicy *
type SecPolicyRef uintptr
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
// (* - A pointer to a function that creates a new instance of a custom transform.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/SecTransformCreateFP
// (* is a block type: struct __CFError *(^(*)(const struct __CFString *, const void *, const struct OpaqueSecTransformImplementation *))(void) SecTransformCreateFP
// Block types are not yet fully supported in Go bindings
type (* uintptr
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
// sec_certificate_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_certificate_t
// sec_certificate_t has base type: NSObject<OS_sec_certificate> *
type sec_certificate_t uintptr
// sec_identity_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_identity_t
// sec_identity_t has base type: NSObject<OS_sec_identity> *
type sec_identity_t uintptr
// sec_object_t - A   is a generic, ARC-able type wrapper for common CoreFoundation Security types.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_object_t
// sec_object_t has base type: NSObject<OS_sec_object> *
type sec_object_t uintptr
// sec_protocol_metadata_t - A   instance conatins read-only properties of a connected and configured   security protocol. Clients use this object to read information about a protocol instance. Properties   include, for example, the negotiated TLS version, ciphersuite, and peer certificates.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_metadata_t
// sec_protocol_metadata_t has base type: NSObject<OS_sec_protocol_metadata> *
type sec_protocol_metadata_t uintptr
// sec_protocol_options_t - A   instance is a container of options for security protocol instances,   such as TLS. Protocol options are used to configure security protocols in the network stack.   For example, clients may set the maximum and minimum allowed TLS versions through protocol   options.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_protocol_options_t
// sec_protocol_options_t has base type: NSObject<OS_sec_protocol_options> *
type sec_protocol_options_t uintptr
// sec_trust_t - These are os_object compatible and ARC-able wrappers around existing CoreFoundation   Security types, including: SecTrustRef, SecIdentityRef, and SecCertificateRef. They allow   clients to use these types in os_object-type APIs and data structures. The underlying   CoreFoundation types may be extracted and used by clients as needed.
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sec_trust_t
// sec_trust_t has base type: NSObject<OS_sec_trust> *
type sec_trust_t uintptr
// sint16 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sint16
// sint16 has base type: int16_t
type sint16 uintptr
// sint32 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sint32
// sint32 has base type: int32_t
type sint32 uintptr
// sint64 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sint64
// sint64 has base type: int64_t
type sint64 uintptr
// sint8 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/sint8
// sint8 has base type: int8_t
type sint8 uintptr
// uint16 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/uint16
// uint16 has base type: uint16_t
type uint16 uintptr
// uint32 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/uint32
// uint32 has base type: uint32_t
type uint32 uintptr
// uint64 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/uint64
// uint64 has base type: uint64_t
type uint64 uintptr
// uint8 type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Security/uint8
// uint8 has base type: uint8_t
type uint8 uintptr

