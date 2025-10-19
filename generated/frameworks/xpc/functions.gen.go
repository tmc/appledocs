// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// XPC Functions (205 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_launch_activate_socket func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_launch_data_alloc func(unsafe.Pointer) unsafe.Pointer
	_launch_data_array_get_count func(unsafe.Pointer) uintptr
	_launch_data_array_get_index func(unsafe.Pointer, uintptr) unsafe.Pointer
	_launch_data_array_set_index func(unsafe.Pointer, unsafe.Pointer, uintptr) bool
	_launch_data_copy func(unsafe.Pointer) unsafe.Pointer
	_launch_data_dict_get_count func(unsafe.Pointer) uintptr
	_launch_data_dict_insert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_launch_data_dict_iterate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_launch_data_dict_lookup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_launch_data_dict_remove func(unsafe.Pointer, unsafe.Pointer) bool
	_launch_data_free func(unsafe.Pointer)
	_launch_data_get_bool func(unsafe.Pointer) bool
	_launch_data_get_errno func(unsafe.Pointer) int
	_launch_data_get_fd func(unsafe.Pointer) int
	_launch_data_get_integer func(unsafe.Pointer) unsafe.Pointer
	_launch_data_get_machport func(unsafe.Pointer) unsafe.Pointer
	_launch_data_get_opaque func(unsafe.Pointer) unsafe.Pointer
	_launch_data_get_opaque_size func(unsafe.Pointer) uintptr
	_launch_data_get_real func(unsafe.Pointer) float64
	_launch_data_get_string func(unsafe.Pointer) unsafe.Pointer
	_launch_data_get_type func(unsafe.Pointer) unsafe.Pointer
	_launch_data_new_bool func(bool) unsafe.Pointer
	_launch_data_new_fd func(int) unsafe.Pointer
	_launch_data_new_integer func(unsafe.Pointer) unsafe.Pointer
	_launch_data_new_machport func(unsafe.Pointer) unsafe.Pointer
	_launch_data_new_opaque func(unsafe.Pointer, uintptr) unsafe.Pointer
	_launch_data_new_real func(float64) unsafe.Pointer
	_launch_data_new_string func(unsafe.Pointer) unsafe.Pointer
	_launch_data_set_bool func(unsafe.Pointer, bool) bool
	_launch_data_set_fd func(unsafe.Pointer, int) bool
	_launch_data_set_integer func(unsafe.Pointer, unsafe.Pointer) bool
	_launch_data_set_machport func(unsafe.Pointer, unsafe.Pointer) bool
	_launch_data_set_opaque func(unsafe.Pointer, unsafe.Pointer, uintptr) bool
	_launch_data_set_real func(unsafe.Pointer, float64) bool
	_launch_data_set_string func(unsafe.Pointer, unsafe.Pointer) bool
	_launch_get_fd func() int
	_launch_msg func(unsafe.Pointer) unsafe.Pointer
	_xpc_activity_copy_criteria func(unsafe.Pointer) unsafe.Pointer
	_xpc_activity_get_state func(unsafe.Pointer) unsafe.Pointer
	_xpc_activity_register func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_activity_set_criteria func(unsafe.Pointer, unsafe.Pointer)
	_xpc_activity_set_state func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_activity_should_defer func(unsafe.Pointer) bool
	_xpc_activity_unregister func(unsafe.Pointer)
	_xpc_array_append_value func(unsafe.Pointer, unsafe.Pointer)
	_xpc_array_apply func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_array_create func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_create_connection func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_create_empty func() unsafe.Pointer
	_xpc_array_dup_fd func(unsafe.Pointer, uintptr) int
	_xpc_array_get_array func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_bool func(unsafe.Pointer, uintptr) bool
	_xpc_array_get_count func(unsafe.Pointer) uintptr
	_xpc_array_get_data func(unsafe.Pointer, uintptr, unsafe.Pointer) unsafe.Pointer
	_xpc_array_get_date func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_dictionary func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_double func(unsafe.Pointer, uintptr) float64
	_xpc_array_get_int64 func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_string func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_uint64 func(unsafe.Pointer, uintptr) uint64
	_xpc_array_get_uuid func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_get_value func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_array_set_bool func(unsafe.Pointer, uintptr, bool)
	_xpc_array_set_connection func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_array_set_data func(unsafe.Pointer, uintptr, unsafe.Pointer, uintptr)
	_xpc_array_set_date func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_array_set_double func(unsafe.Pointer, uintptr, float64)
	_xpc_array_set_fd func(unsafe.Pointer, uintptr, int)
	_xpc_array_set_int64 func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_array_set_string func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_array_set_uint64 func(unsafe.Pointer, uintptr, uint64)
	_xpc_array_set_uuid func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_array_set_value func(unsafe.Pointer, uintptr, unsafe.Pointer)
	_xpc_bool_create func(bool) unsafe.Pointer
	_xpc_bool_get_value func(unsafe.Pointer) bool
	_xpc_connection_activate func(unsafe.Pointer)
	_xpc_connection_cancel func(unsafe.Pointer)
	_xpc_connection_copy_invalidation_reason func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_connection_create_from_endpoint func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_create_mach_service func(unsafe.Pointer, unsafe.Pointer, uint64) unsafe.Pointer
	_xpc_connection_get_asid func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_get_context func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_get_egid func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_get_euid func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_get_name func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_get_pid func(unsafe.Pointer) unsafe.Pointer
	_xpc_connection_resume func(unsafe.Pointer)
	_xpc_connection_send_barrier func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_send_message func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_send_message_with_reply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_send_message_with_reply_sync func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_connection_set_context func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_set_event_handler func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_set_finalizer_f func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_set_peer_code_signing_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_peer_entitlement_exists_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_peer_entitlement_matches_value_requirement func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_peer_lightweight_code_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_peer_platform_identity_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_peer_requirement func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_set_peer_team_identity_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_connection_set_target_queue func(unsafe.Pointer, unsafe.Pointer)
	_xpc_connection_suspend func(unsafe.Pointer)
	_xpc_copy func(unsafe.Pointer) unsafe.Pointer
	_xpc_copy_description func(unsafe.Pointer) unsafe.Pointer
	_xpc_data_create func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_data_create_with_dispatch_data func(unsafe.Pointer) unsafe.Pointer
	_xpc_data_get_bytes func(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr) uintptr
	_xpc_data_get_bytes_ptr func(unsafe.Pointer) unsafe.Pointer
	_xpc_data_get_length func(unsafe.Pointer) uintptr
	_xpc_date_create func(unsafe.Pointer) unsafe.Pointer
	_xpc_date_create_from_current func() unsafe.Pointer
	_xpc_date_get_value func(unsafe.Pointer) unsafe.Pointer
	_xpc_debugger_api_misuse_info func() unsafe.Pointer
	_xpc_dictionary_apply func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_dictionary_copy_mach_send func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_create func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_dictionary_create_connection func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_create_empty func() unsafe.Pointer
	_xpc_dictionary_create_reply func(unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_dup_fd func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_dictionary_get_array func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_bool func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_dictionary_get_count func(unsafe.Pointer) uintptr
	_xpc_dictionary_get_data func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_date func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_dictionary func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_double func(unsafe.Pointer, unsafe.Pointer) float64
	_xpc_dictionary_get_int64 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_remote_connection func(unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_string func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_uint64 func(unsafe.Pointer, unsafe.Pointer) uint64
	_xpc_dictionary_get_uuid func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_get_value func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_dictionary_set_bool func(unsafe.Pointer, unsafe.Pointer, bool)
	_xpc_dictionary_set_connection func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_data func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, uintptr)
	_xpc_dictionary_set_date func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_double func(unsafe.Pointer, unsafe.Pointer, float64)
	_xpc_dictionary_set_fd func(unsafe.Pointer, unsafe.Pointer, int)
	_xpc_dictionary_set_int64 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_mach_send func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_string func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_uint64 func(unsafe.Pointer, unsafe.Pointer, uint64)
	_xpc_dictionary_set_uuid func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_dictionary_set_value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_double_create func(float64) unsafe.Pointer
	_xpc_double_get_value func(unsafe.Pointer) float64
	_xpc_endpoint_create func(unsafe.Pointer) unsafe.Pointer
	_xpc_equal func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_fd_create func(int) unsafe.Pointer
	_xpc_fd_dup func(unsafe.Pointer) int
	_xpc_get_type func(unsafe.Pointer) unsafe.Pointer
	_xpc_hash func(unsafe.Pointer) uintptr
	_xpc_int64_create func(unsafe.Pointer) unsafe.Pointer
	_xpc_int64_get_value func(unsafe.Pointer) unsafe.Pointer
	_xpc_listener_activate func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_listener_cancel func(unsafe.Pointer)
	_xpc_listener_copy_description func(unsafe.Pointer) unsafe.Pointer
	_xpc_listener_create func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_listener_reject_peer func(unsafe.Pointer, unsafe.Pointer)
	_xpc_listener_set_peer_code_signing_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_listener_set_peer_requirement func(unsafe.Pointer, unsafe.Pointer)
	_xpc_main func(unsafe.Pointer)
	_xpc_null_create func() unsafe.Pointer
	_xpc_peer_requirement_create_entitlement_exists func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_peer_requirement_create_entitlement_matches_value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_peer_requirement_create_lwcr func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_peer_requirement_create_platform_identity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_peer_requirement_create_team_identity func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_peer_requirement_match_received_message func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_xpc_release func(unsafe.Pointer)
	_xpc_retain func(unsafe.Pointer) unsafe.Pointer
	_xpc_rich_error_can_retry func(unsafe.Pointer) bool
	_xpc_rich_error_copy_description func(unsafe.Pointer) unsafe.Pointer
	_xpc_session_activate func(unsafe.Pointer, unsafe.Pointer) bool
	_xpc_session_cancel func(unsafe.Pointer)
	_xpc_session_copy_description func(unsafe.Pointer) unsafe.Pointer
	_xpc_session_create_mach_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_session_create_xpc_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_session_send_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_session_send_message_with_reply_async func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_session_send_message_with_reply_sync func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_session_set_cancel_handler func(unsafe.Pointer, unsafe.Pointer)
	_xpc_session_set_incoming_message_handler func(unsafe.Pointer, unsafe.Pointer)
	_xpc_session_set_peer_code_signing_requirement func(unsafe.Pointer, unsafe.Pointer) int
	_xpc_session_set_peer_requirement func(unsafe.Pointer, unsafe.Pointer)
	_xpc_session_set_target_queue func(unsafe.Pointer, unsafe.Pointer)
	_xpc_set_event_stream_handler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_xpc_shmem_create func(unsafe.Pointer, uintptr) unsafe.Pointer
	_xpc_shmem_map func(unsafe.Pointer, unsafe.Pointer) uintptr
	_xpc_string_create func(unsafe.Pointer) unsafe.Pointer
	_xpc_string_create_with_format func(unsafe.Pointer) unsafe.Pointer
	_xpc_string_create_with_format_and_arguments func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_xpc_string_get_length func(unsafe.Pointer) uintptr
	_xpc_string_get_string_ptr func(unsafe.Pointer) unsafe.Pointer
	_xpc_transaction_begin func()
	_xpc_transaction_end func()
	_xpc_type_get_name func(unsafe.Pointer) unsafe.Pointer
	_xpc_uint64_create func(uint64) unsafe.Pointer
	_xpc_uint64_get_value func(unsafe.Pointer) uint64
	_xpc_uuid_create func(unsafe.Pointer) unsafe.Pointer
	_xpc_uuid_get_bytes func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_launch_activate_socket, lib, "launch_activate_socket")
	tryRegister(&_launch_data_alloc, lib, "launch_data_alloc")
	tryRegister(&_launch_data_array_get_count, lib, "launch_data_array_get_count")
	tryRegister(&_launch_data_array_get_index, lib, "launch_data_array_get_index")
	tryRegister(&_launch_data_array_set_index, lib, "launch_data_array_set_index")
	tryRegister(&_launch_data_copy, lib, "launch_data_copy")
	tryRegister(&_launch_data_dict_get_count, lib, "launch_data_dict_get_count")
	tryRegister(&_launch_data_dict_insert, lib, "launch_data_dict_insert")
	tryRegister(&_launch_data_dict_iterate, lib, "launch_data_dict_iterate")
	tryRegister(&_launch_data_dict_lookup, lib, "launch_data_dict_lookup")
	tryRegister(&_launch_data_dict_remove, lib, "launch_data_dict_remove")
	tryRegister(&_launch_data_free, lib, "launch_data_free")
	tryRegister(&_launch_data_get_bool, lib, "launch_data_get_bool")
	tryRegister(&_launch_data_get_errno, lib, "launch_data_get_errno")
	tryRegister(&_launch_data_get_fd, lib, "launch_data_get_fd")
	tryRegister(&_launch_data_get_integer, lib, "launch_data_get_integer")
	tryRegister(&_launch_data_get_machport, lib, "launch_data_get_machport")
	tryRegister(&_launch_data_get_opaque, lib, "launch_data_get_opaque")
	tryRegister(&_launch_data_get_opaque_size, lib, "launch_data_get_opaque_size")
	tryRegister(&_launch_data_get_real, lib, "launch_data_get_real")
	tryRegister(&_launch_data_get_string, lib, "launch_data_get_string")
	tryRegister(&_launch_data_get_type, lib, "launch_data_get_type")
	tryRegister(&_launch_data_new_bool, lib, "launch_data_new_bool")
	tryRegister(&_launch_data_new_fd, lib, "launch_data_new_fd")
	tryRegister(&_launch_data_new_integer, lib, "launch_data_new_integer")
	tryRegister(&_launch_data_new_machport, lib, "launch_data_new_machport")
	tryRegister(&_launch_data_new_opaque, lib, "launch_data_new_opaque")
	tryRegister(&_launch_data_new_real, lib, "launch_data_new_real")
	tryRegister(&_launch_data_new_string, lib, "launch_data_new_string")
	tryRegister(&_launch_data_set_bool, lib, "launch_data_set_bool")
	tryRegister(&_launch_data_set_fd, lib, "launch_data_set_fd")
	tryRegister(&_launch_data_set_integer, lib, "launch_data_set_integer")
	tryRegister(&_launch_data_set_machport, lib, "launch_data_set_machport")
	tryRegister(&_launch_data_set_opaque, lib, "launch_data_set_opaque")
	tryRegister(&_launch_data_set_real, lib, "launch_data_set_real")
	tryRegister(&_launch_data_set_string, lib, "launch_data_set_string")
	tryRegister(&_launch_get_fd, lib, "launch_get_fd")
	tryRegister(&_launch_msg, lib, "launch_msg")
	tryRegister(&_xpc_activity_copy_criteria, lib, "xpc_activity_copy_criteria")
	tryRegister(&_xpc_activity_get_state, lib, "xpc_activity_get_state")
	tryRegister(&_xpc_activity_register, lib, "xpc_activity_register")
	tryRegister(&_xpc_activity_set_criteria, lib, "xpc_activity_set_criteria")
	tryRegister(&_xpc_activity_set_state, lib, "xpc_activity_set_state")
	tryRegister(&_xpc_activity_should_defer, lib, "xpc_activity_should_defer")
	tryRegister(&_xpc_activity_unregister, lib, "xpc_activity_unregister")
	tryRegister(&_xpc_array_append_value, lib, "xpc_array_append_value")
	tryRegister(&_xpc_array_apply, lib, "xpc_array_apply")
	tryRegister(&_xpc_array_create, lib, "xpc_array_create")
	tryRegister(&_xpc_array_create_connection, lib, "xpc_array_create_connection")
	tryRegister(&_xpc_array_create_empty, lib, "xpc_array_create_empty")
	tryRegister(&_xpc_array_dup_fd, lib, "xpc_array_dup_fd")
	tryRegister(&_xpc_array_get_array, lib, "xpc_array_get_array")
	tryRegister(&_xpc_array_get_bool, lib, "xpc_array_get_bool")
	tryRegister(&_xpc_array_get_count, lib, "xpc_array_get_count")
	tryRegister(&_xpc_array_get_data, lib, "xpc_array_get_data")
	tryRegister(&_xpc_array_get_date, lib, "xpc_array_get_date")
	tryRegister(&_xpc_array_get_dictionary, lib, "xpc_array_get_dictionary")
	tryRegister(&_xpc_array_get_double, lib, "xpc_array_get_double")
	tryRegister(&_xpc_array_get_int64, lib, "xpc_array_get_int64")
	tryRegister(&_xpc_array_get_string, lib, "xpc_array_get_string")
	tryRegister(&_xpc_array_get_uint64, lib, "xpc_array_get_uint64")
	tryRegister(&_xpc_array_get_uuid, lib, "xpc_array_get_uuid")
	tryRegister(&_xpc_array_get_value, lib, "xpc_array_get_value")
	tryRegister(&_xpc_array_set_bool, lib, "xpc_array_set_bool")
	tryRegister(&_xpc_array_set_connection, lib, "xpc_array_set_connection")
	tryRegister(&_xpc_array_set_data, lib, "xpc_array_set_data")
	tryRegister(&_xpc_array_set_date, lib, "xpc_array_set_date")
	tryRegister(&_xpc_array_set_double, lib, "xpc_array_set_double")
	tryRegister(&_xpc_array_set_fd, lib, "xpc_array_set_fd")
	tryRegister(&_xpc_array_set_int64, lib, "xpc_array_set_int64")
	tryRegister(&_xpc_array_set_string, lib, "xpc_array_set_string")
	tryRegister(&_xpc_array_set_uint64, lib, "xpc_array_set_uint64")
	tryRegister(&_xpc_array_set_uuid, lib, "xpc_array_set_uuid")
	tryRegister(&_xpc_array_set_value, lib, "xpc_array_set_value")
	tryRegister(&_xpc_bool_create, lib, "xpc_bool_create")
	tryRegister(&_xpc_bool_get_value, lib, "xpc_bool_get_value")
	tryRegister(&_xpc_connection_activate, lib, "xpc_connection_activate")
	tryRegister(&_xpc_connection_cancel, lib, "xpc_connection_cancel")
	tryRegister(&_xpc_connection_copy_invalidation_reason, lib, "xpc_connection_copy_invalidation_reason")
	tryRegister(&_xpc_connection_create, lib, "xpc_connection_create")
	tryRegister(&_xpc_connection_create_from_endpoint, lib, "xpc_connection_create_from_endpoint")
	tryRegister(&_xpc_connection_create_mach_service, lib, "xpc_connection_create_mach_service")
	tryRegister(&_xpc_connection_get_asid, lib, "xpc_connection_get_asid")
	tryRegister(&_xpc_connection_get_context, lib, "xpc_connection_get_context")
	tryRegister(&_xpc_connection_get_egid, lib, "xpc_connection_get_egid")
	tryRegister(&_xpc_connection_get_euid, lib, "xpc_connection_get_euid")
	tryRegister(&_xpc_connection_get_name, lib, "xpc_connection_get_name")
	tryRegister(&_xpc_connection_get_pid, lib, "xpc_connection_get_pid")
	tryRegister(&_xpc_connection_resume, lib, "xpc_connection_resume")
	tryRegister(&_xpc_connection_send_barrier, lib, "xpc_connection_send_barrier")
	tryRegister(&_xpc_connection_send_message, lib, "xpc_connection_send_message")
	tryRegister(&_xpc_connection_send_message_with_reply, lib, "xpc_connection_send_message_with_reply")
	tryRegister(&_xpc_connection_send_message_with_reply_sync, lib, "xpc_connection_send_message_with_reply_sync")
	tryRegister(&_xpc_connection_set_context, lib, "xpc_connection_set_context")
	tryRegister(&_xpc_connection_set_event_handler, lib, "xpc_connection_set_event_handler")
	tryRegister(&_xpc_connection_set_finalizer_f, lib, "xpc_connection_set_finalizer_f")
	tryRegister(&_xpc_connection_set_peer_code_signing_requirement, lib, "xpc_connection_set_peer_code_signing_requirement")
	tryRegister(&_xpc_connection_set_peer_entitlement_exists_requirement, lib, "xpc_connection_set_peer_entitlement_exists_requirement")
	tryRegister(&_xpc_connection_set_peer_entitlement_matches_value_requirement, lib, "xpc_connection_set_peer_entitlement_matches_value_requirement")
	tryRegister(&_xpc_connection_set_peer_lightweight_code_requirement, lib, "xpc_connection_set_peer_lightweight_code_requirement")
	tryRegister(&_xpc_connection_set_peer_platform_identity_requirement, lib, "xpc_connection_set_peer_platform_identity_requirement")
	tryRegister(&_xpc_connection_set_peer_requirement, lib, "xpc_connection_set_peer_requirement")
	tryRegister(&_xpc_connection_set_peer_team_identity_requirement, lib, "xpc_connection_set_peer_team_identity_requirement")
	tryRegister(&_xpc_connection_set_target_queue, lib, "xpc_connection_set_target_queue")
	tryRegister(&_xpc_connection_suspend, lib, "xpc_connection_suspend")
	tryRegister(&_xpc_copy, lib, "xpc_copy")
	tryRegister(&_xpc_copy_description, lib, "xpc_copy_description")
	tryRegister(&_xpc_data_create, lib, "xpc_data_create")
	tryRegister(&_xpc_data_create_with_dispatch_data, lib, "xpc_data_create_with_dispatch_data")
	tryRegister(&_xpc_data_get_bytes, lib, "xpc_data_get_bytes")
	tryRegister(&_xpc_data_get_bytes_ptr, lib, "xpc_data_get_bytes_ptr")
	tryRegister(&_xpc_data_get_length, lib, "xpc_data_get_length")
	tryRegister(&_xpc_date_create, lib, "xpc_date_create")
	tryRegister(&_xpc_date_create_from_current, lib, "xpc_date_create_from_current")
	tryRegister(&_xpc_date_get_value, lib, "xpc_date_get_value")
	tryRegister(&_xpc_debugger_api_misuse_info, lib, "xpc_debugger_api_misuse_info")
	tryRegister(&_xpc_dictionary_apply, lib, "xpc_dictionary_apply")
	tryRegister(&_xpc_dictionary_copy_mach_send, lib, "xpc_dictionary_copy_mach_send")
	tryRegister(&_xpc_dictionary_create, lib, "xpc_dictionary_create")
	tryRegister(&_xpc_dictionary_create_connection, lib, "xpc_dictionary_create_connection")
	tryRegister(&_xpc_dictionary_create_empty, lib, "xpc_dictionary_create_empty")
	tryRegister(&_xpc_dictionary_create_reply, lib, "xpc_dictionary_create_reply")
	tryRegister(&_xpc_dictionary_dup_fd, lib, "xpc_dictionary_dup_fd")
	tryRegister(&_xpc_dictionary_get_array, lib, "xpc_dictionary_get_array")
	tryRegister(&_xpc_dictionary_get_bool, lib, "xpc_dictionary_get_bool")
	tryRegister(&_xpc_dictionary_get_count, lib, "xpc_dictionary_get_count")
	tryRegister(&_xpc_dictionary_get_data, lib, "xpc_dictionary_get_data")
	tryRegister(&_xpc_dictionary_get_date, lib, "xpc_dictionary_get_date")
	tryRegister(&_xpc_dictionary_get_dictionary, lib, "xpc_dictionary_get_dictionary")
	tryRegister(&_xpc_dictionary_get_double, lib, "xpc_dictionary_get_double")
	tryRegister(&_xpc_dictionary_get_int64, lib, "xpc_dictionary_get_int64")
	tryRegister(&_xpc_dictionary_get_remote_connection, lib, "xpc_dictionary_get_remote_connection")
	tryRegister(&_xpc_dictionary_get_string, lib, "xpc_dictionary_get_string")
	tryRegister(&_xpc_dictionary_get_uint64, lib, "xpc_dictionary_get_uint64")
	tryRegister(&_xpc_dictionary_get_uuid, lib, "xpc_dictionary_get_uuid")
	tryRegister(&_xpc_dictionary_get_value, lib, "xpc_dictionary_get_value")
	tryRegister(&_xpc_dictionary_set_bool, lib, "xpc_dictionary_set_bool")
	tryRegister(&_xpc_dictionary_set_connection, lib, "xpc_dictionary_set_connection")
	tryRegister(&_xpc_dictionary_set_data, lib, "xpc_dictionary_set_data")
	tryRegister(&_xpc_dictionary_set_date, lib, "xpc_dictionary_set_date")
	tryRegister(&_xpc_dictionary_set_double, lib, "xpc_dictionary_set_double")
	tryRegister(&_xpc_dictionary_set_fd, lib, "xpc_dictionary_set_fd")
	tryRegister(&_xpc_dictionary_set_int64, lib, "xpc_dictionary_set_int64")
	tryRegister(&_xpc_dictionary_set_mach_send, lib, "xpc_dictionary_set_mach_send")
	tryRegister(&_xpc_dictionary_set_string, lib, "xpc_dictionary_set_string")
	tryRegister(&_xpc_dictionary_set_uint64, lib, "xpc_dictionary_set_uint64")
	tryRegister(&_xpc_dictionary_set_uuid, lib, "xpc_dictionary_set_uuid")
	tryRegister(&_xpc_dictionary_set_value, lib, "xpc_dictionary_set_value")
	tryRegister(&_xpc_double_create, lib, "xpc_double_create")
	tryRegister(&_xpc_double_get_value, lib, "xpc_double_get_value")
	tryRegister(&_xpc_endpoint_create, lib, "xpc_endpoint_create")
	tryRegister(&_xpc_equal, lib, "xpc_equal")
	tryRegister(&_xpc_fd_create, lib, "xpc_fd_create")
	tryRegister(&_xpc_fd_dup, lib, "xpc_fd_dup")
	tryRegister(&_xpc_get_type, lib, "xpc_get_type")
	tryRegister(&_xpc_hash, lib, "xpc_hash")
	tryRegister(&_xpc_int64_create, lib, "xpc_int64_create")
	tryRegister(&_xpc_int64_get_value, lib, "xpc_int64_get_value")
	tryRegister(&_xpc_listener_activate, lib, "xpc_listener_activate")
	tryRegister(&_xpc_listener_cancel, lib, "xpc_listener_cancel")
	tryRegister(&_xpc_listener_copy_description, lib, "xpc_listener_copy_description")
	tryRegister(&_xpc_listener_create, lib, "xpc_listener_create")
	tryRegister(&_xpc_listener_reject_peer, lib, "xpc_listener_reject_peer")
	tryRegister(&_xpc_listener_set_peer_code_signing_requirement, lib, "xpc_listener_set_peer_code_signing_requirement")
	tryRegister(&_xpc_listener_set_peer_requirement, lib, "xpc_listener_set_peer_requirement")
	tryRegister(&_xpc_main, lib, "xpc_main")
	tryRegister(&_xpc_null_create, lib, "xpc_null_create")
	tryRegister(&_xpc_peer_requirement_create_entitlement_exists, lib, "xpc_peer_requirement_create_entitlement_exists")
	tryRegister(&_xpc_peer_requirement_create_entitlement_matches_value, lib, "xpc_peer_requirement_create_entitlement_matches_value")
	tryRegister(&_xpc_peer_requirement_create_lwcr, lib, "xpc_peer_requirement_create_lwcr")
	tryRegister(&_xpc_peer_requirement_create_platform_identity, lib, "xpc_peer_requirement_create_platform_identity")
	tryRegister(&_xpc_peer_requirement_create_team_identity, lib, "xpc_peer_requirement_create_team_identity")
	tryRegister(&_xpc_peer_requirement_match_received_message, lib, "xpc_peer_requirement_match_received_message")
	tryRegister(&_xpc_release, lib, "xpc_release")
	tryRegister(&_xpc_retain, lib, "xpc_retain")
	tryRegister(&_xpc_rich_error_can_retry, lib, "xpc_rich_error_can_retry")
	tryRegister(&_xpc_rich_error_copy_description, lib, "xpc_rich_error_copy_description")
	tryRegister(&_xpc_session_activate, lib, "xpc_session_activate")
	tryRegister(&_xpc_session_cancel, lib, "xpc_session_cancel")
	tryRegister(&_xpc_session_copy_description, lib, "xpc_session_copy_description")
	tryRegister(&_xpc_session_create_mach_service, lib, "xpc_session_create_mach_service")
	tryRegister(&_xpc_session_create_xpc_service, lib, "xpc_session_create_xpc_service")
	tryRegister(&_xpc_session_send_message, lib, "xpc_session_send_message")
	tryRegister(&_xpc_session_send_message_with_reply_async, lib, "xpc_session_send_message_with_reply_async")
	tryRegister(&_xpc_session_send_message_with_reply_sync, lib, "xpc_session_send_message_with_reply_sync")
	tryRegister(&_xpc_session_set_cancel_handler, lib, "xpc_session_set_cancel_handler")
	tryRegister(&_xpc_session_set_incoming_message_handler, lib, "xpc_session_set_incoming_message_handler")
	tryRegister(&_xpc_session_set_peer_code_signing_requirement, lib, "xpc_session_set_peer_code_signing_requirement")
	tryRegister(&_xpc_session_set_peer_requirement, lib, "xpc_session_set_peer_requirement")
	tryRegister(&_xpc_session_set_target_queue, lib, "xpc_session_set_target_queue")
	tryRegister(&_xpc_set_event_stream_handler, lib, "xpc_set_event_stream_handler")
	tryRegister(&_xpc_shmem_create, lib, "xpc_shmem_create")
	tryRegister(&_xpc_shmem_map, lib, "xpc_shmem_map")
	tryRegister(&_xpc_string_create, lib, "xpc_string_create")
	tryRegister(&_xpc_string_create_with_format, lib, "xpc_string_create_with_format")
	tryRegister(&_xpc_string_create_with_format_and_arguments, lib, "xpc_string_create_with_format_and_arguments")
	tryRegister(&_xpc_string_get_length, lib, "xpc_string_get_length")
	tryRegister(&_xpc_string_get_string_ptr, lib, "xpc_string_get_string_ptr")
	tryRegister(&_xpc_transaction_begin, lib, "xpc_transaction_begin")
	tryRegister(&_xpc_transaction_end, lib, "xpc_transaction_end")
	tryRegister(&_xpc_type_get_name, lib, "xpc_type_get_name")
	tryRegister(&_xpc_uint64_create, lib, "xpc_uint64_create")
	tryRegister(&_xpc_uint64_get_value, lib, "xpc_uint64_get_value")
	tryRegister(&_xpc_uuid_create, lib, "xpc_uuid_create")
	tryRegister(&_xpc_uuid_get_bytes, lib, "xpc_uuid_get_bytes")
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


// Retrieves the file descriptors for sockets in the process’s   property list. [Full Topic]
//
// Added in macOS 10.10.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_activate_socket
func launch_activate_socket(name unsafe.Pointer, fds unsafe.Pointer, cnt unsafe.Pointer) int {
	return _launch_activate_socket(name, fds, cnt)
	}


// launch_data_alloc is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_alloc
func launch_data_alloc(type_ unsafe.Pointer) unsafe.Pointer {
	return _launch_data_alloc(type_)
	}


// launch_data_array_get_count is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_array_get_count
func launch_data_array_get_count(larray unsafe.Pointer) uintptr {
	return _launch_data_array_get_count(larray)
	}


// launch_data_array_get_index is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_array_get_index
func launch_data_array_get_index(larray unsafe.Pointer, idx uintptr) unsafe.Pointer {
	return _launch_data_array_get_index(larray, idx)
	}


// launch_data_array_set_index is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_array_set_index
func launch_data_array_set_index(larray unsafe.Pointer, lval unsafe.Pointer, idx uintptr) bool {
	return _launch_data_array_set_index(larray, lval, idx)
	}


// launch_data_copy is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_copy
func launch_data_copy(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_copy(ld)
	}


// launch_data_dict_get_count is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_dict_get_count
func launch_data_dict_get_count(ldict unsafe.Pointer) uintptr {
	return _launch_data_dict_get_count(ldict)
	}


// launch_data_dict_insert is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_dict_insert
func launch_data_dict_insert(ldict unsafe.Pointer, lval unsafe.Pointer, key unsafe.Pointer) bool {
	return _launch_data_dict_insert(ldict, lval, key)
	}


// launch_data_dict_iterate is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_dict_iterate
func launch_data_dict_iterate(ldict unsafe.Pointer, iterator unsafe.Pointer, ctx unsafe.Pointer) {
	_launch_data_dict_iterate(ldict, iterator, ctx)
	}


// launch_data_dict_lookup is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_dict_lookup
func launch_data_dict_lookup(ldict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _launch_data_dict_lookup(ldict, key)
	}


// launch_data_dict_remove is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_dict_remove
func launch_data_dict_remove(ldict unsafe.Pointer, key unsafe.Pointer) bool {
	return _launch_data_dict_remove(ldict, key)
	}


// launch_data_free is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_free
func launch_data_free(ld unsafe.Pointer) {
	_launch_data_free(ld)
	}


// launch_data_get_bool is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_bool
func launch_data_get_bool(ld unsafe.Pointer) bool {
	return _launch_data_get_bool(ld)
	}


// launch_data_get_errno is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_errno
func launch_data_get_errno(ld unsafe.Pointer) int {
	return _launch_data_get_errno(ld)
	}


// launch_data_get_fd is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_fd
func launch_data_get_fd(ld unsafe.Pointer) int {
	return _launch_data_get_fd(ld)
	}


// launch_data_get_integer is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_integer
func launch_data_get_integer(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_get_integer(ld)
	}


// launch_data_get_machport is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_machport
func launch_data_get_machport(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_get_machport(ld)
	}


// launch_data_get_opaque is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_opaque
func launch_data_get_opaque(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_get_opaque(ld)
	}


// launch_data_get_opaque_size is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_opaque_size
func launch_data_get_opaque_size(ld unsafe.Pointer) uintptr {
	return _launch_data_get_opaque_size(ld)
	}


// launch_data_get_real is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_real
func launch_data_get_real(ld unsafe.Pointer) float64 {
	return _launch_data_get_real(ld)
	}


// launch_data_get_string is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_string
func launch_data_get_string(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_get_string(ld)
	}


// launch_data_get_type is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_get_type
func launch_data_get_type(ld unsafe.Pointer) unsafe.Pointer {
	return _launch_data_get_type(ld)
	}


// launch_data_new_bool is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_bool
func launch_data_new_bool(val bool) unsafe.Pointer {
	return _launch_data_new_bool(val)
	}


// launch_data_new_fd is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_fd
func launch_data_new_fd(fd int) unsafe.Pointer {
	return _launch_data_new_fd(fd)
	}


// launch_data_new_integer is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_integer
func launch_data_new_integer(val unsafe.Pointer) unsafe.Pointer {
	return _launch_data_new_integer(val)
	}


// launch_data_new_machport is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_machport
func launch_data_new_machport(val unsafe.Pointer) unsafe.Pointer {
	return _launch_data_new_machport(val)
	}


// launch_data_new_opaque is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_opaque
func launch_data_new_opaque(bytes unsafe.Pointer, sz uintptr) unsafe.Pointer {
	return _launch_data_new_opaque(bytes, sz)
	}


// launch_data_new_real is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_real
func launch_data_new_real(val float64) unsafe.Pointer {
	return _launch_data_new_real(val)
	}


// launch_data_new_string is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_new_string
func launch_data_new_string(val unsafe.Pointer) unsafe.Pointer {
	return _launch_data_new_string(val)
	}


// launch_data_set_bool is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_bool
func launch_data_set_bool(ld unsafe.Pointer, val bool) bool {
	return _launch_data_set_bool(ld, val)
	}


// launch_data_set_fd is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_fd
func launch_data_set_fd(ld unsafe.Pointer, fd int) bool {
	return _launch_data_set_fd(ld, fd)
	}


// launch_data_set_integer is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_integer
func launch_data_set_integer(ld unsafe.Pointer, val unsafe.Pointer) bool {
	return _launch_data_set_integer(ld, val)
	}


// launch_data_set_machport is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_machport
func launch_data_set_machport(ld unsafe.Pointer, mp unsafe.Pointer) bool {
	return _launch_data_set_machport(ld, mp)
	}


// launch_data_set_opaque is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_opaque
func launch_data_set_opaque(ld unsafe.Pointer, bytes unsafe.Pointer, sz uintptr) bool {
	return _launch_data_set_opaque(ld, bytes, sz)
	}


// launch_data_set_real is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_real
func launch_data_set_real(ld unsafe.Pointer, val float64) bool {
	return _launch_data_set_real(ld, val)
	}


// launch_data_set_string is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_data_set_string
func launch_data_set_string(ld unsafe.Pointer, val unsafe.Pointer) bool {
	return _launch_data_set_string(ld, val)
	}


// launch_get_fd is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_get_fd
func launch_get_fd() int {
	return _launch_get_fd()
	}


// launch_msg is a XPC function. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/launch_msg
func launch_msg(request unsafe.Pointer) unsafe.Pointer {
	return _launch_msg(request)
	}


// Returns an XPC dictionary that describes the execution criteria of an activity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_copy_criteria(_:)
func xpc_activity_copy_criteria(activity unsafe.Pointer) unsafe.Pointer {
	return _xpc_activity_copy_criteria(activity)
	}


// Returns the current state of an activity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_get_state(_:)
func xpc_activity_get_state(activity unsafe.Pointer) unsafe.Pointer {
	return _xpc_activity_get_state(activity)
	}


// Registers an activity with the system. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_register(_:_:_:)
func xpc_activity_register(identifier unsafe.Pointer, criteria unsafe.Pointer, handler unsafe.Pointer) {
	_xpc_activity_register(identifier, criteria, handler)
	}


// Modifies the execution criteria of an activity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_set_criteria(_:_:)
func xpc_activity_set_criteria(activity unsafe.Pointer, criteria unsafe.Pointer) {
	_xpc_activity_set_criteria(activity, criteria)
	}


// Updates the current state of an activity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_set_state(_:_:)
func xpc_activity_set_state(activity unsafe.Pointer, state unsafe.Pointer) bool {
	return _xpc_activity_set_state(activity, state)
	}


// Tests whether to defer an activity. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_should_defer(_:)
func xpc_activity_should_defer(activity unsafe.Pointer) bool {
	return _xpc_activity_should_defer(activity)
	}


// Unregisters an activity with the specified identifier. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_activity_unregister(_:)
func xpc_activity_unregister(identifier unsafe.Pointer) {
	_xpc_activity_unregister(identifier)
	}


// Appends an object to an XPC array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_append_value(_:_:)
func xpc_array_append_value(xarray unsafe.Pointer, value unsafe.Pointer) {
	_xpc_array_append_value(xarray, value)
	}


// Invokes the specified block for every value in the array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_apply(_:_:)
func xpc_array_apply(xarray unsafe.Pointer, applier unsafe.Pointer) bool {
	return _xpc_array_apply(xarray, applier)
	}


// Creates an XPC object that represents an array of XPC objects. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_create(_:_:)
func xpc_array_create(objects unsafe.Pointer, count uintptr) unsafe.Pointer {
	return _xpc_array_create(objects, count)
	}


// Creates a connection object from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_create_connection(_:_:)
func xpc_array_create_connection(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_create_connection(xarray, index)
	}


// Creates an XPC object that represents an array of XPC objects. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_create_empty()
func xpc_array_create_empty() unsafe.Pointer {
	return _xpc_array_create_empty()
	}


// Gets a file descriptor from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_dup_fd(_:_:)
func xpc_array_dup_fd(xarray unsafe.Pointer, index uintptr) int {
	return _xpc_array_dup_fd(xarray, index)
	}


// Returns the array at the specified index in the array. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_array(_:_:)
func xpc_array_get_array(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_array(xarray, index)
	}


// Gets a Boolean primitive value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_bool(_:_:)
func xpc_array_get_bool(xarray unsafe.Pointer, index uintptr) bool {
	return _xpc_array_get_bool(xarray, index)
	}


// Returns the count of values in the array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_count(_:)
func xpc_array_get_count(xarray unsafe.Pointer) uintptr {
	return _xpc_array_get_count(xarray)
	}


// Gets a pointer to the raw bytes of a data object from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_data(_:_:_:)
func xpc_array_get_data(xarray unsafe.Pointer, index uintptr, length unsafe.Pointer) unsafe.Pointer {
	return _xpc_array_get_data(xarray, index, length)
	}


// Gets a date interval from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_date(_:_:)
func xpc_array_get_date(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_date(xarray, index)
	}


// Returns the dictionary at the specified index in the array. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_dictionary(_:_:)
func xpc_array_get_dictionary(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_dictionary(xarray, index)
	}


// Gets a double-precision floating point primitive value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_double(_:_:)
func xpc_array_get_double(xarray unsafe.Pointer, index uintptr) float64 {
	return _xpc_array_get_double(xarray, index)
	}


// Gets a 64-bit integer primitive value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_int64(_:_:)
func xpc_array_get_int64(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_int64(xarray, index)
	}


// Gets a C-string value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_string(_:_:)
func xpc_array_get_string(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_string(xarray, index)
	}


// Gets a 64-bit unsigned integer primitive value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_uint64(_:_:)
func xpc_array_get_uint64(xarray unsafe.Pointer, index uintptr) uint64 {
	return _xpc_array_get_uint64(xarray, index)
	}


// Gets a UUID value from an array directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_uuid(_:_:)
func xpc_array_get_uuid(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_uuid(xarray, index)
	}


// Returns the value at the specified index in the array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_get_value(_:_:)
func xpc_array_get_value(xarray unsafe.Pointer, index uintptr) unsafe.Pointer {
	return _xpc_array_get_value(xarray, index)
	}


// Inserts a Boolean primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_bool(_:_:_:)
func xpc_array_set_bool(xarray unsafe.Pointer, index uintptr, value bool) {
	_xpc_array_set_bool(xarray, index, value)
	}


// Inserts a connection into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_connection(_:_:_:)
func xpc_array_set_connection(xarray unsafe.Pointer, index uintptr, connection unsafe.Pointer) {
	_xpc_array_set_connection(xarray, index, connection)
	}


// Inserts a raw data value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_data(_:_:_:_:)
func xpc_array_set_data(xarray unsafe.Pointer, index uintptr, bytes unsafe.Pointer, length uintptr) {
	_xpc_array_set_data(xarray, index, bytes, length)
	}


// Inserts a date value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_date(_:_:_:)
func xpc_array_set_date(xarray unsafe.Pointer, index uintptr, value unsafe.Pointer) {
	_xpc_array_set_date(xarray, index, value)
	}


// Inserts a double-precision floating point primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_double(_:_:_:)
func xpc_array_set_double(xarray unsafe.Pointer, index uintptr, value float64) {
	_xpc_array_set_double(xarray, index, value)
	}


// Inserts a file descriptor into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_fd(_:_:_:)
func xpc_array_set_fd(xarray unsafe.Pointer, index uintptr, fd int) {
	_xpc_array_set_fd(xarray, index, fd)
	}


// Inserts a 64-bit integer primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_int64(_:_:_:)
func xpc_array_set_int64(xarray unsafe.Pointer, index uintptr, value unsafe.Pointer) {
	_xpc_array_set_int64(xarray, index, value)
	}


// Inserts a C-string into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_string(_:_:_:)
func xpc_array_set_string(xarray unsafe.Pointer, index uintptr, string unsafe.Pointer) {
	_xpc_array_set_string(xarray, index, string)
	}


// Inserts a 64-bit unsigned integer primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_uint64(_:_:_:)
func xpc_array_set_uint64(xarray unsafe.Pointer, index uintptr, value uint64) {
	_xpc_array_set_uint64(xarray, index, value)
	}


// Inserts a UUID primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_uuid(_:_:_:)
func xpc_array_set_uuid(xarray unsafe.Pointer, index uintptr, uuid unsafe.Pointer) {
	_xpc_array_set_uuid(xarray, index, uuid)
	}


// Inserts the specified object into the array at the specified index. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_array_set_value(_:_:_:)
func xpc_array_set_value(xarray unsafe.Pointer, index uintptr, value unsafe.Pointer) {
	_xpc_array_set_value(xarray, index, value)
	}


// Creates an XPC Boolean object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_bool_create(_:)
func xpc_bool_create(value bool) unsafe.Pointer {
	return _xpc_bool_create(value)
	}


// Returns the underlying Boolean value from the object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_bool_get_value(_:)
func xpc_bool_get_value(xbool unsafe.Pointer) bool {
	return _xpc_bool_get_value(xbool)
	}


// Activates a new connection. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_activate(_:)
func xpc_connection_activate(connection unsafe.Pointer) {
	_xpc_connection_activate(connection)
	}


// Cancels the connection and ensures that its event handler doesn’t fire again. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_cancel(_:)
func xpc_connection_cancel(connection unsafe.Pointer) {
	_xpc_connection_cancel(connection)
	}


// xpc_connection_copy_invalidation_reason is a XPC function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_copy_invalidation_reason(_:)
func xpc_connection_copy_invalidation_reason(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_copy_invalidation_reason(connection)
	}


// Creates a new connection object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_create(_:_:)
func xpc_connection_create(name unsafe.Pointer, targetq unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_create(name, targetq)
	}


// Creates a new connection from the specified endpoint. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_create_from_endpoint(_:)
func xpc_connection_create_from_endpoint(endpoint unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_create_from_endpoint(endpoint)
	}


// Creates a new connection object that represents a Mach service. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_create_mach_service(_:_:_:)
func xpc_connection_create_mach_service(name unsafe.Pointer, targetq unsafe.Pointer, flags uint64) unsafe.Pointer {
	return _xpc_connection_create_mach_service(name, targetq, flags)
	}


// Returns the audit session identifier of the remote peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_asid(_:)
func xpc_connection_get_asid(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_asid(connection)
	}


// Returns the context for the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_context(_:)
func xpc_connection_get_context(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_context(connection)
	}


// Returns the EGID of the remote peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_egid(_:)
func xpc_connection_get_egid(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_egid(connection)
	}


// Returns the EUID of the remote peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_euid(_:)
func xpc_connection_get_euid(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_euid(connection)
	}


// Returns the name of the remote service that creates the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_name(_:)
func xpc_connection_get_name(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_name(connection)
	}


// Returns the PID of the remote peer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_get_pid(_:)
func xpc_connection_get_pid(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_get_pid(connection)
	}


// Resumes a suspended connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_resume(_:)
func xpc_connection_resume(connection unsafe.Pointer) {
	_xpc_connection_resume(connection)
	}


// Issues a barrier against the connection’s message-send activity. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_send_barrier(_:_:)
func xpc_connection_send_barrier(connection unsafe.Pointer, barrier unsafe.Pointer) {
	_xpc_connection_send_barrier(connection, barrier)
	}


// Sends a message over the connection to the destination service. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_send_message(_:_:)
func xpc_connection_send_message(connection unsafe.Pointer, message unsafe.Pointer) {
	_xpc_connection_send_message(connection, message)
	}


// Sends a message over the connection to the destination service and associates a handler to invoke when the remote service sends a reply message. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_send_message_with_reply(_:_:_:_:)
func xpc_connection_send_message_with_reply(connection unsafe.Pointer, message unsafe.Pointer, replyq unsafe.Pointer, handler unsafe.Pointer) {
	_xpc_connection_send_message_with_reply(connection, message, replyq, handler)
	}


// Sends a message over the connection and blocks the caller until it receives a reply. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_send_message_with_reply_sync(_:_:)
func xpc_connection_send_message_with_reply_sync(connection unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer {
	return _xpc_connection_send_message_with_reply_sync(connection, message)
	}


// Sets a context on the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_context(_:_:)
func xpc_connection_set_context(connection unsafe.Pointer, context unsafe.Pointer) {
	_xpc_connection_set_context(connection, context)
	}


// Sets the event handler block for the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_event_handler(_:_:)
func xpc_connection_set_event_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	_xpc_connection_set_event_handler(connection, handler)
	}


// Sets the finalizer for the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_finalizer_f(_:_:)
func xpc_connection_set_finalizer_f(connection unsafe.Pointer, finalizer unsafe.Pointer) {
	_xpc_connection_set_finalizer_f(connection, finalizer)
	}


// xpc_connection_set_peer_code_signing_requirement is a XPC function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_code_signing_requirement(_:_:)
func xpc_connection_set_peer_code_signing_requirement(connection unsafe.Pointer, requirement unsafe.Pointer) int {
	return _xpc_connection_set_peer_code_signing_requirement(connection, requirement)
	}


// Sets a requirement that the executable for the peer process has a valid code signature that contains an entitlement. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_entitlement_exists_requirement(_:_:)
func xpc_connection_set_peer_entitlement_exists_requirement(connection unsafe.Pointer, entitlement unsafe.Pointer) int {
	return _xpc_connection_set_peer_entitlement_exists_requirement(connection, entitlement)
	}


// Sets a requirement that the executable for the peer process has a valid code signature that contains an entitlement with a specific value. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_entitlement_matches_value_requirement(_:_:_:)
func xpc_connection_set_peer_entitlement_matches_value_requirement(connection unsafe.Pointer, entitlement unsafe.Pointer, value unsafe.Pointer) int {
	return _xpc_connection_set_peer_entitlement_matches_value_requirement(connection, entitlement, value)
	}


// Sets a requirement that the executable for the peer process has a valid code signature that matches the lightweight code requirement. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_lightweight_code_requirement(_:_:)
func xpc_connection_set_peer_lightweight_code_requirement(connection unsafe.Pointer, lwcr unsafe.Pointer) int {
	return _xpc_connection_set_peer_lightweight_code_requirement(connection, lwcr)
	}


// Sets a requirement that the executable for the peer process has a valid code signature that identifies it as an Apple-signed binary with the given signing identifier. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_platform_identity_requirement(_:_:)
func xpc_connection_set_peer_platform_identity_requirement(connection unsafe.Pointer, signing_identifier unsafe.Pointer) int {
	return _xpc_connection_set_peer_platform_identity_requirement(connection, signing_identifier)
	}


// xpc_connection_set_peer_requirement is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_requirement
func xpc_connection_set_peer_requirement(connection unsafe.Pointer, peer_requirement unsafe.Pointer) {
	_xpc_connection_set_peer_requirement(connection, peer_requirement)
	}


// Sets a requirement that the executable for the peer process has a valid code signature and is signed by the same team identifier as the calling process. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_peer_team_identity_requirement(_:_:)
func xpc_connection_set_peer_team_identity_requirement(connection unsafe.Pointer, signing_identifier unsafe.Pointer) int {
	return _xpc_connection_set_peer_team_identity_requirement(connection, signing_identifier)
	}


// Sets the target queue of the connection. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_set_target_queue(_:_:)
func xpc_connection_set_target_queue(connection unsafe.Pointer, targetq unsafe.Pointer) {
	_xpc_connection_set_target_queue(connection, targetq)
	}


// Suspends the connection so the event handler block doesn’t fire and the connection doesn’t attempt to send any messages it has in its queue. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_connection_suspend(_:)
func xpc_connection_suspend(connection unsafe.Pointer) {
	_xpc_connection_suspend(connection)
	}


// Creates a copy of the object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_copy(_:)
func xpc_copy(object unsafe.Pointer) unsafe.Pointer {
	return _xpc_copy(object)
	}


// Copies a debug string that describes the object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_copy_description(_:)
func xpc_copy_description(object unsafe.Pointer) unsafe.Pointer {
	return _xpc_copy_description(object)
	}


// Creates an XPC object that represents a buffer of bytes. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_data_create(_:_:)
func xpc_data_create(bytes unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _xpc_data_create(bytes, length)
	}


// Creates an XPC object that represents a buffer of bytes that the specified GCD data object describes. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_data_create_with_dispatch_data(_:)
func xpc_data_create_with_dispatch_data(ddata unsafe.Pointer) unsafe.Pointer {
	return _xpc_data_create_with_dispatch_data(ddata)
	}


// Copies the bytes in a data object into the specified buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_data_get_bytes(_:_:_:_:)
func xpc_data_get_bytes(xdata unsafe.Pointer, buffer unsafe.Pointer, off uintptr, length uintptr) uintptr {
	return _xpc_data_get_bytes(xdata, buffer, off, length)
	}


// Returns a pointer to the internal storage of a data object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_data_get_bytes_ptr(_:)
func xpc_data_get_bytes_ptr(xdata unsafe.Pointer) unsafe.Pointer {
	return _xpc_data_get_bytes_ptr(xdata)
	}


// Returns the length of the data that an XPC data object encapsulates. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_data_get_length(_:)
func xpc_data_get_length(xdata unsafe.Pointer) uintptr {
	return _xpc_data_get_length(xdata)
	}


// Creates an XPC date object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_date_create(_:)
func xpc_date_create(interval unsafe.Pointer) unsafe.Pointer {
	return _xpc_date_create(interval)
	}


// Creates an XPC date object that represents the current date. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_date_create_from_current()
func xpc_date_create_from_current() unsafe.Pointer {
	return _xpc_date_create_from_current()
	}


// Returns the underlying date interval from an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_date_get_value(_:)
func xpc_date_get_value(xdate unsafe.Pointer) unsafe.Pointer {
	return _xpc_date_get_value(xdate)
	}


// Returns a pointer to a string that describes the reason XPC aborts the calling process. [Full Topic]
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_debugger_api_misuse_info()
func xpc_debugger_api_misuse_info() unsafe.Pointer {
	return _xpc_debugger_api_misuse_info()
	}


// Invokes the specified block for every key-value pair in the dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_apply(_:_:)
func xpc_dictionary_apply(xdict unsafe.Pointer, applier unsafe.Pointer) bool {
	return _xpc_dictionary_apply(xdict, applier)
	}


// xpc_dictionary_copy_mach_send is a XPC function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_copy_mach_send(_:_:)
func xpc_dictionary_copy_mach_send(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_copy_mach_send(xdict, key)
	}


// Creates an XPC object that represents a dictionary of XPC objects keyed to C-strings. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_create(_:_:_:)
func xpc_dictionary_create(keys unsafe.Pointer, values unsafe.Pointer, count uintptr) unsafe.Pointer {
	return _xpc_dictionary_create(keys, values, count)
	}


// Creates a connection from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_create_connection(_:_:)
func xpc_dictionary_create_connection(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_create_connection(xdict, key)
	}


// Creates an XPC object that represents a dictionary of XPC objects keyed to C-strings. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_create_empty()
func xpc_dictionary_create_empty() unsafe.Pointer {
	return _xpc_dictionary_create_empty()
	}


// Creates a dictionary that is in reply to the specified dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_create_reply(_:)
func xpc_dictionary_create_reply(original unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_create_reply(original)
	}


// Creates a file descriptor from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_dup_fd(_:_:)
func xpc_dictionary_dup_fd(xdict unsafe.Pointer, key unsafe.Pointer) int {
	return _xpc_dictionary_dup_fd(xdict, key)
	}


// Returns the array value for the specified key. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_array(_:_:)
func xpc_dictionary_get_array(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_array(xdict, key)
	}


// Gets a Boolean primitive value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_bool(_:_:)
func xpc_dictionary_get_bool(xdict unsafe.Pointer, key unsafe.Pointer) bool {
	return _xpc_dictionary_get_bool(xdict, key)
	}


// Returns the number of values in the dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_count(_:)
func xpc_dictionary_get_count(xdict unsafe.Pointer) uintptr {
	return _xpc_dictionary_get_count(xdict)
	}


// Gets a raw data value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_data(_:_:_:)
func xpc_dictionary_get_data(xdict unsafe.Pointer, key unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_data(xdict, key, length)
	}


// Gets a date value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_date(_:_:)
func xpc_dictionary_get_date(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_date(xdict, key)
	}


// Returns the dictionary value for the specified key. [Full Topic]
//
// Added in macOS 10.11.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_dictionary(_:_:)
func xpc_dictionary_get_dictionary(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_dictionary(xdict, key)
	}


// Gets a double-precision floating point primitive value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_double(_:_:)
func xpc_dictionary_get_double(xdict unsafe.Pointer, key unsafe.Pointer) float64 {
	return _xpc_dictionary_get_double(xdict, key)
	}


// Gets a 64-bit integer primitive value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_int64(_:_:)
func xpc_dictionary_get_int64(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_int64(xdict, key)
	}


// Returns the connection that receives the dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_remote_connection(_:)
func xpc_dictionary_get_remote_connection(xdict unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_remote_connection(xdict)
	}


// Gets a C-string value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_string(_:_:)
func xpc_dictionary_get_string(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_string(xdict, key)
	}


// Gets a 64-bit unsigned integer primitive value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_uint64(_:_:)
func xpc_dictionary_get_uint64(xdict unsafe.Pointer, key unsafe.Pointer) uint64 {
	return _xpc_dictionary_get_uint64(xdict, key)
	}


// Gets a UUID value from a dictionary directly. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_uuid(_:_:)
func xpc_dictionary_get_uuid(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_uuid(xdict, key)
	}


// Returns the value for the specified key. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_get_value(_:_:)
func xpc_dictionary_get_value(xdict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _xpc_dictionary_get_value(xdict, key)
	}


// Inserts a Boolean primitive value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_bool(_:_:_:)
func xpc_dictionary_set_bool(xdict unsafe.Pointer, key unsafe.Pointer, value bool) {
	_xpc_dictionary_set_bool(xdict, key, value)
	}


// Inserts a connection into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_connection(_:_:_:)
func xpc_dictionary_set_connection(xdict unsafe.Pointer, key unsafe.Pointer, connection unsafe.Pointer) {
	_xpc_dictionary_set_connection(xdict, key, connection)
	}


// Inserts a raw data value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_data(_:_:_:_:)
func xpc_dictionary_set_data(xdict unsafe.Pointer, key unsafe.Pointer, bytes unsafe.Pointer, length uintptr) {
	_xpc_dictionary_set_data(xdict, key, bytes, length)
	}


// Inserts a date primitive value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_date(_:_:_:)
func xpc_dictionary_set_date(xdict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_xpc_dictionary_set_date(xdict, key, value)
	}


// Inserts a double-precision floating point primitive value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_double(_:_:_:)
func xpc_dictionary_set_double(xdict unsafe.Pointer, key unsafe.Pointer, value float64) {
	_xpc_dictionary_set_double(xdict, key, value)
	}


// Inserts a file descriptor into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_fd(_:_:_:)
func xpc_dictionary_set_fd(xdict unsafe.Pointer, key unsafe.Pointer, fd int) {
	_xpc_dictionary_set_fd(xdict, key, fd)
	}


// Inserts a 64-bit integer primitive value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_int64(_:_:_:)
func xpc_dictionary_set_int64(xdict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_xpc_dictionary_set_int64(xdict, key, value)
	}


// xpc_dictionary_set_mach_send is a XPC function. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_mach_send(_:_:_:)
func xpc_dictionary_set_mach_send(xdict unsafe.Pointer, key unsafe.Pointer, p unsafe.Pointer) {
	_xpc_dictionary_set_mach_send(xdict, key, p)
	}


// Inserts a C-string value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_string(_:_:_:)
func xpc_dictionary_set_string(xdict unsafe.Pointer, key unsafe.Pointer, string unsafe.Pointer) {
	_xpc_dictionary_set_string(xdict, key, string)
	}


// Inserts a 64-bit unsigned integer primitive value into a dictionary. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_uint64(_:_:_:)
func xpc_dictionary_set_uint64(xdict unsafe.Pointer, key unsafe.Pointer, value uint64) {
	_xpc_dictionary_set_uint64(xdict, key, value)
	}


// Inserts a UUID primitive value into an array. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_uuid(_:_:_:)
func xpc_dictionary_set_uuid(xdict unsafe.Pointer, key unsafe.Pointer, uuid unsafe.Pointer) {
	_xpc_dictionary_set_uuid(xdict, key, uuid)
	}


// Sets the value for the specified key to the specified object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_dictionary_set_value(_:_:_:)
func xpc_dictionary_set_value(xdict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_xpc_dictionary_set_value(xdict, key, value)
	}


// Creates an XPC double object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_double_create(_:)
func xpc_double_create(value float64) unsafe.Pointer {
	return _xpc_double_create(value)
	}


// Returns the underlying double-precision floating point value from an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_double_get_value(_:)
func xpc_double_get_value(xdouble unsafe.Pointer) float64 {
	return _xpc_double_get_value(xdouble)
	}


// Creates a new endpoint from a connection that is suitable for embedding into messages. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_endpoint_create(_:)
func xpc_endpoint_create(connection unsafe.Pointer) unsafe.Pointer {
	return _xpc_endpoint_create(connection)
	}


// Compares two objects for equality. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_equal(_:_:)
func xpc_equal(object1 unsafe.Pointer, object2 unsafe.Pointer) bool {
	return _xpc_equal(object1, object2)
	}


// Creates an XPC object that represents a POSIX file descriptor. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_fd_create(_:)
func xpc_fd_create(fd int) unsafe.Pointer {
	return _xpc_fd_create(fd)
	}


// Returns a file descriptor that is equivalent to the one that the specified file descriptor object boxes. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_fd_dup(_:)
func xpc_fd_dup(xfd unsafe.Pointer) int {
	return _xpc_fd_dup(xfd)
	}


// Returns the type of an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_get_type(_:)
func xpc_get_type(object unsafe.Pointer) unsafe.Pointer {
	return _xpc_get_type(object)
	}


// Calculates a hash value for the specified object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_hash(_:)
func xpc_hash(object unsafe.Pointer) uintptr {
	return _xpc_hash(object)
	}


// Creates an XPC signed integer object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_int64_create(_:)
func xpc_int64_create(value unsafe.Pointer) unsafe.Pointer {
	return _xpc_int64_create(value)
	}


// Returns the underlying signed 64-bit integer value from an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_int64_get_value(_:)
func xpc_int64_get_value(xint unsafe.Pointer) unsafe.Pointer {
	return _xpc_int64_get_value(xint)
	}


// Activates an inactive listener. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_activate
func xpc_listener_activate(listener unsafe.Pointer, error_out unsafe.Pointer) bool {
	return _xpc_listener_activate(listener, error_out)
	}


// Cancels a listener. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_cancel
func xpc_listener_cancel(listener unsafe.Pointer) {
	_xpc_listener_cancel(listener)
	}


// Copies the description string of a listener. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_copy_description
func xpc_listener_copy_description(listener unsafe.Pointer) unsafe.Pointer {
	return _xpc_listener_copy_description(listener)
	}


// Creates the server side of an XPC service using the specified service name. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_create
func xpc_listener_create(service unsafe.Pointer, target_queue unsafe.Pointer, flags unsafe.Pointer, incoming_session_handler unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_listener_create(service, target_queue, flags, incoming_session_handler, error_out)
	}


// Rejects an incoming peer session request. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_reject_peer
func xpc_listener_reject_peer(peer unsafe.Pointer, reason unsafe.Pointer) {
	_xpc_listener_reject_peer(peer, reason)
	}


// xpc_listener_set_peer_code_signing_requirement is a XPC function. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_set_peer_code_signing_requirement(_:_:)
func xpc_listener_set_peer_code_signing_requirement(listener unsafe.Pointer, requirement unsafe.Pointer) int {
	return _xpc_listener_set_peer_code_signing_requirement(listener, requirement)
	}


// xpc_listener_set_peer_requirement is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_listener_set_peer_requirement
func xpc_listener_set_peer_requirement(listener unsafe.Pointer, requirement unsafe.Pointer) {
	_xpc_listener_set_peer_requirement(listener, requirement)
	}


// Starts listening for incoming connections and processes them with the specified event handler. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_main(_:)
func xpc_main(handler unsafe.Pointer) {
	_xpc_main(handler)
	}


// Creates an XPC object that represents the null object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_null_create()
func xpc_null_create() unsafe.Pointer {
	return _xpc_null_create()
	}


// xpc_peer_requirement_create_entitlement_exists is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_create_entitlement_exists
func xpc_peer_requirement_create_entitlement_exists(entitlement unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_peer_requirement_create_entitlement_exists(entitlement, error_out)
	}


// xpc_peer_requirement_create_entitlement_matches_value is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_create_entitlement_matches_value
func xpc_peer_requirement_create_entitlement_matches_value(entitlement unsafe.Pointer, value unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_peer_requirement_create_entitlement_matches_value(entitlement, value, error_out)
	}


// xpc_peer_requirement_create_lwcr is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_create_lwcr
func xpc_peer_requirement_create_lwcr(lwcr unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_peer_requirement_create_lwcr(lwcr, error_out)
	}


// xpc_peer_requirement_create_platform_identity is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_create_platform_identity
func xpc_peer_requirement_create_platform_identity(signing_identifier unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_peer_requirement_create_platform_identity(signing_identifier, error_out)
	}


// xpc_peer_requirement_create_team_identity is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_create_team_identity
func xpc_peer_requirement_create_team_identity(signing_identifier unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_peer_requirement_create_team_identity(signing_identifier, error_out)
	}


// xpc_peer_requirement_match_received_message is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_peer_requirement_match_received_message
func xpc_peer_requirement_match_received_message(peer_requirement unsafe.Pointer, message unsafe.Pointer, error_out unsafe.Pointer) bool {
	return _xpc_peer_requirement_match_received_message(peer_requirement, message, error_out)
	}


// Decrements the reference count of an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_release
func xpc_release(object unsafe.Pointer) {
	_xpc_release(object)
	}


// Increments the reference count of an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_retain
func xpc_retain(object unsafe.Pointer) unsafe.Pointer {
	return _xpc_retain(object)
	}


// Returns a Boolean that indicates whether you can retry the operation that experienced an error. [Full Topic]
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_rich_error_can_retry(_:)
func xpc_rich_error_can_retry(error unsafe.Pointer) bool {
	return _xpc_rich_error_can_retry(error)
	}


// Copies the string description of an error. [Full Topic]
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_rich_error_copy_description(_:)
func xpc_rich_error_copy_description(error unsafe.Pointer) unsafe.Pointer {
	return _xpc_rich_error_copy_description(error)
	}


// Activates a session so you can send messages. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_activate
func xpc_session_activate(session unsafe.Pointer, error_out unsafe.Pointer) bool {
	return _xpc_session_activate(session, error_out)
	}


// Cancels a session, discarding any unsent messages. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_cancel
func xpc_session_cancel(session unsafe.Pointer) {
	_xpc_session_cancel(session)
	}


// Copies the description string of a session. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_copy_description
func xpc_session_copy_description(session unsafe.Pointer) unsafe.Pointer {
	return _xpc_session_copy_description(session)
	}


// Establishes a connection to a launch agent or launch daemon with the name you specify. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_create_mach_service
func xpc_session_create_mach_service(mach_service unsafe.Pointer, target_queue unsafe.Pointer, flags unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_session_create_mach_service(mach_service, target_queue, flags, error_out)
	}


// Establishes a connection to an XPC service with the name you specify. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_create_xpc_service
func xpc_session_create_xpc_service(name unsafe.Pointer, target_queue unsafe.Pointer, flags unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_session_create_xpc_service(name, target_queue, flags, error_out)
	}


// Sends a message over the session to the destination service. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_send_message
func xpc_session_send_message(session unsafe.Pointer, message unsafe.Pointer) unsafe.Pointer {
	return _xpc_session_send_message(session, message)
	}


// Sends a message asynchronously over the session to the destination service, calling a handler after receiving a reply. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_send_message_with_reply_async
func xpc_session_send_message_with_reply_async(session unsafe.Pointer, message unsafe.Pointer, reply_handler unsafe.Pointer) {
	_xpc_session_send_message_with_reply_async(session, message, reply_handler)
	}


// Sends a message over the session to the destination service, blocking the caller until receiving a reply. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_send_message_with_reply_sync
func xpc_session_send_message_with_reply_sync(session unsafe.Pointer, message unsafe.Pointer, error_out unsafe.Pointer) unsafe.Pointer {
	return _xpc_session_send_message_with_reply_sync(session, message, error_out)
	}


// Sets a handler the session calls when it’s canceled. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_set_cancel_handler
func xpc_session_set_cancel_handler(session unsafe.Pointer, cancel_handler unsafe.Pointer) {
	_xpc_session_set_cancel_handler(session, cancel_handler)
	}


// Sets a handler to receive incoming messages for a session. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_set_incoming_message_handler
func xpc_session_set_incoming_message_handler(session unsafe.Pointer, handler unsafe.Pointer) {
	_xpc_session_set_incoming_message_handler(session, handler)
	}


// xpc_session_set_peer_code_signing_requirement is a XPC function. [Full Topic]
//
// Added in macOS 14.4.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_set_peer_code_signing_requirement(_:_:)
func xpc_session_set_peer_code_signing_requirement(session unsafe.Pointer, requirement unsafe.Pointer) int {
	return _xpc_session_set_peer_code_signing_requirement(session, requirement)
	}


// xpc_session_set_peer_requirement is a XPC function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_set_peer_requirement
func xpc_session_set_peer_requirement(session unsafe.Pointer, requirement unsafe.Pointer) {
	_xpc_session_set_peer_requirement(session, requirement)
	}


// Sets the target dispatch queue on an inactive session for processing messages. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_session_set_target_queue
func xpc_session_set_target_queue(session unsafe.Pointer, target_queue unsafe.Pointer) {
	_xpc_session_set_target_queue(session, target_queue)
	}


// Sets the event handler to invoke when receiving streamed events. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_set_event_stream_handler(_:_:_:)
func xpc_set_event_stream_handler(stream unsafe.Pointer, targetq unsafe.Pointer, handler unsafe.Pointer) {
	_xpc_set_event_stream_handler(stream, targetq, handler)
	}


// Creates an XPC object that represents the specified shared memory region. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_shmem_create(_:_:)
func xpc_shmem_create(region unsafe.Pointer, length uintptr) unsafe.Pointer {
	return _xpc_shmem_create(region, length)
	}


// Maps the region that the XPC shared memory object boxes into the caller’s address space. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_shmem_map(_:_:)
func xpc_shmem_map(xshmem unsafe.Pointer, region unsafe.Pointer) uintptr {
	return _xpc_shmem_map(xshmem, region)
	}


// Creates an XPC object that represents a null-terminated C-string. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_string_create(_:)
func xpc_string_create(string unsafe.Pointer) unsafe.Pointer {
	return _xpc_string_create(string)
	}


// Creates an XPC object that represents a C-string that the specified format string and arguments generate. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_string_create_with_format
func xpc_string_create_with_format(fmt unsafe.Pointer) unsafe.Pointer {
	return _xpc_string_create_with_format(fmt)
	}


// Creates an XPC object that represents a C-string that the specified format string and argument list pointer generate. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_string_create_with_format_and_arguments(_:_:)
func xpc_string_create_with_format_and_arguments(fmt unsafe.Pointer, ap unsafe.Pointer) unsafe.Pointer {
	return _xpc_string_create_with_format_and_arguments(fmt, ap)
	}


// Returns the length of the underlying string. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_string_get_length(_:)
func xpc_string_get_length(xstring unsafe.Pointer) uintptr {
	return _xpc_string_get_length(xstring)
	}


// Returns a pointer to the internal storage of a string object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_string_get_string_ptr(_:)
func xpc_string_get_string_ptr(xstring unsafe.Pointer) unsafe.Pointer {
	return _xpc_string_get_string_ptr(xstring)
	}


// Informs the XPC runtime when a transaction begins, indicating that the service isn’t idle. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_transaction_begin()
func xpc_transaction_begin() {
	_xpc_transaction_begin()
	}


// Informs the XPC runtime when a transaction ends. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_transaction_end()
func xpc_transaction_end() {
	_xpc_transaction_end()
	}


// Returns a string that describes an XPC object type. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_type_get_name(_:)
func xpc_type_get_name(type_ unsafe.Pointer) unsafe.Pointer {
	return _xpc_type_get_name(type_)
	}


// Creates an XPC unsigned integer object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_uint64_create(_:)
func xpc_uint64_create(value uint64) unsafe.Pointer {
	return _xpc_uint64_create(value)
	}


// Returns the underlying unsigned 64-bit integer value from an object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_uint64_get_value(_:)
func xpc_uint64_get_value(xuint unsafe.Pointer) uint64 {
	return _xpc_uint64_get_value(xuint)
	}


// Creates an XPC object that represents a universally unique identifier (UUID). [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_uuid_create(_:)
func xpc_uuid_create(uuid unsafe.Pointer) unsafe.Pointer {
	return _xpc_uuid_create(uuid)
	}


// Copies the UUID that an XPC UUID object boxes into the specified UUID buffer. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: doc://com.apple.xpc/documentation/XPC/xpc_uuid_get_bytes(_:)
func xpc_uuid_get_bytes(xuuid unsafe.Pointer) unsafe.Pointer {
	return _xpc_uuid_get_bytes(xuuid)
	}



