// Code generated from Apple documentation for GSS. DO NOT EDIT.

package gss

/* debug [functions.gen.go]: Generating 82 functions for GSS */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GSS Functions (82 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_gss_aapl_change_password func(unsafe.Pointer, Gss_const_OID, DictionaryRef, unsafe.Pointer) OM_uint32
	_gss_aapl_initial_cred func(unsafe.Pointer, Gss_const_OID, DictionaryRef, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_accept_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_acquire_cred func(unsafe.Pointer, unsafe.Pointer, OM_uint32, unsafe.Pointer, Gss_cred_usage_t, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_acquire_cred_with_password func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, OM_uint32, unsafe.Pointer, Gss_cred_usage_t, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_add_buffer_set_member func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_add_cred func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_cred_usage_t, OM_uint32, OM_uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_add_oid_set_member func(unsafe.Pointer, Gss_const_OID, unsafe.Pointer) OM_uint32
	_gss_canonicalize_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_compare_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []int) OM_uint32
	_gss_context_time func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_create_empty_buffer_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_create_empty_oid_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_decapsulate_token func(Gss_const_buffer_t, Gss_const_OID, Gss_buffer_t) OM_uint32
	_gss_delete_sec_context func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_destroy_cred func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_display_mech_attr func(unsafe.Pointer, Gss_const_OID, Gss_buffer_t, Gss_buffer_t, Gss_buffer_t) OM_uint32
	_gss_display_name func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, unsafe.Pointer) OM_uint32
	_gss_display_status func(unsafe.Pointer, OM_uint32, int, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_duplicate_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_duplicate_oid func(unsafe.Pointer, Gss_OID, unsafe.Pointer) OM_uint32
	_gss_encapsulate_token func(Gss_const_buffer_t, Gss_const_OID, Gss_buffer_t) OM_uint32
	_gss_export_cred func(unsafe.Pointer, Gss_cred_id_t, Gss_buffer_t) OM_uint32
	_gss_export_name func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_export_sec_context func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_get_mic func(unsafe.Pointer, unsafe.Pointer, Gss_qop_t, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_import_cred func(unsafe.Pointer, Gss_buffer_t, unsafe.Pointer) OM_uint32
	_gss_import_name func(unsafe.Pointer, unsafe.Pointer, Gss_const_OID, unsafe.Pointer) OM_uint32
	_gss_import_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_indicate_mechs func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_indicate_mechs_by_attrs func(unsafe.Pointer, Gss_const_OID_set, Gss_const_OID_set, Gss_const_OID_set, unsafe.Pointer) OM_uint32
	_gss_init_sec_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, OM_uint32, OM_uint32, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_attrs_for_mech func(unsafe.Pointer, Gss_const_OID, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_context func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, []int, []int) OM_uint32
	_gss_inquire_cred func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_cred_by_mech func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_cred_by_oid func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_mech_for_saslname func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_mechs_for_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_name func(unsafe.Pointer, Gss_name_t, []int, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_inquire_names_for_mech func(unsafe.Pointer, Gss_const_OID, unsafe.Pointer) OM_uint32
	_gss_inquire_saslname_for_mech func(unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, Gss_buffer_t, Gss_buffer_t) OM_uint32
	_gss_inquire_sec_context_by_oid func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_iter_creds func(unsafe.Pointer, OM_uint32, Gss_const_OID) OM_uint32
	_gss_iter_creds_f func(unsafe.Pointer, OM_uint32, Gss_const_OID, unsafe.Pointer) OM_uint32
	_gss_krb5_ccache_name func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_krb5_copy_ccache func(unsafe.Pointer, Gss_cred_id_t, unsafe.Pointer) OM_uint32
	_gss_krb5_export_lucid_sec_context func(unsafe.Pointer, unsafe.Pointer, OM_uint32, unsafe.Pointer) OM_uint32
	_gss_krb5_free_lucid_sec_context func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_krb5_set_allowable_enctypes func(unsafe.Pointer, Gss_cred_id_t, OM_uint32, unsafe.Pointer) OM_uint32
	_gss_oid_equal func(Gss_const_OID, Gss_const_OID) int
	_gss_oid_to_str func(unsafe.Pointer, Gss_OID, Gss_buffer_t) OM_uint32
	_gss_process_context_token func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_pseudo_random func(unsafe.Pointer, Gss_ctx_id_t, int, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_release_buffer func(unsafe.Pointer, Gss_buffer_t) OM_uint32
	_gss_release_buffer_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_release_cred func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_release_name func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_release_oid func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_release_oid_set func(unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_seal func(unsafe.Pointer, Gss_ctx_id_t, int, int, Gss_buffer_t, []int, Gss_buffer_t) OM_uint32
	_gss_set_cred_option func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_set_sec_context_option func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_sign func(unsafe.Pointer, Gss_ctx_id_t, int, Gss_buffer_t, Gss_buffer_t) OM_uint32
	_gss_test_oid_set_member func(unsafe.Pointer, Gss_const_OID, unsafe.Pointer, []int) OM_uint32
	_gss_unseal func(unsafe.Pointer, Gss_ctx_id_t, Gss_buffer_t, Gss_buffer_t, []int, []int) OM_uint32
	_gss_unwrap func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Gss_buffer_t, []int, unsafe.Pointer) OM_uint32
	_gss_userok func(unsafe.Pointer, unsafe.Pointer) int
	_gss_verify func(unsafe.Pointer, Gss_ctx_id_t, Gss_buffer_t, Gss_buffer_t, []int) OM_uint32
	_gss_verify_mic func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) OM_uint32
	_gss_wrap func(unsafe.Pointer, unsafe.Pointer, int, Gss_qop_t, unsafe.Pointer, []int, Gss_buffer_t) OM_uint32
	_gss_wrap_size_limit func(unsafe.Pointer, unsafe.Pointer, int, Gss_qop_t, OM_uint32, unsafe.Pointer) OM_uint32
	_GSSCreateCredentialFromUUID func(UUIDRef) Gss_cred_id_t
	_GSSCreateError func(Gss_const_OID, OM_uint32, OM_uint32) ErrorRef
	_GSSCreateName func(TypeRef, Gss_const_OID, unsafe.Pointer) Gss_name_t
	_GSSCredentialCopyName func(Gss_cred_id_t) Gss_name_t
	_GSSCredentialCopyUUID func(Gss_cred_id_t) UUIDRef
	_GSSCredentialGetLifetime func(Gss_cred_id_t) OM_uint32
	_gsskrb5_extract_authz_data_from_sec_context func(unsafe.Pointer, Gss_ctx_id_t, int, Gss_buffer_t) OM_uint32
	_gsskrb5_register_acceptor_identity func(unsafe.Pointer) OM_uint32
	_GSSNameCreateDisplayString func(Gss_name_t) StringRef
	_krb5_gss_register_acceptor_identity func(unsafe.Pointer) OM_uint32
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_gss_aapl_change_password, lib, "gss_aapl_change_password")
	tryRegister(&_gss_aapl_initial_cred, lib, "gss_aapl_initial_cred")
	tryRegister(&_gss_accept_sec_context, lib, "gss_accept_sec_context")
	tryRegister(&_gss_acquire_cred, lib, "gss_acquire_cred")
	tryRegister(&_gss_acquire_cred_with_password, lib, "gss_acquire_cred_with_password")
	tryRegister(&_gss_add_buffer_set_member, lib, "gss_add_buffer_set_member")
	tryRegister(&_gss_add_cred, lib, "gss_add_cred")
	tryRegister(&_gss_add_oid_set_member, lib, "gss_add_oid_set_member")
	tryRegister(&_gss_canonicalize_name, lib, "gss_canonicalize_name")
	tryRegister(&_gss_compare_name, lib, "gss_compare_name")
	tryRegister(&_gss_context_time, lib, "gss_context_time")
	tryRegister(&_gss_create_empty_buffer_set, lib, "gss_create_empty_buffer_set")
	tryRegister(&_gss_create_empty_oid_set, lib, "gss_create_empty_oid_set")
	tryRegister(&_gss_decapsulate_token, lib, "gss_decapsulate_token")
	tryRegister(&_gss_delete_sec_context, lib, "gss_delete_sec_context")
	tryRegister(&_gss_destroy_cred, lib, "gss_destroy_cred")
	tryRegister(&_gss_display_mech_attr, lib, "gss_display_mech_attr")
	tryRegister(&_gss_display_name, lib, "gss_display_name")
	tryRegister(&_gss_display_status, lib, "gss_display_status")
	tryRegister(&_gss_duplicate_name, lib, "gss_duplicate_name")
	tryRegister(&_gss_duplicate_oid, lib, "gss_duplicate_oid")
	tryRegister(&_gss_encapsulate_token, lib, "gss_encapsulate_token")
	tryRegister(&_gss_export_cred, lib, "gss_export_cred")
	tryRegister(&_gss_export_name, lib, "gss_export_name")
	tryRegister(&_gss_export_sec_context, lib, "gss_export_sec_context")
	tryRegister(&_gss_get_mic, lib, "gss_get_mic")
	tryRegister(&_gss_import_cred, lib, "gss_import_cred")
	tryRegister(&_gss_import_name, lib, "gss_import_name")
	tryRegister(&_gss_import_sec_context, lib, "gss_import_sec_context")
	tryRegister(&_gss_indicate_mechs, lib, "gss_indicate_mechs")
	tryRegister(&_gss_indicate_mechs_by_attrs, lib, "gss_indicate_mechs_by_attrs")
	tryRegister(&_gss_init_sec_context, lib, "gss_init_sec_context")
	tryRegister(&_gss_inquire_attrs_for_mech, lib, "gss_inquire_attrs_for_mech")
	tryRegister(&_gss_inquire_context, lib, "gss_inquire_context")
	tryRegister(&_gss_inquire_cred, lib, "gss_inquire_cred")
	tryRegister(&_gss_inquire_cred_by_mech, lib, "gss_inquire_cred_by_mech")
	tryRegister(&_gss_inquire_cred_by_oid, lib, "gss_inquire_cred_by_oid")
	tryRegister(&_gss_inquire_mech_for_saslname, lib, "gss_inquire_mech_for_saslname")
	tryRegister(&_gss_inquire_mechs_for_name, lib, "gss_inquire_mechs_for_name")
	tryRegister(&_gss_inquire_name, lib, "gss_inquire_name")
	tryRegister(&_gss_inquire_names_for_mech, lib, "gss_inquire_names_for_mech")
	tryRegister(&_gss_inquire_saslname_for_mech, lib, "gss_inquire_saslname_for_mech")
	tryRegister(&_gss_inquire_sec_context_by_oid, lib, "gss_inquire_sec_context_by_oid")
	tryRegister(&_gss_iter_creds, lib, "gss_iter_creds")
	tryRegister(&_gss_iter_creds_f, lib, "gss_iter_creds_f")
	tryRegister(&_gss_krb5_ccache_name, lib, "gss_krb5_ccache_name")
	tryRegister(&_gss_krb5_copy_ccache, lib, "gss_krb5_copy_ccache")
	tryRegister(&_gss_krb5_export_lucid_sec_context, lib, "gss_krb5_export_lucid_sec_context")
	tryRegister(&_gss_krb5_free_lucid_sec_context, lib, "gss_krb5_free_lucid_sec_context")
	tryRegister(&_gss_krb5_set_allowable_enctypes, lib, "gss_krb5_set_allowable_enctypes")
	tryRegister(&_gss_oid_equal, lib, "gss_oid_equal")
	tryRegister(&_gss_oid_to_str, lib, "gss_oid_to_str")
	tryRegister(&_gss_process_context_token, lib, "gss_process_context_token")
	tryRegister(&_gss_pseudo_random, lib, "gss_pseudo_random")
	tryRegister(&_gss_release_buffer, lib, "gss_release_buffer")
	tryRegister(&_gss_release_buffer_set, lib, "gss_release_buffer_set")
	tryRegister(&_gss_release_cred, lib, "gss_release_cred")
	tryRegister(&_gss_release_name, lib, "gss_release_name")
	tryRegister(&_gss_release_oid, lib, "gss_release_oid")
	tryRegister(&_gss_release_oid_set, lib, "gss_release_oid_set")
	tryRegister(&_gss_seal, lib, "gss_seal")
	tryRegister(&_gss_set_cred_option, lib, "gss_set_cred_option")
	tryRegister(&_gss_set_sec_context_option, lib, "gss_set_sec_context_option")
	tryRegister(&_gss_sign, lib, "gss_sign")
	tryRegister(&_gss_test_oid_set_member, lib, "gss_test_oid_set_member")
	tryRegister(&_gss_unseal, lib, "gss_unseal")
	tryRegister(&_gss_unwrap, lib, "gss_unwrap")
	tryRegister(&_gss_userok, lib, "gss_userok")
	tryRegister(&_gss_verify, lib, "gss_verify")
	tryRegister(&_gss_verify_mic, lib, "gss_verify_mic")
	tryRegister(&_gss_wrap, lib, "gss_wrap")
	tryRegister(&_gss_wrap_size_limit, lib, "gss_wrap_size_limit")
	tryRegister(&_GSSCreateCredentialFromUUID, lib, "GSSCreateCredentialFromUUID")
	tryRegister(&_GSSCreateError, lib, "GSSCreateError")
	tryRegister(&_GSSCreateName, lib, "GSSCreateName")
	tryRegister(&_GSSCredentialCopyName, lib, "GSSCredentialCopyName")
	tryRegister(&_GSSCredentialCopyUUID, lib, "GSSCredentialCopyUUID")
	tryRegister(&_GSSCredentialGetLifetime, lib, "GSSCredentialGetLifetime")
	tryRegister(&_gsskrb5_extract_authz_data_from_sec_context, lib, "gsskrb5_extract_authz_data_from_sec_context")
	tryRegister(&_gsskrb5_register_acceptor_identity, lib, "gsskrb5_register_acceptor_identity")
	tryRegister(&_GSSNameCreateDisplayString, lib, "GSSNameCreateDisplayString")
	tryRegister(&_krb5_gss_register_acceptor_identity, lib, "krb5_gss_register_acceptor_identity")
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



// Changes the password associated with a name.
//
// Added in macOS 10.9.
// Changes the password associated with a name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_aapl_change_password(_:_:_:_:)
func gss_aapl_change_password(name unsafe.Pointer, mech Gss_const_OID, attributes DictionaryRef, error_ unsafe.Pointer) OM_uint32 {
	return _gss_aapl_change_password(name, mech, attributes, error_)
}/* debug [functions.gen.go/function]: gss_aapl_change_password */

// Acquires a new credential using a password or certificate.
//
// Added in macOS 10.7.
// Acquires a new credential using a password or certificate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_aapl_initial_cred(_:_:_:_:_:)
func gss_aapl_initial_cred(desired_name unsafe.Pointer, desired_mech Gss_const_OID, attributes DictionaryRef, output_cred_handle unsafe.Pointer, error_ unsafe.Pointer) OM_uint32 {
	return _gss_aapl_initial_cred(desired_name, desired_mech, attributes, output_cred_handle, error_)
}/* debug [functions.gen.go/function]: gss_aapl_initial_cred */

// Accepts a security context initiated by a peer.
//
// Added in macOS 10.7.
// Accepts a security context initiated by a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_accept_sec_context(_:_:_:_:_:_:_:_:_:_:_:)
func gss_accept_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, acceptor_cred_handle unsafe.Pointer, input_token unsafe.Pointer, input_chan_bindings unsafe.Pointer, src_name unsafe.Pointer, mech_type unsafe.Pointer, output_token Gss_buffer_t, ret_flags unsafe.Pointer, time_rec unsafe.Pointer, delegated_cred_handle unsafe.Pointer) OM_uint32 {
	return _gss_accept_sec_context(minor_status, context_handle, acceptor_cred_handle, input_token, input_chan_bindings, src_name, mech_type, output_token, ret_flags, time_rec, delegated_cred_handle)
}/* debug [functions.gen.go/function]: gss_accept_sec_context */

// Acquires a credential for use in establishing a security context.
//
// Added in macOS 10.7.
// Acquires a credential for use in establishing a security context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_acquire_cred(_:_:_:_:_:_:_:_:)
func gss_acquire_cred(minor_status unsafe.Pointer, desired_name unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec unsafe.Pointer) OM_uint32 {
	return _gss_acquire_cred(minor_status, desired_name, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec)
}/* debug [functions.gen.go/function]: gss_acquire_cred */

// Acquires a credential for use in establishing a security context using a password.
//
// Added in macOS 10.7.
// Acquires a credential for use in establishing a security context using a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_acquire_cred_with_password(_:_:_:_:_:_:_:_:_:)
func gss_acquire_cred_with_password(minor_status unsafe.Pointer, desired_name unsafe.Pointer, password unsafe.Pointer, time_req OM_uint32, desired_mechs unsafe.Pointer, cred_usage Gss_cred_usage_t, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, time_rec unsafe.Pointer) OM_uint32 {
	return _gss_acquire_cred_with_password(minor_status, desired_name, password, time_req, desired_mechs, cred_usage, output_cred_handle, actual_mechs, time_rec)
}/* debug [functions.gen.go/function]: gss_acquire_cred_with_password */

// Copies the contents of a buffer into a buffer set.
//
// Added in macOS 10.7.
// Copies the contents of a buffer into a buffer set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_add_buffer_set_member(_:_:_:)
func gss_add_buffer_set_member(minor_status unsafe.Pointer, member_buffer unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32 {
	return _gss_add_buffer_set_member(minor_status, member_buffer, buffer_set)
}/* debug [functions.gen.go/function]: gss_add_buffer_set_member */

// Adds a new credential element to an existing credential.
//
// Added in macOS 10.7.
// Adds a new credential element to an existing credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_add_cred(_:_:_:_:_:_:_:_:_:_:_:)
func gss_add_cred(minor_status unsafe.Pointer, input_cred_handle unsafe.Pointer, desired_name unsafe.Pointer, desired_mech unsafe.Pointer, cred_usage Gss_cred_usage_t, initiator_time_req OM_uint32, acceptor_time_req OM_uint32, output_cred_handle unsafe.Pointer, actual_mechs unsafe.Pointer, initiator_time_rec unsafe.Pointer, acceptor_time_rec unsafe.Pointer) OM_uint32 {
	return _gss_add_cred(minor_status, input_cred_handle, desired_name, desired_mech, cred_usage, initiator_time_req, acceptor_time_req, output_cred_handle, actual_mechs, initiator_time_rec, acceptor_time_rec)
}/* debug [functions.gen.go/function]: gss_add_cred */

// Adds an object identifier into an OID set.
//
// Added in macOS 10.7.
// Adds an object identifier into an OID set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_add_oid_set_member(_:_:_:)
func gss_add_oid_set_member(minor_status unsafe.Pointer, member_oid Gss_const_OID, oid_set unsafe.Pointer) OM_uint32 {
	return _gss_add_oid_set_member(minor_status, member_oid, oid_set)
}/* debug [functions.gen.go/function]: gss_add_oid_set_member */

// Converts an internal name into a mechanism name.
//
// Added in macOS 10.7.
// Converts an internal name into a mechanism name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_canonicalize_name(_:_:_:_:)
func gss_canonicalize_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, mech_type unsafe.Pointer, output_name unsafe.Pointer) OM_uint32 {
	return _gss_canonicalize_name(minor_status, input_name, mech_type, output_name)
}/* debug [functions.gen.go/function]: gss_canonicalize_name */

// Returns a flag that indicates if two names in internal name format refer to the same entity.
//
// Added in macOS 10.7.
// Returns a flag that indicates if two names in internal name format refer to the same entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_compare_name(_:_:_:_:)
func gss_compare_name(minor_status unsafe.Pointer, name1_arg unsafe.Pointer, name2_arg unsafe.Pointer, name_equal []int) OM_uint32 {
	return _gss_compare_name(minor_status, name1_arg, name2_arg, name_equal)
}/* debug [functions.gen.go/function]: gss_compare_name */

// Returns the amount of time remaining before a context expires.
//
// Added in macOS 10.7.
// Returns the amount of time remaining before a context expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_context_time(_:_:_:)
func gss_context_time(minor_status unsafe.Pointer, context_handle unsafe.Pointer, time_rec unsafe.Pointer) OM_uint32 {
	return _gss_context_time(minor_status, context_handle, time_rec)
}/* debug [functions.gen.go/function]: gss_context_time */

// Allocates an empty buffer set descriptor that you use to manage an array of buffers.
//
// Added in macOS 10.7.
// Allocates an empty buffer set descriptor that you use to manage an array of buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_buffer_set(_:_:)
func gss_create_empty_buffer_set(minor_status unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32 {
	return _gss_create_empty_buffer_set(minor_status, buffer_set)
}/* debug [functions.gen.go/function]: gss_create_empty_buffer_set */

// Allocates a new, empty set to hold object identifiers.
//
// Added in macOS 10.7.
// Allocates a new, empty set to hold object identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_create_empty_oid_set(_:_:)
func gss_create_empty_oid_set(minor_status unsafe.Pointer, oid_set unsafe.Pointer) OM_uint32 {
	return _gss_create_empty_oid_set(minor_status, oid_set)
}/* debug [functions.gen.go/function]: gss_create_empty_oid_set */

// Returns a token encapsulated in a buffer.
//
// Added in macOS 10.7.
// Returns a token encapsulated in a buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_decapsulate_token(_:_:_:)
func gss_decapsulate_token(input_token Gss_const_buffer_t, oid Gss_const_OID, output_token Gss_buffer_t) OM_uint32 {
	return _gss_decapsulate_token(input_token, oid, output_token)
}/* debug [functions.gen.go/function]: gss_decapsulate_token */

// Deletes a security context.
//
// Added in macOS 10.7.
// Deletes a security context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_delete_sec_context(_:_:_:)
func gss_delete_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, output_token Gss_buffer_t) OM_uint32 {
	return _gss_delete_sec_context(minor_status, context_handle, output_token)
}/* debug [functions.gen.go/function]: gss_delete_sec_context */

// Purges a credential from memory.
//
// Added in macOS 10.7.
// Purges a credential from memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_destroy_cred(_:_:)
func gss_destroy_cred(min_stat unsafe.Pointer, cred_handle unsafe.Pointer) OM_uint32 {
	return _gss_destroy_cred(min_stat, cred_handle)
}/* debug [functions.gen.go/function]: gss_destroy_cred */

// Returns a human-readable name and description of a mechanism attribute.
//
// Added in macOS 10.7.
// Returns a human-readable name and description of a mechanism attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_display_mech_attr(_:_:_:_:_:)
func gss_display_mech_attr(minor_status unsafe.Pointer, mech_attr Gss_const_OID, name Gss_buffer_t, short_desc Gss_buffer_t, long_desc Gss_buffer_t) OM_uint32 {
	return _gss_display_mech_attr(minor_status, mech_attr, name, short_desc, long_desc)
}/* debug [functions.gen.go/function]: gss_display_mech_attr */

// Converts a name in the internal format to an octet string and the associated name type.
//
// Added in macOS 10.7.
// Converts a name in the internal format to an octet string and the associated name type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_display_name(_:_:_:_:)
func gss_display_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, output_name_buffer Gss_buffer_t, output_name_type unsafe.Pointer) OM_uint32 {
	return _gss_display_name(minor_status, input_name, output_name_buffer, output_name_type)
}/* debug [functions.gen.go/function]: gss_display_name */

// Returns a human readable string for a status code.
//
// Added in macOS 10.7.
// Returns a human readable string for a status code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_display_status(_:_:_:_:_:_:)
func gss_display_status(minor_status unsafe.Pointer, status_value OM_uint32, status_type int, mech_type unsafe.Pointer, message_content unsafe.Pointer, status_string Gss_buffer_t) OM_uint32 {
	return _gss_display_status(minor_status, status_value, status_type, mech_type, message_content, status_string)
}/* debug [functions.gen.go/function]: gss_display_status */

// Returns a copy of an internal name.
//
// Added in macOS 10.7.
// Returns a copy of an internal name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_duplicate_name(_:_:_:)
func gss_duplicate_name(minor_status unsafe.Pointer, src_name unsafe.Pointer, dest_name unsafe.Pointer) OM_uint32 {
	return _gss_duplicate_name(minor_status, src_name, dest_name)
}/* debug [functions.gen.go/function]: gss_duplicate_name */

// Copies an OID into a new object.

// Copies an OID into a new object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_duplicate_oid(_:_:_:)
func gss_duplicate_oid(minor_status unsafe.Pointer, src_oid Gss_OID, dest_oid unsafe.Pointer) OM_uint32 {
	return _gss_duplicate_oid(minor_status, src_oid, dest_oid)
}/* debug [functions.gen.go/function]: gss_duplicate_oid */

// Returns a buffer encapsulating the given token.
//
// Added in macOS 10.7.
// Returns a buffer encapsulating the given token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_encapsulate_token(_:_:_:)
func gss_encapsulate_token(input_token Gss_const_buffer_t, oid Gss_const_OID, output_token Gss_buffer_t) OM_uint32 {
	return _gss_encapsulate_token(input_token, oid, output_token)
}/* debug [functions.gen.go/function]: gss_encapsulate_token */

// Exports a credential to a token.
//
// Added in macOS 10.7.
// Exports a credential to a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_export_cred(_:_:_:)
func gss_export_cred(minor_status unsafe.Pointer, cred_handle Gss_cred_id_t, token Gss_buffer_t) OM_uint32 {
	return _gss_export_cred(minor_status, cred_handle, token)
}/* debug [functions.gen.go/function]: gss_export_cred */

// Returns a mechanism name in contiguous octet format.
//
// Added in macOS 10.7.
// Returns a mechanism name in contiguous octet format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_export_name(_:_:_:)
func gss_export_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, exported_name Gss_buffer_t) OM_uint32 {
	return _gss_export_name(minor_status, input_name, exported_name)
}/* debug [functions.gen.go/function]: gss_export_name */

// Transfers a security context to another process.
//
// Added in macOS 10.7.
// Transfers a security context to another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_export_sec_context(_:_:_:)
func gss_export_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, interprocess_token Gss_buffer_t) OM_uint32 {
	return _gss_export_sec_context(minor_status, context_handle, interprocess_token)
}/* debug [functions.gen.go/function]: gss_export_sec_context */

// Returns a token that contains the MIC for a message.
//
// Added in macOS 10.7.
// Returns a token that contains the MIC for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_get_mic(_:_:_:_:_:)
func gss_get_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, qop_req Gss_qop_t, message_buffer unsafe.Pointer, message_token Gss_buffer_t) OM_uint32 {
	return _gss_get_mic(minor_status, context_handle, qop_req, message_buffer, message_token)
}/* debug [functions.gen.go/function]: gss_get_mic */

// Imports a credential from a token.
//
// Added in macOS 10.7.
// Imports a credential from a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_import_cred(_:_:_:)
func gss_import_cred(minor_status unsafe.Pointer, token Gss_buffer_t, cred_handle unsafe.Pointer) OM_uint32 {
	return _gss_import_cred(minor_status, token, cred_handle)
}/* debug [functions.gen.go/function]: gss_import_cred */

// Converts a name in contiguous octet format to the internal name format.
//
// Added in macOS 10.7.
// Converts a name in contiguous octet format to the internal name format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_import_name(_:_:_:_:)
func gss_import_name(minor_status unsafe.Pointer, input_name_buffer unsafe.Pointer, input_name_type Gss_const_OID, output_name unsafe.Pointer) OM_uint32 {
	return _gss_import_name(minor_status, input_name_buffer, input_name_type, output_name)
}/* debug [functions.gen.go/function]: gss_import_name */

// Imports a security context from another process.
//
// Added in macOS 10.7.
// Imports a security context from another process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_import_sec_context(_:_:_:)
func gss_import_sec_context(minor_status unsafe.Pointer, interprocess_token unsafe.Pointer, context_handle unsafe.Pointer) OM_uint32 {
	return _gss_import_sec_context(minor_status, interprocess_token, context_handle)
}/* debug [functions.gen.go/function]: gss_import_sec_context */

// Returns the list of supported underlying security mechanisms.
//
// Added in macOS 10.7.
// Returns the list of supported underlying security mechanisms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_indicate_mechs(_:_:)
func gss_indicate_mechs(minor_status unsafe.Pointer, mech_set unsafe.Pointer) OM_uint32 {
	return _gss_indicate_mechs(minor_status, mech_set)
}/* debug [functions.gen.go/function]: gss_indicate_mechs */

// Returns the set of mechanisms that fulfill the given criteria.
//
// Added in macOS 10.10.
// Returns the set of mechanisms that fulfill the given criteria.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_indicate_mechs_by_attrs(_:_:_:_:_:)
func gss_indicate_mechs_by_attrs(minor_status unsafe.Pointer, desired_mech_attrs Gss_const_OID_set, except_mech_attrs Gss_const_OID_set, critical_mech_attrs Gss_const_OID_set, mechs unsafe.Pointer) OM_uint32 {
	return _gss_indicate_mechs_by_attrs(minor_status, desired_mech_attrs, except_mech_attrs, critical_mech_attrs, mechs)
}/* debug [functions.gen.go/function]: gss_indicate_mechs_by_attrs */

// Initiates a security context with a peer.
//
// Added in macOS 10.7.
// Initiates a security context with a peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_init_sec_context(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func gss_init_sec_context(minor_status unsafe.Pointer, initiator_cred_handle unsafe.Pointer, context_handle unsafe.Pointer, target_name unsafe.Pointer, input_mech_type unsafe.Pointer, req_flags OM_uint32, time_req OM_uint32, input_chan_bindings unsafe.Pointer, input_token unsafe.Pointer, actual_mech_type unsafe.Pointer, output_token Gss_buffer_t, ret_flags unsafe.Pointer, time_rec unsafe.Pointer) OM_uint32 {
	return _gss_init_sec_context(minor_status, initiator_cred_handle, context_handle, target_name, input_mech_type, req_flags, time_req, input_chan_bindings, input_token, actual_mech_type, output_token, ret_flags, time_rec)
}/* debug [functions.gen.go/function]: gss_init_sec_context */

// Returns the supported attributes for one or all mechanisms.
//
// Added in macOS 10.7.
// Returns the supported attributes for one or all mechanisms.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_attrs_for_mech(_:_:_:_:)
func gss_inquire_attrs_for_mech(minor_status unsafe.Pointer, mech Gss_const_OID, mech_attr unsafe.Pointer, known_mech_attrs unsafe.Pointer) OM_uint32 {
	return _gss_inquire_attrs_for_mech(minor_status, mech, mech_attr, known_mech_attrs)
}/* debug [functions.gen.go/function]: gss_inquire_attrs_for_mech */

// Returns information about a security context.
//
// Added in macOS 10.7.
// Returns information about a security context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_context(_:_:_:_:_:_:_:_:_:)
func gss_inquire_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, src_name unsafe.Pointer, targ_name unsafe.Pointer, lifetime_rec unsafe.Pointer, mech_type unsafe.Pointer, ctx_flags unsafe.Pointer, locally_initiated []int, xopen []int) OM_uint32 {
	return _gss_inquire_context(minor_status, context_handle, src_name, targ_name, lifetime_rec, mech_type, ctx_flags, locally_initiated, xopen)
}/* debug [functions.gen.go/function]: gss_inquire_context */

// Obtains information about a credential.
//
// Added in macOS 10.7.
// Obtains information about a credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_cred(_:_:_:_:_:_:)
func gss_inquire_cred(minor_status unsafe.Pointer, cred_handle unsafe.Pointer, name_ret unsafe.Pointer, lifetime unsafe.Pointer, cred_usage unsafe.Pointer, mechanisms unsafe.Pointer) OM_uint32 {
	return _gss_inquire_cred(minor_status, cred_handle, name_ret, lifetime, cred_usage, mechanisms)
}/* debug [functions.gen.go/function]: gss_inquire_cred */

// Obtains per-mechanism information about a credential.
//
// Added in macOS 10.7.
// Obtains per-mechanism information about a credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_cred_by_mech(_:_:_:_:_:_:_:)
func gss_inquire_cred_by_mech(minor_status unsafe.Pointer, cred_handle unsafe.Pointer, mech_type unsafe.Pointer, cred_name unsafe.Pointer, initiator_lifetime unsafe.Pointer, acceptor_lifetime unsafe.Pointer, cred_usage unsafe.Pointer) OM_uint32 {
	return _gss_inquire_cred_by_mech(minor_status, cred_handle, mech_type, cred_name, initiator_lifetime, acceptor_lifetime, cred_usage)
}/* debug [functions.gen.go/function]: gss_inquire_cred_by_mech */

// Inquires about a particular characteristic of a credential.
//
// Added in macOS 10.7.
// Inquires about a particular characteristic of a credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_cred_by_oid(_:_:_:_:)
func gss_inquire_cred_by_oid(minor_status unsafe.Pointer, cred_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32 {
	return _gss_inquire_cred_by_oid(minor_status, cred_handle, desired_object, data_set)
}/* debug [functions.gen.go/function]: gss_inquire_cred_by_oid */

// Returns the GSS-API mechanism identifier for a given Simple Authentication and Security Layer (SASL) protocol name.
//
// Added in macOS 10.10.
// Returns the GSS-API mechanism identifier for a given Simple Authentication and Security Layer (SASL) protocol name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_mech_for_saslname(_:_:_:)
func gss_inquire_mech_for_saslname(minor_status unsafe.Pointer, sasl_mech_name unsafe.Pointer, mech_type unsafe.Pointer) OM_uint32 {
	return _gss_inquire_mech_for_saslname(minor_status, sasl_mech_name, mech_type)
}/* debug [functions.gen.go/function]: gss_inquire_mech_for_saslname */

// Returns a list of mechanisms that support a particular name type.
//
// Added in macOS 10.7.
// Returns a list of mechanisms that support a particular name type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_mechs_for_name(_:_:_:)
func gss_inquire_mechs_for_name(minor_status unsafe.Pointer, input_name unsafe.Pointer, mech_types unsafe.Pointer) OM_uint32 {
	return _gss_inquire_mechs_for_name(minor_status, input_name, mech_types)
}/* debug [functions.gen.go/function]: gss_inquire_mechs_for_name */

// Returns information about a name.
//
// Added in macOS 10.7.
// Returns information about a name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_name(_:_:_:_:_:)
func gss_inquire_name(minor_status unsafe.Pointer, input_name Gss_name_t, name_is_MN []int, MN_mech unsafe.Pointer, attrs unsafe.Pointer) OM_uint32 {
	return _gss_inquire_name(minor_status, input_name, name_is_MN, MN_mech, attrs)
}/* debug [functions.gen.go/function]: gss_inquire_name */

// Returns a list of name types that a given mechanism supports.
//
// Added in macOS 10.7.
// Returns a list of name types that a given mechanism supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_names_for_mech(_:_:_:)
func gss_inquire_names_for_mech(minor_status unsafe.Pointer, mechanism Gss_const_OID, name_types unsafe.Pointer) OM_uint32 {
	return _gss_inquire_names_for_mech(minor_status, mechanism, name_types)
}/* debug [functions.gen.go/function]: gss_inquire_names_for_mech */

// Returns the Simple Authentication and Security Layer (SASL) protocol name for a given GSS-API mechanism.
//
// Added in macOS 10.10.
// Returns the Simple Authentication and Security Layer (SASL) protocol name for a given GSS-API mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_saslname_for_mech(_:_:_:_:_:)
func gss_inquire_saslname_for_mech(minor_status unsafe.Pointer, desired_mech unsafe.Pointer, sasl_mech_name Gss_buffer_t, mech_name Gss_buffer_t, mech_description Gss_buffer_t) OM_uint32 {
	return _gss_inquire_saslname_for_mech(minor_status, desired_mech, sasl_mech_name, mech_name, mech_description)
}/* debug [functions.gen.go/function]: gss_inquire_saslname_for_mech */

// Returns information about a particular part of a context.
//
// Added in macOS 10.7.
// Returns information about a particular part of a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_inquire_sec_context_by_oid(_:_:_:_:)
func gss_inquire_sec_context_by_oid(minor_status unsafe.Pointer, context_handle unsafe.Pointer, desired_object unsafe.Pointer, data_set unsafe.Pointer) OM_uint32 {
	return _gss_inquire_sec_context_by_oid(minor_status, context_handle, desired_object, data_set)
}/* debug [functions.gen.go/function]: gss_inquire_sec_context_by_oid */

// Iterates over all credentials.
//
// Added in macOS 10.7.
// Iterates over all credentials.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_iter_creds(_:_:_:_:)
func gss_iter_creds(min_stat unsafe.Pointer, flags OM_uint32, mech Gss_const_OID) OM_uint32 {
	return _gss_iter_creds(min_stat, flags, mech)
}/* debug [functions.gen.go/function]: gss_iter_creds */

// Iterates over all credentials with a user context.
//
// Added in macOS 10.7.
// Iterates over all credentials with a user context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_iter_creds_f(_:_:_:_:_:)
func gss_iter_creds_f(min_stat unsafe.Pointer, flags OM_uint32, mech Gss_const_OID, userctx unsafe.Pointer) OM_uint32 {
	return _gss_iter_creds_f(min_stat, flags, mech, userctx)
}/* debug [functions.gen.go/function]: gss_iter_creds_f */

// Sets the internal Kerberos 5 credential cache name.
//
// Added in macOS 10.7.
// Sets the internal Kerberos 5 credential cache name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_ccache_name(_:_:_:)
func gss_krb5_ccache_name(minor_status unsafe.Pointer, name unsafe.Pointer, out_name unsafe.Pointer) OM_uint32 {
	return _gss_krb5_ccache_name(minor_status, name, out_name)
}/* debug [functions.gen.go/function]: gss_krb5_ccache_name */

// Copies Kerberos 5 credentials into the passed cache.

// Copies Kerberos 5 credentials into the passed cache.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_copy_ccache(_:_:_:)
func gss_krb5_copy_ccache(minor_status unsafe.Pointer, cred Gss_cred_id_t, out unsafe.Pointer) OM_uint32 {
	return _gss_krb5_copy_ccache(minor_status, cred, out)
}/* debug [functions.gen.go/function]: gss_krb5_copy_ccache */

// Returns a non-opaque version of the internal context information.
//
// Added in macOS 10.7.
// Returns a non-opaque version of the internal context information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_export_lucid_sec_context(_:_:_:_:)
func gss_krb5_export_lucid_sec_context(minor_status unsafe.Pointer, context_handle unsafe.Pointer, version OM_uint32, rctx unsafe.Pointer) OM_uint32 {
	return _gss_krb5_export_lucid_sec_context(minor_status, context_handle, version, rctx)
}/* debug [functions.gen.go/function]: gss_krb5_export_lucid_sec_context */

// Frees allocated storage associated with an exported context.
//
// Added in macOS 10.7.
// Frees allocated storage associated with an exported context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_free_lucid_sec_context(_:_:)
func gss_krb5_free_lucid_sec_context(minor_status unsafe.Pointer, c unsafe.Pointer) OM_uint32 {
	return _gss_krb5_free_lucid_sec_context(minor_status, c)
}/* debug [functions.gen.go/function]: gss_krb5_free_lucid_sec_context */

// Limits the keys that can be exported to the specified types.
//
// Added in macOS 10.7.
// Limits the keys that can be exported to the specified types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_krb5_set_allowable_enctypes(_:_:_:_:)
func gss_krb5_set_allowable_enctypes(minor_status unsafe.Pointer, cred Gss_cred_id_t, num_enctypes OM_uint32, enctypes unsafe.Pointer) OM_uint32 {
	return _gss_krb5_set_allowable_enctypes(minor_status, cred, num_enctypes, enctypes)
}/* debug [functions.gen.go/function]: gss_krb5_set_allowable_enctypes */

// Returns a flag that indicates whether two object identifiers are the same.
//
// Added in macOS 10.7.
// Returns a flag that indicates whether two object identifiers are the same.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_equal(_:_:)
func gss_oid_equal(a Gss_const_OID, b Gss_const_OID) int {
	return _gss_oid_equal(a, b)
}/* debug [functions.gen.go/function]: gss_oid_equal */

// Converts an OID object to a human-readable string.
//
// Added in macOS 10.7.
// Converts an OID object to a human-readable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_oid_to_str(_:_:_:)
func gss_oid_to_str(minor_status unsafe.Pointer, oid Gss_OID, oid_str Gss_buffer_t) OM_uint32 {
	return _gss_oid_to_str(minor_status, oid, oid_str)
}/* debug [functions.gen.go/function]: gss_oid_to_str */

// Processes a token from a peer asynchronously.
//
// Added in macOS 10.7.
// Processes a token from a peer asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_process_context_token(_:_:_:)
func gss_process_context_token(minor_status unsafe.Pointer, context_handle unsafe.Pointer, token_buffer unsafe.Pointer) OM_uint32 {
	return _gss_process_context_token(minor_status, context_handle, token_buffer)
}/* debug [functions.gen.go/function]: gss_process_context_token */

// Returns a pseudo-random byte stream for keying.
//
// Added in macOS 10.7.
// Returns a pseudo-random byte stream for keying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_pseudo_random(_:_:_:_:_:_:)
func gss_pseudo_random(minor_status unsafe.Pointer, context Gss_ctx_id_t, prf_key int, prf_in unsafe.Pointer, desired_output_len unsafe.Pointer, prf_out Gss_buffer_t) OM_uint32 {
	return _gss_pseudo_random(minor_status, context, prf_key, prf_in, desired_output_len, prf_out)
}/* debug [functions.gen.go/function]: gss_pseudo_random */

// Frees the memory associated with a single buffer descriptor.
//
// Added in macOS 10.7.
// Frees the memory associated with a single buffer descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_buffer(_:_:)
func gss_release_buffer(minor_status unsafe.Pointer, buffer Gss_buffer_t) OM_uint32 {
	return _gss_release_buffer(minor_status, buffer)
}/* debug [functions.gen.go/function]: gss_release_buffer */

// Frees the memory associated with a buffer set descriptor and all the buffers it contains.
//
// Added in macOS 10.7.
// Frees the memory associated with a buffer set descriptor and all the buffers it contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_buffer_set(_:_:)
func gss_release_buffer_set(minor_status unsafe.Pointer, buffer_set unsafe.Pointer) OM_uint32 {
	return _gss_release_buffer_set(minor_status, buffer_set)
}/* debug [functions.gen.go/function]: gss_release_buffer_set */

// Releases the memory of a credential.
//
// Added in macOS 10.7.
// Releases the memory of a credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_cred(_:_:)
func gss_release_cred(minor_status unsafe.Pointer, cred_handle unsafe.Pointer) OM_uint32 {
	return _gss_release_cred(minor_status, cred_handle)
}/* debug [functions.gen.go/function]: gss_release_cred */

// Frees the resources associated with a name object.
//
// Added in macOS 10.7.
// Frees the resources associated with a name object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_name(_:_:)
func gss_release_name(minor_status unsafe.Pointer, input_name unsafe.Pointer) OM_uint32 {
	return _gss_release_name(minor_status, input_name)
}/* debug [functions.gen.go/function]: gss_release_name */

// Releases the memory associated with an object identifier.

// Releases the memory associated with an object identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_oid(_:_:)
func gss_release_oid(minor_status unsafe.Pointer, oid unsafe.Pointer) OM_uint32 {
	return _gss_release_oid(minor_status, oid)
}/* debug [functions.gen.go/function]: gss_release_oid */

// Releases the memory associated with an OID set.
//
// Added in macOS 10.7.
// Releases the memory associated with an OID set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_release_oid_set(_:_:)
func gss_release_oid_set(minor_status unsafe.Pointer, set unsafe.Pointer) OM_uint32 {
	return _gss_release_oid_set(minor_status, set)
}/* debug [functions.gen.go/function]: gss_release_oid_set */

// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.

// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_seal(_:_:_:_:_:_:_:)
func gss_seal(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, conf_req_flag int, qop_req int, input_message_buffer Gss_buffer_t, conf_state []int, output_message_buffer Gss_buffer_t) OM_uint32 {
	return _gss_seal(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
}/* debug [functions.gen.go/function]: gss_seal */

// Changes a credential option.
//
// Added in macOS 10.7.
// Changes a credential option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_set_cred_option(_:_:_:_:)
func gss_set_cred_option(minor_status unsafe.Pointer, cred_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32 {
	return _gss_set_cred_option(minor_status, cred_handle, object, value)
}/* debug [functions.gen.go/function]: gss_set_cred_option */

// Sets an option on a context.
//
// Added in macOS 10.7.
// Sets an option on a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_set_sec_context_option(_:_:_:_:)
func gss_set_sec_context_option(minor_status unsafe.Pointer, context_handle unsafe.Pointer, object unsafe.Pointer, value unsafe.Pointer) OM_uint32 {
	return _gss_set_sec_context_option(minor_status, context_handle, object, value)
}/* debug [functions.gen.go/function]: gss_set_sec_context_option */

// Returns a digital signature for a message.

// Returns a digital signature for a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_sign(_:_:_:_:_:)
func gss_sign(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, qop_req int, message_buffer Gss_buffer_t, message_token Gss_buffer_t) OM_uint32 {
	return _gss_sign(minor_status, context_handle, qop_req, message_buffer, message_token)
}/* debug [functions.gen.go/function]: gss_sign */

// Returns a flag that indicates if an OID is present in an OID set.
//
// Added in macOS 10.7.
// Returns a flag that indicates if an OID is present in an OID set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_test_oid_set_member(_:_:_:_:)
func gss_test_oid_set_member(minor_status unsafe.Pointer, member Gss_const_OID, set unsafe.Pointer, present []int) OM_uint32 {
	return _gss_test_oid_set_member(minor_status, member, set, present)
}/* debug [functions.gen.go/function]: gss_test_oid_set_member */

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_unseal(_:_:_:_:_:_:)
func gss_unseal(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, input_message_buffer Gss_buffer_t, output_message_buffer Gss_buffer_t, conf_state []int, qop_state []int) OM_uint32 {
	return _gss_unseal(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
}/* debug [functions.gen.go/function]: gss_unseal */

// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// Added in macOS 10.7.
// Returns the original version of a secure message by optionally decrypting it and then extracting and verifying the attached MIC.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_unwrap(_:_:_:_:_:_:)
func gss_unwrap(minor_status unsafe.Pointer, context_handle unsafe.Pointer, input_message_buffer unsafe.Pointer, output_message_buffer Gss_buffer_t, conf_state []int, qop_state unsafe.Pointer) OM_uint32 {
	return _gss_unwrap(minor_status, context_handle, input_message_buffer, output_message_buffer, conf_state, qop_state)
}/* debug [functions.gen.go/function]: gss_unwrap */

// Returns a flag that indicates if a given user is authorized.
//
// Added in macOS 10.9.
// Returns a flag that indicates if a given user is authorized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_userok(_:_:)
func gss_userok(name unsafe.Pointer, user unsafe.Pointer) int {
	return _gss_userok(name, user)
}/* debug [functions.gen.go/function]: gss_userok */

// Returns a flag that indicates the integrity of a message’s digital signature.

// Returns a flag that indicates the integrity of a message’s digital signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_verify(_:_:_:_:_:)
func gss_verify(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, message_buffer Gss_buffer_t, token_buffer Gss_buffer_t, qop_state []int) OM_uint32 {
	return _gss_verify(minor_status, context_handle, message_buffer, token_buffer, qop_state)
}/* debug [functions.gen.go/function]: gss_verify */

// Returns an indication of whether the integrity of a message is intact, given its MIC token.
//
// Added in macOS 10.7.
// Returns an indication of whether the integrity of a message is intact, given its MIC token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_verify_mic(_:_:_:_:_:)
func gss_verify_mic(minor_status unsafe.Pointer, context_handle unsafe.Pointer, message_buffer unsafe.Pointer, token_buffer unsafe.Pointer, qop_state unsafe.Pointer) OM_uint32 {
	return _gss_verify_mic(minor_status, context_handle, message_buffer, token_buffer, qop_state)
}/* debug [functions.gen.go/function]: gss_verify_mic */

// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// Added in macOS 10.7.
// Returns a secure message created by calculating and attaching a MIC to the input message, and then optionally encrypting it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_wrap(_:_:_:_:_:_:_:)
func gss_wrap(minor_status unsafe.Pointer, context_handle unsafe.Pointer, conf_req_flag int, qop_req Gss_qop_t, input_message_buffer unsafe.Pointer, conf_state []int, output_message_buffer Gss_buffer_t) OM_uint32 {
	return _gss_wrap(minor_status, context_handle, conf_req_flag, qop_req, input_message_buffer, conf_state, output_message_buffer)
}/* debug [functions.gen.go/function]: gss_wrap */

// Returns the largest allowable wrap size for a given set of constraints.
//
// Added in macOS 10.7.
// Returns the largest allowable wrap size for a given set of constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gss_wrap_size_limit(_:_:_:_:_:_:)
func gss_wrap_size_limit(minor_status unsafe.Pointer, context_handle unsafe.Pointer, conf_req_flag int, qop_req Gss_qop_t, req_output_size OM_uint32, max_input_size unsafe.Pointer) OM_uint32 {
	return _gss_wrap_size_limit(minor_status, context_handle, conf_req_flag, qop_req, req_output_size, max_input_size)
}/* debug [functions.gen.go/function]: gss_wrap_size_limit */

// Creates a credential from a universally unique identifier.
//
// Added in macOS 10.9.
// Creates a credential from a universally unique identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCreateCredentialFromUUID(_:)
func GSSCreateCredentialFromUUID(uuid UUIDRef) Gss_cred_id_t {
	return _GSSCreateCredentialFromUUID(uuid)
}/* debug [functions.gen.go/function]: GSSCreateCredentialFromUUID */

// Returns an error object based on GSS-API major and minor status codes.
//
// Added in macOS 10.10.
// Returns an error object based on GSS-API major and minor status codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCreateError(_:_:_:)
func GSSCreateError(mech Gss_const_OID, major_status OM_uint32, minor_status OM_uint32) ErrorRef {
	return _GSSCreateError(mech, major_status, minor_status)
}/* debug [functions.gen.go/function]: GSSCreateError */

// Returns a GSS name given a buffer and a type.
//
// Added in macOS 10.9.
// Returns a GSS name given a buffer and a type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCreateName(_:_:_:)
func GSSCreateName(name TypeRef, name_type Gss_const_OID, error_ unsafe.Pointer) Gss_name_t {
	return _GSSCreateName(name, name_type, error_)
}/* debug [functions.gen.go/function]: GSSCreateName */

// Returns the name describing the credential.
//
// Added in macOS 10.9.
// Returns the name describing the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCredentialCopyName(_:)
func GSSCredentialCopyName(cred Gss_cred_id_t) Gss_name_t {
	return _GSSCredentialCopyName(cred)
}/* debug [functions.gen.go/function]: GSSCredentialCopyName */

// Returns a copy of the universally unique identifier corresponding to a GSS credential.
//
// Added in macOS 10.9.
// Returns a copy of the universally unique identifier corresponding to a GSS credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCredentialCopyUUID(_:)
func GSSCredentialCopyUUID(credential Gss_cred_id_t) UUIDRef {
	return _GSSCredentialCopyUUID(credential)
}/* debug [functions.gen.go/function]: GSSCredentialCopyUUID */

// Returns the remaining time in seconds before the credential expires.
//
// Added in macOS 10.9.
// Returns the remaining time in seconds before the credential expires.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSCredentialGetLifetime(_:)
func GSSCredentialGetLifetime(cred Gss_cred_id_t) OM_uint32 {
	return _GSSCredentialGetLifetime(cred)
}/* debug [functions.gen.go/function]: GSSCredentialGetLifetime */

// Extracts Kerberos authorization data stored within the context.
//
// Added in macOS 10.7.
// Extracts Kerberos authorization data stored within the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_extract_authz_data_from_sec_context(_:_:_:_:)
func gsskrb5_extract_authz_data_from_sec_context(minor_status unsafe.Pointer, context_handle Gss_ctx_id_t, ad_type int, ad_data Gss_buffer_t) OM_uint32 {
	return _gsskrb5_extract_authz_data_from_sec_context(minor_status, context_handle, ad_type, ad_data)
}/* debug [functions.gen.go/function]: gsskrb5_extract_authz_data_from_sec_context */

// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// Added in macOS 10.7.
// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/gsskrb5_register_acceptor_identity(_:)
func gsskrb5_register_acceptor_identity(identity unsafe.Pointer) OM_uint32 {
	return _gsskrb5_register_acceptor_identity(identity)
}/* debug [functions.gen.go/function]: gsskrb5_register_acceptor_identity */

// Returns a string suitable for displaying to the user from a GSS name.
//
// Added in macOS 10.9.
// Returns a string suitable for displaying to the user from a GSS name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/GSSNameCreateDisplayString(_:)
func GSSNameCreateDisplayString(name Gss_name_t) StringRef {
	return _GSSNameCreateDisplayString(name)
}/* debug [functions.gen.go/function]: GSSNameCreateDisplayString */

// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// Added in macOS 10.7.
// Sets the Kerberos 5 file-based key that the acceptor will use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GSS/krb5_gss_register_acceptor_identity(_:)
func krb5_gss_register_acceptor_identity(identity unsafe.Pointer) OM_uint32 {
	return _krb5_gss_register_acceptor_identity(identity)
}/* debug [functions.gen.go/function]: krb5_gss_register_acceptor_identity */




