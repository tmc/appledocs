// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import "github.com/ebitengine/purego/objc"

// handleSendMessageProtocol is the handleSendMessage: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 12.0+
//   - visionOS 1.0+
//   - watchOS 3.2+
//
// Use this protocol when registering custom classes that conform to handleSendMessage:.
var handleSendMessageProtocol *objc.Protocol

func init() {
	handleSendMessageProtocol = objc.GetProtocol("handleSendMessage:")
}
