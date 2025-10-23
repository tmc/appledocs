// Code generated from Apple documentation for EndpointSecurity. DO NOT EDIT.

package endpointsecurity

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// EndpointSecurity Functions (39 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_es_clear_cache func(unsafe.Pointer) unsafe.Pointer
	_es_copy_message func(unsafe.Pointer) unsafe.Pointer
	_es_delete_client func(unsafe.Pointer) unsafe.Pointer
	_es_exec_arg func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_exec_arg_count func(unsafe.Pointer) unsafe.Pointer
	_es_exec_env func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_exec_env_count func(unsafe.Pointer) unsafe.Pointer
	_es_exec_fd func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_exec_fd_count func(unsafe.Pointer) unsafe.Pointer
	_es_free_message func(unsafe.Pointer) unsafe.Pointer
	_es_invert_muting func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_message_size func(unsafe.Pointer) unsafe.Pointer
	_es_mute_path func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_mute_path_events func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_mute_path_literal func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_mute_path_prefix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_mute_process func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_mute_process_events func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_muted_paths_events func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_muted_processes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_muted_processes_events func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_muting_inverted func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_new_client func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_release_message func(unsafe.Pointer) unsafe.Pointer
	_es_release_muted_paths func(unsafe.Pointer) unsafe.Pointer
	_es_release_muted_processes func(unsafe.Pointer) unsafe.Pointer
	_es_respond_auth_result func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_es_respond_flags_result func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_es_retain_message func(unsafe.Pointer) unsafe.Pointer
	_es_subscribe func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_subscriptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unmute_all_paths func(unsafe.Pointer) unsafe.Pointer
	_es_unmute_all_target_paths func(unsafe.Pointer) unsafe.Pointer
	_es_unmute_path func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unmute_path_events func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unmute_process func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unmute_process_events func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unsubscribe func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_es_unsubscribe_all func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_es_clear_cache, lib, "es_clear_cache")
	tryRegister(&_es_copy_message, lib, "es_copy_message")
	tryRegister(&_es_delete_client, lib, "es_delete_client")
	tryRegister(&_es_exec_arg, lib, "es_exec_arg")
	tryRegister(&_es_exec_arg_count, lib, "es_exec_arg_count")
	tryRegister(&_es_exec_env, lib, "es_exec_env")
	tryRegister(&_es_exec_env_count, lib, "es_exec_env_count")
	tryRegister(&_es_exec_fd, lib, "es_exec_fd")
	tryRegister(&_es_exec_fd_count, lib, "es_exec_fd_count")
	tryRegister(&_es_free_message, lib, "es_free_message")
	tryRegister(&_es_invert_muting, lib, "es_invert_muting")
	tryRegister(&_es_message_size, lib, "es_message_size")
	tryRegister(&_es_mute_path, lib, "es_mute_path")
	tryRegister(&_es_mute_path_events, lib, "es_mute_path_events")
	tryRegister(&_es_mute_path_literal, lib, "es_mute_path_literal")
	tryRegister(&_es_mute_path_prefix, lib, "es_mute_path_prefix")
	tryRegister(&_es_mute_process, lib, "es_mute_process")
	tryRegister(&_es_mute_process_events, lib, "es_mute_process_events")
	tryRegister(&_es_muted_paths_events, lib, "es_muted_paths_events")
	tryRegister(&_es_muted_processes, lib, "es_muted_processes")
	tryRegister(&_es_muted_processes_events, lib, "es_muted_processes_events")
	tryRegister(&_es_muting_inverted, lib, "es_muting_inverted")
	tryRegister(&_es_new_client, lib, "es_new_client")
	tryRegister(&_es_release_message, lib, "es_release_message")
	tryRegister(&_es_release_muted_paths, lib, "es_release_muted_paths")
	tryRegister(&_es_release_muted_processes, lib, "es_release_muted_processes")
	tryRegister(&_es_respond_auth_result, lib, "es_respond_auth_result")
	tryRegister(&_es_respond_flags_result, lib, "es_respond_flags_result")
	tryRegister(&_es_retain_message, lib, "es_retain_message")
	tryRegister(&_es_subscribe, lib, "es_subscribe")
	tryRegister(&_es_subscriptions, lib, "es_subscriptions")
	tryRegister(&_es_unmute_all_paths, lib, "es_unmute_all_paths")
	tryRegister(&_es_unmute_all_target_paths, lib, "es_unmute_all_target_paths")
	tryRegister(&_es_unmute_path, lib, "es_unmute_path")
	tryRegister(&_es_unmute_path_events, lib, "es_unmute_path_events")
	tryRegister(&_es_unmute_process, lib, "es_unmute_process")
	tryRegister(&_es_unmute_process_events, lib, "es_unmute_process_events")
	tryRegister(&_es_unsubscribe, lib, "es_unsubscribe")
	tryRegister(&_es_unsubscribe_all, lib, "es_unsubscribe_all")
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



// Clears all cached results for all clients.
//
// Added in macOS 10.15.

// Clears all cached results for all clients.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_clear_cache(_:)
func es_clear_cache(client unsafe.Pointer) unsafe.Pointer {
	return _es_clear_cache(client)
	}


// Copies a message, by allocating new memory.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.15.

// Copies a message, by allocating new memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_copy_message(_:)
func es_copy_message(msg unsafe.Pointer) unsafe.Pointer {
	return _es_copy_message(msg)
	}


// Destroys and disconnects a client instance from the Endpoint Security system.
//
// Added in macOS 10.15.

// Destroys and disconnects a client instance from the Endpoint Security system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_delete_client(_:)
func es_delete_client(client unsafe.Pointer) unsafe.Pointer {
	return _es_delete_client(client)
	}


// Gets the argument at the specified position from a process execution event.
//
// Added in macOS 10.15.

// Gets the argument at the specified position from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg(_:_:)
func es_exec_arg(event unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _es_exec_arg(event, index)
	}


// Gets the number of arguments from a process execution event.
//
// Added in macOS 10.15.

// Gets the number of arguments from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_arg_count(_:)
func es_exec_arg_count(event unsafe.Pointer) unsafe.Pointer {
	return _es_exec_arg_count(event)
	}


// Gets the environment variable at the specified position from a process execution event.
//
// Added in macOS 10.15.

// Gets the environment variable at the specified position from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env(_:_:)
func es_exec_env(event unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _es_exec_env(event, index)
	}


// Gets the number of environment variables from a process execution event.
//
// Added in macOS 10.15.

// Gets the number of environment variables from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_env_count(_:)
func es_exec_env_count(event unsafe.Pointer) unsafe.Pointer {
	return _es_exec_env_count(event)
	}


// Gets the file descriptor at the specified position from a process execution event.
//
// Added in macOS 11.0.

// Gets the file descriptor at the specified position from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd(_:_:)
func es_exec_fd(event unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _es_exec_fd(event, index)
	}


// Gets the number of file descriptors from a process execution event.
//
// Added in macOS 11.0.

// Gets the number of file descriptors from a process execution event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_exec_fd_count(_:)
func es_exec_fd_count(event unsafe.Pointer) unsafe.Pointer {
	return _es_exec_fd_count(event)
	}


// Frees the memory allocated for the given message.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.15.

// Frees the memory allocated for the given message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_free_message(_:)
func es_free_message(msg unsafe.Pointer) {
	_es_free_message(msg)
	}


// es_invert_muting is a EndpointSecurity function.
//
// Added in macOS 13.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_invert_muting(_:_:)
func es_invert_muting(client unsafe.Pointer, mute_type unsafe.Pointer) unsafe.Pointer {
	return _es_invert_muting(client, mute_type)
	}


// Calculates the size of a message structure.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.15.

// Calculates the size of a message structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_message_size(_:)
func es_message_size(msg unsafe.Pointer) unsafe.Pointer {
	return _es_message_size(msg)
	}


// Suppresses events from executables that match a given path.
//
// Added in macOS 12.0.

// Suppresses events from executables that match a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path(_:_:_:)
func es_mute_path(client unsafe.Pointer, path unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _es_mute_path(client, path, type_)
	}


// Suppresses a subset of events from executables that match a given path.
//
// Added in macOS 12.0.

// Suppresses a subset of events from executables that match a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_events(_:_:_:_:_:)
func es_mute_path_events(client unsafe.Pointer, path unsafe.Pointer, type_ unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_mute_path_events(client, path, type_, events, event_count)
	}


// Suppresses events from executables matching a path literal.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.15.

// Suppresses events from executables matching a path literal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_literal(_:_:)
func es_mute_path_literal(client unsafe.Pointer, path_literal unsafe.Pointer) unsafe.Pointer {
	return _es_mute_path_literal(client, path_literal)
	}


// Suppresses events from executables matching a path prefix.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.15.

// Suppresses events from executables matching a path prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_path_prefix(_:_:)
func es_mute_path_prefix(client unsafe.Pointer, path_prefix unsafe.Pointer) unsafe.Pointer {
	return _es_mute_path_prefix(client, path_prefix)
	}


// Suppresses events from a given process.
//
// Added in macOS 10.15.

// Suppresses events from a given process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process(_:_:)
func es_mute_process(client unsafe.Pointer, audit_token unsafe.Pointer) unsafe.Pointer {
	return _es_mute_process(client, audit_token)
	}


// Suppresses a subset of events from a given process.
//
// Added in macOS 12.0.

// Suppresses a subset of events from a given process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_mute_process_events(_:_:_:_:)
func es_mute_process_events(client unsafe.Pointer, audit_token unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_mute_process_events(client, audit_token, events, event_count)
	}


// Retrieve a list of all muted paths.
//
// Added in macOS 12.0.

// Retrieve a list of all muted paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_paths_events(_:_:)
func es_muted_paths_events(client unsafe.Pointer, muted_paths unsafe.Pointer) unsafe.Pointer {
	return _es_muted_paths_events(client, muted_paths)
	}


// Generates a list of muted processes.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.15.

// Generates a list of muted processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes(_:_:_:)
func es_muted_processes(client unsafe.Pointer, count unsafe.Pointer, audit_tokens unsafe.Pointer) unsafe.Pointer {
	return _es_muted_processes(client, count, audit_tokens)
	}


// Retrieve a list of all muted processes.
//
// Added in macOS 12.0.

// Retrieve a list of all muted processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muted_processes_events(_:_:)
func es_muted_processes_events(client unsafe.Pointer, muted_processes unsafe.Pointer) unsafe.Pointer {
	return _es_muted_processes_events(client, muted_processes)
	}


// es_muting_inverted is a EndpointSecurity function.
//
// Added in macOS 13.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_muting_inverted(_:_:)
func es_muting_inverted(client unsafe.Pointer, mute_type unsafe.Pointer) unsafe.Pointer {
	return _es_muting_inverted(client, mute_type)
	}


// Creates a new client instance and connects it to the Endpoint Security system.
//
// Added in macOS 10.15.

// Creates a new client instance and connects it to the Endpoint Security system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_new_client(_:_:)
func es_new_client(client unsafe.Pointer, handler unsafe.Pointer) unsafe.Pointer {
	return _es_new_client(client, handler)
	}


// Releases a previously-retained message.
//
// Added in macOS 11.0.

// Releases a previously-retained message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_release_message(_:)
func es_release_message(msg unsafe.Pointer) {
	_es_release_message(msg)
	}


// Frees resources associated with a set of previously-retrieved muted paths.
//
// Added in macOS 12.0.

// Frees resources associated with a set of previously-retrieved muted paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_paths(_:)
func es_release_muted_paths(muted_paths unsafe.Pointer) {
	_es_release_muted_paths(muted_paths)
	}


// Frees resources associated with a set of previously-retrieved muted processes.
//
// Added in macOS 12.0.

// Frees resources associated with a set of previously-retrieved muted processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_release_muted_processes(_:)
func es_release_muted_processes(muted_processes unsafe.Pointer) {
	_es_release_muted_processes(muted_processes)
	}


// Responds to an event that requires an authorization response.
//
// Added in macOS 10.15.

// Responds to an event that requires an authorization response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_respond_auth_result(_:_:_:_:)
func es_respond_auth_result(client unsafe.Pointer, message unsafe.Pointer, result unsafe.Pointer, cache bool) unsafe.Pointer {
	return _es_respond_auth_result(client, message, result, cache)
	}


// Responds to an event that requires authorization flags as a response.
//
// Added in macOS 10.15.

// Responds to an event that requires authorization flags as a response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_respond_flags_result(_:_:_:_:)
func es_respond_flags_result(client unsafe.Pointer, message unsafe.Pointer, authorized_flags unsafe.Pointer, cache bool) unsafe.Pointer {
	return _es_respond_flags_result(client, message, authorized_flags, cache)
	}


// Retains the given message, extending its lifetime until released.
//
// Added in macOS 11.0.

// Retains the given message, extending its lifetime until released.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_retain_message(_:)
func es_retain_message(msg unsafe.Pointer) {
	_es_retain_message(msg)
	}


// Subscribes a client to a set of events.
//
// Added in macOS 10.15.

// Subscribes a client to a set of events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_subscribe(_:_:_:)
func es_subscribe(client unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_subscribe(client, events, event_count)
	}


// Returns a list of the client’s subscriptions.
//
// Added in macOS 10.15.

// Returns a list of the client’s subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_subscriptions(_:_:_:)
func es_subscriptions(client unsafe.Pointer, count unsafe.Pointer, subscriptions unsafe.Pointer) unsafe.Pointer {
	return _es_subscriptions(client, count, subscriptions)
	}


// Restores event delivery from previously-muted paths.
//
// Added in macOS 10.15.

// Restores event delivery from previously-muted paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_paths(_:)
func es_unmute_all_paths(client unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_all_paths(client)
	}


// es_unmute_all_target_paths is a EndpointSecurity function.
//
// Added in macOS 13.0.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_all_target_paths(_:)
func es_unmute_all_target_paths(client unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_all_target_paths(client)
	}


// Restores event delivery from a previously-muted path.
//
// Added in macOS 12.0.

// Restores event delivery from a previously-muted path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path(_:_:_:)
func es_unmute_path(client unsafe.Pointer, path unsafe.Pointer, type_ unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_path(client, path, type_)
	}


// Restores event delivery of a subset of events from a previously-muted path.
//
// Added in macOS 12.0.

// Restores event delivery of a subset of events from a previously-muted path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_path_events(_:_:_:_:_:)
func es_unmute_path_events(client unsafe.Pointer, path unsafe.Pointer, type_ unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_path_events(client, path, type_, events, event_count)
	}


// Restores event delivery from a previously-muted process.
//
// Added in macOS 10.15.

// Restores event delivery from a previously-muted process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process(_:_:)
func es_unmute_process(client unsafe.Pointer, audit_token unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_process(client, audit_token)
	}


// Restores event delivery of a subset of events from a previously-muted process.
//
// Added in macOS 12.0.

// Restores event delivery of a subset of events from a previously-muted process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unmute_process_events(_:_:_:_:)
func es_unmute_process_events(client unsafe.Pointer, audit_token unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_unmute_process_events(client, audit_token, events, event_count)
	}


// Unsubscribes the provided client from a set of events.
//
// Added in macOS 10.15.

// Unsubscribes the provided client from a set of events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe(_:_:_:)
func es_unsubscribe(client unsafe.Pointer, events unsafe.Pointer, event_count unsafe.Pointer) unsafe.Pointer {
	return _es_unsubscribe(client, events, event_count)
	}


// Unsubscribes a client from all events.
//
// Added in macOS 10.15.

// Unsubscribes a client from all events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EndpointSecurity/es_unsubscribe_all(_:)
func es_unsubscribe_all(client unsafe.Pointer) unsafe.Pointer {
	return _es_unsubscribe_all(client)
	}




