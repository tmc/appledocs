// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PLocking is the NSLocking protocol interface.
//
// The elementary methods adopted by classes that define lock objects.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSLocking
type PLocking interface {
	// Required methods
	Lock()/* debug [protocol_interface/required_method]: Lock */
	Unlock()/* debug [protocol_interface/required_method]: Unlock */
}
