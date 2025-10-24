// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// attributedStringProtocol is the attributedString protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to attributedString.
var attributedStringProtocol *objc.Protocol

func init() {
	attributedStringProtocol = objc.GetProtocol("attributedString")
}
