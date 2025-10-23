// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// pasteboardProtocol is the pasteboard: protocol.
//
// Availability:
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to pasteboard:.
var pasteboardProtocol *objc.Protocol

func init() {
	pasteboardProtocol = objc.GetProtocol("pasteboard:")
}
