// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

// POSSystemExtensionsWorkspaceObserver is the OSSystemExtensionsWorkspaceObserver protocol interface.
//
// Availability:
//   - macOS 15.1+
//
// See: doc://com.apple.systemextensions/documentation/SystemExtensions/OSSystemExtensionsWorkspaceObserver
type POSSystemExtensionsWorkspaceObserver interface {
	// Optional methods
	SystemExtensionWillBecomeDisabled(systemExtensionInfo IOSSystemExtensionInfo)
	HasSystemExtensionWillBecomeDisabled() bool
	SystemExtensionWillBecomeEnabled(systemExtensionInfo IOSSystemExtensionInfo)
	HasSystemExtensionWillBecomeEnabled() bool
	SystemExtensionWillBecomeInactive(systemExtensionInfo IOSSystemExtensionInfo)
	HasSystemExtensionWillBecomeInactive() bool
}
