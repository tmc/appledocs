// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// attachmentProtocol is the attachment: protocol.
//
// Availability:
//   - macOS 14.0+
//
// Use this protocol when registering custom classes that conform to attachment:.
var attachmentProtocol *objc.Protocol

func init() {
	attachmentProtocol = objc.GetProtocol("attachment:")
}

