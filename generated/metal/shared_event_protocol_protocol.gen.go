// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

// PSharedEvent is the MTLSharedEvent protocol interface.
//
// An instance you use to synchronize access to Metal resources across multiple CPUs, GPUs, and processes.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.metal/documentation/Metal/MTLSharedEvent
type PSharedEvent interface {
	// Required methods
	NewSharedEventHandle() SharedEventHandle/* debug [protocol_interface/required_method]: NewSharedEventHandle */
	NotifyListenerAtValueBlock(listener IMTLSharedEventListener, value uint64, block SharedEventNotificationBlock /* not a class type */)/* debug [protocol_interface/required_method]: NotifyListenerAtValueBlock */
	WaitUntilSignaledValueTimeoutMS(value uint64, milliseconds uint64) bool/* debug [protocol_interface/required_method]: WaitUntilSignaledValueTimeoutMS */
}
