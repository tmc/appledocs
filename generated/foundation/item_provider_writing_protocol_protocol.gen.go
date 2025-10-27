// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
)

// PItemProviderWriting is the NSItemProviderWriting protocol interface.
//
// The protocol for implementing a class to allow an item provider to retrieve data from an instance of the class.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSItemProviderWriting
type PItemProviderWriting interface {
	// Required methods
	LoadDataWithTypeIdentifierForItemProviderCompletionHandler(typeIdentifier IString, completionHandler unsafe.Pointer) IProgress
	// Optional methods
	ItemProviderVisibilityForRepresentationWithTypeIdentifier(typeIdentifier IString) ItemProviderRepresentationVisibility
	HasItemProviderVisibilityForRepresentationWithTypeIdentifier() bool
}
