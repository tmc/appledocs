// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import "github.com/ebitengine/purego/objc"

// attachmentProtocol is the attachment protocol.
//
// Availability:
//   - macOS 10.0+
//
// Use this protocol when registering custom classes that conform to attachment.
var attachmentProtocol *objc.Protocol

func init() {
	attachmentProtocol = objc.GetProtocol("attachment")
}
