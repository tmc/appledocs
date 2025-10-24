// Code generated from Apple documentation for Security. DO NOT EDIT.

package security
import (
	"unsafe"
)


// C struct types
// AuthorizationCallbacks - The interface implemented by the Security Server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationCallbacks
type AuthorizationCallbacks struct {
	DidDeactivate unsafe.Pointer // Reports the successful deactivation of an authorization mechanism.
	SetContextValue unsafe.Pointer // Stores data collected during authorization as a key-value pair.
	SetResult unsafe.Pointer // Returns the result of an authorization operation.
	GetContextValue unsafe.Pointer // Reads a value collected during authorization.
	RequestInterrupt unsafe.Pointer // Requests the authorization engine to interrupt the currently active authorization mechanism.
	SetHintValue unsafe.Pointer // Stores data needed during authorization as a key-value pair.
	GetArguments unsafe.Pointer // Reads the arguments for this authorization mechanism from the authorization policy database.
	GetHintValue unsafe.Pointer // Reads a value stored by the plug-in authorization mechanism.
	DidDeactivate unsafe.Pointer // Reports the successful deactivation of an authorization mechanism.
	GetArguments unsafe.Pointer // Reads the arguments for this authorization mechanism from the authorization policy database.
	GetContextValue unsafe.Pointer // Reads a value collected during authorization.
	GetHintValue unsafe.Pointer // Reads a value stored by the plug-in authorization mechanism.
	GetImmutableHintValue unsafe.Pointer // Reads an immutable value stored by the plug-in authorization mechanism.
	GetLAContext unsafe.Pointer // Constructs a local authentication context.
	GetSessionId unsafe.Pointer // Reads the session ID.
	GetTKTokenWatcher unsafe.Pointer // Constructs a token watcher.
	GetTokenIdentities unsafe.Pointer // Returns an array of identities available on tokens.
	RemoveContextValue unsafe.Pointer // Removes a value collected during authorization.
	RemoveHintValue unsafe.Pointer // Removes a value stored by the plug-in authorization mechanism.
	RequestInterrupt unsafe.Pointer // Requests the authorization engine to interrupt the currently active authorization mechanism.
	SetContextValue unsafe.Pointer // Stores data collected during authorization as a key-value pair.
	SetHintValue unsafe.Pointer // Stores data needed during authorization as a key-value pair.
	SetResult unsafe.Pointer // Returns the result of an authorization operation.
	Version unsafe.Pointer // The engine callback version.
}// AuthorizationExternalForm - The external representation of an authorization reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationExternalForm
type AuthorizationExternalForm struct {
	Bytes unsafe.Pointer // An array of characters representing the external form of an authorization reference.
}// AuthorizationItem - A structure containing information about an authorization right or the authorization environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItem
type AuthorizationItem struct {
	Flags unsafe.Pointer // Reserved option bits.
	Name AuthorizationString // The required name of the authorization right or environment data.
	Value unsafe.Pointer // A pointer to information pertaining to the name field.
	ValueLength uintptr // The number of bytes in the value field.
}// AuthorizationItemSet - A structure containing a set of authorization items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationItemSet
type AuthorizationItemSet struct {
	Count unsafe.Pointer // The number of elements in the   array.
	Items unsafe.Pointer // A pointer to an array of authorization items.
}// AuthorizationValue - A structure used to pass data between the authorization engine and the plug-in mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValue
type AuthorizationValue struct {
	Data unsafe.Pointer
	Length uintptr
}// AuthorizationValueVector - A structure used to pass arguments from the authorization policy database to the authorization mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationValueVector
type AuthorizationValueVector struct {
	Count unsafe.Pointer
	Values unsafe.Pointer
}// _CE_AccessDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AccessDescription
type _CE_AccessDescription struct {
}// _CE_AuthorityInfoAccess
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityInfoAccess
type _CE_AuthorityInfoAccess struct {
}// _CE_AuthorityKeyID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_AuthorityKeyID
type _CE_AuthorityKeyID struct {
}// _CE_BasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_BasicConstraints
type _CE_BasicConstraints struct {
}// _CE_CRLDistPointsSyntax
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistPointsSyntax
type _CE_CRLDistPointsSyntax struct {
}// _CE_CRLDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CRLDistributionPoint
type _CE_CRLDistributionPoint struct {
}// _CE_CertPolicies
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_CertPolicies
type _CE_CertPolicies struct {
}// _CE_DataAndType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DataAndType
type _CE_DataAndType struct {
}// _CE_DistributionPointName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_DistributionPointName
type _CE_DistributionPointName struct {
}// _CE_ExtendedKeyUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_ExtendedKeyUsage-c.struct
type _CE_ExtendedKeyUsage struct {
	NumPurposes Uint32
	Purposes unsafe.Pointer
}// _CE_GeneralName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralName
type _CE_GeneralName struct {
}// _CE_GeneralNames
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralNames
type _CE_GeneralNames struct {
}// _CE_GeneralSubtree
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtree
type _CE_GeneralSubtree struct {
}// _CE_GeneralSubtrees
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_GeneralSubtrees
type _CE_GeneralSubtrees struct {
}// _CE_IssuingDistributionPoint
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_IssuingDistributionPoint
type _CE_IssuingDistributionPoint struct {
}// _CE_NameConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_NameConstraints
type _CE_NameConstraints struct {
}// _CE_OtherName
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_OtherName
type _CE_OtherName struct {
}// _CE_PolicyConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyConstraints
type _CE_PolicyConstraints struct {
}// _CE_PolicyInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyInformation
type _CE_PolicyInformation struct {
}// _CE_PolicyMapping
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMapping
type _CE_PolicyMapping struct {
}// _CE_PolicyMappings
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyMappings
type _CE_PolicyMappings struct {
}// _CE_PolicyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_PolicyQualifierInfo
type _CE_PolicyQualifierInfo struct {
}// _CE_QC_Statement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statement
type _CE_QC_Statement struct {
}// _CE_QC_Statements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_QC_Statements
type _CE_QC_Statements struct {
}// _CE_SemanticsInformation
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CE_SemanticsInformation
type _CE_SemanticsInformation struct {
}// CSSM_APPLE_CL_CSR_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_CL_CSR_REQUEST
type CSSM_APPLE_CL_CSR_REQUEST struct {
}// CSSM_APPLE_TP_ACTION_DATA
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_ACTION_DATA
type CSSM_APPLE_TP_ACTION_DATA struct {
}// CSSM_APPLE_TP_CERT_REQUEST
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CERT_REQUEST
type CSSM_APPLE_TP_CERT_REQUEST struct {
}// CSSM_APPLE_TP_CRL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_CRL_OPTIONS
type CSSM_APPLE_TP_CRL_OPTIONS struct {
}// CSSM_APPLE_TP_NAME_OID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_NAME_OID
type CSSM_APPLE_TP_NAME_OID struct {
}// CSSM_APPLE_TP_SMIME_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SMIME_OPTIONS
type CSSM_APPLE_TP_SMIME_OPTIONS struct {
}// CSSM_APPLE_TP_SSL_OPTIONS
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_APPLE_TP_SSL_OPTIONS
type CSSM_APPLE_TP_SSL_OPTIONS struct {
}// CSSM_TP_APPLE_EVIDENCE_HEADER
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_EVIDENCE_HEADER
type CSSM_TP_APPLE_EVIDENCE_HEADER struct {
}// CSSM_TP_APPLE_EVIDENCE_INFO
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TP_APPLE_EVIDENCE_INFO
type CSSM_TP_APPLE_EVIDENCE_INFO struct {
}// CSSM_TUPLE
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/CSSM_TUPLE
type CSSM_TUPLE struct {
}// SecAsn1AlgId - A structure identifying an ASN.1 algorithm by its OID, and its corresponding parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1AlgId
type SecAsn1AlgId struct {
}// SecAsn1PubKeyInfo - A structure containing a public key and its associated algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1PubKeyInfo
type SecAsn1PubKeyInfo struct {
}// SecAsn1Template_struct - A structure that defines one element of a BER or DER encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecAsn1Template_struct
type SecAsn1Template_struct struct {
}// SecCEBasicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEBasicConstraints
type SecCEBasicConstraints struct {
	Present bool
}// SecCEPolicyConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecCEPolicyConstraints
type SecCEPolicyConstraints struct {
	Critical bool
	InhibitPolicyMapping uint32
	RequireExplicitPolicy uint32
}// SecItemImportExportKeyParameters - The import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecItemImportExportKeyParameters
type SecItemImportExportKeyParameters struct {
}// SecKeyImportExportParameters - The legacy import/export parameter structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeyImportExportParameters
type SecKeyImportExportParameters struct {
}// SecKeychainAttribute - A structure that holds a single keychain attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttribute
type SecKeychainAttribute struct {
}// SecKeychainAttributeInfo - A structure that represents an attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeInfo
type SecKeychainAttributeInfo struct {
}// SecKeychainAttributeList - A list of keychain attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/SecKeychainAttributeList
type SecKeychainAttributeList struct {
}// AuthorizationPluginInterface - The interface that must be implemented by your plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/AuthorizationPluginInterface
type AuthorizationPluginInterface struct {
	MechanismDestroy unsafe.Pointer // Destroys an authorization mechanism.
	MechanismInvoke unsafe.Pointer // Invokes an authorization mechanism to perform an authorization operation.
	MechanismDeactivate unsafe.Pointer // Deactivates an authorization mechanism.
	MechanismCreate unsafe.Pointer // Creates an authorization mechanism.
	MechanismCreate unsafe.Pointer // Creates an authorization mechanism.
	MechanismDeactivate unsafe.Pointer // Deactivates an authorization mechanism.
	MechanismDestroy unsafe.Pointer // Destroys an authorization mechanism.
	MechanismInvoke unsafe.Pointer // Invokes an authorization mechanism to perform an authorization operation.
	PluginDestroy unsafe.Pointer // Notifies the plug-in that it is about to be unloaded.
	Version unsafe.Pointer // The plug-in interface version.
}// cssm_access_credentials
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_access_credentials-c.struct
type cssm_access_credentials struct {
}// cssm_acl_edit
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_edit-c.struct
type cssm_acl_edit struct {
}// cssm_acl_entry_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_info-c.struct
type cssm_acl_entry_info struct {
}// cssm_acl_entry_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_input-c.struct
type cssm_acl_entry_input struct {
}// cssm_acl_entry_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_entry_prototype-c.struct
type cssm_acl_entry_prototype struct {
}// cssm_acl_keychain_prompt_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_keychain_prompt_selector-swift.struct
type cssm_acl_keychain_prompt_selector struct {
}// cssm_acl_owner_prototype
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_owner_prototype-c.struct
type cssm_acl_owner_prototype struct {
}// cssm_acl_process_subject_selector
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_process_subject_selector-swift.struct
type cssm_acl_process_subject_selector struct {
}// cssm_acl_validity_period
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_acl_validity_period-c.struct
type cssm_acl_validity_period struct {
}// cssm_applecspdl_db_change_password_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_change_password_parameters-swift.struct
type cssm_applecspdl_db_change_password_parameters struct {
}// cssm_applecspdl_db_is_locked_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_is_locked_parameters-swift.struct
type cssm_applecspdl_db_is_locked_parameters struct {
}// cssm_applecspdl_db_settings_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_applecspdl_db_settings_parameters-swift.struct
type cssm_applecspdl_db_settings_parameters struct {
}// cssm_appledl_open_parameters
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_appledl_open_parameters-swift.struct
type cssm_appledl_open_parameters struct {
}// cssm_authorizationgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_authorizationgroup-swift.struct
type cssm_authorizationgroup struct {
}// cssm_base_certs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_base_certs-c.struct
type cssm_base_certs struct {
}// cssm_cert_bundle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle-c.struct
type cssm_cert_bundle struct {
}// cssm_cert_bundle_header
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_bundle_header-c.struct
type cssm_cert_bundle_header struct {
}// cssm_cert_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_cert_pair-c.struct
type cssm_cert_pair struct {
}// cssm_certgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_certgroup-c.struct
type cssm_certgroup struct {
}// cssm_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context-c.struct
type cssm_context struct {
}// cssm_context_attribute
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_context_attribute-c.struct
type cssm_context_attribute struct {
}// cssm_crl_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crl_pair-c.struct
type cssm_crl_pair struct {
}// cssm_crlgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crlgroup-c.struct
type cssm_crlgroup struct {
}// cssm_crypto_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_crypto_data-c.struct
type cssm_crypto_data struct {
}// cssm_csp_operational_statistics
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_csp_operational_statistics-swift.struct
type cssm_csp_operational_statistics struct {
}// cssm_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_data-swift.struct
type cssm_data struct {
}// cssm_date
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_date-swift.struct
type cssm_date struct {
}// cssm_db_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_data-c.struct
type cssm_db_attribute_data struct {
}// cssm_db_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_attribute_info-c.struct
type cssm_db_attribute_info struct {
}// cssm_db_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_index_info-c.struct
type cssm_db_index_info struct {
}// cssm_db_parsing_module_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_parsing_module_info-c.struct
type cssm_db_parsing_module_info struct {
}// cssm_db_record_attribute_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_data-c.struct
type cssm_db_record_attribute_data struct {
}// cssm_db_record_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_attribute_info-c.struct
type cssm_db_record_attribute_info struct {
}// cssm_db_record_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_record_index_info-c.struct
type cssm_db_record_index_info struct {
}// cssm_db_schema_attribute_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_attribute_info-c.struct
type cssm_db_schema_attribute_info struct {
}// cssm_db_schema_index_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_schema_index_info-swift.struct
type cssm_db_schema_index_info struct {
}// cssm_db_unique_record
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_db_unique_record-c.struct
type cssm_db_unique_record struct {
}// cssm_dbinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dbinfo-c.struct
type cssm_dbinfo struct {
}// cssm_dl_db_handle
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_db_handle-swift.struct
type cssm_dl_db_handle struct {
}// cssm_dl_db_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_db_list-c.struct
type cssm_dl_db_list struct {
}// cssm_dl_pkcs11_attributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_dl_pkcs11_attributes
type cssm_dl_pkcs11_attributes struct {
}// cssm_encoded_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_encoded_cert-c.struct
type cssm_encoded_cert struct {
}// cssm_encoded_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_encoded_crl-c.struct
type cssm_encoded_crl struct {
}// cssm_evidence
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_evidence-c.struct
type cssm_evidence struct {
}// cssm_field
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_field-c.struct
type cssm_field struct {
}// cssm_fieldgroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_fieldgroup-c.struct
type cssm_fieldgroup struct {
}// cssm_func_name_addr
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_func_name_addr-swift.struct
type cssm_func_name_addr struct {
}// cssm_guid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_guid-swift.struct
type cssm_guid struct {
}// cssm_kea_derive_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kea_derive_params-c.struct
type cssm_kea_derive_params struct {
}// cssm_key
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key-c.struct
type cssm_key struct {
}// cssm_key_size
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_key_size-swift.struct
type cssm_key_size struct {
}// cssm_keyheader
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_keyheader-c.struct
type cssm_keyheader struct {
}// cssm_kr_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_name-swift.struct
type cssm_kr_name struct {
}// cssm_kr_policy_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_info-c.struct
type cssm_kr_policy_info struct {
}// cssm_kr_policy_list_item
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_policy_list_item-c.struct
type cssm_kr_policy_list_item struct {
}// cssm_kr_profile
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_profile-c.struct
type cssm_kr_profile struct {
}// cssm_kr_wrappedproductinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_kr_wrappedproductinfo
type cssm_kr_wrappedproductinfo struct {
}// cssm_krsubservice
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_krsubservice-c.struct
type cssm_krsubservice struct {
}// cssm_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list-swift.struct
type cssm_list struct {
}// cssm_list_element
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_list_element-c.struct
type cssm_list_element struct {
}// cssm_manager_event_notification
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_event_notification-c.struct
type cssm_manager_event_notification struct {
}// cssm_manager_registration_info
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_manager_registration_info-c.struct
type cssm_manager_registration_info struct {
}// cssm_memory_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_memory_funcs-swift.struct
type cssm_memory_funcs struct {
}// cssm_module_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_module_funcs-c.struct
type cssm_module_funcs struct {
}// cssm_name_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_name_list-swift.struct
type cssm_name_list struct {
}// cssm_net_address
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_net_address-c.struct
type cssm_net_address struct {
}// cssm_parsed_cert
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_cert-swift.struct
type cssm_parsed_cert struct {
}// cssm_parsed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_parsed_crl-swift.struct
type cssm_parsed_crl struct {
}// cssm_pkcs1_oaep_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs1_oaep_params-c.struct
type cssm_pkcs1_oaep_params struct {
}// cssm_pkcs5_pbkdf1_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf1_params-c.struct
type cssm_pkcs5_pbkdf1_params struct {
}// cssm_pkcs5_pbkdf2_params
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_pkcs5_pbkdf2_params-c.struct
type cssm_pkcs5_pbkdf2_params struct {
}// cssm_query
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query-c.struct
type cssm_query struct {
}// cssm_query_limits
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_limits-c.struct
type cssm_query_limits struct {
}// cssm_query_size_data
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_query_size_data-swift.struct
type cssm_query_size_data struct {
}// cssm_range
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_range-swift.struct
type cssm_range struct {
}// cssm_resource_control_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_resource_control_context-c.struct
type cssm_resource_control_context struct {
}// cssm_sample
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_sample-c.struct
type cssm_sample struct {
}// cssm_samplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_samplegroup-c.struct
type cssm_samplegroup struct {
}// cssm_selection_predicate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_selection_predicate-c.struct
type cssm_selection_predicate struct {
}// cssm_spi_ac_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_ac_funcs-c.struct
type cssm_spi_ac_funcs struct {
}// cssm_spi_cl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_cl_funcs-c.struct
type cssm_spi_cl_funcs struct {
}// cssm_spi_csp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_csp_funcs-c.struct
type cssm_spi_csp_funcs struct {
}// cssm_spi_dl_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_dl_funcs-c.struct
type cssm_spi_dl_funcs struct {
}// cssm_spi_kr_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_kr_funcs-c.struct
type cssm_spi_kr_funcs struct {
}// cssm_spi_tp_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_spi_tp_funcs-c.struct
type cssm_spi_tp_funcs struct {
}// cssm_state_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_state_funcs-c.struct
type cssm_state_funcs struct {
}// cssm_subservice_uid
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_subservice_uid-c.struct
type cssm_subservice_uid struct {
}// cssm_tp_authority_id
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_authority_id-c.struct
type cssm_tp_authority_id struct {
}// cssm_tp_callerauth_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_callerauth_context-c.struct
type cssm_tp_callerauth_context struct {
}// cssm_tp_certchange_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_input-c.struct
type cssm_tp_certchange_input struct {
}// cssm_tp_certchange_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certchange_output-c.struct
type cssm_tp_certchange_output struct {
}// cssm_tp_certissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_input-c.struct
type cssm_tp_certissue_input struct {
}// cssm_tp_certissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certissue_output-c.struct
type cssm_tp_certissue_output struct {
}// cssm_tp_certnotarize_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_input-c.struct
type cssm_tp_certnotarize_input struct {
}// cssm_tp_certnotarize_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certnotarize_output-c.struct
type cssm_tp_certnotarize_output struct {
}// cssm_tp_certreclaim_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_input-c.struct
type cssm_tp_certreclaim_input struct {
}// cssm_tp_certreclaim_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certreclaim_output-c.struct
type cssm_tp_certreclaim_output struct {
}// cssm_tp_certverify_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_input-c.struct
type cssm_tp_certverify_input struct {
}// cssm_tp_certverify_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_certverify_output-c.struct
type cssm_tp_certverify_output struct {
}// cssm_tp_confirm_response
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_confirm_response-c.struct
type cssm_tp_confirm_response struct {
}// cssm_tp_crlissue_input
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_input-c.struct
type cssm_tp_crlissue_input struct {
}// cssm_tp_crlissue_output
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_crlissue_output-c.struct
type cssm_tp_crlissue_output struct {
}// cssm_tp_policyinfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_policyinfo-c.struct
type cssm_tp_policyinfo struct {
}// cssm_tp_request_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_request_set-c.struct
type cssm_tp_request_set struct {
}// cssm_tp_result_set
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_result_set-swift.struct
type cssm_tp_result_set struct {
}// cssm_tp_verify_context
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context-c.struct
type cssm_tp_verify_context struct {
}// cssm_tp_verify_context_result
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tp_verify_context_result-c.struct
type cssm_tp_verify_context_result struct {
}// cssm_tuplegroup
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_tuplegroup-c.struct
type cssm_tuplegroup struct {
}// cssm_upcalls
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_upcalls-c.struct
type cssm_upcalls struct {
}// cssm_version
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_version-swift.struct
type cssm_version struct {
}// cssm_x509_extension
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extension-c.struct
type cssm_x509_extension struct {
}// cssm_x509_extensionTagAndValue
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensionTagAndValue
type cssm_x509_extensionTagAndValue struct {
}// cssm_x509_extensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_extensions-c.struct
type cssm_x509_extensions struct {
}// cssm_x509_name
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_name-c.struct
type cssm_x509_name struct {
}// cssm_x509_rdn
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_rdn-c.struct
type cssm_x509_rdn struct {
}// cssm_x509_revoked_cert_entry
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_entry-c.struct
type cssm_x509_revoked_cert_entry struct {
}// cssm_x509_revoked_cert_list
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_revoked_cert_list-c.struct
type cssm_x509_revoked_cert_list struct {
}// cssm_x509_signature
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signature-c.struct
type cssm_x509_signature struct {
}// cssm_x509_signed_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_certificate-c.struct
type cssm_x509_signed_certificate struct {
}// cssm_x509_signed_crl
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_signed_crl-c.struct
type cssm_x509_signed_crl struct {
}// cssm_x509_tbs_certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certificate-c.struct
type cssm_x509_tbs_certificate struct {
}// cssm_x509_tbs_certlist
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_tbs_certlist-c.struct
type cssm_x509_tbs_certlist struct {
}// cssm_x509_time
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_time-c.struct
type cssm_x509_time struct {
}// cssm_x509_type_value_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509_type_value_pair-c.struct
type cssm_x509_type_value_pair struct {
}// cssm_x509ext_basicConstraints
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_basicConstraints-c.struct
type cssm_x509ext_basicConstraints struct {
}// cssm_x509ext_pair
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_pair-c.struct
type cssm_x509ext_pair struct {
}// cssm_x509ext_policyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyInfo-c.struct
type cssm_x509ext_policyInfo struct {
}// cssm_x509ext_policyQualifierInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifierInfo-c.struct
type cssm_x509ext_policyQualifierInfo struct {
}// cssm_x509ext_policyQualifiers
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/cssm_x509ext_policyQualifiers-c.struct
type cssm_x509ext_policyQualifiers struct {
}// mds_funcs
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/mds_funcs-c.struct
type mds_funcs struct {
}// x509_validity
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Security/x509_validity
type x509_validity struct {
}



