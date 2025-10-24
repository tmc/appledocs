// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

// PContentKeyRecipient is the AVContentKeyRecipient protocol interface.
//
// A protocol for requiring decryption keys for media data.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.3+
//   - iPadOS 10.3+
//   - macOS 10.12.4+
//   - tvOS 10.2+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVContentKeyRecipient
type PContentKeyRecipient interface {
	// Optional methods
	ContentKeySessionDidProvideContentKey(contentKeySession IAVContentKeySession, contentKey IAVContentKey)
	HasContentKeySessionDidProvideContentKey() bool
}
