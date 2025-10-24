// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PXPCProxyCreating is the NSXPCProxyCreating protocol interface.
//
// Methods for creating new proxy objects.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSXPCProxyCreating
type PXPCProxyCreating interface {
	// Required methods
	RemoteObjectProxy() objc.ID
	RemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
	// Optional methods
	SynchronousRemoteObjectProxyWithErrorHandler(handler unsafe.Pointer) objc.ID
	HasSynchronousRemoteObjectProxyWithErrorHandler() bool
}
