// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"unsafe"
)

// PMEExtension is the MEExtension protocol interface.
//
// A type that provides objects for manipulating email messages, such as performing actions on messages or blocking content when users view messages.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEExtension
type PMEExtension interface {
	// Optional methods
	HandlerForComposeSession(session IMEComposeSession) unsafe.Pointer
	HasHandlerForComposeSession() bool
	HandlerForContentBlocker() unsafe.Pointer
	HasHandlerForContentBlocker() bool
	HandlerForMessageActions() unsafe.Pointer
	HasHandlerForMessageActions() bool
	HandlerForMessageSecurity() unsafe.Pointer
	HasHandlerForMessageSecurity() bool
}
