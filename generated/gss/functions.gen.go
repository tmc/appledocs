// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GSS Functions (21 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GSSCredentialGetLifetime func(unsafe.Pointer) OM_uint32
	_GSSNameCreateDisplayString func(unsafe.Pointer) unsafe.Pointer
	_gss_add_oid_set_member func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_create_empty_buffer_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_create_empty_oid_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_display_name func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, unsafe.Pointer) OM_uint32
	_gss_export_name func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_get_mic func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_init_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, OM_uint32, OM_uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_saslname_for_mech func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, Gss_buffer_t, Gss_buffer_t) OM_uint32
	_gss_krb5_copy_ccache func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_krb5_export_lucid_sec_context func(unsafe.Pointer, unsafe.Pointer, OM_uint32, unsafe.Pointer) OM_uint32
	_gss_oid_equal func(unsafe.Pointer, unsafe.Pointer) int
	_gss_oid_to_str func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_test_oid_set_member func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []int) OM_uint32
	_gss_unseal func(unsafe.Pointer, Gss_ctx_id_t, Gss_buffer_t, Gss_buffer_t, []int, []int) OM_uint32
	_gss_unwrap func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, []int, unsafe.Pointer) OM_uint32
	_gss_verify_mic func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_wrap func(unsafe.Pointer, unsafe.Pointer, int, unsafe.Pointer, unsafe.Pointer, []int, Gss_buffer_t) OM_uint32
	_gsskrb5_extract_authz_data_from_sec_context func(unsafe.Pointer, Gss_ctx_id_t, int, Gss_buffer_t) OM_uint32
	_gsskrb5_register_acceptor_identity func(unsafe.Pointer) OM_uint32
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_GSSCredentialGetLifetime, lib, "GSSCredentialGetLifetime")
	tryRegister(&_GSSNameCreateDisplayString, lib, "GSSNameCreateDisplayString")
	tryRegister(&_gss_add_oid_set_member, lib, "gss_add_oid_set_member")
	tryRegister(&_gss_create_empty_buffer_set, lib, "gss_create_empty_buffer_set")
	tryRegister(&_gss_create_empty_oid_set, lib, "gss_create_empty_oid_set")
	tryRegister(&_gss_display_name, lib, "gss_display_name")
	tryRegister(&_gss_export_name, lib, "gss_export_name")
	tryRegister(&_gss_get_mic, lib, "gss_get_mic")
	tryRegister(&_gss_init_sec_context, lib, "gss_init_sec_context")
	tryRegister(&_gss_inquire_saslname_for_mech, lib, "gss_inquire_saslname_for_mech")
	tryRegister(&_gss_krb5_copy_ccache, lib, "gss_krb5_copy_ccache")
	tryRegister(&_gss_krb5_export_lucid_sec_context, lib, "gss_krb5_export_lucid_sec_context")
	tryRegister(&_gss_oid_equal, lib, "gss_oid_equal")
	tryRegister(&_gss_oid_to_str, lib, "gss_oid_to_str")
	tryRegister(&_gss_test_oid_set_member, lib, "gss_test_oid_set_member")
	tryRegister(&_gss_unseal, lib, "gss_unseal")
	tryRegister(&_gss_unwrap, lib, "gss_unwrap")
	tryRegister(&_gss_verify_mic, lib, "gss_verify_mic")
	tryRegister(&_gss_wrap, lib, "gss_wrap")
	tryRegister(&_gsskrb5_extract_authz_data_from_sec_context, lib, "gsskrb5_extract_authz_data_from_sec_context")
	tryRegister(&_gsskrb5_register_acceptor_identity, lib, "gsskrb5_register_acceptor_identity")
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



// Returns the remaining time in seconds before the credential expires.
//
// Added in macOS 10.9.
// Returns the remaining time in seconds before the credential expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCredentialGetLifetime(_:)
func GSSCredentialGetLifetime(cred unsafe.Pointer) OM_uint32 {
	return _GSSCredentialGetLifetime(cred)
}

// Returns a string suitable for displaying to the user from a GSS name.
//
// Added in macOS 10.9.
// Returns a string suitable for displaying to the user from a GSS name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSNameCreateDisplayString(_:)
func GSSNameCreateDisplayString(name unsafe.Pointer) unsafe.Pointer {
	return _GSSNameCreateDisplayString(name)
}

// Adds an object identifier into an OID set.
//
// Added in macOS 10.7.
// Adds an object identifier into an OID set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_add_oid_set_member(_:_:_:)
func gss_add_oid_set_member(minor_status unsafe.Pointer, member_oid unsafe.Pointer, oid_set unsafe.Pointer) OM_uint32 {
	return _gss_add_oid_set_member(minor_status, member_oid, oid_set)
}

// Allocates an empty buffer set descriptor that you use to manage an array of buffers.
//
// Added in macOS 10.7.
// Allocates an empty buffer set descriptor that you use to manage an array of buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_buffer_set(_:_:)
func gss_create_empty_buffer_set(minor_status unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32 {
	return _gss_create_empty_buffer_set(minor_status, buffer_set)
}

// Allocates a new, empty set to hold object identifiers.
//
// Added in macOS 10.7.
// Allocates a new, empty set to hold object identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_oid_set(_:_:)
func gss_create_empty_oid_set(minor_status unsafe.Pointer, oid_set unsafe.Pointer) OM_uint32 {
	return _gss_create_empty_oid_set(minor_status, oid_set)
}

// Converts a name in the internal format to an octet string and the associated name type.
//
// Added in macOS 10.7.
// Converts a name in the internal format to an octet string and the associated name type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_display_name(_:_:_:_:)
func gss_display_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, output_name_buffer Gss_buffer_t, output_name_type unsafe.Pointer) OM_uint32 {
	return _gss_display_name(minor_status, input_name, output_name_buffer, output_name_type)
}

// Returns a mechanism name in contiguous octet format.
//
// Added in macOS 10.7.
// Returns a mechanism name in contiguous octet format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_export_name(_:_:_:)
func gss_export_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, exported_name Gss_buffer_t) OM_uint32 {
	return _gss_export_name(minor_status, input_name, exported_name)
}

// Returns a token that contains the MIC for a message.
//
// Added in macOS 10.7.
// Returns a token that contains the MIC for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_get_mic(_:_:_:_:_:)
func gss_get_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, qop_req unsafe.Pointer, message_buffer unsafe.Pointer, message_token Gss_buffer_t) OM_uint32 {
	return _gss_get_mic(minor_status, context_handle, qop_req, message_buffer, message_token)
}

// Initiates a security context with a peer.
//
// Added in macOS 10.7.
// Initiates a security context with a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_init_sec_context(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func gss_init_sec_context(minor_status unsafe.Pointer, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags OM_uint32, time_req OM_uint32, input_chan_bindings unsafe.Pointer, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token Gss_buffer_t, ret_flags unsafe.Pointer, time_rec unsafe.Pointer) OM_uint32 {
	return _gss_init_sec_context(minor_status, initiator_cred_handle, context_handle, target_name, input_mech_type, req_flags, time_req, input_chan_bindings, input_token, actual_mech_type, output_token, ret_flags, time_rec)
}

// Returns the Simple Authentication and Security Layer (SASL) protocol name for a given GSS-API mechanism.
//
// Added in macOS 10.10.
// Returns the Simple Authentication and Security Layer (SASL) protocol name for a given GSS-API mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_saslname_for_mech(_:_:_:_:_:)
func gss_inquire_saslname_for_mech(minor_status unsafe.Pointer, desired_mech unsafe.Pointer, sasl_mech_name Gss_buffer_t, mech_name Gss_buffer_t, mech_description Gss_buffer_t) OM_uint32 {
	return _gss_inquire_saslname_for_mech(minor_status, desired_mech, sasl_mech_name, mech_name, mech_description)
}

// Copies Kerberos 5 credentials into the passed cache.

// Copies Kerberos 5 credentials into the passed cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_copy_ccache(_:_:_:)
func gss_krb5_copy_ccache(minor_status unsafe.Pointer, cred unsafe.Pointer, out unsafe.Pointer) OM_uint32 {
	return _gss_krb5_copy_ccache(minor_status, cred, out)
}

// Returns a non-opaque version of the internal context information.
//
// Added in macOS 10.7.
// Returns a non-opaque version of the internal context information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_export_lucid_sec_context(_:_:_:_:)
func gss_krb5_export_lucid_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, version OM_uint32, rctx unsafe.Pointer) OM_uint32 {
	return _gss_krb5_export_lucid_sec_context(minor_status, context_handle, version, rctx)
}

// Returns a flag that indicates whether two object identifiers are the same.
//
// Added in macOS 10.7.
// Returns a flag that indicates whether two object identifiers are the same.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_equal(_:_:)
func gss_oid_equal(a unsafe.Pointer, b unsafe.Pointer) int {
	return _gss_oid_equal(a, b)
}

// Converts an OID object to a human-readable string.
//
// Added in macOS 10.7.
// Converts an OID object to a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_to_str(_:_:_:)
func gss_oid_to_str(minor_status unsafe.Pointer, oid unsafe.Pointer, oid_str Gss_buffer_t) OM_uint32 {
	return _gss_oid_to_str(minor_status, oid, oid_str)
}

// Returns a flag that indicates if an OID is present in an OID set.
//
// Added in macOS 10.7.
// Returns a flag that indicates if an OID is present in an OID set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_test_oid_set_member(_:_:_:_:)
func gss_test_oid_set_member(minor_status unsafe.Pointer, member unsafe.Pointer, set unsafe.Pointer, present []int) OM_uint32 {
	return _gss_test_oid_set_member(minor_status, member, set, present)
}

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_unseal(_:_:_:_:_:_:)
func gss_unseal(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, input_message_buffer Gss_buffer_t, output_message_buffer Gss_buffer_t, conf_state []int, qop_state []int) OM_uint32 {
	return _gss_unseal(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
}

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// Added in macOS 10.7.
// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_unwrap(_:_:_:_:_:_:)
func gss_unwrap(minor_status unsafe.Pointer, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer Gss_buffer_t, conf_state []int, qop_state unsafe.Pointer) OM_uint32 {
	return _gss_unwrap(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
}

// Returns an indication of whether the integrity of a message is intact, given its MIC token.
//
// Added in macOS 10.7.
// Returns an indication of whether the integrity of a message is intact, given its MIC token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_verify_mic(_:_:_:_:_:)
func gss_verify_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state unsafe.Pointer) OM_uint32 {
	return _gss_verify_mic(minor_status, context_handle, message_buffer, token_buffer, qop_state)
}

// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// Added in macOS 10.7.
// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_wrap(_:_:_:_:_:_:_:)
func gss_wrap(minor_status unsafe.Pointer, context_handle unsafe.Pointer, conf_req_flag int, qop_req unsafe.Pointer, input_message_buffer unsafe.Pointer, conf_state []int, output_message_buffer Gss_buffer_t) OM_uint32 {
	return _gss_wrap(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
}

// Extracts Kerberos authorization data stored within the context.
//
// Added in macOS 10.7.
// Extracts Kerberos authorization data stored within the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_extract_authz_data_from_sec_context(_:_:_:_:)
func gsskrb5_extract_authz_data_from_sec_context(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, ad_type int, ad_data Gss_buffer_t) OM_uint32 {
	return _gsskrb5_extract_authz_data_from_sec_context(minor_status, context_handle, ad_type, ad_data)
}

// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// Added in macOS 10.7.
// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_register_acceptor_identity(_:)
func gsskrb5_register_acceptor_identity(identity unsafe.Pointer) OM_uint32 {
	return _gsskrb5_register_acceptor_identity(identity)
}



