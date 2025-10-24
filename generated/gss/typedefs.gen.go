// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

// Type aliases and typedefs
// Gss_auth_identity_t - A pointer to an opaque object used to manage authentication identities.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_auth_identity_t
// gss_auth_identity_t has base type: struct gss_auth_identity *
type Gss_auth_identity_t uintptr
// Gss_buffer_desc - The buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_desc
// gss_buffer_desc has base type: struct gss_buffer_desc_struct
type Gss_buffer_desc uintptr
// Gss_buffer_set_desc - The descriptor that you use to manage an array of buffer descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_set_desc
// gss_buffer_set_desc has base type: struct gss_buffer_set_desc_struct
type Gss_buffer_set_desc uintptr
// Gss_buffer_set_t - A pointer to the descriptor that you use to manage an array of buffer descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_set_t
// gss_buffer_set_t has base type: struct gss_buffer_set_desc_struct *
type Gss_buffer_set_t uintptr
// Gss_buffer_t - A pointer to a buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_t
// gss_buffer_t has base type: struct gss_buffer_desc_struct *
type Gss_buffer_t uintptr
// Gss_channel_bindings_t - A pointer to a channel bindings descriptor that specifies the communications channel used to carry a context.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_channel_bindings_t
// gss_channel_bindings_t has base type: struct gss_channel_bindings_struct *
type Gss_channel_bindings_t uintptr
// Gss_const_buffer_t - A pointer to an immutable buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_buffer_t
// gss_const_buffer_t has base type: const gss_buffer_desc *
type Gss_const_buffer_t uintptr
// Gss_const_channel_bindings_t - A pointer to an immutable channel bindings descriptor that you use to specify the communications channel used to carry a context.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_channel_bindings_t
// gss_const_channel_bindings_t has base type: const struct gss_channel_bindings_struct *
type Gss_const_channel_bindings_t uintptr
// Gss_const_cred_id_t - A pointer to an immutable opaque type that you use to exchange a credential object with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_cred_id_t
// gss_const_cred_id_t has base type: const struct gss_cred_id_t_desc_struct *
type Gss_const_cred_id_t uintptr
// Gss_const_name_t - A pointer to an immutable version of the opaque descriptor used to exchange name objects with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_name_t
// gss_const_name_t has base type: const struct gss_name_t_desc_struct *
type Gss_const_name_t uintptr
// Gss_const_OID - A pointer to an immutable OID descriptor exchanges object identifiers with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_OID
// gss_const_OID has base type: const gss_OID_desc *
type Gss_const_OID uintptr
// Gss_const_OID_set - A pointer to an immutable descriptor manages an array of OID descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_OID_set
// gss_const_OID_set has base type: const gss_OID_set_desc *
type Gss_const_OID_set uintptr
// Gss_cred_id_t - A pointer to an opaque type that you use to exchange a credential object with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_cred_id_t
// gss_cred_id_t has base type: struct gss_cred_id_t_desc_struct *
type Gss_cred_id_t uintptr
// Gss_cred_usage_t - A credential usage value.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_cred_usage_t
type Gss_cred_usage_t int32
// Gss_ctx_id_t - A pointer to an opaque type that you use to communicate context pointers with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_ctx_id_t
// gss_ctx_id_t has base type: struct gss_ctx_id_t_desc_struct *
type Gss_ctx_id_t uintptr
// Gss_iov_buffer_desc - The structure for a vectored I/O buffer and its defined type.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_iov_buffer_desc
// gss_iov_buffer_desc has base type: struct gss_iov_buffer_desc_struct
type Gss_iov_buffer_desc uintptr
// Gss_iov_buffer_t - The structure for a vectored I/O buffer and its defined type.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_iov_buffer_t
// gss_iov_buffer_t has base type: struct gss_iov_buffer_desc_struct *
type Gss_iov_buffer_t uintptr
// Gss_krb5_cfx_keydata_t - The structure of a Kerberos context and acceptor-asserted key.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_cfx_keydata_t
// gss_krb5_cfx_keydata_t has base type: struct gss_krb5_cfx_keydata
type Gss_krb5_cfx_keydata_t uintptr
// Gss_krb5_lucid_context_v1_t - The structure of a Kerberos context.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_context_v1_t
// gss_krb5_lucid_context_v1_t has base type: struct gss_krb5_lucid_context_v1
type Gss_krb5_lucid_context_v1_t uintptr
// Gss_krb5_lucid_key_t - The structure for a Kerberos encryption key.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_lucid_key_t
// gss_krb5_lucid_key_t has base type: struct gss_krb5_lucid_key
type Gss_krb5_lucid_key_t uintptr
// Gss_krb5_rfc1964_keydata_t - The structure for an RFC 1964-compliant Kerberos encryption key.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_rfc1964_keydata_t
// gss_krb5_rfc1964_keydata_t has base type: struct gss_krb5_rfc1964_keydata
type Gss_krb5_rfc1964_keydata_t uintptr
// Gss_name_t - A pointer to an opaque type that you use to communicate name objects with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_name_t
// gss_name_t has base type: struct gss_name_t_desc_struct *
type Gss_name_t uintptr
// Gss_OID - A pointer to the OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID
// gss_OID has base type: struct gss_OID_desc_struct *
type Gss_OID uintptr
// Gss_OID_desc - The OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_desc
// gss_OID_desc has base type: struct gss_OID_desc_struct
type Gss_OID_desc uintptr
// Gss_OID_set - A pointer to a descriptor that manages an array of OID descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_set
// gss_OID_set has base type: struct gss_OID_set_desc_struct *
type Gss_OID_set uintptr
// Gss_OID_set_desc - The descriptor that manages an array of OID descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_set_desc
// gss_OID_set_desc has base type: struct gss_OID_set_desc_struct
type Gss_OID_set_desc uintptr
// Gss_qop_t - A quality of protection setting.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_qop_t
// gss_qop_t has base type: OM_uint32
type Gss_qop_t uintptr
// Gss_status_id_t - A pointer to a status result.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_status_id_t
// gss_status_id_t has base type: OM_uint32 *
type Gss_status_id_t uintptr
// Gss_uint32 - A 32-bit unsigned integer.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_uint32
// gss_uint32 has base type: uint32_t
type Gss_uint32 uintptr
// OM_uint32 - A 32-bit unsigned integer.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/OM_uint32
// OM_uint32 has base type: uint32_t
type OM_uint32 uintptr
// OM_uint64 - A 64-bit unsigned integer.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/OM_uint64
// OM_uint64 has base type: uint64_t
type OM_uint64 uintptr

