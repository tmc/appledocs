// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

// Type aliases and typedefs
// OM_uint32 - A 32-bit unsigned integer.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/OM_uint32
// OM_uint32 has base type: uint32_t
type OM_uint32 uintptr
// Gss_OID_desc - The OID descriptor that exchanges object identifiers with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_OID_desc
// gss_OID_desc has base type: struct gss_OID_desc_struct
type Gss_OID_desc uintptr
// Gss_buffer_desc - The buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_desc
// gss_buffer_desc has base type: struct gss_buffer_desc_struct
type Gss_buffer_desc uintptr
// Gss_buffer_t - A pointer to a buffer descriptor that you use to exchange octet streams with many GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_buffer_t
// gss_buffer_t has base type: struct gss_buffer_desc_struct *
type Gss_buffer_t uintptr
// Gss_const_OID_set - A pointer to an immutable descriptor manages an array of OID descriptors.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_OID_set
// gss_const_OID_set has base type: const gss_OID_set_desc *
type Gss_const_OID_set uintptr
// Gss_const_cred_id_t - A pointer to an immutable opaque type that you use to exchange a credential object with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_const_cred_id_t
// gss_const_cred_id_t has base type: const struct gss_cred_id_t_desc_struct *
type Gss_const_cred_id_t uintptr
// Gss_ctx_id_t - A pointer to an opaque type that you use to communicate context pointers with GSS-API functions.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_ctx_id_t
// gss_ctx_id_t has base type: struct gss_ctx_id_t_desc_struct *
type Gss_ctx_id_t uintptr
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

