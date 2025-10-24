// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// activeConversationChangedProtocol is the activeConversationChanged: protocol.
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Use this protocol when registering custom classes that conform to activeConversationChanged:.
var activeConversationChangedProtocol *objc.Protocol

func init() {
	activeConversationChangedProtocol = objc.GetProtocol("activeConversationChanged:")
}
