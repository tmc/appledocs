// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import "github.com/ebitengine/purego/objc"

// tokenDriverProtocol is the tokenDriver: protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 10.15+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// Use this protocol when registering custom classes that conform to tokenDriver:.
var tokenDriverProtocol *objc.Protocol

func init() {
	tokenDriverProtocol = objc.GetProtocol("tokenDriver:")
}

