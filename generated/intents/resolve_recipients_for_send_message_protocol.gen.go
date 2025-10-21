// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import "github.com/ebitengine/purego/objc"

// resolveRecipientsForSendMessageProtocol is the resolveRecipientsForSendMessage: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 10.0+ (Deprecated in 11.0)
//   - iPadOS 10.0+ (Deprecated in 11.0)
//   - macOS 12.0+
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 3.2+ (Deprecated in 4.0)
//
// Use this protocol when registering custom classes that conform to resolveRecipientsForSendMessage:.
var resolveRecipientsForSendMessageProtocol *objc.Protocol

func init() {
	resolveRecipientsForSendMessageProtocol = objc.GetProtocol("resolveRecipientsForSendMessage:")
}
