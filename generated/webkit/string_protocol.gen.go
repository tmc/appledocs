// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// stringProtocol is the string protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to string.
var stringProtocol *objc.Protocol

func init() {
	stringProtocol = objc.GetProtocol("string")
}
