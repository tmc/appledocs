// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// pasteProtocol is the paste: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to paste:.
var pasteProtocol *objc.Protocol

func init() {
	pasteProtocol = objc.GetProtocol("paste:")
}
