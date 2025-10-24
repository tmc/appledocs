// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// PDiscardableContent is the NSDiscardableContent protocol interface.
//
// You implement this protocol when a class’s objects have subcomponents that can be discarded when not being used, thereby giving an application a smaller memory footprint.
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
// See: doc://com.apple.foundation/documentation/Foundation/NSDiscardableContent
type PDiscardableContent interface {
	// Required methods
	BeginContentAccess() bool
	DiscardContentIfPossible()
	EndContentAccess()
	IsContentDiscarded() bool
}
