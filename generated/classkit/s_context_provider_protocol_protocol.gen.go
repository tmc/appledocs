// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"unsafe"
)

// PSContextProvider is the CLSContextProvider protocol interface.
//
// An interface used to tell your ClassKit context provider app extension to update contexts.
//
// Availability:
//   - Mac Catalyst 12.2+
//   - iOS 12.2+
//   - iPadOS 12.2+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.classkit/documentation/ClassKit/CLSContextProvider
type PSContextProvider interface {
	// Required methods
	UpdateDescendantsOfContextCompletion(context ICLSContext, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: UpdateDescendantsOfContextCompletion */
}
