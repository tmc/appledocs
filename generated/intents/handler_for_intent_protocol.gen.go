// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import "github.com/ebitengine/purego/objc"

// handlerForIntentProtocol is the handlerForIntent: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 3.2+
//
// Use this protocol when registering custom classes that conform to handlerForIntent:.
var handlerForIntentProtocol *objc.Protocol

func init() {
	handlerForIntentProtocol = objc.GetProtocol("handlerForIntent:")
}
