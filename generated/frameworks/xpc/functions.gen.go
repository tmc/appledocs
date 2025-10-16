// Code generated from Apple documentation for XPC. DO NOT EDIT.

package xpc

// XPC Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (14 total):

// xpc_connection_set_peer_entitlement_exists_requirement(connection _, entitlement :  xpc_connection_t,  _, :  UnsafePointer< CChar>) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+

// xpc_connection_set_peer_entitlement_matches_value_requirement(connection _, entitlement :  xpc_connection_t,  _, value :  UnsafePointer< CChar>,  _, :  xpc_object_t) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+

// xpc_connection_set_peer_lightweight_code_requirement(connection _, lwcr :  xpc_connection_t,  _, :  xpc_object_t) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+

// xpc_connection_set_peer_platform_identity_requirement(connection _, signing_identifier :  xpc_connection_t,  _, :  UnsafePointer< CChar>?) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+

// xpc_connection_set_peer_requirement(connection xpc_connection_t, peer_requirement ,  xpc_peer_requirement_t, );) extern   void
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_connection_set_peer_team_identity_requirement(connection _, signing_identifier :  xpc_connection_t,  _, :  UnsafePointer< CChar>?) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 17.4+
//   - iOS 17.4+
//   - iPadOS 17.4+
//   - macOS 14.4+

// xpc_listener_set_peer_requirement(listener xpc_listener_t, requirement ,  xpc_peer_requirement_t, );) extern   void
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_create_entitlement_exists(entitlement const  char *, error_out ,  xpc_rich_error_t *, );) extern   xpc_peer_requirement_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_create_entitlement_matches_value(entitlement const  char *, value ,  xpc_object_t, error_out ,  xpc_rich_error_t *, );) extern   xpc_peer_requirement_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_create_lwcr(lwcr xpc_object_t, error_out ,  xpc_rich_error_t *, );) extern   xpc_peer_requirement_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_create_platform_identity(signing_identifier const  char *, error_out ,  xpc_rich_error_t *, );) extern   xpc_peer_requirement_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_create_team_identity(signing_identifier const  char *, error_out ,  xpc_rich_error_t *, );) extern   xpc_peer_requirement_t
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_peer_requirement_match_received_message(peer_requirement xpc_peer_requirement_t, message ,  xpc_object_t, error_out ,  xpc_rich_error_t *, );) extern   bool
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+

// xpc_session_set_peer_requirement(session xpc_session_t, requirement ,  xpc_peer_requirement_t, );) extern   void
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
