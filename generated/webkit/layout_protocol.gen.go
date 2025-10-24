// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// layoutProtocol is the layout protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to layout.
var layoutProtocol *objc.Protocol

func init() {
	layoutProtocol = objc.GetProtocol("layout")
}

