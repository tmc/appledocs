// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GSS Functions (27 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GSSCreateName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GSSCredentialCopyUUID func(unsafe.Pointer) unsafe.Pointer
	_GSSNameCreateDisplayString func(unsafe.Pointer) unsafe.Pointer
	_gss_aapl_initial_cred func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_accept_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_compare_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_create_empty_buffer_set func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_create_empty_oid_set func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_duplicate_oid func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_export_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_get_mic func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_init_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_inquire_mechs_for_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_krb5_export_lucid_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_krb5_free_lucid_sec_context func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_oid_equal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_oid_to_str func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_release_buffer_set func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_release_name func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_seal func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_test_oid_set_member func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_unseal func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_verify func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_verify_mic func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gss_wrap func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gsskrb5_extract_authz_data_from_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_gsskrb5_register_acceptor_identity func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_GSSCreateName, lib, "GSSCreateName")
	tryRegister(&_GSSCredentialCopyUUID, lib, "GSSCredentialCopyUUID")
	tryRegister(&_GSSNameCreateDisplayString, lib, "GSSNameCreateDisplayString")
	tryRegister(&_gss_aapl_initial_cred, lib, "gss_aapl_initial_cred")
	tryRegister(&_gss_accept_sec_context, lib, "gss_accept_sec_context")
	tryRegister(&_gss_compare_name, lib, "gss_compare_name")
	tryRegister(&_gss_create_empty_buffer_set, lib, "gss_create_empty_buffer_set")
	tryRegister(&_gss_create_empty_oid_set, lib, "gss_create_empty_oid_set")
	tryRegister(&_gss_duplicate_oid, lib, "gss_duplicate_oid")
	tryRegister(&_gss_export_name, lib, "gss_export_name")
	tryRegister(&_gss_get_mic, lib, "gss_get_mic")
	tryRegister(&_gss_init_sec_context, lib, "gss_init_sec_context")
	tryRegister(&_gss_inquire_mechs_for_name, lib, "gss_inquire_mechs_for_name")
	tryRegister(&_gss_krb5_export_lucid_sec_context, lib, "gss_krb5_export_lucid_sec_context")
	tryRegister(&_gss_krb5_free_lucid_sec_context, lib, "gss_krb5_free_lucid_sec_context")
	tryRegister(&_gss_oid_equal, lib, "gss_oid_equal")
	tryRegister(&_gss_oid_to_str, lib, "gss_oid_to_str")
	tryRegister(&_gss_release_buffer_set, lib, "gss_release_buffer_set")
	tryRegister(&_gss_release_name, lib, "gss_release_name")
	tryRegister(&_gss_seal, lib, "gss_seal")
	tryRegister(&_gss_test_oid_set_member, lib, "gss_test_oid_set_member")
	tryRegister(&_gss_unseal, lib, "gss_unseal")
	tryRegister(&_gss_verify, lib, "gss_verify")
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



// Returns a GSS name given a buffer and a type. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCreateName(_:_:_:)
func GSSCreateName(name unsafe.Pointer, name_type unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _GSSCreateName(name, name_type, error_)
	}


// Returns a copy of the universally unique identifier corresponding to a GSS credential. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCredentialCopyUUID(_:)
func GSSCredentialCopyUUID(credential unsafe.Pointer) unsafe.Pointer {
	return _GSSCredentialCopyUUID(credential)
	}


// Returns a string suitable for displaying to the user from a GSS name. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSNameCreateDisplayString(_:)
func GSSNameCreateDisplayString(name unsafe.Pointer) unsafe.Pointer {
	return _GSSNameCreateDisplayString(name)
	}


// Acquires a new credential using a password or certificate. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_aapl_initial_cred(_:_:_:_:_:)
func gss_aapl_initial_cred(desired_name unsafe.Pointer, desired_mech unsafe.Pointer, attributes unsafe.Pointer, output_cred_handle unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _gss_aapl_initial_cred(desired_name, desired_mech, attributes, output_cred_handle, error_)
	}


// Accepts a security context initiated by a peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_accept_sec_context(_:_:_:_:_:_:_:_:_:_:_:)
func gss_accept_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, acceptor_cred_handle unsafe.Pointer, input_token unsafe.Pointer, input_chan_bindings unsafe.Pointer, src_name unsafe.Pointer, mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags unsafe.Pointer, time_rec unsafe.Pointer, delegated_cred_handle unsafe.Pointer) unsafe.Pointer {
	return _gss_accept_sec_context(minor_status, context_handle, acceptor_cred_handle, input_token, input_chan_bindings, src_name, mech_type, output_token, ret_flags, time_rec, delegated_cred_handle)
	}


// Returns a flag that indicates if two names in internal name format refer to the same entity. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_compare_name(_:_:_:_:)
func gss_compare_name(minor_status unsafe.Pointer, name1_arg unsafe.Pointer, name2_arg unsafe.Pointer, name_equal unsafe.Pointer) unsafe.Pointer {
	return _gss_compare_name(minor_status, name1_arg, name2_arg, name_equal)
	}


// Allocates an empty buffer set descriptor that you use to manage an array of buffers. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_buffer_set(_:_:)
func gss_create_empty_buffer_set(minor_status unsafe.Pointer, buffer_set unsafe.Pointer) unsafe.Pointer {
	return _gss_create_empty_buffer_set(minor_status, buffer_set)
	}


// Allocates a new, empty set to hold object identifiers. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_oid_set(_:_:)
func gss_create_empty_oid_set(minor_status unsafe.Pointer, oid_set unsafe.Pointer) unsafe.Pointer {
	return _gss_create_empty_oid_set(minor_status, oid_set)
	}


// Copies an OID into a new object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_duplicate_oid(_:_:_:)
func gss_duplicate_oid(minor_status unsafe.Pointer, src_oid unsafe.Pointer, dest_oid unsafe.Pointer) unsafe.Pointer {
	return _gss_duplicate_oid(minor_status, src_oid, dest_oid)
	}


// Returns a mechanism name in contiguous octet format. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_export_name(_:_:_:)
func gss_export_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, exported_name unsafe.Pointer) unsafe.Pointer {
	return _gss_export_name(minor_status, input_name, exported_name)
	}


// Returns a token that contains the MIC for a message. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_get_mic(_:_:_:_:_:)
func gss_get_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, qop_req unsafe.Pointer, message_buffer unsafe.Pointer, message_token unsafe.Pointer) unsafe.Pointer {
	return _gss_get_mic(minor_status, context_handle, qop_req, message_buffer, message_token)
	}


// Initiates a security context with a peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_init_sec_context(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func gss_init_sec_context(minor_status unsafe.Pointer, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags unsafe.Pointer, time_req unsafe.Pointer, input_chan_bindings unsafe.Pointer, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token unsafe.Pointer, ret_flags unsafe.Pointer, time_rec unsafe.Pointer) unsafe.Pointer {
	return _gss_init_sec_context(minor_status, initiator_cred_handle, context_handle, target_name, input_mech_type, req_flags, time_req, input_chan_bindings, input_token, actual_mech_type, output_token, ret_flags, time_rec)
	}


// Returns a list of mechanisms that support a particular name type. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_mechs_for_name(_:_:_:)
func gss_inquire_mechs_for_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, mech_types unsafe.Pointer) unsafe.Pointer {
	return _gss_inquire_mechs_for_name(minor_status, input_name, mech_types)
	}


// Returns a non-opaque version of the internal context information. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_export_lucid_sec_context(_:_:_:_:)
func gss_krb5_export_lucid_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, version unsafe.Pointer, rctx unsafe.Pointer) unsafe.Pointer {
	return _gss_krb5_export_lucid_sec_context(minor_status, context_handle, version, rctx)
	}


// Frees allocated storage associated with an exported context. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_free_lucid_sec_context(_:_:)
func gss_krb5_free_lucid_sec_context(minor_status unsafe.Pointer, c unsafe.Pointer) unsafe.Pointer {
	return _gss_krb5_free_lucid_sec_context(minor_status, c)
	}


// Returns a flag that indicates whether two object identifiers are the same. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_equal(_:_:)
func gss_oid_equal(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _gss_oid_equal(a, b)
	}


// Converts an OID object to a human-readable string. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_to_str(_:_:_:)
func gss_oid_to_str(minor_status unsafe.Pointer, oid unsafe.Pointer, oid_str unsafe.Pointer) unsafe.Pointer {
	return _gss_oid_to_str(minor_status, oid, oid_str)
	}


// Frees the memory associated with a buffer set descriptor and all the buffers it contains. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_buffer_set(_:_:)
func gss_release_buffer_set(minor_status unsafe.Pointer, buffer_set unsafe.Pointer) unsafe.Pointer {
	return _gss_release_buffer_set(minor_status, buffer_set)
	}


// Frees the resources associated with a name object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_name(_:_:)
func gss_release_name(minor_status unsafe.Pointer, input_name unsafe.Pointer) unsafe.Pointer {
	return _gss_release_name(minor_status, input_name)
	}


// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_seal(_:_:_:_:_:_:_:)
func gss_seal(minor_status unsafe.Pointer, context_handle unsafe.Pointer, conf_req_flag unsafe.Pointer, qop_req unsafe.Pointer, input_message_buffer unsafe.Pointer, conf_state unsafe.Pointer, output_message_buffer unsafe.Pointer) unsafe.Pointer {
	return _gss_seal(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
	}


// Returns a flag that indicates if an OID is present in an OID set. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_test_oid_set_member(_:_:_:_:)
func gss_test_oid_set_member(minor_status unsafe.Pointer, member unsafe.Pointer, set unsafe.Pointer, present unsafe.Pointer) unsafe.Pointer {
	return _gss_test_oid_set_member(minor_status, member, set, present)
	}


// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_unseal(_:_:_:_:_:_:)
func gss_unseal(minor_status unsafe.Pointer, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer unsafe.Pointer, conf_state unsafe.Pointer, qop_state unsafe.Pointer) unsafe.Pointer {
	return _gss_unseal(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
	}


// Returns a flag that indicates the integrity of a message’s digital signature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_verify(_:_:_:_:_:)
func gss_verify(minor_status unsafe.Pointer, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state unsafe.Pointer) unsafe.Pointer {
	return _gss_verify(minor_status, context_handle, message_buffer, token_buffer, qop_state)
	}


// Returns an indication of whether the integrity of a message is intact, given its MIC token. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_verify_mic(_:_:_:_:_:)
func gss_verify_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state unsafe.Pointer) unsafe.Pointer {
	return _gss_verify_mic(minor_status, context_handle, message_buffer, token_buffer, qop_state)
	}


// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_wrap(_:_:_:_:_:_:_:)
func gss_wrap(minor_status unsafe.Pointer, context_handle unsafe.Pointer, conf_req_flag unsafe.Pointer, qop_req unsafe.Pointer, input_message_buffer unsafe.Pointer, conf_state unsafe.Pointer, output_message_buffer unsafe.Pointer) unsafe.Pointer {
	return _gss_wrap(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
	}


// Extracts Kerberos authorization data stored within the context. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_extract_authz_data_from_sec_context(_:_:_:_:)
func gsskrb5_extract_authz_data_from_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, ad_type unsafe.Pointer, ad_data unsafe.Pointer) unsafe.Pointer {
	return _gsskrb5_extract_authz_data_from_sec_context(minor_status, context_handle, ad_type, ad_data)
	}


// Sets the Kerberos 5 file-based key that the acceptor will use. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_register_acceptor_identity(_:)
func gsskrb5_register_acceptor_identity(identity unsafe.Pointer) unsafe.Pointer {
	return _gsskrb5_register_acceptor_identity(identity)
	}




